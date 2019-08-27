package types

type XinFaType int32

const (
	//生命
	XinFaTypeLife XinFaType = 1 + iota
	//攻击
	XinFaTypeAttack
	//防御
	XinFaTypeDefense
	//暴击
	XinFaTypeCrit
	//坚韧
	XinFaTypeTough
	//破格
	XinFaTypeAbnormality
	//格挡
	XinFaTypeBlock
	//中毒抗性
	XinFaTypePoisoningRes
	//破甲抗性
	XinFaTypeSunderArmorRes
	//冰冻抗性
	XinFaTypeFrozenRes
	//昏迷抗性
	XinFaTypeComaRes
	//失明抗性
	XinFaTypeBlindRes
	//傀儡抗性
	XinFaTypePuppetRes
	//缴械抗性
	XinFaTypeDisarmRes
	//枯竭抗性
	XinFaTypeDriedUpRes
	//虚弱抗性
	XinFaTypeWeakRes
)

func (xft XinFaType) Valid() bool {
	switch xft {
	case XinFaTypeLife,
		XinFaTypeAttack,
		XinFaTypeDefense,
		XinFaTypeCrit,
		XinFaTypeTough,
		XinFaTypeAbnormality,
		XinFaTypeBlock,
		XinFaTypePoisoningRes,
		XinFaTypeSunderArmorRes,
		XinFaTypeFrozenRes,
		XinFaTypeComaRes,
		XinFaTypeBlindRes,
		XinFaTypePuppetRes,
		XinFaTypeDisarmRes,
		XinFaTypeDriedUpRes,
		XinFaTypeWeakRes:
		return true
	default:
		return false
	}
}
