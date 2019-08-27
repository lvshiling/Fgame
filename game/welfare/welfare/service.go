package welfare

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	dummytemplate "fgame/fgame/game/dummy/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/game/welfare/dao"
	drewcrazyboxtemplate "fgame/fgame/game/welfare/drew/crazy_box/template"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfareranktemplate "fgame/fgame/game/welfare/rank/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
	"fgame/fgame/pkg/timeutils"
	"sync"

	log "github.com/Sirupsen/logrus"
)

//开服活动
type WelfareService interface {
	Heartbeat()
	Star() error
	// GM重置邮件
	GMResetRankEmailRecord()
	// 计算活动时间
	CountOpenActivityTime(groupId int32) (startTime int64, endTime int64)
	// 是否有领奖次数
	IsHadReceiveTimes(groupId int32, rewIndex, timesMax, addTimes int32) bool
	// 记录领奖次数
	AddReceiveTimes(groupId int32, rewIndex int32, times int32)
	// 获取已领取次数
	GetLeftReceiveTimes(groupId int32, rewIndex int32) int32
	// 获取领奖次数
	GetReceiveTimesList(groupId int32) []*welfaretypes.TimesLimitInfo
	// 折扣：是否剩余购买次数
	IsHadDiscountTimes(groupId int32, rewIndex, timesMax, addTimes int32) bool
	// 折扣：记录剩余购买次数
	AddDiscountTimes(groupId int32, rewIndex int32, times int32)
	// 折扣：获取购买次数
	GetDiscountTimes(groupId int32) []*welfaretypes.TimesLimitInfo
	// GM设置购买次数
	GMSetGlobalReceiveTimes(groupId, rewIndex, times int32)
	GMSetGlobalDiscountTimes(groupId, rewIndex, times int32)
	//添加拉霸日志
	AddLaBaLog(groupId int32, plName string, useGold, rewGold int32)
	//获取拉霸日志
	GetLaBaLogByTime(groupId int32, time int64) []*GoldLaBaLogObject
	//清空日志
	GMClearLog()
	//添加活动开启邮件
	AddStartMailRecord(group int32) bool
	//添加抽奖日志
	AddDrewLog(groupId int32, plName string, itemId, itemNum int32)
	//获取抽奖日志
	GetDrewLogByTime(groupId int32, time int64) []*DrewLogObject
	//添加疯狂宝箱日志
	AddCrazyBoxLog(groupId int32, plName string, itemId, itemNum int32)
	//获取疯狂宝箱日志
	GetCrazyBoxLogByTime(groupId int32, time int64) []*CrazyBoxLogObject
	//清空疯狂宝箱日志
	GMClearCrazyBoxLog()
	//BOSS已首杀记录
	GetBossFirstKillRecord(groupId int32) []int32
	//是否首杀
	IsBossFirstKill(groupId, bossId int32) bool
	//添加BOSS击杀记录
	AddBossKillRecord(groupId, bossId int32) bool
	//是否运营活动城战胜利
	IsAllianceWin(groupId int32, allianceId int64) bool
	//添加城战胜利记录
	AddAllianceWinRecord(groupId int32, allianceId int64) bool

	//获取活动开服时间
	GetServerStartTime() int64
	//获取活动合服时间
	GetServerMergeTime() int64

	// 循环活动
	GetXunHuanInfo() (arrIndex int32, day int32)
	// 是否循环
	IsOnXunHuan() bool
	// GM设置随机组
	GMSetXunHuanGroupArr(arrGroup int32)
}

type welfareService struct {
	//读写锁
	rwm sync.RWMutex
	//排行榜邮件记录
	rankEmailRecordMap map[int32]*OpenActivityEmailRecordObject
	//全服奖励次数限制
	rewardsLimitMap map[int32]*OpenActivityRewardsLimitObject
	//全服折扣商店次数限制
	discountLimitMap map[int32]*OpenActivityDiscountLimitObject
	//活动开启邮件记录
	startEmailRecordMap map[int32]*OpenActivityStartEmailObject
	//拉霸日志
	labaLogMap map[int32][]*GoldLaBaLogObject
	//抽奖日志
	drewLogMap map[int32][]*DrewLogObject
	//疯狂宝箱日志
	crazyBoxLogMap map[int32][]*CrazyBoxLogObject
	//上次拉霸插入日志时间
	lastDummyLaBaLogLime int64
	//上次抽奖插入日志时间
	lastDummyDrewLogLime int64
	//上次疯狂宝箱插入日志时间
	lastDummyCrazyBoxLogLime int64
	// BOSS首杀记录
	bossKillMap map[int32]*BossKillRecordObject
	// 城战助威
	allianceCheerMap map[int32]*AllianceCheerObject
	// 重置活动开服时间标志
	isHadResetActivityOpenServerTime bool
	// 循环活动数据
	xunHuanObj *ActivityXunHuanObject
	//
	hbRunner heartbeat.HeartbeatTaskRunner
}

const (
	noticeMinute = 30
	maxLogLen    = 10
)

//初始化
func (s *welfareService) init() (err error) {
	s.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	s.hbRunner.AddTask(CreateLaBaDummyLogTask(s))
	s.hbRunner.AddTask(CreateAddRankRewTask(s))
	s.hbRunner.AddTask(CreateStartEmailTask(s))
	s.hbRunner.AddTask(CreateOpenTimeChangedTask(s))
	s.hbRunner.AddTask(CreateRefreshXunHuanTask(s))
	s.labaLogMap = make(map[int32][]*GoldLaBaLogObject)
	s.drewLogMap = make(map[int32][]*DrewLogObject)
	s.crazyBoxLogMap = make(map[int32][]*CrazyBoxLogObject)

	//加载排行榜邮件记录
	//TODO:xzk:加载多久的活动，定时间
	s.initEmailRecord()

	//加载全服奖励次数限制
	s.initRewardsLimit()

	//加载全服折扣商店次数限制
	s.initDiscountLimit()

	//加载元宝拉霸日志
	// s.initGoldLaBaLog()

	//加载活动开启邮件通知记录
	s.initStartEmailRecord()

	//加载抽奖日志
	// s.initDrewLog()

	//加载疯狂宝箱日志
	// s.initCrazyBoxLog()

	//加载BOSS首杀记录
	s.initBossFirstKill()

	//加载城战助威
	s.initAllianceCheer()

	//初始化活动开服时间重置标志
	s.initResetFlag()

	//加载循环活动数据
	err = s.initXunHuanActivity()
	return
}

