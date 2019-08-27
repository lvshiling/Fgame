package types

//基础属性
type BasePropertyType int32

const (
	BasePropertyTypeHP BasePropertyType = 1 + iota
	BasePropertyTypeTP
	BasePropertyTypeLevel
	BasePropertyTypeExp
	BasePropertyTypeSilver
	BasePropertyTypeGold
	BasePropertyTypeBindGold
	BasePropertyTypeEvil
	BasePropertyTypeZhuanSheng
	BasePropertyTypeEvilOnlineTime
	BasePropertyTypeCharm
	BasePropertyTypeGoldYuanLevel
	BasePropertyTypeGoldYuanExp
)

var (
	basePropertyTypeMap = map[BasePropertyType]string{
		BasePropertyTypeHP:             "生命",
		BasePropertyTypeTP:             "体力",
		BasePropertyTypeLevel:          "等级",
		BasePropertyTypeExp:            "经验",
		BasePropertyTypeSilver:         "银两",
		BasePropertyTypeGold:           "元宝",
		BasePropertyTypeBindGold:       "绑定元宝",
		BasePropertyTypeEvil:           "罪恶",
		BasePropertyTypeZhuanSheng:     "转生",
		BasePropertyTypeEvilOnlineTime: "红名值累计在线时间",
		BasePropertyTypeCharm:          "魅力值",
		BasePropertyTypeGoldYuanLevel:  "元神等级",
		BasePropertyTypeGoldYuanExp:    "元神经验",
	}
)

func (pt BasePropertyType) IsValid() bool {
	switch pt {
	case BasePropertyTypeHP,
		BasePropertyTypeTP,
		BasePropertyTypeLevel,
		BasePropertyTypeExp,
		BasePropertyTypeSilver,
		BasePropertyTypeGold,
		BasePropertyTypeBindGold,
		BasePropertyTypeEvil,
		BasePropertyTypeZhuanSheng,
		BasePropertyTypeEvilOnlineTime,
		BasePropertyTypeCharm,
		BasePropertyTypeGoldYuanLevel,
		BasePropertyTypeGoldYuanExp:
		return true

	}
	return false
}

func (pt BasePropertyType) String() string {
	return basePropertyTypeMap[pt]
}

const (
	MaxCharmLimit = 999999
)

//战斗属性
type BattlePropertyType int32

const (
	BattlePropertyTypeMaxHP BattlePropertyType = 1 + iota
	BattlePropertyTypeAttack
	BattlePropertyTypeDefend
	BattlePropertyTypeMaxTP
	BattlePropertyTypeMoveSpeed
	BattlePropertyTypeCrit
	BattlePropertyTypeTough
	BattlePropertyTypeAbnormality
	BattlePropertyTypeBlock
	BattlePropertyTypeDodge
	BattlePropertyTypeHit
	BattlePropertyTypeHuanYunAttack
	BattlePropertyTypeHuanYunDef
	BattlePropertyTypeBingDongRes
	BattlePropertyTypePoJiaRes
	BattlePropertyTypeKuiLeiRes
	BattlePropertyTypeKuJieRes
	BattlePropertyTypeShiMingRes
	BattlePropertyTypeHunMiRes
	BattlePropertyTypeXuRuoRes
	BattlePropertyTypeJiaoXieRes
	BattlePropertyTypeZhongDuRes
	BattlePropertyTypeDamageAdd
	BattlePropertyTypeDamageAddPercent
	BattlePropertyTypeDamageDefend
	BattlePropertyTypeDamageDefendPercent
	BattlePropertyTypeForce
	BattlePropertyTypeFanTan
	BattlePropertyTypeFanTanPercent
	BattlePropertyTypeCritRatePercent
	BattlePropertyTypeCritHarmPercent
	BattlePropertyTypeHitRatePercent
	BattlePropertyTypeDodgeRatePercent
	BattlePropertyTypeSpellCdPercent
	BattlePropertyTypeAddExp
	BattlePropertyTypeBlockRatePercent
	BattlePropertyTypeZhuoShaoRes
	BattlePropertyTypeJianSuRes
	BattlePropertyTypeDingShenRes
)

const (
	MinBattlePropertyType = BattlePropertyTypeMaxHP
	MaxBattlePropertyType = BattlePropertyTypeDingShenRes
)

