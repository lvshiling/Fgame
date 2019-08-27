package types

type SmeltType int32

const (
	//光
	SmeltTypeLight SmeltType = iota
	//暗
	SmeltTypeDark
	//风
	SmeltTypeWind
	//雷
	SmeltTypeThunder
)

const (
	MinSmeltType = SmeltTypeLight
	MaxSmeltType = SmeltTypeThunder
)

var (
	smeltTypeMap = map[SmeltType]string{
		SmeltTypeLight:   "光",
		SmeltTypeDark:    "暗",
		SmeltTypeWind:    "风",
		SmeltTypeThunder: "雷",
	}
)

func (spt SmeltType) Valid() bool {
	switch spt {
	case SmeltTypeLight,
		SmeltTypeDark,
		SmeltTypeWind,
		SmeltTypeThunder:
		return true
	}
	return false
}

func (spt SmeltType) String() string {
	return smeltTypeMap[spt]
}
