package types

const (
	//祝福值有效时间
	EFFCTIVE_TIME = 24 * 60 * 60 * 1000
)

//仙体类型
type XianTiType int32

const (
	//进阶仙体
	XianTiTypeAdvanced XianTiType = 1 + iota
	//仙体皮肤
	XianTiTypeSkin
)

func (mt XianTiType) Valid() bool {
	switch mt {
	case XianTiTypeAdvanced,
		XianTiTypeSkin:
		return true
	}
	return false
}

//幻化条件
type XianTiUCondType int32

const (
	//没有限制
	XianTiUcondTypeZ XianTiUCondType = iota
	//仙体阶别
	XianTiUCondTypeX
	//食用幻化丹数量
	XianTiUCondTypeU
	//消耗物品
	XianTiUCondTypeI
	//关联系统阶数
	XianTiUCondTypeW
)

func (muct XianTiUCondType) Valid() bool {
	switch muct {
	case XianTiUcondTypeZ,
		XianTiUCondTypeX,
		XianTiUCondTypeU,
		XianTiUCondTypeI,
		XianTiUCondTypeW:
		return true
	}
	return false
}
