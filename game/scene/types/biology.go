package types

var (
	biologyImmuneMap = map[BiologySetType]BuffImmuneType{
		BiologySetTypePlayer:          BuffImmuneTypePlayer,
		BiologySetTypeNPC:             BuffImmuneTypeMonster,
		BiologySetTypeMonster:         BuffImmuneTypeMonster,
		BiologySetTypeBoss:            BuffImmuneTypeBoss,
		BiologySetTypeWorldBoss:       BuffImmuneTypeWorldBoss,
		BiologySetTypeRealm:           BuffImmuneTypeRealmBoss,
		BiologySetTypeSpecialBiaoChe:  BuffImmuneTypeSpecialBiologyBiaoChe,
		BiologySetTypeSpecialChengMen: BuffImmuneTypeSpecialBiologyChengMen,
		BiologySetTypeSpecialMonster:  BuffImmuneTypeSpecialMonster,
		BiologySetTypeCangJingGeBoss:  BuffImmuneTypeCangJingGeBoss,
	}
)

type BiologySetType int32

const (
	BiologySetTypePlayer BiologySetType = 0
	//npc
	BiologySetTypeNPC             = 1
	BiologySetTypeMonster         = 2
	BiologySetTypeBoss            = 3
	BiologySetTypeWorldBoss       = 4
	BiologySetTypeRealm           = 5
	BiologySetTypeSpecialBiaoChe  = 6
	BiologySetTypeSpecialChengMen = 7
	BiologySetTypeSpecialMonster  = 8
	BiologySetTypePet             = 9
	BiologySetTypeFlyPet          = 10
	BiologySetTypeCangJingGeBoss  = 11
	BiologySetTypeExp             = 12
)

func (t BiologySetType) Valid() bool {
	switch t {
	case BiologySetTypeNPC,
		BiologySetTypeMonster,
		BiologySetTypeBoss,
		BiologySetTypeWorldBoss,
		BiologySetTypeRealm,
		BiologySetTypeSpecialBiaoChe,
		BiologySetTypeSpecialChengMen,
		BiologySetTypeSpecialMonster,
		BiologySetTypePet,
		BiologySetTypeFlyPet,
		BiologySetTypeCangJingGeBoss,
		BiologySetTypeExp:
		return true
	}
	return false
}

var (
	biologySetTypeMap = map[BiologySetType]string{
		BiologySetTypePlayer:          "玩家",
		BiologySetTypeNPC:             "npc",
		BiologySetTypeMonster:         "怪物",
		BiologySetTypeBoss:            "boss",
		BiologySetTypeWorldBoss:       "世界boss",
		BiologySetTypeRealm:           "天劫塔怪物",
		BiologySetTypeSpecialBiaoChe:  "非生物类型(镖车)",
		BiologySetTypeSpecialChengMen: "非生物类型(城门)",
		BiologySetTypeSpecialMonster:  "普通怪物(邪将传令官)",
		BiologySetTypePet:             "宠物",
		BiologySetTypeFlyPet:          "飞宠",
		BiologySetTypeCangJingGeBoss:  "藏经阁boss",
		BiologySetTypeExp:             "经验副本怪物",
	}
)

func (bt BiologySetType) Immune() int32 {
	immune, ok := biologyImmuneMap[bt]
	if !ok {
		return 0
	}
	return int32(immune)
}

//客户端使用
type BiologyType int32

var (
	biologyTypeMap = map[BiologyType]string{
		BiologyTypeNPC:             "npc",
		BiologyTypeFuncNPC:         "功能npc",
		BiologyTypeBattleNPC:       "战斗npc",
		BiologyTypeMonster:         "怪物",
		BiologyTypeCountMonster:    "统计怪物",
		BiologyTypeSeniorMonster:   "精英怪",
		BiologyTypeBoss:            "boss",
		BiologyTypeBuildingMonster: "建筑怪",
		BiologyTypeClickCollect:    "可以点击的采集物",
		BiologyTypeChenZhanCollect: "城战复活点",
		BiologyTransmissionArray:   "传送阵",
		BiologyTransmissionPoint:   "传送点",
		BiologyTypePlayer:          "玩家",
		BiologyTypeItem:            "物品",
	}
)

