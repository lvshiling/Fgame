package quest

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	itemtypes "fgame/fgame/game/item/types"
)

const (
	KaiFuMuBiaoDayMax    = 7
	QiYuEquipGiftGroupId = 79
	QiYuEquipGiftIndex   = 0
)

type QuestType int32

const (
	//一次性
	QuestTypeOnce QuestType = iota + 1
	//支线
	QuestTypeBranch
	//屠魔
	QuestTypeTuMo
	//天机牌
	QuestTypeTianJiPai
	//活跃度任务
	QuestTypeLiveness
	//日环任务
	QuestTypeDaily
	//开服目标任务
	QuestTypeKaiFuMuBiao
	//仙盟日环任务
	QuestTypeDailyAlliance
	//奇遇任务---9
	QuestTypeQiYu
	//运营活动目标
	QuestTypeYunYingGoal
	//创世任务
	QuestTypeChuangShi
)

func (qt QuestType) Valid() bool {
	switch qt {
	case QuestTypeOnce,
		QuestTypeBranch,
		QuestTypeTuMo,
		QuestTypeTianJiPai,
		QuestTypeLiveness,
		QuestTypeDaily,
		QuestTypeKaiFuMuBiao,
		QuestTypeDailyAlliance,
		QuestTypeQiYu,
		QuestTypeYunYingGoal,
		QuestTypeChuangShi:
		return true
	}
	return false
}

var (
	questTypeMap = map[QuestType]string{
		QuestTypeOnce:          "一次性任务",
		QuestTypeBranch:        "支线",
		QuestTypeTuMo:          "屠魔",
		QuestTypeTianJiPai:     "天机牌",
		QuestTypeLiveness:      "活跃度任务",
		QuestTypeDaily:         "日环任务",
		QuestTypeKaiFuMuBiao:   "开服目标任务",
		QuestTypeDailyAlliance: "仙盟日环任务",
		QuestTypeQiYu:          "奇遇任务",
		QuestTypeYunYingGoal:   "运营活动目标任务",
		QuestTypeChuangShi:     "创世之战任务",
	}
)

//任务嵌套任务
var questNestedQuestMap = map[QuestType]QuestSubType{
	QuestTypeTuMo:          QuestSubTypeFinishTuMo,
	QuestTypeTianJiPai:     QuestSubTypeFinishSecretCard,
	QuestTypeDaily:         QuestSubTypeFinishDailyQuestNum,
	QuestTypeDailyAlliance: QuestSubTypeFinishDailyAllianceQuestNum,
}

func (qt QuestType) QuestNestedSubType() (subTyp QuestSubType, sucess bool) {
	subTyp, sucess = questNestedQuestMap[qt]
	if !sucess {
		return
	}
	return
}

func (qt QuestType) GetDailyTag() (tag QuestDailyTag, flag bool) {
	dailyTag, ok := dailyTagMap[qt]
	if !ok {
		return
	}
	flag = true
	tag = dailyTag
	return
}

type SystemReachXType int32

const (
	//坐骑系统
	SystemReachXTypeMount SystemReachXType = iota
	//暗器系统
	SystemReachXTypeAnQi
	//战翼系统
	SystemReachXTypeWing
	//护体盾系统
	SystemReachXTypeBodyShield
	//领域系统
	SystemReachXTypeLingYu
	//身法系统
	SystemReachXTypeShenFa
	//法宝系统
	SystemReachXTypeFaBao
	//仙体系统
	SystemReachXTypeXianTi
	//戮仙刃
	SystemReachXTypeLuXianRen
	//血盾
	SystemReachXTypeXueDun
	//点星
	SystemReachXTypeDianXing //10
	//噬魂幡
	SystemReachXTypeShiHunFan
	//天魔体
	SystemReachXTypeTianMoTi
)

const (
	//灵童兵魂
	SystemReachXTypeLingTongWeapon SystemReachXType = iota + 101
	//灵童坐骑
	SystemReachXTypeLingTongMount
	//灵童战翼
	SystemReachXTypeLingTongWing
	//灵童身法
	SystemReachXTypeLingTongShenFa
	//灵童领域
	SystemReachXTypeLingTongLingYu
	//灵童法宝
	SystemReachXTypeLingTongFaBao
	//灵童仙体
	SystemReachXTypeLingTongXianTi
)

