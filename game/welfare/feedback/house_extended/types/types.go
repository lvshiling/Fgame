package types

type HouseRewType int32

const (
	HouseRewTypeActivate HouseRewType = iota //激活礼包
	HouseRewTypeUplevel                      //升级礼包
)

func (t HouseRewType) Valid() bool {
	switch t {
	case HouseRewTypeActivate,
		HouseRewTypeUplevel:
		return true
	default:
		return false
	}
}