func (bt BiologyType) String() string {
	return biologyTypeMap[bt]
}

const (
	//npc
	BiologyTypeNPC             BiologyType = 0
	BiologyTypeFuncNPC                     = 1
	BiologyTypeBattleNPC                   = 2
	BiologyTypeMonster                     = 3
	BiologyTypeCountMonster                = 4
	BiologyTypeSeniorMonster               = 5
	BiologyTypeBoss                        = 6
	BiologyTypeBuildingMonster             = 7
	BiologyTypeClickCollect                = 8
	BiologyTypeChenZhanCollect             = 9
	BiologyTypeCollect                     = 20
	BiologyTransmissionArray               = 21
	BiologyTransmissionPoint               = 22
	BiologyTypeItem                        = 98
	BiologyTypePlayer                      = 100
	BiologyTypePet                         = 101
	BiologyTypeFlyPet                      = 102
)

func (bt BiologyType) Valid() bool {
	switch bt {
	case BiologyTypeNPC,
		BiologyTypeFuncNPC,
		BiologyTypeMonster,
		BiologyTypeBoss,
		BiologyTypeCountMonster,
		BiologyTypeSeniorMonster,
		BiologyTypeBuildingMonster,
		BiologyTypeBattleNPC,
		BiologyTypeClickCollect,
		BiologyTypeChenZhanCollect,
		BiologyTypeCollect,
		BiologyTransmissionArray,
		BiologyTransmissionPoint,
		BiologyTypePlayer,
		BiologyTypeItem,
		BiologyTypePet,
		BiologyTypeFlyPet:
		return true
	}
	return false
}

var (
	biologyTypeDeadIgnoreMap = map[BiologyType]bool{
		BiologyTypeNPC:             true,
		BiologyTypeFuncNPC:         true,
		BiologyTypeBattleNPC:       true,
		BiologyTypeMonster:         true,
		BiologyTypeCountMonster:    true,
		BiologyTypeSeniorMonster:   true,
		BiologyTypeBoss:            true,
		BiologyTypeBuildingMonster: true,
		BiologyTypeClickCollect:    false,
		BiologyTypeChenZhanCollect: true,
		BiologyTransmissionArray:   true,
		BiologyTransmissionPoint:   true,
		BiologyTypePlayer:          true,
		BiologyTypeItem:            true,
	}
)

func (bt BiologyType) DeadIgnore() bool {
	return biologyTypeDeadIgnoreMap[bt]
}

type BiologyScriptType int32

