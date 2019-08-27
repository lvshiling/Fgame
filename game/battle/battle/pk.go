package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	pkcommon "fgame/fgame/game/pk/common"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

//pk管理器
type PlayerPKManager struct {
	p            scene.Player
	pkValue      int32
	pkState      pktypes.PkState
	pkCamp       pktypes.PkCamp
	killNum      int32
	lastKillTime int64
	onlineTime   int64
	loginTime    int64 //上次扣除时间
}

func (m *PlayerPKManager) GetPkState() pktypes.PkState {
	return m.pkState
}

func (m *PlayerPKManager) GetPkCamp() pktypes.PkCamp {
	return m.pkCamp
}

func (m *PlayerPKManager) GetPkRedState() pktypes.PkRedState {
	return pktypes.PkRedStateFromValue(m.pkValue)
}

func (m *PlayerPKManager) GetPkValue() int32 {
	return m.pkValue
}

func (m *PlayerPKManager) GetKillNum() int32 {
	return m.killNum
}

func (m *PlayerPKManager) GetPkOnlineTime() int64 {
	return m.onlineTime
}

func (m *PlayerPKManager) GetPkLoginTime() int64 {
	return m.loginTime
}

func (m *PlayerPKManager) GetLastKillTime() int64 {
	return m.lastKillTime
}

//切换pk模式
func (m *PlayerPKManager) SwitchPkState(pkState pktypes.PkState, camp pktypes.PkCamp) (flag bool) {
	if !pkState.Valid() {
		panic(fmt.Errorf("pk:state [%d] invalid", pkState))
	}
	if m.pkState == pkState && m.pkCamp == camp {
		return
	}

	m.pkState = pkState
	m.pkCamp = camp
	m.p.ResetEnemy()
	//发送事件pk状态改变
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerPkStateSwitch, m.p, nil)
	flag = true
	return
}

const (
	maxPKValue = 80
)

//击杀
func (m *PlayerPKManager) Kill(white bool) {
	if white {
		firstKill := m.pkValue == 0
		if m.pkValue < maxPKValue {
			m.pkValue += 1
			m.killNum += 1
			now := global.GetGame().GetTimeService().Now()
			m.lastKillTime = now

			if firstKill {
				m.onlineTime = 0
				m.loginTime = now
			}
			m.pkValueChanged()
		}
	}
	//TODO 添加杀人数
	//TODO 发送事件
	return
}

func (m *PlayerPKManager) pkValueChanged() {
	//发送事件pk值改变
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerPkValueChanged, m.p, nil)
}

//刷新
func (m *PlayerPKManager) refresh() {
	if m.GetPkRedState() == pktypes.PkRedStateInit {
		return
	}

	clearTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypePkValueClearTime)
	now := global.GetGame().GetTimeService().Now()
	elapseTime := now - m.loginTime
	elapseTime += m.onlineTime
	clearNum := int32(elapseTime / int64(clearTime))
	if clearNum <= 0 {
		return
	}
	if clearNum >= m.pkValue {
		m.pkValue = 0
		m.onlineTime = 0
	} else {
		m.pkValue -= clearNum
		m.onlineTime = elapseTime - (int64(clearTime) * int64(clearNum))
		m.loginTime = now
	}

	m.pkValueChanged()
	return
}

func (m *PlayerPKManager) ReducePkValue(clearNum int32) {
	if clearNum < 0 {
		panic(fmt.Errorf("clearNum不能小于0, clearNum:%d", clearNum))
	}

	if clearNum >= m.pkValue {
		m.pkValue = 0
		m.onlineTime = 0
	} else {
		m.pkValue -= clearNum
	}

	m.pkValueChanged()
}

func (m *PlayerPKManager) Logout() {
	now := global.GetGame().GetTimeService().Now()
	m.onlineTime += now - m.loginTime
	m.pkValueChanged()
}

//心跳
func (m *PlayerPKManager) Heartbeat() {
	m.refresh()
}

func CreatePlayerPKManagerWithObject(p scene.Player, pkObj pkcommon.PlayerPkObject) *PlayerPKManager {
	m := &PlayerPKManager{}
	m.p = p
	m.killNum = pkObj.GetKillNum()
	m.lastKillTime = pkObj.GetLastKillTime()
	now := global.GetGame().GetTimeService().Now()
	m.loginTime = now
	m.pkValue = pkObj.GetPkValue()
	m.pkState = pkObj.GetPkState()
	m.pkCamp = pkObj.GetPkCamp()
	m.onlineTime = pkObj.GetOnlineTime()
	m.lastKillTime = pkObj.GetLastKillTime()
	return m
}

func CreatePlayerPKManager(p scene.Player) *PlayerPKManager {
	m := &PlayerPKManager{}
	m.p = p
	m.pkState = pktypes.PkStatePeach
	m.pkCamp = pktypes.PkCommonCampDefault
	return m

}