func (s *welfareService) initEmailRecord() {
	s.rankEmailRecordMap = make(map[int32]*OpenActivityEmailRecordObject)

	recordEntityList, err := dao.GetWelfareDao().GetOpenActivityEmailRecordList()
	if err != nil {
		return
	}
	for _, entity := range recordEntityList {
		obj := newOpenActivityEmailRecordObject()
		obj.FromEntity(entity)
		s.rankEmailRecordMap[obj.groupId] = obj
	}
}

func (s *welfareService) initStartEmailRecord() {
	s.startEmailRecordMap = make(map[int32]*OpenActivityStartEmailObject)

	recordEntityList, err := dao.GetWelfareDao().GetOpenActivityStartEmailList()
	if err != nil {
		return
	}
	for _, entity := range recordEntityList {
		obj := newOpenActivityStartEmailObject()
		obj.FromEntity(entity)
		s.startEmailRecordMap[obj.groupId] = obj
	}
}

func (s *welfareService) initRewardsLimit() {
	s.rewardsLimitMap = make(map[int32]*OpenActivityRewardsLimitObject)

	entityList, err := dao.GetWelfareDao().GetOpenActivityRewardsLimitList()
	if err != nil {
		return
	}
	for _, entity := range entityList {
		obj := newOpenActivityRewardsLimitObject()
		obj.FromEntity(entity)
		s.rewardsLimitMap[obj.groupId] = obj
	}
}

func (s *welfareService) initDiscountLimit() {
	s.discountLimitMap = make(map[int32]*OpenActivityDiscountLimitObject)

	entityList, err := dao.GetWelfareDao().GetOpenActivityDiscountLimitList()
	if err != nil {
		return
	}
	for _, entity := range entityList {
		obj := newOpenActivityDiscountLimitObject()
		obj.FromEntity(entity)
		s.discountLimitMap[obj.groupId] = obj
	}
}

// func (s *welfareService) initGoldLaBaLog() {
// 	s.labaLogMap = make(map[int32][]*GoldLaBaLogObject)
// 	entityList, err := dao.GetWelfareDao().GetOpenActivityLaBa()
// 	if err != nil {
// 		return
// 	}
// 	for _, entity := range entityList {
// 		obj := newGoldLaBaLogObject()
// 		obj.FromEntity(entity)
// 		s.labaLogMap[obj.groupId] = append(s.labaLogMap[obj.groupId], obj)
// 	}
// }

// func (s *welfareService) initDrewLog() {
// 	s.drewLogMap = make(map[int32][]*DrewLogObject)
// 	entityList, err := dao.GetWelfareDao().GetOpenActivityDrewLog(global.GetGame().GetServerIndex())
// 	if err != nil {
// 		return
// 	}
// 	for _, entity := range entityList {
// 		obj := newDrewLogObject()
// 		obj.FromEntity(entity)
// 		s.drewLogMap[obj.groupId] = append(s.drewLogMap[obj.groupId], obj)
// 	}
// }

// func (s *welfareService) initCrazyBoxLog() {
// 	s.crazyBoxLogMap = make(map[int32][]*CrazyBoxLogObject)
// 	entityList, err := dao.GetWelfareDao().GetOpenActivityCrazyBoxLog(global.GetGame().GetServerIndex())
// 	if err != nil {
// 		return
// 	}
// 	for _, entity := range entityList {
// 		obj := newCrazyBoxLogObject()
// 		obj.FromEntity(entity)
// 		s.crazyBoxLogMap[obj.groupId] = append(s.crazyBoxLogMap[obj.groupId], obj)
// 	}
// }

func (s *welfareService) initBossFirstKill() {
	s.bossKillMap = make(map[int32]*BossKillRecordObject)
	entityList, err := dao.GetWelfareDao().GetOpenActivityBossKilledIdList(global.GetGame().GetServerIndex())
	if err != nil {
		return
	}
	for _, entity := range entityList {
		obj := newBossKillRecordObject()
		obj.FromEntity(entity)
		s.bossKillMap[obj.groupId] = obj
	}
}

func (s *welfareService) initAllianceCheer() {
	s.allianceCheerMap = make(map[int32]*AllianceCheerObject)
	entityList, err := dao.GetWelfareDao().GetOpenActivityAllianceCheerList(global.GetGame().GetServerIndex())
	if err != nil {
		return
	}
	for _, entity := range entityList {
		obj := newAllianceCheerObject()
		obj.FromEntity(entity)
		s.allianceCheerMap[obj.groupId] = obj
	}
}

func (s *welfareService) initResetFlag() {
	now := global.GetGame().GetTimeService().Now()
	activityResetTime := constant.GetConstantService().GetActivityResetTime()
	if now > activityResetTime {
		s.isHadResetActivityOpenServerTime = true
	}
}

func (s *welfareService) initXunHuanActivity() (err error) {
	entity, err := dao.GetWelfareDao().GetOpenActivityXunHuan(global.GetGame().GetServerIndex())
	if err != nil {
		return
	}

	obj := newActivityXunHuanObject()
	if entity != nil {
		obj.FromEntity(entity)
	} else {
		now := global.GetGame().GetTimeService().Now()
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.activityDay = 0
		obj.arrGroup = 0
		obj.createTime = now
		obj.SetModified()
	}
	s.xunHuanObj = obj

	return
}

func (s *welfareService) Star() (err error) {
	err = s.registerAllRank()
	if err != nil {
		return
	}
	s.registerTempRank()
	return
}

func (s *welfareService) registerAllRank() (err error) {
	//注册排行榜
	rankTimeTmepMap := welfaretemplate.GetWelfareTemplateService().GetRankActivityTimeTemplate()
	for _, timeTempLsit := range rankTimeTmepMap {
		for _, timeTemp := range timeTempLsit {
			if timeTemp.IsTempRank() {
				continue
			}

			s.registRank(timeTemp.Group)
		}
	}

	return
}

