package types

//投资计划类型
type InvestLevelType int32

const (
	InvesetLevelTypeJunior InvestLevelType = iota //初级投资计划
	InvesetLevelTypeSenior                        //高级投资计划
)

func (t InvestLevelType) Valid() bool {
	switch t {
	case InvesetLevelTypeJunior,
		InvesetLevelTypeSenior:
		return true
	default:
		return false
	}
}
