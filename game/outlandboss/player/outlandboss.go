package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/outlandboss/dao"
	outlandbosseventtypes "fgame/fgame/game/outlandboss/event/types"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家外域BOSS管理器
type PlayerOutlandBossDataManager struct {
	p                       player.Player
	playerOutlandBossObject *PlayerOutlandBossObject
}

func (m *PlayerOutlandBossDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerOutlandBossDataManager) Load() (err error) {
	//加载玩家外域BOSS信息
	mybossEntity, err := dao.GetOutlandBossDao().GetOutlandBossEntity(m.p.GetId())
	if err != nil {
		return
	}
	if mybossEntity == nil {
		m.initPlayerOutlandBossObject()
	} else {
		m.playerOutlandBossObject = NewPlayerOutlandBossObject(m.p)
		m.playerOutlandBossObject.FromEntity(mybossEntity)
	}
	return nil
}

//第一次初始化
func (m *PlayerOutlandBossDataManager) initPlayerOutlandBossObject() {
	o := NewPlayerOutlandBossObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.zhuoqiNum = 0
	o.createTime = now
	o.SetModified()

	m.playerOutlandBossObject = o
}

//加载后
func (m *PlayerOutlandBossDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerOutlandBossDataManager) Heartbeat() {
}

//击杀外域boss
func (m *PlayerOutlandBossDataManager) KillZhuoQiBoss(biologyId int32) {
	boss := outlandboss.GetOutlandBossService().GetOutlandBoss(biologyId)
	if boss == nil {
		return
	}
	oldZhuoQi := m.playerOutlandBossObject.zhuoqiNum
	addZhuoQi := boss.GetBiologyTemplate().ZhuoQi
	total := oldZhuoQi + addZhuoQi

	now := global.GetGame().GetTimeService().Now()
	m.playerOutlandBossObject.zhuoqiNum = total
	m.playerOutlandBossObject.updateTime = now
	m.playerOutlandBossObject.SetModified()

	gameevent.Emit(outlandbosseventtypes.EventTypeZhuoQiChanged, m.p, total)
}

func (m *PlayerOutlandBossDataManager) GetCurZhuoQiNum() int32 {
	return m.playerOutlandBossObject.zhuoqiNum
}

func (m *PlayerOutlandBossDataManager) CanUseJingLingDan() bool {
	return m.playerOutlandBossObject.zhuoqiNum > 0
}

func (m *PlayerOutlandBossDataManager) RefreshZhuoQi() {
	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameFive(now, m.playerOutlandBossObject.updateTime)
	if !isSame {
		m.playerOutlandBossObject.zhuoqiNum = 0
		m.playerOutlandBossObject.updateTime = now
		m.playerOutlandBossObject.SetModified()

		gameevent.Emit(outlandbosseventtypes.EventTypeZhuoQiChanged, m.p, m.playerOutlandBossObject.zhuoqiNum)
	}
}

// 使用净灵丹改变浊气值
func (m *PlayerOutlandBossDataManager) SetZhuoQiByJingLingDan(zhuoqiNum int32) {
	now := global.GetGame().GetTimeService().Now()
	if m.playerOutlandBossObject.zhuoqiNum <= zhuoqiNum {
		m.playerOutlandBossObject.zhuoqiNum = 0
	} else {
		m.playerOutlandBossObject.zhuoqiNum -= zhuoqiNum
	}

	m.playerOutlandBossObject.updateTime = now
	m.playerOutlandBossObject.SetModified()

	gameevent.Emit(outlandbosseventtypes.EventTypeZhuoQiChanged, m.p, zhuoqiNum)
}

func (m *PlayerOutlandBossDataManager) GMSetZhuoQi(zhuoqiNum int32) {
	now := global.GetGame().GetTimeService().Now()
	m.playerOutlandBossObject.zhuoqiNum = zhuoqiNum
	m.playerOutlandBossObject.SetModified()
	m.playerOutlandBossObject.updateTime = now

	gameevent.Emit(outlandbosseventtypes.EventTypeZhuoQiChanged, m.p, zhuoqiNum)
}

func CreatePlayerOutlandBossDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerOutlandBossDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerOutlandBossDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerOutlandBossDataManager))
}