func (s *welfareService) registerTempRank() (err error) {
	_, err = s.registerXunHuanRank()
	if err != nil {
		return
	}
	err = s.registerOpenServerNoMergeRank()
	if err != nil {
		return
	}

	err = s.registerOpenServerNoOpenRank()
	if err != nil {
		return
	}
	//TODO:zrc 注册周活动
	return
}

func (s *welfareService) registerXunHuanRank() (newGroupIdList []int32, err error) {
	openServerTime := s.GetServerStartTime()
	cycleDay := welfaretemplate.GetWelfareTemplateService().GetCurActivityXunHuanDay(openServerTime)
	if cycleDay != s.xunHuanObj.activityDay {
		return
	}

	// 同一天的活动
	xunhuanTemp := welfaretemplate.GetWelfareTemplateService().GetActivityXunHuanTemplate(openServerTime, s.xunHuanObj.arrGroup, s.xunHuanObj.activityDay)
	if xunhuanTemp == nil {
		return
	}
	// 注册循环活动的排行榜活动
	newGroupIdList = xunhuanTemp.GetGroupIdList()
	for _, groupId := range newGroupIdList {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		if timeTemp.GetOpenType() != welfaretypes.OpenActivityTypeRank {
			continue
		}
		s.registRank(groupId)
	}

	return
}

func (s *welfareService) registerOpenServerNoMergeRank() (err error) {
	//合服了
	mergeTime := s.GetServerMergeTime()
	if mergeTime != 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	startTime := s.GetServerStartTime()
	rankTimeTmepMap := welfaretemplate.GetWelfareTemplateService().GetRankActivityTimeTemplate()
	for _, timeTempLsit := range rankTimeTmepMap {
		for _, timeTemp := range timeTempLsit {
			if timeTemp.GetOpenTimeType() != welfaretypes.OpenTimeTypeOpenActivityNoMerge {
				continue
			}
			flag, err := timeTemp.IsOnTime(now, startTime)
			if err != nil {
				return err
			}
			if !flag {
				continue
			}
			s.registRank(timeTemp.Group)
		}
	}

	return
}

func (s *welfareService) registerOpenServerNoOpenRank() (err error) {

	now := global.GetGame().GetTimeService().Now()
	startTime := s.GetServerStartTime()
	rankTimeTmepMap := welfaretemplate.GetWelfareTemplateService().GetRankActivityTimeTemplate()
	for _, timeTempLsit := range rankTimeTmepMap {
		for _, timeTemp := range timeTempLsit {
			if timeTemp.CloseOpenDay <= 0 {
				continue
			}
			openServerTime := s.GetServerStartTime()
			openDays, _ := timeutils.DiffDay(now, openServerTime)
			if openDays < timeTemp.CloseOpenDay {
				continue
			}
			flag, err := timeTemp.IsOnTime(now, startTime)
			if err != nil {
				return err
			}
			if !flag {
				continue
			}
			s.registRank(timeTemp.Group)
		}
	}

	return
}

func (s *welfareService) registRank(groupId int32) {
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		return
	}
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.Infof("welfare_rank:排行榜模板不存在，groupId:%d", groupId)
		return 
	}
	groupTemp := groupInterface.(*welfareranktemplate.GroupTemplateRank)
	minCondition := groupTemp.GetRankMinLimitCondition()

	openServerTime := s.GetServerStartTime() //global.GetGame().GetServerTime()
	now := global.GetGame().GetTimeService().Now()
	mxExpireTime, _ := welfaretemplate.GetWelfareTemplateService().GetRankMaxEndTime(openServerTime)
	// str, _ := timeTemp.GetBeginTime(now, openServerTime)
	// end, _ := timeTemp.GetEndTime(now, openServerTime)
	str, end := s.CountOpenActivityTime(groupId)
	rankType, ok := timeTemp.GetOpenSubType().(welfaretypes.OpenActivityRankSubType)
	if !ok {
		return
	}

	config := ranktypes.NewRankConfig()
	if !timeTemp.IsRankCharge() && !timeTemp.IsRankCost() && !timeTemp.IsRankCharm() && !timeTemp.IsRankMarryDevelop() {
		if now > mxExpireTime {
			log.Infof("超过%s最大结束时间;groupId:%d, EndTime:%d", rankType.RankType().String(), groupId, mxExpireTime)
			return
		}
		config.IsCanExpire = true
	}
	config.GroupId = groupId
	config.StartTime = str
	config.EndTime = end
	config.MaxExpireTime = mxExpireTime
	config.MinCondition = minCondition
	config.RefreshTime = ranktypes.RankClassTypeLocalActivity.RankRefreshTime()
	config.ClassType = ranktypes.RankClassTypeLocalActivity
	rank.GetRankService().RegisterActivityRank(rankType.RankType(), config)
}

//心跳
func (s *welfareService) Heartbeat() {
	s.hbRunner.Heartbeat()
}

func (s *welfareService) addOpenRankRewards() (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	// openServerTime := s.GetServerStartTime() //global.GetGame().GetServerTime()
	rankTimeTmepMap := welfaretemplate.GetWelfareTemplateService().GetRankActivityTimeTemplate()

	for subType, timeTmepList := range rankTimeTmepMap {
		for _, timeTemp := range timeTmepList {
			if timeTemp.IsTempRank() {
				continue
			}
			_, endTime := s.CountOpenActivityTime(timeTemp.Group)
			// endTime, err := timeTemp.GetEndTime(now, openServerTime)
			// if err != nil {
			// 	return err
			// }

			if s.isExistRecord(timeTemp.Group, endTime) {
				continue
			}

			if now < endTime {
				continue
			}

			gameevent.Emit(welfareeventtypes.EventTypeRankEnd, timeTemp.Group, subType.RankType())

			// 添加奖励记录
			s.addRecord(timeTemp.Group, endTime)
		}
	}

	//重新注册
	err = s.registerAllRank()
	return
}

