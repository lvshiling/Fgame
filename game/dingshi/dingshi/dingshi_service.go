package dingshi

import (
	"fgame/fgame/game/dingshi/dao"
	dingshieventtypes "fgame/fgame/game/dingshi/event/types"
	dingshitemplate "fgame/fgame/game/dingshi/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/idutil"
	"sort"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type DingShiService interface {
	Start()
	//获取藏经阁boss列表
	GetDingShiBossList() []scene.NPC
	//获取地图藏经阁boss列表
	GetDingShiBossListGroupByMap(mapId int32) []scene.NPC
	//获取藏经阁boss
	GetDingShiBoss(biologyId int32) scene.NPC
	//筛选boss
	GetGuaiJiDingShiBossList(force int64) []scene.NPC

	GetBossDeadTime(mapId int32, bossId int32) int64
	BossDead(mapId int32, bossId int32)
}

type sortDingShiBossList []scene.NPC

func (s sortDingShiBossList) Len() int {
	return len(s)
}

func (s sortDingShiBossList) Less(i, j int) bool {
	a := dingshitemplate.GetDingShiTemplateService().GetDingShiBossTemplateByBiologyId(int32(s[i].GetBiologyTemplate().Id))
	b := dingshitemplate.GetDingShiTemplateService().GetDingShiBossTemplateByBiologyId(int32(s[j].GetBiologyTemplate().Id))
	if a.RecForce == b.RecForce {
		return a.Id < b.Id
	}
	return a.RecForce < b.RecForce
}

func (s sortDingShiBossList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type dingShiService struct {
	rwm sync.RWMutex
	//藏经阁boss
	dingShiBossList []scene.NPC
	//按战斗力排序
	sortDingShiBossList []scene.NPC
	//已经入库的定时boss列表
	dingShiBossMapOfMap map[int32]map[int32]*DingShiBossObject
}

func (s *dingShiService) init() (err error) {
	err = s.loadDingShiBossObject()
	if err != nil {
		return
	}
	return
}

func (s *dingShiService) loadDingShiBossObject() (err error) {
	s.dingShiBossMapOfMap = make(map[int32]map[int32]*DingShiBossObject)
	serverId := global.GetGame().GetServerIndex()
	bossEntityList, err := dao.GetDingShiDao().GetDingShiBossEntityList(serverId)
	if err != nil {
		return
	}
	for _, bossEntity := range bossEntityList {
		bossObj := newDingShiBossObject()
		err = bossObj.FromEntity(bossEntity)
		if err != nil {
			return
		}
		s.addBoss(bossObj)
	}
	return nil
}

func (s *dingShiService) addBoss(obj *DingShiBossObject) {
	dingShiBossMap, ok := s.dingShiBossMapOfMap[obj.mapId]
	if !ok {
		dingShiBossMap = make(map[int32]*DingShiBossObject)
		s.dingShiBossMapOfMap[obj.mapId] = dingShiBossMap
	}
	dingShiBossMap[obj.bossId] = obj
}

func (s *dingShiService) BossDead(mapId int32, bossId int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	var bossObj *DingShiBossObject
	now := global.GetGame().GetTimeService().Now()
	dingShiBossMap, ok := s.dingShiBossMapOfMap[mapId]
	if !ok {
		goto Create
	}

	bossObj, ok = dingShiBossMap[bossId]
	if ok {
		bossObj.lastKillTime = now
		bossObj.updateTime = now
		bossObj.SetModified()
		return
	}
Create:
	bossObj = newDingShiBossObject()
	bossObj.id, _ = idutil.GetId()
	bossObj.mapId = mapId
	bossObj.serverId = global.GetGame().GetServerIndex()
	bossObj.bossId = bossId
	bossObj.lastKillTime = now
	bossObj.createTime = now
	bossObj.SetModified()
	s.addBoss(bossObj)
	return
}

func (s *dingShiService) GetBossDeadTime(mapId int32, bossId int32) int64 {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	var bossObj *DingShiBossObject
	dingShiBossMap, ok := s.dingShiBossMapOfMap[mapId]
	if !ok {
		goto Create
	}

	bossObj, ok = dingShiBossMap[bossId]
	if !ok {
		goto Create
	}
	return bossObj.lastKillTime
Create:
	now := global.GetGame().GetTimeService().Now()
	bossObj = newDingShiBossObject()
	bossObj.id, _ = idutil.GetId()
	bossObj.mapId = mapId
	bossObj.serverId = global.GetGame().GetServerIndex()
	bossObj.bossId = bossId
	bossObj.lastKillTime = now
	bossObj.createTime = now
	bossObj.SetModified()
	s.addBoss(bossObj)
	return bossObj.lastKillTime
}

func (s *dingShiService) GetDingShiBossList() []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.dingShiBossList
}

func (s *dingShiService) GetDingShiBossListGroupByMap(mapId int32) []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	var bossArr []scene.NPC
	for _, boss := range s.dingShiBossList {
		if boss.GetScene().MapId() == mapId {
			bossArr = append(bossArr, boss)
		}
	}
	return bossArr
}

func (s *dingShiService) GetDingShiBoss(biologyId int32) scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.getBoss(biologyId)
}

func (s *dingShiService) Start() {
	//添加定时boss
	dingShiMap := dingshitemplate.GetDingShiTemplateService().GetDingShiMap()
	for _, dingShiObj := range dingShiMap {
		sb := scene.GetSceneService().GetBossSceneByMapId(dingShiObj.MapId)

		if sb == nil {
			continue
		}
		sceneTemplate := scenetemplate.GetSceneTemplateService().GetDingShiSceneTemplate(dingShiObj.MapId, dingShiObj.BiologyId)
		deadTime := s.GetBossDeadTime(dingShiObj.MapId, dingShiObj.BiologyId)
		n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, int32(sceneTemplate.IndexID), sceneTemplate.GetBiology(), sceneTemplate.GetPos(), sceneTemplate.Angle, deadTime)
		if n == nil {
			log.WithFields(
				log.Fields{
					"mapId":       sb.MapId(),
					"biologyType": sceneTemplate.GetBiology().GetBiologyScriptType(),
				}).Warnln("创建npc,不存在")
			continue
		}

		s.dingShiBossList = append(s.dingShiBossList, n)
		s.sortDingShiBossList = append(s.sortDingShiBossList, n)
		gameevent.Emit(dingshieventtypes.EventTypeDingShiInit, sb, n)
	}

	sort.Sort(sortDingShiBossList(s.sortDingShiBossList))
	return
}

func (s *dingShiService) getBoss(biologyId int32) (n scene.NPC) {
	for _, boss := range s.dingShiBossList {
		bossBiologyId := int32(boss.GetBiologyTemplate().TemplateId())
		if bossBiologyId == biologyId {
			return boss
		}
	}
	return nil
}

func (s *dingShiService) GetGuaiJiDingShiBossList(force int64) []scene.NPC {
	for index, boss := range s.sortDingShiBossList {
		template := dingshitemplate.GetDingShiTemplateService().GetDingShiBossTemplateByBiologyId(int32(boss.GetBiologyTemplate().Id))
		if int64(template.RecForce) > force {
			return s.sortDingShiBossList[:index]
		}
	}
	return s.sortDingShiBossList
}

var (
	once sync.Once
	ds   *dingShiService
)

func Init() (err error) {
	once.Do(func() {
		ds = &dingShiService{}
		err = ds.init()
	})
	return err
}

func GetDingShiService() DingShiService {
	return ds
}
