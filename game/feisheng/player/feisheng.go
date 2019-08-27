package player

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/feisheng/dao"
	feishengeventtypes "fgame/fgame/game/feisheng/event/types"
	feishengtemplate "fgame/fgame/game/feisheng/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

const (
	GONGDE_MAX = 99999999999999
)

//玩家飞升管理器
type PlayerFeiShengDataManager struct {
	p                           player.Player
	playerFeiShengObject        *PlayerFeiShengObject
	playerFeiShengReceiveObject *PlayerFeiShengReceiveObject
}

func (m *PlayerFeiShengDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerFeiShengDataManager) Load() (err error) {
	//加载玩家飞升信息
	feishengEntity, err := dao.GetFeiShengDao().GetFeiShengEntity(m.p.GetId())
	if err != nil {
		return
	}
	if feishengEntity == nil {
		m.initPlayerFeiShengObject()
	} else {
		m.playerFeiShengObject = NewPlayerFeiShengObject(m.p)
		m.playerFeiShengObject.FromEntity(feishengEntity)
	}
	//加载玩家飞升信息
	feishengReceiveEntity, err := dao.GetFeiShengDao().GetFeiShengReceiveEntity(m.p.GetId())
	if err != nil {
		return
	}
	if feishengReceiveEntity == nil {
		m.initPlayerFeiShengReceiveObject()
	} else {
		m.playerFeiShengReceiveObject = NewPlayerFeiShengReceiveObject(m.p)
		m.playerFeiShengReceiveObject.FromEntity(feishengReceiveEntity)
	}

	return nil
}

//第一次初始化
func (m *PlayerFeiShengDataManager) initPlayerFeiShengObject() {
	o := NewPlayerFeiShengObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.feiLevel = 0
	o.addRate = 0
	o.gongDeNum = 0
	o.leftPotential = 0
	o.tiZhi = 0
	o.liDao = 0
	o.jinGu = 0
	o.createTime = now
	m.playerFeiShengObject = o
	o.SetModified()
}

//第一次初始化
func (m *PlayerFeiShengDataManager) initPlayerFeiShengReceiveObject() {
	o := NewPlayerFeiShengReceiveObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id

	o.createTime = now
	m.playerFeiShengReceiveObject = o
	o.SetModified()
}

//加载后
func (m *PlayerFeiShengDataManager) AfterLoad() (err error) {
	m.refreshFeiShengReceiveObject()

	return nil
}

//第一次初始化
func (m *PlayerFeiShengDataManager) refreshFeiShengReceiveObject() {

	now := global.GetGame().GetTimeService().Now()
	flag, _ := timeutils.IsSameFive(now, m.playerFeiShengReceiveObject.updateTime)
	if flag {
		return
	}
	m.playerFeiShengReceiveObject.updateTime = now
	m.playerFeiShengReceiveObject.num = 0

	m.playerFeiShengReceiveObject.SetModified()
}

func (m *PlayerFeiShengDataManager) ReceiveFeiSheng() {
	m.refreshFeiShengReceiveObject()
	now := global.GetGame().GetTimeService().Now()
	m.playerFeiShengReceiveObject.num += 1
	m.playerFeiShengReceiveObject.updateTime = now
	m.playerFeiShengReceiveObject.SetModified()
}

func (m *PlayerFeiShengDataManager) GetFeiShengReceiveNum() int32 {
	m.refreshFeiShengReceiveObject()
	return m.playerFeiShengReceiveObject.num
}

//心跳
func (m *PlayerFeiShengDataManager) Heartbeat() {

}

//飞升等级
func (m *PlayerFeiShengDataManager) GetFeiShengLevel() int32 {
	return m.playerFeiShengObject.feiLevel
}

//飞升信息
func (m *PlayerFeiShengDataManager) GetFeiShengInfo() *PlayerFeiShengObject {
	return m.playerFeiShengObject
}

//是否食概率丹
func (m *PlayerFeiShengDataManager) IsFullRate() bool {
	nextFeiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(m.playerFeiShengObject.feiLevel + 1)
	if nextFeiTemplate == nil {
		return true
	}
	curRate := m.playerFeiShengObject.addRate + nextFeiTemplate.Rate
	return curRate >= common.MAX_RATE
}