func (st SystemReachXType) Valid() bool {
	switch st {
	case SystemReachXTypeMount,
		SystemReachXTypeAnQi,
		SystemReachXTypeWing,
		SystemReachXTypeBodyShield,
		SystemReachXTypeLingYu,
		SystemReachXTypeShenFa,
		SystemReachXTypeFaBao,
		SystemReachXTypeXianTi,
		SystemReachXTypeLuXianRen,
		SystemReachXTypeXueDun,
		SystemReachXTypeDianXing,
		SystemReachXTypeShiHunFan,
		SystemReachXTypeTianMoTi,
		SystemReachXTypeLingTongWeapon,
		SystemReachXTypeLingTongMount,
		SystemReachXTypeLingTongWing,
		SystemReachXTypeLingTongShenFa,
		SystemReachXTypeLingTongLingYu,
		SystemReachXTypeLingTongFaBao,
		SystemReachXTypeLingTongXianTi:
		return true
	}
	return false
}

type QuestSubType int32

const (
	QuestSubTypeDialog QuestSubType = iota //0
	QuestSubTypeKillMonster
	QuestSubTypeCollect
	QuestSubTypeLevel
	QuestSubTypeHurtMonster
	QuestSubTypeWorldChannel
	QuestSubTypeSendFlower
	QuestSubTypeSyntheticTuMoToken
	QuestSubTypeFinishTuMo
	QuestSubTypeSoulRuins
	QuestSubTypeSpecialXianFu //10
	QuestSubTypeReserve
	QuestSubTypeGamble
	QuestSubTypeRealm
	QuestSubTypeDart
	QuestSubType1V1 //15
	QuestSubType3V3
	QuestSubTypeMonsterHunt
	QuestSubTypeMoonLove
	QuestSubTypeAlliance
	QuestSubTypeEmperor //20
	QuestSubTypeDragonChess
	QuestSubTypeFourMysterious
	QuestSubTypeDartCar
	QuestSubTypeWorldBoss
	QuestSubTypeClickSkillUpgradeButton //25
	QuestSubTypeProfessionalSkillLevel
	QuestSubTypeProfessionalSkillTotalLevel
	QuestSubTypeClickEquipmentStrengthenButton
	QuestSubTypeEquipmentStrengthenLevel
	QuestSubTypeEquipmentStrengthenTotalLevel //30
	QuestSubTypeClickEquipmentUpgradeButton
	QuestSubTypeEquipmentUpgradeStar
	QuestSubTypeClickEquipmentUpgradeStarButton
	QuestSubTypeSpecifiedSoulRuins
	QuestSubTypeSoulActive //35
	QuestSubTypeClickSoulStrengthenButton
	QuestSubTypeSoulStrengthenLevel
	QuestSubTypeSoulSpecialEmbed
	QuestSubTypeClickSoulUpgradeButton
	QuestSubTypeSoulUpgradeLevel //40
	QuestSubTypeXianFu
	QuestSubTypePlayerLevel
	QuestSubTypeJoinAlliance
	QuestSubTypeApplyJoinAlliance
	QuestSubTypeCollectItemTotal //45
	QuestSubTypeCollectAllItem
	QuestSubTypeEquipmentUpgradeLevel
	QuestSubTypeEnterSpecialXianFu
	QuestSubTypeEnterXianFu
	QuestSubTypeEnterSoulRuins //50
	QuestSubTypeEnterRealm
	QuestSubTypeUpgradeSpecialXianFu
	QuestSubTypeFinishSecretCard
	QuestSubTypeRealmLevel
	QuestSubTypeEmbedGem //55
	QuestSubTypeZhuanSheng
	QuestSubTypeFinishFirstCharge
	QuestSubTypeActiveSkill
	QuestSubTypeSystemX
	QuestSubTypeEmbedQualityTwoGoldEquipNum //60
	QuestSubTypeEmbedQualityThreeGoldEquipNum
	QuestSubTypeEmbedQualityFourGoldEquipNum
	QuestSubTypeSomeMonster
	QuestSubTypeDayFirstCharge
	QuestSubTypeDaBaoTimePast //65
	QuestSubTypeAttandActivityPlay
	QuestSubTypeAttandKillSpecialMonster
	QuestSubTypeXianFuPersonal
	QuestSubTypechallengeMaterialFuBen
	QuestSubTypeBossKilled //70
	QuestSubTypeBiologyTypeKilled
	QuestSubTypechallengeSpecialMaterialFuBen
	QuestSubTypeAdvancedOperation
	QuestSubTypeEnterTower
	QuestSubTypeUpgradeSysOperation //75
	QuestSubTypeGoldEquipmentStrength
	QuestSubTypeFinishDailyQuestNum
	QuestSubTypeGoldEquipmentResolve
	QuestSubTypeBaGuaMiJingLevel
	QuestSubTypeSetTypeKilled //80
	QuestSubTypeLingTongClick
	QuestSubTypePassTeamCopy
	QuestSubTypeLingTongActivateNum
	QuestSubTypeGemTotalLevel
	QuestSubTypeFuBenMonsterGroup //85
	QuestSubTypeActivateSoulNum
	QuestSubTypeAwakenSoulNum
	QuestSubTypeChengZhanWinNum
	QuestSubType3V3WinNum
	QuestSubType3V3LianSheng //90
	QuestSubTypeShenLongFuHua
	QuestSubTypeActivateYiChuSuitNum
	QuestSubTypeOneArenaOccupyTime
	QuestSubTypeFeiShengLevel
	QuestSubTypeGoldEquipmentTotalLevel //95
	QuestSubTypeActivateYiChuSuitOne
	QuestSubTypeActivateYiChuSuitTwo
	QuestSubTypeActivateYiChuSuitThree
	QuestSubTypeActivateYiChuSuitFour
	QuestSubTypeBuyHuiYuan //100
	QuestSubTypeBuyInvest
	QuestSubTypeBuyEquipGift
	QuestSubTypeFinishDailyAllianceQuestNum
	QuestSubTypeAttendMajorFuBen
	QuestSubTypeAttendFuQiFuBen // 105
	QuestSubTypeFinishMajorFuBen
	QuestSubTypeFinishFuQiFuBen
	QuestSubTypeAllianceChat
	QuestSubTypeEquipBaoKuAttend
	QuestSubTypeFeiShengSanGong // 110
	QuestSubTypeAttendTeamFuBen
	QuestSubTypeFinishTeamFuBen
	QuestSubTypeGuideReplica
	QuestSubTypeWeekBuy
	QuestSubTypeTradUpload // 115
	QuestSubTypeEquipBaoKuDuiHuan
	QuestSubTypeChargeTimes
	QuestSubTypeForceGrow
	QuestSubTypeSendHongBao //119
	QuestSubTypeTradItem
	QuestSubTypeZuDuiFuBen
	QuestSubTypeSkillXinFa
	QuestSubTypeFinishJieYi
	QuestSubTypeCostGold //124
	QuestSubTypeArenaPVPJiFen
	QuestSubTypeXiongDiWeiMing
	QuestSubTypeXiongDiToken //127
	QuestSubTypeMarryTimes
	QuestSubTypeBiaoBaiTimes
	QuestSubTypeCoupleToeknForce
	QuestSubTypeChuangShiCollect
	QuestSubTypeChuangShiBiaoChe
	QuestSubTypeKillPlayer
)

