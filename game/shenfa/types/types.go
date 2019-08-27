package types

const (
	//祝福值有效时间
	EFFCTIVE_TIME = 24 * 60 * 60 * 1000
)

type ShenfaType int32

const (
	//进阶身法
	ShenfaTypeAdvanced ShenfaType = 1 + iota
	//身法皮肤
	ShenfaTypeSkin
)

func (t ShenfaType) Valid() bool {
	switch t {
	case ShenfaTypeAdvanced,
		ShenfaTypeSkin:
		return true
	}
	return false
}

//幻化条件
type ShenfaUCondType int32

const (
	//没有限制
	ShenfaUCondTypeZ ShenfaUCondType = iota
	//身法阶别
	ShenfaUCondTypeX
	//食用幻化丹数量
	ShenfaUCondTypeU
	//消耗物品
	ShenfaUCondTypeI
)

func (t ShenfaUCondType) Valid() bool {
	switch t {
	case ShenfaUCondTypeZ,
		ShenfaUCondTypeX,
		ShenfaUCondTypeU,
		ShenfaUCondTypeI:
		return true
	}
	return false
}
