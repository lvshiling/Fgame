package buff

import (
	"fgame/fgame/core/heartbeat"
	buffcommon "fgame/fgame/game/buff/common"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	bufftemplate "fgame/fgame/game/buff/template"

	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"math"
)

//特殊效果
type effectBuff struct {
	buffId    int32
	effectNum int64
}

func createEffectBuff(buffId int32, effectNum int64) *effectBuff {
	b := &effectBuff{
		buffId:    buffId,
		effectNum: effectNum,
	}
	return b
}

//buff管理器
type BuffDataManager struct {
	bo scene.BattleObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
	//列表
	buffMap map[int32]buffcommon.BuffObject
	//特殊效果 列表
	effectBuffMap map[scenetypes.BuffEffectType][]*effectBuff
}

func (m *BuffDataManager) GetBattleLimit() int64 {
	limitMask := int64(0)
	for _, b := range m.buffMap {
		buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(b.GetBuffId())
		limitMask |= buffTemplate.TypeLimit
	}
	return limitMask
}

//刷新buff
func (m *BuffDataManager) RefreshBuff() {
	for _, b := range m.buffMap {
		m.tickBuff(b)
	}
	return
}

//刷新buff
func (m *BuffDataManager) tickBuff(bo buffcommon.BuffObject) {
	tb := bo.(*buffObject)
	buffId := tb.GetBuffId()
	//判断是否到期了
	if tb.IsExpired() {
		//移除buff
		m.RemoveBuff(buffId)
		return
	}

	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	//不是定时触发的
	if !buffTemplate.IsTimer() {
		return
	}
	now := global.GetGame().GetTimeService().Now()

	if tb.LastTouchTime == 0 {
		elaspseTime := now - tb.StartTime
		if elaspseTime >= buffTemplate.FirstTouch {
			goto Touch
		}
		return
	} else {
		elaspseTime := now - tb.LastTouchTime
		if elaspseTime >= buffTemplate.Frequency {
			goto Touch
		}
		return
	}
Touch:
	tb.LastTouchTime = now

	//发送事件
	gameevent.Emit(buffeventtypes.EventTypeBuffTouch, m.bo, tb)
	return
}

//获取效果
func (m *BuffDataManager) getEffectBuffByBuffId(effectType scenetypes.BuffEffectType, buffId int32) (ebuff *effectBuff, index int32) {
	effectList := m.effectBuffMap[effectType]
	for i, effect := range effectList {
		if effect.buffId == buffId {
			ebuff = effect
			index = int32(i)
			return
		}
	}
	return nil, -1
}

func (m *BuffDataManager) addBuffEffect(effectType scenetypes.BuffEffectType, buffId int32, effectNum int64) {

	effect, index := m.getEffectBuffByBuffId(effectType, buffId)
	newEffect := createEffectBuff(buffId, effectNum)
	if effect == nil {
		m.effectBuffMap[effectType] = append(m.effectBuffMap[effectType], newEffect)
		return
	}

	//取代
	m.effectBuffMap[effectType][index] = newEffect

}

//移除效果
func (m *BuffDataManager) removeEffectByBuffId(effectType scenetypes.BuffEffectType, buffId int32) {
	_, index := m.getEffectBuffByBuffId(effectType, buffId)
	if index == -1 {
		return
	}
	m.effectBuffMap[effectType] = append(m.effectBuffMap[effectType][:index], m.effectBuffMap[effectType][index+1:]...)
	//TODO 效果移除事件
}

//移除buff
func (m *BuffDataManager) RemoveBuff(buffId int32) {
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemplate == nil {
		return
	}
	b := m.getBuff(buffTemplate.Group)
	if b == nil {
		return
	}

	delete(m.buffMap, buffTemplate.Group)
	//移除效果
	if buffTemplate.GetBuffEffectType() != scenetypes.BuffEffectTypeNone {
		m.removeEffectByBuffId(buffTemplate.GetBuffEffectType(), buffId)
	}
	//发送事件
	gameevent.Emit(buffeventtypes.EventTypeBuffRemove, m.bo, b)

	return
}

