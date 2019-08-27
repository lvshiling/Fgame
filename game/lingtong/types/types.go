package types

//灵童激活类型
type LingTongActivateType int32

const (
	LingTongActivateTypeItem LingTongActivateType = iota
)

func (t LingTongActivateType) Valid() bool {
	switch t {
	case LingTongActivateTypeItem:
		return true
	}
	return false
}