//食概率丹
func (m *PlayerFeiShengDataManager) EatRateDan(eatNum int32) {
	if eatNum <= 0 {
		panic(fmt.Errorf("飞升：食丹数量不能为0,eatNum:%d", eatNum))
	}
	feiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(m.playerFeiShengObject.feiLevel)
	addRate := feiTemplate.AddRate * eatNum

	now := global.GetGame().GetTimeService().Now()
	m.playerFeiShengObject.addRate += addRate
	m.playerFeiShengObject.updateTime = now
	m.playerFeiShengObject.SetModified()
	return
}

//重置潜能
func (m *PlayerFeiShengDataManager) ResetQn() {
	remainQn := m.playerFeiShengObject.tiZhi
	remainQn += m.playerFeiShengObject.jinGu
	remainQn += m.playerFeiShengObject.liDao

	now := global.GetGame().GetTimeService().Now()
	m.playerFeiShengObject.leftPotential += remainQn
	m.playerFeiShengObject.tiZhi = 0
	m.playerFeiShengObject.liDao = 0
	m.playerFeiShengObject.jinGu = 0
	m.playerFeiShengObject.updateTime = now
	m.playerFeiShengObject.SetModified()

	return
}

//设置潜能
func (m *PlayerFeiShengDataManager) SaveQn(ti, li, gu int32) {
	if ti+li+gu <= 0 {
		panic(fmt.Errorf("保存失败，潜能设置不能为负数,ti:%d,li:%d,gu:%d", ti, li, gu))
	}

	remainQn := ti
	remainQn += li
	remainQn += gu

	now := global.GetGame().GetTimeService().Now()
	m.playerFeiShengObject.leftPotential -= remainQn
	m.playerFeiShengObject.tiZhi += ti
	m.playerFeiShengObject.liDao += li
	m.playerFeiShengObject.jinGu += gu
	m.playerFeiShengObject.updateTime = now
	m.playerFeiShengObject.SetModified()

	return
}

//飞升渡劫
func (m *PlayerFeiShengDataManager) FeiShengDuJie(isSuccess bool) {
	nextFeiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(m.playerFeiShengObject.feiLevel + 1)
	if nextFeiTemplate == nil {
		return
	}
	beforeGongDe := m.playerFeiShengObject.gongDeNum
	beforeLevel := m.playerFeiShengObject.feiLevel
	if isSuccess {
		m.playerFeiShengObject.feiLevel += 1
		m.playerFeiShengObject.leftPotential += nextFeiTemplate.QnAdd
		m.playerFeiShengObject.addRate = 0

		gameevent.Emit(feishengeventtypes.EventTypePlayerFeiSheng, m.p, m.playerFeiShengObject.feiLevel)
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerFeiShengObject.gongDeNum -= int64(nextFeiTemplate.GongDe)
	m.playerFeiShengObject.updateTime = now
	m.playerFeiShengObject.SetModified()

	reason := commonlog.FeiShengLogReasonDuJi
	reasonText := fmt.Sprintf(reason.String(), isSuccess)
	eventData := feishengeventtypes.CreatePlayerFeiShengLogEventData(beforeLevel, beforeGongDe, reason, reasonText)
	gameevent.Emit(feishengeventtypes.EventTypePlayerFeiShengLog, m.p, eventData)

	return
}

//飞升散功
func (m *PlayerFeiShengDataManager) FeiShengSanGong(costExp int64) int64 {
	feiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(m.playerFeiShengObject.feiLevel)
	if feiTemplate == nil {
		return 0
	}

	addGongDe := costExp / int64(feiTemplate.GongdeRatio)
	m.AddGongDe(addGongDe)

	gameevent.Emit(feishengeventtypes.EventTypePlayerSanGong, m.p, costExp)
	return addGongDe
}

//添加功德
func (m *PlayerFeiShengDataManager) AddGongDe(num int64) {
	now := global.GetGame().GetTimeService().Now()
	oldNum := m.playerFeiShengObject.gongDeNum
	newNum := oldNum + num
	if newNum > GONGDE_MAX {
		newNum = GONGDE_MAX
	}
	m.playerFeiShengObject.gongDeNum = int64(newNum)
	m.playerFeiShengObject.updateTime = now
	m.playerFeiShengObject.SetModified()
}

//GM设置功德
func (m *PlayerFeiShengDataManager) GMSetGongDe(num int64) {
	now := global.GetGame().GetTimeService().Now()
	if num > GONGDE_MAX {
		num = GONGDE_MAX
	}
	m.playerFeiShengObject.gongDeNum = int64(num)
	m.playerFeiShengObject.updateTime = now
	m.playerFeiShengObject.SetModified()
	return
}

func CreatePlayerFeiShengDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerFeiShengDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerFeiShengDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerFeiShengDataManager))
}