func (s *welfareService) isExistRecord(groupId int32, endTime int64) bool {
	recordObj, ok := s.rankEmailRecordMap[groupId]
	if !ok {
		return false
	}

	preEndTime := recordObj.endTime
	if endTime != preEndTime {
		return false
	}

	return true
}

func (s *welfareService) addRecord(groupId int32, endTime int64) bool {
	now := global.GetGame().GetTimeService().Now()
	obj, ok := s.rankEmailRecordMap[groupId]
	if !ok {
		obj = newOpenActivityEmailRecordObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.groupId = groupId
		obj.createTime = now
		s.rankEmailRecordMap[groupId] = obj
	}

	obj.endTime = endTime
	obj.updateTime = now
	obj.SetModified()

	return true
}

func (s *welfareService) GMResetRankEmailRecord() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range s.rankEmailRecordMap {
		obj.deleteTime = now
		obj.SetModified()
		delete(s.rankEmailRecordMap, obj.groupId)
	}
}

func (s *welfareService) CountOpenActivityTime(groupId int32) (startTime int64, endTime int64) {
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	openServerTime := s.GetServerStartTime()
	//开服几天内不开启
	if timeTemp.CloseOpenDay != 0 {
		openDays, _ := timeutils.DiffDay(now, openServerTime)
		if openDays < timeTemp.CloseOpenDay {
			return 0, 0
		}
	}

	switch timeTemp.GetOpenTimeType() {
	case welfaretypes.OpenTimeTypeMerge:
		{
			mergeTime := s.GetServerMergeTime() //merge.GetMergeService().GetMergeTime()
			startTime, _ = timeTemp.GetBeginTime(now, mergeTime)
			endTime, _ = timeTemp.GetEndTime(now, mergeTime)
		}
	case welfaretypes.OpenTimeTypeOpenActivity,
		welfaretypes.OpenTimeTypeOpenActivityNoMerge:
		{
			openServerTime := s.GetServerStartTime() //global.GetGame().GetServerTime()
			startTime, _ = timeTemp.GetBeginTime(now, openServerTime)
			endTime, _ = timeTemp.GetEndTime(now, openServerTime)
		}
	case welfaretypes.OpenTimeTypeNotTimeliness,
		welfaretypes.OpenTimeTypeSchedule,
		welfaretypes.OpenTimeTypeXunHuan,
		welfaretypes.OpenTimeTypeMergeXunHuan,
		welfaretypes.OpenTimeTypeWeek,
		welfaretypes.OpenTimeTypeMonth:
		{
			startTime, _ = timeTemp.GetBeginTime(now, 0)
			endTime, _ = timeTemp.GetEndTime(now, 0)
		}
	}
	return
}

func (s *welfareService) IsHadReceiveTimes(groupId int32, rewIndex, timesMax, addTimes int32) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	times := int32(0)
	endTime := int64(0)
	obj, ok := s.rewardsLimitMap[groupId]
	if !ok {
		goto Check
	}

	_, endTime = s.CountOpenActivityTime(groupId)
	if obj.endTime != endTime {
		goto Check
	}
	times = obj.timesMap[rewIndex]

Check:
	if times+addTimes <= timesMax {
		return true
	}

	return false
}

func (s *welfareService) AddReceiveTimes(groupId int32, rewIndex, times int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	startTime, endTime := s.CountOpenActivityTime(groupId)
	obj, ok := s.rewardsLimitMap[groupId]
	if !ok {
		id, _ := idutil.GetId()
		obj = newOpenActivityRewardsLimitObject()
		obj.id = id
		obj.groupId = groupId
		obj.serverId = global.GetGame().GetServerIndex()
		obj.timesMap = map[int32]int32{}
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now

		s.rewardsLimitMap[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.startTime = startTime
		obj.endTime = endTime
		obj.updateTime = now
		obj.timesMap = map[int32]int32{}
	}

	obj.updateTime = now
	obj.timesMap[rewIndex] += times
	obj.SetModified()
}

func (s *welfareService) GetReceiveTimesList(groupId int32) (timesList []*welfaretypes.TimesLimitInfo) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	obj, ok := s.rewardsLimitMap[groupId]
	if !ok {
		return nil
	}

	_, endTime := s.CountOpenActivityTime(groupId)
	if obj.endTime != endTime {
		return nil
	}

	for key, times := range obj.timesMap {
		data := &welfaretypes.TimesLimitInfo{}
		data.Key = key
		data.Times = times
		//构造虚拟数字
		timesList = append(timesList, data)
	}

	return
}

func (s *welfareService) GetLeftReceiveTimes(groupId, rewIndex int32) int32 {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	obj, ok := s.rewardsLimitMap[groupId]
	if !ok {
		return 0
	}

	_, endTime := s.CountOpenActivityTime(groupId)
	if obj.endTime != endTime {
		return 0
	}

	times, ok := obj.timesMap[rewIndex]
	if !ok {
		return 0
	}

	return times
}

func (s *welfareService) IsHadDiscountTimes(groupId int32, rewIndex, timesMax, addTimes int32) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	times := int32(0)
	_, endTime := s.CountOpenActivityTime(groupId)
	curDay := s.countDiscountDay(groupId)
	obj, ok := s.discountLimitMap[groupId]
	if !ok {
		goto Check
	}

	if obj.endTime != endTime {
		goto Check
	}

	if curDay != obj.discountDay {
		goto Check
	}

	times = obj.timesMap[rewIndex]

Check:
	if times+addTimes <= timesMax {
		return true
	}

	return false
}

func (s *welfareService) AddDiscountTimes(groupId int32, rewIndex, times int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.refreshDiscountTimesMap(groupId)

	now := global.GetGame().GetTimeService().Now()
	startTime, endTime := s.CountOpenActivityTime(groupId)
	obj, ok := s.discountLimitMap[groupId]
	if !ok {
		id, _ := idutil.GetId()
		obj = newOpenActivityDiscountLimitObject()
		obj.id = id
		obj.groupId = groupId
		obj.discountDay = s.countDiscountDay(groupId)
		obj.serverId = global.GetGame().GetServerIndex()
		obj.timesMap = map[int32]int32{}
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now

		s.discountLimitMap[groupId] = obj
	}

	obj.timesMap[rewIndex] += times
	obj.updateTime = now
	obj.SetModified()
}