const (
	//npc
	BiologyScriptTypeNPC                    BiologyScriptType = 0
	BiologyScriptTypeMonster                                  = 1
	BiologyScriptTypeBattleNPC                                = 2
	BiologyScriptTypeRobber                                   = 3
	BiologyScriptTypeBiaoChe                                  = 4
	BiologyScriptTypeFourGodSpecial                           = 5
	BiologyScriptTypeFourGodCollect                           = 6
	BiologyScriptTypeRelivePoint                              = 7
	BiologyScriptTypeXianMengNPC                              = 8
	BiologyScriptTypeBuildingMonster                          = 9
	BiologyScriptTypeFourGodBoss                              = 10
	BiologyScriptTypeSoulBoss                                 = 11
	BiologyScriptTypeWorldBoss                                = 12
	BiologyScriptTypeWedBanquet                               = 13
	BiologyScriptTypeWeddingCar                               = 14
	BiologyScriptTypeOneArenaGuardian                         = 15
	BiologyScriptTypeArenaExpTree                             = 16
	BiologyScriptTypeArenaShengShou                           = 17
	BiologyScriptTypeArenaTreasure                            = 18
	BiologyScriptTypeCrossWorldBoss                           = 19
	BiologyScriptTypeCrossBigBoss                             = 20
	BiologyScriptTypeCrossSmallBoss                           = 21
	BiologyScriptTypeCrossBigEgg                              = 22
	BiologyScriptTypeCrossSmallEgg                            = 23
	BiologyScriptTypeBossCallTicket                           = 24
	BiologyScriptTypeGeneralCollect                           = 25
	BiologyScriptTypeCrossLianYuBoss                          = 26
	BiologyScriptTypePet                                      = 27
	BiologyScriptTypeTowerMonster                             = 28
	BiologyScriptTypeTowerBoss                                = 29
	BiologyScriptTypeMyBoss                                   = 30
	BiologyScriptTypeVIPMyBoss                                = 31
	BiologyScriptTypeGodSiegeBoss                             = 32
	BiologyScriptTypeUnrealBoss                               = 33
	BiologyScriptTypeOutlandBoss                              = 34
	BiologyScriptTypeFeiChong                                 = 35
	BiologyScriptTypeCangJingGeBoss                           = 36
	BiologyScriptTypeAllianceBoss                             = 37
	BiologyScriptTypeAllianceShengTan                         = 38
	BiologyScriptTypeXianTaoQianNianCollect                   = 39
	BiologyScriptTypeXianTaoBaiNianCollect                    = 40
	BiologyScriptTypeShenYuBoss                               = 41
	BiologyScriptTypeLongGongBoss                             = 42
	BiologyScriptTypePearl                                    = 43
	BiologyScriptTypeLongGongTreasure                         = 44
	BiologyScriptTypeQiYuDao                                  = 45
	BiologyScriptTypeZhenXiBoss                               = 46
	BiologyScriptTypeDingShiBoss                              = 47
	BiologyScriptTypeArenaBossHuWei                           = 48
)

func (bt BiologyScriptType) Valid() bool {
	switch bt {
	case BiologyScriptTypeNPC,
		BiologyScriptTypeMonster,
		BiologyScriptTypeBattleNPC,
		BiologyScriptTypeRobber,
		BiologyScriptTypeBiaoChe,
		BiologyScriptTypeFourGodSpecial,
		BiologyScriptTypeFourGodCollect,
		BiologyScriptTypeRelivePoint,
		BiologyScriptTypeXianMengNPC,
		BiologyScriptTypeBuildingMonster,
		BiologyScriptTypeFourGodBoss,
		BiologyScriptTypeSoulBoss,
		BiologyScriptTypeWorldBoss,
		BiologyScriptTypeWedBanquet,
		BiologyScriptTypeWeddingCar,
		BiologyScriptTypeOneArenaGuardian,
		BiologyScriptTypeArenaExpTree,
		BiologyScriptTypeArenaShengShou,
		BiologyScriptTypeArenaTreasure,
		BiologyScriptTypeCrossWorldBoss,
		BiologyScriptTypeCrossBigBoss,
		BiologyScriptTypeCrossSmallBoss,
		BiologyScriptTypeCrossBigEgg,
		BiologyScriptTypeCrossSmallEgg,
		BiologyScriptTypeBossCallTicket,
		BiologyScriptTypeGeneralCollect,
		BiologyScriptTypeCrossLianYuBoss,
		BiologyScriptTypePet,
		BiologyScriptTypeTowerMonster,
		BiologyScriptTypeTowerBoss,
		BiologyScriptTypeMyBoss,
		BiologyScriptTypeGodSiegeBoss,
		BiologyScriptTypeUnrealBoss,
		BiologyScriptTypeOutlandBoss,
		BiologyScriptTypeFeiChong,
		BiologyScriptTypeCangJingGeBoss,
		BiologyScriptTypeAllianceBoss,
		BiologyScriptTypeAllianceShengTan,
		BiologyScriptTypeXianTaoQianNianCollect,
		BiologyScriptTypeXianTaoBaiNianCollect,
		BiologyScriptTypeShenYuBoss,
		BiologyScriptTypeLongGongBoss,
		BiologyScriptTypePearl,
		BiologyScriptTypeLongGongTreasure,
		BiologyScriptTypeQiYuDao,
		BiologyScriptTypeZhenXiBoss,
		BiologyScriptTypeDingShiBoss,
		BiologyScriptTypeArenaBossHuWei:
		return true
	}
	return false
}