func (qt QuestSubType) Valid() bool {
	switch qt {
	case QuestSubTypeDialog,
		QuestSubTypeKillMonster,
		QuestSubTypeCollect,
		QuestSubTypeLevel,
		QuestSubTypeHurtMonster,
		QuestSubTypeWorldChannel,
		QuestSubTypeSendFlower,
		QuestSubTypeSyntheticTuMoToken,
		QuestSubTypeFinishTuMo,
		QuestSubTypeSoulRuins,
		QuestSubTypeSpecialXianFu,
		QuestSubTypeReserve,
		QuestSubTypeGamble,
		QuestSubTypeRealm,
		QuestSubTypeDart,
		QuestSubType1V1,
		QuestSubType3V3,
		QuestSubTypeMonsterHunt,
		QuestSubTypeMoonLove,
		QuestSubTypeAlliance,
		QuestSubTypeEmperor,
		QuestSubTypeDragonChess,
		QuestSubTypeFourMysterious,
		QuestSubTypeDartCar,
		QuestSubTypeWorldBoss,
		QuestSubTypeClickSkillUpgradeButton,
		QuestSubTypeProfessionalSkillLevel,
		QuestSubTypeProfessionalSkillTotalLevel,
		QuestSubTypeClickEquipmentStrengthenButton,
		QuestSubTypeEquipmentStrengthenLevel,
		QuestSubTypeEquipmentStrengthenTotalLevel,
		QuestSubTypeClickEquipmentUpgradeButton,
		QuestSubTypeEquipmentUpgradeStar,
		QuestSubTypeClickEquipmentUpgradeStarButton,
		QuestSubTypeSpecifiedSoulRuins,
		QuestSubTypeSoulActive,
		QuestSubTypeClickSoulStrengthenButton,
		QuestSubTypeSoulStrengthenLevel,
		QuestSubTypeSoulSpecialEmbed,
		QuestSubTypeClickSoulUpgradeButton,
		QuestSubTypeSoulUpgradeLevel,
		QuestSubTypeXianFu,
		QuestSubTypePlayerLevel,
		QuestSubTypeJoinAlliance,
		QuestSubTypeApplyJoinAlliance,
		QuestSubTypeCollectItemTotal,
		QuestSubTypeCollectAllItem,
		QuestSubTypeEquipmentUpgradeLevel,
		QuestSubTypeEnterSpecialXianFu,
		QuestSubTypeEnterXianFu,
		QuestSubTypeEnterSoulRuins,
		QuestSubTypeEnterRealm,
		QuestSubTypeUpgradeSpecialXianFu,
		QuestSubTypeFinishSecretCard,
		QuestSubTypeRealmLevel,
		QuestSubTypeEmbedGem,
		QuestSubTypeZhuanSheng,
		QuestSubTypeFinishFirstCharge,
		QuestSubTypeActiveSkill,
		QuestSubTypeSystemX,
		QuestSubTypeEmbedQualityTwoGoldEquipNum,
		QuestSubTypeEmbedQualityThreeGoldEquipNum,
		QuestSubTypeEmbedQualityFourGoldEquipNum,
		QuestSubTypeSomeMonster,
		QuestSubTypeDayFirstCharge,
		QuestSubTypeDaBaoTimePast,
		QuestSubTypeAttandActivityPlay,
		QuestSubTypeAttandKillSpecialMonster,
		QuestSubTypeXianFuPersonal,
		QuestSubTypechallengeMaterialFuBen,
		QuestSubTypeBossKilled,
		QuestSubTypeBiologyTypeKilled,
		QuestSubTypeAdvancedOperation,
		QuestSubTypechallengeSpecialMaterialFuBen,
		QuestSubTypeEnterTower,
		QuestSubTypeUpgradeSysOperation,
		QuestSubTypeGoldEquipmentStrength,
		QuestSubTypeFinishDailyQuestNum,
		QuestSubTypeGoldEquipmentResolve,
		QuestSubTypeBaGuaMiJingLevel,
		QuestSubTypeSetTypeKilled,
		QuestSubTypeLingTongClick,
		QuestSubTypePassTeamCopy,
		QuestSubTypeLingTongActivateNum,
		QuestSubTypeGemTotalLevel,
		QuestSubTypeFuBenMonsterGroup,
		QuestSubTypeActivateSoulNum,
		QuestSubTypeAwakenSoulNum,
		QuestSubTypeChengZhanWinNum,
		QuestSubType3V3WinNum,
		QuestSubType3V3LianSheng,
		QuestSubTypeShenLongFuHua,
		QuestSubTypeActivateYiChuSuitNum,
		QuestSubTypeOneArenaOccupyTime,
		QuestSubTypeFeiShengLevel,
		QuestSubTypeGoldEquipmentTotalLevel,
		QuestSubTypeActivateYiChuSuitOne,
		QuestSubTypeActivateYiChuSuitTwo,
		QuestSubTypeActivateYiChuSuitThree,
		QuestSubTypeActivateYiChuSuitFour,
		QuestSubTypeBuyHuiYuan,
		QuestSubTypeBuyInvest,
		QuestSubTypeBuyEquipGift,
		QuestSubTypeFinishDailyAllianceQuestNum,
		QuestSubTypeAttendMajorFuBen,
		QuestSubTypeAttendFuQiFuBen,
		QuestSubTypeFinishMajorFuBen,
		QuestSubTypeFinishFuQiFuBen,
		QuestSubTypeAllianceChat,
		QuestSubTypeEquipBaoKuAttend,
		QuestSubTypeFeiShengSanGong,
		QuestSubTypeAttendTeamFuBen,
		QuestSubTypeFinishTeamFuBen,
		QuestSubTypeGuideReplica,
		QuestSubTypeWeekBuy,
		QuestSubTypeTradUpload,
		QuestSubTypeEquipBaoKuDuiHuan,
		QuestSubTypeChargeTimes,
		QuestSubTypeForceGrow,
		QuestSubTypeSendHongBao,
		QuestSubTypeTradItem,
		QuestSubTypeZuDuiFuBen,
		QuestSubTypeSkillXinFa,
		QuestSubTypeFinishJieYi,
		QuestSubTypeCostGold,
		QuestSubTypeArenaPVPJiFen,
		QuestSubTypeXiongDiWeiMing,
		QuestSubTypeXiongDiToken,
		QuestSubTypeMarryTimes,
		QuestSubTypeBiaoBaiTimes,
		QuestSubTypeCoupleToeknForce,
		QuestSubTypeKillPlayer,
		QuestSubTypeChuangShiCollect,
		QuestSubTypeChuangShiBiaoChe:
		return true
	}
	return false
}

