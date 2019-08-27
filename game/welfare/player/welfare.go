package player

import (
	"fgame/fgame/core/heartbeat"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/welfare/dao"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//玩家开服活动管理器
type PlayerWelfareManager struct {
	p player.Player
	//活动数据
	openActivityByGroupMap map[int32]*PlayerOpenActivityObject
	//活动充值记录
	openActivityChargeByGroupMap map[int32]*PlayerOpenActivityChargeObject
	//活动消费记录
	openActivityCostByGroupMap map[int32]*PlayerOpenActivityCostObject
	//首充数据
	firstChargeObject *PlayerFirstChargeObject
	//次数数据记录
	activityNumRecordObject map[int32]*PlayerActivityNumRecordObject
	//活动开启提醒邮件记录
	openActivityMailRecord map[int32]*PlayerActivityOpenMailObject
	//增长数据记录
	activityAddNumObject map[int32]*PlayerActivityAddNumObject
	//
	hbRunner heartbeat.HeartbeatTaskRunner
}

func (m *PlayerWelfareManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerWelfareManager) Load() (err error) {
	//加载玩家福利厅
	if err = m.loadOpenActivity(); err != nil {
		return
	}
	//加载活动充值数据
	if err = m.loadOpenActivityChargeList(); err != nil {
		return
	}

	//加载活动消费数据
	if err = m.loadOpenActivityCostList(); err != nil {
		return
	}

	//加载活动抽奖数据
	if err = m.loadActivityNumRecordList(); err != nil {
		return
	}

	//加载活动抽奖数据
	if err = m.loadActivityAddNumList(); err != nil {
		return
	}

	//加载首充数据
	if err = m.loadFirstCharge(); err != nil {
		return
	}

	//加载活动开启邮件数据
	if err = m.loadActivityOpenMailRecordList(); err != nil {
		return
	}

	return nil
}

//加载玩家福利厅
func (m *PlayerWelfareManager) loadOpenActivity() (err error) {
	openActivityEntityList, err := dao.GetWelfareDao().GetOpenActivityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range openActivityEntityList {
		obj := newPlayerOpenActivityObject(m.p)
		err = obj.FromEntity(entity)
		if err != nil {
			log.WithFields(
				log.Fields{
					"playerId": entity.PlayerId,
					"groupId":  entity.GroupId,
					"err":      err,
				}).Error("welfare:加载活动失败")
			continue
		}
		m.addOpenActivityObj(obj)
	}
	return nil
}

//加载活动充值数据
func (m *PlayerWelfareManager) loadOpenActivityChargeList() (err error) {
	chargeEntityList, err := dao.GetWelfareDao().GetOpenActivityChargeList(m.p.GetId())
	if err != nil {
		return
	}
	for _, chargeEntity := range chargeEntityList {
		chargeObj := newPlayerOpenActivityChargeObject(m.p)
		chargeObj.FromEntity(chargeEntity)
		m.openActivityChargeByGroupMap[chargeObj.groupId] = chargeObj
	}
	return nil
}

//加载活动消费数据
func (m *PlayerWelfareManager) loadOpenActivityCostList() (err error) {
	costEntityList, err := dao.GetWelfareDao().GetOpenActivityCostList(m.p.GetId())
	if err != nil {
		return
	}
	for _, costEntity := range costEntityList {
		costObj := newPlayerOpenActivityCostObject(m.p)
		costObj.FromEntity(costEntity)
		m.openActivityCostByGroupMap[costObj.groupId] = costObj
	}

	return nil
}

//加载活动抽奖数据
func (m *PlayerWelfareManager) loadActivityNumRecordList() (err error) {
	drewEntityList, err := dao.GetWelfareDao().GetActivityNumRecordList(m.p.GetId())
	if err != nil {
		return
	}
	for _, drewEntity := range drewEntityList {
		drewObj := newPlayerActivityNumRecordObject(m.p)
		drewObj.FromEntity(drewEntity)
		m.activityNumRecordObject[drewObj.groupId] = drewObj
	}

	return nil
}

//加载活动增长数据
func (m *PlayerWelfareManager) loadActivityAddNumList() (err error) {
	addNumEntityList, err := dao.GetWelfareDao().GetActivityAddNumList(m.p.GetId())
	if err != nil {
		return
	}
	for _, addNumEntity := range addNumEntityList {
		addNumObj := newPlayerActivityAddNumObject(m.p)
		addNumObj.FromEntity(addNumEntity)
		m.activityAddNumObject[addNumObj.groupId] = addNumObj
	}

	return nil
}

//加载活动开启邮件数据
func (m *PlayerWelfareManager) loadActivityOpenMailRecordList() (err error) {
	mailEntityList, err := dao.GetWelfareDao().GetPlayerActivityOpenMailRecordList(m.p.GetId())
	if err != nil {
		return
	}
	for _, entity := range mailEntityList {
		obj := newPlayerActivityOpenMailObject(m.p)
		obj.FromEntity(entity)
		m.openActivityMailRecord[obj.group] = obj
	}

	return nil
}

//加载首冲
func (m *PlayerWelfareManager) loadFirstCharge() (err error) {
	//加载首充数据
	firstChargeEntity, err := dao.GetWelfareDao().GetPlayerFirstCharge(m.p.GetId())
	if err != nil {
		return
	}

	if firstChargeEntity != nil {
		obj := newPlayerFirstChargeObject(m.p)
		obj.FromEntity(firstChargeEntity)
		m.firstChargeObject = obj
	}
	return nil
}

func (m *PlayerWelfareManager) newOpenActivityObj(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, groupId int32) *PlayerOpenActivityObject {
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)

	obj := newPlayerOpenActivityObject(m.p)
	obj.id = id
	obj.activityType = typ
	obj.activitySubType = subType
	obj.groupId = groupId

	newData := CreateEmptyOpenActivityData(typ, subType)
	if newData == nil {
		panic(fmt.Errorf("welfare:创建初始活动数据错误，typ:%d,subType:%d", typ, subType))
	}
	obj.activityData = newData
	obj.startTime = startTime
	obj.endTime = endTime
	obj.createTime = now
	obj.updateTime = now

	initHandler := GetInfoInitHandler(typ, subType)
	if initHandler != nil {
		initHandler.InitInfo(obj)
	}

	obj.SetModified()
	return obj
}