var (
	BiologyScriptTypeMap = map[BiologyScriptType]string{
		BiologyScriptTypeNPC:                    "npc",
		BiologyScriptTypeMonster:                "怪物",
		BiologyScriptTypeBattleNPC:              "战斗npc",
		BiologyScriptTypeRobber:                 "robber",
		BiologyScriptTypeBiaoChe:                "镖车",
		BiologyScriptTypeFourGodSpecial:         "四神特殊怪",
		BiologyScriptTypeFourGodCollect:         "四神采集",
		BiologyScriptTypeRelivePoint:            "复活点",
		BiologyScriptTypeXianMengNPC:            "皇城战守卫",
		BiologyScriptTypeBuildingMonster:        "建筑怪",
		BiologyScriptTypeFourGodBoss:            "四神boss",
		BiologyScriptTypeSoulBoss:               "帝陵boss",
		BiologyScriptTypeWorldBoss:              "世界boss",
		BiologyScriptTypeWedBanquet:             "结婚酒席",
		BiologyScriptTypeWeddingCar:             "婚车",
		BiologyScriptTypeOneArenaGuardian:       "灵池守护者",
		BiologyScriptTypeArenaExpTree:           "四圣兽经验树",
		BiologyScriptTypeArenaShengShou:         "四圣兽",
		BiologyScriptTypeArenaTreasure:          "四圣兽宝箱",
		BiologyScriptTypeCrossWorldBoss:         "跨服世界boss",
		BiologyScriptTypeCrossBigBoss:           "跨服屠龙大boss",
		BiologyScriptTypeCrossSmallBoss:         "跨服屠龙小boss",
		BiologyScriptTypeCrossBigEgg:            "跨服屠龙大龙蛋",
		BiologyScriptTypeCrossSmallEgg:          "跨服屠龙小龙蛋",
		BiologyScriptTypeBossCallTicket:         "boss召唤券",
		BiologyScriptTypeGeneralCollect:         "通用采集物",
		BiologyScriptTypeCrossLianYuBoss:        "无间炼狱boss",
		BiologyScriptTypePet:                    "幻兽",
		BiologyScriptTypeTowerMonster:           "打宝塔小怪",
		BiologyScriptTypeTowerBoss:              "打宝塔boss",
		BiologyScriptTypeMyBoss:                 "个人boss",
		BiologyScriptTypeGodSiegeBoss:           "神兽攻城boss",
		BiologyScriptTypeUnrealBoss:             "幻境boss",
		BiologyScriptTypeOutlandBoss:            "外域boss",
		BiologyScriptTypeFeiChong:               "飞宠",
		BiologyScriptTypeCangJingGeBoss:         "藏经阁boss",
		BiologyScriptTypeAllianceBoss:           "仙盟Boss",
		BiologyScriptTypeAllianceShengTan:       "仙盟圣坛",
		BiologyScriptTypeXianTaoQianNianCollect: "仙桃千年采集点",
		BiologyScriptTypeXianTaoBaiNianCollect:  "仙桃百年采集点",
		BiologyScriptTypeShenYuBoss:             "神域BOSS",
		BiologyScriptTypeLongGongBoss:           "龙宫BOSS",
		BiologyScriptTypePearl:                  "珍珠采集物",
		BiologyScriptTypeLongGongTreasure:       "黑龙财宝采集点",
		BiologyScriptTypeQiYuDao:                "奇遇岛怪物",
		BiologyScriptTypeZhenXiBoss:             "珍稀BOSS",
		BiologyScriptTypeDingShiBoss:            "定时boss",
		BiologyScriptTypeArenaBossHuWei:         "圣兽护卫",
	}
)

func (bt BiologyScriptType) String() string {
	return BiologyScriptTypeMap[bt]
}

//阵营
type FactionType int32

