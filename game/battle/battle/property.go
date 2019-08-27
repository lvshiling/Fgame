package battle

import (
	"fgame/fgame/core/heartbeat"
	coretypes "fgame/fgame/core/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	propertycommon "fgame/fgame/game/property/common"
	propertytypes "fgame/fgame/game/property/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"math"
)

type SystemPropertyManager struct {
	bo            scene.BattleObject
	finalProperty *propertycommon.BattlePropertySegment
}

//获取系统属性
func (m *SystemPropertyManager) GetSystemBattleProperty(pt propertytypes.BattlePropertyType) int64 {
	return m.finalProperty.Get(pt)
}

func (m *SystemPropertyManager) GetAllSystemBattleProperties() map[int32]int64 {
	properties := make(map[int32]int64)
	for typ := propertytypes.MinBattlePropertyType; typ <= propertytypes.MaxBattlePropertyType; typ++ {
		val := m.finalProperty.Get(typ)
		properties[int32(typ)] = val
	}
	return properties
}

func (m *SystemPropertyManager) UpdateSystemBattleProperty(properties map[int32]int64) {
	for k, v := range properties {
		pt := propertytypes.BattlePropertyType(k)
		if !pt.IsValid() {
			continue
		}
		m.finalProperty.Set(pt, v)
	}
}

func (m *SystemPropertyManager) GetSystemBattlePropertyChangedTypesAndReset() map[int32]int64 {
	battleChanged := m.finalProperty.GetChangedTypes()
	m.finalProperty.ResetChanged()

	return battleChanged
}

func CreateSystemPropertyManager(bo scene.BattleObject) *SystemPropertyManager {
	m := &SystemPropertyManager{
		bo: bo,
	}
	m.finalProperty = propertycommon.NewBattlePropertySegment()
	return m
}

func CreateSystemPropertyManagerWithData(bo scene.BattleObject, ps map[int32]int64) *SystemPropertyManager {
	m := &SystemPropertyManager{
		bo: bo,
	}
	m.finalProperty = propertycommon.NewBattlePropertySegment()
	m.UpdateSystemBattleProperty(ps)
	return m
}

//战斗属性管理器
type BattlePropertyManager struct {
	bo scene.BattleObject
	//原始血量
	originHP int64
	//原始体力
	originTP int64
	//首次加载
	firstLoad bool
	//血量
	currentHP int64
	//体力
	currentTP int64
	//战力
	power int64
	//buff属性
	buffProperty *propertycommon.BattlePropertySegment
	//buff万分比
	buffPercentProperty *propertycommon.BattlePropertySegment
	//最终属性
	finalProperty *propertycommon.BattlePropertySegment
	//额外的移动速度
	extraSpeed int64
	//最大血量
	maxHP int64
	//最大体力
	maxTP int64
}

//获取当前hp
func (m *BattlePropertyManager) GetMaxHP() int64 {
	return m.maxHP
}

//获取当前hp
func (m *BattlePropertyManager) GetMaxTP() int64 {
	return m.maxTP
}

//获取当前hp
func (m *BattlePropertyManager) GetHP() int64 {
	return m.currentHP
}

//获取当前hp
func (m *BattlePropertyManager) CostHP(hp int64, attackId int64) bool {
	if hp <= 0 {
		return false
	}

	remainHp := m.bo.CostEffectNum(scenetypes.BuffEffectTypeHuDun, hp)
	if remainHp <= 0 {
		return false
	}

	m.currentHP -= remainHp
	if m.currentHP <= 0 {
		m.currentHP = 0
	}

	if m.currentHP == 0 {
		return true
	}
	return false
}

func (m *BattlePropertyManager) Dead(attackId int64) bool {
	return false
}

func (m *BattlePropertyManager) ResetHP() {
	m.recover()
}

//获取当前hp
func (m *BattlePropertyManager) AddHP(hp int64) int64 {
	if hp <= 0 {
		return 0
	}

	oldHp := m.currentHP
	currentHp := oldHp + hp
	hpMax := m.bo.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	if currentHp > hpMax {
		currentHp = hpMax
	}
	m.currentHP = currentHp

	return m.currentHP - oldHp
}

