package types

import (
	inventorytypes "fgame/fgame/game/inventory/types"
)

type ItemType int32

const (
	//默认
	ItemTypeDefault ItemType = iota
	//血量
	ItemTypeHp = 1
	//经验
	ItemTypeExp = 2
	//银两
	ItemTypeSilver = 3
	//元宝
	ItemTypeGold = 4
	//技能
	ItemTypeSkill = 5
	//buff
	ItemTypeBuff = 6
	//装备
	ItemTypeEquipment = 7
	//时装
	ItemTypeFashion = 8
	//兵魂
	ItemTypeBingHun = 9
	//坐骑装备
	ItemTypePetEquipment = 10
	//食丹物品
	ItemTypeDan = 11
	//藏宝
	ItemTypeTreasure = 12
	//装备宝石
	ItemTypeEquipmentGem = 13
	//坐骑
	ItemTypePet = 14
	//生命之源
	ItemTypeLifeOrigin = 15
	//技能要诀
	ItemTypeSkillPoint = 16
	//宝石
	ItemTypeGem = 17
	//装备升阶
	ItemTypeEquipmentUpgrade = 18
	//战翼
	ItemTypeWing = 19
	//冰魂
	ItemTypeSoul = 20
	//帝魂
	ItemTypeKingSoul = 21
	//护体盾
	ItemTypeDun = 22
	//称号
	ItemTypeTitle = 23
	//自动使用资源类型
	ItemTypeAutoUseRes = 24
	//鲜花
	ItemTypeXianHua = 25
	//屠魔令
	ItemTypeTuMoLing = 26
	//绝学
	ItemTypeJueXue = 27
	//心法
	ItemTypeXinFa = 28
	//副本门票
	ItemTypeEctypal = 29
	//仙盟
	ItemTypeAlliance = 30
	//礼包
	ItemTypeGiftBag = 31
	//神龙现世
	ItemTypeShenLong = 32
	//喜糖
	ItemTypeCandy = 33
	//婚戒
	ItemTypeWedRing = 34
	//领域
	ItemTypeLingyu = 35
	//身法
	ItemTypeShenFa = 36
	//婚戒升阶物品
	ItemTypeRingAdvaced = 37
	//元神金装
	ItemTypeGoldEquip = 38
	//爱情树升阶物品
	ItemTypeTreeAdvanced = 39
	//坤
	ItemTypeKun = 40
	//pk技能升级道具
	ItemTypePKSkillUpgrade = 41
	//求婚成功的戒指
	ItemTypeWedSuccRing = 42
	//挖矿
	ItemTypeWaKuang = 43
	//棋局破解卷
	ItemTypeQiJu = 44
	//银两副本升级
	ItemTypeSilverFuBenUpgrade = 45
	//暗器相关道具
	ItemTypeAnQi = 46
	//祝福丹
	ItemTypeBlessDan = 47
	//等级直升丹
	ItemTypeUpLevelDan = 48
	//直升券
	ItemTypeAdvancedTicket = 49
	//Boss召唤券
	ItemTypeBossCallTicket = 50
	//红名值道具
	ItemTypePkValue = 51
	//战翼激活卡
	ItemTypeWingActivateCard = 52
	//金装相关
	ItemTypeGoldEquipStrengthen = 53
	//宝石防爆符
	ItemTypeGemAvoidBomb = 54
	//幸运符
	ItemTypeLuckyRate = 55
	//性别卡
	ItemTypeSexCard = 56
	//改名卡
	ItemTypeRenameCard = 57
	//收益卡
	ItemTypeResouceCard = 58
	//打宝相关
	ItemTypeTower = 59
	//个人BOSS
	ItemTypeMyBoss = 60
	//坐骑装备
	ItemTypeMountEquip = 61
	//战翼符石
	ItemTypeWingStone = 62
	//暗器机关
	ItemTypeAnqiJiguan = 63
	//技能书
	ItemTypeJiNengShu = 64
	//强化石
	ItemTypeQiangHuaShi = 65
	//追踪符
	ItemTypeZhuiZongFu = 66
	//天书相关
	ItemTypeTianShu = 67
	//碎片相关
	ItemTypeDebris = 68
	//法宝配饰
	ItemTypeFaBaoSuit = 69
	//法宝相关
	ItemTypeFaBao = 70
	//仙体相关
	ItemTypeXianTi = 71
	//仙体灵玉
	ItemTypeXianTiLingYu = 72
	//血盾系统
	ItemTypeXueDun = 73
	//扶持
	ItemTypeFuChi = 74
	//点星系统相关
	ItemTypeDianXing = 75
	//天魔
	ItemTypeTianMo = 76
	//血魔
	ItemTypeXueMo = 77
	//衣橱套装
	ItemTypeWardrobe = 78
	//领域装备
	ItemTypeLingyuEquip = 79
	//身法装备
	ItemTypeShenfaEquip = 80
	//噬魂幡装备
	ItemTypeShiHunFanEquip = 81
	//天魔体装备
	ItemTypeTianMoTiEquip = 82
	//灵童改名卡
	ItemTypeLingTongReNameCard = 83
	//灵童道具
	ItemTypeLingTong = 84
	//灵童时装道具
	ItemTypeLingTongFashion = 85
	//灵童兵魂道具
	ItemTypeLingTongWeapon = 86
	//灵童坐骑道具
	ItemTypeLingTongMount = 87
	//灵童战翼道具
	ItemTypeLingTongWing = 88
	//灵童身法道具
	ItemTypeLingTongShenFa = 89
	//灵童领域道具
	ItemTypeLingTongLingYu = 90
	//灵童法宝道具
	ItemTypeLingTongFaBao = 91
	//灵童仙体道具
	ItemTypeLingTongXianTi = 92
	//灵童兵器装备
	ItemTypeLingTongWeaponEquip = 93
	//灵童坐骑装备
	ItemTypeLingTongMountEquip = 94
	//灵童战翼装备
	ItemTypeLingTongWingEquip = 95
	//灵童身法装备
	ItemTypeLingTongShenFaEquip = 96
	//灵童领域装备
	ItemTypeLingTongLingYuEquip = 97
	//灵童法宝装备
	ItemTypeLingTongFaBaoEquip = 98
	//灵童仙体装备
	ItemTypeLingTongXianTiEquip = 99
	//灵童技能书
	ItemTypeLingTongJiNengShu = 100
	//飞升相关
	ItemTypeFeiSheng = 101
	//红包
	ItemTypeHongBao = 102
	//化灵丹
	ItemTypeHuaLingDan = 103
	//天赋升级丹
	ItemTypeTianFuDan = 104
	//装备宝库抽奖券
	ItemTypeEquipBaoKuTicket = 105
	//圣痕青龙装备
	ItemTypeShengHenEquipQingLong = 106
	//圣痕白虎装备
	ItemTypeShengHenEquipBaiHu = 107
	//圣痕朱雀装备
	ItemTypeShengHenEquipZhuQue = 108
	//圣痕玄武装备
	ItemTypeShengHenEquipXuanWu = 109
	//圣痕相关道具
	ItemTypeSheng = 110
	//命格
	ItemTypeMingGe = 111
	//神器相关道具
	ItemTypeShenQi = 112
	//屠龙装备
	ItemTypeTuLongEquip = 113
	// 寻宝道具
	ItemTypeHunt = 114
	// 屠龙装备相关
	ItemTypeTuLongEquipItem = 118
	// 英灵谱相关
	ItemTypeYingLingPu = 119
	//阵法相关
	ItemTypeZhenFa = 120
	//表白
	ItemTypeBiaoBai = 121
	//宝宝相关
	ItemTypeBaoBao = 122
	//宝宝玩具
	ItemTypeBabyToy = 123
	//定情信物
	ItemTypeDingQing = 124
	//宝宝卡
	ItemTypeBaoBaoCard = 125
	//附加系统通灵丹
	ItemTypeAdditionsysTongLingDan = 126
	//附加系统觉醒丹
	ItemTypeAdditionsysJueXingDan = 127
	//元宝卡
	ItemTypeYuanBaoKa = 128
	//背包扩充符
	ItemTypeExpendBagSlotCard = 129
	//八卦符石
	ItemTypeFuShi = 130
	//物品技能
	ItemTypeItemSkill = 131
	//Boss道具
	ItemTypeBossItem = 132
	//结义道具
	ItemTypeJieYiItem = 133
	//现金元宝卡
	ItemTypeXianJinYuanBaoKa = 134
	//创世之战道具
	ItemTypeChuangShiZhiZhan = 135
	//BOSS密藏道具
	ItemTypeBOSSMiZang = 136
	//大力丸
	ItemTypeVigorousPill = 137
	//无双神器
	ItemTypeWushuangWeapon = 138
	//无双神器精华
	ItemTypeWushuangWeaponEssence = 139
	//神铸装备
	ItemTypeGodCastingEquip = 140
	//满汉全席合成物品
	ItemTypeManHanQuanXi = 141
	//灵童装备
	ItemTypeLingTongEquip = 142
	// 特戒相关
	ItemTypeTeRing = 143
	//上古之灵相关
	ItemTypeShangGuZhiLing = 144
)

var itemTypeMap = map[ItemType]string{
	ItemTypeDefault:                "默认",
	ItemTypeHp:                     "血量物品",
	ItemTypeExp:                    "经验物品",
	ItemTypeSilver:                 "银两物品",
	ItemTypeGold:                   "元宝物品",
	ItemTypeSkill:                  "技能物品",
	ItemTypeBuff:                   "buff物品",
	ItemTypeEquipment:              "装备物品",
	ItemTypeFashion:                "时装物品",
	ItemTypeBingHun:                "冰魂物品",
	ItemTypePetEquipment:           "坐骑装备",
	ItemTypeDan:                    "食丹物品",
	ItemTypeTreasure:               "藏宝物品",
	ItemTypeEquipmentGem:           "装备宝石",
	ItemTypePet:                    "坐骑",
	ItemTypeLifeOrigin:             "生命之源",
	ItemTypeSkillPoint:             "技能要诀",
	ItemTypeGem:                    "宝石",
	ItemTypeEquipmentUpgrade:       "装备升阶",
	ItemTypeWing:                   "战翼",
	ItemTypeSoul:                   "冰魂",
	ItemTypeKingSoul:               "帝魂",
	ItemTypeDun:                    "护体盾",
	ItemTypeTitle:                  "称号",
	ItemTypeAutoUseRes:             "自动使用资源类型",
	ItemTypeXianHua:                "鲜花",
	ItemTypeTuMoLing:               "屠魔令",
	ItemTypeJueXue:                 "绝学",
	ItemTypeXinFa:                  "心法",
	ItemTypeEctypal:                "副本门票",
	ItemTypeAlliance:               "仙盟",
	ItemTypeGiftBag:                "礼包",
	ItemTypeShenLong:               "神龙现世",
	ItemTypeCandy:                  "喜糖",
	ItemTypeWedRing:                "婚戒",
	ItemTypeLingyu:                 "领域",
	ItemTypeShenFa:                 "身法",
	ItemTypeRingAdvaced:            "婚戒升阶物品",
	ItemTypeGoldEquip:              "元神金装",
	ItemTypeTreeAdvanced:           "爱情树升阶物品",
	ItemTypeKun:                    "坤",
	ItemTypePKSkillUpgrade:         "pk技能升级道具",
	ItemTypeWedSuccRing:            "求婚成功的戒指",
	ItemTypeWaKuang:                "挖矿",
	ItemTypeQiJu:                   "棋局",
	ItemTypeSilverFuBenUpgrade:     "银两副本升级",
	ItemTypeAnQi:                   "暗器相关道具",
	ItemTypeBlessDan:               "祝福丹",
	ItemTypeUpLevelDan:             "等级直升丹",
	ItemTypeAdvancedTicket:         "直升券",
	ItemTypeBossCallTicket:         "Boss召唤券",
	ItemTypePkValue:                "红名道具",
	ItemTypeWingActivateCard:       "战翼 激活卡",
	ItemTypeGoldEquipStrengthen:    "金装相关",
	ItemTypeGemAvoidBomb:           "宝石防爆符",
	ItemTypeLuckyRate:              "幸运符",
	ItemTypeSexCard:                "性别卡",
	ItemTypeRenameCard:             "改名卡",
	ItemTypeResouceCard:            "收益卡",
	ItemTypeTower:                  "打宝相关",
	ItemTypeMyBoss:                 "个人BOSS",
	ItemTypeMountEquip:             "坐骑装备",
	ItemTypeWingStone:              "战翼符石",
	ItemTypeAnqiJiguan:             "暗器机关",
	ItemTypeJiNengShu:              "技能书",
	ItemTypeQiangHuaShi:            "强化石",
	ItemTypeZhuiZongFu:             "追踪符",
	ItemTypeTianShu:                "天书相关",
	ItemTypeDebris:                 "碎片相关",
	ItemTypeFaBaoSuit:              "法宝配饰",
	ItemTypeFaBao:                  "法宝相关",
	ItemTypeXianTi:                 "仙体相关",
	ItemTypeXianTiLingYu:           "仙体灵玉",
	ItemTypeXueDun:                 "血盾系统",
	ItemTypeFuChi:                  "扶持",
	ItemTypeDianXing:               "点星系统相关",
	ItemTypeTianMo:                 "天魔",
	ItemTypeXueMo:                  "血魔",
	ItemTypeWardrobe:               "衣橱套装相关",
	ItemTypeLingyuEquip:            "领域装备",
	ItemTypeShenfaEquip:            "身法装备",
	ItemTypeShiHunFanEquip:         "噬魂幡装备",
	ItemTypeTianMoTiEquip:          "天魔体装备",
	ItemTypeLingTongReNameCard:     "灵童改名卡",
	ItemTypeLingTong:               "灵童道具",
	ItemTypeLingTongFashion:        "灵童时装道具",
	ItemTypeLingTongWeapon:         "灵童兵魂道具",
	ItemTypeLingTongMount:          "灵童坐骑道具",
	ItemTypeLingTongWing:           "灵童战翼道具",
	ItemTypeLingTongShenFa:         "灵童身法道具",
	ItemTypeLingTongLingYu:         "灵童领域道具",
	ItemTypeLingTongFaBao:          "灵童法宝道具",
	ItemTypeLingTongXianTi:         "灵童仙体道具",
	ItemTypeLingTongMountEquip:     "灵童坐骑装备",
	ItemTypeLingTongWingEquip:      "灵童战翼装备",
	ItemTypeLingTongShenFaEquip:    "灵童身法装备",
	ItemTypeLingTongLingYuEquip:    "灵童领域装备",
	ItemTypeLingTongFaBaoEquip:     "灵童法宝装备",
	ItemTypeLingTongXianTiEquip:    "灵童仙体装备",
	ItemTypeLingTongJiNengShu:      "灵童技能书",
	ItemTypeLingTongWeaponEquip:    "灵童兵器装备",
	ItemTypeFeiSheng:               "飞升相关",
	ItemTypeHongBao:                "红包",
	ItemTypeHuaLingDan:             "化灵丹",
	ItemTypeTianFuDan:              "天赋升级丹",
	ItemTypeEquipBaoKuTicket:       "装备宝库抽奖券",
	ItemTypeShengHenEquipQingLong:  "圣痕青龙装备",
	ItemTypeShengHenEquipBaiHu:     "圣痕白虎装备",
	ItemTypeShengHenEquipZhuQue:    "圣痕朱雀装备",
	ItemTypeShengHenEquipXuanWu:    "圣痕玄武装备",
	ItemTypeSheng:                  "圣痕升级道具",
	ItemTypeMingGe:                 "命格",
	ItemTypeShenQi:                 "神器相关道具",
	ItemTypeTuLongEquip:            "屠龙装备",
	ItemTypeTuLongEquipItem:        "屠龙装备相关",
	ItemTypeYingLingPu:             "英灵谱相关",
	ItemTypeZhenFa:                 "阵法相关",
	ItemTypeHunt:                   "寻宝道具",
	ItemTypeBiaoBai:                "表白道具",
	ItemTypeBaoBao:                 "宝宝相关",
	ItemTypeBabyToy:                "宝宝玩具",
	ItemTypeDingQing:               "定情信物",
	ItemTypeBaoBaoCard:             "宝宝卡",
	ItemTypeAdditionsysTongLingDan: "附加系统通灵丹",
	ItemTypeAdditionsysJueXingDan:  "附加系统觉醒丹",
	ItemTypeYuanBaoKa:              "元宝卡",
	ItemTypeExpendBagSlotCard:      "背包扩充符",
	ItemTypeFuShi:                  "八卦符石",
	ItemTypeItemSkill:              "物品技能",
	ItemTypeBossItem:               "外域Boss、幻境Boss道具",
	ItemTypeJieYiItem:              "结义道具",
	ItemTypeXianJinYuanBaoKa:       "现金元宝卡",
	ItemTypeChuangShiZhiZhan:       "创世之战道具",
	ItemTypeBOSSMiZang:             "BOSS密藏道具",
	ItemTypeVigorousPill:           "大力丸",
	ItemTypeWushuangWeapon:         "无双神器",
	ItemTypeWushuangWeaponEssence:  "无双神器精华",
	ItemTypeGodCastingEquip:        "神铸装备",
	ItemTypeManHanQuanXi:           "满汉全席合成物品",
	ItemTypeLingTongEquip:          "灵童装备",
	ItemTypeTeRing:                 "特戒相关",
	ItemTypeShangGuZhiLing:         "上古之灵相关",
}

func (it ItemType) String() string {
	return itemTypeMap[it]
}