const (
	//普通玩家
	FactionTypePlayer FactionType = iota + 1
	//npc
	FactionTypeNPC
	//怪物
	FactionTypeMonster
	//与玩家同阵营的npc
	FactionTypeNPCCampOfPlayer
	//城战npc
	FactionTypeChengZhanNPC          //城战npc
	FactionTypeChengZhanAttackPlayer //城战攻方
	FactionTypeChengZhanDefendPlayer //城战守方
	FactionTypeChengZhanDefendNPC    //城战守方
	FactionTypeModel                 //雕像

)

func (ft FactionType) Valid() bool {
	switch ft {
	case FactionTypePlayer,
		FactionTypeNPC,
		FactionTypeMonster,
		FactionTypeNPCCampOfPlayer,
		FactionTypeChengZhanNPC,
		FactionTypeChengZhanAttackPlayer,
		FactionTypeChengZhanDefendPlayer,
		FactionTypeChengZhanDefendNPC,
		FactionTypeModel:

		return true
	}
	return false
}

var (
	factionTypeMap = map[FactionType]map[FactionType]FactionRelationType{
		FactionTypePlayer: map[FactionType]FactionRelationType{
			FactionTypeMonster: FactionRelationTypeEnemy,
		},
		FactionTypeNPC: map[FactionType]FactionRelationType{},
		FactionTypeMonster: map[FactionType]FactionRelationType{
			FactionTypePlayer:                FactionRelationTypeEnemy,
			FactionTypeNPCCampOfPlayer:       FactionRelationTypeEnemy,
			FactionTypeChengZhanDefendPlayer: FactionRelationTypeEnemy,
			FactionTypeChengZhanAttackPlayer: FactionRelationTypeEnemy,
		},
		FactionTypeNPCCampOfPlayer: map[FactionType]FactionRelationType{
			FactionTypeMonster: FactionRelationTypeEnemy,
		},
		FactionTypeChengZhanNPC: map[FactionType]FactionRelationType{},
		FactionTypeChengZhanAttackPlayer: map[FactionType]FactionRelationType{
			FactionTypeChengZhanDefendNPC:    FactionRelationTypeEnemy,
			FactionTypeChengZhanDefendPlayer: FactionRelationTypeEnemy,
		},
		FactionTypeChengZhanDefendPlayer: map[FactionType]FactionRelationType{
			FactionTypeChengZhanAttackPlayer: FactionRelationTypeEnemy,
		},
		FactionTypeChengZhanDefendNPC: map[FactionType]FactionRelationType{
			FactionTypeChengZhanAttackPlayer: FactionRelationTypeEnemy,
		},
	}
)

func (ft FactionType) IsEnemy(ft2 FactionType) bool {

	return factionTypeMap[ft][ft2] == FactionRelationTypeEnemy
}

//阵营关系
type FactionRelationType int32

const (
	FactionRelationTypeAlliance FactionRelationType = iota
	FactionRelationTypeEnemy
)

func (ft FactionRelationType) Valid() bool {
	switch ft {
	case FactionRelationTypeAlliance,
		FactionRelationTypeEnemy:
		return true
	}
	return false
}

//主被动
type ThreatType int32

const (
	ThreateTypePositive ThreatType = iota + 1
	ThreateTypePasstive
)

func (tt ThreatType) Valid() bool {
	switch tt {
	case ThreateTypePositive,
		ThreateTypePasstive:
		return true
	}
	return false
}

type BiologyAutoRecoverType int32

const (
	BiologyAutoRecoverTypeYes = iota + 1
	BiologyAutoRecoverTypeNo
)

func (tt BiologyAutoRecoverType) Valid() bool {
	switch tt {
	case BiologyAutoRecoverTypeYes,
		BiologyAutoRecoverTypeNo:
		return true
	}
	return false
}

type BiologyRebornType int32

const (
	//几秒后
	BiologyRebornTypeSecond BiologyRebornType = iota
	//定时
	BiologyRebornTypeTime
	//召唤
	BiologyRebornTypeCall
)

func (brt BiologyRebornType) Valid() bool {
	switch brt {
	case BiologyRebornTypeSecond,
		BiologyRebornTypeTime,
		BiologyRebornTypeCall:
		return true
	}
	return false
}
