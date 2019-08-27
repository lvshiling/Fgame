package outlandboss

import (
	"fgame/fgame/core/runner"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	dummytemplate "fgame/fgame/game/dummy/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	outlandbosstemplate "fgame/fgame/game/outlandboss/template"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
	"sort"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type OutlandBossService interface {
	Start()
	//获取外域boss列表
	GetOutlandBossList() []scene.NPC
	//获取地图外域boss列表
	GetOutlandBossListGroupByMap(mapId int32) []scene.NPC
	//获取外域boss
	GetOutlandBoss(biologyId int32) scene.NPC
	//---------
	runner.Task
	//记录列表
	GetDropRecordsList() []*OutlandBossDropRecordsObject
	//添加记录
	AddDropRecords(plName string, biologyId int32, mapId int32, dropTime int64, dropTemplateList []*droptemplate.DropItemData)
	//获取记录
	GetDropRecordsByTime(time int64) (int32, *OutlandBossDropRecordsObject)
	//清空记录
	GMClearDropRecords()

	//筛选boss
	GetGuaiJiOutlandBossList(force int64) []scene.NPC
}

type outlandBossService struct {
	rwm sync.RWMutex
	//外域boss
	outlandBossList []scene.NPC
	//--------
	rwmLog sync.RWMutex
	//记录列表
	dropRecordsList []*OutlandBossDropRecordsObject
	//上次系统插入记录时间
	lastAddDummyDropRecordsTime int64

	//按战斗力排序
	sortOutlandBossList []scene.NPC
}

type sortOutlandBossList []scene.NPC

func (s sortOutlandBossList) Len() int {
	return len(s)
}

func (s sortOutlandBossList) Less(i, j int) bool {
	a := outlandbosstemplate.GetOutlandBossTemplateService().GetOutlandBossTemplate(int32(s[i].GetBiologyTemplate().Id))
	b := outlandbosstemplate.GetOutlandBossTemplateService().GetOutlandBossTemplate(int32(s[j].GetBiologyTemplate().Id))
	if a.RecForce == b.RecForce {
		return a.Id < b.Id
	}
	return a.RecForce < b.RecForce
}

func (s sortOutlandBossList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *outlandBossService) init() (err error) {
	// entityList, err := dao.GetOutlandBossDao().GetOutlandBossDropRecordsList(global.GetGame().GetServerIndex())
	// if err != nil {
	// 	return
	// }

	// for _, entity := range entityList {
	// 	logObj := NewOutlandBossDropRecordsObject()
	// 	logObj.FromEntity(entity)
	// 	s.dropRecordsList = append(s.dropRecordsList, logObj)
	// }
	return
}

func (s *outlandBossService) GetOutlandBossList() []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.outlandBossList
}

func (s *outlandBossService) GetOutlandBossListGroupByMap(mapId int32) []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	var bossArr []scene.NPC
	for _, boss := range s.outlandBossList {
		if boss.GetScene().MapId() == mapId {
			bossArr = append(bossArr, boss)
		}
	}

	return bossArr
}

func (s *outlandBossService) GetOutlandBoss(biologyId int32) scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.getBoss(biologyId)
}

func (s *outlandBossService) Start() {
	mapIdList := outlandbosstemplate.GetOutlandBossTemplateService().GetMapIdList()
	for _, mapId := range mapIdList {
		sc := scene.GetSceneService().GetBossSceneByMapId(mapId)
		if sc == nil {
			continue
		}
		//TODO:xzk:修改优化
		bossList := sc.GetNPCS(scenetypes.BiologyScriptTypeOutlandBoss)
		for _, boss := range bossList {
			s.outlandBossList = append(s.outlandBossList, boss)
			s.sortOutlandBossList = append(s.sortOutlandBossList, boss)
		}
	}
	sort.Sort(sortOutlandBossList(s.sortOutlandBossList))
	return
}

func (s *outlandBossService) getBoss(biologyId int32) (n scene.NPC) {
	for _, boss := range s.outlandBossList {
		bossBiologyId := int32(boss.GetBiologyTemplate().TemplateId())
		if bossBiologyId == biologyId {
			return boss
		}
	}
	return nil
}

//TODO 修改为n秒加一次
//心跳
func (s *outlandBossService) Heartbeat() {

	err := s.addDummyDropRecords()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("outlandboss:系统生成假记录,错误")
		return
	}
}

//生成系统假记录
func (s *outlandBossService) addDummyDropRecords() (err error) {
	now := global.GetGame().GetTimeService().Now()
	lastTime := s.lastAddDummyDropRecordsTime
	diffTime := now - lastTime
	randTime := s.getRandomDropRecordsTime()
	if diffTime < randTime {
		return
	}

	name := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	template := outlandbosstemplate.GetOutlandBossTemplateService().GetOutlandbossTemplateRandom(0)
	dropDataList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(template.GetBiologyTemplate().GetDropIdList())

	s.AddDropRecords(name, template.BiologyId, template.MapId, now, dropDataList)

	s.lastAddDummyDropRecordsTime = now
	return
}

