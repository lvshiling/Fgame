package log

type GoldLogReason int32

const (
	GoldLogReasonGM GoldLogReason = iota + 1
	GoldLogReasonDanAccelerate
	GoldLogReasonQuestAccept
	GoldLogReasonQuestCommit
	GoldLogReasonQuestReward
	GoldLogReasonBuySlots
	GoldLogReasonShopBuyItem
	GoldLogReasonMountAdvanced
	GoldLogReasonWingAdvanced
	GoldLogReasonBodyShieldAdvanced
	GoldLogReasonWeapUpstar //11
	GoldLogReasonEquipSlotStrengthUpgradeAuto
	GoldLogReasonGemMineActive
	GoldLogReasonGemGamble
	GoldLogReasonGemGambleDrop
	GoldLogReasonSynthesis
	GoldLogReasonGiftBuy
	GoldLogReasonTuMoBuyNum
	GoldLogReasonTuMoFinishAll
	GoldLogReasonEmailAttachment
	GoldLogReasonSoulRuinsRewChapter //21
	GoldLogReasonSoulRuinsSweep
	GoldLogReasonSoulRuinsFinishAll
	GoldLogReasonSoulRuinsSweepDrop
	GoldLogReasonSoulRuinsFinishAllDrop
	GoldLogReasonSoulRuins
	GoldLogReasonSoulRuinsBuyNum
	GoldLogReasonSoulRuinsFinishAllCost
	GoldLogReasonXianfuUpgrade
	GoldLogReasonXianfuAccelerate
	GoldLogReasonXianfuSaodang //31
	GoldLogReasonXianfuSaodangRewards
	GoldLogReasonRealmToKillRew
	GoldLogReasonMonsterKilled
	GoldLogReasonEmperorWorship
	GoldLogReasonRobEmperor
	GoldLogReasonMoonloveGivePresents
	GoldLogReasonMoonloveRankRew
	GoldLogReasonMoonloveTickRew
	GoldLogReasonAllianceDonate
	GoldLogReasonAllianceCallGuard //41
	GoldLogReasonAllianceSceneDoorReward
	GoldLogReasonSecretCardStarRew
	GoldLogReasonSecretCardFinishAllStarRew
	GoldLogReasonSecretCardFinishAll
	GoldLogReasonSecretCardFinishAllCost
	GoldLogReasonSecretCardImmediateFinishCost
	GoldLogReasonSecretCardFinish
	GoldLogReasonFoundResource
	GoldLogReasonFoundResourceCost
	GoldLogReasonFoundResourceBatchCost //51
	GoldLogReasonTransportationCost
	GoldLogReasonTransportationRobRew
	GoldLogReasonTransportationFailRew
	GoldLogReasonTransportationFinishRew
	GoldLogReasonFourGodOpenBox
	GoldLogReasonFeatherAdvanced
	GoldLogReasonShieldAdvanced
	GoldLogReasonShenfaAdvanced
	GoldLogReasonLingyuAdvanced
	GoldLogReasonMarryRingFeedCost //61
	GoldLogReasonMarryTreeFeedCost
	GoldLogReasonMarryWeddingGrade
	GoldLogReasonXianfuFinishAllUse
	GoldLogReasonXianfuFinishAllRew
	GoldLogReasonBoxUse
	GoldLogReasonBoxGet
	GoldLogReasonChessUse
	GoldLogReasonBuyDepotSlots
	GoldLogReasonOneArenaSellKun
	GoldLogReasonChessGet //71
	GoldLogReasonFirstCharge
	GoldLogReasonBuyInvest
	GoldLogReasonBuyHuiYuan
	GoldLogReasonHuiYuanRew
	GoldLogReasonOpenActivityRew
	GoldLogReasonAnqiAdvanced
	GoldLogReasonGiftCode
	GoldLogReasonCharge
	GoldLogReasonFirstLevelCharge
	GoldLogReasonXueChiAutoBuy //81
	GoldLogReasonTuLongKillBossRew
	GoldLogReasonDrewUse
	GoldLogReasonBuyDiscountUse
	GoldLogReasonReliveAutoBuy
	GoldLogReasonCollectRew
	GoldLogReasonBuyVipGift
	GoldLogReasonTuMoImmediateFinishCost
	GoldLogReasonTuMoImmediateFinish
	GoldLogReasonFireworks
	GoldLogReasonBuyXianHua //91
	GoldLogReasonEmperorReward
	GoldLogReasonPrivilegeCharge
	GoldLogReasonPrivilegeItem
	GoldLogReasonMassacreAdvanced
	GoldLogReasonTianShuActivate
	GoldLogReasonTianShuReceive
	GoldLogReasonGoldEquipAutoBuy
	GoldLogReasonBuyJunior
	GoldLogReasonBuySenior
	GoldLogReasonSystemSkillActive
	GoldLogReasonSystemSkillUpgrade
	GoldLogReasonChargeDrew
	GoldLogReasonFaBaoAdvanced
	GoldLogReasonAdditionSysLevel
	GoldLogReasonMaterialRew
	GoldLogReasonXianTiAdvanced
	GoldLogReasonLivenessStarRew
	GoldLogReasonUnrealBossBuyPilao
	GoldLogReasonFoeFeedbackBuyProtect
	GoldLogReasonUnrealBossDrop
	GoldLogReasonDailyQuestCommitDouble
	GoldLogReasonDailyQuestReward
	GoldLogReasonDailyQuestImmediateFinishAll
	GoldLogReasonDailyQuestFinishAll
	GoldLogReasonBaGuaToKillRew
	GoldLogReasonOutlandBossDrop
	GoldLogReasonLaBaUse
	GoldLogReasonLaBaGet
	GoldLogReasonSongBuTingRew
	GoldLogReasonInventoryResourceUse
	GoldLogReasonTeamCopyReward
	GoldLogReasonQuizAnswer
	GoldLogReasonDianXingAdvanced
	GoldLogReasonDianXingJieFengAdvanced
	GoldLogReasonMaterialSaoDangUse
	GoldLogReasonTianMoAdvanced
	GoldLogReasonShiHunFanAdvanced
	GoldLogReasonActivityTickRew
	GoldLogReasonLingTongAdvanced
	GoldLogReasonActivityCost
	GoldLogReasonOpenActivityChargeReturn
	GoldLogReasonLingyuActivate
	GoldLogReasonFeiShengResetQn
	GoldLogReasonShenMoWarRankReward
	GoldLogReasonHongBaoSnatch
	GoldLogReasonChatAward
	GoldLogReasonAdditionSysHuaLing
	GoldLogReasonGroupCollectRew
	GoldLogReasonEquipBaoKuUse
	GoldLogReasonEquipBaoKuGet
	GoldLogReasonEquipBaoKuLuckyBoxGet
	GoldLogReasonMingGeSynthesisCost
	GoldLogReasonMingGeRefinedCost
	GoldLogReasonMingGeBaptizeCost
	GoldLogReasonHuntCost
	GoldLogReasonZhenFaShengJiCost
	GoldLogReasonZhenQiXianHuoShengJiCost
	GoldLogReasonShenQiDebrisUpCost
	GoldLogReasonShenQiSmeltUpCost
	GoldLogReasonFuncopenRew
	GoldLogReasonMarryPreGift
	GoldLogReasonMarryProposal
	GoldLogReasonMarryProposalReturn
	GoldLogReasonBabyAccelerateUse
	GoldLogReasonBabyChaoShengUse
	GoldLogReasonBabyActivateSkillUse
	GoldLogReasonBabyLearnBuyUse
	GoldLogReasonBabyLockSkillUse
	GoldLogReasonTrade
	GoldLogReasonMarryDingQingJiHuo
	GoldLogReasonQuestQiYuRew
	GoldLogReasonBuyMiddle
	GoldLogReasonBuyHigher
	GoldLogReasonBuyHighest
	GoldLogReasonXianTaoCommit
	GoldLogReasonHouseSell
	GoldLogReasonHouseRent
	GoldLogReasonAdditionSysShenZhuCost
	GoldLogReasonOpenActivityCost
	GoldLogReasonAdditionSysTongLingCost
	GoldLogReasonGuideRew
	GoldLogReasonInvestUpLevelMiddle
	GoldLogReasonInvestUpLevelHigher
	GoldLogReasonInvestUpLevelHighest
	GoldLogReasonShenYuRoundRew
	GoldLogReasonUseYuanBaoKa
	GoldLogReasonBuyWeekCost
	GoldLogReasonBuyWeekRew
	GoldLogReasonWeekDayRew
	GoldLogReasonMajorSaoDangUse
	GoldLogReasonMajorSaoDangGet
	GoldLogReasonQiXueAdvanced
	GoldLogReasonArenaRankReward
	GoldLogReasonBeachShopBuy
	GoldLogReasonArenapvpAttendGuess
	GoldLogReasonMaterialSaodangGet
	GoldLogReasonJieYiTokenChange
	GoldLogReasonJieYiDaoJuChange
	GoldLogReasonArenapvpRew
	GoldLogReasonXianJinExchange
	GoldLogReasonAllianceNewCreateCost
	GoldLogReasonJieYiTokenUpLev
	GoldLogReasonCastingSpiritUpLevel
	GoldLogReasonNewFirstChargeReturn
	GoldLogReasonRingTypeChangeCost
	GoldLogReasonYunYinShopBuy
	GoldLogReasonArenapvpBuyTicket
	GoldLogReasonDrewPools
	GoldLogReasonXianZunCardBuy
	GoldLogReasonXianZunActiviteAdd
	GoldLogReasonXianZunReceiveAdd
	GoldLogReasonRewPoolsDrew
	GoldLogReasonLingshouUplevel
	GoldLogReasonLingwenUplevel
	GoldLogReasonLingshouUpRank
	GoldLogReasonLingshouLinglian
	GoldLogReasonRingAdvance
	GoldLogReasonRingFuse
	GoldLogReasonRingBaoKuUse
	GoldLogReasonRingBaoKuGet
	GoldLogReasonRingBaoKuLuckyBoxGet
)

