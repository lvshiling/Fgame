package types

type HuntType int32

const (
	HuntTypeMingGe               HuntType = iota //命格
	HuntTypeShenQi                               //神器
	HuntTypeShengHen                             //圣痕
	HuntTypeTuLongEquip                          //屠龙装
	HuntTypeShangGuZhiLingMiling                 //上古之灵觅灵
)

func (t HuntType) Valid() bool {
	switch t {
	case HuntTypeMingGe,
		HuntTypeShenQi,
		HuntTypeShengHen,
		HuntTypeTuLongEquip,
		HuntTypeShangGuZhiLingMiling:
		break
	default:
		return false
	}
	return true
}

const (
	MinHuntType = HuntTypeMingGe
	MaxHuntType = HuntTypeShangGuZhiLingMiling
)