//添加buff
func (m *BuffDataManager) UpdateBuff(bo buffcommon.BuffObject) {
	buffId := bo.GetBuffId()
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemplate == nil {
		return
	}

	b := m.getBuff(buffTemplate.Group)
	ownerId := bo.GetOwnerId()

	group := bo.GetGroupId()
	startTime := bo.GetStartTime()
	duration := bo.GetDuration()
	times := bo.GetCulTime()
	if b == nil {
		bo := newBuffObject(ownerId, buffId, group, startTime, duration, times, nil)
		m.buffMap[bo.GroupId] = bo
	} else {
		b.Duration = duration
		b.CulTime = times
	}

}

//添加buff
func (m *BuffDataManager) AddBuff(buffId int32, ownerId int64, times int32, tianFuList []int32) (flag bool) {
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemplate == nil {
		return
	}
	timeDuration := bufftemplate.GetBuffTemplateService().GetBuffTimeDuration(buffId, tianFuList)
	now := global.GetGame().GetTimeService().Now()
	b := m.getBuff(buffTemplate.Group)
	var previousBuffTemplate *gametemplate.BuffTemplate
	//buff不存在
	if b == nil {
		//添加
		goto Replace
	}

	previousBuffTemplate = bufftemplate.GetBuffTemplateService().GetBuff(b.GetBuffId())
	if previousBuffTemplate == nil {
		//TODO panic
		return
	}

	//可以覆盖
	if buffTemplate.Replace != 0 {
		if previousBuffTemplate.Lev > buffTemplate.Lev {
			return
		}
		m.RemoveBuff(buffId)
		goto Replace
	} else {
		//叠加
		if b.GetCulTime() < buffTemplate.StackMax {
			//使用低等级的
			if buffTemplate.Lev < previousBuffTemplate.Lev {
				b.BuffId = buffId
			}
			b.CulTime += times
			//修改天赋
			b.TianFuList = tianFuList
			//更新时间
			if buffTemplate.StackType&int32(scenetypes.BuffStackTypeTime) != 0 {
				b.Duration += timeDuration * int64(times)
			}
			flag = true
			//发送事件
			gameevent.Emit(buffeventtypes.EventTypeBuffUpdate, m.bo, b)
			return
		}
		return
	}
Replace:
	duration := timeDuration
	if buffTemplate.Replace != 0 {
		times = 1
	} else {
		if times > buffTemplate.StackMax {
			times = buffTemplate.StackMax
		}
	}

	//更新时间
	if buffTemplate.StackType&int32(scenetypes.BuffStackTypeTime) != 0 {
		duration = timeDuration * int64(times)
	}
	bo := newBuffObject(ownerId, buffId, buffTemplate.Group, now, duration, times, tianFuList)

	m.buffMap[bo.GroupId] = bo
	//添加特殊效果
	if buffTemplate.GetBuffEffectType() != scenetypes.BuffEffectTypeNone {
		effectVal := int64(buffTemplate.EffectTypeBase)
		percentVal := int64(0)
		if buffTemplate.GetBuffEffectType() == scenetypes.BuffEffectTypeHuDun {
			maxHp := m.bo.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
			percentVal = int64(math.Ceil(float64(buffTemplate.EffectTypePercent) / float64(common.MAX_RATE) * float64(maxHp)))
		}

		effectVal += percentVal
		if effectVal > 0 {
			m.addBuffEffect(buffTemplate.GetBuffEffectType(), buffId, effectVal)
		}
	}
	flag = true

	//发送事件
	gameevent.Emit(buffeventtypes.EventTypeBuffAdd, m.bo, bo)

	return
}

func (m *BuffDataManager) TouchBuff(buffId int32) {
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemplate == nil {
		return
	}
	b := m.getBuff(buffTemplate.Group)
	if b == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	b.LastTouchTime = now

	//发送事件
	gameevent.Emit(buffeventtypes.EventTypeBuffTouch, m.bo, b)

	switch buffTemplate.GetTouchType() {
	case scenetypes.BuffTouchTypeImmediate:
		break
	case scenetypes.BuffTouchTypeTimer:
		break
	case scenetypes.BuffTouchTypeHurtedOther:
		break
	case scenetypes.BuffTouchTypeHurted,
		scenetypes.BuffTouchTypeBuff,
		scenetypes.BuffTouchTypeDead,
		scenetypes.BuffTouchTypeObjectDamage,
		scenetypes.BuffTouchTypeObjectDamageSelf,
		scenetypes.BuffTouchTypeJump:
		m.removeBuffOnce(buffId)
		break

	}
	return
}

