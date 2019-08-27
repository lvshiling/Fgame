package lianyu

import (
	coretypes "fgame/fgame/core/types"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	lianyuscene "fgame/fgame/game/lianyu/scene"
	lianyutemplate "fgame/fgame/game/lianyu/template"
	lianyutypes "fgame/fgame/game/lianyu/types"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/scene/scene"
	"sync"
)

type LianYuService interface {
	Heartbeat()
	//获无间炼狱取场景
	GetLianYuScene() (s scene.Scene)
	//创建无间炼狱场景
	CreateLianYuScene(mapId int32, endTime int64) scene.Scene
	//无间炼狱活动结束
	LianYuSceneFinish()
	//参加无间炼狱
	Attend(playerId int64) (lineUpPos int32, flag bool)
	//获取复活点复活位置
	GetRebornPos(playerId int64) (pos coretypes.Position, flag bool)
	//玩家是否有排队
	GetHasLineUp(playerId int64) (lineUpPos int32, flag bool)
	//玩家取消排队
	CancleLineUp(playerId int64) (flag bool)
	//同步场景人数
	SyncSceneNum(scenePlayerNum int32)
	//获取第一个排队玩家
	RemoveFirstLineUpPlayer(scenePlayerNum int32)
	//获取当前还在排队的
	GetAllLineUpList() (lineUpList []int64)
}

type lianYuService struct {
	rwm             sync.RWMutex
	lianYuSceneData lianyuscene.LianYuSceneData
	//参加记录
	attendMap map[int64]struct{}
	//排队记录
	lineList []int64
	//出生统计
	bornMap map[lianyutypes.LianYuPosType]int32
	//场景人数
	num int32
	//总人数
	totalNum int32
	//活动类型
	activityType activitytypes.ActivityType
}

func (s *lianYuService) init(activityType activitytypes.ActivityType) (err error) {
	s.attendMap = make(map[int64]struct{})
	s.lineList = make([]int64, 0, 8)
	s.bornMap = make(map[lianyutypes.LianYuPosType]int32)
	s.num = 0
	s.totalNum = 0
	s.lianYuSceneData = nil
	s.activityType = activityType
	return nil
}

func (s *lianYuService) clearData() (err error) {
	s.attendMap = make(map[int64]struct{})
	s.lineList = make([]int64, 0, 8)
	s.bornMap = make(map[lianyutypes.LianYuPosType]int32)
	s.num = 0
	s.totalNum = 0
	s.lianYuSceneData = nil
	return nil
}

func (s *lianYuService) CreateLianYuScene(mapId int32, endTime int64) (se scene.Scene) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if s.lianYuSceneData != nil {
		se = s.lianYuSceneData.GetScene()
		return
	}

	consTemp := lianyutemplate.GetLianYuTemplateService().GetConstantTemplate(s.activityType)
	if consTemp == nil {
		return
	}

	s.totalNum = consTemp.PlayerLimitCount
	s.lineList = make([]int64, 0, 8)
	s.lianYuSceneData = lianyuscene.CreateLianYuSceneData(s.activityType)
	se = lianyuscene.CreateLianYuScene(mapId, endTime, s.lianYuSceneData)
	return
}

func (s *lianYuService) LianYuSceneFinish() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.clearData()
}

func (s *lianYuService) Attend(playerId int64) (lineUpPos int32, lineUpFlag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	lineLen := int32(len(s.lineList))
	if s.num >= s.totalNum {
		lineUpFlag = true
		s.lineList = append(s.lineList, playerId)
		lineUpPos = lineLen
		return
	}
	s.attendMap[playerId] = struct{}{}
	s.num++
	return 0, false
}

func (s *lianYuService) getHasLineUp(playerId int64) (lineUpPos int32, flag bool) {
	for index, curPlayerId := range s.lineList {
		if playerId == curPlayerId {
			lineUpPos = int32(index)
			flag = true
			break
		}
	}
	return
}

func (s *lianYuService) GetHasLineUp(playerId int64) (lineUpPos int32, flag bool) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	lineUpPos, flag = s.getHasLineUp(playerId)
	return
}

func (s *lianYuService) GetLianYuScene() (se scene.Scene) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if s.lianYuSceneData == nil {
		return
	}
	return s.lianYuSceneData.GetScene()
}

func (s *lianYuService) GetRebornPos(playerId int64) (pos coretypes.Position, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	_, exist := s.attendMap[playerId]
	if !exist {
		return
	}

	initPos := lianyutypes.LianYuBornTypePlayer
	if s.activityType == activitytypes.ActivityTypeLianYu {
		initPos = lianyutypes.LianYuBornTypeCrossPlayer
	}

	posTemplate := lianyutemplate.GetLianYuTemplateService().GetBornPosTemplate(initPos)
	if posTemplate != nil {
		pos = posTemplate.GetPos()
		flag = true
	}
	return
}

func (s *lianYuService) CancleLineUp(playerId int64) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	index, flag := s.getHasLineUp(playerId)
	if !flag {
		return
	}
	if index == 0 && len(s.lineList) <= 1 {
		s.lineList = make([]int64, 0, 8)
		return true
	}
	s.lineList = append(s.lineList[:index], s.lineList[index+1:]...)
	gameevent.Emit(lianyueventtypes.EventTypeLianYuCancleLineUp, s.lineList, int32(index))
	flag = true
	return
}

func (s *lianYuService) SyncSceneNum(scenePlayerNum int32) {
	if scenePlayerNum < 0 || scenePlayerNum > s.totalNum {
		return
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.num = scenePlayerNum
}

func (s *lianYuService) GetAllLineUpList() (lineUpList []int64) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.lineList
}

func (s *lianYuService) RemoveFirstLineUpPlayer(scenePlayerNum int32) {
	if scenePlayerNum < 0 || scenePlayerNum > s.totalNum {
		return
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()

	playerId := int64(0)
	s.num = scenePlayerNum
	lineLen := int32(len(s.lineList))
	if s.num < s.totalNum && lineLen != 0 {
		playerId = s.lineList[0]
		if lineLen == 1 {
			s.lineList = make([]int64, 0, 8)
		} else {
			s.lineList = s.lineList[1:]
		}
	}

	if playerId != 0 {
		s.attendMap[playerId] = struct{}{}
		gameevent.Emit(lianyueventtypes.EventTypeLianYuPlayerLineUpFinish, s.lineList, playerId)
	}
	return
}

func (s *lianYuService) checkLianYuActivity(activityType activitytypes.ActivityType) (err error) {
	if s.lianYuSceneData != nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activityType)
	if activityTemplate == nil {
		return
	}

	openTime := int64(0)
	mergeTime := int64(0)
	if activityType == activitytypes.ActivityTypeLocalLianYu {
		openTime = global.GetGame().GetServerTime()
		mergeTime = merge.GetMergeService().GetMergeTime()
	}
	activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
	if err != nil {
		return
	}
	if activityTimeTemplate == nil {
		return
	}
	endTime, err := activityTimeTemplate.GetEndTime(now)
	if err != nil {
		return
	}
	s.CreateLianYuScene(activityTemplate.Mapid, endTime)
	return nil
}

func (s *lianYuService) Heartbeat() {
	s.checkLianYuActivity(s.activityType)
}

var (
	once sync.Once
	cs   *lianYuService
)

func Init(activityType activitytypes.ActivityType) (err error) {
	once.Do(func() {
		cs = &lianYuService{}
		err = cs.init(activityType)
	})
	return err
}

func GetLianYuService() LianYuService {
	return cs
}