func (s *welfareService) GetDiscountTimes(groupId int32) (timesList []*welfaretypes.TimesLimitInfo) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.refreshDiscountTimesMap(groupId)

	obj, ok := s.discountLimitMap[groupId]
	if !ok {
		return
	}

	// startTime, endTime := s.CountOpenActivityTime(groupId)
	// if obj.endTime != endTime {
	// 	now := global.GetGame().GetTimeService().Now()
	// 	obj.startTime = startTime
	// 	obj.endTime = endTime
	// 	obj.updateTime = now
	// 	obj.timesMap = map[int32]int32{}
	// 	obj.SetModified()
	// 	return
	// }

	for key, times := range obj.timesMap {
		data := &welfaretypes.TimesLimitInfo{}
		data.Key = key
		data.Times = times
		timesList = append(timesList, data)
	}

	return
}

func (s *welfareService) GMSetGlobalReceiveTimes(groupId, key, times int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	obj, ok := s.rewardsLimitMap[groupId]
	if !ok {
		startTime, endTime := s.CountOpenActivityTime(groupId)
		id, _ := idutil.GetId()
		obj = newOpenActivityRewardsLimitObject()
		obj.id = id
		obj.groupId = groupId
		obj.serverId = global.GetGame().GetServerIndex()
		obj.timesMap = map[int32]int32{}
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		s.rewardsLimitMap[groupId] = obj
	}

	obj.timesMap[key] = times
	obj.updateTime = now
	obj.SetModified()

	return
}

func (s *welfareService) GMSetGlobalDiscountTimes(groupId, key, times int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.refreshDiscountTimesMap(groupId)

	now := global.GetGame().GetTimeService().Now()
	obj, ok := s.discountLimitMap[groupId]
	if !ok {
		startTime, endTime := s.CountOpenActivityTime(groupId)
		id, _ := idutil.GetId()
		obj = newOpenActivityDiscountLimitObject()
		obj.id = id
		obj.groupId = groupId
		obj.discountDay = s.countDiscountDay(groupId)
		obj.serverId = global.GetGame().GetServerIndex()
		obj.timesMap = map[int32]int32{}
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		s.discountLimitMap[groupId] = obj
	}

	obj.timesMap[key] = times
	obj.updateTime = now
	obj.SetModified()

	return
}

func (s *welfareService) GMClearLog() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.labaLogMap = map[int32][]*GoldLaBaLogObject{}
}

func (s *welfareService) GMClearCrazyBoxLog() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.crazyBoxLogMap = map[int32][]*CrazyBoxLogObject{}
}

func (s *welfareService) refreshDiscountTimesMap(groupId int32) {
	obj, ok := s.discountLimitMap[groupId]
	if !ok {
		return
	}

	startTime, endTime := s.CountOpenActivityTime(groupId)
	curDay := s.countDiscountDay(groupId)
	if obj.endTime != endTime {
		obj.startTime = startTime
		obj.endTime = endTime
		obj.discountDay = curDay
		obj.timesMap = map[int32]int32{}
	}

	if curDay != obj.discountDay {
		obj.timesMap = map[int32]int32{}
		obj.discountDay = curDay
	}

	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
}

func (s *welfareService) countDiscountDay(groupId int32) int32 {
	now := global.GetGame().GetTimeService().Now()
	startTime, _ := s.CountOpenActivityTime(groupId)
	if startTime == 0 {
		return 0
	}
	if now < startTime {
		return 0
	}

	curDay, _ := timeutils.DiffDay(now, startTime)
	return curDay
}

//生成拉霸虚拟日志
func (s *welfareService) addDummyLaBaLog() (err error) {
	now := global.GetGame().GetTimeService().Now()
	diffTime := now - s.lastDummyLaBaLogLime
	minTime := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeLaBaDummyLogAddTimeMin))
	maxTime := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeLaBaDummyLogAddTimeMax))
	randTime := int64(mathutils.RandomRange(minTime, maxTime))
	if diffTime < randTime {
		return
	}

	name := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeGoldLaBa
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		labaTemp := welfaretemplate.GetWelfareTemplateService().GetRandomLaBaTemplate(groupId)
		if labaTemp == nil {
			continue
		}
		data := welfareeventtypes.CreateLaBaAddLogEventData(name, labaTemp.Investment, int32(labaTemp.RandomRuleGold()))
		gameevent.Emit(welfareeventtypes.EventTypeLaBaAddLog, groupId, data)
	}

	s.lastDummyLaBaLogLime = now
	return
}

//生成充值抽奖虚拟日志
func (s *welfareService) addDummyDrewLog() (err error) {
	now := global.GetGame().GetTimeService().Now()
	diffTime := now - s.lastDummyDrewLogLime
	minTime := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDrewDummyLogAddTimeMin))
	maxTime := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDrewDummyLogAddTimeMax))
	randTime := int64(mathutils.RandomRange(minTime, maxTime))
	if diffTime < randTime {
		return
	}

	plName := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeChargeDrew
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		drewTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplate(groupId)
		if drewTemp == nil {
			continue
		}
		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(drewTemp.DropId)
		if dropData == nil {
			continue
		}
		drewLogEventData := welfareeventtypes.CreateDrewAddLogEventData(plName, dropData.ItemId, dropData.Num)
		gameevent.Emit(welfareeventtypes.EventTypeDrewAddLog, groupId, drewLogEventData)
	}

	s.lastDummyDrewLogLime = now
	return
}