//获取当前hp
func (m *BattlePropertyManager) GetTP() int64 {
	return m.currentTP
}

//获取当前hp
func (m *BattlePropertyManager) CostTP(tp int64) bool {
	if tp <= 0 {
		return false
	}
	if m.currentTP < tp {
		return false
	}
	m.currentTP -= tp

	return true
}

//获取当前hp
func (m *BattlePropertyManager) AddTP(tp int64) int64 {
	if tp <= 0 {
		return 0
	}

	oldTp := m.currentTP
	currentTp := oldTp + tp
	tpMax := m.bo.GetBattleProperty(propertytypes.BattlePropertyTypeMaxTP)
	if currentTp > tpMax {
		currentTp = tpMax
	}
	m.currentTP = currentTp

	return m.currentTP - oldTp
}

func (m *BattlePropertyManager) Reborn() {
	m.recover()
}

func (m *BattlePropertyManager) recover() {
	hpMax := m.bo.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	m.currentHP = hpMax
}

func (m *BattlePropertyManager) calculateProperty() {
	hpMaxOld := m.finalProperty.Get(propertytypes.BattlePropertyTypeMaxHP)
	tpMaxOld := m.finalProperty.Get(propertytypes.BattlePropertyTypeMaxTP)

	//获取系统属性
	for pt := propertytypes.MinBattlePropertyType; pt <= propertytypes.MaxBattlePropertyType; pt++ {
		total := int64(0)
		total += m.bo.GetSystemBattleProperty(pt)
		buffPercent := m.buffPercentProperty.Get(pt)
		buffVal := m.buffProperty.Get(pt)

		extra := int64(math.Ceil(float64(total)*(float64(buffPercent)/float64(common.MAX_RATE)))) + buffVal
		total += extra
		if pt == propertytypes.BattlePropertyTypeMoveSpeed {
			total += m.bo.GetExtraSpeed()
		}
		if pt.MoreThanZero() && total <= 1 {
			total = int64(1)
		}
		if pt.NoLessZero() && total < 0 {
			total = int64(0)
		}
		m.finalProperty.Set(pt, total)
	}

	// if pt.Penalize() {
	// 	totalInt64 = int64(math.Ceil(float64(total) * (float64(rate) / float64(common.MAX_RATE))))
	// } else {
	// 	totalInt64 = int64(math.Ceil(float64(total)))
	// }
	//更新惩罚机制
	// ppdm.battlePropertyGroup.UpdateProperty(common.MAX_RATE - ppdm.p.GetPkRedState().Penalize())

	//更新基础属性
	hpChanged := m.finalProperty.IsTypeChanged(propertytypes.BattlePropertyTypeMaxHP)
	if hpChanged {
		hpMaxNow := m.finalProperty.Get(propertytypes.BattlePropertyTypeMaxHP)
		hpNow := hpMaxNow
		m.maxHP = hpNow
		if hpMaxOld != 0 {
			hpNow = int64(math.Ceil(float64(hpMaxNow) * (float64(m.currentHP) / float64(hpMaxOld))))
		}
		m.currentHP = hpNow
		if m.firstLoad {
			if m.originHP > 0 {
				m.currentHP = m.originHP
			}
			if m.currentHP > hpMaxNow {
				m.currentHP = hpMaxNow
			}
		}
	}

	tpChanged := m.finalProperty.IsTypeChanged(propertytypes.BattlePropertyTypeMaxTP)
	if tpChanged {
		tpMaxNow := m.finalProperty.Get(propertytypes.BattlePropertyTypeMaxTP)
		tpNow := tpMaxNow
		m.maxTP = tpMaxNow
		if tpMaxOld != 0 {
			tpNow = int64(math.Ceil(float64(tpMaxNow) * (float64(m.currentTP)) / float64(tpMaxOld)))
		}
		m.currentTP = tpNow
		if m.firstLoad {
			if m.originTP > 0 {
				m.currentTP = m.originTP
			}
			if m.currentTP > tpMaxNow {
				m.currentTP = tpMaxNow
			}
		}
	}
	if m.firstLoad {
		m.firstLoad = false
	}
}