var (
	questSubTypeMap = map[QuestSubType]string{
		QuestSubTypeDialog:                          "对话",
		QuestSubTypeKillMonster:                     "杀怪",
		QuestSubTypeCollect:                         "采集",
		QuestSubTypeLevel:                           "等级",
		QuestSubTypeHurtMonster:                     "对怪造成伤害",
		QuestSubTypeWorldChannel:                    "在世界频道发言X次",
		QuestSubTypeSendFlower:                      "送花X次",
		QuestSubTypeSyntheticTuMoToken:              "合成X次屠魔令",
		QuestSubTypeFinishTuMo:                      "完成X次屠魔任务",
		QuestSubTypeSoulRuins:                       "帝魂遗迹副本通关X次",
		QuestSubTypeSpecialXianFu:                   "进入指定秘境仙府x次数(仅限免费次数)",
		QuestSubTypeReserve:                         "保留",
		QuestSubTypeGamble:                          "进行X次赌石",
		QuestSubTypeRealm:                           "挑战成功X次天劫塔",
		QuestSubTypeDart:                            "进行X次押镖",
		QuestSubType1V1:                             "参加1V1竞技场X次",
		QuestSubType3V3:                             "参加3V3竞技场X次",
		QuestSubTypeMonsterHunt:                     "参加捉妖记活动X次",
		QuestSubTypeMoonLove:                        "参加月下情缘活动X次",
		QuestSubTypeAlliance:                        "参加仙盟活动(屠龙，城战，仙盟镖都算)X次",
		QuestSubTypeEmperor:                         "参加抢龙椅X次",
		QuestSubTypeDragonChess:                     "进行苍龙棋局抽奖X次",
		QuestSubTypeFourMysterious:                  "参加四神秘境X次",
		QuestSubTypeDartCar:                         "所在仙盟开启X次仙盟镖车",
		QuestSubTypeWorldBoss:                       "击杀X只世界BOSS",
		QuestSubTypeClickSkillUpgradeButton:         "一键升级或者升级职业技能X次",
		QuestSubTypeProfessionalSkillLevel:          "指定职业技能达到X级",
		QuestSubTypeProfessionalSkillTotalLevel:     "职业技能总等级为X级",
		QuestSubTypeClickEquipmentStrengthenButton:  "一键强化或者强化装备X次",
		QuestSubTypeEquipmentStrengthenLevel:        "指定部位装备强化等级为X级",
		QuestSubTypeEquipmentStrengthenTotalLevel:   "强化总等级为X级",
		QuestSubTypeClickEquipmentUpgradeButton:     "进阶装备X次",
		QuestSubTypeEquipmentUpgradeStar:            "指定装备星级达到X星",
		QuestSubTypeClickEquipmentUpgradeStarButton: "装备升星的次数",
		QuestSubTypeSpecifiedSoulRuins:              "通关指定帝魂副本X次",
		QuestSubTypeSoulActive:                      "激活指定帝魂",
		QuestSubTypeClickSoulStrengthenButton:       "强化帝魂X次",
		QuestSubTypeSoulStrengthenLevel:             "指定帝魂强化等级为X级",
		QuestSubTypeSoulSpecialEmbed:                "装备指定帝魂",
		QuestSubTypeClickSoulUpgradeButton:          "魂技升级X次",
		QuestSubTypeSoulUpgradeLevel:                "指定魂技达到X级",
		QuestSubTypeXianFu:                          "通关X次秘境仙府",
		QuestSubTypePlayerLevel:                     "玩家达到X级",
		QuestSubTypeJoinAlliance:                    "加入仙盟",
		QuestSubTypeApplyJoinAlliance:               "申请加入仙盟",
		QuestSubTypeCollectItemTotal:                "收集物品总数",
		QuestSubTypeCollectAllItem:                  "收集所有物品",
		QuestSubTypeEquipmentUpgradeLevel:           "装备升到X阶",
		QuestSubTypeEnterSpecialXianFu:              "进入指定X次秘境仙府",
		QuestSubTypeEnterXianFu:                     "进入X次秘境仙府",
		QuestSubTypeEnterSoulRuins:                  "进入X次帝陵遗迹副本",
		QuestSubTypeEnterRealm:                      "进入天劫塔X次",
		QuestSubTypeUpgradeSpecialXianFu:            "升级指定秘境仙府",
		QuestSubTypeFinishSecretCard:                "完成X次天机牌",
		QuestSubTypeRealmLevel:                      "通关天劫塔第x层",
		QuestSubTypeEmbedGem:                        "镶嵌宝石",
		QuestSubTypeZhuanSheng:                      "转生数达到X转",
		QuestSubTypeFinishFirstCharge:               "完成首冲",
		QuestSubTypeActiveSkill:                     "激活技能",
		QuestSubTypeSystemX:                         "系统升到X阶",
		QuestSubTypeEmbedQualityTwoGoldEquipNum:     "穿戴品质达到2及以上的元神金装x件()",
		QuestSubTypeEmbedQualityThreeGoldEquipNum:   "穿戴品质达到3及以上的元神金装x件()",
		QuestSubTypeEmbedQualityFourGoldEquipNum:    "穿戴品质达到4及以上的元神金装x件()",
		QuestSubTypeSomeMonster:                     "对某种类型的怪物造成伤害",
		QuestSubTypeDayFirstCharge:                  "完成每日首冲",
		QuestSubTypeDaBaoTimePast:                   "消耗打宝时间",
		QuestSubTypeAttandActivityPlay:              "参与活动玩法",
		QuestSubTypeAttandKillSpecialMonster:        "击杀个人boss和付费boss",
		QuestSubTypeXianFuPersonal:                  "仙府个人副本(进入)",
		QuestSubTypechallengeMaterialFuBen:          "挑战材料副本x次",
		QuestSubTypeBossKilled:                      "击杀世界boss,跨服boss",
		QuestSubTypeBiologyTypeKilled:               "击杀某种类型怪物",
		QuestSubTypeAdvancedOperation:               "升阶操作",
		QuestSubTypechallengeSpecialMaterialFuBen:   "进入指定的材料副本",
		QuestSubTypeEnterTower:                      "进入打宝塔",
		QuestSubTypeUpgradeSysOperation:             "升级系统点击X次",
		QuestSubTypeGoldEquipmentStrength:           "原元神金装装备强化x次",
		QuestSubTypeFinishDailyQuestNum:             "完成x个日环任务",
		QuestSubTypeBaGuaMiJingLevel:                "通关八卦秘境第x层",
		QuestSubTypeSetTypeKilled:                   "击杀指定的策划类型怪物",
		QuestSubTypeLingTongClick:                   "灵童点击次数",
		QuestSubTypePassTeamCopy:                    "通关组队副本",
		QuestSubTypeLingTongActivateNum:             "激活灵童数量",
		QuestSubTypeGemTotalLevel:                   "宝石总等级",
		QuestSubTypeFuBenMonsterGroup:               "指定副本达到X波怪",
		QuestSubTypeActivateSoulNum:                 "激活X个帝魂",
		QuestSubTypeAwakenSoulNum:                   "觉醒X个帝魂",
		QuestSubTypeChengZhanWinNum:                 "获得城战胜利X次",
		QuestSubType3V3WinNum:                       "获得3v3胜利X次",
		QuestSubType3V3LianSheng:                    "3v3连胜X场",
		QuestSubTypeShenLongFuHua:                   "孵化升龙到第X阶段",
		QuestSubTypeActivateYiChuSuitNum:            "激活指定衣橱时装X件属性",
		QuestSubTypeOneArenaOccupyTime:              "占领X阶灵池持续Y分钟",
		QuestSubTypeFeiShengLevel:                   "飞升等级达到",
		QuestSubTypeGoldEquipmentTotalLevel:         "元神金装强化总等级",
		QuestSubTypeActivateYiChuSuitOne:            "激活1条属性的衣橱套装X件",
		QuestSubTypeActivateYiChuSuitTwo:            "激活2条属性的衣橱套装X件",
		QuestSubTypeActivateYiChuSuitThree:          "激活3条属性的衣橱套装X件",
		QuestSubTypeActivateYiChuSuitFour:           "激活4条属性的衣橱套装X件",
		QuestSubTypeBuyHuiYuan:                      "成为至尊会员",
		QuestSubTypeBuyInvest:                       "购买任意等级投资",
		QuestSubTypeBuyEquipGift:                    "购买绝版首饰套装",
		QuestSubTypeFinishDailyAllianceQuestNum:     "完成X个仙盟日环任务",
		QuestSubTypeAttendMajorFuBen:                "参与双修副本X次",
		QuestSubTypeAttendFuQiFuBen:                 "参与夫妻副本X次",
		QuestSubTypeFinishMajorFuBen:                "通关双修副本X次",
		QuestSubTypeFinishFuQiFuBen:                 "通关夫妻副本X次",
		QuestSubTypeAllianceChat:                    "仙盟频道发言X次",
		QuestSubTypeEquipBaoKuAttend:                "装备宝库抽奖X次",
		QuestSubTypeFeiShengSanGong:                 "进行飞升散功x次",
		QuestSubTypeAttendTeamFuBen:                 "参与组队副本X次（不限制副本类型）",
		QuestSubTypeFinishTeamFuBen:                 "通关组队副本X次（不限制副本类型）",
		QuestSubTypeGuideReplica:                    "引导副本",
		QuestSubTypeWeekBuy:                         "购买周卡",
		QuestSubTypeTradUpload:                      "交易行上架物品",
		QuestSubTypeEquipBaoKuDuiHuan:               "装备宝库消耗积分兑换",
		QuestSubTypeChargeTimes:                     "完成充值（任意金额）X次",
		QuestSubTypeForceGrow:                       "提升战力X万（城战期间不算）",
		QuestSubTypeSendHongBao:                     "发送红包X次",
		QuestSubTypeTradItem:                        "交易行购买物品",
		QuestSubTypeZuDuiFuBen:                      "组队副本类型，只需进入便可通过",
		QuestSubTypeSkillXinFa:                      "激活x个技能心法",
		QuestSubTypeFinishJieYi:                     "完成结义x次",
		QuestSubTypeCostGold:                        "消费元宝",
		QuestSubTypeArenaPVPJiFen:                   "比武大会海选积分累计获得x",
		QuestSubTypeXiongDiWeiMing:                  "威名等级达到x级",
		QuestSubTypeXiongDiToken:                    "兄弟信物等级达到x级",
		QuestSubTypeMarryTimes:                      "完成结婚的次数",
		QuestSubTypeBiaoBaiTimes:                    "赠送表白道具的次数",
		QuestSubTypeCoupleToeknForce:                "夫妻信物总战力达到x",
		QuestSubTypeKillPlayer:                      "击杀玩家X次",
	}
)