//生成虚拟疯狂宝箱日志
func (s *welfareService) addDummyCrazyBoxLog() (err error) {
	now := global.GetGame().GetTimeService().Now()
	diffTime := now - s.lastDummyCrazyBoxLogLime
	minTime := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDrewDummyLogAddTimeMin))
	maxTime := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDrewDummyLogAddTimeMax))
	randTime := int64(mathutils.RandomRange(minTime, maxTime))
	if diffTime < randTime {
		return
	}

	plName := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeCrazyBox
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*drewcrazyboxtemplate.GroupTemplateCrazyBox)

		randomLevel := groupTemp.GetRandomOpenActivityCrazyBoxLevel()
		openActivityTemp := groupTemp.GetOpenActivityCrazyBox(randomLevel)
		if openActivityTemp == nil {
			continue
		}

		drewTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplateByArg(groupId, randomLevel)
		if drewTemp == nil {
			continue
		}
		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(drewTemp.DropId)
		if dropData == nil {
			continue
		}
		logEventData := welfareeventtypes.CreateCrazyBoxAddLogEventData(plName, dropData.ItemId, dropData.Num)
		gameevent.Emit(welfareeventtypes.EventTypeCrazyBoxAddLog, groupId, logEventData)
	}

	s.lastDummyCrazyBoxLogLime = now
	return
}

func (s *welfareService) AddLaBaLog(groupId int32, plName string, needGold, rewGold int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj := s.createLogObj(groupId, plName, needGold, rewGold)
	s.labaLogMap[groupId] = append(s.labaLogMap[groupId], obj)
}

func (s *welfareService) GetLaBaLogByTime(groupId int32, time int64) []*GoldLaBaLogObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	logList := s.labaLogMap[groupId]
	for index, log := range logList {
		if time < log.updateTime {
			return logList[index:]
		}
	}

	return nil
}

func (s *welfareService) createLogObj(groupId int32, playerName string, needGold, rewGold int32) *GoldLaBaLogObject {
	now := global.GetGame().GetTimeService().Now()
	logList := s.labaLogMap[groupId]
	var obj *GoldLaBaLogObject
	if len(logList) < maxLogLen {
		obj = newGoldLaBaLogObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.createTime = now
	} else {
		obj = logList[0]
		s.labaLogMap[groupId] = logList[1:]
	}

	obj.groupId = groupId
	obj.playerName = playerName
	obj.costGold = needGold
	obj.rewGold = rewGold
	obj.updateTime = now
	obj.SetModified()

	return obj
}

func (s *welfareService) AddDrewLog(groupId int32, plName string, itemId, itemNum int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj := s.createDrewLogObj(groupId, plName, itemId, itemNum)
	s.drewLogMap[groupId] = append(s.drewLogMap[groupId], obj)
}

func (s *welfareService) GetDrewLogByTime(groupId int32, time int64) []*DrewLogObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	logList := s.drewLogMap[groupId]
	for index, log := range logList {
		if time < log.updateTime {
			return logList[index:]
		}
	}

	return nil
}

func (s *welfareService) createDrewLogObj(groupId int32, playerName string, itemId, itemNum int32) *DrewLogObject {
	now := global.GetGame().GetTimeService().Now()
	logList := s.drewLogMap[groupId]
	var obj *DrewLogObject
	if len(logList) < maxLogLen {
		obj = newDrewLogObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.createTime = now
	} else {
		obj = logList[0]
		s.drewLogMap[groupId] = logList[1:]
	}

	obj.groupId = groupId
	obj.playerName = playerName
	obj.itemId = itemId
	obj.itemNum = itemNum
	obj.updateTime = now
	obj.SetModified()

	return obj
}

func (s *welfareService) AddCrazyBoxLog(groupId int32, plName string, itemId, itemNum int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj := s.createCrazyBoxLogObj(groupId, plName, itemId, itemNum)
	s.crazyBoxLogMap[groupId] = append(s.crazyBoxLogMap[groupId], obj)
}

func (s *welfareService) GetCrazyBoxLogByTime(groupId int32, time int64) []*CrazyBoxLogObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	logList := s.crazyBoxLogMap[groupId]
	for index, log := range logList {
		if time < log.updateTime {
			return logList[index:]
		}
	}

	return nil
}

func (s *welfareService) createCrazyBoxLogObj(groupId int32, playerName string, itemId, itemNum int32) *CrazyBoxLogObject {
	now := global.GetGame().GetTimeService().Now()
	logList := s.crazyBoxLogMap[groupId]
	var obj *CrazyBoxLogObject
	if len(logList) < maxLogLen {
		obj = newCrazyBoxLogObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.createTime = now
	} else {
		obj = logList[0]
		s.crazyBoxLogMap[groupId] = logList[1:]
	}

	obj.groupId = groupId
	obj.playerName = playerName
	obj.itemId = itemId
	obj.itemNum = itemNum
	obj.updateTime = now
	obj.SetModified()

	return obj
}

func (s *welfareService) checkActivityStartMail() {
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetAllActivityTimeTemplate()
	for _, timeTemp := range timeTempList {
		if len(timeTemp.MailDes) <= 0 {
			continue
		}

		group := timeTemp.Group
		_, endTime := s.CountOpenActivityTime(group)
		if s.isExistStartMailRecord(group, endTime) {
			continue
		}

		gameevent.Emit(welfareeventtypes.EventTypeCheckActivityOpenMail, group, nil)
	}
}

func (s *welfareService) isExistStartMailRecord(groupId int32, endTime int64) bool {
	recordObj, ok := s.startEmailRecordMap[groupId]
	if !ok {
		return false
	}

	preEndTime := recordObj.endTime
	if endTime != preEndTime {
		return false
	}

	return true
}

func (s *welfareService) AddStartMailRecord(groupId int32) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	obj, ok := s.startEmailRecordMap[groupId]
	if !ok {
		obj = newOpenActivityStartEmailObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.groupId = groupId
		obj.createTime = now
	}
	_, endTime := s.CountOpenActivityTime(groupId)
	obj.endTime = endTime
	obj.updateTime = now
	obj.SetModified()

	s.startEmailRecordMap[groupId] = obj
	return true
}

func (s *welfareService) GetBossFirstKillRecord(groupId int32) []int32 {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	obj, ok := s.bossKillMap[groupId]
	if !ok {
		return nil
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return nil
	}
	openTime := s.GetServerStartTime()
	mergeTime := s.GetServerMergeTime()
	_, endTime := groupInterface.GetActivityTime(openTime, mergeTime)
	if obj.endTime != endTime {
		return nil
	}

	return obj.bossIdList
}