func (it ItemType) Valid() bool {
	switch it {
	case ItemTypeDefault,
		ItemTypeHp,
		ItemTypeExp,
		ItemTypeSilver,
		ItemTypeGold,
		ItemTypeSkill,
		ItemTypeBuff,
		ItemTypeEquipment,
		ItemTypeFashion,
		ItemTypeBingHun,
		ItemTypePetEquipment,
		ItemTypeDan,
		ItemTypeTreasure,
		ItemTypeEquipmentGem,
		ItemTypePet,
		ItemTypeLifeOrigin,
		ItemTypeSkillPoint,
		ItemTypeGem,
		ItemTypeEquipmentUpgrade,
		ItemTypeWing,
		ItemTypeSoul,
		ItemTypeKingSoul,
		ItemTypeDun,
		ItemTypeTitle,
		ItemTypeAutoUseRes,
		ItemTypeXianHua,
		ItemTypeTuMoLing,
		ItemTypeJueXue,
		ItemTypeXinFa,
		ItemTypeEctypal,
		ItemTypeAlliance,
		ItemTypeGiftBag,
		ItemTypeShenLong,
		ItemTypeCandy,
		ItemTypeWedRing,
		ItemTypeShenFa,
		ItemTypeLingyu,
		ItemTypeRingAdvaced,
		ItemTypeGoldEquip,
		ItemTypeTreeAdvanced,
		ItemTypeKun,
		ItemTypePKSkillUpgrade,
		ItemTypeWedSuccRing,
		ItemTypeWaKuang,
		ItemTypeQiJu,
		ItemTypeSilverFuBenUpgrade,
		ItemTypeAnQi,
		ItemTypeBlessDan,
		ItemTypeUpLevelDan,
		ItemTypeAdvancedTicket,
		ItemTypeBossCallTicket,
		ItemTypePkValue,
		ItemTypeWingActivateCard,
		ItemTypeGoldEquipStrengthen,
		ItemTypeGemAvoidBomb,
		ItemTypeLuckyRate,
		ItemTypeSexCard,
		ItemTypeRenameCard,
		ItemTypeResouceCard,
		ItemTypeTower,
		ItemTypeMyBoss,
		ItemTypeMountEquip,
		ItemTypeWingStone,
		ItemTypeAnqiJiguan,
		ItemTypeJiNengShu,
		ItemTypeQiangHuaShi,
		ItemTypeZhuiZongFu,
		ItemTypeTianShu,
		ItemTypeDebris,
		ItemTypeFaBaoSuit,
		ItemTypeFaBao,
		ItemTypeXianTi,
		ItemTypeXianTiLingYu,
		ItemTypeXueDun,
		ItemTypeFuChi,
		ItemTypeDianXing,
		ItemTypeTianMo,
		ItemTypeXueMo,
		ItemTypeWardrobe,
		ItemTypeLingyuEquip,
		ItemTypeShenfaEquip,
		ItemTypeShiHunFanEquip,
		ItemTypeTianMoTiEquip,
		ItemTypeLingTongReNameCard,
		ItemTypeLingTong,
		ItemTypeLingTongFashion,
		ItemTypeLingTongWeapon,
		ItemTypeLingTongMount,
		ItemTypeLingTongWing,
		ItemTypeLingTongShenFa,
		ItemTypeLingTongLingYu,
		ItemTypeLingTongFaBao,
		ItemTypeLingTongXianTi,
		ItemTypeLingTongMountEquip,
		ItemTypeLingTongWingEquip,
		ItemTypeLingTongShenFaEquip,
		ItemTypeLingTongLingYuEquip,
		ItemTypeLingTongFaBaoEquip,
		ItemTypeLingTongXianTiEquip,
		ItemTypeLingTongJiNengShu,
		ItemTypeLingTongWeaponEquip,
		ItemTypeFeiSheng,
		ItemTypeHongBao,
		ItemTypeHuaLingDan,
		ItemTypeTianFuDan,
		ItemTypeEquipBaoKuTicket,
		ItemTypeShengHenEquipQingLong,
		ItemTypeShengHenEquipBaiHu,
		ItemTypeShengHenEquipZhuQue,
		ItemTypeShengHenEquipXuanWu,
		ItemTypeSheng,
		ItemTypeMingGe,
		ItemTypeShenQi,
		ItemTypeTuLongEquip,
		ItemTypeTuLongEquipItem,
		ItemTypeYingLingPu,
		ItemTypeZhenFa,
		ItemTypeHunt,
		ItemTypeBiaoBai,
		ItemTypeBaoBao,
		ItemTypeBabyToy,
		ItemTypeDingQing,
		ItemTypeBaoBaoCard,
		ItemTypeAdditionsysTongLingDan,
		ItemTypeAdditionsysJueXingDan,
		ItemTypeYuanBaoKa,
		ItemTypeExpendBagSlotCard,
		ItemTypeFuShi,
		ItemTypeItemSkill,
		ItemTypeBossItem,
		ItemTypeJieYiItem,
		ItemTypeXianJinYuanBaoKa,
		ItemTypeChuangShiZhiZhan,
		ItemTypeBOSSMiZang,
		ItemTypeVigorousPill,
		ItemTypeWushuangWeapon,
		ItemTypeWushuangWeaponEssence,
		ItemTypeGodCastingEquip,
		ItemTypeManHanQuanXi,
		ItemTypeLingTongEquip,
		ItemTypeTeRing,
		ItemTypeShangGuZhiLing:
		return true
	}
	return false
}

type ItemSubType interface {
	SubType() int32
	Valid() bool
}

type ItemSubTypeFactory interface {
	CreateItemSubType(subType int32) ItemSubType
}

type ItemSubTypeFactoryFunc func(subType int32) ItemSubType

func (istff ItemSubTypeFactoryFunc) CreateItemSubType(subType int32) ItemSubType {
	return istff(subType)
}

//默认子类型
type ItemDefaultSubType int32

const (
	ItemDefaultSubTypeDefault ItemDefaultSubType = iota
)

func (idst ItemDefaultSubType) SubType() int32 {
	return int32(idst)
}

func (idst ItemDefaultSubType) Valid() bool {
	return true
}

func CreateItemDefaultSubType(subType int32) ItemSubType {
	return ItemDefaultSubType(subType)
}

//装备子类型
type ItemEquipmentSubType int32

const (
	//武器
	ItemEquipmentSubTypeWeapon ItemEquipmentSubType = iota
	//战袍
	ItemEquipmentSubTypeArmor
	//头盔
	ItemEquipmentSubTypeHelmet
	//战靴
	ItemEquipmentSubTypeShoe
	//腰带
	ItemEquipmentSubTypeBelt
	//护手
	ItemEquipmentSubTypeHandGuard
	//项链
	ItemEquipmentSubTypeNecklace
	//戒指
	ItemEquipmentSubTypeRing
)

var (
	equipmentBodyMap = map[ItemEquipmentSubType]inventorytypes.BodyPositionType{
		ItemEquipmentSubTypeWeapon: inventorytypes.BodyPositionTypeWeapon,

		ItemEquipmentSubTypeArmor: inventorytypes.BodyPositionTypeArmor,

		ItemEquipmentSubTypeHelmet: inventorytypes.BodyPositionTypeHelmet,

		ItemEquipmentSubTypeShoe: inventorytypes.BodyPositionTypeShoe,

		ItemEquipmentSubTypeBelt: inventorytypes.BodyPositionTypeBelt,

		ItemEquipmentSubTypeHandGuard: inventorytypes.BodyPositionTypeHandGuard,

		ItemEquipmentSubTypeNecklace: inventorytypes.BodyPositionTypeNecklace,

		ItemEquipmentSubTypeRing: inventorytypes.BodyPositionTypeRing,
	}
)

func (iest ItemEquipmentSubType) Position() inventorytypes.BodyPositionType {
	return equipmentBodyMap[iest]
}

func (iest ItemEquipmentSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemEquipmentSubType) Valid() bool {
	switch iest {
	case ItemEquipmentSubTypeWeapon,
		ItemEquipmentSubTypeArmor,
		ItemEquipmentSubTypeHelmet,
		ItemEquipmentSubTypeShoe,
		ItemEquipmentSubTypeBelt,
		ItemEquipmentSubTypeHandGuard,
		ItemEquipmentSubTypeNecklace,
		ItemEquipmentSubTypeRing:
		return true
	}
	return false
}

func CreateItemEquipmentSubType(subType int32) ItemSubType {
	return ItemEquipmentSubType(subType)
}

// 时装子类型
type ItemFashionSubType int32

const (
	//时装激活卡
	ItemFashionSubTypeActivate ItemFashionSubType = iota
	//时装碎片
	ItemFashionSubTypeMaterial
	//时装试用卡
	ItemFashionSubTypeTrialCard
)

func (iest ItemFashionSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemFashionSubType) Valid() bool {
	switch iest {
	case ItemFashionSubTypeActivate,
		ItemFashionSubTypeMaterial,
		ItemFashionSubTypeTrialCard:
		return true
	}
	return false
}

func CreateItemFashionSubType(subType int32) ItemSubType {
	return ItemFashionSubType(subType)
}

//丹药子类型
type ItemDanSubType int32

const (
	//丹药
	ItemDanSubTypeEat ItemDanSubType = iota
	//材料
	ItemDanSubTypeMaterial
)

func (iest ItemDanSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemDanSubType) Valid() bool {
	switch iest {
	case ItemDanSubTypeEat,
		ItemDanSubTypeMaterial:
		return true
	}
	return false
}

func CreateItemDanSubType(subType int32) ItemSubType {
	return ItemDanSubType(subType)
}

//坐骑子类型
type ItemMountSubType int32

const (
	//坐骑升阶丹
	ItemMountSubTypeAdvanced ItemMountSubType = iota
	//坐骑幻化丹
	ItemMountSubTypeUnreal
	//坐骑培养丹
	ItemMountSubTypeCul
	//坐骑幻化卡
	ItemMountSubTypeUnrealCard
	//坐骑碎片
	ItemMountSubTypeChip
)

func (iest ItemMountSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemMountSubType) Valid() bool {
	switch iest {
	case ItemMountSubTypeAdvanced,
		ItemMountSubTypeUnreal,
		ItemMountSubTypeCul,
		ItemMountSubTypeUnrealCard,
		ItemMountSubTypeChip:
		return true
	}
	return false
}

func CreateItemMountSubType(subType int32) ItemSubType {
	return ItemMountSubType(subType)
}

//战翼子类型
type ItemWingSubType int32

const (
	//升阶丹
	ItemWingSubTypeAdvanced ItemWingSubType = iota
	//幻化丹
	ItemWingSubTypeUnreal
	//独立幻化卡
	ItemWingSubTypeUnrealIndepend
	//战翼试用卡
	ItemWingSubTypeTrialCard
	//战翼碎片
	ItemWingSubTypeChip
)

func (iwst ItemWingSubType) SubType() int32 {
	return int32(iwst)
}

func (iwst ItemWingSubType) Valid() bool {
	switch iwst {
	case ItemWingSubTypeAdvanced,
		ItemWingSubTypeUnreal,
		ItemWingSubTypeUnrealIndepend,
		ItemWingSubTypeTrialCard,
		ItemWingSubTypeChip:
		return true
	}
	return false
}

func CreateItemWingSubType(subType int32) ItemSubType {
	return ItemWingSubType(subType)
}

//兵魂子类型
type ItemSoulSubType int32

const (
	//兵魂碎片
	ItemSoulSubTypeDebris ItemSoulSubType = iota
	//培养丹
	ItemSoulSubTypeCul
	//兵魂觉醒物品
	ItemSoulSubTypeAwaken
	//特殊兵魂激活道具
	ItemSoulSubTypeSpecialProps
	//定制冰魂激活道具
	ItemSoulSubTypeDingZhi
)

func (isst ItemSoulSubType) SubType() int32 {
	return int32(isst)
}

func (isst ItemSoulSubType) Valid() bool {
	switch isst {
	case ItemSoulSubTypeDebris,
		ItemSoulSubTypeCul,
		ItemSoulSubTypeAwaken,
		ItemSoulSubTypeSpecialProps,
		ItemSoulSubTypeDingZhi:
		return true
	}
	return false
}

func CreateItemSoulSubType(subType int32) ItemSubType {
	return ItemSoulSubType(subType)
}

//护体盾子类型
type ItemBodyShieldSubType int32

const (
	//进阶丹
	ItemBodyShieldSubTypeAdvanced ItemBodyShieldSubType = iota
	//金甲丹
	ItemBodyShieldSubTypeJJDan
)

func (ibsst ItemBodyShieldSubType) SubType() int32 {
	return int32(ibsst)
}

func (ibsst ItemBodyShieldSubType) Valid() bool {
	switch ibsst {
	case ItemBodyShieldSubTypeAdvanced,
		ItemBodyShieldSubTypeJJDan:
		return true
	}
	return false
}

func CreateItemBodyShieldSubType(subType int32) ItemSubType {
	return ItemBodyShieldSubType(subType)
}

//宝石子类型
type ItemGemSubType int32

const (
	//生命宝石
	ItemGemSubTypeHP ItemGemSubType = iota
	//攻击宝石
	ItemGemSubTypeAttack
	//防御宝石
	ItemGemSubTypeDefence
	//暴击宝石
	ItemGemSubTypeCrit
	//免暴宝石
	ItemGemSubTypeTough
	//格挡宝石
	ItemGemSubTypeBlock
	//破格宝石
	ItemGemSubTypeBreak
)

func (ibsst ItemGemSubType) SubType() int32 {
	return int32(ibsst)
}

func (ibsst ItemGemSubType) Valid() bool {
	switch ibsst {
	case ItemGemSubTypeHP,
		ItemGemSubTypeAttack,
		ItemGemSubTypeDefence,
		ItemGemSubTypeCrit,
		ItemGemSubTypeTough,
		ItemGemSubTypeBlock,
		ItemGemSubTypeBreak:
		return true
	}
	return false
}

func CreateItemGemSubType(subType int32) ItemSubType {
	return ItemGemSubType(subType)
}

//自动使用资源子类型
type ItemAutoUseResSubType int32

const (
	//0银两
	ItemAutoUseResSubTypeSilver ItemAutoUseResSubType = iota
	//元宝
	ItemAutoUseResSubTypeGold
	//绑元
	ItemAutoUseResSubTypeBindGold
	//腰牌
	ItemAutoUseResSubTypeYaoPai
	//钥匙
	ItemAutoUseResSubTypeKey
	//经验--5
	ItemAutoUseResSubTypeExp
	//原石
	ItemAutoUseResSubTypeStorage
	//普通烟花
	ItemAutoUseResSubTypeNormalFireworks
	//高级烟花
	ItemAutoUseResSubTypeSeniorFireworks
	//杀气
	ItemAutoUseResSubTypeShaQi
	//血炼值--10
	ItemAutoUseResSubTypeBloodZhi
	//浊气值
	ItemAutoUseResSubTypeZhuoQi
	//星尘值
	ItemAutoUseResSubTypeXingChen
	//功勋
	ItemAutoUseResSubTypeGongXun
	//功德值
	ItemAutoUseResSubTypeGongDe
	//装备宝库积分--15
	ItemAutoUseResSubTypeEquipBaoKuAttendPoints
	//器灵值
	ItemAutoUseResSubTypeLingQi
	//千年仙桃
	ItemAutoUseResSubTypeQianNianXianTao
	//百年仙桃
	ItemAutoUseResSubTypeBaiNianXianTao
	//神域钥匙
	ItemAutoUseResSubTypeShenYuKey
	//杀戮心-------------20
	ItemAutoUseResSubTypeShaLuXin
	//3v3积分
	ItemAutoUseResSubTypeArenaPoint
	//声威值
	ItemAutoUseResSubTypeShengWei
	//比武大会积分
	ItemAutoUseResSubTypeArenapvpJiFen
	//材料宝库积分
	ItemAutoUseResSubTypeMaterialBaoKuAttendPoints
	//创世积分----------------25
	ItemAutoUseResSubTypeChuangShiJifen
	// 经验点
	ItemAutoUseResSubTypeExpDian
	//特戒寻宝积分
	ItemAutoUseResSubTypeTeWingJiFen
)

func (iaurst ItemAutoUseResSubType) SubType() int32 {
	return int32(iaurst)
}

func (iaurst ItemAutoUseResSubType) Valid() bool {
	switch iaurst {
	case ItemAutoUseResSubTypeSilver,
		ItemAutoUseResSubTypeGold,
		ItemAutoUseResSubTypeBindGold,
		ItemAutoUseResSubTypeYaoPai,
		ItemAutoUseResSubTypeKey,
		ItemAutoUseResSubTypeExp,
		ItemAutoUseResSubTypeStorage,
		ItemAutoUseResSubTypeNormalFireworks,
		ItemAutoUseResSubTypeSeniorFireworks,
		ItemAutoUseResSubTypeShaQi,
		ItemAutoUseResSubTypeBloodZhi,
		ItemAutoUseResSubTypeZhuoQi,
		ItemAutoUseResSubTypeXingChen,
		ItemAutoUseResSubTypeGongXun,
		ItemAutoUseResSubTypeGongDe,
		ItemAutoUseResSubTypeEquipBaoKuAttendPoints,
		ItemAutoUseResSubTypeLingQi,
		ItemAutoUseResSubTypeQianNianXianTao,
		ItemAutoUseResSubTypeBaiNianXianTao,
		ItemAutoUseResSubTypeShenYuKey,
		ItemAutoUseResSubTypeShaLuXin,
		ItemAutoUseResSubTypeArenaPoint,
		ItemAutoUseResSubTypeShengWei,
		ItemAutoUseResSubTypeArenapvpJiFen,
		ItemAutoUseResSubTypeMaterialBaoKuAttendPoints,
		ItemAutoUseResSubTypeChuangShiJifen,
		ItemAutoUseResSubTypeExpDian,
		ItemAutoUseResSubTypeTeWingJiFen:
		return true
	}
	return false
}

