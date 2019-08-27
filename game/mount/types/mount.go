package types

const (
	//祝福值有效时间
	EFFCTIVE_TIME = 24 * 60 * 60 * 1000
)

//坐骑类型
type MountType int32

const (
	//进阶坐骑
	MountTypeAdvanced MountType = 1 + iota
	//神龙现世
	MountTypeShenLong
	//特殊处理
	MountTypeLingTongMount
)

func (mt MountType) Valid() bool {
	switch mt {
	case MountTypeAdvanced,
		MountTypeShenLong,
		MountTypeLingTongMount:
		return true
	}
	return false
}

//幻化条件
type MountUCondType int32

const (
	//没有限制
	MountUcondTypeZ MountUCondType = iota
	//坐骑阶别
	MountUCondTypeX
	//食用幻化丹数量
	MountUCondTypeU
	//消耗物品
	MountUCondTypeI
)

func (muct MountUCondType) Valid() bool {
	switch muct {
	case MountUcondTypeZ,
		MountUCondTypeX,
		MountUCondTypeU,
		MountUCondTypeI:
		return true
	}
	return false
}