func (qt QuestSubType) String() string {
	return questSubTypeMap[qt]
}

type QuestSpecialType int32

const (
	QuestSpecialTypeMarried QuestSpecialType = 1 << iota
	QuestSpecialTypeBangHui
)

var (
	questSpecialTypeMap = map[QuestSpecialType]string{
		QuestSpecialTypeMarried: "结婚",
		QuestSpecialTypeBangHui: "帮会",
	}
)

func (qst QuestSpecialType) String() string {
	return questSpecialTypeMap[qst]
}

//任务状态
type QuestState int32

const (
	//初始化
	QuestStateInit QuestState = iota
	//激活
	QuestStateActive
	//接受
	QuestStateAccept
	//完成
	QuestStateFinish
	//交付
	QuestStateCommit
	//放弃
	QuestStateDiscard
)

var questStateMap = map[QuestState]string{
	QuestStateInit:    "初始化",
	QuestStateActive:  "激活",
	QuestStateAccept:  "接受",
	QuestStateFinish:  "完成",
	QuestStateCommit:  "交付",
	QuestStateDiscard: "放弃",
}

func (qt QuestState) String() string {
	return questStateMap[qt]
}

//任务品质
type QuestLevelType int32

const (
	//绿
	QuestLevelTypeTuMoGreen QuestLevelType = iota
	//蓝
	QuestLevelTypeTuMoBlue
	//紫
	QuestLevelTypeTuMoPurple
	//橙橙
	QuestLevelTypeTuMoOrange
)