func CreateItemAutoUseResSubType(subType int32) ItemSubType {
	return ItemAutoUseResSubType(subType)
}

//屠魔令子类型
type ItemTuMoLingSubType int32

const (
	//绿
	ItemTuMoLingSubTypeGreen ItemTuMoLingSubType = iota
	//蓝
	ItemTuMoLingSubTypeBlue
	//紫
	ItemTuMoLingSubTypePurple
	//橙
	ItemTuMoLingSubTypeOrange
)

func (t ItemTuMoLingSubType) SubType() int32 {
	return int32(t)
}

func (t ItemTuMoLingSubType) Valid() bool {
	switch t {
	case ItemTuMoLingSubTypeGreen,
		ItemTuMoLingSubTypeBlue,
		ItemTuMoLingSubTypePurple,
		ItemTuMoLingSubTypeOrange:
		return true
	}
	return false
}

func CreateItemTuMoLingSubType(subType int32) ItemSubType {
	return ItemTuMoLingSubType(subType)
}

//绝学子类型
type ItemJueXueSubType int32

const (
	//绝学技能书
	ItemJueXueSubTypeJiNeng ItemJueXueSubType = iota
	//升级秘籍
	ItemJueXueSubTypeUpgrade
	//顿悟
	ItemJueXueSubTypeDunWu
	//绝学技能书碎片
	ItemJueXueSubTypeSkillBookDebris
)

func (t ItemJueXueSubType) SubType() int32 {
	return int32(t)
}

func (t ItemJueXueSubType) Valid() bool {
	switch t {
	case ItemJueXueSubTypeJiNeng,
		ItemJueXueSubTypeUpgrade,
		ItemJueXueSubTypeDunWu,
		ItemJueXueSubTypeSkillBookDebris:
		return true
	}
	return false
}

func CreateItemJueXueSubType(subType int32) ItemSubType {
	return ItemJueXueSubType(subType)
}

//心法子类型
type ItemXinFaSubType int32

const (
	//心法书
	ItemXinFaSubTypeActive ItemXinFaSubType = iota
	//升级秘籍
	ItemXinFaSubTypeUpgrade
)

func (t ItemXinFaSubType) SubType() int32 {
	return int32(t)
}

func (t ItemXinFaSubType) Valid() bool {
	switch t {
	case ItemXinFaSubTypeActive,
		ItemXinFaSubTypeUpgrade:
		return true
	}
	return false
}

func CreateItemXinFaSubType(subType int32) ItemSubType {
	return ItemXinFaSubType(subType)
}

//副本门票子类型
type ItemEctypalSubType int32

const (
	//银两副本入场券
	ItemEctypalSubTypeSilver ItemEctypalSubType = iota
	//经验副本入场券
	ItemEctypalSubTypeExp
	//材料副本入场券
	ItemEctypalSubTypeMaterial
)

func (t ItemEctypalSubType) SubType() int32 {
	return int32(t)
}

func (t ItemEctypalSubType) Valid() bool {
	switch t {
	case ItemEctypalSubTypeSilver,
		ItemEctypalSubTypeExp,
		ItemEctypalSubTypeMaterial:
		return true
	}
	return false
}

func CreateItemEctypalSubType(subType int32) ItemSubType {
	return ItemEctypalSubType(subType)
}

//仙盟类型
type ItemAllianceSubType int32

const (
	//仙盟创建
	ItemAllianceSubTypeCreate ItemAllianceSubType = iota
	//仙盟建设
	ItemAllianceSubTypeBuild
	//虎符
	ItemAllianceSubTypeHuFu
	//虎符碎片
	ItemAllianceSubTypeHuFuPart
	//酒酿
	ItemAllianceSubTypeJiuNiang
	//合盟改名卡
	ItemAllianceSubTypeHeMengGaiMingKa
)

func (t ItemAllianceSubType) SubType() int32 {
	return int32(t)
}

func (t ItemAllianceSubType) Valid() bool {
	switch t {
	case ItemAllianceSubTypeCreate,
		ItemAllianceSubTypeBuild,
		ItemAllianceSubTypeHuFu,
		ItemAllianceSubTypeHuFuPart,
		ItemAllianceSubTypeJiuNiang,
		ItemAllianceSubTypeHeMengGaiMingKa:
		return true
	}
	return false
}

func CreateItemAllianceSubType(subType int32) ItemSubType {
	return ItemAllianceSubType(subType)
}

//宝箱子类型
type ItemGiftBagSubType int32

const (
	//普通宝箱
	ItemGiftBagSubTypeBox ItemGiftBagSubType = iota
	//乾坤袋
	ItemGiftBagSubTypeQianKunDai
	//跨服宝箱
	ItemGiftBagSubTypeCorss
	//消耗选择宝箱
	ItemGiftBagSubTypeCostChoose
)

func (t ItemGiftBagSubType) SubType() int32 {
	return int32(t)
}

func (t ItemGiftBagSubType) Valid() bool {
	switch t {
	case ItemGiftBagSubTypeBox,
		ItemGiftBagSubTypeQianKunDai,
		ItemGiftBagSubTypeCorss,
		ItemGiftBagSubTypeCostChoose:
		return true
	}
	return false
}

func CreateItemGiftBagSubType(subType int32) ItemSubType {
	return ItemGiftBagSubType(subType)
}

//喜糖子类型
type ItemCandySubType int32

const (
	//经验喜糖
	ItemCandySubTypeExp ItemCandySubType = iota
	//银两喜糖
	ItemCandySubTypeSilver
)

func (t ItemCandySubType) SubType() int32 {
	return int32(t)
}

func (t ItemCandySubType) Valid() bool {
	switch t {
	case ItemCandySubTypeExp,
		ItemCandySubTypeSilver:
		return true
	}
	return false
}

func CreateItemCandySubType(subType int32) ItemSubType {
	return ItemCandySubType(subType)
}

//婚戒子类型
type ItemWedRingSubType int32

const (
	//青铜对戒
	ItemWedRingSubTypeBronze ItemWedRingSubType = iota
	//紫金对戒
	ItemWedRingSubTypePurple
	//龙风对戒
	ItemWedRingSubTypeLongFeng
	//青铜对戒
	ItemWedRingSubTypeBronzeCheap
	//紫金对戒
	ItemWedRingSubTypePurpleCheap
	//龙风对戒
	ItemWedRingSubTypeLongFengCheap
	//豪华青铜对戒
	ItemWedRingSubTypeBronzeLuxury
	//豪华紫金对戒
	ItemWedRingSubTypePurpleLuxury
	//豪华龙风对戒
	ItemWedRingSubTypeLongFengLuxury
)

func (t ItemWedRingSubType) SubType() int32 {
	return int32(t)
}

func (t ItemWedRingSubType) Valid() bool {
	switch t {
	case ItemWedRingSubTypeBronze,
		ItemWedRingSubTypePurple,
		ItemWedRingSubTypeLongFeng,
		ItemWedRingSubTypeBronzeCheap,
		ItemWedRingSubTypePurpleCheap,
		ItemWedRingSubTypeLongFengCheap,
		ItemWedRingSubTypeBronzeLuxury,
		ItemWedRingSubTypePurpleLuxury,
		ItemWedRingSubTypeLongFengLuxury:
		return true
	}
	return false
}

func CreateItemWedRingSubType(subType int32) ItemSubType {
	return ItemWedRingSubType(subType)
}

//身法子类型
type ItemShenFaSubType int32

const (
	//升阶丹
	ItemShenfaSubTypeAdvanced ItemShenFaSubType = iota
	//幻化丹
	ItemShenfaSubTypeUnreal
	//独立的幻化卡
	ItemShenfaSubTyoeUnrealCard
)

func (t ItemShenFaSubType) SubType() int32 {
	return int32(t)
}

func (t ItemShenFaSubType) Valid() bool {
	switch t {
	case ItemShenfaSubTypeAdvanced,
		ItemShenfaSubTypeUnreal,
		ItemShenfaSubTyoeUnrealCard:
		return true
	}
	return false
}

func CreateItemShenFaSubType(subType int32) ItemSubType {
	return ItemShenFaSubType(subType)
}

//领域子类型
type ItemLingyuSubType int32

const (
	//升阶丹
	ItemLingyuSubTypeAdvanced ItemLingyuSubType = iota
	//幻化丹
	ItemLingyuSubTypeUnreal
	//独立的幻化卡
	ItemLingyuSubTypeUnrealCard
)

func (t ItemLingyuSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingyuSubType) Valid() bool {
	switch t {
	case ItemLingyuSubTypeAdvanced,
		ItemLingyuSubTypeUnreal,
		ItemLingyuSubTypeUnrealCard:
		return true
	}
	return false
}

func CreateItemLingyuSubType(subType int32) ItemSubType {
	return ItemLingyuSubType(subType)
}

//元神金装子类型
type ItemGoldEquipSubType int32

const (
	//武器
	ItemGoldEquipSubTypeWeapon ItemGoldEquipSubType = iota
	//战袍
	ItemGoldEquipSubTypeArmor
	//头盔
	ItemGoldEquipSubTypeHelmet
	//战靴
	ItemGoldEquipSubTypeShoe
	//护腿
	ItemGoldEquipSubTypeBelt
	//护手
	ItemGoldEquipSubTypeHandGuard
	//玉坠
	ItemGoldEquipSubTypeNecklace
	//项链
	ItemGoldEquipSubTypeRing
)

var (
	goldEquipBodyMap = map[ItemGoldEquipSubType]inventorytypes.BodyPositionType{
		ItemGoldEquipSubTypeWeapon:    inventorytypes.BodyPositionTypeWeapon,
		ItemGoldEquipSubTypeArmor:     inventorytypes.BodyPositionTypeArmor,
		ItemGoldEquipSubTypeHelmet:    inventorytypes.BodyPositionTypeHelmet,
		ItemGoldEquipSubTypeShoe:      inventorytypes.BodyPositionTypeShoe,
		ItemGoldEquipSubTypeBelt:      inventorytypes.BodyPositionTypeBelt,
		ItemGoldEquipSubTypeHandGuard: inventorytypes.BodyPositionTypeHandGuard,
		ItemGoldEquipSubTypeNecklace:  inventorytypes.BodyPositionTypeNecklace,
		ItemGoldEquipSubTypeRing:      inventorytypes.BodyPositionTypeRing,
	}
)

func (iest ItemGoldEquipSubType) Position() inventorytypes.BodyPositionType {
	return goldEquipBodyMap[iest]
}

func (t ItemGoldEquipSubType) SubType() int32 {
	return int32(t)
}

func (t ItemGoldEquipSubType) Valid() bool {
	switch t {
	case ItemGoldEquipSubTypeWeapon,
		ItemGoldEquipSubTypeArmor,
		ItemGoldEquipSubTypeHelmet,
		ItemGoldEquipSubTypeShoe,
		ItemGoldEquipSubTypeBelt,
		ItemGoldEquipSubTypeHandGuard,
		ItemGoldEquipSubTypeNecklace,
		ItemGoldEquipSubTypeRing:
		return true
	}
	return false
}

func CreateItemGoldEquipSubType(subType int32) ItemSubType {
	return ItemGoldEquipSubType(subType)
}

//暗器子类型
type ItemAnQiSubType int32

const (
	//升阶丹
	ItemAnQiSubTypeAdvanced ItemAnQiSubType = iota
	//培养丹
	ItemAnQiSubTypePeiYang
)

func (t ItemAnQiSubType) SubType() int32 {
	return int32(t)
}

func (t ItemAnQiSubType) Valid() bool {
	switch t {
	case ItemAnQiSubTypeAdvanced,
		ItemAnQiSubTypePeiYang:
		return true
	}
	return false
}

func CreateItemAnQiSubType(subType int32) ItemSubType {
	return ItemAnQiSubType(subType)
}

//祝福丹子类型
type ItemBlessDanSubType int32

const (
	// 0通用祝福丹
	ItemBlessDanSubTypeCommon ItemBlessDanSubType = iota
	// 1坐骑祝福丹
	ItemBlessDanSubTypeMount ItemBlessDanSubType = 1
	// 2战翼祝福丹
	ItemBlessDanSubTypeWing ItemBlessDanSubType = 2
	// 3暗器祝福丹
	ItemBlessDanSubTypeAnqi ItemBlessDanSubType = 3
	// 4护体盾祝福丹
	ItemBlessDanSubTypeBodyShield ItemBlessDanSubType = 4
	// 5身法祝福丹
	ItemBlessDanSubTypeShenfa ItemBlessDanSubType = 5
	// 6领域祝福丹
	ItemBlessDanSubTypeLingyu ItemBlessDanSubType = 6
	// 7仙羽祝福丹
	ItemBlessDanSubTypeFeather ItemBlessDanSubType = 7
	// 8盾刺祝福丹
	ItemBlessDanSubTypeShield ItemBlessDanSubType = 8
	// 9法宝祝福丹
	ItemBlessDanSubTypeFaBao ItemBlessDanSubType = 9
	// 10仙体祝福丹
	ItemBlessDanSubTypeXianTi ItemBlessDanSubType = 10
	// 11噬魂幡祝福丹
	ItemBlessDanSubTypeShiHunFan ItemBlessDanSubType = 11
	// 12天魔体祝福丹
	ItemBlessDanSubTypeTianMoTi ItemBlessDanSubType = 12

	// 101灵兵祝福丹
	ItemBlessDanSubTypeLingBing ItemBlessDanSubType = 101
	// 102灵骑祝福丹
	ItemBlessDanSubTypeLingQi ItemBlessDanSubType = 102
	// 103灵翼祝福丹
	ItemBlessDanSubTypeLingYi ItemBlessDanSubType = 103
	// 104灵身祝福丹
	ItemBlessDanSubTypeLingShen ItemBlessDanSubType = 104
	// 105灵域祝福丹
	ItemBlessDanSubTypeLingTongYu ItemBlessDanSubType = 105
	// 106灵宝祝福丹
	ItemBlessDanSubTypeLingBao ItemBlessDanSubType = 106
	// 107灵体祝福丹
	ItemBlessDanSubTypeLingTi ItemBlessDanSubType = 107
)

func (t ItemBlessDanSubType) SubType() int32 {
	return int32(t)
}

func (t ItemBlessDanSubType) Valid() bool {
	switch t {
	case ItemBlessDanSubTypeCommon,
		ItemBlessDanSubTypeMount,
		ItemBlessDanSubTypeWing,
		ItemBlessDanSubTypeAnqi,
		ItemBlessDanSubTypeBodyShield,
		ItemBlessDanSubTypeShenfa,
		ItemBlessDanSubTypeLingyu,
		ItemBlessDanSubTypeFeather,
		ItemBlessDanSubTypeShield,
		ItemBlessDanSubTypeFaBao,
		ItemBlessDanSubTypeXianTi,
		ItemBlessDanSubTypeShiHunFan,
		ItemBlessDanSubTypeTianMoTi,
		ItemBlessDanSubTypeLingBing,
		ItemBlessDanSubTypeLingQi,
		ItemBlessDanSubTypeLingYi,
		ItemBlessDanSubTypeLingShen,
		ItemBlessDanSubTypeLingTongYu,
		ItemBlessDanSubTypeLingBao,
		ItemBlessDanSubTypeLingTi:
		return true
	}
	return false
}

func CreateItemBlessDanSubType(subType int32) ItemSubType {
	return ItemBlessDanSubType(subType)
}

//直升券子类型
type ItemAdvancedTicketSubType int32

const (
	// 0通用直升券
	ItemAdvancedTicketSubTypeCommon ItemAdvancedTicketSubType = iota
	// 1坐骑直升券
	ItemAdvancedTicketSubTypeMount ItemAdvancedTicketSubType = 1
	// 2战翼直升券
	ItemAdvancedTicketSubTypeWing ItemAdvancedTicketSubType = 2
	// 3暗器直升券
	ItemAdvancedTicketSubTypeAnqi ItemAdvancedTicketSubType = 3
	// 4护体盾直升券
	ItemAdvancedTicketSubTypeBodyShield ItemAdvancedTicketSubType = 4
	// 5身法直升券
	ItemAdvancedTicketSubTypeShenfa ItemAdvancedTicketSubType = 5
	// 6领域直升券
	ItemAdvancedTicketSubTypeLingyu ItemAdvancedTicketSubType = 6
	// 7仙羽直升券
	ItemAdvancedTicketSubTypeFeather ItemAdvancedTicketSubType = 7
	// 8盾刺直升券
	ItemAdvancedTicketSubTypeShield ItemAdvancedTicketSubType = 8
	// 9法宝直升券
	ItemAdvancedTicketSubTypeFaBao ItemAdvancedTicketSubType = 9
	// 10仙体直升券
	ItemAdvancedTicketSubTypeXianTi ItemAdvancedTicketSubType = 10
	// 11噬魂幡直升券
	ItemAdvancedTicketSubTypeShiHunFan ItemAdvancedTicketSubType = 11
	// 12天魔体直升券
	ItemAdvancedTicketSubTypeTianMoTi ItemAdvancedTicketSubType = 12

	// 101灵兵直升券
	ItemAdvancedTicketSubTypeLingBing ItemAdvancedTicketSubType = 101
	// 102灵骑直升券
	ItemAdvancedTicketSubTypeLingQi ItemAdvancedTicketSubType = 102
	// 103灵翼直升券
	ItemAdvancedTicketSubTypeLingYi ItemAdvancedTicketSubType = 103
	// 104灵身直升券
	ItemAdvancedTicketSubTypeLingShen ItemAdvancedTicketSubType = 104
	// 105灵域直升券
	ItemAdvancedTicketSubTypeLingTongYu ItemAdvancedTicketSubType = 105
	// 106灵宝直升券
	ItemAdvancedTicketSubTypeLingBao ItemAdvancedTicketSubType = 106
	// 107灵体直升券
	ItemAdvancedTicketSubTypeLingTi ItemAdvancedTicketSubType = 107
)

