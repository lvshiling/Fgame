package types

type FaBaoType int32

const (
	//进阶法宝
	FaBaoTypeAdvanced FaBaoType = 1 + iota
	//法宝皮肤
	FaBaoTypeSkin
)

func (t FaBaoType) Valid() bool {
	switch t {
	case FaBaoTypeAdvanced,
		FaBaoTypeSkin:
		return true
	}
	return false
}

//幻化条件
type FaBaoUCondType int32

const (
	//没有限制
	FaBaoUCondTypeZ FaBaoUCondType = iota
	//法宝阶别
	FaBaoUCondTypeX
	//食用幻化丹数量
	FaBaoUCondTypeU
	//消耗物品
	FaBaoUCondTypeI
	//关联系统阶数
	FaBaoUCondTypeW
)

func (t FaBaoUCondType) Valid() bool {
	switch t {
	case FaBaoUCondTypeZ,
		FaBaoUCondTypeX,
		FaBaoUCondTypeU,
		FaBaoUCondTypeI,
		FaBaoUCondTypeW:
		return true
	}
	return false
}
