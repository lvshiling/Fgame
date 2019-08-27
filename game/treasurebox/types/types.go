package types

type BoxType int32

//宝箱类型
const (
	ObtainRes BoxType = iota //开启后直接获得奖励
	ChooseRes                //开启后从多个道具中自由选择奖励
)

func (t BoxType) Valid() bool {
	switch t {
	case ObtainRes, ChooseRes:
		return true
	default:
		return false
	}
}

//可选消耗类型
type BoxCostType int32

const (
	BoxCostTypeFree BoxCostType = iota //免费开启
	BoxCostTypeGold                    //元宝开启
)

func (t BoxCostType) Valid() bool {
	switch t {
	case BoxCostTypeFree,
		BoxCostTypeGold:
		return true
	default:
		return false
	}
}
