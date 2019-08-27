package types

const (
	//祝福值有效时间
	EFFCTIVE_TIME = 24 * 60 * 60 * 1000
)

type LingyuType int32

const (
	//进阶领域
	LingyuTypeAdvanced LingyuType = 1 + iota
	//领域皮肤
	LingyuTypeSkin
)

func (t LingyuType) Valid() bool {
	switch t {
	case LingyuTypeAdvanced,
		LingyuTypeSkin:
		return true
	}
	return false
}

//幻化条件
type LingyuUCondType int32

const (
	//没有限制
	LingyuUCondTypeZ LingyuUCondType = iota
	//领域阶别
	LingyuUCondTypeX
	//食用幻化丹数量
	LingyuUCondTypeU
	//消耗物品
	LingyuUCondTypeI
)

func (t LingyuUCondType) Valid() bool {
	switch t {
	case LingyuUCondTypeZ,
		LingyuUCondTypeX,
		LingyuUCondTypeU,
		LingyuUCondTypeI:
		return true
	}
	return false
}
