package log

type SilverLogReason int32

const (
	SilverLogReasonGM SilverLogReason = iota + 1
	SilverLogReasonEquipSlotStrengthUpgrade
	SilverLogReasonEquipSlotStrengthStar
	SilverLogReasonQuestAccept
	SilverLogReasonQuestCommit
	SilverLogReasonQuestReward
	SilverLogReasonItemSell
	SilverLogReasonShopBuyItem
	SilverLogReasonMountAdvanced
	SilverLogReasonWingAdvanced //10
	SilverLogReasonBodyShieldAdvanced
	SilverLogReasonWeapUpstar
	SilverLogReasonEquipSlotStrengthUpgradeAuto
	SilverLogReasonSkillUpgrade
	SilverLogReasonSkillUpgradeAll
	SilverLogReasonXinFaActive
	SilverLogReasonXinFaUpgrade
	SilverLogReasonGemMineActive
	SilverLogReasonGemGamble
	SilverLogReasonGemGambleDrop //20
	SilverLogReasonSynthesis
	SilverLogReasonGiftBuy
	SilverLogReasonEmailAttachment
	SilverLogReasonSoulRuinsRewChapter
	SilverLogReasonSoulRuinsSweep
	SilverLogReasonSoulRuinsFinishAll
	SilverLogReasonSoulRuinsSweepDrop
	SilverLogReasonSoulRuinsFinishAllDrop
	SilverLogReasonSoulRuins
	SilverLogReasonSoulRuinsRobber //30
	SilverLogReasonXianfuUpgrade
	SilverLogReasonXianfuSaodangRewards
	SilverLogReasonRealmToKillRew
	SilverLogReasonMonsterKilled
	SilverLogReasonEmperorStorageGet
	SilverLogReasonEmperorWorship
	SilverLogReasonMoonloveRankRew
	SilverLogReasonMoonloveTickRew
	SilverLogReasonAllianceDonate
	SilverLogReasonAllianceCallGuard //40
	SilverLogReasonAllianceSceneDoorReward
	SilverLogReasonAllianceSkillUpgrade
	SilverLogReasonSecretCardStarRew
	SilverLogReasonSecretCardFinishAllStarRew
	SilverLogReasonSecretCardFinishAll
	SilverLogReasonSecretCardFinish
	SilverLogReasonFoundResourceUse
	SilverLogReasonFoundResourceBatchUse
	SilverLogReasonFoundResource
	SilverLogReasonTransportationCost //50
	SilverLogReasonTransportationRobRew
	SilverLogReasonTuMoFinishAll
	SilverLogResonFourGodOpenBox
	SilverLogReasonTransportationFailRew
	SilverLogReasonTransportationFinishRew
	SilverLogReasonFeatherAdvanced
	SilverLogReasonShieldAdvanced
	SilverLogReasonShenfaAdvanced
	SilverLogReasonLingyuAdvanced
	SilverLogReasonMarryRingFeedCost //60
	SilverLogReasonMarryTreeFeedCost
	SilverLogReasonMarryWedGift
	SilverLogReasonMarryWeddingGrade
	SilverLogReasonXianfuFinishAllRew
	SilverLogReasonBoxUse
	SilverLogReasonBoxGet
	SilverLogReasonSoulStrengthen
	SilverLogReasonChessUse
	SilverLogReasonOneArenaSellKun
	SilverLogReasonItemSellBatch //70
	SilverLogReasonChangeScene
	SilverLogReasonChessGet
	SilverLogReasonFirstCharge
	SilverLogReasonWelfareRew
	SilverLogReasonInvestRew
	SilverLogReasonFeedbackRew
	SilverLogReasonArena
	SilverLogReasonOpenActivityRew
	SilverLogReasonAnqiAdvanced
	SilverLogReasonGiftCode //80
	SilverLogReasonTuLongKillBossRew
	SilverLogReasonCollectRew
	SilverLogReasonTuMoImmediateFinish
	SilverLogReasonFireworks
	SilverLogReasonBuyXianHua
	SilverLogReasonVipGiftRew
	SilverLogReasonReliveAutoBuy
	SilverLogReasonVipFreeGiftRew
	SilverLogReasonEmperorReward
	SilverLogReasonMassacreAdvanced //90
	SilverLogReasonTianShuRew
	SilverLogReasonGoldEquipAutoBuy
	SilverLogReasonSystemSkillActive
	SilverLogReasonSystemSkillUpgrade
	SilverLogReasonFaBaoAdvanced
	SilverLogReasonAdditionSysLevel
	SilverLogReasonMaterialRew
	SilverLogReasonXianTiAdvanced
	SilverLogReasonXueDunUpgrade
	SilverLogReasonLivenessStarRew //100
	SilverLogReasonFoeFeedbackCost
	SilverLogReasonHuiYuanRew
	SilverLogReasonAddFriendRew
	SilverLogReasonUnrealBossDrop
	SilverLogResonDailyQuestReward
	SilverLogReasonDailyQuestFinishAll
	SilverLogReasonBaGuaToKillRew
	SilverLogReasonOutlandBossDrop
	SilverLogReasonSongBuTingRew
	SilverLogReasonInventoryResourceUse
	SilverLogResonTeamCopyReward
	SilverLogReasonQuizAnswer
	SilverLogReasonDianXingAdvanced
	SilverLogReasonDianXingJieFengAdvanced
	SilverLogReasonTianMoAdvanced
	SilverLogReasonShiHunFanAdvanced
	SilverLogReasonActivityTickRew
	SilverLogReasonLingTongAdvanced
	SilverLogReasonDrewUse
	SilverLogReasonShenMoWarRankReward
	SilverLogReasonKaiFuMuBiaoReward
	SilverLogReasonHongBaoSnatch
	SilverLogReasonChatAward
	SilverLogReasonAdditionSysHuaLing
	SilverLogReasonGroupCollectRew
	SilverLogReasonEquipBaoKuUse
	SilverLogReasonEquipBaoKuGet
	SilverLogReasonEquipBaoKuLuckyBoxGet
	SilverLogReasonMingGeSynthesisCost
	SilverLogReasonMingGeRefinedCost
	SilverLogReasonMingGeMingLiBaptizeCost
	SilverLogReasonZhenFaShengJiCost
	SilverLogReasonZhenFaXianHuoShengJiCost
	SilverLogReasonShenQiDebrisUpCost
	SilverLogReasonShenQiSmeltUpCost
	SilverLogReasonFuncopenRew
	SilverLogReasonBabyLearnBuyUse
	SilverLogReasonMarryDingQingJiHuo
	SilverLogReasonQuestQiYuRew
	SilverLogReasonXianTaoCommit
	SilverLogReasonHouseSell
	SilverLogReasonHouseRent
	SilverLogReasonMarryProposal
	SilverLogReasonAdditionSysShenZhuCost
	SilverLogReasonAdditionSysTongLingCost
	SilverLogReasonGuideRew
	SilverLogReasonShenYuRoundRew
	SilverLogReasonBuyWeekRew
	SilverLogReasonWeekDayRew
	SilverLogReasonMajorSaoDangGet
	SilverLogReasonQiXueAdvanced
	SilverLogReasonArenaRankReward
	SilverLogReasonMaterialSaodangGet
	SilverLogReasonJieYiTokenChange
	SilverLogReasonJieYiDaoJuChange
	SilverLogReasonArenapvpRew
	SilverLogReasonJieYiTokenUpLev
	SilverLogReasonCastingSpiritUplevel
	SilverLogReasonRingTypeChangeCost
	SilverLogReasonRewPools
	SilverLogReasonXianZunCardActiviteRew
	SilverLogReasonXianZunCardReceiveRew
	SilverLogReasonXueChiAutoBuy
	SilverLogReasonRewPoolsDrew
	SilverLogReasonLingShouUplevel
	SilverLogReasonLingwenUplevel
	SilverLogReasonLingShouUpRank
	SilverLogReasonLingLinglian
	SilverLogReasonRingAdvance
	SilverLogReasonRingFuse
	SilverLogReasonRingBaoKuUse
	SilverLogReasonRingBaoKuGet
	SilverLogReasonRingBaoKuLuckyBoxGet
)