func (glr GoldLogReason) Reason() int32 {
	return int32(glr)
}

var (
	GoldLogReasonMap = map[GoldLogReason]string{
		GoldLogReasonGM:                            "gm修改",
		GoldLogReasonDanAccelerate:                 "加速炼丹,丹药id:%d,加速炼丹个数num:%d",
		GoldLogReasonQuestAccept:                   "任务接取消耗,任务id:%d",
		GoldLogReasonQuestCommit:                   "任务完成消耗,任务id:%d",
		GoldLogReasonQuestReward:                   "任务完成奖励,任务id:%d",
		GoldLogReasonBuySlots:                      "购买槽位",
		GoldLogReasonBuyDepotSlots:                 "购买仓库槽位",
		GoldLogReasonShopBuyItem:                   "商城购买shopId:%d,商品名字:%s,购买个数:%d",
		GoldLogReasonMountAdvanced:                 "坐骑进阶",
		GoldLogReasonWingAdvanced:                  "战翼进阶",
		GoldLogReasonBodyShieldAdvanced:            "护体盾进阶",
		GoldLogReasonWeapUpstar:                    "兵魂升星购买道具",
		GoldLogReasonEquipSlotStrengthUpgradeAuto:  "装备槽自动升级,位置:%s,等级:%d",
		GoldLogReasonGemMineActive:                 "矿工激活",
		GoldLogReasonGemGamble:                     "赌石元宝消耗",
		GoldLogReasonGemGambleDrop:                 "赌石掉落元宝",
		GoldLogReasonSynthesis:                     "合成物品，合成id:%d,数量:%d",
		GoldLogReasonGiftBuy:                       "礼物购买,礼物id:%d,数量:%d",
		GoldLogReasonTuMoBuyNum:                    "购买屠魔任务次数",
		GoldLogReasonTuMoFinishAll:                 "屠魔任务一次完成,次数:%d",
		GoldLogReasonEmailAttachment:               "邮件附件奖励",
		GoldLogReasonSoulRuinsRewChapter:           "帝陵遗迹章节奖励领取chapter:%d typ:%d",
		GoldLogReasonSoulRuinsSweep:                "帝陵遗迹扫荡,扫荡关卡chapter:%d,typ:%d,level:%d,扫荡次数:%d",
		GoldLogReasonSoulRuinsFinishAll:            "帝陵遗迹一键完成 chapter:%d,typ:%d,level:%d,次数:%d",
		GoldLogReasonSoulRuinsSweepDrop:            "帝陵遗迹扫荡掉落,扫荡关卡chapter:%d,typ:%d,level:%d,扫荡次数:%d",
		GoldLogReasonSoulRuinsFinishAllDrop:        "帝陵遗迹一键完成掉落chapter:%d,typ:%d,level:%d,次数:%d",
		GoldLogReasonSoulRuins:                     "帝陵遗迹挑战成功奖励chapter:%d typ:%d level:%d",
		GoldLogReasonSoulRuinsBuyNum:               "帝陵遗迹购买挑战次数",
		GoldLogReasonSoulRuinsFinishAllCost:        "帝陵遗迹一键完成消耗",
		GoldLogReasonXianfuUpgrade:                 "秘境仙府升级,仙府id:%d,类型:%d",
		GoldLogReasonXianfuAccelerate:              "秘境仙府加速,仙府id:%d,类型:%d",
		GoldLogReasonXianfuSaodang:                 "秘境仙府扫荡消耗,仙府id:%d,类型:%d",
		GoldLogReasonXianfuSaodangRewards:          "秘境仙府扫荡奖励,仙府id:%d,类型:%d",
		GoldLogReasonRealmToKillRew:                "天劫塔前往击杀奖励,层数:%d",
		GoldLogReasonMonsterKilled:                 "杀怪奖励",
		GoldLogReasonEmperorWorship:                "膜拜帝王奖励",
		GoldLogReasonRobEmperor:                    "抢龙椅消耗",
		GoldLogReasonMoonloveGivePresents:          "月下情缘赠送自动购买,物品Id:%d,数量:%d",
		GoldLogReasonMoonloveRankRew:               "月下情缘排行榜奖励，名次:%d,类型:%d",
		GoldLogReasonMoonloveTickRew:               "月下情缘定时奖励",
		GoldLogReasonAllianceDonate:                "仙盟捐献,仙盟id:%d",
		GoldLogReasonAllianceCallGuard:             "城战召唤守卫,仙盟id:%d,地图id:%d,守卫id:%d",
		GoldLogReasonAllianceSceneDoorReward:       "城战城门奖励,仙盟id:%d,门:%d",
		GoldLogReasonSecretCardStarRew:             "天机牌领取运势箱奖励",
		GoldLogReasonSecretCardFinishAllStarRew:    "一键牌一键完成剩余运势箱奖励",
		GoldLogReasonSecretCardFinishAll:           "天机牌一键完成奖励",
		GoldLogReasonSecretCardFinishAllCost:       "天机牌一键完成消耗",
		GoldLogReasonSecretCardImmediateFinishCost: "天机牌直接完成消耗",
		GoldLogReasonSecretCardFinish:              "天机牌任务完成奖励",
		GoldLogReasonTransportationCost:            "押镖消耗,镖车类型：%s",
		GoldLogReasonTransportationRobRew:          "劫镖奖励，镖车类型：%s",
		GoldLogReasonTransportationFailRew:         "押镖失败奖励，镖车类型：%s",
		GoldLogReasonTransportationFinishRew:       "押镖成功奖励，镖车类型：%s",
		GoldLogReasonFoundResource:                 "资源找回奖励,资源id:%d,次数:%d,波数:%d",
		GoldLogReasonFoundResourceCost:             "资源找回消耗，资源id:%d,次数:%d",
		GoldLogReasonFoundResourceBatchCost:        "资源找回批量消耗，资源map：%v",
		GoldLogReasonFourGodOpenBox:                "四神遗迹开启宝箱奖励,钥匙数:%d",
		GoldLogReasonFeatherAdvanced:               "护体仙羽进阶",
		GoldLogReasonShieldAdvanced:                "神盾尖刺进阶",
		GoldLogReasonShenfaAdvanced:                "身法进阶",
		GoldLogReasonLingyuAdvanced:                "领域进阶",
		GoldLogReasonMarryRingFeedCost:             "婚戒培养消耗",
		GoldLogReasonMarryTreeFeedCost:             "爱情树培养消耗",
		GoldLogReasonMarryWeddingGrade:             "婚宴预定,预定婚宴档次:%d,婚车档次:%d,喜糖档次:%d",
		GoldLogReasonXianfuFinishAllUse:            "秘境仙府一键完成消耗,仙府id:%d,类型:%d",
		GoldLogReasonXianfuFinishAllRew:            "秘境仙府一键完成奖励,仙府id:%d,类型:%d",
		GoldLogReasonBoxUse:                        "宝箱开启消耗",
		GoldLogReasonBoxGet:                        "宝箱开启奖励",
		GoldLogReasonChessUse:                      "苍龙棋局消耗,类型：%s",
		GoldLogReasonOneArenaSellKun:               "鲲一键卖出获得",
		GoldLogReasonChessGet:                      "苍龙棋局奖励,类型：%s",
		GoldLogReasonFirstCharge:                   "首冲奖励",
		GoldLogReasonFirstLevelCharge:              "档次首充奖励,充值档次:%d",
		GoldLogReasonBuyInvest:                     "购买七日投资计划,typ:%d",
		GoldLogReasonBuyHuiYuan:                    "购买会员消耗，会员类型：%d",
		GoldLogReasonHuiYuanRew:                    "会员每日奖励，会员类型：%d",
		GoldLogReasonOpenActivityRew:               "开服活动奖励,类型:%d,子类型:%d",
		GoldLogReasonAnqiAdvanced:                  "暗器进阶",
		GoldLogReasonGiftCode:                      "礼包兑换奖励",
		GoldLogReasonCharge:                        "充值奖励",
		GoldLogReasonXueChiAutoBuy:                 "血池自动购买生命瓶:物品id:%d,数量:%d",
		GoldLogReasonTuLongKillBossRew:             "跨服屠龙击杀boss奖励",
		GoldLogReasonDrewUse:                       "抽奖活动消耗,类型:%d",
		GoldLogReasonBuyDiscountUse:                "购买折扣礼包消耗,typ:%d,subType:%d",
		GoldLogReasonReliveAutoBuy:                 "复活丹自动购买:%d",
		GoldLogReasonCollectRew:                    "采集物奖励",
		GoldLogReasonBuyVipGift:                    "购买vip礼包消耗",
		GoldLogReasonTuMoImmediateFinishCost:       "屠魔任务直接完成消耗",
		GoldLogReasonTuMoImmediateFinish:           "屠魔任务直接完成",
		GoldLogReasonFireworks:                     "放烟花,物品id:%d,数量:%d",
		GoldLogReasonBuyXianHua:                    "购买鲜花",
		GoldLogReasonEmperorReward:                 "帝王奖励",
		GoldLogReasonPrivilegeCharge:               "后台充值",
		GoldLogReasonPrivilegeItem:                 "后台扶持",
		GoldLogReasonMassacreAdvanced:              "戮仙刃消耗，当前%d阶%d星",
		GoldLogReasonTianShuActivate:               "天书激活消耗,%s",
		GoldLogReasonTianShuReceive:                "天书每日奖励,%s",
		GoldLogReasonGoldEquipAutoBuy:              "金装开光购买道具",
		GoldLogReasonBuyJunior:                     "购买初级投资计划消耗",
		GoldLogReasonBuySenior:                     "购买高级投资计划消耗",
		GoldLogReasonSystemSkillActive:             "系统技能激活消耗,模板id:%d",
		GoldLogReasonSystemSkillUpgrade:            "系统技能升级消耗,模板id:%d",
		GoldLogReasonFaBaoAdvanced:                 "法宝进阶",
		GoldLogReasonAdditionSysLevel:              "%s系統消耗，当前%d等级%d进度%d次数",
		GoldLogReasonMaterialRew:                   "材料副本奖励,typ:%s",
		GoldLogReasonXianTiAdvanced:                "仙体进阶",
		GoldLogReasonLivenessStarRew:               "活跃度领取宝箱奖励",
		GoldLogReasonUnrealBossBuyPilao:            "幻境boss购买疲劳值",
		GoldLogReasonFoeFeedbackBuyProtect:         "购买仇人反馈保护消耗",
		GoldLogReasonUnrealBossDrop:                "幻境boss掉落包",
		GoldLogReasonDailyQuestCommitDouble:        "日环任务领取双倍奖励消耗",
		GoldLogReasonDailyQuestReward:              "完成日环任务奖励,日环类型：%s",
		GoldLogReasonDailyQuestImmediateFinishAll:  "一键完成%s任务消耗,完成次数:%d",
		GoldLogReasonDailyQuestFinishAll:           "一键完成%s任务奖励",
		GoldLogReasonBaGuaToKillRew:                "八卦秘境前往击杀奖励,层数:%d",
		GoldLogReasonOutlandBossDrop:               "外域boss掉落包",
		GoldLogReasonLaBaUse:                       "参与拉霸消耗,次数:%d",
		GoldLogReasonLaBaGet:                       "参与拉霸奖励,次数:%d",
		GoldLogReasonSongBuTingRew:                 "领取元宝送不停奖励",
		GoldLogReasonInventoryResourceUse:          "背包资源使用获得",
		GoldLogReasonTeamCopyReward:                "组队副本奖励,副本类型:%d",
		GoldLogReasonQuizAnswer:                    "仙尊问答奖励,玩家等级:%d ,答题结果:%d",
		GoldLogReasonDianXingAdvanced:              "点星系统升级，当前%d星谱%d星%d进度%d次数",
		GoldLogReasonDianXingJieFengAdvanced:       "点星解封升级，当前%d阶级%d进度%d次数",
		GoldLogReasonMaterialSaoDangUse:            "材料副本扫荡%d次消耗，副本类型：%s",
		GoldLogReasonTianMoAdvanced:                "天魔体进阶消耗",
		GoldLogReasonShiHunFanAdvanced:             "噬魂幡进阶消耗，当前%d阶",
		GoldLogReasonActivityTickRew:               "活动定时奖励,活动类型:%d",
		GoldLogReasonLingTongAdvanced:              "灵童%s进阶消耗",
		GoldLogReasonActivityCost:                  "开服活动消耗,类型:%d,子类型:%d",
		GoldLogReasonOpenActivityChargeReturn:      "运营活动充值返还",
		GoldLogReasonLingyuActivate:                "领域激活",
		GoldLogReasonFeiShengResetQn:               "重置飞升潜能，飞升等级:%d",
		GoldLogReasonShenMoWarRankReward:           "神魔战场周排名奖励,仙盟排名:%d",
		GoldLogReasonHongBaoSnatch:                 "抢送红包获得",
		GoldLogReasonChatAward:                     "发言奖励",
		GoldLogReasonAdditionSysHuaLing:            "%s系統化灵食用消耗，当前%d等级%d进度%d次数",
		GoldLogReasonGroupCollectRew:               "摸金卡牌收集奖励，卡牌组合类型：%d",
		GoldLogReasonEquipBaoKuUse:                 "装备宝库消耗,等级：%d，转数：%d",
		GoldLogReasonEquipBaoKuGet:                 "装备宝库奖励,等级：%d，转数：%d",
		GoldLogReasonEquipBaoKuLuckyBoxGet:         "装备宝库幸运宝箱奖励",
		GoldLogReasonMingGeSynthesisCost:           "命格合成消耗",
		GoldLogReasonMingGeRefinedCost:             "命盘祭炼消耗",
		GoldLogReasonMingGeBaptizeCost:             "命格命理洗练消耗",
		GoldLogReasonHuntCost:                      "寻宝消耗，类型：%d",
		GoldLogReasonZhenFaShengJiCost:             "阵法升级消耗",
		GoldLogReasonZhenQiXianHuoShengJiCost:      "阵法仙火升级消耗",
		GoldLogReasonShenQiDebrisUpCost:            "神器碎片升级消耗,类型：%s，部位：%s",
		GoldLogReasonShenQiSmeltUpCost:             "神器淬炼升级消耗,类型：%s，部位：%s,",
		GoldLogReasonFuncopenRew:                   "功能开启奖励,功能id:%d",
		GoldLogReasonMarryPreGift:                  "结婚游车送贺礼消耗",
		GoldLogReasonMarryProposal:                 "求婚消耗",
		GoldLogReasonMarryProposalReturn:           "求婚失败退回",
		GoldLogReasonBabyAccelerateUse:             "宝宝加速出生消耗",
		GoldLogReasonBabyChaoShengUse:              "宝宝超生消耗",
		GoldLogReasonBabyActivateSkillUse:          "宝宝激活天赋技能消耗，babyId:%d,第%d次",
		GoldLogReasonBabyLearnBuyUse:               "宝宝读书物品购买消耗，物品id：%d，数量：%d，宝宝id:%d,等级:%d,品质：%d",
		GoldLogReasonBabyLockSkillUse:              "宝宝锁定天赋技能消耗,babyId:%d,第%d次",
		GoldLogReasonTrade:                         "交易花费,交易id[%d]",
		GoldLogReasonMarryDingQingJiHuo:            "定情信物购买消耗",
		GoldLogReasonQuestQiYuRew:                  "奇遇任务奖励，qiyuId:%d",
		GoldLogReasonBuyMiddle:                     "购买中级投资计划消耗",
		GoldLogReasonBuyHigher:                     "购买高级投资计划消耗",
		GoldLogReasonBuyHighest:                    "购买顶级投资计划消耗",
		GoldLogReasonXianTaoCommit:                 "提交仙桃奖励",
		GoldLogReasonHouseSell:                     "房子出售奖励，序号：%d",
		GoldLogReasonHouseRent:                     "房子租金奖励，序号：%d",
		GoldLogReasonAdditionSysShenZhuCost:        "%s系统%s部位神铸消耗，当前%d等级%d进度%d次数",
		GoldLogReasonOpenActivityCost:              "开服活动消耗,类型:%d,子类型:%d",
		GoldLogReasonAdditionSysTongLingCost:       "%s系统通灵消耗，当前%d等级%d进度%d次数",
		GoldLogReasonGuideRew:                      "引导副本奖励，类型：%s",
		GoldLogReasonInvestUpLevelMiddle:           "升级等级投资计划到中级",
		GoldLogReasonInvestUpLevelHigher:           "升级等级投资计划到高级",
		GoldLogReasonInvestUpLevelHighest:          "升级等级投资计划到最高级",
		GoldLogReasonShenYuRoundRew:                "神域之战奖励，参赛轮：%d",
		GoldLogReasonUseYuanBaoKa:                  "使用元宝卡增加元宝,物品id[%d],数量[%d]",
		GoldLogReasonBuyWeekCost:                   "购买周卡消耗，类型：%s",
		GoldLogReasonBuyWeekRew:                    "购买周卡奖励，类型：%s",
		GoldLogReasonWeekDayRew:                    "周卡每日奖励，类型：%s",
		GoldLogReasonMajorSaoDangUse:               "夫妻副本扫荡%d次消耗,type:%s,fubenId:%d",
		GoldLogReasonMajorSaoDangGet:               "夫妻副本扫荡%d次奖励,type:%s,fubenId:%d",
		GoldLogReasonQiXueAdvanced:                 "泣血枪消耗，当前%d阶%d星",
		GoldLogReasonArenaRankReward:               "领取3v3周排名奖励,排名:%d",
		GoldLogReasonBeachShopBuy:                  "购买沙滩商店商品消耗,活动type：%d,活动subType: %d, 商品type：%d",
		GoldLogReasonArenapvpAttendGuess:           "比武大会参与竞猜消耗，类型：%s",
		GoldLogReasonMaterialSaodangGet:            "材料副本扫荡%d次奖励,typ:%s",
		GoldLogReasonJieYiTokenChange:              "结义信物替换，旧类型：%s, 新类型：%s, 消耗方式：%s",
		GoldLogReasonJieYiDaoJuChange:              "结义道具替换，旧类型：%s, 新类型：%s, 消耗方式：%s",
		GoldLogReasonArenapvpRew:                   "比武大会奖励，获胜：%v， 类型：%s",
		GoldLogReasonXianJinExchange:               "现金兑换,记录[%d],兑换金额[%d]",
		GoldLogReasonAllianceNewCreateCost:         "创建新版本仙盟花费元宝",
		GoldLogReasonJieYiTokenUpLev:               "结义信物升级自动购买消耗元宝",
		GoldLogReasonCastingSpiritUpLevel:          "铸灵升级消耗元宝",
		GoldLogReasonNewFirstChargeReturn:          "新首充活动类型返利,充值[%d]",
		GoldLogReasonRingTypeChangeCost:            "结婚戒指替换消耗，旧类型：%s, 新类型：%s",
		GoldLogReasonYunYinShopBuy:                 "购买沙滩商店商品消耗,活动type：%d,活动subType: %d, 商品type：%d",
		GoldLogReasonArenapvpBuyTicket:             "比武大会购买门票消耗",
		GoldLogReasonDrewPools:                     "奖池抽奖，奖池位置[%d]",
		GoldLogReasonXianZunCardBuy:                "购买仙尊特权卡消耗，类型：%s",
		GoldLogReasonXianZunActiviteAdd:            "购买仙尊特权卡增加，类型：%s",
		GoldLogReasonXianZunReceiveAdd:             "领取每日仙尊特权卡奖励，类型：%s",
		GoldLogReasonRewPoolsDrew:                  "仙人指路-奖池抽奖",
		GoldLogReasonLingshouUplevel:               "上古之灵灵兽升级，灵兽类型[%s]",
		GoldLogReasonLingwenUplevel:                "上古之灵灵纹升级，灵兽类型[%s], 灵纹类型[%s]",
		GoldLogReasonLingshouUpRank:                "上古之灵灵兽进阶，灵兽类型[%s]",
		GoldLogReasonLingshouLinglian:              "上古之灵灵兽灵炼，灵兽类型[%s]",
		GoldLogReasonRingAdvance:                   "特戒进阶消耗，特戒类型：%s",
		GoldLogReasonRingFuse:                      "特戒融合消耗，特戒类型：%s",
		GoldLogReasonRingBaoKuUse:                  "特戒宝库探索消耗",
		GoldLogReasonRingBaoKuGet:                  "特戒宝库探索增加",
		GoldLogReasonRingBaoKuLuckyBoxGet:         "特戒宝库幸运宝箱奖励",
	}
)

