package types

// 次数限制类型
type TimesLimitType int32

const (
	TimesLimitTypePersonal TimesLimitType = iota //个人次数
	TimesLimitTypeGlobal                         //全服次数
)

func (t TimesLimitType) Valid() bool {
	switch t {
	case TimesLimitTypePersonal,
		TimesLimitTypeGlobal:
		return true
	default:
		return false
	}
}
