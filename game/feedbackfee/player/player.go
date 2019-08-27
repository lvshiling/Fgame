package player

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/feedbackfee/dao"
	feedbackeventtypes "fgame/fgame/game/feedbackfee/event/types"
	feedbacktypes "fgame/fgame/game/feedbackfee/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家逆付费管理器
type PlayerFeedbackFeeManager struct {
	p                 player.Player
	feedbackfeeObject *PlayerFeedbackFeeObject      //逆付费数据
	recordList        []*PlayerFeedbackRecordObject //正在进行中的兑换记录
}

func (m *PlayerFeedbackFeeManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerFeedbackFeeManager) Load() (err error) {
	err = m.loadPlayerFeedbackObj()
	if err != nil {
		return
	}
	err = m.loadFeedbackRecordList()
	if err != nil {
		return
	}
	return nil
}

//加载逆付费数据
func (m *PlayerFeedbackFeeManager) loadPlayerFeedbackObj() (err error) {

	entity, err := dao.GetFeedbackFeeDao().GetFeedbackFeeEntity(m.p.GetId())
	if err != nil {
		return
	}

	if entity != nil {
		obj := newPlayerFeedbackFeeObject(m.p)
		obj.FromEntity(entity)
		m.feedbackfeeObject = obj
	} else {
		m.initFeedbackFeeObj()
	}
	return
}

//初始化付费数据
func (m *PlayerFeedbackFeeManager) initFeedbackFeeObj() {
	obj := newPlayerFeedbackFeeObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	obj.money = 0
	obj.createTime = now
	obj.SetModified()

	m.feedbackfeeObject = obj
}

//加载兑换过程记录
func (m *PlayerFeedbackFeeManager) loadFeedbackRecordList() (err error) {
	m.recordList = make([]*PlayerFeedbackRecordObject, 0, 8)
	recordEntityList, err := dao.GetFeedbackFeeDao().GetFeedbackRecordList(m.p.GetId(), int32(feedbacktypes.FeedbackStatusFinish))
	if err != nil {
		return
	}

	for _, recordEntity := range recordEntityList {
		recordObj := newPlayerFeedbackRecordObject(m.p)
		err = recordObj.FromEntity(recordEntity)
		if err != nil {
			return
		}
		m.recordList = append(m.recordList, recordObj)
	}
	return
}

//加载后
func (m *PlayerFeedbackFeeManager) AfterLoad() (err error) {
	m.refresh()
	return
}

//心跳
func (m *PlayerFeedbackFeeManager) Heartbeat() {
}

func (m *PlayerFeedbackFeeManager) refresh() {
	now := global.GetGame().GetTimeService().Now()
	flag, _ := timeutils.IsSameDay(now, m.feedbackfeeObject.useTime)
	if flag {
		return
	}
	m.feedbackfeeObject.useTime = now
	m.feedbackfeeObject.todayUseNum = 0
	m.feedbackfeeObject.updateTime = now
	m.feedbackfeeObject.SetModified()
}

func (m *PlayerFeedbackFeeManager) GetFeedbackFeeInfo() *PlayerFeedbackFeeObject {
	m.refresh()
	return m.feedbackfeeObject
}

//获取正在进行的列表
func (m *PlayerFeedbackFeeManager) GetCurrentRecordList() []*PlayerFeedbackRecordObject {
	return m.recordList
}

//是否可以兑换
func (m *PlayerFeedbackFeeManager) IfCanExchange(money int32) bool {
	if money <= 0 {
		return false
	}
	m.refresh()
	//库存不够
	if m.feedbackfeeObject.money < money {
		return false
	}
	exchangeLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCodeExchangeLimit)
	todayRemain := exchangeLimit - m.feedbackfeeObject.todayUseNum
	if todayRemain < money {
		return false
	}

	return true
}

func (m *PlayerFeedbackFeeManager) getRecord(id int64) *PlayerFeedbackRecordObject {
	for _, record := range m.recordList {
		if record.id == id {
			return record
		}
	}
	return nil
}

