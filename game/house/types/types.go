package types

import (
	"fgame/fgame/pkg/mathutils"
)

type HouseType int32

const (
	HouseTypeSilver   HouseType = 1 + iota //银两
	HouseTypeBindGold                      //绑元
)

// 校验枚举类型
func (t HouseType) Valid() bool {
	switch t {
	case HouseTypeSilver,
		HouseTypeBindGold:
		return true
	default:
		return false
	}
}

//房子操作类型
type HouseOperateType int32

const (
	HouseOperateTypeActivate HouseOperateType = iota //激活
	HouseOperateTypeUplevel                          //升级
	HouseOperateTypeSell                             //出售
)

// 随机操作类型
const (
	MinOperateType = HouseOperateTypeUplevel
	MaxOperateType = HouseOperateTypeSell
)

func RandomOperateTypeExcludeActivate() HouseOperateType {
	min := int(MinOperateType)
	max := int(MaxOperateType)
	randomTypeInt := mathutils.RandomRange(min, max+1)
	return HouseOperateType(randomTypeInt)
}

//随机操作类型
var (
	operateTypeList = []HouseOperateType{
		HouseOperateTypeActivate,
		HouseOperateTypeSell,
	}
)

func RandomOperateTypeExcludeUplevel() HouseOperateType {
	var weights []int64
	for i := 1; i <= len(operateTypeList); i++ {
		weights = append(weights, int64(1))
	}

	randomIndex := mathutils.RandomWeights(weights)
	return operateTypeList[randomIndex]
}

//初始房子数据
const (
	InitHouseIndex = 0
	InitHouseLevel = 1
)
