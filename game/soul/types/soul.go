package types

//帝魂种类
type SoulKindType int32

const (
	//攻击型
	SoulKindTypeAttack SoulKindType = 1 + iota
	//防御型
	SoulKindTypeDefense
	//辅助型
	SoulKindTypeAssist
)

func (skt SoulKindType) Valid() bool {
	switch skt {
	case SoulKindTypeAttack,
		SoulKindTypeDefense,
		SoulKindTypeAssist:
		return true
	}
	return false
}

//帝魂类型
type SoulType int32

const (
	//魂·重华
	SoulTypeChongHua SoulType = iota
	//魂·大禹
	SoulTypeDaYu
	//魂·伊祁
	SoulTypeYiQi
	//魂·帝俊
	SoulTypeDiJun
	//魂·颛顼
	SoulTypeZhuanXu
	//魂·神农
	SoulTypeShenNong
	//魂·轩辕
	SoulTypeXuanYuan
	//魂·女娲
	SoulTypeNuWa
	//魂·伏羲
	SoulTypeFuXi
)

func (st SoulType) Valid() bool {
	switch st {
	case SoulTypeChongHua,
		SoulTypeDaYu,
		SoulTypeYiQi,
		SoulTypeDiJun,
		SoulTypeZhuanXu,
		SoulTypeShenNong,
		SoulTypeXuanYuan,
		SoulTypeNuWa,
		SoulTypeFuXi:
		return true
	}
	return false
}

const (
	SoulTypeMin = SoulTypeChongHua
	SoulTypeMax = SoulTypeFuXi
)

//效果类型
type SoulEffectType int32

const (
	//增加帝魂技能伤害
	SoulEffectTypeAddDamage SoulEffectType = iota
)

func (set SoulEffectType) Valid() bool {
	switch set {
	case SoulEffectTypeAddDamage:
		return true
	}
	return false
}