func (m *PlayerFeedbackFeeManager) removeRecord(id int64) {
	index := -1
	for tempIndex, record := range m.recordList {
		if record.id == id {
			index = tempIndex
			break
		}
	}
	if index < 0 {
		return
	}
	m.recordList = append(m.recordList[:index], m.recordList[index+1:]...)
}

//是否可以兑换
func (m *PlayerFeedbackFeeManager) Exchange(money int32) *PlayerFeedbackRecordObject {
	if !m.IfCanExchange(money) {
		return nil
	}
	now := global.GetGame().GetTimeService().Now()
	m.feedbackfeeObject.todayUseNum += money
	m.feedbackfeeObject.useTime = now
	m.feedbackfeeObject.updateTime = now
	m.feedbackfeeObject.SetModified()
	reason := commonlog.FeedbackLogReasonExchange
	reasonText := reason.String()
	flag := m.costMoney(money, reason, reasonText)
	if !flag {
		panic(fmt.Errorf("feedback:兑换应该成功"))
	}
	//添加记录
	recordObj := newPlayerFeedbackRecordObject(m.p)
	recordObj.id, _ = idutil.GetId()
	recordObj.money = money
	recordObj.status = feedbacktypes.FeedbackStatusInit
	validTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCodeExpireTime))
	recordObj.expiredTime = now + validTime
	recordObj.createTime = now
	recordObj.typ = feedbacktypes.FeedbackFeeTypeCash
	recordObj.SetModified()
	m.recordList = append(m.recordList, recordObj)
	return recordObj
}

func (m *PlayerFeedbackFeeManager) ExchangeGold(money int32, gold int64) *PlayerFeedbackRecordObject {
	if !m.IfCanExchange(money) {
		return nil
	}
	now := global.GetGame().GetTimeService().Now()
	m.feedbackfeeObject.todayUseNum += money
	m.feedbackfeeObject.useTime = now
	m.feedbackfeeObject.updateTime = now
	m.feedbackfeeObject.SetModified()
	reason := commonlog.FeedbackLogReasonExchangeGold
	reasonText := fmt.Sprintf(reason.String(), gold)
	flag := m.costMoney(money, reason, reasonText)
	if !flag {
		panic(fmt.Errorf("feedback:兑换应该成功"))
	}

	//添加记录
	recordObj := newPlayerFeedbackRecordObject(m.p)
	recordObj.id, _ = idutil.GetId()
	recordObj.typ = feedbacktypes.FeedbackFeeTypeGold
	recordObj.money = money
	recordObj.status = feedbacktypes.FeedbackStatusFinish
	validTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCodeExpireTime))
	recordObj.expiredTime = now + validTime
	recordObj.createTime = now
	recordObj.SetModified()
	return recordObj
}

func (m *PlayerFeedbackFeeManager) AddMoney(money int32, reason commonlog.FeedbackLogReason, reasonText string) {
	m.addMoney(money, reason, reasonText)
	m.addTotalMoney(money)
}

//兑换过期
func (m *PlayerFeedbackFeeManager) Expire(id int64) *PlayerFeedbackRecordObject {
	record := m.getRecord(id)
	if record == nil {
		return nil
	}
	flag := record.Expire()
	if !flag {
		return nil
	}
	m.refresh()
	m.removeRecord(id)
	//退还金额
	reason := commonlog.FeedbackLogReasonExchangeRefund
	reasonText := fmt.Sprintf(commonlog.FeedbackLogReasonExchangeRefund.String(), record.GetCode())
	m.addMoney(record.GetMoney(), reason, reasonText)
	m.feedbackfeeObject.todayUseNum -= record.GetMoney()
	if m.feedbackfeeObject.todayUseNum < 0 {
		m.feedbackfeeObject.todayUseNum = 0
	}
	now := global.GetGame().GetTimeService().Now()
	m.feedbackfeeObject.updateTime = now
	m.feedbackfeeObject.SetModified()
	return record
}