func (t ItemAdvancedTicketSubType) SubType() int32 {
	return int32(t)
}

func (t ItemAdvancedTicketSubType) Valid() bool {
	switch t {
	case ItemAdvancedTicketSubTypeCommon,
		ItemAdvancedTicketSubTypeMount,
		ItemAdvancedTicketSubTypeWing,
		ItemAdvancedTicketSubTypeAnqi,
		ItemAdvancedTicketSubTypeBodyShield,
		ItemAdvancedTicketSubTypeShenfa,
		ItemAdvancedTicketSubTypeLingyu,
		ItemAdvancedTicketSubTypeFeather,
		ItemAdvancedTicketSubTypeShield,
		ItemAdvancedTicketSubTypeFaBao,
		ItemAdvancedTicketSubTypeXianTi,
		ItemAdvancedTicketSubTypeShiHunFan,
		ItemAdvancedTicketSubTypeTianMoTi,
		ItemAdvancedTicketSubTypeLingBing,
		ItemAdvancedTicketSubTypeLingQi,
		ItemAdvancedTicketSubTypeLingYi,
		ItemAdvancedTicketSubTypeLingShen,
		ItemAdvancedTicketSubTypeLingTongYu,
		ItemAdvancedTicketSubTypeLingBao,
		ItemAdvancedTicketSubTypeLingTi:
		return true
	}
	return false
}

func CreateItemAdvancedTicketSubType(subType int32) ItemSubType {
	return ItemAdvancedTicketSubType(subType)
}

//鲜花子类型
type ItemXueHuaSubType int32

const (
	//普通鲜花
	ItemXueHuaSubTypeNormal ItemXueHuaSubType = iota
	//普通花束
	ItemXueHuaSubTypeBouquet
	//中档花束
	ItemXueHuaSubTypeMidBouquet
	//高级花束
	ItemXueHuaSubTypeSeniorBouquet
)

func (iest ItemXueHuaSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemXueHuaSubType) Valid() bool {
	switch iest {
	case ItemXueHuaSubTypeNormal,
		ItemXueHuaSubTypeBouquet,
		ItemXueHuaSubTypeMidBouquet,
		ItemXueHuaSubTypeSeniorBouquet:
		return true
	}
	return false
}

func CreateItemXueHuaSubType(subType int32) ItemSubType {
	return ItemXueHuaSubType(subType)
}

//幸运符子类型
type ItemLuckySubType int32

const (
	//0攻击宝石幸运符
	ItemLuckySubTypeAttackLucky ItemLuckySubType = iota
	//1生命宝石幸运符
	ItemLuckySubTypeHpLucky ItemLuckySubType = 1
	//2防御宝石幸运符
	ItemLuckySubTypeDefenceLucky ItemLuckySubType = 2
	//3坐骑幸运符
	ItemLuckySubTypeMountLucky ItemLuckySubType = 3
	//4战翼幸运符
	ItemLuckySubTypeWingLucky ItemLuckySubType = 4
	//5暗器幸运符
	ItemLuckySubTypeAnqiLucky ItemLuckySubType = 5
	//6护体盾幸运符
	ItemLuckySubTypeBodyShieldLucky ItemLuckySubType = 6
	//7身法幸运符
	ItemLuckySubTypeShenfaLucky ItemLuckySubType = 7
	//8领域幸运符
	ItemLuckySubTypeLingyuLucky ItemLuckySubType = 8
	//9仙羽幸运符
	ItemLuckySubTypeFeatherLucky ItemLuckySubType = 9
	//10盾刺幸运符
	ItemLuckySubTypeShieldLucky ItemLuckySubType = 10
	//11法宝幸运符
	ItemLuckySubTypeFaBaoLucky ItemLuckySubType = 11
	//12仙体幸运符
	ItemLuckySubTypeXianTiLucky ItemLuckySubType = 12
	//13噬魂幡幸运符
	ItemLuckySubTypeShiHunFanLucky ItemLuckySubType = 13
	//14天魔体幸运符
	ItemLuckySubTypeTianMoTiLucky ItemLuckySubType = 14

	//101灵兵幸运符
	ItemLuckySubTypeLingBingLucky ItemLuckySubType = 101
	//102灵骑幸运符
	ItemLuckySubTypeLingQiLucky ItemLuckySubType = 102
	//103灵翼幸运符
	ItemLuckySubTypeLingYiLucky ItemLuckySubType = 103
	//104灵身幸运符
	ItemLuckySubTypeLingShenLucky ItemLuckySubType = 104
	//105灵域幸运符
	ItemLuckySubTypeLingTongYuLucky ItemLuckySubType = 105
	//106灵宝幸运符
	ItemLuckySubTypeLingBaoLucky ItemLuckySubType = 106
	//107灵体幸运符
	ItemLuckySubTypeLingTiLucky ItemLuckySubType = 107
)

func (iest ItemLuckySubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemLuckySubType) Valid() bool {
	switch iest {
	case ItemLuckySubTypeAttackLucky,
		ItemLuckySubTypeHpLucky,
		ItemLuckySubTypeDefenceLucky,
		ItemLuckySubTypeMountLucky,
		ItemLuckySubTypeWingLucky,
		ItemLuckySubTypeAnqiLucky,
		ItemLuckySubTypeBodyShieldLucky,
		ItemLuckySubTypeShenfaLucky,
		ItemLuckySubTypeLingyuLucky,
		ItemLuckySubTypeFeatherLucky,
		ItemLuckySubTypeShieldLucky,
		ItemLuckySubTypeFaBaoLucky,
		ItemLuckySubTypeXianTiLucky,
		ItemLuckySubTypeShiHunFanLucky,
		ItemLuckySubTypeTianMoTiLucky,
		ItemLuckySubTypeLingBingLucky,
		ItemLuckySubTypeLingQiLucky,
		ItemLuckySubTypeLingYiLucky,
		ItemLuckySubTypeLingShenLucky,
		ItemLuckySubTypeLingTongYuLucky,
		ItemLuckySubTypeLingBaoLucky,
		ItemLuckySubTypeLingTiLucky:
		return true
	}
	return false
}

func CreateItemLuckySubType(subType int32) ItemSubType {
	return ItemLuckySubType(subType)
}

//收益卡子类型
type ItemResourceCardSubType int32

const (
	//经验符
	ItemResourceCardSubTypeExp ItemResourceCardSubType = iota
	//掉宝符
	ItemResourceCardSubTypeDrop
)

func (iest ItemResourceCardSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemResourceCardSubType) Valid() bool {
	switch iest {
	case ItemResourceCardSubTypeExp,
		ItemResourceCardSubTypeDrop:
		return true
	}
	return false
}

func CreateItemResourceCardSubType(subType int32) ItemSubType {
	return ItemResourceCardSubType(subType)
}

//打宝塔相关子类型
type ItemTowerSubType int32

const (
	//时间沙漏
	ItemTowerSubTypeTimeCard ItemTowerSubType = iota
	//直飞符
	ItemTowerSubTypeJumpCard
)

func (iest ItemTowerSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemTowerSubType) Valid() bool {
	switch iest {
	case ItemTowerSubTypeTimeCard,
		ItemTowerSubTypeJumpCard:
		return true
	}
	return false
}

func CreateItemTowerSubType(subType int32) ItemSubType {
	return ItemTowerSubType(subType)
}

//天书相关子类型
type ItemTianShuSubType int32

const (
	//财富残卷
	ItemTianShuSubTypeSilver ItemTianShuSubType = iota
	//经验残卷
	ItemTianShuSubTypeExp
	//掉宝残卷
	ItemTianShuSubTypeDrop
	//祝福残卷
	ItemTianShuSubTypeAdvanced
	//BOSS残卷
	ItemTianShuSubTypeBoss
	//绑元残卷
	ItemTianShuSubTypeBindGold
	//元宝残卷
	ItemTianShuSubTypeGold
	//创世残卷
	ItemTianShuSubTypeChuangShi
)

func (iest ItemTianShuSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemTianShuSubType) Valid() bool {
	switch iest {
	case ItemTianShuSubTypeSilver,
		ItemTianShuSubTypeExp,
		ItemTianShuSubTypeDrop,
		ItemTianShuSubTypeAdvanced,
		ItemTianShuSubTypeBoss,
		ItemTianShuSubTypeBindGold,
		ItemTianShuSubTypeGold,
		ItemTianShuSubTypeChuangShi:
		return true
	}
	return false
}

func CreateItemTianShuSubType(subType int32) ItemSubType {
	return ItemTianShuSubType(subType)
}

//金装强化相关子类型
type ItemGoldEquipStrengthenSubType int32

const (
	//天工锤
	ItemGoldEquipStrengthenSubTypeChuiZi ItemGoldEquipStrengthenSubType = iota
	//金装强化石
	ItemGoldEquipStrengthenSubTypeQiangHuaShi
	//金装开光石
	ItemGoldEquipStrengthenSubTypeKaiGuangShi
	//金装升星石
	ItemGoldEquipStrengthenSubTypeShengXingShi
	//开光钻
	ItemGoldEquipStrengthenSubTypeKaiGuangZuan
	//5强化圣石
	ItemGoldEquipStrengthenSubTypeQiangHuaShengShi
	//装备·继承
	ItemGoldEquipStrengthenSubTypeExtend
)

func (iest ItemGoldEquipStrengthenSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemGoldEquipStrengthenSubType) Valid() bool {
	switch iest {
	case ItemGoldEquipStrengthenSubTypeChuiZi,
		ItemGoldEquipStrengthenSubTypeQiangHuaShi,
		ItemGoldEquipStrengthenSubTypeKaiGuangShi,
		ItemGoldEquipStrengthenSubTypeShengXingShi,
		ItemGoldEquipStrengthenSubTypeKaiGuangZuan,
		ItemGoldEquipStrengthenSubTypeQiangHuaShengShi,
		ItemGoldEquipStrengthenSubTypeExtend:
		return true
	}
	return false
}

func CreateItemGoldEquipStrengthenSubType(subType int32) ItemSubType {
	return ItemGoldEquipStrengthenSubType(subType)
}

//法宝配饰子类型
type ItemFaBaoSuitSubType int32

const (
	//灵魂
	ItemFaBaoSuitSubTypeLingHun ItemFaBaoSuitSubType = iota
	//灵华
	ItemFaBaoSuitSubTypeLingHua
	//灵饰
	ItemFaBaoSuitSubTypeLingShi
	//灵器
	ItemFaBaoSuitSubTypeLingQi
)

func (t ItemFaBaoSuitSubType) SubType() int32 {
	return int32(t)
}

func (t ItemFaBaoSuitSubType) Valid() bool {
	switch t {
	case ItemFaBaoSuitSubTypeLingHun,
		ItemFaBaoSuitSubTypeLingHua,
		ItemFaBaoSuitSubTypeLingShi,
		ItemFaBaoSuitSubTypeLingQi:
		return true
	}
	return false
}

func CreateItemFaBaoSuitSubType(subType int32) ItemSubType {
	return ItemFaBaoSuitSubType(subType)
}

//法宝子类型
type ItemFaBaoSubType int32

const (
	//法宝升阶丹
	ItemFaBaoSubTypeAdvanced ItemFaBaoSubType = iota
	//法宝幻化丹
	ItemFaBaoSubTypeUnreal
	//法宝精华
	ItemFaBaoSubTypeJingHua
	//法宝碎片
	ItemFaBaoSubTypeDebris
	//法宝皮肤碎片
	ItemFaBaoSubTypeSkinDebris
)

func (t ItemFaBaoSubType) SubType() int32 {
	return int32(t)
}

func (t ItemFaBaoSubType) Valid() bool {
	switch t {
	case ItemFaBaoSubTypeAdvanced,
		ItemFaBaoSubTypeUnreal,
		ItemFaBaoSubTypeJingHua,
		ItemFaBaoSubTypeDebris,
		ItemFaBaoSubTypeSkinDebris:
		return true
	}
	return false
}

func CreateItemFaBaoSubType(subType int32) ItemSubType {
	return ItemFaBaoSubType(subType)
}

//血盾子类型
type ItemXueDunSubType int32

const (
	//血盾碎片
	ItemXueDunSubTypeDebris ItemXueDunSubType = iota
	//血盾升星
	ItemXueDunSubTypeUpstar
)

func (t ItemXueDunSubType) SubType() int32 {
	return int32(t)
}

func (t ItemXueDunSubType) Valid() bool {
	switch t {
	case ItemXueDunSubTypeUpstar,
		ItemXueDunSubTypeDebris:
		return true
	}
	return false
}

func CreateItemXueDunSubType(subType int32) ItemSubType {
	return ItemXueDunSubType(subType)
}

//帝魂子类型
type ItemKingSoulSubType int32

const (
	//帝魂碎片
	ItemKingSoulSubTypeDebris ItemKingSoulSubType = iota
	//帝魂激活物
	ItemKingSoulSubTypeActive
	//帝魂觉醒残页
	ItemKingSoulSubTypeAwakenPage
	//帝魂觉醒物品
	ItemKingSoulSubTypeAwaken
)

func (t ItemKingSoulSubType) SubType() int32 {
	return int32(t)
}

func (t ItemKingSoulSubType) Valid() bool {
	switch t {
	case ItemKingSoulSubTypeDebris,
		ItemKingSoulSubTypeActive,
		ItemKingSoulSubTypeAwakenPage,
		ItemKingSoulSubTypeAwaken:
		return true
	}
	return false
}

func CreateItemKingSoulSubType(subType int32) ItemSubType {
	return ItemKingSoulSubType(subType)
}

//天魔子类型
type ItemTianMoSubType int32

const (
	//天魔升阶丹
	ItemTianMoSubTypeAdvanced ItemTianMoSubType = iota
	//天魔培养丹
	ItemTianMoSubTypePeiYang
)

func (t ItemTianMoSubType) SubType() int32 {
	return int32(t)
}

func (t ItemTianMoSubType) Valid() bool {
	switch t {
	case ItemTianMoSubTypeAdvanced,
		ItemTianMoSubTypePeiYang:
		return true
	}
	return false
}

func CreateItemTianMoSubType(subType int32) ItemSubType {
	return ItemTianMoSubType(subType)
}

//创世之战子类型
type ItemChuangShiZhiZhanSubType int32

const (
	//仙盟阵营更改令
	ItemChuangShiZhiZhanSubTypeChangeZhenYing ItemChuangShiZhiZhanSubType = iota
	//小喇叭
	ItemChuangShiZhiZhanSubTypeXiaoLaBa
	//城防建设令
	ItemChuangShiZhiZhanSubTypeChengFang
	//雇佣兵召唤令
	ItemChuangShiZhiZhanSubTypeGuYongBing
	//天气激活道具
	ItemChuangShiZhiZhanSubTypeTianQi
	//万法自然道具书
	ItemChuangShiZhiZhanSubTypeWanFa
	//玩家阵营更改令
	ItemChuangShiZhiZhanSubTypeChangeZhenYingPlayer
)

func (t ItemChuangShiZhiZhanSubType) SubType() int32 {
	return int32(t)
}

func (t ItemChuangShiZhiZhanSubType) Valid() bool {
	switch t {
	case ItemChuangShiZhiZhanSubTypeChangeZhenYing,
		ItemChuangShiZhiZhanSubTypeXiaoLaBa,
		ItemChuangShiZhiZhanSubTypeChengFang,
		ItemChuangShiZhiZhanSubTypeGuYongBing,
		ItemChuangShiZhiZhanSubTypeTianQi,
		ItemChuangShiZhiZhanSubTypeWanFa,
		ItemChuangShiZhiZhanSubTypeChangeZhenYingPlayer:
		return true
	}
	return false
}

func CreateItemChuangShiZhiZhanSubType(subType int32) ItemSubType {
	return ItemChuangShiZhiZhanSubType(subType)
}

//无双神器子类型
type ItemWushuangWeaponSubType int32

const (
	ItemWushuangWeaponSubTypeWeapon ItemWushuangWeaponSubType = iota
	ItemWushuangWeaponSubTypeCloths
	ItemWushuangWeaponSubTypeHead
	ItemWushuangWeaponSubTypeShoes
	ItemWushuangWeaponSubTypeNecklace
	ItemWushuangWeaponSubTypePendant
)

func (t ItemWushuangWeaponSubType) SubType() int32 {
	return int32(t)
}

func (t ItemWushuangWeaponSubType) Valid() bool {
	switch t {
	case ItemWushuangWeaponSubTypeWeapon,
		ItemWushuangWeaponSubTypeCloths,
		ItemWushuangWeaponSubTypeHead,
		ItemWushuangWeaponSubTypeShoes,
		ItemWushuangWeaponSubTypeNecklace,
		ItemWushuangWeaponSubTypePendant:
		return true
	}
	return false
}

func CreateItemWushuangWeaponSubType(subType int32) ItemSubType {
	return ItemWushuangWeaponSubType(subType)
}

// 无双神器精华子类
type ItemWushuangWeaponEssenceSubType int32

