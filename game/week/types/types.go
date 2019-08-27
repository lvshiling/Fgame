package types

//周卡类型
type WeekType int32

const (
	WeekTypeSenior WeekType = iota // 高级周卡
	WeekTypeJunior                 // 普通周卡
)

func (t WeekType) Valid() bool {
	switch t {
	case
		WeekTypeSenior,
		WeekTypeJunior:
		return true
	default:
		return false
	}
}

const (
	MinType = WeekTypeSenior
	MaxType = WeekTypeJunior
)

var (
	weekStringMap = map[WeekType]string{
		WeekTypeJunior: "普通周卡",
		WeekTypeSenior: "超级周卡",
	}
)

func (t WeekType) String() string {
	return weekStringMap[t]
}