func (qtmlt QuestLevelType) Valid() bool {
	switch qtmlt {
	case QuestLevelTypeTuMoGreen,
		QuestLevelTypeTuMoBlue,
		QuestLevelTypeTuMoPurple,
		QuestLevelTypeTuMoOrange:
		return true
	}
	return false
}

var QuestTuMoMap = map[QuestLevelType]string{
	QuestLevelTypeTuMoGreen:  "绿",
	QuestLevelTypeTuMoBlue:   "蓝",
	QuestLevelTypeTuMoPurple: "紫",
	QuestLevelTypeTuMoOrange: "橙",
}

func (qtmlt QuestLevelType) String() string {
	return QuestTuMoMap[qtmlt]
}

var questTuMoItemSubMap = map[QuestLevelType]itemtypes.ItemTuMoLingSubType{
	QuestLevelTypeTuMoGreen:  itemtypes.ItemTuMoLingSubTypeGreen,
	QuestLevelTypeTuMoBlue:   itemtypes.ItemTuMoLingSubTypeBlue,
	QuestLevelTypeTuMoPurple: itemtypes.ItemTuMoLingSubTypePurple,
	QuestLevelTypeTuMoOrange: itemtypes.ItemTuMoLingSubTypeOrange,
}

func (qtmlt QuestLevelType) ItemTumoSubType() itemtypes.ItemTuMoLingSubType {
	return questTuMoItemSubMap[qtmlt]
}

