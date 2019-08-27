package arenaboss

import (
	"fgame/fgame/cross/arenaboss/dao"
	arenabosseventtypes "fgame/fgame/cross/arenaboss/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	sharebosstemplate "fgame/fgame/game/shareboss/template"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/pkg/idutil"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type ArenaBossService interface {
	Start()
	GetBossDeadTime(mapId int32, bossId int32) int64
	BossDead(mapId int32, bossId int32)
	Stop()

	GetBossList() []scene.NPC
}

type arenaBossService struct {
	rwm sync.RWMutex

	//已经入库的定时boss列表
	arenaBossMapOfMap map[int32]map[int32]*ArenaBossObject
	arenaBossList     []scene.NPC
}

func (s *arenaBossService) init() (err error) {
	err = s.loadArenaBossObject()
	if err != nil {
		return
	}
	return
}

func (s *arenaBossService) loadArenaBossObject() (err error) {
	s.arenaBossMapOfMap = make(map[int32]map[int32]*ArenaBossObject)
	serverId := global.GetGame().GetServerIndex()
	platform := global.GetGame().GetPlatform()
	bossEntityList, err := dao.GetArenaBossDao().GetArenaBossEntityList(platform, serverId)
	if err != nil {
		return
	}
	for _, bossEntity := range bossEntityList {
		bossObj := newArenaBossObject()
		err = bossObj.FromEntity(bossEntity)
		if err != nil {
			return
		}
		s.addBoss(bossObj)
	}
	return nil
}

func (s *arenaBossService) addBoss(obj *ArenaBossObject) {
	arenaBossMap, ok := s.arenaBossMapOfMap[obj.mapId]
	if !ok {
		arenaBossMap = make(map[int32]*ArenaBossObject)
		s.arenaBossMapOfMap[obj.mapId] = arenaBossMap
	}
	arenaBossMap[obj.bossId] = obj
}

func (s *arenaBossService) BossDead(mapId int32, bossId int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	var bossObj *ArenaBossObject
	now := global.GetGame().GetTimeService().Now()
	arenaBossMap, ok := s.arenaBossMapOfMap[mapId]
	if !ok {
		goto Create
	}

	bossObj, ok = arenaBossMap[bossId]
	if ok {
		bossObj.lastKillTime = now
		bossObj.updateTime = now
		bossObj.SetModified()
		return
	}
Create:
	bossObj = newArenaBossObject()
	bossObj.id, _ = idutil.GetId()
	bossObj.mapId = mapId
	bossObj.platform = global.GetGame().GetPlatform()
	bossObj.serverId = global.GetGame().GetServerIndex()
	bossObj.bossId = bossId
	bossObj.lastKillTime = now
	bossObj.createTime = now
	bossObj.SetModified()
	s.addBoss(bossObj)
	return
}

func (s *arenaBossService) GetBossDeadTime(mapId int32, bossId int32) int64 {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	var bossObj *ArenaBossObject
	arenaBossMap, ok := s.arenaBossMapOfMap[mapId]
	if !ok {
		goto Create
	}

	bossObj, ok = arenaBossMap[bossId]
	if !ok {
		goto Create
	}
	return bossObj.lastKillTime
Create:
	now := global.GetGame().GetTimeService().Now()
	bossObj = newArenaBossObject()
	bossObj.id, _ = idutil.GetId()
	bossObj.mapId = mapId
	bossObj.platform = global.GetGame().GetPlatform()
	bossObj.serverId = global.GetGame().GetServerIndex()
	bossObj.bossId = bossId
	bossObj.lastKillTime = now
	bossObj.createTime = now
	bossObj.SetModified()
	s.addBoss(bossObj)
	return bossObj.lastKillTime
}

func (s *arenaBossService) GetBossList() []scene.NPC {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	return s.arenaBossList
}

func (s *arenaBossService) Start() {
	//添加定时boss
	arenaBossMap := sharebosstemplate.GetShareBossTemplateService().GetShareBossTemplateMap(worldbosstypes.BossTypeArena)
	for _, arenaBossObj := range arenaBossMap {
		sb := scene.GetSceneService().GetBossSceneByMapId(arenaBossObj.GetMapId())
		if sb == nil {
			continue
		}
		sceneTemplate := scenetemplate.GetSceneTemplateService().GetDingShiSceneTemplate(arenaBossObj.GetMapId(), arenaBossObj.GetBiologyId())
		if sceneTemplate == nil {
			npcList := sb.GetNPCListByBiology(arenaBossObj.GetBiologyId())
			if len(npcList) == 0 {
				continue
			}
			s.arenaBossList = append(s.arenaBossList, npcList[0])
			continue
		}
		deadTime := s.GetBossDeadTime(arenaBossObj.GetMapId(), arenaBossObj.GetBiologyId())
		n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, int32(sceneTemplate.IndexID), sceneTemplate.GetBiology(), sceneTemplate.GetPos(), sceneTemplate.Angle, deadTime)
		if n == nil {
			log.WithFields(
				log.Fields{
					"mapId":       sb.MapId(),
					"biologyType": sceneTemplate.GetBiology().GetBiologyScriptType(),
				}).Warnln("创建npc,不存在")
			continue
		}

		s.arenaBossList = append(s.arenaBossList, n)

		gameevent.Emit(arenabosseventtypes.EventTypeArenaBossInit, sb, n)
	}

	return
}

func (s *arenaBossService) Stop() {
	return
}

var (
	once sync.Once
	as   *arenaBossService
)

func Init() (err error) {
	once.Do(func() {
		as = &arenaBossService{}
		err = as.init()
	})
	return err
}

func GetArenaBossService() ArenaBossService {
	return as
}
