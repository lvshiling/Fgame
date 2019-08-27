package arena

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/arena/dao"
	"fgame/fgame/game/arena/types"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"sort"
	"sync"
)

type ArenaService interface {
	Heartbeat()
	Star() (err error)
	CheckRefreshWeekRank()

	//获取我的排名
	GetMyRank(rankType types.RankTimeType, servrId int32, playerId int64) (pos, val int32, rankTime int64)
	//获取排名
	GetRankList(rankType types.RankTimeType, page int32) (dataList []*ArenaRankObject, rankTime int64)
	//获取排行时间
	GetRankTime(rankType types.RankTimeType) (rankTime int64)

	// //获取排行榜积分
	// GetJiFenNum(allianceId int64) int32
	// 添加获胜次数
	UpdateWinCount(win bool, playerId int64, playerName string)
}

type arenaService struct {
	rwm      sync.RWMutex
	hbRunner heartbeat.HeartbeatTaskRunner
	// -------------------

	//本周排行榜数据
	thisArenaRankList []*ArenaRankObject
	//上周排行榜数据
	lastArenaRankList []*ArenaRankObject
	//排行榜刷新时间对象
	arenaRankTimeObj *ArenaRankTimeObject
}

const (
	pageLimit = 50 //分页大小
)

func (s *arenaService) init() (err error) {
	s.thisArenaRankList = make([]*ArenaRankObject, 0, pageLimit)
	s.lastArenaRankList = make([]*ArenaRankObject, 0, pageLimit)

	s.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	s.hbRunner.AddTask(CreateArenaRankTask(s))

	err = s.loadRankData()
	if err != nil {
		return
	}
	err = s.loadRankTime()
	if err != nil {
		return
	}

	return nil
}

func (s *arenaService) loadRankTime() (err error) {
	//排行刷新时间
	serverId := global.GetGame().GetServerIndex()
	arenaRankTimeEntity, err := dao.GetArenaDao().GetArenaRankTimeEntity(serverId)
	if err != nil {
		return
	}
	if arenaRankTimeEntity == nil {
		s.initArenaRankTimeObject()
	} else {
		s.arenaRankTimeObj = NewArenaRankTimeObject()
		s.arenaRankTimeObj.FromEntity(arenaRankTimeEntity)
	}

	return
}

func (s *arenaService) loadRankData() (err error) {
	serverId := global.GetGame().GetServerIndex()
	//排行榜列表
	arenaRankList, err := dao.GetArenaDao().GetArenaRankList(serverId)
	if err != nil {
		return
	}
	for _, arenaRank := range arenaRankList {
		obj := NewArenaRankObject()
		obj.FromEntity(arenaRank)
		s.thisArenaRankList = append(s.thisArenaRankList, obj)
		s.lastArenaRankList = append(s.lastArenaRankList, obj)
	}

	sort.Sort(sort.Reverse(ThisArenaRankObjectList(s.thisArenaRankList)))
	sort.Sort(sort.Reverse(LastArenaRankObjectList(s.lastArenaRankList)))
	return
}

//第一次初始化
func (s *arenaService) initArenaRankTimeObject() {
	po := NewArenaRankTimeObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	po.Id = id
	po.ServerId = global.GetGame().GetServerIndex()
	po.ThisTime = 0
	po.LastTime = 0
	po.CreateTime = now
	s.arenaRankTimeObj = po
	po.SetModified()
}

//更新连胜排行榜
func (s *arenaService) UpdateWinCount(win bool, playerId int64, playerName string) {
	serverId := global.GetGame().GetServerIndex()
	rankObj, flag := s.getPlayerRankData(playerId)
	if !flag {
		if !win {
			return
		}
		rankObj = initArenaRankObject(serverId, playerId, playerName)
		s.thisArenaRankList = append(s.thisArenaRankList, rankObj)
		s.lastArenaRankList = append(s.lastArenaRankList, rankObj)
	} else {
		if win {
			curWin := rankObj.CurWinCount + 1
			rankObj.CurWinCount = curWin
			if curWin > rankObj.WinCount {
				rankObj.WinCount = curWin
			}
		} else {
			rankObj.CurWinCount = 0
		}
	}

	now := global.GetGame().GetTimeService().Now()
	rankObj.UpdateTime = now
	rankObj.SetModified()

	sort.Sort(sort.Reverse(ThisArenaRankObjectList(s.thisArenaRankList)))
}