func (m *PlayerWelfareManager) addOpenActivityObj(obj *PlayerOpenActivityObject) {
	_, ok := m.openActivityByGroupMap[obj.groupId]
	if ok {
		panic(fmt.Errorf("welfare:groupId:[%d]已经存在", obj.groupId))
	}
	m.openActivityByGroupMap[obj.groupId] = obj
}

//加载后
func (m *PlayerWelfareManager) AfterLoad() (err error) {
	m.hbRunner.AddTask(CreateTimeReddotNoticeTask(m.p))
	m.hbRunner.AddTask(CreateRefreshActivityTask(m.p))
	return
}

//心跳
func (m *PlayerWelfareManager) Heartbeat() {
	m.hbRunner.Heartbeat()
}

// 刷新运营活动数据
func (m *PlayerWelfareManager) RefreshActivityData(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) (err error) {
	refreshHandler := GetInfoRefreshHandler(typ, subType)
	if refreshHandler != nil {
		for _, obj := range m.openActivityByGroupMap {
			if obj.activityType != typ || obj.activitySubType != subType {
				continue
			}

			err = refreshHandler.RefreshInfo(obj)
			if err != nil {
				return
			}
		}
	}

	return
}

const (
	sevenDay = 7
)

// 刷新运营活动数据
func (m *PlayerWelfareManager) RefreshActivityDataByGroupId(groupId int32) (err error) {
	obj, ok := m.openActivityByGroupMap[groupId]
	if !ok {
		return
	}
	if obj.GetEndTime() != 0 {
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetEndTime())
		if diff > sevenDay {
			return
		}
	}

	refreshHandler := GetInfoRefreshHandler(obj.activityType, obj.activitySubType)
	if refreshHandler == nil {
		return
	}

	err = refreshHandler.RefreshInfo(obj)
	if err != nil {
		return
	}

	return
}

