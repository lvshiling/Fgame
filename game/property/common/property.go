package property

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
	"math"

	"github.com/willf/bitset"
)

//属性值
type PropertySegmentBase struct {
	properties map[uint]int64
	bs         *bitset.BitSet
}

//TODO 优化对象池
func (psb *PropertySegmentBase) GetChangedTypes() map[int32]int64 {
	changed := make(map[int32]int64)
	for i, valid := psb.bs.NextSet(0); valid; i, valid = psb.bs.NextSet(i + 1) {
		changed[int32(i)] = psb.Get(i)
	}
	return changed
}

func (psb *PropertySegmentBase) IsChanged() bool {
	return psb.bs.Count() != 0
}

func (psb *PropertySegmentBase) Set(typ uint, val int64) {
	oldVal, exist := psb.properties[typ]
	if exist {
		if oldVal == val {
			return
		}
	}
	//设置标记位 改变
	psb.bs.Set(typ)
	psb.properties[typ] = val
}

func (psb *PropertySegmentBase) Get(typ uint) int64 {
	val, exist := psb.properties[typ]
	if !exist {
		return 0
	}
	return val

}
func (psb *PropertySegmentBase) IsTypeChanged(typ uint) bool {
	return psb.bs.Test(typ)
}

func (ps *PropertySegmentBase) Clear() {
	for prop, _ := range ps.properties {
		ps.Set(prop, int64(0))
	}
}

func (psb *PropertySegmentBase) ResetChanged() {
	psb.bs.ClearAll()
}

func NewPropertySegmentBase() *PropertySegmentBase {
	psb := &PropertySegmentBase{}
	psb.properties = make(map[uint]int64)
	psb.bs = bitset.New(64)
	return psb
}

//基础属性
type BasePropertySegment struct {
	*PropertySegmentBase
}

func (p *BasePropertySegment) Set(pt propertytypes.BasePropertyType, val int64) {
	p.PropertySegmentBase.Set(uint(pt), val)
}

func (p *BasePropertySegment) Get(pt propertytypes.BasePropertyType) int64 {
	return p.PropertySegmentBase.Get(uint(pt))
}

func (p *BasePropertySegment) IsTypeChanged(typ propertytypes.BasePropertyType) bool {
	return p.PropertySegmentBase.IsTypeChanged(uint(typ))
}

func NewBasePropertySegment() *BasePropertySegment {
	p := &BasePropertySegment{
		PropertySegmentBase: NewPropertySegmentBase(),
	}
	return p
}

//战斗
type BattlePropertySegment struct {
	*PropertySegmentBase
}

func (p *BattlePropertySegment) Set(pt propertytypes.BattlePropertyType, val int64) {
	p.PropertySegmentBase.Set(uint(pt), int64(val))
}

func (p *BattlePropertySegment) Get(pt propertytypes.BattlePropertyType) int64 {
	return p.PropertySegmentBase.Get(uint(pt))
}

func (p *BattlePropertySegment) IsTypeChanged(typ propertytypes.BattlePropertyType) bool {
	return p.PropertySegmentBase.IsTypeChanged(uint(typ))
}

func (p *BattlePropertySegment) GetForce() int64 {
	force := float64(0)
	forceTemplate := constant.GetConstantService().GetForceTemplate()
	for k, v := range forceTemplate.GetAllForceProperty() {
		val := p.Get(k)
		tempForce := float64(val) / common.MAX_RATE * float64(v)
		force += tempForce
	}
	return int64(math.Floor(force)) + p.Get(propertytypes.BattlePropertyTypeForce)
}

func NewBattlePropertySegment() *BattlePropertySegment {
	p := &BattlePropertySegment{
		PropertySegmentBase: NewPropertySegmentBase(),
	}
	return p
}

//战斗
type SystemPropertySegment struct {
	//基础属性
	baseProperty *PropertySegmentBase
	//基础属性万分比
	basePropertyPercent *PropertySegmentBase
	//外部万分比
	propertySegmentMap map[PropertyEffectorType]*PropertySegmentBase
	//更新外部万分比
	externalPropertyPercent *PropertySegmentBase
	//全局属性万分比
	globalPropertyPercent *PropertySegmentBase
	//全局属性值
	globalProperty *PropertySegmentBase
	//模块属性
	moduleProperty *PropertySegmentBase
}

