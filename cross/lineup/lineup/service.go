package lineup

import (
	lineupeventtypes "fgame/fgame/cross/lineup/event/types"
	crosstypes "fgame/fgame/game/cross/types"
	gameevent "fgame/fgame/game/event"
	"sync"
)

type LineupService interface {
	//参加活动
	Attend(crossType crosstypes.CrossType, sceneId, playerId int64) (lineUpPos int32)
	//玩家是否有排队
	GetHasLineUp(crossType crosstypes.CrossType, sceneId, playerId int64) (lineUpPos int32, flag bool)
	//玩家取消排队
	CancleLineUp(crossType crosstypes.CrossType, sceneId, playerId int64) (flag bool)
	//获取第一个排队玩家
	RemoveFirstLineUpPlayer(crossType crosstypes.CrossType, sceneId int64)
	//获取当前还在排队的
	GetAllLineUpList(crossType crosstypes.CrossType, sceneId int64) (lineUpList []int64)
	//清空排队
	ClearAllLineupList(crossType crosstypes.CrossType, sceneId int64)
}

//排队数据
type LineupData struct {
	//排队记录
	lineListMap map[int64][]int64
	//类型
	crossType crosstypes.CrossType
}

func CreateLineupData(crossType crosstypes.CrossType) *LineupData {
	d := &LineupData{}
	d.crossType = crossType
	d.lineListMap = map[int64][]int64{}
	return d
}

func (d *LineupData) GetLineupList(sceneId int64) []int64 {
	return d.lineListMap[sceneId]
}

func (d *LineupData) GetCrossType() crosstypes.CrossType {
	return d.crossType
}

func (d *LineupData) getHasLineUp(sceneId int64, playerId int64) (lineUpPos int32, flag bool) {
	lineList := d.lineListMap[sceneId]
	for index, curPlayerId := range lineList {
		if playerId == curPlayerId {
			lineUpPos = int32(index)
			flag = true
			break
		}
	}
	return
}

type lineupService struct {
	rwm       sync.RWMutex
	lineupMap map[crosstypes.CrossType]*LineupData
}

func (s *lineupService) init() (err error) {
	s.lineupMap = make(map[crosstypes.CrossType]*LineupData)
	return nil
}

func (s *lineupService) Attend(crossType crosstypes.CrossType, sceneId, playerId int64) (lineUpPos int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	lineData, ok := s.lineupMap[crossType]
	if !ok {
		lineData = CreateLineupData(crossType)
		s.lineupMap[crossType] = lineData
	}

	// 排队
	lineList := lineData.lineListMap[sceneId]
	lineLen := int32(len(lineList))
	lineData.lineListMap[sceneId] = append(lineData.lineListMap[sceneId], playerId)
	lineUpPos = lineLen
	return
}

func (s *lineupService) GetHasLineUp(crossType crosstypes.CrossType, sceneId, playerId int64) (lineUpPos int32, flag bool) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	ld := s.getLineupData(crossType)
	if ld == nil {
		return
	}

	lineUpPos, flag = ld.getHasLineUp(sceneId, playerId)
	return
}

func (s *lineupService) CancleLineUp(crossType crosstypes.CrossType, sceneId, playerId int64) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	ld := s.getLineupData(crossType)
	if ld == nil {
		return
	}

	index, flag := ld.getHasLineUp(sceneId, playerId)
	if !flag {
		return
	}

	lineList := ld.lineListMap[sceneId]
	if index == 0 && len(lineList) <= 1 {
		ld.lineListMap[sceneId] = make([]int64, 0, 8)
		return true
	}

	//通知后面的玩家
	lineList = append(lineList[:index], lineList[index+1:]...)
	ld.lineListMap[sceneId] = lineList

	eventData := lineupeventtypes.CreateCancleLineUpEventData(int32(ld.crossType), int32(index))
	gameevent.Emit(lineupeventtypes.EventTypeLineupCancleLineUp, lineList, eventData)
	flag = true
	return
}

func (s *lineupService) GetAllLineUpList(crossType crosstypes.CrossType, sceneId int64) (lineUpList []int64) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	ld := s.getLineupData(crossType)
	if ld == nil {
		return
	}

	return ld.lineListMap[sceneId]
}

func (s *lineupService) ClearAllLineupList(crossType crosstypes.CrossType, sceneId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	ld := s.getLineupData(crossType)
	if ld == nil {
		return
	}

	ld.lineListMap[sceneId] = make([]int64, 0, 8)
}

func (s *lineupService) RemoveFirstLineUpPlayer(crossType crosstypes.CrossType, sceneId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	ld, ok := s.lineupMap[crossType]
	if !ok {
		return
	}

	// 排队长度
	lineList := ld.lineListMap[sceneId]
	lineLen := int32(len(lineList))
	if lineLen == 0 {
		return
	}

	playerId := lineList[0]
	if lineLen == 1 {
		ld.lineListMap[sceneId] = make([]int64, 0, 8)
	} else {
		ld.lineListMap[sceneId] = lineList[1:]
	}

	eventData := lineupeventtypes.CreatePlayerLineUpFinishEventData(int32(ld.crossType), playerId)
	gameevent.Emit(lineupeventtypes.EventTypeLineupPlayerLineUpFinish, lineList[1:], eventData)
	return
}

func (s *lineupService) getLineupData(crossType crosstypes.CrossType) *LineupData {
	return s.lineupMap[crossType]
}

var (
	once sync.Once
	cs   *lineupService
)

func Init() (err error) {
	once.Do(func() {
		cs = &lineupService{}
		err = cs.init()
	})
	return err
}

func GetLineupService() LineupService {
	return cs
}