func (m *BuffDataManager) removeBuffOnce(buffId int32) {
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemplate == nil {
		return
	}
	b := m.getBuff(buffTemplate.Group)
	if b == nil {
		return
	}
	if b.CulTime > 1 {
		b.CulTime -= 1

		//发送事件
		gameevent.Emit(buffeventtypes.EventTypeBuffUpdate, m.bo, b)

	}
	m.RemoveBuff(buffId)
}

//获取buff
func (m *BuffDataManager) getBuff(groupId int32) *buffObject {
	b, ok := m.buffMap[groupId]
	if !ok {
		return nil
	}
	tb, ok := b.(*buffObject)
	if !ok {
		return nil
	}
	return tb
}

//获取buff
func (m *BuffDataManager) GetBuff(groupId int32) buffcommon.BuffObject {
	b, ok := m.buffMap[groupId]
	if !ok {
		return nil
	}
	return b
}

//获取buff
func (m *BuffDataManager) GetBuffs() map[int32]buffcommon.BuffObject {
	return m.buffMap
}

//心跳
func (m *BuffDataManager) Heartbeat() {
	m.heartbeatRunner.Heartbeat()
}

//扣除效果值
func (m *BuffDataManager) CostEffectNum(effectType scenetypes.BuffEffectType, huDunVal int64) (remain int64) {

	remain = huDunVal
	effectList := m.effectBuffMap[effectType]
	if len(effectList) == 0 {
		return remain
	}

	removeBuffIdList := make([]int32, 0, len(effectList))
	for _, effect := range effectList {
		if remain < effect.effectNum {
			//TODO:zrc 暂时不修改,等策划决定
			remain = 0
			effect.effectNum -= remain
			break
		} else {
			remain -= effect.effectNum
			removeBuffIdList = append(removeBuffIdList, effect.buffId)
		}
	}
	m.effectBuffMap[effectType] = effectList
	for _, removeBuffId := range removeBuffIdList {
		m.RemoveBuff(removeBuffId)
	}

	costEffectNum := huDunVal - remain

	eventData := buffeventtypes.CraetBuffEffectCostEventData(effectType, costEffectNum)
	gameevent.Emit(buffeventtypes.EventTypeBuffEffectCost, m.bo, eventData)
	return remain
}

func (m *BuffDataManager) GetEffectNum(effectType scenetypes.BuffEffectType) int64 {
	effectVal := int64(0)
	effectList := m.effectBuffMap[effectType]
	for _, effect := range effectList {
		effectVal += effect.effectNum
	}
	return effectVal
}

func CreateBuffDataManager(bo scene.BattleObject) *BuffDataManager {
	m := &BuffDataManager{}
	m.bo = bo
	m.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	m.heartbeatRunner.AddTask(CreateBuffTask(m.bo))
	m.buffMap = make(map[int32]buffcommon.BuffObject)
	m.effectBuffMap = make(map[scenetypes.BuffEffectType][]*effectBuff)
	return m
}

func CreateBuffDataManagerWithBuffs(bo scene.BattleObject, buffs map[int32]buffcommon.BuffObject) *BuffDataManager {
	m := CreateBuffDataManager(bo)
	for groupId, buf := range buffs {
		m.buffMap[groupId] = copyFromBuffObject(buf)
	}
	return m
}

func CreateBuffDataManagerWithBuffList(bo scene.BattleObject, buffs []buffcommon.BuffObject) *BuffDataManager {
	m := CreateBuffDataManager(bo)
	for _, buf := range buffs {
		m.buffMap[buf.GetGroupId()] = copyFromBuffObject(buf)
	}
	return m
}