func (s *arenaService) getPlayerRankData(playerId int64) (rankObj *ArenaRankObject, flag bool) {
	if len(s.thisArenaRankList) == 0 {
		return
	}
	for _, rankObj := range s.thisArenaRankList {
		if rankObj.PlayerId == playerId {
			return rankObj, true
		}
	}
	return
}

func (s *arenaService) GetMyRank(rankType types.RankTimeType, servrId int32, playerId int64) (pos, val int32, rankTime int64) {
	pos = 0
	if rankType == types.RankTimeTypeLast {
		for index, rankObj := range s.lastArenaRankList {
			if rankObj.PlayerId == playerId {
				pos = int32(index + 1)
				val = rankObj.LastWinCount
			}
		}

		rankTime = s.arenaRankTimeObj.LastTime
	} else {
		for index, rankObj := range s.thisArenaRankList {
			if rankObj.PlayerId == playerId {
				pos = int32(index + 1)
				val = rankObj.WinCount
			}
		}
		rankTime = s.arenaRankTimeObj.ThisTime
	}

	return
}

func (s *arenaService) GetRankList(rankType types.RankTimeType, page int32) (dataList []*ArenaRankObject, rankTime int64) {
	var rankList []*ArenaRankObject
	if rankType == types.RankTimeTypeLast {
		rankList = s.lastArenaRankList
		rankTime = s.arenaRankTimeObj.LastTime
	} else {
		rankList = s.thisArenaRankList
		rankTime = s.arenaRankTimeObj.ThisTime
	}
	pageIndex := page * pageLimit
	len := len(rankList)
	if pageIndex >= int32(len) {
		return
	}
	addLen := pageIndex + pageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	dataList = rankList[pageIndex:addLen]
	return
}

func (s *arenaService) GetRankTime(rankType types.RankTimeType) (rankTime int64) {
	if rankType == types.RankTimeTypeLast {
		rankTime = s.arenaRankTimeObj.LastTime
	} else {
		rankTime = s.arenaRankTimeObj.ThisTime
	}
	return
}

func (s *arenaService) Star() (err error) {
	return
}

func (s *arenaService) Heartbeat() {
	s.hbRunner.Heartbeat()
}

func (s *arenaService) CheckRefreshWeekRank() {
	now := global.GetGame().GetTimeService().Now()
	if s.arenaRankTimeObj.ThisTime == 0 {
		timeStamp, _ := timeutils.MondayFivePointTime(now)
		s.arenaRankTimeObj.ThisTime = timeStamp
		s.arenaRankTimeObj.UpdateTime = now
		s.arenaRankTimeObj.SetModified()
		return
	}

	diffTime := now - s.arenaRankTimeObj.ThisTime
	if diffTime < 0 {
		diffTime *= -1 //绝对值
	}
	if diffTime < int64(7*timeutils.DAY) {
		return
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.refreshWeekRank(now)
}

func (s *arenaService) refreshWeekRank(now int64) {
	timeStamp, _ := timeutils.MondayFivePointTime(now)
	s.arenaRankTimeObj.LastTime = timeStamp - int64(7*timeutils.DAY)
	s.arenaRankTimeObj.ThisTime = timeStamp
	s.arenaRankTimeObj.UpdateTime = now
	s.arenaRankTimeObj.SetModified()

	for _, rankObj := range s.thisArenaRankList {
		rankObj.LastWinCount = rankObj.WinCount
		rankObj.WinCount = 0
		rankObj.LastTime = now
		rankObj.UpdateTime = now
		rankObj.SetModified()
	}

	if len(s.thisArenaRankList) != 0 {
		sort.Sort(sort.Reverse(ThisArenaRankObjectList(s.thisArenaRankList)))
	}
	if len(s.lastArenaRankList) != 0 {
		sort.Sort(sort.Reverse(LastArenaRankObjectList(s.lastArenaRankList)))
	}
}

var (
	once sync.Once
	cs   *arenaService
)

func Init() (err error) {
	once.Do(func() {
		cs = &arenaService{}
		err = cs.init()
	})
	return err
}

func GetArenaService() ArenaService {
	return cs
}