func (glr GoldLogReason) String() string {
	return GoldLogReasonMap[glr]
}

var (
	GoldLogReasonRewardMap = map[GoldLogReason]bool{
		GoldLogReasonQuestReward: true,

		GoldLogReasonEmailAttachment:     true,
		GoldLogReasonSoulRuinsRewChapter: true,

		GoldLogReasonSoulRuinsSweepDrop:     true,
		GoldLogReasonSoulRuinsFinishAllDrop: true,
		GoldLogReasonSoulRuins:              true,

		GoldLogReasonXianfuSaodangRewards: true,
		GoldLogReasonRealmToKillRew:       true,
		GoldLogReasonMonsterKilled:        true,
		GoldLogReasonEmperorWorship:       true,

		GoldLogReasonMoonloveRankRew: true,
		GoldLogReasonMoonloveTickRew: true,

		GoldLogReasonAllianceSceneDoorReward:    true,
		GoldLogReasonSecretCardStarRew:          true,
		GoldLogReasonSecretCardFinishAllStarRew: true,
		GoldLogReasonSecretCardFinishAll:        true,

		GoldLogReasonSecretCardFinish: true,

		GoldLogReasonTransportationRobRew:    true,
		GoldLogReasonTransportationFailRew:   true,
		GoldLogReasonTransportationFinishRew: true,
		GoldLogReasonFoundResource:           true,

		GoldLogReasonFourGodOpenBox: true,

		GoldLogReasonBoxGet: true,

		GoldLogReasonOneArenaSellKun: true,

		GoldLogReasonFirstCharge:      true,
		GoldLogReasonFirstLevelCharge: true,

		GoldLogReasonOpenActivityRew:  true,
		GoldLogReasonOpenActivityCost: true,

		GoldLogReasonGiftCode: true,
		GoldLogReasonCharge:   true,

		GoldLogReasonTuLongKillBossRew: true,

		GoldLogReasonEmperorReward:   true,
		GoldLogReasonPrivilegeCharge: true,

		GoldLogReasonTianShuReceive: true,
	}
)
