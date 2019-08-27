package godsiege

import (
	coretypes "fgame/fgame/core/types"
	activitytemplate "fgame/fgame/game/activity/template"
	"fgame/fgame/game/global"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"fgame/fgame/game/scene/scene"
	"sync"
)

type GodSiegeService interface {
	Heartbeat()
	//创建神兽攻城数据
	CreateGodSiegeScene(godType godsiegetypes.GodSiegeType, mapId int32, endTime int64) scene.Scene
	//获取神兽攻城场景
	GetGodSiegeScene(godType godsiegetypes.GodSiegeType) (s scene.Scene)
	//神兽攻城活动结束
	GodSiegeSceneFinish(godType godsiegetypes.GodSiegeType)
	//参加神兽攻城
	Attend(godType godsiegetypes.GodSiegeType, playerId int64) (lineUpPos int32, isLineUp bool, flag bool)
	//获取复活点复活位置
	GetRebornPos(godType godsiegetypes.GodSiegeType, playerId int64) (pos coretypes.Position, flag bool)
	//玩家是否有排队
	GetHasLineUp(godType godsiegetypes.GodSiegeType, playerId int64) (lineUpPos int32, isLineUp bool, flag bool)
	//玩家取消排队
	CancleLineUp(godType godsiegetypes.GodSiegeType, playerId int64) (flag bool)
	//同步场景人数
	SyncSceneNum(godType godsiegetypes.GodSiegeType, scenePlayerNum int32)
	//获取第一个排队玩家
	RemoveFirstLineUpPlayer(godType godsiegetypes.GodSiegeType, scenePlayerNum int32)
	//获取当前还在排队的
	GetAllLineUpList(godType godsiegetypes.GodSiegeType) (lineUpList []int64)
	//活动是否开始
	IsGodSiegeActivityTime(godType godsiegetypes.GodSiegeType) (flag bool)
}

type godSiegeService struct {
	rwm                sync.RWMutex
	godSiegeDataList   []*godSiegeData
	godSiegeServerType godsiegetypes.GodSiegeServerType
}

func (s *godSiegeService) init(godSiegeServerType godsiegetypes.GodSiegeServerType) (err error) {
	s.godSiegeServerType = godSiegeServerType

	for _, godSiegeData := range s.godSiegeDataList {
		godSiegeData.init()
	}
	return nil
}

func (s *godSiegeService) CreateGodSiegeScene(godType godsiegetypes.GodSiegeType, mapId int32, endTime int64) scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		godSiegeData = createGodSiegeData(mapId, godType)
		s.godSiegeDataList = append(s.godSiegeDataList, godSiegeData)
	}

	se := godSiegeData.getGodSiegeScene()
	if se == nil {
		se = godSiegeData.createGodSiegeScene(mapId, endTime, godType)
	}

	return se
}

func (s *godSiegeService) getGodSiegeData(godType godsiegetypes.GodSiegeType) *godSiegeData {
	for _, godSiegeData := range s.godSiegeDataList {
		if godSiegeData.godType == godType {
			return godSiegeData
		}
	}
	return nil
}

func (s *godSiegeService) GodSiegeSceneFinish(godType godsiegetypes.GodSiegeType) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	godSiegeData.init()
}

func (s *godSiegeService) Attend(godType godsiegetypes.GodSiegeType, playerId int64) (lineUpPos int32, lineUpFlag bool, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	flag = true
	lineUpPos, lineUpFlag = godSiegeData.attend(playerId)
	return
}

func (s *godSiegeService) GetHasLineUp(godType godsiegetypes.GodSiegeType, playerId int64) (lineUpPos int32, isLineUp bool, flag bool) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	flag = true
	lineUpPos, isLineUp = godSiegeData.getHasLineUp(playerId)
	return
}

func (s *godSiegeService) GetGodSiegeScene(godType godsiegetypes.GodSiegeType) (se scene.Scene) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}

	se = godSiegeData.getGodSiegeScene()
	return
}

func (s *godSiegeService) GetRebornPos(godType godsiegetypes.GodSiegeType, playerId int64) (pos coretypes.Position, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	pos, flag = godSiegeData.getRebornPos(playerId)
	return
}

func (s *godSiegeService) CancleLineUp(godType godsiegetypes.GodSiegeType, playerId int64) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	flag = godSiegeData.cancleLineUp(playerId)
	return
}

func (s *godSiegeService) SyncSceneNum(godType godsiegetypes.GodSiegeType, scenePlayerNum int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	godSiegeData.syncSceneNum(scenePlayerNum)
}

func (s *godSiegeService) GetAllLineUpList(godType godsiegetypes.GodSiegeType) (lineUpList []int64) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	return godSiegeData.getAllLineUpList()
}

func (s *godSiegeService) IsGodSiegeActivityTime(godType godsiegetypes.GodSiegeType) (flag bool) {
	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	flag = godSiegeData.isGodSiegeActivityTime()
	return
}

func (s *godSiegeService) RemoveFirstLineUpPlayer(godType godsiegetypes.GodSiegeType, scenePlayerNum int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData == nil {
		return
	}
	godSiegeData.removeFirstLineUpPlayer(scenePlayerNum)
	return
}

func (s *godSiegeService) checkGodSiegeActivityExist(godType godsiegetypes.GodSiegeType) (flag bool) {
	godSiegeData := s.getGodSiegeData(godType)
	if godSiegeData != nil {
		flag = true
	}
	return
}

func (s *godSiegeService) checkGodSiegeActivity() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for _, godType := range s.godSiegeServerType.GetGodSiegeTypeList() {
		flag := s.checkGodSiegeActivityExist(godType)
		if flag {
			continue
		}

		activityType, ok := godType.GetActivityType()
		if !ok {
			continue
		}
		activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activityType)
		if activityTemplate == nil {
			continue
		}
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
		if err != nil {
			return err
		}
		if activityTimeTemplate == nil {
			continue
		}
		endTime, err := activityTimeTemplate.GetEndTime(now)
		if err != nil {
			return err
		}
		s.CreateGodSiegeScene(godType, activityTemplate.Mapid, endTime)
	}

	return nil
}

func (s *godSiegeService) Heartbeat() {
	s.checkGodSiegeActivity()
}

var (
	once sync.Once
	cs   *godSiegeService
)

func Init(godSiegeServerType godsiegetypes.GodSiegeServerType) (err error) {
	once.Do(func() {
		cs = &godSiegeService{}
		err = cs.init(godSiegeServerType)
	})
	return err
}

func GetGodSiegeService() GodSiegeService {
	return cs
}