const (
	ItemWushuangWeaponEssenceSubTypeEssence ItemWushuangWeaponEssenceSubType = iota
	ItemWushuangWeaponEssenceSubTypeMaterial
)

func (t ItemWushuangWeaponEssenceSubType) SubType() int32 {
	return int32(t)
}

func (t ItemWushuangWeaponEssenceSubType) Valid() bool {
	switch t {
	case ItemWushuangWeaponEssenceSubTypeEssence,
		ItemWushuangWeaponEssenceSubTypeMaterial:
		return true
	}
	return false
}

func CreateItemWushuangWeaponEssenceSubType(subType int32) ItemSubType {
	return ItemWushuangWeaponEssenceSubType(subType)
}

//BOSS密藏子类型
type ItemBOSSMiZangSubType int32

const (
	//银铲子
	ItemBOSSMiZangSubTypeYinChanZi ItemBOSSMiZangSubType = iota
	//金铲子
	ItemBOSSMiZangSubTypeJinChanZi
)

func (t ItemBOSSMiZangSubType) SubType() int32 {
	return int32(t)
}

func (t ItemBOSSMiZangSubType) Valid() bool {
	switch t {
	case ItemBOSSMiZangSubTypeYinChanZi,
		ItemBOSSMiZangSubTypeJinChanZi:
		return true
	}
	return false
}

func CreateItemBOSSMiZangSubTypee(subType int32) ItemSubType {
	return ItemBOSSMiZangSubType(subType)
}

//幻境、外域boss物品子类型
type ItemBossSubType int32

const (
	// 净灵丹
	ItemBossSubTypeJingLingDan ItemBossSubType = iota
	// 醒神丹
	ItemBossSubTypeXingShenDan
)

func (t ItemBossSubType) SubType() int32 {
	return int32(t)
}

func (t ItemBossSubType) Valid() bool {
	switch t {
	case ItemBossSubTypeJingLingDan,
		ItemBossSubTypeXingShenDan:
		return true
	}
	return false
}

func CreateItemBossSubType(subType int32) ItemSubType {
	return ItemBossSubType(subType)
}

// 结义子类型
type ItemJieYiSubType int32

const (
	ItemJieYiSubTypeDaoJu ItemJieYiSubType = iota
	ItemJieYiSubTypeToken
	ItemJieYiSubTypeTokenLevel
	ItemJieYiSubTypeChangeName
)

func (t ItemJieYiSubType) SubType() int32 {
	return int32(t)
}

func (t ItemJieYiSubType) Valid() bool {
	switch t {
	case ItemJieYiSubTypeDaoJu,
		ItemJieYiSubTypeToken,
		ItemJieYiSubTypeTokenLevel,
		ItemJieYiSubTypeChangeName:
		return true
	}
	return false
}

func CreateItemJieYiSubType(subType int32) ItemSubType {
	return ItemJieYiSubType(subType)
}

// 特戒子类型
type ItemTeRingSubType int32

const (
	ItemTeRingSubTypeRing ItemTeRingSubType = iota
	ItemTeRingSubTypeStrengthen
	ItemTeRingSubTypeAdvance
	ItemTeRingSubTypeJingLing
)

func (t ItemTeRingSubType) SubType() int32 {
	return int32(t)
}

func (t ItemTeRingSubType) Valid() bool {
	switch t {
	case ItemTeRingSubTypeRing,
		ItemTeRingSubTypeStrengthen,
		ItemTeRingSubTypeAdvance,
		ItemTeRingSubTypeJingLing:
		return true
	}
	return false
}

func CreateItemTeRingSubType(subType int32) ItemSubType {
	return ItemTeRingSubType(subType)
}

//血魔子类型
type ItemXueMoSubType int32

const (
	//血魔升阶丹
	ItemXueMoSubTypeAdvanced ItemXueMoSubType = iota
	//血魔培养丹
	ItemXueMoSubTypePeiYang
)

func (t ItemXueMoSubType) SubType() int32 {
	return int32(t)
}

func (t ItemXueMoSubType) Valid() bool {
	switch t {
	case ItemXueMoSubTypeAdvanced,
		ItemXueMoSubTypePeiYang:
		return true
	}
	return false
}

func CreateItemXueMoSubType(subType int32) ItemSubType {
	return ItemXueMoSubType(subType)
}

//灵童子类型
type ItemLingTongSubType int32

const (
	//灵童激活卡
	ItemLingTongSubTypeActivate ItemLingTongSubType = iota
	//灵童升级物品
	ItemLingTongSubTypeUpgrade
	//灵童培养物品
	ItemLingTongSubTypePeiYang
	//灵童碎片
	ItemLingTongSubTypeSuiPian
	//五行灵珠
	ItemLingTongSubTypeWuXingLingZhu
)

func (t ItemLingTongSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongSubType) Valid() bool {
	switch t {
	case ItemLingTongSubTypeActivate,
		ItemLingTongSubTypeUpgrade,
		ItemLingTongSubTypePeiYang,
		ItemLingTongSubTypeSuiPian,
		ItemLingTongSubTypeWuXingLingZhu:
		return true
	}
	return false
}

func CreateItemLingTongSubType(subType int32) ItemSubType {
	return ItemLingTongSubType(subType)
}

//灵童时装子类型
type ItemLingTongFashionSubType int32

const (
	//灵童时装幻化卡
	ItemLingTongFashionSubTypeUnrealCard ItemLingTongFashionSubType = iota
	//灵童时装碎片
	ItemLingTongFashionSubTypeSuiPian
	//灵童时装试用卡
	ItemLingTongFashionSubTypeTrialCard
)

func (t ItemLingTongFashionSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongFashionSubType) Valid() bool {
	switch t {
	case ItemLingTongFashionSubTypeUnrealCard,
		ItemLingTongFashionSubTypeSuiPian,
		ItemLingTongFashionSubTypeTrialCard:
		return true
	}
	return false
}

func CreateItemLingTongFashionSubType(subType int32) ItemSubType {
	return ItemLingTongFashionSubType(subType)
}

//灵童兵魂子类型
type ItemLingTongWeaponSubType int32

const (
	//灵童兵魂幻化卡
	ItemLingTongWeaponSubTypeUnrealCard ItemLingTongWeaponSubType = iota
	//灵童兵魂升阶丹
	ItemLingTongWeaponSubTypeAdvancedDan
	//灵童兵魂幻化丹
	ItemLingTongWeaponSubTypeUnrealDan
	//灵童兵魂培养丹
	ItemLingTongWeaponSubTypePeiYangDan
	//灵童兵魂通灵丹
	ItemLingTongWeaponSubTypeTongLingDan
	//灵童兵魂碎片
	ItemLingTongWeaponSubTypeTongSuiPian
)

func (t ItemLingTongWeaponSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongWeaponSubType) Valid() bool {
	switch t {
	case ItemLingTongWeaponSubTypeUnrealCard,
		ItemLingTongWeaponSubTypeAdvancedDan,
		ItemLingTongWeaponSubTypeUnrealDan,
		ItemLingTongWeaponSubTypePeiYangDan,
		ItemLingTongWeaponSubTypeTongLingDan,
		ItemLingTongWeaponSubTypeTongSuiPian:
		return true
	}
	return false
}

func CreateItemLingTongWeaponSubType(subType int32) ItemSubType {
	return ItemLingTongWeaponSubType(subType)
}

//灵童坐骑子类型
type ItemLingTongMountSubType int32

const (
	//灵童坐骑幻化卡
	ItemLingTongMountSubTypeUnrealCard ItemLingTongMountSubType = iota
	//灵童坐骑升阶丹
	ItemLingTongMountSubTypeAdvancedDan
	//灵童坐骑幻化丹
	ItemLingTongMountSubTypeUnrealDan
	//灵童坐骑培养丹
	ItemLingTongMountSubTypePeiYangDan
	//灵童坐骑通灵
	ItemLingTongMountSubTypeTongLingDan
	//灵童坐骑碎片
	ItemLingTongMountSubTypeSuiPian
)

func (t ItemLingTongMountSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongMountSubType) Valid() bool {
	switch t {
	case ItemLingTongMountSubTypeUnrealCard,
		ItemLingTongMountSubTypeAdvancedDan,
		ItemLingTongMountSubTypeUnrealDan,
		ItemLingTongMountSubTypePeiYangDan,
		ItemLingTongMountSubTypeTongLingDan,
		ItemLingTongMountSubTypeSuiPian:
		return true
	}
	return false
}

func CreateItemLingTongMountSubType(subType int32) ItemSubType {
	return ItemLingTongMountSubType(subType)
}

//灵童战翼子类型
type ItemLingTongWingSubType int32

const (
	//灵童战翼幻化卡
	ItemLingTongWingSubTypeUnrealCard ItemLingTongWingSubType = iota
	//灵童战翼升阶丹
	ItemLingTongWingSubTypeAdvancedDan
	//灵童战翼幻化丹
	ItemLingTongWingSubTypeUnrealDan
	//灵童战翼培养丹
	ItemLingTongWingSubTypePeiYangDan
	//灵童战翼通灵
	ItemLingTongWingSubTypeTongLingDan
	//灵童战翼碎片
	ItemLingTongWingSubTypeSuiPian
)

func (t ItemLingTongWingSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongWingSubType) Valid() bool {
	switch t {
	case ItemLingTongWingSubTypeUnrealCard,
		ItemLingTongWingSubTypeAdvancedDan,
		ItemLingTongWingSubTypeUnrealDan,
		ItemLingTongWingSubTypePeiYangDan,
		ItemLingTongWingSubTypeTongLingDan,
		ItemLingTongWingSubTypeSuiPian:
		return true
	}
	return false
}

func CreateItemLingTongWingSubType(subType int32) ItemSubType {
	return ItemLingTongWingSubType(subType)
}

//灵童身法子类型
type ItemLingTongShenFaSubType int32

const (
	//灵童身法幻化卡
	ItemLingTongShenFaSubTypeUnrealCard ItemLingTongShenFaSubType = iota
	//灵童身法升阶丹
	ItemLingTongShenFaSubTypeAdvancedDan
	//灵童身法幻化丹
	ItemLingTongShenFaSubTypeUnrealDan
	//灵童身法培养丹
	ItemLingTongShenFaSubTypePeiYangDan
	//灵童身法通灵
	ItemLingTongShenFaSubTypeTongLingDan
	//灵童身法碎片
	ItemLingTongShenFaSubTypeSuiPian
)

func (t ItemLingTongShenFaSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongShenFaSubType) Valid() bool {
	switch t {
	case ItemLingTongShenFaSubTypeUnrealCard,
		ItemLingTongShenFaSubTypeAdvancedDan,
		ItemLingTongShenFaSubTypeUnrealDan,
		ItemLingTongShenFaSubTypePeiYangDan,
		ItemLingTongShenFaSubTypeTongLingDan,
		ItemLingTongShenFaSubTypeSuiPian:
		return true
	}
	return false
}

func CreateItemLingTongShenFaSubType(subType int32) ItemSubType {
	return ItemLingTongShenFaSubType(subType)
}

//灵童领域子类型
type ItemLingTongLingYuSubType int32

const (
	//灵童领域幻化卡
	ItemLingTongLingYuSubTypeUnrealCard ItemLingTongLingYuSubType = iota
	//灵童领域幻化丹
	ItemLingTongLingYuSubTypeAdvancedDan
	//灵童领域幻化丹
	ItemLingTongLingYuSubTypeUnrealDan
	//灵童领域培养丹
	ItemLingTongLingYuSubTypePeiYangDan
	//灵童领域通灵丹
	ItemLingTongLingYuSubTypeTongLingDan
	//灵童领域碎片
	ItemLingTongLingYuSubTypeSuiPian
)

func (t ItemLingTongLingYuSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongLingYuSubType) Valid() bool {
	switch t {
	case ItemLingTongLingYuSubTypeUnrealCard,
		ItemLingTongLingYuSubTypeAdvancedDan,
		ItemLingTongLingYuSubTypeUnrealDan,
		ItemLingTongLingYuSubTypePeiYangDan,
		ItemLingTongLingYuSubTypeTongLingDan,
		ItemLingTongLingYuSubTypeSuiPian:
		return true
	}
	return false
}

func CreateItemLingTongLingYuSubType(subType int32) ItemSubType {
	return ItemLingTongLingYuSubType(subType)
}

//灵童法宝子类型
type ItemLingTongFaBaoSubType int32

const (
	//灵童法宝幻化卡
	ItemLingTongFaBaoSubTypeUnrealCard ItemLingTongFaBaoSubType = iota
	//灵童法宝升阶丹
	ItemLingTongFaBaoSubTypeAdvancedDan
	//灵童法宝幻化丹
	ItemLingTongFaBaoSubTypeUnrealDan
	//灵童法宝培养丹
	ItemLingTongFaBaoSubTypePeiYangDan
	//灵童法宝通灵丹
	ItemLingTongFaBaoSubTypeTongLingDan
	//灵童法宝碎片
	ItemLingTongFaBaoSubTypeSuiPian
)

func (t ItemLingTongFaBaoSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongFaBaoSubType) Valid() bool {
	switch t {
	case ItemLingTongFaBaoSubTypeUnrealCard,
		ItemLingTongFaBaoSubTypeAdvancedDan,
		ItemLingTongFaBaoSubTypeUnrealDan,
		ItemLingTongFaBaoSubTypePeiYangDan,
		ItemLingTongFaBaoSubTypeTongLingDan,
		ItemLingTongFaBaoSubTypeSuiPian:
		return true
	}
	return false
}

func CreateItemLingTongFaBaoSubType(subType int32) ItemSubType {
	return ItemLingTongFaBaoSubType(subType)
}

//灵童仙体子类型
type ItemLingTongXianTiSubType int32

const (
	//灵童仙体幻化卡
	ItemLingTongXianTiSubTypeUnrealCard ItemLingTongXianTiSubType = iota
	//灵童仙体升阶丹
	ItemLingTongXianTiSubTypeAdvancedDan
	//灵童仙体幻化丹
	ItemLingTongXianTiSubTypeUnrealDan
	//灵童仙体培养丹
	ItemLingTongXianTiSubTypePeiYangDan
	//灵童仙体通灵丹
	ItemLingTongXianTiSubTypeTongLingDan
	//灵童仙体碎片
	ItemLingTongXianTiSubTypeSuiPian
)

func (t ItemLingTongXianTiSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongXianTiSubType) Valid() bool {
	switch t {
	case ItemLingTongXianTiSubTypeUnrealCard,
		ItemLingTongXianTiSubTypeAdvancedDan,
		ItemLingTongXianTiSubTypeUnrealDan,
		ItemLingTongXianTiSubTypePeiYangDan,
		ItemLingTongXianTiSubTypeTongLingDan,
		ItemLingTongXianTiSubTypeSuiPian:
		return true
	}
	return false
}

func CreateItemLingTongXianTiSubType(subType int32) ItemSubType {
	return ItemLingTongXianTiSubType(subType)
}

//直升丹子类型
type ItemUpLevelDanSubType int32

const (
	ItemUpLevelDanSubTypeUpLevelDan ItemUpLevelDanSubType = iota
	ItemUpLevelDanSubTypeSuperUpLevelDan
)

func (t ItemUpLevelDanSubType) SubType() int32 {
	return int32(t)
}

func (t ItemUpLevelDanSubType) Valid() bool {
	switch t {
	case ItemUpLevelDanSubTypeUpLevelDan,
		ItemUpLevelDanSubTypeSuperUpLevelDan:
		return true
	}
	return false
}

func CreateItemUpLevelDanSubType(subType int32) ItemSubType {
	return ItemUpLevelDanSubType(subType)
}

//灵童装备子类型
type ItemLingTongEquipSubType int32

const (
	//灵环(部位一)
	ItemLingTongEquipSubTypeOne ItemLingTongEquipSubType = iota
	//灵器(部位二)
	ItemLingTongEquipSubTypeTwo
	//灵盾(部位三)
	ItemLingTongEquipSubTypeThree
	//灵饰(部位四)
	ItemLingTongEquipSubTypeFour
)

func (t ItemLingTongEquipSubType) Valid() bool {
	switch t {
	case ItemLingTongEquipSubTypeOne,
		ItemLingTongEquipSubTypeTwo,
		ItemLingTongEquipSubTypeThree,
		ItemLingTongEquipSubTypeFour:
		return true
	}
	return false
}

func (t ItemLingTongEquipSubType) SubType() int32 {
	return int32(t)
}

func CreateItemLingTongEquipSubType(subType int32) ItemSubType {
	return ItemLingTongEquipSubType(subType)
}

//灵童技能书子类型
type ItemLingTongJiNengShuSubType int32

const (
	//灵童兵魂技能书
	ItemLingTongJiNengShuSubTypeWeapon ItemLingTongJiNengShuSubType = iota
	//灵童坐骑技能书
	ItemLingTongJiNengShuSubTypeMount
	//灵童战翼技能书
	ItemLingTongJiNengShuSubTypeWing
	//灵童身法技能书
	ItemLingTongJiNengShuSubTypeShenFa
	//灵童领域技能书
	ItemLingTongJiNengShuSubTypeLingYu
	//灵童法宝技能书
	ItemLingTongJiNengShuSubTypeFaBao
	//灵童仙体技能书
	ItemLingTongJiNengShuSubTypeXianTi
)

func (t ItemLingTongJiNengShuSubType) SubType() int32 {
	return int32(t)
}