type QuestDailyType int32

const (
	QuestDailyTypeMin QuestDailyType = 1
	QuestDailyTypeMax QuestDailyType = 20
)

func (t QuestDailyType) Valid() bool {
	if t >= QuestDailyTypeMin && t <= QuestDailyTypeMax {
		return true
	}
	return false
}

type QuestDailyTag int32

const (
	QuestDailyTagPerson QuestDailyTag = iota + 1
	QuestDailyTagAlliance
)

func (t QuestDailyTag) Valid() bool {
	switch t {
	case QuestDailyTagAlliance,
		QuestDailyTagPerson:
		return true
	}
	return false
}

const (
	QuestDailyTagMin = QuestDailyTagPerson
	QuestDailyTagMax = QuestDailyTagAlliance
)

var dailyTagMap = map[QuestType]QuestDailyTag{
	QuestTypeDaily:         QuestDailyTagPerson,
	QuestTypeDailyAlliance: QuestDailyTagAlliance,
}

var dailyTagReverseMap = map[QuestDailyTag]QuestType{
	QuestDailyTagPerson:   QuestTypeDaily,
	QuestDailyTagAlliance: QuestTypeDailyAlliance,
}

var dailyTagFuncopenMap = map[QuestDailyTag]funcopentypes.FuncOpenType{
	QuestDailyTagPerson:   funcopentypes.FuncOpenTypeDailyQuest,
	QuestDailyTagAlliance: funcopentypes.FuncOpenTypeAllianceDaily,
}

var dailyTagStringMap = map[QuestDailyTag]string{
	QuestDailyTagPerson:   "日环",
	QuestDailyTagAlliance: "仙盟日常",
}

func (t QuestDailyTag) GetQuestType() QuestType {
	return dailyTagReverseMap[t]
}

func (t QuestDailyTag) GetFuncOpen() funcopentypes.FuncOpenType {
	return dailyTagFuncopenMap[t]
}

func (t QuestDailyTag) String() string {
	return dailyTagStringMap[t]
}