//获取次数记录
func (m *PlayerWelfareManager) GetActivityCountNum(groupId int32) int32 {
	obj := m.getActivityNumRecordObj(groupId)
	if obj == nil {
		return 0
	}

	_, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	if obj.endTime != endTime {
		return 0
	}

	return obj.times
}

//获取增长数据
func (m *PlayerWelfareManager) GetActivityAddNumVal(groupId int32) int32 {
	obj := m.getActivityAddNumObj(groupId)
	if obj == nil {
		return 0
	}

	_, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	if obj.endTime != endTime {
		return 0
	}

	return obj.addNum
}

//添加活动充值记录
func (m *PlayerWelfareManager) AddOpenActivityChargeRecord(groupId int32, goldNum int32) int32 {
	now := global.GetGame().GetTimeService().Now()
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := m.getOpenActivityCharge(groupId)
	if obj == nil {
		obj = newPlayerOpenActivityChargeObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.groupId = groupId
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		m.openActivityChargeByGroupMap[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.goldNum = 0
		obj.startTime = startTime
		obj.endTime = endTime
	}

	obj.goldNum += goldNum
	obj.updateTime = now
	obj.SetModified()

	return obj.goldNum
}

//同步充值记录
func (m *PlayerWelfareManager) SyncFirstDayChargeRecord(goldNum int32) {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.openActivityChargeByGroupMap {
		groupId := obj.groupId
		startTime, _ := welfare.GetWelfareService().CountOpenActivityTime(groupId)
		flag, _ := timeutils.IsSameDay(startTime, now)
		if !flag {
			continue
		}
		if obj.goldNum == goldNum {
			continue
		}
		obj.goldNum = goldNum
		obj.updateTime = now
		obj.SetModified()
	}
}

// 同步消费记录
func (m *PlayerWelfareManager) SyncFirstDayCostRecord(costNum int64) {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.openActivityCostByGroupMap {
		groupId := obj.groupId
		startTime, _ := welfare.GetWelfareService().CountOpenActivityTime(groupId)
		flag, _ := timeutils.IsSameDay(startTime, now)
		if !flag {
			continue
		}
		if obj.goldNum == costNum {
			continue
		}
		obj.goldNum = costNum
		obj.updateTime = now
		obj.SetModified()
	}
}

//获取活动充值金额
func (m *PlayerWelfareManager) GetRankChargeNum(groupId int32) int32 {
	obj := m.getOpenActivityCharge(groupId)
	if obj == nil {
		return 0
	}

	_, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	if obj.endTime != endTime {
		return 0
	}

	return obj.goldNum
}

//获取活动消费金额
func (m *PlayerWelfareManager) GetRankCostNum(groupId int32) int64 {
	obj := m.getOpenActivityCost(groupId)
	if obj == nil {
		return 0
	}

	_, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	if obj.endTime != endTime {
		return 0
	}

	return obj.goldNum
}

//添加活动消费记录
func (m *PlayerWelfareManager) AddOpenActivityCostRecord(groupId int32, goldNum int64) int64 {
	now := global.GetGame().GetTimeService().Now()
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := m.getOpenActivityCost(groupId)
	if obj == nil {
		obj = newPlayerOpenActivityCostObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.groupId = groupId
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		m.openActivityCostByGroupMap[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.goldNum = 0
		obj.startTime = startTime
		obj.endTime = endTime
	}

	obj.goldNum += goldNum
	obj.updateTime = now
	obj.SetModified()

	return obj.goldNum
}

//是否首充
func (m *PlayerWelfareManager) IsFirstCharge() bool {
	if m.firstChargeObject == nil {
		return false
	}
	return true
}

//是否领取首充
func (m *PlayerWelfareManager) IsReceiveFirstCharge() bool {
	if m.firstChargeObject == nil {
		return false
	}
	return m.firstChargeObject.isReceive
}

//添加首充记录
func (m *PlayerWelfareManager) AddFirstCharge() (err error) {
	if m.firstChargeObject != nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	newObj := newPlayerFirstChargeObject(m.p)
	id, _ := idutil.GetId()
	newObj.id = id
	newObj.isReceive = false
	newObj.createTime = now
	newObj.SetModified()

	m.firstChargeObject = newObj
	gameevent.Emit(welfareeventtypes.EventTypeFinishFirstCharge, m.p, nil)
	return
}

//更新首充信息
func (m *PlayerWelfareManager) ReceiveFirstCharge() (err error) {
	if m.firstChargeObject == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.firstChargeObject.isReceive = true
	m.firstChargeObject.updateTime = now
	m.firstChargeObject.SetModified()
	return
}

//累计活动数据排行
func (m *PlayerWelfareManager) AddActivityNumRecordRecord(groupId int32, times int32) int32 {
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	now := global.GetGame().GetTimeService().Now()
	obj := m.getActivityNumRecord(groupId)
	if obj == nil {
		obj = newPlayerActivityNumRecordObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.groupId = groupId
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		m.activityNumRecordObject[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.times = 0
		obj.startTime = startTime
		obj.endTime = endTime
	}

	obj.updateTime = now
	obj.times += times
	obj.SetModified()

	return obj.times
}

//设置活动数据排行
func (m *PlayerWelfareManager) SetActivityNumRecordRecord(groupId int32, times int32) int32 {
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	now := global.GetGame().GetTimeService().Now()
	obj := m.getActivityNumRecord(groupId)
	if obj == nil {
		obj = newPlayerActivityNumRecordObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.groupId = groupId
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		m.activityNumRecordObject[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.times = 0
		obj.startTime = startTime
		obj.endTime = endTime
	}

	obj.updateTime = now
	obj.times = times
	obj.SetModified()

	return obj.times
}

//添加活动增长值(与 AddActivityNumRecordRecord 作用相同，表不同)
func (m *PlayerWelfareManager) AddActivityAddNumVal(groupId int32, val int32) int32 {
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	now := global.GetGame().GetTimeService().Now()
	obj := m.getActivityAddNumObj(groupId)
	if obj == nil {
		obj = newPlayerActivityAddNumObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.groupId = groupId
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		m.activityAddNumObject[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.addNum = 0
		obj.startTime = startTime
		obj.endTime = endTime
	}

	obj.updateTime = now
	obj.addNum += val
	obj.SetModified()

	return obj.addNum
}

//设置活动增长值(与 SetActivityNumRecordRecord 作用相同，表不同)
func (m *PlayerWelfareManager) SetActivityAddNumVal(groupId int32, val int32) int32 {
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	now := global.GetGame().GetTimeService().Now()
	obj := m.getActivityAddNumObj(groupId)
	if obj == nil {
		obj = newPlayerActivityAddNumObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.groupId = groupId
		obj.startTime = startTime
		obj.endTime = endTime
		obj.createTime = now
		m.activityAddNumObject[groupId] = obj
	}

	if obj.endTime != endTime {
		obj.addNum = 0
		obj.startTime = startTime
		obj.endTime = endTime
	}

	obj.updateTime = now
	obj.addNum = val
	obj.SetModified()

	return obj.addNum
}

// 添加活动开启邮件记录
func (m *PlayerWelfareManager) AddOpenMailRecord(groupId int32) {
	obj := m.getOpenMailObj(groupId)
	if obj != nil {
		return
	}

	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj = newPlayerActivityOpenMailObject(m.p)
	obj.id = id
	obj.group = groupId
	obj.createTime = now
	obj.SetModified()

	m.openActivityMailRecord[groupId] = obj
}

// 是否发开启邮件
func (m *PlayerWelfareManager) IsOpenMailRecord(groupId int32) bool {
	obj := m.getOpenMailObj(groupId)
	if obj != nil {
		return true
	}
	return false
}

func (m *PlayerWelfareManager) GetOpenActivity(groupId int32) *PlayerOpenActivityObject {

	obj := m.getOpenActivityByGorup(groupId)
	return obj
}

func (m *PlayerWelfareManager) GetOpenActivityIfNotCreate(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, groupId int32) *PlayerOpenActivityObject {

	obj := m.getOpenActivityByGorup(groupId)
	if obj == nil {
		obj = m.newOpenActivityObj(typ, subType, groupId)
		m.addOpenActivityObj(obj)
	}
	return obj
}

func (m *PlayerWelfareManager) UpdateObj(obj *PlayerOpenActivityObject) {
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(welfareeventtypes.EventTypeActivityDataChanged, m.p, obj)
}

// GM重置数据
func (m *PlayerWelfareManager) GMResetActivity(groupId int32) {
	obj := m.getOpenActivityByGorup(groupId)
	if obj == nil {
		return
	}

	obj.activityData = CreateEmptyOpenActivityData(obj.GetActivityType(), obj.GetActivitySubType())
	initHandler := GetInfoInitHandler(obj.GetActivityType(), obj.GetActivitySubType())
	if initHandler != nil {
		initHandler.InitInfo(obj)
	}

	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
}

// 运营活动对象复用
func (m *PlayerWelfareManager) checkReUseOpenActivityObj(obj *PlayerOpenActivityObject) {

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(obj.groupId)
	if obj.endTime != endTime {
		initHandler := GetInfoInitHandler(obj.activityType, obj.activitySubType)
		if initHandler == nil {
			return
		}

		obj.startTime = startTime
		obj.endTime = endTime
		initHandler.InitInfo(obj)

		obj.SetModified()
	}
}

func (m *PlayerWelfareManager) getOpenActivityByGorup(groupId int32) *PlayerOpenActivityObject {
	obj, ok := m.openActivityByGroupMap[groupId]
	if !ok {
		return nil
	}

	m.checkReUseOpenActivityObj(obj)
	m.RefreshActivityDataByGroupId(groupId)

	return obj
}

func (m *PlayerWelfareManager) getOpenMailObj(groupId int32) *PlayerActivityOpenMailObject {
	obj, ok := m.openActivityMailRecord[groupId]
	if !ok {
		return nil
	}

	return obj
}

func (m *PlayerWelfareManager) getOpenActivityCharge(groupId int32) *PlayerOpenActivityChargeObject {
	obj, ok := m.openActivityChargeByGroupMap[groupId]
	if !ok {
		return nil
	}

	return obj
}

func (m *PlayerWelfareManager) getActivityNumRecordObj(groupId int32) *PlayerActivityNumRecordObject {
	obj, ok := m.activityNumRecordObject[groupId]
	if !ok {
		return nil
	}

	return obj
}

func (m *PlayerWelfareManager) getActivityAddNumObj(groupId int32) *PlayerActivityAddNumObject {
	obj, ok := m.activityAddNumObject[groupId]
	if !ok {
		return nil
	}

	return obj
}

func (m *PlayerWelfareManager) getOpenActivityCost(groupId int32) *PlayerOpenActivityCostObject {
	obj, ok := m.openActivityCostByGroupMap[groupId]
	if !ok {
		return nil
	}

	return obj
}

func (m *PlayerWelfareManager) getActivityNumRecord(groupId int32) *PlayerActivityNumRecordObject {
	obj, ok := m.activityNumRecordObject[groupId]
	if !ok {
		return nil
	}

	return obj
}

func createPlayerWelfareDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerWelfareManager{}
	m.p = p
	m.openActivityByGroupMap = make(map[int32]*PlayerOpenActivityObject)
	m.openActivityChargeByGroupMap = make(map[int32]*PlayerOpenActivityChargeObject)
	m.openActivityCostByGroupMap = make(map[int32]*PlayerOpenActivityCostObject)
	m.activityNumRecordObject = make(map[int32]*PlayerActivityNumRecordObject)
	m.openActivityMailRecord = make(map[int32]*PlayerActivityOpenMailObject)
	m.activityAddNumObject = make(map[int32]*PlayerActivityAddNumObject)
	m.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerWelfareDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerWelfareDataManager))
}