//更新buff属性
func (m *BattlePropertyManager) UpdateBuffProperty() {
	m.buffProperty.Clear()
	buffPropertyEffect(m.bo, m.buffProperty)
	m.buffPercentProperty.Clear()
	buffPercentPropertyEffect(m.bo, m.buffPercentProperty)
}

func (m *BattlePropertyManager) Calculate() {
	m.calculateProperty()
}

//获取战斗属性变更
func (m *BattlePropertyManager) GetBattlePropertyChangedTypesAndReset() (battleChanged map[int32]int64) {
	battleChanged = m.finalProperty.GetChangedTypes()
	m.finalProperty.ResetChanged()
	return
}

func (m *BattlePropertyManager) IsTypeChanged(typ propertytypes.BattlePropertyType) bool {
	return m.finalProperty.IsTypeChanged(typ)

}

//获取战斗属性
func (m *BattlePropertyManager) GetBattleProperty(battlePropertyType propertytypes.BattlePropertyType) int64 {
	return m.finalProperty.Get(battlePropertyType)
}

//获取战斗属性
func (m *BattlePropertyManager) GetAllBattleProperties() map[int32]int64 {
	properties := make(map[int32]int64)
	for typ := propertytypes.MinBattlePropertyType; typ <= propertytypes.MaxBattlePropertyType; typ++ {
		val := m.finalProperty.Get(typ)
		properties[int32(typ)] = val
	}
	return properties
}

func (m *BattlePropertyManager) GetForce() int64 {
	return m.power
}

func CreateBattlePropertyManager(bo scene.BattleObject, hp int64, tp int64) *BattlePropertyManager {
	m := &BattlePropertyManager{
		bo:        bo,
		originHP:  hp,
		originTP:  tp,
		firstLoad: true,
	}
	m.buffProperty = propertycommon.NewBattlePropertySegment()
	m.buffPercentProperty = propertycommon.NewBattlePropertySegment()
	m.finalProperty = propertycommon.NewBattlePropertySegment()
	return m
}

//战斗属性管理器
type PlayerBattlePropertyManager struct {
	*BattlePropertyManager
	p scene.Player
	//死亡
	isDead bool
	//死亡时间
	deadTime int64
	//战力
	power  int64
	runner heartbeat.HeartbeatTaskRunner
}

func (m *PlayerBattlePropertyManager) Heartbeat() {
	m.runner.Heartbeat()
	if m.isDead {
		s := m.p.GetScene()
		if s == nil {
			return
		}

		mapTemplate := s.MapTemplate()
		now := global.GetGame().GetTimeService().Now()
		elapse := now - m.deadTime
		if elapse >= int64(mapTemplate.ResurrectionTime) {
			//自动回城复活
			gameevent.Emit(battleeventtypes.EventTypeBattlePlayerAutoReborn, m.p, nil)
		}

	}
}

//获取当前hp
func (m *PlayerBattlePropertyManager) CostHP(hp int64, attackId int64) bool {
	if hp <= 0 {
		return false
	}
	if m.isDead {
		return false
	}
	dead := m.BattlePropertyManager.CostHP(hp, attackId)

	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerHPChanged, m.p, nil)
	if dead {
		//TODO 修改
		//清除pvp
		m.p.ClearPvpBattle()
		//清除技能管理器
		m.p.ClearAllSkillAction()
		//死亡
		m.p.GuaJiDead()
		m.isDead = true
		m.deadTime = global.GetGame().GetTimeService().Now()

		return true
	}
	return dead
}

func (m *PlayerBattlePropertyManager) Dead(attackId int64) bool {
	if !m.isDead {
		return false
	}
	//死亡事件
	gameevent.Emit(sceneeventtypes.EventTypeBattleObjectDead, m.p, attackId)
	return true
}

