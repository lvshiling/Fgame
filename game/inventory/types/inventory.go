package types

const MAX_SLOT_NUM = 9999

type InventoryTag int32

const (
	//其它
	InventoryTagOther InventoryTag = iota
	//装备
	InventoryTagEquipment
	//消耗
	InventoryTagConsume
	//宝石
	InventoryTagGem
	//升阶系统装备
	InventoryTagSystemEquipment
)

func (it InventoryTag) Valid() bool {
	switch it {
	case InventoryTagOther,
		InventoryTagEquipment,
		InventoryTagConsume,
		InventoryTagGem,
		InventoryTagSystemEquipment:
		return true
	}
	return false
}

type BagType int32

const (
	//主背包
	BagTypePrim BagType = iota
	//宝石
	BagTypeGem
	//鲲背包
	BagTypeKun
	//命格
	BagTypeMingGe
	//神器
	BagTypeShenQi
	//屠龙装备---------5
	BagTypeTuLongEquip
	//器灵
	BagTypeQiLing
	//英灵谱
	BagTypeYingLingPu
)

func (it BagType) Valid() bool {
	switch it {
	case BagTypePrim,
		BagTypeGem,
		BagTypeKun,
		BagTypeMingGe,
		BagTypeShenQi,
		BagTypeTuLongEquip,
		BagTypeQiLing,
		BagTypeYingLingPu:
		return true
	}
	return false
}

var bagTypeMap = map[BagType]string{
	BagTypePrim:        "主背包",
	BagTypeGem:         "宝石背包",
	BagTypeKun:         "鲲背包",
	BagTypeMingGe:      "命格",
	BagTypeShenQi:      "神器",
	BagTypeTuLongEquip: "屠龙装备",
	BagTypeQiLing:      "器灵",
	BagTypeYingLingPu:  "英灵谱",
}

func (it BagType) String() string {
	return bagTypeMap[it]
}

//过期时间类型(运营活动)
type NewItemLimitTimeType int32

const (
	//无过期时间
	NewItemLimitTimeTypeNone NewItemLimitTimeType = iota
	//当日指定时间过期
	NewItemLimitTimeTypeExpiredToday
	//获得物品多少秒后过期
	NewItemLimitTimeTypeExpiredAfterTime
	//指定时间日期过期
	NewItemLimitTimeTypeExpiredDate
)

func (t NewItemLimitTimeType) Valid() bool {
	switch t {
	case NewItemLimitTimeTypeNone,
		NewItemLimitTimeTypeExpiredToday,
		NewItemLimitTimeTypeExpiredAfterTime,
		NewItemLimitTimeTypeExpiredDate:
		return true
	}
	return false
}

type IsDepotType int32

const (
	//主背包
	IsDepotTypePrim IsDepotType = iota
	//仓库
	IsDepotTypeDepot
	//秘宝仓库
	IsDepotTypeMiBao
	//材料仓库
	IsDepotTypeMaterial
)

func (it IsDepotType) Valid() bool {
	switch it {
	case IsDepotTypePrim,
		IsDepotTypeDepot,
		IsDepotTypeMiBao,
		IsDepotTypeMaterial:
		return true
	}
	return false
}

var isDepotTypeMap = map[IsDepotType]string{
	IsDepotTypePrim:     "主背包",
	IsDepotTypeDepot:    "仓库",
	IsDepotTypeMiBao:    "秘宝仓库",
	IsDepotTypeMaterial: "材料仓库",
}

func (it IsDepotType) String() string {
	return isDepotTypeMap[it]
}