func (t ItemLingTongJiNengShuSubType) Valid() bool {
	switch t {
	case ItemLingTongJiNengShuSubTypeWeapon,
		ItemLingTongJiNengShuSubTypeMount,
		ItemLingTongJiNengShuSubTypeWing,
		ItemLingTongJiNengShuSubTypeShenFa,
		ItemLingTongJiNengShuSubTypeLingYu,
		ItemLingTongJiNengShuSubTypeFaBao,
		ItemLingTongJiNengShuSubTypeXianTi:
		return true
	}
	return false
}

func CreateItemLingTongJiNengShuSubType(subType int32) ItemSubType {
	return ItemLingTongJiNengShuSubType(subType)
}

var (
	itemSubTypeFactoryMap = make(map[ItemType]ItemSubTypeFactory)
)

func CreateItemSubType(typ ItemType, subType int32) ItemSubType {
	factory, ok := itemSubTypeFactoryMap[typ]
	if !ok {
		return CreateItemDefaultSubType(subType)
	}
	return factory.CreateItemSubType(subType)
}

func init() {
	itemSubTypeFactoryMap[ItemTypeDefault] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeHp] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeSilver] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeGold] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeSkill] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeBuff] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeFashion] = ItemSubTypeFactoryFunc(CreateItemFashionSubType)
	itemSubTypeFactoryMap[ItemTypeBingHun] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypePetEquipment] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeDan] = ItemSubTypeFactoryFunc(CreateItemDanSubType)
	itemSubTypeFactoryMap[ItemTypeTreasure] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeEquipmentGem] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)

	itemSubTypeFactoryMap[ItemTypeGem] = ItemSubTypeFactoryFunc(CreateItemGemSubType)
	itemSubTypeFactoryMap[ItemTypePet] = ItemSubTypeFactoryFunc(CreateItemMountSubType)
	itemSubTypeFactoryMap[ItemTypeLifeOrigin] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeSkillPoint] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeEquipment] = ItemSubTypeFactoryFunc(CreateItemEquipmentSubType)
	itemSubTypeFactoryMap[ItemTypeEquipmentUpgrade] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)

	itemSubTypeFactoryMap[ItemTypeWing] = ItemSubTypeFactoryFunc(CreateItemWingSubType)
	itemSubTypeFactoryMap[ItemTypeSoul] = ItemSubTypeFactoryFunc(CreateItemSoulSubType)
	itemSubTypeFactoryMap[ItemTypeKingSoul] = ItemSubTypeFactoryFunc(CreateItemKingSoulSubType)
	itemSubTypeFactoryMap[ItemTypeDun] = ItemSubTypeFactoryFunc(CreateItemBodyShieldSubType)
	itemSubTypeFactoryMap[ItemTypeTitle] = ItemSubTypeFactoryFunc(CreateItemTitleSubType)

	itemSubTypeFactoryMap[ItemTypeAutoUseRes] = ItemSubTypeFactoryFunc(CreateItemAutoUseResSubType)
	itemSubTypeFactoryMap[ItemTypeXianHua] = ItemSubTypeFactoryFunc(CreateItemXueHuaSubType)
	itemSubTypeFactoryMap[ItemTypeTuMoLing] = ItemSubTypeFactoryFunc(CreateItemTuMoLingSubType)

	itemSubTypeFactoryMap[ItemTypeJueXue] = ItemSubTypeFactoryFunc(CreateItemJueXueSubType)
	itemSubTypeFactoryMap[ItemTypeXinFa] = ItemSubTypeFactoryFunc(CreateItemXinFaSubType)
	itemSubTypeFactoryMap[ItemTypeEctypal] = ItemSubTypeFactoryFunc(CreateItemEctypalSubType)
	itemSubTypeFactoryMap[ItemTypeAlliance] = ItemSubTypeFactoryFunc(CreateItemAllianceSubType)

	itemSubTypeFactoryMap[ItemTypeGiftBag] = ItemSubTypeFactoryFunc(CreateItemGiftBagSubType)
	itemSubTypeFactoryMap[ItemTypeShenLong] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)

	itemSubTypeFactoryMap[ItemTypeCandy] = ItemSubTypeFactoryFunc(CreateItemCandySubType)
	itemSubTypeFactoryMap[ItemTypeWedRing] = ItemSubTypeFactoryFunc(CreateItemWedRingSubType)
	itemSubTypeFactoryMap[ItemTypeRingAdvaced] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)

	itemSubTypeFactoryMap[ItemTypeShenFa] = ItemSubTypeFactoryFunc(CreateItemShenFaSubType)
	itemSubTypeFactoryMap[ItemTypeLingyu] = ItemSubTypeFactoryFunc(CreateItemLingyuSubType)

	itemSubTypeFactoryMap[ItemTypeGoldEquip] = ItemSubTypeFactoryFunc(CreateItemGoldEquipSubType)
	itemSubTypeFactoryMap[ItemTypeTreeAdvanced] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeKun] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypePKSkillUpgrade] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeWedSuccRing] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeWaKuang] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)

	itemSubTypeFactoryMap[ItemTypeAnQi] = ItemSubTypeFactoryFunc(CreateItemAnQiSubType)
	itemSubTypeFactoryMap[ItemTypeBlessDan] = ItemSubTypeFactoryFunc(CreateItemBlessDanSubType)
	itemSubTypeFactoryMap[ItemTypeAdvancedTicket] = ItemSubTypeFactoryFunc(CreateItemAdvancedTicketSubType)
	itemSubTypeFactoryMap[ItemTypeUpLevelDan] = ItemSubTypeFactoryFunc(CreateItemUpLevelDanSubType)
	itemSubTypeFactoryMap[ItemTypeBossCallTicket] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypePkValue] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeWingActivateCard] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeGoldEquipStrengthen] = ItemSubTypeFactoryFunc(CreateItemGoldEquipStrengthenSubType)
	itemSubTypeFactoryMap[ItemTypeGemAvoidBomb] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeLuckyRate] = ItemSubTypeFactoryFunc(CreateItemLuckySubType)
	itemSubTypeFactoryMap[ItemTypeSexCard] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeRenameCard] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeResouceCard] = ItemSubTypeFactoryFunc(CreateItemResourceCardSubType)
	itemSubTypeFactoryMap[ItemTypeTower] = ItemSubTypeFactoryFunc(CreateItemTowerSubType)
	itemSubTypeFactoryMap[ItemTypeMyBoss] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeMountEquip] = ItemSubTypeFactoryFunc(CreateMountEquipmentSubType)
	itemSubTypeFactoryMap[ItemTypeWingStone] = ItemSubTypeFactoryFunc(CreateWingStoneSubType)
	itemSubTypeFactoryMap[ItemTypeAnqiJiguan] = ItemSubTypeFactoryFunc(CreateAnqiJiguanSubType)
	itemSubTypeFactoryMap[ItemTypeJiNengShu] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeQiangHuaShi] = ItemSubTypeFactoryFunc(CreateStrengthenStoneSubType)
	itemSubTypeFactoryMap[ItemTypeZhuiZongFu] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeTianShu] = ItemSubTypeFactoryFunc(CreateItemTianShuSubType)

	itemSubTypeFactoryMap[ItemTypeFaBaoSuit] = ItemSubTypeFactoryFunc(CreateItemFaBaoSuitSubType)
	itemSubTypeFactoryMap[ItemTypeFaBao] = ItemSubTypeFactoryFunc(CreateItemFaBaoSubType)
	itemSubTypeFactoryMap[ItemTypeXianTi] = ItemSubTypeFactoryFunc(CreateItemXianTiSubType)
	itemSubTypeFactoryMap[ItemTypeXianTiLingYu] = ItemSubTypeFactoryFunc(CreateXiantiLingyuSubType)

	itemSubTypeFactoryMap[ItemTypeXueDun] = ItemSubTypeFactoryFunc(CreateItemXueDunSubType)
	itemSubTypeFactoryMap[ItemTypeFuChi] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeDianXing] = ItemSubTypeFactoryFunc(CreateItemDianXingSubType)

	itemSubTypeFactoryMap[ItemTypeTianMo] = ItemSubTypeFactoryFunc(CreateItemTianMoSubType)
	itemSubTypeFactoryMap[ItemTypeXueMo] = ItemSubTypeFactoryFunc(CreateItemXueMoSubType)
	itemSubTypeFactoryMap[ItemTypeLingyuEquip] = ItemSubTypeFactoryFunc(CreateLingyuEquipSubType)
	itemSubTypeFactoryMap[ItemTypeShenfaEquip] = ItemSubTypeFactoryFunc(CreateShenfaEquipSubType)
	itemSubTypeFactoryMap[ItemTypeShiHunFanEquip] = ItemSubTypeFactoryFunc(CreateShiHunFanEquipSubType)
	itemSubTypeFactoryMap[ItemTypeTianMoTiEquip] = ItemSubTypeFactoryFunc(CreateTianMoTiEquipSubType)

	itemSubTypeFactoryMap[ItemTypeLingTongReNameCard] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeLingTong] = ItemSubTypeFactoryFunc(CreateItemLingTongSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongFashion] = ItemSubTypeFactoryFunc(CreateItemLingTongFashionSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongWeapon] = ItemSubTypeFactoryFunc(CreateItemLingTongWeaponSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongMount] = ItemSubTypeFactoryFunc(CreateItemLingTongMountSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongWing] = ItemSubTypeFactoryFunc(CreateItemLingTongWingSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongShenFa] = ItemSubTypeFactoryFunc(CreateItemLingTongShenFaSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongLingYu] = ItemSubTypeFactoryFunc(CreateItemLingTongLingYuSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongFaBao] = ItemSubTypeFactoryFunc(CreateItemLingTongFaBaoSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongXianTi] = ItemSubTypeFactoryFunc(CreateItemLingTongXianTiSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongMountEquip] = ItemSubTypeFactoryFunc(CreateLingTongMountEquipSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongWingEquip] = ItemSubTypeFactoryFunc(CreateLingTongWingEquipSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongShenFaEquip] = ItemSubTypeFactoryFunc(CreateLingTongShenFaEquipSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongLingYuEquip] = ItemSubTypeFactoryFunc(CreateLingTongLingYuEquipSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongFaBaoEquip] = ItemSubTypeFactoryFunc(CreateLingTongFaBaoEquipSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongXianTiEquip] = ItemSubTypeFactoryFunc(CreateLingTongXianTiEquipSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongWeaponEquip] = ItemSubTypeFactoryFunc(CreateLingTongWeaponEquipSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongJiNengShu] = ItemSubTypeFactoryFunc(CreateItemLingTongJiNengShuSubType)
	itemSubTypeFactoryMap[ItemTypeLingTongEquip] = ItemSubTypeFactoryFunc(CreateItemLingTongEquipSubType)

	itemSubTypeFactoryMap[ItemTypeFeiSheng] = ItemSubTypeFactoryFunc(CreateItemFeiShengSubType)
	itemSubTypeFactoryMap[ItemTypeHongBao] = ItemSubTypeFactoryFunc(CreateItemHongBaoSubType)
	itemSubTypeFactoryMap[ItemTypeHuaLingDan] = ItemSubTypeFactoryFunc(CreateItemHuaLingDanSubType)
	itemSubTypeFactoryMap[ItemTypeTianFuDan] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeEquipBaoKuTicket] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeShengHenEquipQingLong] = ItemSubTypeFactoryFunc(CreateShengHenEquipQingLongSubType)
	itemSubTypeFactoryMap[ItemTypeShengHenEquipBaiHu] = ItemSubTypeFactoryFunc(CreateShengHenEquipBaiHuSubType)
	itemSubTypeFactoryMap[ItemTypeShengHenEquipZhuQue] = ItemSubTypeFactoryFunc(CreateShengHenEquipZhuQueSubType)
	itemSubTypeFactoryMap[ItemTypeShengHenEquipXuanWu] = ItemSubTypeFactoryFunc(CreateShengHenEquipXuanWuSubType)
	itemSubTypeFactoryMap[ItemTypeSheng] = ItemSubTypeFactoryFunc(CreateShengHenSubType)
	itemSubTypeFactoryMap[ItemTypeMingGe] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeShenQi] = ItemSubTypeFactoryFunc(CreateItemShenQiSubType)
	itemSubTypeFactoryMap[ItemTypeYingLingPu] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeZhenFa] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeTuLongEquipItem] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeTuLongEquip] = ItemSubTypeFactoryFunc(CreateTuLongEquipSubType)
	itemSubTypeFactoryMap[ItemTypeHunt] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeBiaoBai] = ItemSubTypeFactoryFunc(CreateItemBiaoBaiSubType)
	itemSubTypeFactoryMap[ItemTypeBaoBao] = ItemSubTypeFactoryFunc(CreateItemBaoBaoSubType)
	itemSubTypeFactoryMap[ItemTypeBabyToy] = ItemSubTypeFactoryFunc(CreateItemBabyToySubType)
	itemSubTypeFactoryMap[ItemTypeDingQing] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeBaoBaoCard] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeAdditionsysTongLingDan] = ItemSubTypeFactoryFunc(CreateItemAdditionsysTongLingDanSubType)
	itemSubTypeFactoryMap[ItemTypeAdditionsysJueXingDan] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeYuanBaoKa] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeExpendBagSlotCard] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeFuShi] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeItemSkill] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeBossItem] = ItemSubTypeFactoryFunc(CreateItemBossSubType)
	itemSubTypeFactoryMap[ItemTypeJieYiItem] = ItemSubTypeFactoryFunc(CreateItemJieYiSubType)
	itemSubTypeFactoryMap[ItemTypeXianJinYuanBaoKa] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeChuangShiZhiZhan] = ItemSubTypeFactoryFunc(CreateItemChuangShiZhiZhanSubType)
	itemSubTypeFactoryMap[ItemTypeBOSSMiZang] = ItemSubTypeFactoryFunc(CreateItemBOSSMiZangSubTypee)
	itemSubTypeFactoryMap[ItemTypeVigorousPill] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
	itemSubTypeFactoryMap[ItemTypeWushuangWeapon] = ItemSubTypeFactoryFunc(CreateItemWushuangWeaponSubType)
	itemSubTypeFactoryMap[ItemTypeWushuangWeaponEssence] = ItemSubTypeFactoryFunc(CreateItemWushuangWeaponEssenceSubType)
	itemSubTypeFactoryMap[ItemTypeGodCastingEquip] = ItemSubTypeFactoryFunc(CreateItemGodCastingEquipSubType)
	itemSubTypeFactoryMap[ItemTypeManHanQuanXi] = ItemSubTypeFactoryFunc(CreateItemManHanQuanXiSubType)
	itemSubTypeFactoryMap[ItemTypeTeRing] = ItemSubTypeFactoryFunc(CreateItemTeRingSubType)
	itemSubTypeFactoryMap[ItemTypeShangGuZhiLing] = ItemSubTypeFactoryFunc(CreateItemDefaultSubType)
}

type ItemUseFlag int32

const (
	//可以使用
	ItemUseFlagUse = 1 << iota
	//使用后删除
	ItemUseFlagDelete
	//日志记录
	ItemUseFlagLog
	//人民币物品
	ItemUseFlagRMB
	//失效后删除
	ItemUseFlagExpiredDelete
)

func (iuf ItemUseFlag) Valid() bool {
	switch iuf {
	case ItemUseFlagUse,
		ItemUseFlagDelete,
		ItemUseFlagLog,
		ItemUseFlagRMB,
		ItemUseFlagExpiredDelete:
		return true
	}
	return false
}

//过期时间类型(item表)
type ItemLimitTimeType int32

const (
	//无过期时间
	ItemLimitTimeTypeNone ItemLimitTimeType = iota
	//多久之后过期
	ItemLimitTimeTypeExpired
	//指定时间过期
	ItemLimitTimeTypeExpiredSpecialTime
	//当天几点后过期
	ItemLimitTimeTypeTodayExpired
)

func (t ItemLimitTimeType) Valid() bool {
	switch t {
	case ItemLimitTimeTypeNone,
		ItemLimitTimeTypeExpired,
		ItemLimitTimeTypeTodayExpired:
		return true
	}
	return false
}

var (
	toNewExpireTypeMap = map[ItemLimitTimeType]inventorytypes.NewItemLimitTimeType{
		ItemLimitTimeTypeNone:         inventorytypes.NewItemLimitTimeTypeNone,
		ItemLimitTimeTypeExpired:      inventorytypes.NewItemLimitTimeTypeExpiredAfterTime,
		ItemLimitTimeTypeTodayExpired: inventorytypes.NewItemLimitTimeTypeExpiredToday,
	}
)

func (t ItemLimitTimeType) ConvertToNewItemLimitTimeType() inventorytypes.NewItemLimitTimeType {
	newType, ok := toNewExpireTypeMap[t]
	if !ok {
		return inventorytypes.NewItemLimitTimeTypeNone
	}
	return newType
}

//物品品质
type ItemQualityType int32

const (
	ItemQualityTypeWhite ItemQualityType = iota
	ItemQualityTypeGreen
	ItemQualityTypeBlue
	ItemQualityTypePurple
	ItemQualityTypeOrange
	ItemQualityTypeRed
)

var qalityColorMap = map[ItemQualityType]string{
	ItemQualityTypeWhite:  "#bdd3e0",
	ItemQualityTypeGreen:  "#4fda6f",
	ItemQualityTypeBlue:   "#3ab7ff",
	ItemQualityTypePurple: "#dc4bf6",
	ItemQualityTypeOrange: "#f6ac4b",
	ItemQualityTypeRed:    "#f64b4b",
}

func (t ItemQualityType) GetColor() string {
	return qalityColorMap[t]
}