//获取当前hp
func (m *PlayerBattlePropertyManager) AddHP(hp int64) int64 {
	if hp <= 0 {
		return 0
	}
	addHp := m.BattlePropertyManager.AddHP(hp)
	if addHp <= 0 {
		return addHp
	}

	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerHPChanged, m.p, nil)
	return addHp
}

//重置hp
func (m *PlayerBattlePropertyManager) ResetHP() {
	m.recover()
	return
}

//获取当前hp
func (m *PlayerBattlePropertyManager) CostTP(tp int64) bool {
	if tp <= 0 {
		return false
	}
	flag := m.BattlePropertyManager.CostTP(tp)
	if !flag {
		return false
	}
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerTPChanged, m.p, nil)
	return true
}

//获取当前hp
func (m *PlayerBattlePropertyManager) AddTP(tp int64) int64 {
	if tp <= 0 {
		return 0
	}
	addTp := m.BattlePropertyManager.AddTP(tp)
	if addTp <= 0 {
		return addTp
	}
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerTPChanged, m.p, nil)
	return addTp
}

//是否死亡
func (m *PlayerBattlePropertyManager) IsDead() bool {
	return m.isDead
}

func (m *PlayerBattlePropertyManager) GetDeadTime() int64 {
	return m.deadTime
}

func (m *PlayerBattlePropertyManager) Reborn(pos coretypes.Position) {
	m.isDead = false
	m.recover()
	gameevent.Emit(sceneeventtypes.EventTypeBattleObjectReborn, m.p, pos)
	m.p.ResetEnemy()
	m.p.GuaJiIdle()
}

func (m *PlayerBattlePropertyManager) recover() {
	m.BattlePropertyManager.recover()

	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerHPChanged, m.p, nil)
	//发送事件
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerMaxHPChanged, m.p, nil)
}

func (m *PlayerBattlePropertyManager) calculateProperty() {
	//发送事件
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerHPChanged, m.p, nil)
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerMaxHPChanged, m.p, nil)

	speedChanged := m.finalProperty.IsTypeChanged(propertytypes.BattlePropertyTypeMoveSpeed)
	if speedChanged {
		gameevent.Emit(battleeventtypes.EventTypeBattlePlayerSpeedChanged, m.p, nil)
	}
}

//重新计算战力
// func (m *PlayerBattlePropertyManager) updateForce() {
// 	force := float64(0)
// 	forceTemplate := constant.GetConstantService().GetForceTemplate()
// 	for k, v := range forceTemplate.GetAllForceProperty() {
// 		val := m.BattlePropertyManager.GetBattleProperty(k)
// 		if k == propertytypes.BattlePropertyTypeMoveSpeed {
// 			initSpeed := constant.GetConstantService().GetConstant(constantypes.ConstantTypeInitMoveSpeed)
// 			val -= int64(initSpeed)
// 		}
// 		if k == propertytypes.BattlePropertyTypeHit {
// 			hit := constant.GetConstantService().GetConstant(constantypes.ConstantTypeInitalHit)
// 			val -= int64(hit)
// 		}
// 		tempForce := float64(val) / common.MAX_RATE * float64(v)
// 		force += tempForce
// 	}
// 	power := int64(math.Floor(force)) + m.BattlePropertyManager.GetBattleProperty(propertytypes.BattlePropertyTypeForce)
// 	m.power = power
// 	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerForceChanged, m.p, nil)
// }

//更新buff属性
// func (m *PlayerBattlePropertyManager) UpdateBuffProperty() {
// 	m.buffProperty.Clear()
// 	buffPropertyEffect(m.p, m.buffProperty)
// 	m.buffPercentProperty.Clear()
// 	buffPercentPropertyEffect(m.p, m.buffPercentProperty)
// }