//是否记录掉落物品
func (s *outlandBossService) IsDropRecordItem(dropDataList []*droptemplate.DropItemData) bool {
	for _, itemData := range dropDataList {
		itemId := itemData.GetItemId()

		itemTemplate := item.GetItemService().GetItem(int(itemId))
		quality := itemtypes.ItemQualityType(itemTemplate.Quality)
		if quality >= itemtypes.ItemQualityTypePurple {
			return true
		}
	}
	return false
}

func (s *outlandBossService) GetDropRecordsList() []*OutlandBossDropRecordsObject {
	s.rwmLog.RLock()
	defer s.rwmLog.RUnlock()
	return s.dropRecordsList
}

func (s *outlandBossService) AddDropRecords(plName string, biologyId int32, mapId int32, dropTime int64, dropTemplateList []*droptemplate.DropItemData) {
	if s.IsDropRecordItem(dropTemplateList) {
		s.addDropRecordsStart(plName, biologyId, mapId, dropTime, dropTemplateList)
	}
}

func (s *outlandBossService) addDropRecordsStart(plName string, biologyId int32, mapId int32, dropTime int64, dropTemplateList []*droptemplate.DropItemData) {
	s.rwmLog.Lock()
	defer s.rwmLog.Unlock()

	s.appendDropRecords(plName, biologyId, mapId, dropTime, dropTemplateList)
}

func (s *outlandBossService) GetDropRecordsByTime(time int64) (int32, *OutlandBossDropRecordsObject) {
	s.rwmLog.RLock()
	defer s.rwmLog.RUnlock()

	if time == 0 {
		return -1, nil
	}

	endIndex := int32(0)
	var endLog *OutlandBossDropRecordsObject
	for index, log := range s.dropRecordsList {
		if time == log.dropTime {
			endIndex = int32(index)
			endLog = log
		}
	}

	return endIndex, endLog
}

func (s *outlandBossService) GMClearDropRecords() {
	s.rwmLog.Lock()
	defer s.rwmLog.Unlock()

	var empty []*OutlandBossDropRecordsObject
	s.dropRecordsList = empty
}

func (s *outlandBossService) appendDropRecords(plName string, biologyId int32, mapId int32, dropTime int64, dropTemplateList []*droptemplate.DropItemData) {
	obj := s.createDropRecordsObj(plName, biologyId, mapId, dropTime, dropTemplateList)
	s.dropRecordsList = append(s.dropRecordsList, obj)
}

func (s *outlandBossService) createDropRecordsObj(plName string, biologyId int32, mapId int32, dropTime int64, dropTemplateList []*droptemplate.DropItemData) *OutlandBossDropRecordsObject {
	now := global.GetGame().GetTimeService().Now()
	maxLen := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeOutlandBossDropRecordsCount)
	var obj *OutlandBossDropRecordsObject
	if len(s.dropRecordsList) >= int(maxLen) {
		obj = s.dropRecordsList[0]
		s.dropRecordsList = s.dropRecordsList[1:]
	} else {
		obj = NewOutlandBossDropRecordsObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.createTime = now
	}

	obj.killerName = plName
	obj.biologyId = biologyId
	obj.mapId = mapId
	obj.dropTime = dropTime
	obj.ItemMap = make(map[int32]int32)
	for _, dropItem := range dropTemplateList {
		itemId := dropItem.ItemId
		num := dropItem.Num
		obj.ItemMap[itemId] += num
	}
	// obj.SetModified()

	return obj
}

//系统假记录生成间隔
func (s *outlandBossService) getRandomDropRecordsTime() int64 {
	min := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeOutlandBossDropRecordsAddTimeMin))
	max := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeOutlandBossDropRecordsAddTimeMax))
	randTime := int64(mathutils.RandomRange(min, max))
	return randTime
}

func (s *outlandBossService) GetGuaiJiOutlandBossList(force int64) []scene.NPC {
	for index, boss := range s.sortOutlandBossList {
		outlandbossTemplate := outlandbosstemplate.GetOutlandBossTemplateService().GetOutlandBossTemplate(int32(boss.GetBiologyTemplate().Id))
		if int64(outlandbossTemplate.RecForce) > force {
			return s.sortOutlandBossList[:index]
		}
	}
	return s.sortOutlandBossList
}

var (
	once sync.Once
	ws   *outlandBossService
)

func Init() (err error) {
	once.Do(func() {
		ws = &outlandBossService{}
		err = ws.init()
	})
	return err
}

func GetOutlandBossService() OutlandBossService {
	return ws
}
