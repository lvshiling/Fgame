package types

type GuaJiAdvanceType int32

const (
	GuaJiAdvanceTypeMount GuaJiAdvanceType = iota
	GuaJiAdvanceTypeWing
	GuaJiAdvanceTypeBodyshield
	GuaJiAdvanceTypeAnqi
	GuaJiAdvanceTypeFabao
	GuaJiAdvanceTypeShenfa
	GuaJiAdvanceTypeXianti
	GuaJiAdvanceTypeLingyu
	GuaJiAdvanceTypeShihunfan
	GuaJiAdvanceTypeTianmoti
	GuaJiAdvanceTypeFeather
	GuaJiAdvanceTypeShield
	GuaJiAdvanceTypeMassacre
)

const (
	GuaJiAdvanceTypeLingTongWeapon GuaJiAdvanceType = iota + 100
	GuaJiAdvanceTypeLingTongMount
	GuaJiAdvanceTypeLingTongWing
	GuaJiAdvanceTypeLingTongShenFa
	GuaJiAdvanceTypeLingTongLingYu
	GuaJiAdvanceTypeLingTongFaBao
	GuaJiAdvanceTypeLingTongXianTi
)

var (
	guaJiAdvanceMap = map[GuaJiAdvanceType]string{
		GuaJiAdvanceTypeMount:          "坐骑",
		GuaJiAdvanceTypeWing:           "战翼",
		GuaJiAdvanceTypeBodyshield:     "护体盾",
		GuaJiAdvanceTypeAnqi:           "暗器",
		GuaJiAdvanceTypeFabao:          "法宝",
		GuaJiAdvanceTypeShenfa:         "身法",
		GuaJiAdvanceTypeXianti:         "仙体",
		GuaJiAdvanceTypeLingyu:         "领域",
		GuaJiAdvanceTypeShihunfan:      "噬魂番",
		GuaJiAdvanceTypeTianmoti:       "天魔体",
		GuaJiAdvanceTypeFeather:        "仙羽",
		GuaJiAdvanceTypeShield:         "盾刺",
		GuaJiAdvanceTypeMassacre:       "戮仙刃",
		GuaJiAdvanceTypeLingTongWeapon: "灵兵",
		GuaJiAdvanceTypeLingTongMount:  "灵骑",
		GuaJiAdvanceTypeLingTongWing:   "灵翼",
		GuaJiAdvanceTypeLingTongShenFa: "灵身",
		GuaJiAdvanceTypeLingTongLingYu: "灵域",
		GuaJiAdvanceTypeLingTongFaBao:  "灵宝",
		GuaJiAdvanceTypeLingTongXianTi: "灵体",
	}
)

func (t GuaJiAdvanceType) String() string {
	return guaJiAdvanceMap[t]
}

func (t GuaJiAdvanceType) Valid() bool {
	switch t {
	case GuaJiAdvanceTypeMount,
		GuaJiAdvanceTypeWing,
		GuaJiAdvanceTypeBodyshield,
		GuaJiAdvanceTypeAnqi,
		GuaJiAdvanceTypeFabao,
		GuaJiAdvanceTypeShenfa,
		GuaJiAdvanceTypeXianti,
		GuaJiAdvanceTypeLingyu,
		GuaJiAdvanceTypeShihunfan,
		GuaJiAdvanceTypeTianmoti,
		GuaJiAdvanceTypeFeather,
		GuaJiAdvanceTypeShield,
		GuaJiAdvanceTypeMassacre,
		GuaJiAdvanceTypeLingTongWeapon,
		GuaJiAdvanceTypeLingTongMount,
		GuaJiAdvanceTypeLingTongWing,
		GuaJiAdvanceTypeLingTongShenFa,
		GuaJiAdvanceTypeLingTongLingYu,
		GuaJiAdvanceTypeLingTongFaBao,
		GuaJiAdvanceTypeLingTongXianTi:
		return true
	}
	return false
}

func GetGuaJiAdvanceMap() map[GuaJiAdvanceType]string {
	return guaJiAdvanceMap
}