func (m *PlayerBattlePropertyManager) Calculate() {

	m.BattlePropertyManager.Calculate()
	// m.updateForce()

	//发送事件
	if m.BattlePropertyManager.IsTypeChanged(propertytypes.BattlePropertyTypeMaxHP) {
		gameevent.Emit(battleeventtypes.EventTypeBattlePlayerHPChanged, m.p, nil)
		gameevent.Emit(battleeventtypes.EventTypeBattlePlayerMaxHPChanged, m.p, nil)
	}

	// speedChanged := m.finalProperty.IsTypeChanged(propertytypes.BattlePropertyTypeMoveSpeed)
	// if speedChanged {
	if m.BattlePropertyManager.IsTypeChanged(propertytypes.BattlePropertyTypeMoveSpeed) {
		gameevent.Emit(battleeventtypes.EventTypeBattlePlayerSpeedChanged, m.p, nil)
	}
	// }

	//发送属性变化
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerPropertyChanged, m.p, nil)

}

func (m *PlayerBattlePropertyManager) GetForce() int64 {
	return m.power
}

func (m *PlayerBattlePropertyManager) UpdateForce(force int64) {
	if m.power != force {
		m.power = force
		gameevent.Emit(battleeventtypes.EventTypeBattlePlayerForceChanged, m.p, nil)
	}
}

func CreatePlayerBattlePropertyManager(p scene.Player, hp int64, tp int64, power int64) *PlayerBattlePropertyManager {
	m := &PlayerBattlePropertyManager{
		p:     p,
		power: power,
	}
	m.BattlePropertyManager = CreateBattlePropertyManager(p, hp, tp)
	m.runner = heartbeat.NewHeartbeatTaskRunner()
	m.runner.AddTask(CreatePropertyTask(p))
	return m
}

//buff 万分比
func buffPercentPropertyEffect(bo scene.BattleObject, prop *propertycommon.BattlePropertySegment) {
	buffs := bo.GetBuffs()
	for _, b := range buffs {
		buffId := b.GetBuffId()

		buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
		if buffTemplate.GetTouchType() != scenetypes.BuffTouchTypeImmediate {
			continue
		}

		for typ, val := range buffTemplate.GetBattlePropertyPercentMap() {
			total := prop.Get(typ)
			tempVal := val
			for _, buffDongTaiId := range b.GetTianFuList() {
				buffDongTaiTemplate := bufftemplate.GetBuffTemplateService().GetBuffDongTai(buffDongTaiId)
				replaceVal := buffDongTaiTemplate.GetBattlePropertyPercentMap()[typ]
				if replaceVal == 0 {
					continue
				}
				tempVal = replaceVal
				break
			}
			//特殊处理移动速度用乘法
			if typ == propertytypes.BattlePropertyTypeMoveSpeed {
				percent := float64(common.MAX_RATE+total) / float64(common.MAX_RATE)
				percent *= math.Pow(float64(common.MAX_RATE+tempVal)/float64(common.MAX_RATE), float64(b.GetCulTime()))
				total = int64(math.Ceil(percent*common.MAX_RATE)) - common.MAX_RATE
			} else {
				if buffTemplate.StackType&int32(scenetypes.BuffStackTypeEffect) != 0 {
					total += tempVal * int64(b.GetCulTime())
				} else {
					total += tempVal
				}
			}
			prop.Set(typ, total)
		}
	}
}

//buff 万分比
func buffPropertyEffect(bo scene.BattleObject, prop *propertycommon.BattlePropertySegment) {
	buffs := bo.GetBuffs()
	for _, b := range buffs {
		buffId := b.GetBuffId()
		buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
		if buffTemplate.GetTouchType() != scenetypes.BuffTouchTypeImmediate {
			continue
		}
		for typ, val := range buffTemplate.GetBattlePropertyMap() {
			total := prop.Get(typ)
			tempVal := val
			for _, buffDongTaiId := range b.GetTianFuList() {
				buffDongTaiTemplate := bufftemplate.GetBuffTemplateService().GetBuffDongTai(buffDongTaiId)
				replaceVal := buffDongTaiTemplate.GetBattlePropertyMap()[typ]
				if replaceVal == 0 {
					continue
				}
				tempVal = replaceVal
				break
			}
			if buffTemplate.StackType&int32(scenetypes.BuffStackTypeEffect) != 0 {
				total += tempVal * int64(b.GetCulTime())
			} else {
				total += tempVal
			}
			prop.Set(typ, total)
		}
	}
}