func (t ItemQualityType) Valid() bool {
	switch t {
	case ItemQualityTypeWhite,
		ItemQualityTypeGreen,
		ItemQualityTypeBlue,
		ItemQualityTypePurple,
		ItemQualityTypeOrange,
		ItemQualityTypeRed:
		return true
	}
	return false
}

// // 颜色模块
// type ColorModuleType int32

// const (
// 	ColorModuleTypeChat   ColorModuleType = iota //聊天
// 	ColorModuleTypeNotice                        //公告
// )

//绑定属性
type ItemBindType int32

const (
	ItemBindTypeUnBind ItemBindType = iota //不绑定
	ItemBindTypeBind                       //绑定
)

func (t ItemBindType) Valid() bool {
	switch t {
	case ItemBindTypeUnBind,
		ItemBindTypeBind:
		return true
	}
	return false
}

//坐骑装备子类型
type MountEquipmentSubType int32

const (
	//马鞭
	MountEquipmentSubTypeRomal MountEquipmentSubType = iota
	//马鞍
	MountEquipmentSubTypeSaddle
	//马镫
	MountEquipmentSubTypeStirrup
	//马蹄
	MountEquipmentSubTypeHoof
)

func (iest MountEquipmentSubType) SubType() int32 {
	return int32(iest)
}

func (iest MountEquipmentSubType) Valid() bool {
	switch iest {
	case MountEquipmentSubTypeRomal,
		MountEquipmentSubTypeSaddle,
		MountEquipmentSubTypeStirrup,
		MountEquipmentSubTypeHoof:
		return true
	}
	return false
}

func CreateMountEquipmentSubType(subType int32) ItemSubType {
	return MountEquipmentSubType(subType)
}

//战翼符石子类型
type WingStoneSubType int32

const (
	//天
	WingStoneSubTypeTian WingStoneSubType = iota
	//地
	WingStoneSubTypeDi
	//玄
	WingStoneSubTypeXuan
	//黄
	WingStoneSubTypeHuang
)

func (iest WingStoneSubType) SubType() int32 {
	return int32(iest)
}

func (iest WingStoneSubType) Valid() bool {
	switch iest {
	case WingStoneSubTypeTian,
		WingStoneSubTypeDi,
		WingStoneSubTypeXuan,
		WingStoneSubTypeHuang:
		return true
	}
	return false
}

func CreateWingStoneSubType(subType int32) ItemSubType {
	return WingStoneSubType(subType)
}

//暗器机关子类型
type AnqiJiguanSubType int32

const (
	//机括
	AnqiJiguanSubTypeJiKuo AnqiJiguanSubType = iota
	//暗格
	AnqiJiguanSubTypeAnGe
	//毒药
	AnqiJiguanSubTypeDuYao
	//钉刺
	AnqiJiguanSubTypeDingCi
)

func (iest AnqiJiguanSubType) SubType() int32 {
	return int32(iest)
}

func (iest AnqiJiguanSubType) Valid() bool {
	switch iest {
	case AnqiJiguanSubTypeJiKuo,
		AnqiJiguanSubTypeAnGe,
		AnqiJiguanSubTypeDuYao,
		AnqiJiguanSubTypeDingCi:
		return true
	}
	return false
}

func CreateAnqiJiguanSubType(subType int32) ItemSubType {
	return AnqiJiguanSubType(subType)
}

//强化石子类型
type StrengthenStoneSubType int32

const (
	//坐骑强化石
	StrengthenStoneSubTypeMount StrengthenStoneSubType = iota
	//翅膀强化石
	StrengthenStoneSubTypeWing
	//暗器强化石
	StrengthenStoneSubTypeAnqi
)

func (iest StrengthenStoneSubType) SubType() int32 {
	return int32(iest)
}

func (iest StrengthenStoneSubType) Valid() bool {
	switch iest {
	case StrengthenStoneSubTypeMount,
		StrengthenStoneSubTypeWing,
		StrengthenStoneSubTypeAnqi:
		return true
	}
	return false
}

func CreateStrengthenStoneSubType(subType int32) ItemSubType {
	return StrengthenStoneSubType(subType)
}

//仙体子类型
type ItemXianTiSubType int32

const (
	//仙体升阶丹
	ItemXianTiSubTypeAdvanced ItemXianTiSubType = iota
	//仙体幻化丹
	ItemXianTiSubTypeUnreal
	//仙体培养丹
	ItemXianTiSubTypeCul
	//仙体幻化卡
	ItemXianTiSubTypeUnrealCard
	//仙体碎片
	ItemXianTiSubTypeChip
)

func (iest ItemXianTiSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemXianTiSubType) Valid() bool {
	switch iest {
	case ItemXianTiSubTypeAdvanced,
		ItemXianTiSubTypeUnreal,
		ItemXianTiSubTypeCul,
		ItemXianTiSubTypeUnrealCard,
		ItemXianTiSubTypeChip:
		return true
	}
	return false
}

func CreateItemXianTiSubType(subType int32) ItemSubType {
	return ItemXianTiSubType(subType)
}

//仙体灵玉子类型
type XiantiLingyuSubType int32

const (
	//天罡
	XiantiLingyuSubTypeTianGang XiantiLingyuSubType = iota
	//地煞
	XiantiLingyuSubTypeDiSha
	//浮屠
	XiantiLingyuSubTypeFuTu
	//燃灭
	XiantiLingyuSubTypeRanMie
)

func (iest XiantiLingyuSubType) SubType() int32 {
	return int32(iest)
}

func (iest XiantiLingyuSubType) Valid() bool {
	switch iest {
	case XiantiLingyuSubTypeTianGang,
		XiantiLingyuSubTypeDiSha,
		XiantiLingyuSubTypeFuTu,
		XiantiLingyuSubTypeRanMie:
		return true
	}
	return false
}

func CreateXiantiLingyuSubType(subType int32) ItemSubType {
	return XiantiLingyuSubType(subType)
}

//点星系统子类型
type ItemDianXingSubType int32

const (
	//星云符
	ItemDianXingSubTypeXingYunFu ItemDianXingSubType = iota
	//解封石
	ItemDianXingSubTypeJieFengShi
)

func (iest ItemDianXingSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemDianXingSubType) Valid() bool {
	switch iest {
	case ItemDianXingSubTypeXingYunFu,
		ItemDianXingSubTypeJieFengShi:
		return true
	}
	return false
}

func CreateItemDianXingSubType(subType int32) ItemSubType {
	return ItemDianXingSubType(subType)
}

//领域装备子类型
type LingyuEquipSubType int32

const (
	//装备1
	LingyuEquipSubTypeOne LingyuEquipSubType = iota
	//装备2
	LingyuEquipSubTypeTwo
	//装备3
	LingyuEquipSubTypeThree
	//装备4
	LingyuEquipSubTypeFour
)

func (iest LingyuEquipSubType) SubType() int32 {
	return int32(iest)
}

func (iest LingyuEquipSubType) Valid() bool {
	switch iest {
	case LingyuEquipSubTypeOne,
		LingyuEquipSubTypeTwo,
		LingyuEquipSubTypeThree,
		LingyuEquipSubTypeFour:
		return true
	}
	return false
}

func CreateLingyuEquipSubType(subType int32) ItemSubType {
	return LingyuEquipSubType(subType)
}

//身法装备子类型
type ShenfaEquipSubType int32

const (
	//装备1
	ShenfaEquipSubTypeOne ShenfaEquipSubType = iota
	//装备2
	ShenfaEquipSubTypeTwo
	//装备3
	ShenfaEquipSubTypeThree
	//装备4
	ShenfaEquipSubTypeFour
)

func (iest ShenfaEquipSubType) SubType() int32 {
	return int32(iest)
}

func (iest ShenfaEquipSubType) Valid() bool {
	switch iest {
	case ShenfaEquipSubTypeOne,
		ShenfaEquipSubTypeTwo,
		ShenfaEquipSubTypeThree,
		ShenfaEquipSubTypeFour:
		return true
	}
	return false
}

func CreateShenfaEquipSubType(subType int32) ItemSubType {
	return ShenfaEquipSubType(subType)
}

//噬魂幡装备子类型
type ShiHunFanEquipSubType int32

const (
	//装备1
	ShiHunFanEquipSubTypeOne ShiHunFanEquipSubType = iota
	//装备2
	ShiHunFanEquipSubTypeTwo
	//装备3
	ShiHunFanEquipSubTypeThree
	//装备4
	ShiHunFanEquipSubTypeFour
)

func (iest ShiHunFanEquipSubType) SubType() int32 {
	return int32(iest)
}

func (iest ShiHunFanEquipSubType) Valid() bool {
	switch iest {
	case ShiHunFanEquipSubTypeOne,
		ShiHunFanEquipSubTypeTwo,
		ShiHunFanEquipSubTypeThree,
		ShiHunFanEquipSubTypeFour:
		return true
	}
	return false
}

func CreateShiHunFanEquipSubType(subType int32) ItemSubType {
	return ShiHunFanEquipSubType(subType)
}

//天魔体装备子类型
type TianMoTiEquipSubType int32

const (
	//装备1
	TianMoTiEquipSubTypeOne TianMoTiEquipSubType = iota
	//装备2
	TianMoTiEquipSubTypeTwo
	//装备3
	TianMoTiEquipSubTypeThree
	//装备4
	TianMoTiEquipSubTypeFour
)

func (t TianMoTiEquipSubType) SubType() int32 {
	return int32(t)
}

func (t TianMoTiEquipSubType) Valid() bool {
	switch t {
	case TianMoTiEquipSubTypeOne,
		TianMoTiEquipSubTypeTwo,
		TianMoTiEquipSubTypeThree,
		TianMoTiEquipSubTypeFour:
		return true
	}
	return false
}

func CreateTianMoTiEquipSubType(subType int32) ItemSubType {
	return TianMoTiEquipSubType(subType)
}

//灵童坐骑装备子类型
type LingTongMountEquipSubType int32

const (
	//装备1
	LingTongMountEquipSubTypeOne LingTongMountEquipSubType = iota
	//装备2
	LingTongMountEquipSubTypeTwo
	//装备3
	LingTongMountEquipSubTypeThree
	//装备4
	LingTongMountEquipSubTypeFour
)

func (t LingTongMountEquipSubType) SubType() int32 {
	return int32(t)
}

func (t LingTongMountEquipSubType) Valid() bool {
	switch t {
	case LingTongMountEquipSubTypeOne,
		LingTongMountEquipSubTypeTwo,
		LingTongMountEquipSubTypeThree,
		LingTongMountEquipSubTypeFour:
		return true
	}
	return false
}

func CreateLingTongMountEquipSubType(subType int32) ItemSubType {
	return LingTongMountEquipSubType(subType)
}

//灵童战翼装备子类型
type LingTongWingEquipSubType int32

const (
	//装备1
	LingTongWingEquipSubTypeOne LingTongWingEquipSubType = iota
	//装备2
	LingTongWingEquipSubTypeTwo
	//装备3
	LingTongWingEquipSubTypeThree
	//装备4
	LingTongWingEquipSubTypeFour
)

func (t LingTongWingEquipSubType) SubType() int32 {
	return int32(t)
}

func (t LingTongWingEquipSubType) Valid() bool {
	switch t {
	case LingTongWingEquipSubTypeOne,
		LingTongWingEquipSubTypeTwo,
		LingTongWingEquipSubTypeThree,
		LingTongWingEquipSubTypeFour:
		return true
	}
	return false
}

func CreateLingTongWingEquipSubType(subType int32) ItemSubType {
	return LingTongWingEquipSubType(subType)
}

//灵童领域装备子类型
type LingTongLingYuEquipSubType int32

const (
	//装备1
	LingTongLingYuEquipSubTypeOne LingTongLingYuEquipSubType = iota
	//装备2
	LingTongLingYuEquipSubTypeTwo
	//装备3
	LingTongLingYuEquipSubTypeThree
	//装备4
	LingTongLingYuEquipSubTypeFour
)

func (t LingTongLingYuEquipSubType) SubType() int32 {
	return int32(t)
}

func (t LingTongLingYuEquipSubType) Valid() bool {
	switch t {
	case LingTongLingYuEquipSubTypeOne,
		LingTongLingYuEquipSubTypeTwo,
		LingTongLingYuEquipSubTypeThree,
		LingTongLingYuEquipSubTypeFour:
		return true
	}
	return false
}

func CreateLingTongLingYuEquipSubType(subType int32) ItemSubType {
	return LingTongLingYuEquipSubType(subType)
}

//灵童身法装备子类型
type LingTongShenFaEquipSubType int32

const (
	//装备1
	LingTongShenFaEquipSubTypeOne LingTongShenFaEquipSubType = iota
	//装备2
	LingTongShenFaEquipSubTypeTwo
	//装备3
	LingTongShenFaEquipSubTypeThree
	//装备4
	LingTongShenFaEquipSubTypeFour
)

func (t LingTongShenFaEquipSubType) SubType() int32 {
	return int32(t)
}

func (t LingTongShenFaEquipSubType) Valid() bool {
	switch t {
	case LingTongShenFaEquipSubTypeOne,
		LingTongShenFaEquipSubTypeTwo,
		LingTongShenFaEquipSubTypeThree,
		LingTongShenFaEquipSubTypeFour:
		return true
	}
	return false
}

func CreateLingTongShenFaEquipSubType(subType int32) ItemSubType {
	return LingTongShenFaEquipSubType(subType)
}

//灵童法宝装备子类型
type LingTongFaBaoEquipSubType int32

const (
	//装备1
	LingTongFaBaoEquipSubTypeOne LingTongFaBaoEquipSubType = iota
	//装备2
	LingTongFaBaoEquipSubTypeTwo
	//装备3
	LingTongFaBaoEquipSubTypeThree
	//装备4
	LingTongFaBaoEquipSubTypeFour
)

func (t LingTongFaBaoEquipSubType) SubType() int32 {
	return int32(t)
}

func (t LingTongFaBaoEquipSubType) Valid() bool {
	switch t {
	case LingTongFaBaoEquipSubTypeOne,
		LingTongFaBaoEquipSubTypeTwo,
		LingTongFaBaoEquipSubTypeThree,
		LingTongFaBaoEquipSubTypeFour:
		return true
	}
	return false
}

func CreateLingTongFaBaoEquipSubType(subType int32) ItemSubType {
	return LingTongFaBaoEquipSubType(subType)
}

//灵童仙体装备子类型
type LingTongXianTiEquipSubType int32

const (
	//装备1
	LingTongXianTiEquipSubTypeOne LingTongXianTiEquipSubType = iota
	//装备2
	LingTongXianTiEquipSubTypeTwo
	//装备3
	LingTongXianTiEquipSubTypeThree
	//装备4
	LingTongXianTiEquipSubTypeFour
)

func (t LingTongXianTiEquipSubType) SubType() int32 {
	return int32(t)
}

func (t LingTongXianTiEquipSubType) Valid() bool {
	switch t {
	case LingTongXianTiEquipSubTypeOne,
		LingTongXianTiEquipSubTypeTwo,
		LingTongXianTiEquipSubTypeThree,
		LingTongXianTiEquipSubTypeFour:
		return true
	}
	return false
}

func CreateLingTongXianTiEquipSubType(subType int32) ItemSubType {
	return LingTongXianTiEquipSubType(subType)
}

//灵童兵魂装备子类型
type LingTongWeaponEquipSubType int32

const (
	//装备1
	LingTongWeaponEquipSubTypeOne LingTongWeaponEquipSubType = iota
	//装备2
	LingTongWeaponEquipSubTypeTwo
	//装备3
	LingTongWeaponEquipSubTypeThree
	//装备4
	LingTongWeaponEquipSubTypeFour
)

func (t LingTongWeaponEquipSubType) SubType() int32 {
	return int32(t)
}

func (t LingTongWeaponEquipSubType) Valid() bool {
	switch t {
	case LingTongWeaponEquipSubTypeOne,
		LingTongWeaponEquipSubTypeTwo,
		LingTongWeaponEquipSubTypeThree,
		LingTongWeaponEquipSubTypeFour:
		return true
	}
	return false
}

func CreateLingTongWeaponEquipSubType(subType int32) ItemSubType {
	return LingTongWeaponEquipSubType(subType)
}

// 飞升子类型
type ItemFeiShengSubType int32

const (
	//飞升丹
	ItemFeiShengSubTypeJinDan ItemFeiShengSubType = iota
	//功德丹
	ItemFeiShengSubTypeGongDeDan
)

func (iest ItemFeiShengSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemFeiShengSubType) Valid() bool {
	switch iest {
	case ItemFeiShengSubTypeJinDan,
		ItemFeiShengSubTypeGongDeDan:
		return true
	}
	return false
}

func CreateItemFeiShengSubType(subType int32) ItemSubType {
	return ItemFeiShengSubType(subType)
}

// 红包子类型
type ItemHongBaoSubType int32

const (
	//银两红包
	ItemHongBaoSubTypeSilver ItemHongBaoSubType = iota
	//绑元红包
	ItemHongBaoSubTypeGold
	//珍稀红包
	ItemHongBaoSubTypeZhenXi
)

func (iest ItemHongBaoSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemHongBaoSubType) Valid() bool {
	switch iest {
	case ItemHongBaoSubTypeSilver,
		ItemHongBaoSubTypeGold,
		ItemHongBaoSubTypeZhenXi:
		return true
	}
	return false
}

func CreateItemHongBaoSubType(subType int32) ItemSubType {
	return ItemHongBaoSubType(subType)
}

// 化灵丹子类型
type ItemHuaLingDanSubType int32

