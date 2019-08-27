package types

type BodyPositionType int32

const (
	//武器
	BodyPositionTypeWeapon BodyPositionType = iota
	//上衣
	BodyPositionTypeArmor
	//头盔
	BodyPositionTypeHelmet
	//战靴
	BodyPositionTypeShoe
	//腰带
	BodyPositionTypeBelt
	//护手
	BodyPositionTypeHandGuard
	//项链
	BodyPositionTypeNecklace
	//戒指
	BodyPositionTypeRing
)

const (
	minPosition = BodyPositionTypeWeapon
	maxPosition = BodyPositionTypeRing
	MinToyPos   = BodyPositionTypeWeapon
	MaxToyPos   = BodyPositionTypeHandGuard
)

func GetMinPosition() BodyPositionType {
	return minPosition
}

func GetMaxPosition() BodyPositionType {
	return maxPosition
}

var (
	bodyPositionTypeMap = map[BodyPositionType]string{
		BodyPositionTypeWeapon:    "武器",
		BodyPositionTypeArmor:     "上衣",
		BodyPositionTypeHelmet:    "头盔",
		BodyPositionTypeShoe:      "战靴",
		BodyPositionTypeBelt:      "腰带",
		BodyPositionTypeHandGuard: "护手",
		BodyPositionTypeNecklace:  "项链",
		BodyPositionTypeRing:      "戒指",
	}
)

func (bpt BodyPositionType) Valid() bool {
	switch bpt {
	case BodyPositionTypeWeapon,
		BodyPositionTypeArmor,
		BodyPositionTypeHelmet,
		BodyPositionTypeShoe,
		BodyPositionTypeBelt,
		BodyPositionTypeHandGuard,
		BodyPositionTypeNecklace,
		BodyPositionTypeRing:
		return true
	}
	return false
}

func (bpt BodyPositionType) String() string {
	return bodyPositionTypeMap[bpt]
}

//强化结果
type StrengthenResult struct {
	Pos    BodyPositionType
	Result EquipmentStrengthenResultType
}
