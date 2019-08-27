package types

//进阶方式
type AdvancedType int32

const (
	AdvancedTypeJinJieDan        AdvancedType = iota + 1 //进阶丹
	AdvancedTypeBlessDan                                 //祝福丹
	AdvancedTypeTicket                                   //直升券
	AdvancedTypeActivateCard                             //激活卡
	AdvancedTypeShaQi                                    //杀气
	AdvancedTypeSilver                                   //银两
	AdvancedTypeGold                                     //元宝
	AdvancedTypeBindGold                                 //绑元
	AdvancedTypeXingChen                                 //星尘
	AdvancedTypeOpenActivity                             //运营活动激活
	AdvancedTypeDebris                                   //碎片
	AdvancedTypeElement                                  //元素之魂
	AdvancedTypeLingQi                                   //灵气
	AdvancedTypeAdditionsysEquip                         //附加系统装备
)

var (
	typeMap = map[AdvancedType]string{
		AdvancedTypeJinJieDan:        "进阶丹",
		AdvancedTypeBlessDan:         "祝福丹",
		AdvancedTypeTicket:           "直升券",
		AdvancedTypeActivateCard:     "激活卡",
		AdvancedTypeShaQi:            "杀气",
		AdvancedTypeSilver:           "银两",
		AdvancedTypeGold:             "元宝",
		AdvancedTypeBindGold:         "绑元",
		AdvancedTypeXingChen:         "星尘",
		AdvancedTypeOpenActivity:     "运营活动",
		AdvancedTypeDebris:           "碎片",
		AdvancedTypeElement:          "元素之魂",
		AdvancedTypeLingQi:           "灵气",
		AdvancedTypeAdditionsysEquip: "附加系统装备",
	}
)

func (t AdvancedType) String() string {
	return typeMap[t]
}

type SpecialAdvancedType int32

const (
	//普通升阶
	SpecialAdvancedTypeDefault SpecialAdvancedType = iota
	//充值升阶
	SpecialAdvancedTypeCharge
	//花费升阶
	SpecialAdvancedTypeCost
)

var (
	specialAdvancedTypeMap = map[SpecialAdvancedType]string{
		SpecialAdvancedTypeDefault: "普通升阶",
		SpecialAdvancedTypeCharge:  "充值升阶",
		SpecialAdvancedTypeCost:    "花费升阶",
	}
)

func (spt SpecialAdvancedType) Valid() bool {
	switch spt {
	case SpecialAdvancedTypeDefault,
		SpecialAdvancedTypeCharge,
		SpecialAdvancedTypeCost:
		return true
	}
	return false
}

func (spt SpecialAdvancedType) String() string {
	return specialAdvancedTypeMap[spt]
}

//升阶关联模块激活的皮肤类型
type AdvancedUnitePiFuType int32

const (
	//默认0为没有
	AdvancedUnitePiFuTypeDefault AdvancedUnitePiFuType = iota
	//兵魂
	AdvancedUnitePiFuTypeWeapon
	//时装
	AdvancedUnitePiFuTypeFashion
	//坐骑
	AdvancedUnitePiFuTypeMount
	//战翼
	AdvancedUnitePiFuTypeWing
	//暗器
	AdvancedUnitePiFuTypeAnQi
	//法宝
	AdvancedUnitePiFuTypeFaBao
	//仙体
	AdvancedUnitePiFuTypeXianTi
	//领域
	AdvancedUnitePiFuTypeLingYu
	//身法
	AdvancedUnitePiFuTypeShenFa
	//称号
	AdvancedUnitePiFuTypeTitle
)

var (
	advancedUnitePiFuTypeMap = map[AdvancedUnitePiFuType]string{
		AdvancedUnitePiFuTypeDefault: "默认0为没有",
		AdvancedUnitePiFuTypeWeapon:  "兵魂",
		AdvancedUnitePiFuTypeFashion: "时装",
		AdvancedUnitePiFuTypeMount:   "坐骑",
		AdvancedUnitePiFuTypeWing:    "战翼",
		AdvancedUnitePiFuTypeAnQi:    "暗器",
		AdvancedUnitePiFuTypeFaBao:   "法宝",
		AdvancedUnitePiFuTypeXianTi:  "仙体",
		AdvancedUnitePiFuTypeLingYu:  "领域",
		AdvancedUnitePiFuTypeShenFa:  "身法",
		AdvancedUnitePiFuTypeTitle:   "称号",
	}
)

func (spt AdvancedUnitePiFuType) Valid() bool {
	switch spt {
	case AdvancedUnitePiFuTypeDefault,
		AdvancedUnitePiFuTypeWeapon,
		AdvancedUnitePiFuTypeFashion,
		AdvancedUnitePiFuTypeMount,
		AdvancedUnitePiFuTypeWing,
		AdvancedUnitePiFuTypeAnQi,
		AdvancedUnitePiFuTypeFaBao,
		AdvancedUnitePiFuTypeXianTi,
		AdvancedUnitePiFuTypeLingYu,
		AdvancedUnitePiFuTypeShenFa,
		AdvancedUnitePiFuTypeTitle:
		return true
	}
	return false
}

func (spt AdvancedUnitePiFuType) String() string {
	return advancedUnitePiFuTypeMap[spt]
}

//变化方式
type ChangeType int32

const (
	ChangeTypeAttendGet ChangeType = iota //抽奖获得
	ChangeTypeItemGet                     //物品获得
	ChangeTypeRefresh                     //刷新清空
	ChangeTypeUse                         //使用
	ChangeTypeExchange                    //兑换
)

var (
	changeTypeMap = map[ChangeType]string{
		ChangeTypeAttendGet: "抽奖获得",
		ChangeTypeItemGet:   "物品获得",
		ChangeTypeRefresh:   "刷新清空",
		ChangeTypeUse:       "使用",
		ChangeTypeExchange:  "兑换",
	}
)

func (t ChangeType) String() string {
	return changeTypeMap[t]
}

// 结果
type ResultType int32

const (
	ResultTypeSuccess ResultType = iota //强化成功
	ResultTypeFailed                    //强化失败
	ResultTypeBack                      //强化回退
)
