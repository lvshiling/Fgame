package lingTong

//属性作用器效果
type PropertyEffectorType uint32

const (
	//初始化
	LingTongPropertyEffectorTypeInit PropertyEffectorType = iota
	LingTongPropertyEffectorTypeFashion
	LingTongPropertyEffectorTypeWeapon
	LingTongPropertyEffectorTypeMount
	LingTongPropertyEffectorTypeWing
	LingTongPropertyEffectorTypeShenFa
	LingTongPropertyEffectorTypeLingYu
	LingTongPropertyEffectorTypeXianTi
	LingTongPropertyEffectorTypeFaBao
)

func (pet PropertyEffectorType) EffectorType() uint32 {
	return uint32(pet)
}

var (
	propertyEffectoryTypeMap = map[PropertyEffectorType]string{
		LingTongPropertyEffectorTypeInit:    "初始化",
		LingTongPropertyEffectorTypeFashion: "时装",
		LingTongPropertyEffectorTypeWeapon:  "冰魂",
		LingTongPropertyEffectorTypeMount:   "坐骑",
		LingTongPropertyEffectorTypeWing:    "战翼",
		LingTongPropertyEffectorTypeShenFa:  "身法",
		LingTongPropertyEffectorTypeLingYu:  "领域",
		LingTongPropertyEffectorTypeXianTi:  "仙体",
		LingTongPropertyEffectorTypeFaBao:   "法宝",
	}
)

func (pet PropertyEffectorType) String() string {
	return propertyEffectoryTypeMap[pet]
}

func (pet PropertyEffectorType) Mask() uint64 {
	return 1 << uint(pet)
}

//所有作用器
const (
	PropertyEffectorTypeMaskAll = 1<<63 - 1
)