func (slr SilverLogReason) Reason() int32 {
	return int32(slr)
}

var (
	silverLogReasonMap = map[SilverLogReason]string{
		SilverLogReasonGM:                           "gm修改",
		SilverLogReasonEquipSlotStrengthUpgrade:     "装备槽强化升级",
		SilverLogReasonEquipSlotStrengthStar:        "装备槽强化升星",
		SilverLogReasonQuestAccept:                  "任务接取消耗,任务id:%d",
		SilverLogReasonQuestCommit:                  "任务完成消耗,任务id:%d",
		SilverLogReasonQuestReward:                  "任务完成奖励,任务id:%d",
		SilverLogReasonItemSell:                     "物品出售[%d],个数[%d]",
		SilverLogReasonItemSellBatch:                "物品批量出售，%v",
		SilverLogReasonShopBuyItem:                  "商城购买shopId:%d,,商品名字:%s,购买个数:%d",
		SilverLogReasonMountAdvanced:                "坐骑进阶",
		SilverLogReasonWingAdvanced:                 "战翼进阶",
		SilverLogReasonBodyShieldAdvanced:           "护体盾进阶",
		SilverLogReasonWeapUpstar:                   "兵魂升星购买道具",
		SilverLogReasonEquipSlotStrengthUpgradeAuto: "装备槽自动升级,位置:%s,等级:%d",
		SilverLogReasonSkillUpgrade:                 "技能升级,技能id:%d,技能等级:%d",
		SilverLogReasonSkillUpgradeAll:              "技能全部升级",
		SilverLogReasonXinFaActive:                  "心法激活",
		SilverLogReasonXinFaUpgrade:                 "心法升级",
		SilverLogReasonGemMineActive:                "矿工激活",
		SilverLogReasonGemGamble:                    "赌石银两消耗",
		SilverLogReasonGemGambleDrop:                "赌石掉落银两",
		SilverLogReasonSynthesis:                    "合成物品，合成id:%d,数量:%d",
		SilverLogReasonGiftBuy:                      "礼物购买,礼物id:%d,数量:%d",
		SilverLogReasonEmailAttachment:              "邮件附件奖励",
		SilverLogReasonSoulRuinsRewChapter:          "帝陵遗迹章节奖励领取chapter:%d typ:%d",
		SilverLogReasonSoulRuinsSweep:               "帝陵遗迹扫荡,扫荡关卡chapter:%d typ:%d level:%d,扫荡次数:%d",
		SilverLogReasonSoulRuinsFinishAll:           "帝陵遗迹一键完成chapter:%d typ:%d level:%d,次数:%d",
		SilverLogReasonSoulRuinsSweepDrop:           "帝陵遗迹扫荡掉落,扫荡关卡chapter:%d typ:%d level:%d,扫荡次数:%d",
		SilverLogReasonSoulRuinsFinishAllDrop:       "帝陵遗迹一键完成掉落chapter:%d typ:%d level:%d,次数:%d",
		SilverLogReasonSoulRuins:                    "帝陵遗迹挑战成功奖励chapter:%d typ:%d level:%d",
		SilverLogReasonSoulRuinsRobber:              "帝陵遗迹马贼收取银两:chapter:%d typ:%d level:%d",
		SilverLogReasonXianfuUpgrade:                "秘境仙府升级,仙府id:%d,类型:%d",
		SilverLogReasonXianfuSaodangRewards:         "秘境仙府扫荡奖励,仙府id:%d,类型:%d",
		SilverLogReasonRealmToKillRew:               "天劫塔前往击杀奖励,层数:%d",
		SilverLogReasonMonsterKilled:                "杀怪奖励",
		SilverLogReasonEmperorStorageGet:            "帝王领取库存",
		SilverLogReasonEmperorWorship:               "膜拜帝王领取奖励",
		SilverLogReasonMoonloveRankRew:              "月下情缘排行榜奖励，名次:%d,类型:%d",
		SilverLogReasonMoonloveTickRew:              "月下情缘定时奖励",
		SilverLogReasonAllianceDonate:               "仙盟捐献,仙盟id:%d",
		SilverLogReasonAllianceCallGuard:            "城战召唤守卫,仙盟id:%d,地图id:%d,守卫id:%d",
		SilverLogReasonAllianceSceneDoorReward:      "城战城门奖励,仙盟id:%d,门:%d",
		SilverLogReasonAllianceSkillUpgrade:         "仙术升级消耗",
		SilverLogReasonSecretCardStarRew:            "天机牌领取运势箱奖励",
		SilverLogReasonSecretCardFinishAllStarRew:   "天机牌一键完成剩余运势箱奖励",
		SilverLogReasonSecretCardFinishAll:          "天机牌一键完成奖励",
		SilverLogReasonSecretCardFinish:             "天机牌任务完成",
		SilverLogReasonFoundResourceUse:             "资源找回消耗,资源id:%d,次数:%d",
		SilverLogReasonFoundResourceBatchUse:        "资源找回批量消耗，资源map：%v",
		SilverLogReasonFoundResource:                "资源找回奖励,资源id:%d,次数:%,波数：%d",
		SilverLogReasonTransportationCost:           "押镖消耗,镖车类型：%s",
		SilverLogReasonTuMoFinishAll:                "屠魔任务一键完成,次数:%d",
		SilverLogReasonTransportationRobRew:         "劫镖奖励，镖车类型：%s",
		SilverLogResonFourGodOpenBox:                "四神遗迹开启宝箱奖励,钥匙数:%d",
		SilverLogReasonTransportationFailRew:        "押镖失败奖励，镖车类型：%s",
		SilverLogReasonTransportationFinishRew:      "押镖成功奖励，镖车类型：%s",
		SilverLogReasonFeatherAdvanced:              "护体仙羽进阶",
		SilverLogReasonShieldAdvanced:               "神盾尖刺进阶",
		SilverLogReasonShenfaAdvanced:               "身法进阶",
		SilverLogReasonLingyuAdvanced:               "领域进阶",
		SilverLogReasonMarryRingFeedCost:            "婚戒培养消耗",
		SilverLogReasonMarryTreeFeedCost:            "爱情树培养消耗",
		SilverLogReasonMarryWedGift:                 "婚宴赠送贺礼",
		SilverLogReasonMarryWeddingGrade:            "婚宴预定,预定婚宴档次:%d,婚车档次:%d,喜糖档次:%d",
		SilverLogReasonXianfuFinishAllRew:           "秘境仙府一键完成奖励,仙府id:%d,类型:%d",
		SilverLogReasonBoxUse:                       "宝箱开启消耗",
		SilverLogReasonBoxGet:                       "宝箱开启奖励",
		SilverLogReasonSoulStrengthen:               "帝魂强化",
		SilverLogReasonChessUse:                     "苍龙棋局消耗，类型：%s",
		SilverLogReasonOneArenaSellKun:              "鲲一键卖出获得",
		SilverLogReasonChangeScene:                  "切换世界地图消耗,地图:%d",
		SilverLogReasonChessGet:                     "苍龙棋局奖励,类型：%s",
		SilverLogReasonFirstCharge:                  "首冲奖励",
		SilverLogReasonOpenActivityRew:              "开服活动奖励，类型:%d,子类型:%d",
		SilverLogReasonArena:                        "3v3竞技场,层数:%d",
		SilverLogReasonAnqiAdvanced:                 "暗器进阶",
		SilverLogReasonGiftCode:                     "礼包兑换奖励",
		SilverLogReasonTuLongKillBossRew:            "跨服屠龙击杀boss奖励",
		SilverLogReasonCollectRew:                   "采集物奖励",
		SilverLogReasonTuMoImmediateFinish:          "屠魔任务直接完成",
		SilverLogReasonFireworks:                    "放烟花,物品id:%d,数量:%d",
		SilverLogReasonBuyXianHua:                   "购买鲜花",
		SilverLogReasonVipGiftRew:                   "vip礼包奖励",
		SilverLogReasonReliveAutoBuy:                "复活丹自动购买:%d",
		SilverLogReasonVipFreeGiftRew:               "vip免费礼包奖励",
		SilverLogReasonEmperorReward:                "帝王奖励",
		SilverLogReasonMassacreAdvanced:             "戮仙刃消耗，当前%d阶%d星",
		SilverLogReasonTianShuRew:                   "天书每日奖励,typ:%s",
		SilverLogReasonGoldEquipAutoBuy:             "金装开光购买道具",
		SilverLogReasonSystemSkillActive:            "系统技能激活消耗,模板id:%d",
		SilverLogReasonSystemSkillUpgrade:           "系统技能升级消耗,模板id:%d",
		SilverLogReasonFaBaoAdvanced:                "法宝进阶",
		SilverLogReasonAdditionSysLevel:             "%s系統消耗，当前%d等级%d进度%d次数",
		SilverLogReasonMaterialRew:                  "材料副本奖励，typ:%s",
		SilverLogReasonXianTiAdvanced:               "仙体进阶",
		SilverLogReasonXueDunUpgrade:                "血盾升阶消耗",
		SilverLogReasonLivenessStarRew:              "活跃度领取宝箱奖励",
		SilverLogReasonFoeFeedbackCost:              "仇人推送反馈消耗",
		SilverLogReasonHuiYuanRew:                   "会员每日奖励,类型:%d",
		SilverLogReasonAddFriendRew:                 "添加好友奖励",
		SilverLogReasonUnrealBossDrop:               "幻境boss掉落",
		SilverLogResonDailyQuestReward:              "完成日环任务奖励，类型：%s",
		SilverLogReasonDailyQuestFinishAll:          "一键完成%s任务奖励",
		SilverLogReasonBaGuaToKillRew:               "八卦秘境前往击杀奖励,层数:%d",
		SilverLogReasonOutlandBossDrop:              "外域boss掉落",
		SilverLogReasonSongBuTingRew:                "领取元宝送不停奖励",
		SilverLogReasonInventoryResourceUse:         "背包资源使用获得",
		SilverLogResonTeamCopyReward:                "组队副本奖励,副本类型:%d",
		SilverLogReasonQuizAnswer:                   "仙尊问答奖励,玩家等级:%d ,答题结果:%d",
		SilverLogReasonDianXingAdvanced:             "点星系统升级，当前%d星谱%d星%d进度%d次数",
		SilverLogReasonDianXingJieFengAdvanced:      "点星解封升级，当前%d阶级%d进度%d次数",
		SilverLogReasonTianMoAdvanced:               "天魔体进阶消耗",
		SilverLogReasonShiHunFanAdvanced:            "噬魂幡进阶消耗，当前%d阶",
		SilverLogReasonActivityTickRew:              "活动定时奖励,活动类型:%d",
		SilverLogReasonLingTongAdvanced:             "灵童%s进阶消耗",
		SilverLogReasonDrewUse:                      "抽奖活动消耗,类型:%d",
		SilverLogReasonShenMoWarRankReward:          "领取神魔战场周排名奖励,仙盟排名:%d",
		SilverLogReasonKaiFuMuBiaoReward:            "开服目标,领取组奖励:%d",
		SilverLogReasonHongBaoSnatch:                "抢送红包获得",
		SilverLogReasonChatAward:                    "发言奖励",
		SilverLogReasonAdditionSysHuaLing:           "%s系統化灵食用消耗，当前%d等级%d进度%d次数",
		SilverLogReasonGroupCollectRew:              "摸金卡牌收集奖励，卡牌组合类型：%d",
		SilverLogReasonEquipBaoKuUse:                "装备宝库消耗，等级：%d，转数：%d",
		SilverLogReasonEquipBaoKuGet:                "装备宝库奖励，等级：%d，转数：%d",
		SilverLogReasonEquipBaoKuLuckyBoxGet:        "装备宝库幸运宝箱奖励",
		SilverLogReasonMingGeSynthesisCost:          "命格合成消耗",
		SilverLogReasonMingGeRefinedCost:            "命盘祭炼消耗",
		SilverLogReasonMingGeMingLiBaptizeCost:      "命格命理洗炼消耗",
		SilverLogReasonZhenFaShengJiCost:            "阵法升级消耗",
		SilverLogReasonZhenFaXianHuoShengJiCost:     "阵法仙火升级消耗",
		SilverLogReasonShenQiDebrisUpCost:           "神器碎片升级消耗,类型：%s，部位：%s",
		SilverLogReasonShenQiSmeltUpCost:            "神器淬炼升级消耗,类型：%s，部位：%s,",
		SilverLogReasonFuncopenRew:                  "功能开启奖励,功能id:%d",
		SilverLogReasonBabyLearnBuyUse:              "宝宝读书物品购买消耗,物品id：%d，数量：%d，宝宝id:%d,等级:%d,品质：%d",
		SilverLogReasonMarryDingQingJiHuo:           "定情信物激活消耗",
		SilverLogReasonQuestQiYuRew:                 "奇遇任务奖励，qiyuId:%d",
		SilverLogReasonXianTaoCommit:                "提交仙桃奖励",
		SilverLogReasonHouseSell:                    "房子出售奖励，序号：%d，类型：%d，等级：%d",
		SilverLogReasonHouseRent:                    "房子租金奖励，序号：%d，类型：%d，等级：%d",
		SilverLogReasonMarryProposal:                "求婚消耗",
		SilverLogReasonAdditionSysShenZhuCost:       "%s系统%s部位神铸消耗，当前%d等级%d进度%d次数",
		SilverLogReasonAdditionSysTongLingCost:      "%s系统通灵消耗，当前%d等级%d进度%d次数",
		SilverLogReasonGuideRew:                     "引导副本奖励，类型：%s",
		SilverLogReasonShenYuRoundRew:               "神域之战奖励，参赛轮：%d",
		SilverLogReasonBuyWeekRew:                   "购买周卡奖励，类型：%s",
		SilverLogReasonWeekDayRew:                   "周卡每日奖励，类型：%s",
		SilverLogReasonMajorSaoDangGet:              "夫妻副本扫荡%d次奖励,type:%s,fubenId:%d",
		SilverLogReasonQiXueAdvanced:                "泣血枪消耗，当前%d阶%d星",
		SilverLogReasonArenaRankReward:              "领取3v3周排名奖励,排名:%d",
		SilverLogReasonMaterialSaodangGet:           "材料副本扫荡%d次奖励，typ:%s",
		SilverLogReasonJieYiTokenChange:             "结义信物消耗银两，旧类型：%s,新类型：%s, 消耗方式：%s",
		SilverLogReasonJieYiDaoJuChange:             "结义道具消耗银两，旧类型：%s,新类型：%s, 消耗方式：%s",
		SilverLogReasonArenapvpRew:                  "比武大会奖励，获胜：%v， 类型：%s",
		SilverLogReasonJieYiTokenUpLev:              "结义信物升级自动购买消耗银两",
		SilverLogReasonCastingSpiritUplevel:         "铸灵自动升级消耗银两",
		SilverLogReasonRingTypeChangeCost:           "结婚戒指替换消耗，旧类型：%s,新类型：%s",
		SilverLogReasonRewPools:                     "奖池抽奖，位置[%d]",
		SilverLogReasonXianZunCardActiviteRew:       "仙尊特权卡激活奖励，类型：%s",
		SilverLogReasonXianZunCardReceiveRew:        "仙尊特权卡每日领取奖励，类型：%s",
		SilverLogReasonXueChiAutoBuy:                "血池自动购买生命瓶:物品id:%d,数量:%d",
		SilverLogReasonRewPoolsDrew:                 "仙人指路-奖池抽奖",
		SilverLogReasonLingShouUplevel:              "上古之灵灵兽升级，灵兽类型[%s]",
		SilverLogReasonLingwenUplevel:               "上古之灵灵纹升级，灵兽类型[%s], 灵纹类型[%s]",
		SilverLogReasonLingShouUpRank:               "上古之灵灵兽进阶，灵兽类型[%s]",
		SilverLogReasonLingLinglian:                 "上古之灵灵兽灵炼，灵兽类型[%s]",
		SilverLogReasonRingAdvance:                  "特戒进阶消耗，特戒类型：%s",
		SilverLogReasonRingFuse:                     "特戒融合消耗，特戒类型：%s",
		SilverLogReasonRingBaoKuUse:                 "特戒宝库探索消耗",
		SilverLogReasonRingBaoKuGet:                 "特戒宝库探索获得",
		SilverLogReasonRingBaoKuLuckyBoxGet:        "特戒宝库幸运宝箱奖励",
	}
)

func (slr SilverLogReason) String() string {
	return silverLogReasonMap[slr]
}