func (s *welfareService) AddBossKillRecord(groupId, bossId int32) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	openTime := s.GetServerStartTime()
	mergeTime := s.GetServerMergeTime()
	startTime, endTime := groupInterface.GetActivityTime(openTime, mergeTime)
	now := global.GetGame().GetTimeService().Now()
	obj, ok := s.bossKillMap[groupId]
	if !ok {
		obj = newBossKillRecordObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.groupId = groupId
		obj.bossIdList = []int32{}
		obj.serverId = global.GetGame().GetServerIndex()
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		s.bossKillMap[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.bossIdList = []int32{}
		obj.startTime = startTime
		obj.endTime = endTime
	}

	obj.bossIdList = append(obj.bossIdList, bossId)
	obj.updateTime = now
	obj.SetModified()

	return true
}

func (s *welfareService) IsBossFirstKill(groupId, bossId int32) (flag bool) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	openTime := s.GetServerStartTime()
	mergeTime := s.GetServerMergeTime()
	_, endTime := groupInterface.GetActivityTime(openTime, mergeTime)
	obj, ok := s.bossKillMap[groupId]
	if !ok {
		return false
	}

	if endTime != obj.endTime {
		return false
	}

	isFirstKill := obj.isExistBossId(bossId)
	return isFirstKill
}

func (s *welfareService) AddAllianceWinRecord(groupId int32, allianceId int64) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	openTime := s.GetServerStartTime()
	mergeTime := s.GetServerMergeTime()
	startTime, endTime := groupInterface.GetActivityTime(openTime, mergeTime)
	now := global.GetGame().GetTimeService().Now()
	obj, ok := s.allianceCheerMap[groupId]
	if !ok {
		obj = newAllianceCheerObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.groupId = groupId
		obj.allianceId = allianceId
		obj.serverId = global.GetGame().GetServerIndex()
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		s.allianceCheerMap[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.allianceId = allianceId
		obj.startTime = startTime
		obj.endTime = endTime
	}

	obj.updateTime = now
	obj.SetModified()

	return true
}

func (s *welfareService) IsAllianceWin(groupId int32, allianceId int64) (flag bool) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return false
	}
	openTime := s.GetServerStartTime()
	mergeTime := s.GetServerMergeTime()
	_, endTime := groupInterface.GetActivityTime(openTime, mergeTime)
	obj, ok := s.allianceCheerMap[groupId]
	if !ok {
		return false
	}

	if endTime != obj.endTime {
		return false
	}

	if allianceId == 0 {
		return false
	}

	if obj.allianceId != allianceId {
		return false
	}

	return true
}

const (
	noResetDay = 16
)

func (s *welfareService) GetServerStartTime() int64 {
	activityResetTime := constant.GetConstantService().GetActivityResetTime()
	startTime := global.GetGame().GetServerTime()
	//当前时间小于重置时间
	now := global.GetGame().GetTimeService().Now()
	if now < activityResetTime {
		return startTime
	}

	//重置时间小于开服时间
	if activityResetTime < startTime {
		return startTime
	}

	diffDay, _ := timeutils.DiffDay(activityResetTime, startTime)
	if diffDay > noResetDay {
		return activityResetTime
	}

	return startTime
}

const (
	startMergeDiffDay = 7
)

func (s *welfareService) GetServerMergeTime() int64 {
	mergeTime := merge.GetMergeService().GetMergeTime()
	if mergeTime <= 0 {
		return 0
	}

	startTime := s.GetServerStartTime()
	atLeastMergeTime, _ := timeutils.AfterNDayOfTime(startTime, startMergeDiffDay)
	if atLeastMergeTime > mergeTime {
		return atLeastMergeTime
	}

	return mergeTime
}

func (s *welfareService) GetXunHuanInfo() (arrGroup int32, day int32) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.xunHuanObj.arrGroup, s.xunHuanObj.activityDay
}

func (s *welfareService) GMSetXunHuanGroupArr(arrGroup int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	s.xunHuanObj.arrGroup = arrGroup
	s.xunHuanObj.updateTime = now
	s.xunHuanObj.SetModified()

	// 注册循环活动的排行榜活动
	openServerTime := s.GetServerStartTime()
	xunhuanTemp := welfaretemplate.GetWelfareTemplateService().GetActivityXunHuanTemplate(openServerTime, arrGroup, s.xunHuanObj.activityDay)
	if xunhuanTemp == nil {
		return
	}
	for _, groupId := range xunhuanTemp.GetGroupIdList() {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		if timeTemp.GetOpenType() != welfaretypes.OpenActivityTypeRank {
			continue
		}

		s.registRank(groupId)
	}
	return
}

func (s *welfareService) IsOnXunHuan() bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.isOnXunHuan()
}

const (
	xunHuanOpenDay  = 16
	xunHuanMergeDay = 7
)

func (s *welfareService) isOnXunHuan() bool {
	// //合服第一天有循环活动
	// if s.isMergeXunHuan() {
	// 	return true
	// }

	// TODO:xzk:日期走配置
	// 服务器16天后开启
	now := global.GetGame().GetTimeService().Now()
	openServerTime := s.GetServerStartTime()
	openDiff, _ := timeutils.DiffDay(now, openServerTime)
	if openDiff+1 <= xunHuanOpenDay {
		return false
	}

	// // 合服7天后开启
	// mergeServerTime := s.GetServerMergeTime()
	// if now > mergeServerTime {
	// 	mergeDiff, _ := timeutils.DiffDay(now, mergeServerTime)
	// 	if mergeDiff+1 <= xunHuanMergeDay {
	// 		return false
	// 	}
	// }

	return true
}

// // 特殊处理：合服第一天有循环活动
// func (s *welfareService) isMergeXunHuan() bool {
// 	now := global.GetGame().GetTimeService().Now()
// 	mergeServerTime := s.GetServerMergeTime()
// 	mergeDiff, _ := timeutils.DiffDay(now, mergeServerTime)

// 	//合服第一天有循环活动
// 	if mergeDiff == 0 && s.xunHuanObj.isOnXunHuan() {
// 		return true
// 	}