func StringFromPropertySegmentBase(base *PropertySegmentBase) string {
	returnStr := ""
	for k, v := range base.properties {
		returnStr += fmt.Sprintf("%s=%d", propertytypes.BattlePropertyType(k).String(), v)
		returnStr += ","
	}
	return returnStr
}

func (p *SystemPropertySegment) String() string {
	return fmt.Sprintf("基础:%s\n, 基础万分比:%s\n, 全局万分比:%s\n, 全局:%s\n",
		StringFromPropertySegmentBase(p.baseProperty),
		StringFromPropertySegmentBase(p.basePropertyPercent),
		StringFromPropertySegmentBase(p.globalPropertyPercent),
		StringFromPropertySegmentBase(p.globalProperty))
}

func (p *SystemPropertySegment) Clear() {
	//清除基础属性
	p.baseProperty.Clear()
	//清除基础属性万分比
	p.basePropertyPercent.Clear()
	//清除全局属性万分比
	p.globalPropertyPercent.Clear()
	//清除
	p.globalProperty.Clear()
}

func (p *SystemPropertySegment) SetBase(pt propertytypes.BattlePropertyType, val int64) {
	p.baseProperty.Set(uint(pt), int64(val))
}

func (p *SystemPropertySegment) GetBase(pt propertytypes.BattlePropertyType) int64 {
	return p.baseProperty.Get(uint(pt))
}

func (p *SystemPropertySegment) SetBasePercent(pt propertytypes.BattlePropertyType, val int64) {
	p.basePropertyPercent.Set(uint(pt), int64(val))
}

func (p *SystemPropertySegment) GetBasePercent(pt propertytypes.BattlePropertyType) int64 {
	return p.basePropertyPercent.Get(uint(pt))
}

func (p *SystemPropertySegment) GetExternalPropertySegment(t PropertyEffectorType) *PropertySegmentBase {
	segment, exist := p.propertySegmentMap[t]
	if exist {
		return segment
	}
	segment = NewPropertySegmentBase()
	p.propertySegmentMap[t] = segment
	return segment
}

func (p *SystemPropertySegment) SetGlobalPercent(pt propertytypes.BattlePropertyType, val int64) {
	p.globalPropertyPercent.Set(uint(pt), int64(val))
}

func (p *SystemPropertySegment) GetGlobalPercent(pt propertytypes.BattlePropertyType) int64 {
	return p.globalPropertyPercent.Get(uint(pt))
}

func (p *SystemPropertySegment) SetGlobal(pt propertytypes.BattlePropertyType, val int64) {
	p.globalProperty.Set(uint(pt), int64(val))
}

func (p *SystemPropertySegment) GetGlobal(pt propertytypes.BattlePropertyType) int64 {
	return p.globalProperty.Get(uint(pt))
}

func (p *SystemPropertySegment) GetPercent(pt propertytypes.BattlePropertyType) int64 {
	return p.externalPropertyPercent.Get(uint(pt))
}

func (p *SystemPropertySegment) UpdatePercent() {
	//合并万分比
	for pt := propertytypes.MinBattlePropertyType; pt <= propertytypes.MaxBattlePropertyType; pt++ {
		total := int64(0)
		for _, propertySegment := range p.propertySegmentMap {
			val := propertySegment.Get(uint(pt))
			total += val
		}
		p.externalPropertyPercent.Set(uint(pt), total)
	}

}

func (p *SystemPropertySegment) UpdateModuleProperty() {
	p.moduleProperty.Clear()
	//合并系统属性
	for pt := propertytypes.MinBattlePropertyType; pt <= propertytypes.MaxBattlePropertyType; pt++ {
		total := float64(0)
		val := p.GetBase(pt)
		internalPercent := p.GetBasePercent(pt)
		externalPercent := p.GetPercent(pt)
		percent := internalPercent + externalPercent
		extra := p.GetGlobal(pt)
		total += float64(val) + float64(val)*(float64(percent)/float64(common.MAX_RATE)) + float64(extra)
		totalInt := int64(math.Ceil(total))
		p.moduleProperty.Set(uint(pt), totalInt)
	}
}

