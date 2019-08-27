package tulong

import (
	"context"
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/core/heartbeat"
	tulongclient "fgame/fgame/cross/tulong/client"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/center/center"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	tulongeventtypes "fgame/fgame/game/tulong/event/types"
	"fmt"
	"sync"
)

type TuLongService interface {
	Heartbeat()
	SyncRemoteRankListTask() error
	CheckTuLongActivityStartTask()
	Star() (err error)

	//获取排行榜
	GetRankList(servrId int32, allianceId int64) (dataList []*TuLongRankData, pos int32)
}

type tuLongService struct {
	rwm          sync.RWMutex
	tulongClient tulongclient.TuLongClient
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner

	//屠龙排行榜数据
	tuLongRankList []*TuLongRankData
	//活动开始重置排行榜标识
	resetFlag bool
}

func (s *tuLongService) init() (err error) {
	s.tuLongRankList = make([]*TuLongRankData, 0, 8)

	s.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	s.heartbeatRunner.AddTask(CreateSyncTask(s))
	s.heartbeatRunner.AddTask(CreateTuLongStartTask(s))

	err = s.resetClient()
	if err != nil {
		return
	} 

	return
}

func (s *tuLongService) resetClient() (err error) {
	conn := center.GetCenterService().GetCross(centertypes.GameServerTypeRegion)
	if conn == nil {
		return fmt.Errorf("tulong:跨服连接不存在")
	}

	//TODO 修改可能连接变化了
	s.tulongClient = tulongclient.NewTuLongClient(conn)
	return
}

func (s *tuLongService) GetRankList(servrId int32, allianceId int64) (dataList []*TuLongRankData, pos int32) {
	pos = 0
	for index, rankObj := range s.tuLongRankList {
		if rankObj.serverId == servrId && rankObj.allianceId == allianceId {
			pos = int32(index + 1)
		}
	}
	dataList = s.tuLongRankList
	return
}

func (s *tuLongService) Star() (err error) {
	err = s.SyncRemoteRankListTask()
	if err != nil {
		return
	}
	return
}

func (s *tuLongService) ifOnTuLongActivity() (isActiveTime bool) {
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeCoressTuLong)
	activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
	if err != nil {
		return
	}
	if activityTimeTemplate == nil {
		return
	}

	isActiveTime = true
	return
}

//活动开始
func (s *tuLongService) CheckTuLongActivityStartTask() {
	if !s.ifOnTuLongActivity() {
		s.resetFlag = false
		return
	}

	if !s.resetFlag {
		//清空跨服排行榜
		s.resetFlag = true
		s.tuLongRankList = make([]*TuLongRankData, 0, 8)

		gameevent.Emit(tulongeventtypes.EventTypeTuLongActivityStart, nil, nil)
	}
}

//定时同步排行榜列表
func (s *tuLongService) SyncRemoteRankListTask() (err error) {
	if s.tulongClient == nil {
		err = s.resetClient()
		if err != nil {
			return
		}
	}

	//TODO 超时
	ctx := context.TODO()
	resp, err := s.tulongClient.GetTuLongRankList(ctx)
	if err != nil {
		return
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.tuLongRankList = convertFromRankInfoList(resp.RankInfoList)
	return nil
}

func (s *tuLongService) Heartbeat() {
	s.heartbeatRunner.Heartbeat()
}

var (
	once sync.Once
	cs   *tuLongService
)

func Init() (err error) {
	once.Do(func() {
		cs = &tuLongService{}
		err = cs.init()
	})
	return err
}

func GetTuLongService() TuLongService {
	return cs
}
