package activity

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/activity/dao"
	activityeventtypes "fgame/fgame/game/activity/event/types"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sync"
)

type ActivityService interface {
	Heartbeat()
	//活动开始提醒
	CheckStartActivity() (err error)
	//活动结束提醒
	CheckEndActivity() (err error)
	//活动结束记录
	AddEndRecord(activityType activitytypes.ActivityType, endTime int64)
	//活动是否已结束(活动提前结束处理)
	IsActivityEnd(activityType activitytypes.ActivityType, endTime int64) bool
	//GM：清除结束记录
	GMClearEndRecord()
}

type activityService struct {
	rwm            sync.RWMutex
	runner         heartbeat.HeartbeatTaskRunner
	startNoticeMap map[activitytypes.ActivityType]int64
	endNoticeMap   map[activitytypes.ActivityType]int64
	endRecordMap   map[activitytypes.ActivityType]*ActivityEndRecordObject
}

func (as *activityService) init() (err error) {
	as.startNoticeMap = make(map[activitytypes.ActivityType]int64)
	as.endNoticeMap = make(map[activitytypes.ActivityType]int64)

	//添加定时任务
	as.runner = heartbeat.NewHeartbeatTaskRunner()
	as.runner.AddTask(CreateActivityNoticeTask(as))

	// 初始化活动结束记录
	err = as.initActivityEndRecord()
	if err != nil {
		return
	}

	return nil
}

func (as *activityService) initActivityEndRecord() error {

	as.endRecordMap = make(map[activitytypes.ActivityType]*ActivityEndRecordObject)
	entityList, err := dao.GetActivityDao().GetActivityEndRecordList(global.GetGame().GetServerIndex())
	if err != nil {
		return err
	}

	for _, entity := range entityList {
		obj := newActivityEndRecordObject()
		obj.FromEntity(entity)
		as.endRecordMap[obj.activityType] = obj
	}

	return nil
}

func (as *activityService) Heartbeat() {
	as.runner.Heartbeat()
}

func (as *activityService) Start() {
	return
}

func (as *activityService) CheckStartActivity() (err error) {
	as.rwm.Lock()
	defer as.rwm.Unlock()

	activityMap := activitytemplate.GetActivityTemplateService().GetActiveAll()
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	for _, temp := range activityMap {
		timeTemp := temp.GetOnDateTimeTemplate(now, openTime, mergeTime)
		if timeTemp == nil {
			continue
		}

		if timeTemp.BeginNoticeTime == 0 {
			continue
		}

		begin, err := timeTemp.GetBeginTime(now)
		if err != nil {
			err = fmt.Errorf("activity:活动开始提醒，获取活动开始时间错误")
			return err
		}

		if now > begin {
			continue
		}

		afterTime := now + timeTemp.BeginNoticeTime
		if afterTime < begin {
			continue
		}

		preBegin, ok := as.startNoticeMap[temp.GetActivityType()]
		if ok {
			flag, err := timeutils.IsSameDay(preBegin, begin)
			if err != nil {
				err = fmt.Errorf("activity:活动开始提醒，时间对比解析错误")
				return err
			}
			if flag {
				continue
			}
		}

		gameevent.Emit(activityeventtypes.EventTypeActivityNoticeStart, nil, timeTemp)

		as.startNoticeMap[temp.GetActivityType()] = begin
	}
	return
}

func (as *activityService) CheckEndActivity() (err error) {
	as.rwm.Lock()
	defer as.rwm.Unlock()

	activityMap := activitytemplate.GetActivityTemplateService().GetActiveAll()
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	for _, temp := range activityMap {
		timeTemp := temp.GetOnDateTimeTemplate(now, openTime, mergeTime)
		if timeTemp == nil {
			continue
		}

		if timeTemp.EndNoticeTime == 0 {
			continue
		}

		end, err := timeTemp.GetEndTime(now)
		if err != nil {
			err = fmt.Errorf("activity:活动结束提醒，获取活动结束时间错误")
			return err
		}

		if now > end {
			continue
		}

		afterTime := now + timeTemp.EndNoticeTime
		if afterTime < end {
			continue
		}

		preEnd, ok := as.endNoticeMap[temp.GetActivityType()]
		if ok {
			flag, err := timeutils.IsSameDay(preEnd, end)
			if err != nil {
				err = fmt.Errorf("activity:活动结束提醒，活动时间对比解析错误")
				return err
			}
			if flag {
				continue
			}
		}

		gameevent.Emit(activityeventtypes.EventTypeActivityNoticeEnd, nil, timeTemp)

		as.endNoticeMap[temp.GetActivityType()] = end
	}
	return
}

func (as *activityService) AddEndRecord(activityType activitytypes.ActivityType, endTime int64) {
	as.rwm.Lock()
	defer as.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	recordObj, ok := as.endRecordMap[activityType]
	if !ok {
		recordObj = newActivityEndRecordObject()
		id, _ := idutil.GetId()
		recordObj.id = id
		recordObj.serverId = global.GetGame().GetServerIndex()
		recordObj.activityType = activityType
		recordObj.endTime = endTime
		recordObj.createTime = now
		recordObj.SetModified()
		as.endRecordMap[activityType] = recordObj
		return
	}

	if recordObj.endTime == endTime {
		return
	}

	recordObj.endTime = endTime
	recordObj.updateTime = now
	recordObj.SetModified()
}

func (as *activityService) IsActivityEnd(activityType activitytypes.ActivityType, endTime int64) (flag bool) {
	as.rwm.RLock()
	defer as.rwm.RUnlock()

	recordObj, ok := as.endRecordMap[activityType]
	if !ok {
		return
	}

	if recordObj.endTime != endTime {
		return
	}

	flag = true
	return
}

func (as *activityService) GMClearEndRecord() {
	as.rwm.Lock()
	defer as.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	for _, recordObj := range as.endRecordMap {
		recordObj.endTime = 0
		recordObj.updateTime = now
		recordObj.SetModified()
	}

	return
}

var (
	once sync.Once
	as   *activityService
)

func Init() (err error) {
	once.Do(func() {
		as = &activityService{}
		err = as.init()
	})
	return err
}

func GetActivityService() ActivityService {
	return as
}
