package types

const (
	//祝福值有效时间
	EFFCTIVE_TIME = 24 * 60 * 60 * 1000
)

type WingType int32

const (
	//进阶战翼
	WingTypeAdvanced WingType = 1 + iota
	//战翼皮肤
	WingTypeSkin
)

func (wt WingType) Valid() bool {
	switch wt {
	case WingTypeAdvanced,
		WingTypeSkin:
		return true
	}
	return false
}

//幻化条件
type WingUCondType int32

const (
	//没有限制
	WingUCondTypeZ WingUCondType = iota
	//战翼阶别
	WingUCondTypeX
	//食用幻化丹数量
	WingUCondTypeU
	//消耗物品
	WingUCondTypeI
)

func (wuct WingUCondType) Valid() bool {
	switch wuct {
	case WingUCondTypeZ,
		WingUCondTypeX,
		WingUCondTypeU,
		WingUCondTypeI:
		return true
	}
	return false
}