const (
	//坐骑化灵丹
	ItemHuaLingDanSubTypeMount ItemHuaLingDanSubType = iota
	//战翼化灵丹
	ItemHuaLingDanSubTypeWing
	//暗器化灵丹
	ItemHuaLingDanSubTypeAnQi
	//法宝化灵丹
	ItemHuaLingDanSubTypeFaBao
	//仙体化灵丹
	ItemHuaLingDanSubTypeXianTi
	//领域化灵丹
	ItemHuaLingDanSubTypeLingYu
	//身法化灵丹
	ItemHuaLingDanSubTypeShenFa
	//噬魂幡化灵丹
	ItemHuaLingDanSubTypeShiHunFan
	//天魔体化灵丹
	ItemHuaLingDanSubTypeTianMoTi
	//灵兵化灵丹
	ItemHuaLingDanSubTypeLingBing
	//灵骑化灵丹
	ItemHuaLingDanSubTypeLingQi
	//灵翼化灵丹
	ItemHuaLingDanSubTypeLingWing
	//灵身化灵丹
	ItemHuaLingDanSubTypeLingShen
	//灵域化灵丹
	ItemHuaLingDanSubTypeLingArea
	//灵宝化灵丹
	ItemHuaLingDanSubTypeLingBao
	//灵体化灵丹
	ItemHuaLingDanSubTypeLingTi
)

func (iest ItemHuaLingDanSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemHuaLingDanSubType) Valid() bool {
	switch iest {
	case ItemHuaLingDanSubTypeMount,
		ItemHuaLingDanSubTypeWing,
		ItemHuaLingDanSubTypeAnQi,
		ItemHuaLingDanSubTypeFaBao,
		ItemHuaLingDanSubTypeXianTi,
		ItemHuaLingDanSubTypeLingYu,
		ItemHuaLingDanSubTypeShenFa,
		ItemHuaLingDanSubTypeShiHunFan,
		ItemHuaLingDanSubTypeTianMoTi,
		ItemHuaLingDanSubTypeLingBing,
		ItemHuaLingDanSubTypeLingQi,
		ItemHuaLingDanSubTypeLingWing,
		ItemHuaLingDanSubTypeLingShen,
		ItemHuaLingDanSubTypeLingArea,
		ItemHuaLingDanSubTypeLingBao,
		ItemHuaLingDanSubTypeLingTi:
		return true
	}
	return false
}

func CreateItemHuaLingDanSubType(subType int32) ItemSubType {
	return ItemHuaLingDanSubType(subType)
}

// 附加系统通灵丹子类型
type ItemAdditionsysTongLingDanSubType int32

const (
	//坐骑化灵丹
	ItemAdditionsysTongLingDanSubTypeMount ItemAdditionsysTongLingDanSubType = iota
	//战翼化灵丹
	ItemAdditionsysTongLingDanSubTypeWing
	//暗器化灵丹
	ItemAdditionsysTongLingDanSubTypeAnQi
	//法宝化灵丹
	ItemAdditionsysTongLingDanSubTypeFaBao
	//仙体化灵丹
	ItemAdditionsysTongLingDanSubTypeXianTi
	//领域化灵丹
	ItemAdditionsysTongLingDanSubTypeLingYu
	//身法化灵丹
	ItemAdditionsysTongLingDanSubTypeShenFa
	//噬魂幡化灵丹
	ItemAdditionsysTongLingDanSubTypeShiHunFan
	//天魔体化灵丹
	ItemAdditionsysTongLingDanSubTypeTianMoTi
	//灵兵化灵丹
	ItemAdditionsysTongLingDanSubTypeLingBing
	//灵骑化灵丹
	ItemAdditionsysTongLingDanSubTypeLingQi
	//灵翼化灵丹
	ItemAdditionsysTongLingDanSubTypeLingWing
	//灵身化灵丹
	ItemAdditionsysTongLingDanSubTypeLingShen
	//灵域化灵丹
	ItemAdditionsysTongLingDanSubTypeLingArea
	//灵宝化灵丹
	ItemAdditionsysTongLingDanSubTypeLingBao
	//灵体化灵丹
	ItemAdditionsysTongLingDanSubTypeLingTi
)

func (iest ItemAdditionsysTongLingDanSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemAdditionsysTongLingDanSubType) Valid() bool {
	switch iest {
	case ItemAdditionsysTongLingDanSubTypeMount,
		ItemAdditionsysTongLingDanSubTypeWing,
		ItemAdditionsysTongLingDanSubTypeAnQi,
		ItemAdditionsysTongLingDanSubTypeFaBao,
		ItemAdditionsysTongLingDanSubTypeXianTi,
		ItemAdditionsysTongLingDanSubTypeLingYu,
		ItemAdditionsysTongLingDanSubTypeShenFa,
		ItemAdditionsysTongLingDanSubTypeShiHunFan,
		ItemAdditionsysTongLingDanSubTypeTianMoTi,
		ItemAdditionsysTongLingDanSubTypeLingBing,
		ItemAdditionsysTongLingDanSubTypeLingQi,
		ItemAdditionsysTongLingDanSubTypeLingWing,
		ItemAdditionsysTongLingDanSubTypeLingShen,
		ItemAdditionsysTongLingDanSubTypeLingArea,
		ItemAdditionsysTongLingDanSubTypeLingBao,
		ItemAdditionsysTongLingDanSubTypeLingTi:
		return true
	}
	return false
}

func CreateItemAdditionsysTongLingDanSubType(subType int32) ItemSubType {
	return ItemAdditionsysTongLingDanSubType(subType)
}

// 称号子类型
type ItemTitleSubType int32

const (
	//称号卡
	ItemTitleSubTypeCard ItemTitleSubType = iota
	//定制称号卡
	ItemTitleSubTypeDingZhiCard
)

func (t ItemTitleSubType) SubType() int32 {
	return int32(t)
}

func (t ItemTitleSubType) Valid() bool {
	switch t {
	case ItemTitleSubTypeCard,
		ItemTitleSubTypeDingZhiCard:
		return true
	}
	return false
}

func CreateItemTitleSubType(subType int32) ItemSubType {
	return ItemTitleSubType(subType)
}

//青龙装备子类型
type ShengHenEquipQingLongSubType int32

const (
	ShengHenEquipQingLongSubTypeOne   ShengHenEquipQingLongSubType = iota //装备1
	ShengHenEquipQingLongSubTypeTwo                                       //装备2
	ShengHenEquipQingLongSubTypeThree                                     //装备3
	ShengHenEquipQingLongSubTypeFour                                      //装备4
)

func (t ShengHenEquipQingLongSubType) SubType() int32 {
	return int32(t)
}

func (t ShengHenEquipQingLongSubType) Valid() bool {
	switch t {
	case ShengHenEquipQingLongSubTypeOne,
		ShengHenEquipQingLongSubTypeTwo,
		ShengHenEquipQingLongSubTypeThree,
		ShengHenEquipQingLongSubTypeFour:
		return true
	}
	return false
}

func CreateShengHenEquipQingLongSubType(subType int32) ItemSubType {
	return ShengHenEquipQingLongSubType(subType)
}

//白虎装备子类型
type ShengHenEquipBaiHuSubType int32

const (
	ShengHenEquipBaiHuSubTypeOne   ShengHenEquipBaiHuSubType = iota //装备1
	ShengHenEquipBaiHuSubTypeTwo                                    //装备2
	ShengHenEquipBaiHuSubTypeThree                                  //装备3
	ShengHenEquipBaiHuSubTypeFour                                   //装备4
)

func (t ShengHenEquipBaiHuSubType) SubType() int32 {
	return int32(t)
}

func (t ShengHenEquipBaiHuSubType) Valid() bool {
	switch t {
	case ShengHenEquipBaiHuSubTypeOne,
		ShengHenEquipBaiHuSubTypeTwo,
		ShengHenEquipBaiHuSubTypeThree,
		ShengHenEquipBaiHuSubTypeFour:
		return true
	}
	return false
}

func CreateShengHenEquipBaiHuSubType(subType int32) ItemSubType {
	return ShengHenEquipBaiHuSubType(subType)
}

//朱雀装备子类型
type ShengHenEquipZhuQueSubType int32

const (
	ShengHenEquipZhuQueSubTypeOne   ShengHenEquipZhuQueSubType = iota //装备1
	ShengHenEquipZhuQueSubTypeTwo                                     //装备2
	ShengHenEquipZhuQueSubTypeThree                                   //装备3
	ShengHenEquipZhuQueSubTypeFour                                    //装备4
)

func (t ShengHenEquipZhuQueSubType) SubType() int32 {
	return int32(t)
}

func (t ShengHenEquipZhuQueSubType) Valid() bool {
	switch t {
	case ShengHenEquipZhuQueSubTypeOne,
		ShengHenEquipZhuQueSubTypeTwo,
		ShengHenEquipZhuQueSubTypeThree,
		ShengHenEquipZhuQueSubTypeFour:
		return true
	}
	return false
}

func CreateShengHenEquipZhuQueSubType(subType int32) ItemSubType {
	return ShengHenEquipZhuQueSubType(subType)
}

//玄武装备子类型
type ShengHenEquipXuanWuSubType int32

const (
	ShengHenEquipXuanWuSubTypeOne   ShengHenEquipXuanWuSubType = iota //装备1
	ShengHenEquipXuanWuSubTypeTwo                                     //装备2
	ShengHenEquipXuanWuSubTypeThree                                   //装备3
	ShengHenEquipXuanWuSubTypeFour                                    //装备4
)

func (t ShengHenEquipXuanWuSubType) SubType() int32 {
	return int32(t)
}

func (t ShengHenEquipXuanWuSubType) Valid() bool {
	switch t {
	case ShengHenEquipXuanWuSubTypeOne,
		ShengHenEquipXuanWuSubTypeTwo,
		ShengHenEquipXuanWuSubTypeThree,
		ShengHenEquipXuanWuSubTypeFour:
		return true
	}
	return false
}

func CreateShengHenEquipXuanWuSubType(subType int32) ItemSubType {
	return ShengHenEquipXuanWuSubType(subType)
}

//圣痕道具子类型
type ShengHenSubType int32

const (
	ShengHenSubTypeQingLong ShengHenSubType = iota //青龙
	ShengHenSubTypeBaiHu                           //白虎
	ShengHenSubTypeZhuQue                          //朱雀
	ShengHenSubTypeXuanWu                          //玄武
)

func (t ShengHenSubType) SubType() int32 {
	return int32(t)
}

func (t ShengHenSubType) Valid() bool {
	switch t {
	case ShengHenSubTypeQingLong,
		ShengHenSubTypeBaiHu,
		ShengHenSubTypeZhuQue,
		ShengHenSubTypeXuanWu:
		return true
	}
	return false
}

func CreateShengHenSubType(subType int32) ItemSubType {
	return ShengHenSubType(subType)
}

//屠龙装备子类型
type TuLongEquipSubType int32

const (
	TuLongEquipSubTypeLongYa  TuLongEquipSubType = iota //屠龙装备-龙牙套
	TuLongEquipSubTypeLongLin                           //屠龙装备-龙鳞套
	TuLongEquipSubTypeLongXie                           //屠龙装备-龙血套
	TuLongEquipSubTypeLongDan                           //屠龙装备-龙胆套
	TuLongEquipSubTypeLongPo                            //屠龙装备-龙魂套
)

func (t TuLongEquipSubType) SubType() int32 {
	return int32(t)
}

func (t TuLongEquipSubType) Valid() bool {
	switch t {
	case TuLongEquipSubTypeLongYa,
		TuLongEquipSubTypeLongLin,
		TuLongEquipSubTypeLongXie,
		TuLongEquipSubTypeLongDan,
		TuLongEquipSubTypeLongPo:
		return true
	}
	return false
}

func CreateTuLongEquipSubType(subType int32) ItemSubType {
	return TuLongEquipSubType(subType)
}

//神器子类型
type ItemShenQiSubType int32

const (
	//神器碎片
	ItemShenQiSubTypeDebris ItemShenQiSubType = iota
	//神器器灵
	ItemShenQiSubTypeQiLing
	//神器元素之魂
	ItemShenQiSubTypeElem
)

func (iest ItemShenQiSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemShenQiSubType) Valid() bool {
	switch iest {
	case ItemShenQiSubTypeDebris,
		ItemShenQiSubTypeQiLing,
		ItemShenQiSubTypeElem:
		return true
	}
	return false
}

func CreateItemShenQiSubType(subType int32) ItemSubType {
	return ItemShenQiSubType(subType)
}

//宝宝相关子类型
type ItemBaoBaoSubType int32

const (
	ItemBaoBaoSubTypeDongFang       ItemBaoBaoSubType = iota //0洞房道具
	ItemBaoBaoSubTypeRiver                                   //子母河道具
	ItemBaoBaoSubTypeTonic                                   //补品道具
	ItemBaoBaoSubTypeBabyRenameCard                          //宝宝改名卡
	ItemBaoBaoSubTypeBabyLearn                               //宝宝读书道具
	ItemBaoBaoSubTypeBabyXiLian                              //5宝宝洗练道具
	ItemBaoBaoSubTypeBaby                                    //宝宝
)

func (iest ItemBaoBaoSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemBaoBaoSubType) Valid() bool {
	switch iest {
	case ItemBaoBaoSubTypeDongFang,
		ItemBaoBaoSubTypeRiver,
		ItemBaoBaoSubTypeTonic,
		ItemBaoBaoSubTypeBabyRenameCard,
		ItemBaoBaoSubTypeBabyLearn,
		ItemBaoBaoSubTypeBabyXiLian,
		ItemBaoBaoSubTypeBaby:
		return true
	}
	return false
}

func CreateItemBaoBaoSubType(subType int32) ItemSubType {
	return ItemBaoBaoSubType(subType)
}

//玩具相关子类型
type ItemBabyToySubType int32

const (
	ItemBabyToySubTypeOne   ItemBabyToySubType = iota //套餐1
	ItemBabyToySubTypeTwo                             //套餐2
	ItemBabyToySubTypeThree                           //套餐3
	ItemBabyToySubTypeFour                            //套餐4
	ItemBabyToySubTypeFive                            //套餐5
	ItemBabyToySubTypeSix                             //套餐6
)

func (iest ItemBabyToySubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemBabyToySubType) Valid() bool {
	switch iest {
	case ItemBabyToySubTypeOne,
		ItemBabyToySubTypeTwo,
		ItemBabyToySubTypeThree,
		ItemBabyToySubTypeFour,
		ItemBabyToySubTypeFive,
		ItemBabyToySubTypeSix:
		return true
	}
	return false
}

func CreateItemBabyToySubType(subType int32) ItemSubType {
	return ItemBabyToySubType(subType)
}

//表白相关子类型
type ItemBiaoBaiSubType int32

const (
	ItemBiaoBaiSubTypeCash ItemBiaoBaiSubType = iota + 1 //钞票
	ItemBiaoBaiSubTypeCar                                //兰博基尼
	ItemBiaoBaiSubTypeShip                               //游艇
)

func (iest ItemBiaoBaiSubType) SubType() int32 {
	return int32(iest)
}

func (iest ItemBiaoBaiSubType) Valid() bool {
	switch iest {
	case ItemBiaoBaiSubTypeCash,
		ItemBiaoBaiSubTypeCar,
		ItemBiaoBaiSubTypeShip:
		return true
	}
	return false
}

func CreateItemBiaoBaiSubType(subType int32) ItemSubType {
	return ItemBiaoBaiSubType(subType)
}

//神铸目标子类型
type ItemGodCastingEquipSubType int32

const (
	ItemGodCastingEquipSubTypeEquip              ItemGodCastingEquipSubType = iota //神铸装备
	ItemGodCastingEquipSubTypeCrystal                                              //神铸结晶
	ItemGodCastingEquipSubTypeGodCastingStone                                      //神铸石
	ItemGodCastingEquipSubTypeForgeSoulStone                                       //锻魂石
	ItemGodCastingEquipSubTypeCastingSpiritStone                                   //铸灵石
	ItemGodCastingEquipSubTypeBossTicket                                           //神铸BOSS入场卷
)

func (i ItemGodCastingEquipSubType) SubType() int32 {
	return int32(i)
}

func (subType ItemGodCastingEquipSubType) Valid() bool {
	switch subType {
	case ItemGodCastingEquipSubTypeEquip,
		ItemGodCastingEquipSubTypeCrystal,
		ItemGodCastingEquipSubTypeGodCastingStone,
		ItemGodCastingEquipSubTypeForgeSoulStone,
		ItemGodCastingEquipSubTypeCastingSpiritStone,
		ItemGodCastingEquipSubTypeBossTicket:
		return true
	}
	return false
}

func CreateItemGodCastingEquipSubType(subType int32) ItemSubType {
	return ItemGodCastingEquipSubType(subType)
}

//满汉全席合成物品子类型
type ItemManHanQuanXiSubType int32

const (
	ItemManHanQuanXiSubTypeZhuShi  ItemManHanQuanXiSubType = iota // 主食
	ItemManHanQuanXiSubTypePeiCai                                 // 配菜
	ItemManHanQuanXiSubTypePeiLiao                                // 配料
)

func (i ItemManHanQuanXiSubType) SubType() int32 {
	return int32(i)
}

func (subType ItemManHanQuanXiSubType) Valid() bool {
	switch subType {
	case ItemManHanQuanXiSubTypeZhuShi,
		ItemManHanQuanXiSubTypePeiCai,
		ItemManHanQuanXiSubTypePeiLiao:
		return true
	}
	return false
}

func CreateItemManHanQuanXiSubType(subType int32) ItemSubType {
	return ItemManHanQuanXiSubType(subType)
}