func (p *SystemPropertySegment) GetForce() int64 {
	force := float64(0)
	forceTemplate := constant.GetConstantService().GetForceTemplate()
	for k, v := range forceTemplate.GetAllForceProperty() {
		val := p.moduleProperty.Get(uint(k))
		tempForce := float64(val) / common.MAX_RATE * float64(v)
		force += tempForce
	}
	return int64(math.Floor(force)) + p.moduleProperty.Get(uint(propertytypes.BattlePropertyTypeForce))
}

func NewSystemPropertySegment() *SystemPropertySegment {
	segment := &SystemPropertySegment{}
	segment.baseProperty = NewPropertySegmentBase()
	segment.basePropertyPercent = NewPropertySegmentBase()
	segment.externalPropertyPercent = NewPropertySegmentBase()
	segment.globalPropertyPercent = NewPropertySegmentBase()
	segment.propertySegmentMap = make(map[PropertyEffectorType]*PropertySegmentBase)
	segment.globalProperty = NewPropertySegmentBase()
	segment.moduleProperty = NewPropertySegmentBase()
	return segment
}

type PropertyEffectorType interface {
	EffectorType() uint32
	// IsPercent() bool
}

//战斗属性
type BattlePropertyGroup struct {
	//各个系统战斗属性
	propertySegmentMap map[PropertyEffectorType]*SystemPropertySegment
	//系统属性
	systemProperty *BattlePropertySegment
	//系统万分比
	systemPercentProperty *BattlePropertySegment
	//最终属性
	finalProperty *BattlePropertySegment
}

func (bp *BattlePropertyGroup) GetPropertySegment(pet PropertyEffectorType) *SystemPropertySegment {
	segment, exist := bp.propertySegmentMap[pet]
	if exist {
		return segment
	}
	segment = NewSystemPropertySegment()
	bp.propertySegmentMap[pet] = segment
	return segment
}

func (bp *BattlePropertyGroup) UpdateProperty() {

	//合并万分比
	for pt := propertytypes.MinBattlePropertyType; pt <= propertytypes.MaxBattlePropertyType; pt++ {
		total := int64(0)
		for _, propertySegment := range bp.propertySegmentMap {
			val := propertySegment.GetGlobalPercent(pt)
			total += val
		}
		bp.systemPercentProperty.Set(pt, total)
	}

	for _, propertySegment := range bp.propertySegmentMap {
		propertySegment.UpdatePercent()
	}

	//合并系统属性
	for pt := propertytypes.MinBattlePropertyType; pt <= propertytypes.MaxBattlePropertyType; pt++ {
		total := float64(0)
		for _, propertySegment := range bp.propertySegmentMap {
			val := propertySegment.GetBase(pt)
			internalPercent := propertySegment.GetBasePercent(pt)
			externalPercent := propertySegment.GetPercent(pt)
			allPercent := bp.systemPercentProperty.Get(pt)
			percent := internalPercent + externalPercent + allPercent
			extra := propertySegment.GetGlobal(pt)
			total += float64(val) + float64(val)*(float64(percent)/float64(common.MAX_RATE)) + float64(extra)
		}

		totalInt := int64(math.Ceil(total))
		bp.finalProperty.Set(pt, totalInt)
	}

}

//重置最终属性
func (bp *BattlePropertyGroup) ResetChanged() {
	bp.finalProperty.ResetChanged()
}

//是否最终属性改变了
func (bp *BattlePropertyGroup) IsChanged() bool {
	return bp.finalProperty.IsChanged()
}

func (bp *BattlePropertyGroup) GetChangedTypes() map[int32]int64 {
	return bp.finalProperty.GetChangedTypes()
}

func (bp *BattlePropertyGroup) IsTypeChanged(typ propertytypes.BattlePropertyType) bool {
	return bp.finalProperty.IsTypeChanged(typ)
}

func (bp *BattlePropertyGroup) Get(typ propertytypes.BattlePropertyType) int64 {
	val := bp.finalProperty.Get(typ)
	return val
}

//创建战斗属性组
func NewBattlePropertyGroup() *BattlePropertyGroup {
	bpg := &BattlePropertyGroup{
		propertySegmentMap: make(map[PropertyEffectorType]*SystemPropertySegment),
		finalProperty:      NewBattlePropertySegment(),
		//系统属性
		systemProperty: NewBattlePropertySegment(),
		//系统万分比
		systemPercentProperty: NewBattlePropertySegment(),
	}
	return bpg
}