// 	return false
// }

func (s *welfareService) checkRefreshXunHuan() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if !s.isOnXunHuan() {
		return
	}

	// if s.isMergeXunHuan() {
	// 	return
	// }

	now := global.GetGame().GetTimeService().Now()
	curArrGroup := s.xunHuanObj.arrGroup
	curDay := s.xunHuanObj.activityDay
	openServerTime := s.GetServerStartTime()
	cycleDay := welfaretemplate.GetWelfareTemplateService().GetCurActivityXunHuanDay(openServerTime)
	if cycleDay == curDay {
		return
	}

	// 发送循环活动的排行榜奖励
	curXunhuanTemp := welfaretemplate.GetWelfareTemplateService().GetActivityXunHuanTemplate(openServerTime, curArrGroup, curDay)
	if curXunhuanTemp != nil {
		oldEndTime := s.xunHuanObj.endTime
		for _, groupId := range curXunhuanTemp.GetGroupIdList() {
			timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
			if timeTemp.GetOpenType() != welfaretypes.OpenActivityTypeRank {
				continue
			}

			if s.isExistRecord(groupId, oldEndTime) {
				continue
			}
			subType := timeTemp.GetOpenSubType().(welfaretypes.OpenActivityRankSubType)
			gameevent.Emit(welfareeventtypes.EventTypeRankEnd, groupId, subType.RankType())
			// 添加奖励记录
			s.addRecord(groupId, oldEndTime)
		}
	}

	// 刷新数据
	newStartTime, _ := timeutils.BeginOfNow(now)
	newEndTime := newStartTime + int64(common.DAY)
	s.xunHuanObj.activityDay = welfaretemplate.GetWelfareTemplateService().GetCurActivityXunHuanDay(openServerTime)
	s.xunHuanObj.arrGroup = welfaretemplate.GetWelfareTemplateService().GetRandomActivityXunHuanArrGroup(openServerTime)
	s.xunHuanObj.startTime = newStartTime
	s.xunHuanObj.endTime = newEndTime
	s.xunHuanObj.updateTime = now
	s.xunHuanObj.SetModified()

	// newXunhuanTemp := welfaretemplate.GetWelfareTemplateService().GetActivityXunHuanTemplate(openServerTime, s.xunHuanObj.arrGroup, s.xunHuanObj.activityDay)
	// if newXunhuanTemp == nil {
	// 	return
	// }
	// // 注册循环活动的排行榜活动
	// newGroupIdList := newXunhuanTemp.GetGroupIdList()
	// for _, groupId := range newGroupIdList {
	// 	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	// 	if timeTemp.GetOpenType() != welfaretypes.OpenActivityTypeRank {
	// 		continue
	// 	}
	// 	s.registRank(groupId)
	// }
	newGroupIdList, err := s.registerXunHuanRank()
	if err != nil {
		return
	}
	gameevent.Emit(welfareeventtypes.EventTypeRefreshXunHuanActivity, nil, newGroupIdList)

	//TODO:zrc 注册周活动
}

func (s *welfareService) checkTempRank() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	openServerTime := s.GetServerStartTime()
	//发送开服未合服的排行榜
	mergeTime := s.GetServerMergeTime()
	if mergeTime == 0 {
		rankTimeTmepMap := welfaretemplate.GetWelfareTemplateService().GetRankActivityTimeTemplate()
		for _, timeTempLsit := range rankTimeTmepMap {
			for _, timeTemp := range timeTempLsit {
				if timeTemp.GetOpenTimeType() != welfaretypes.OpenTimeTypeOpenActivityNoMerge {
					continue
				}
				endTime, err := timeTemp.GetEndTime(now, openServerTime)
				if err != nil {
					return
				}

				if s.isExistRecord(timeTemp.Group, endTime) {
					continue
				}

				if now < endTime {
					continue
				}
				subType := timeTemp.GetOpenSubType().(welfaretypes.OpenActivityRankSubType)
				gameevent.Emit(welfareeventtypes.EventTypeRankEnd, timeTemp.Group, subType.RankType())

				// 添加奖励记录
				s.addRecord(timeTemp.Group, endTime)
			}
		}
	}

	rankTimeTmepMap := welfaretemplate.GetWelfareTemplateService().GetRankActivityTimeTemplate()
	for _, timeTempLsit := range rankTimeTmepMap {
		for _, timeTemp := range timeTempLsit {
			if timeTemp.CloseOpenDay <= 0 {
				continue
			}
			openServerTime := s.GetServerStartTime()
			openDays, _ := timeutils.DiffDay(now, openServerTime)
			if openDays < timeTemp.CloseOpenDay {
				continue
			}

			endTime, err := timeTemp.GetEndTime(now, openServerTime)
			if err != nil {
				return
			}

			if s.isExistRecord(timeTemp.Group, endTime) {
				continue
			}

			if now < endTime {
				continue
			}
			subType := timeTemp.GetOpenSubType().(welfaretypes.OpenActivityRankSubType)
			gameevent.Emit(welfareeventtypes.EventTypeRankEnd, timeTemp.Group, subType.RankType())

			// 添加奖励记录
			s.addRecord(timeTemp.Group, endTime)
		}
	}

	err := s.registerOpenServerNoMergeRank()
	if err != nil {
		return
	}
	err = s.registerOpenServerNoOpenRank()
	if err != nil {
		return
	}
}

func (s *welfareService) checkopenTimeChanged() {
	now := global.GetGame().GetTimeService().Now()
	activityResetTime := constant.GetConstantService().GetActivityResetTime()
	if now < activityResetTime {
		return
	}

	if s.isHadResetActivityOpenServerTime {
		return
	}

	s.isHadResetActivityOpenServerTime = true
	gameevent.Emit(welfareeventtypes.EventTypeActivityOpenTimeChanged, nil, nil)
	log.Infoln("重置运营活动开服时间")
}

var (
	once sync.Once
	s    *welfareService
)

func Init() (err error) {
	once.Do(func() {
		s = &welfareService{}
		err = s.init()
	})
	return err
}

func GetWelfareService() WelfareService {
	return s
}
