package player

import (
	"fgame/fgame/game/constant/constant"
	constantypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/unrealboss/dao"
	unrealbosseventtypes "fgame/fgame/game/unrealboss/event/types"
	"fgame/fgame/game/unrealboss/unrealboss"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家幻境BOSS管理器
type PlayerUnrealBossDataManager struct {
	p                      player.Player
	playerUnrealBossObject *PlayerUnrealBossObject
}

func (m *PlayerUnrealBossDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerUnrealBossDataManager) Load() (err error) {
	//加载玩家幻境BOSS信息
	mybossEntity, err := dao.GetUnrealBossDao().GetUnrealBossEntity(m.p.GetId())
	if err != nil {
		return
	}
	if mybossEntity == nil {
		m.initPlayerUnrealBossObject()
	} else {
		m.playerUnrealBossObject = NewPlayerUnrealBossObject(m.p)
		m.playerUnrealBossObject.FromEntity(mybossEntity)
	}
	return nil
}

//第一次初始化
func (m *PlayerUnrealBossDataManager) initPlayerUnrealBossObject() {
	o := NewPlayerUnrealBossObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.pilaoNum = 0
	o.buyPiLaoNum = 0
	o.createTime = now
	o.SetModified()

	m.playerUnrealBossObject = o
}

//加载后
func (m *PlayerUnrealBossDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerUnrealBossDataManager) Heartbeat() {
}

//击杀幻境boss
func (m *PlayerUnrealBossDataManager) KillPiLaoBoss(biologyId int32) {
	boss := unrealboss.GetUnrealBossService().GetUnrealBoss(biologyId)
	if boss == nil {
		return
	}
	initPilao := constant.GetConstantService().GetConstant(constantypes.ConstantTypePiLaoInitNum)
	curPilao := m.playerUnrealBossObject.pilaoNum
	leftPilao := initPilao - curPilao
	costPilaoNum := boss.GetBiologyTemplate().NeedPilao

	if leftPilao > costPilaoNum {
		leftPilao -= costPilaoNum
		costPilaoNum = 0
	} else {
		costPilaoNum -= leftPilao
		leftPilao = 0
	}

	curBuyPilao := m.playerUnrealBossObject.buyPiLaoNum
	if costPilaoNum > 0 {
		curBuyPilao -= costPilaoNum
	}

	if curBuyPilao < 0 {
		panic("unrealboss:购买的疲劳值不应为负数")
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerUnrealBossObject.pilaoNum = initPilao - leftPilao
	m.playerUnrealBossObject.buyPiLaoNum = curBuyPilao
	m.playerUnrealBossObject.updateTime = now
	m.playerUnrealBossObject.SetModified()

	total := m.GetCurPilaoNum()
	gameevent.Emit(unrealbosseventtypes.EventTypeUnrealPilaoChanged, m.p, total)
}

func (m *PlayerUnrealBossDataManager) GetPilaoBuyTimes() int32 {
	return m.playerUnrealBossObject.buyPiLaoTimes
}

func (m *PlayerUnrealBossDataManager) GetCurPilaoNum() int32 {
	initPilao := constant.GetConstantService().GetConstant(constantypes.ConstantTypePiLaoInitNum)
	curPilao := m.playerUnrealBossObject.pilaoNum
	leftPilao := initPilao - curPilao
	curBuyPilao := m.playerUnrealBossObject.buyPiLaoNum
	return leftPilao + curBuyPilao
}

func (m *PlayerUnrealBossDataManager) BuyPilaoNum(buyTimes int32) {
	m.RefreshPilao()

	if buyTimes <= 0 {
		panic("购买次数不能小于1")
	}
	num := constant.GetConstantService().GetConstant(constantypes.ConstantTypePiLaoBuyNum)
	addNum := num * buyTimes
	now := global.GetGame().GetTimeService().Now()

	m.playerUnrealBossObject.buyPiLaoNum += addNum
	m.playerUnrealBossObject.buyPiLaoTimes += buyTimes
	m.playerUnrealBossObject.updateTime = now
	m.playerUnrealBossObject.SetModified()

	total := m.GetCurPilaoNum()
	gameevent.Emit(unrealbosseventtypes.EventTypeUnrealPilaoChanged, m.p, total)
}

func (m *PlayerUnrealBossDataManager) RefreshPilao() {
	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameFive(now, m.playerUnrealBossObject.updateTime)
	if !isSame {
		m.playerUnrealBossObject.pilaoNum = 0
		m.playerUnrealBossObject.buyPiLaoNum = 0
		m.playerUnrealBossObject.buyPiLaoTimes = 0
		m.playerUnrealBossObject.updateTime = now
		m.playerUnrealBossObject.SetModified()

		total := m.GetCurPilaoNum()
		gameevent.Emit(unrealbosseventtypes.EventTypeUnrealPilaoChanged, m.p, total)
	}
}

// 使用醒神丹改变疲劳值
func (m *PlayerUnrealBossDataManager) SetPiLaoByXingShenDan(pilao int32) {
	now := global.GetGame().GetTimeService().Now()
	m.playerUnrealBossObject.buyPiLaoNum += pilao
	m.playerUnrealBossObject.updateTime = now
	m.playerUnrealBossObject.SetModified()

	total := m.GetCurPilaoNum()
	gameevent.Emit(unrealbosseventtypes.EventTypeUnrealPilaoChanged, m.p, total)
}

func (m *PlayerUnrealBossDataManager) GMSetPilao(pilao int32) {
	now := global.GetGame().GetTimeService().Now()
	m.playerUnrealBossObject.pilaoNum = 0
	m.playerUnrealBossObject.SetModified()
	m.playerUnrealBossObject.updateTime = now
	total := m.GetCurPilaoNum()
	gameevent.Emit(unrealbosseventtypes.EventTypeUnrealPilaoChanged, m.p, total)
}

func CreatePlayerUnrealBossDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerUnrealBossDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerUnrealBossDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerUnrealBossDataManager))
}