func (m *PlayerFeedbackFeeManager) CodeGenerate(id int64, code string) *PlayerFeedbackRecordObject {
	record := m.getRecord(id)
	if record == nil {
		return nil
	}
	flag := record.Code(code)
	if !flag {
		return nil
	}
	return record
}

func (m *PlayerFeedbackFeeManager) CodeFinish(id int64) *PlayerFeedbackRecordObject {
	record := m.getRecord(id)
	if record == nil {
		return nil
	}
	flag := record.Finish()
	if !flag {
		return nil
	}
	m.removeRecord(id)
	return record
}

//获取初始化
func (m *PlayerFeedbackFeeManager) GetInitRecordList() (recordList []*PlayerFeedbackRecordObject) {
	for _, recordObj := range m.recordList {
		if recordObj.status == feedbacktypes.FeedbackStatusInit {
			recordList = append(recordList, recordObj)
		}
	}
	return recordList
}

//获取当前兑换记录
func (m *PlayerFeedbackFeeManager) GetCurrentRecord() *PlayerFeedbackRecordObject {
	if len(m.recordList) <= 0 {
		return nil
	}
	return m.recordList[0]
}

//gm修改库存
func (m *PlayerFeedbackFeeManager) GMSetMoney(money int32) {
	addMoney := m.feedbackfeeObject.money - money
	if addMoney == 0 {
		return
	}
	if addMoney < 0 {
		reason := commonlog.FeedbackLogReasonGM
		reasonText := commonlog.FeedbackLogReasonGM.String()
		m.addMoney(-addMoney, reason, reasonText)
		m.addTotalMoney(-addMoney)
		return
	} else {
		reason := commonlog.FeedbackLogReasonGM
		reasonText := commonlog.FeedbackLogReasonGM.String()
		m.costMoney(addMoney, reason, reasonText)
		return
	}

	return
}

//扣除库存
func (m *PlayerFeedbackFeeManager) costMoney(money int32, reason commonlog.FeedbackLogReason, reasonText string) bool {
	if m.feedbackfeeObject.money < money {
		return false
	}
	beforeMoney := m.feedbackfeeObject.money
	now := global.GetGame().GetTimeService().Now()
	m.feedbackfeeObject.money -= money
	m.feedbackfeeObject.updateTime = now
	m.feedbackfeeObject.SetModified()
	curMoney := m.feedbackfeeObject.money
	eventData := feedbackeventtypes.CreateFeedbackfeeExchangeLogEventData(curMoney, beforeMoney, -money, reason, reasonText)
	gameevent.Emit(feedbackeventtypes.EventTypeFeedbackfeeExchangeLog, m.p, eventData)
	//记录日志
	return true
}

func (m *PlayerFeedbackFeeManager) addMoney(money int32, reason commonlog.FeedbackLogReason, reasonText string) {
	beforeMoney := m.feedbackfeeObject.money
	now := global.GetGame().GetTimeService().Now()
	m.feedbackfeeObject.money += money
	m.feedbackfeeObject.updateTime = now
	m.feedbackfeeObject.SetModified()
	curMoney := m.feedbackfeeObject.money
	eventData := feedbackeventtypes.CreateFeedbackfeeExchangeLogEventData(curMoney, beforeMoney, money, reason, reasonText)
	gameevent.Emit(feedbackeventtypes.EventTypeFeedbackfeeExchangeLog, m.p, eventData)
}

//添加历史总获得
func (m *PlayerFeedbackFeeManager) addTotalMoney(money int32) {
	now := global.GetGame().GetTimeService().Now()
	m.feedbackfeeObject.totalGetMoney += int64(money)
	m.feedbackfeeObject.updateTime = now
	m.feedbackfeeObject.SetModified()
}

func createPlayerFeedbackFeeDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerFeedbackFeeManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerFeedbackFeeDataManager))
}