var (
	propertyTypeMap = map[BattlePropertyType]string{
		BattlePropertyTypeMaxHP:               "最大生命",
		BattlePropertyTypeAttack:              "攻击",
		BattlePropertyTypeDefend:              "防御",
		BattlePropertyTypeMaxTP:               "最大体力",
		BattlePropertyTypeMoveSpeed:           "移动速度",
		BattlePropertyTypeCrit:                "暴击",
		BattlePropertyTypeTough:               "坚韧",
		BattlePropertyTypeAbnormality:         "破格",
		BattlePropertyTypeBlock:               "格挡",
		BattlePropertyTypeDodge:               "闪避",
		BattlePropertyTypeHit:                 "命中",
		BattlePropertyTypeHuanYunAttack:       "混元伤害",
		BattlePropertyTypeHuanYunDef:          "混元抗性",
		BattlePropertyTypeBingDongRes:         "冰冻抗性",
		BattlePropertyTypePoJiaRes:            "破解抗性",
		BattlePropertyTypeKuiLeiRes:           "傀儡抗性",
		BattlePropertyTypeKuJieRes:            "枯竭抗性",
		BattlePropertyTypeShiMingRes:          "失明抗性",
		BattlePropertyTypeHunMiRes:            "昏迷抗性",
		BattlePropertyTypeXuRuoRes:            "虚弱抗性",
		BattlePropertyTypeJiaoXieRes:          "缴械抗性",
		BattlePropertyTypeZhongDuRes:          "中毒抗性",
		BattlePropertyTypeDamageAdd:           "增伤值",
		BattlePropertyTypeDamageAddPercent:    "增伤万分比",
		BattlePropertyTypeDamageDefend:        "减伤值",
		BattlePropertyTypeDamageDefendPercent: "减伤万分比",
		BattlePropertyTypeForce:               "战力",
		BattlePropertyTypeFanTan:              "反弹",
		BattlePropertyTypeFanTanPercent:       "反弹万分比",
		BattlePropertyTypeCritRatePercent:     "最终暴击比例",
		BattlePropertyTypeCritHarmPercent:     "最终暴击伤害比例",
		BattlePropertyTypeHitRatePercent:      "最终命中比例",
		BattlePropertyTypeDodgeRatePercent:    "最终闪避比例",
		BattlePropertyTypeSpellCdPercent:      "技能cd时长",
		BattlePropertyTypeAddExp:              "增加经验",
		BattlePropertyTypeBlockRatePercent:    "强制格挡",
		BattlePropertyTypeZhuoShaoRes:         "灼烧抗性",
		BattlePropertyTypeJianSuRes:           "减速抗性",
		BattlePropertyTypeDingShenRes:         "定身抗性",
	}
)

func (pt BattlePropertyType) Penalize() bool {
	switch pt {
	case BattlePropertyTypeMaxHP,
		BattlePropertyTypeAttack,
		BattlePropertyTypeDefend,
		BattlePropertyTypeCrit,
		BattlePropertyTypeTough,
		BattlePropertyTypeBlock,
		BattlePropertyTypeAbnormality:
		return true
	}
	return false
}

func (pt BattlePropertyType) IsValid() bool {
	switch pt {
	case BattlePropertyTypeMaxHP,
		BattlePropertyTypeAttack,
		BattlePropertyTypeDefend,
		BattlePropertyTypeMaxTP,
		BattlePropertyTypeMoveSpeed,
		BattlePropertyTypeCrit,
		BattlePropertyTypeTough,
		BattlePropertyTypeBlock,
		BattlePropertyTypeAbnormality,
		BattlePropertyTypeDodge,
		BattlePropertyTypeHit,
		BattlePropertyTypeHuanYunAttack,
		BattlePropertyTypeHuanYunDef,
		BattlePropertyTypeBingDongRes,
		BattlePropertyTypePoJiaRes,
		BattlePropertyTypeKuiLeiRes,
		BattlePropertyTypeKuJieRes,
		BattlePropertyTypeShiMingRes,
		BattlePropertyTypeHunMiRes,
		BattlePropertyTypeXuRuoRes,
		BattlePropertyTypeJiaoXieRes,
		BattlePropertyTypeZhongDuRes,
		BattlePropertyTypeDamageAdd,
		BattlePropertyTypeDamageAddPercent,
		BattlePropertyTypeDamageDefend,
		BattlePropertyTypeDamageDefendPercent,
		BattlePropertyTypeForce,
		BattlePropertyTypeFanTan,
		BattlePropertyTypeFanTanPercent,
		BattlePropertyTypeCritRatePercent,
		BattlePropertyTypeCritHarmPercent,
		BattlePropertyTypeHitRatePercent,
		BattlePropertyTypeDodgeRatePercent,
		BattlePropertyTypeSpellCdPercent,
		BattlePropertyTypeAddExp,
		BattlePropertyTypeBlockRatePercent,
		BattlePropertyTypeZhuoShaoRes,
		BattlePropertyTypeJianSuRes,
		BattlePropertyTypeDingShenRes:
		return true
	}
	return false
}

var (
	propertyTypeMoreThanZeroMap = map[BattlePropertyType]bool{
		BattlePropertyTypeMaxHP:  true,
		BattlePropertyTypeAttack: true,
		BattlePropertyTypeDefend: true,
		BattlePropertyTypeMaxTP:  true,
	}
)

var (
	propertyTypeNoLessZeroMap = map[BattlePropertyType]bool{
		BattlePropertyTypeCrit:        true,
		BattlePropertyTypeTough:       true,
		BattlePropertyTypeAbnormality: true,
		BattlePropertyTypeBlock:       true,
	}
)

func (pt BattlePropertyType) MoreThanZero() bool {
	return propertyTypeMoreThanZeroMap[pt]
}

func (pt BattlePropertyType) NoLessZero() bool {
	return propertyTypeNoLessZeroMap[pt]
}

func (pt BattlePropertyType) String() string {
	return propertyTypeMap[pt]
}
