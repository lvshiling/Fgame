package types

//在线抽奖类型
type OnlineDrewType int32

const (
	OnlineDrewTypeSilver OnlineDrewType = iota
	OnlineDrewTypeGold
)

func (t OnlineDrewType) Valid() bool {
	switch t {
	case OnlineDrewTypeGold,
		OnlineDrewTypeSilver:
		return true
	default:
		return false
	}
}
