package log

type InventoryLogReason int32

const (
	//gm添加
	InventoryLogReasonGM InventoryLogReason = iota + 1
	InventoryLogReasonPutOn
	InventoryLogReasonTakeOff
	InventoryLogReasonSell
	InventoryLogReasonUse
	InventoryLogReasonEquipSlotStrengthUpgrade
	InventoryLogReasonEquipSlotStrengthStar
	InventoryLogReasonEquipUpgrade
	InventoryLogReasonDanEat
	InventoryLogReasonAchemyStart
	InventoryLogReasonAchemyRec
	InventoryLogReasonShieldAdvanced
	InventoryLogReasonShieldJJDan
	InventoryLogReasonMountAdvanced
	InventoryLogReasonMountEatClu
	InventoryLogReasonMountEatUn
	InventoryLogReasonMountUnreal
	InventoryLogReasonWingAdvanced
	InventoryLogReasonWingUseTrialCard
	InventoryLogReasonWingEatUn
	InventoryLogReasonWingUnreal
	InventoryLogReasonTitleActive
	InventoryLogReasonQuestAccept
	InventoryLogReasonQuestCommit
	InventoryLogReasonQuestReward
	InventoryLogReasonTuMoFinishAllRew
	InventoryLogReasonShopBuy
	InventoryLogReasonFashionActive
	InventoryLogReasonWeaponActive
	InventoryLogReasonWeaponEatDan
	InventoryLogReasonWeaponUpstar
	InventoryLogReasonWeaponAwaken
	InventoryLogReasonSoulActive
	InventoryLogReasonSoulFeed
	InventoryLogReasonSoulUpgrade
	InventoryLogReasonSoulAwaken
	InventoryLogReasonPutOnGem
	InventoryLogReasonTakeOffGem
	InventoryLogReasonJueXueActive
	InventoryLogReasonJueXueUpgrade
	InventoryLogReasonXinFaActive
	InventoryLogReasonXinFaUpgrade
	InventoryLogReasonGemMineActive
	InventoryLogReasonGemGambleDrop
	InventoryLogReasonSynthesisStart
	InventoryLogReasonSynthesisReceive
	InventoryLogReasonGift
	InventoryLogReasonUseToken
	InventoryLogReasonEmailAttachment
	InventoryLogReasonSoulRuinsRewChapter
	InventoryLogReasonSoulRuinsSweepDrop
	InventoryLogReasonSoulRuinsFinishAll
	InventoryLogReasonSoulRuinsSweep
	InventoryLogReasonSoulRuins
	InventoryLogReasonSoulRuinsForceGet
	InventoryLogReasonXianfuUpgrade
	InventoryLogReasonXianfuSaodang
	InventoryLogReasonXianfuSaodangRewards
	InventoryLogReasonXianfuChallengeSuccessRewards
	InventoryLogReasonXianfuChallenge
	InventoryLogReasonRealmToKillRew
	InventoryLogReasonMonsterKilled
	InventoryLogReasonMoonloveGivePresents
	InventoryLogReasonMoonloveTickRew
	InventoryLogReasonAllianceCreate
	InventoryLogReasonAllianceDonate
	InventoryLogReasonAllianceDonateHuFu
	InventoryLogReasonAllianceCallGuard
	InventoryLogReasonAllianceSceneDoorReward
	InventoryLogReasonAllianceSkillUpgrade
	InventoryLogReasonAllianceConvert
	InventoryLogReasonSecretCardStarRew
	InventoryLogReasonSecretCardFinishAll
	InventoryLogReasonSecretCardFinish
	InventoryLogReasonFoundResource
	InventoryLogReasonDragonFeed
	InventoryLogReasonDragonAdvancedRew
	InventoryLogReasonFourGodOpenBox
	InventoryLogReasonFourGodBlack
	InventoryLogReasonFeatherAdvanced
	InventoryLogReasonShieldSpikesAdvanced
	InventoryLogReasonLingyuEatUn
	InventoryLogReasonLingyuAdvanced
	InventoryLogReasonLingyuUnreal
	InventoryLogReasonMarryProposal
	InventoryLogReasonMarryProposalFail
	InventoryLogReasonMarryRingFeedUse
	InventoryLogReasonMarryTreeFeedUse
	InventoryLogReasonMarryWedGift
	InventoryLogReasonMarryRingReplace
	InventoryLogReasonMarryGift
	InventoryLogReasonBoxUse
	InventoryLogReasonBoxRew
	InventoryLogReasonGoldEquipStrengthUpgrade
	InventoryLogReasonGoldEquipChongzhuUse
	InventoryLogReasonGoldEquipChongzhuGet
	InventoryLogReasonChessGet
	InventoryLogReasonShenfaAdvanced
	InventoryLogReasonShenfaUnreal
	InventoryLogReasonShenfaEatUn
	InventoryLogReasonSaveInDepot
	InventoryLogReasonTakeOutDepot
	InventoryLogReasonOneArenaOutputKun
	InventoryLogReasonOneArenaSellKun
	InventoryLogReasonChangeScene
	InventoryLogReasonRelive
	InventoryLogReasonRecover
	InventoryLogReasonFirstCharge
	InventoryLogReasonChessAttend
	InventoryLogReasonWelfareRew
	InventoryLogReasonInvestRew
	InventoryLogReasonFeedbackRew
	InventoryLogReasonArena
	InventoryLogReasonArenaExtra
	InventoryLogReasonOpenActivityRew
	InventoryLogReasonArenaRelive
	InventoryLogReasonMajorRew
	InventoryLogReasonAnqiAdvanced
	InventoryLogReasonAnqiEatDan
	InventoryLogReasonFashionUpstar
	InventoryLogReasonBlessDan
	InventoryLogReasonTransportRew
	InventoryLogReasonTransportFailRew
	InventoryLogReasonMountUpstar
	InventoryLogReasonFuncopenRew
	InventoryLogReasonCollect
	InventoryLogReasonGoldEquipEat
	InventoryLogReasonVipGiftGet
	InventoryLogReasonTuMoImmediateFinishRew
	InventoryLogReasonVipFreeGiftGet
	InventoryLogReasonWingUpstar
	InventoryLogReasonLingYuUpstar
	InventoryLogReasonShenFaUpstar
	InventoryLogReasonEmperorRobReward
	InventoryLogReasonEmperorOpenBox
	InventoryLogReasonShootFireworks
	InventoryLogReasonSaveInAllianceDepot
	InventoryLogReasonTakeOutAllianceDepot
	InventoryLogReasonMassacreAdvanced
	InventoryLogReasonEnterTower
	InventoryLogReasonTianShuUplevel
	InventoryLogReasonMyBossUse
	InventoryLogReasonGoldEquipOpenLight
	InventoryLogReasonGoldEquipTunShiReturn
	InventoryLogReasonGoldEquipUpstar
	InventoryLogReasonSystemSkillActive
	InventoryLogReasonSystemSkillUpgrade
	InventoryLogReasonFoeTrackUse
	InventoryLogReasonDeliverUse
	InventoryLogReasonQuestQuickTransfer
	InventoryLogReasonUnrealBossGet
	InventoryLogReasonFaBaoAdvanced
	InventoryLogReasonFaBaoUnreal
	InventoryLogReasonFaBaoEatUn
	InventoryLogReasonFaBaoUpstar
	InventoryLogReasonFaBaoTongLing
	InventoryLogReasonMaterialChallenge
	InventoryLogReasonMaterialSaodangUse
	InventoryLogReasonMaterialSaodangGet
	InventoryLogReasonXueDunEatClu
	InventoryLogReasonXianTiAdvanced
	InventoryLogReasonXianTiEatUn
	InventoryLogReasonXianTiUnreal
	InventoryLogReasonXianTiUpstar
	InventoryLogReasonXueDunUpgrade
	InventoryLogReasonLivenessStarRew
	InventoryLogReasonAdditionSysShengJi
	InventoryLogReasonHuiYuanRew
	InventoryLogReasonTianShuRew
	InventoryLogReasonAddFriendRew
	InventoryLogReasonDailyQuestReward
	InventoryLogReasonDailyQuestFinishAllReward
	InventoryLogReasonBaGuaToKillRew
	InventoryLogReasonOutlandBossGet
	InventoryLogReasonSongBuTingRew
	InventoryLogReasonTeamCopyReward
	InventoryLogReasonQuizAnswer
	InventoryLogReasonPlayerTransfer
	InventoryLogReasonDianXingAdvanced
	InventoryLogReasonFriendNoticeFeedback
	InventoryLogReasonFriendGiveGet
	InventoryLogReasonFriendGiveUse
	InventoryLogReasonDianXingJieFengAdvanced
	InventoryLogReasonWardrobeEatDan
	InventoryLogReasonTianMoAdvanced
	InventoryLogReasonTianMoEatDan
	InventoryLogReasonShiHunFanAdvanced
	InventoryLogReasonShiHunFanEatDan
	InventoryLogReasonActivityTickRew
	InventoryLogReasonLingTongAdvaced
	InventoryLogReasonLingTongDevTongLing
	InventoryLogReasonLingTongDevUnreal
	InventoryLogReasonLingTongDevEatUn
	InventoryLogReasonLingTongDevUpstar
	InventoryLogReasonLingTongDevEatClu
	InventoryLogReasonOpenActivityUse
	InventoryLogReasonLingTongActivate
	InventoryLogReasonLingTongUpgrade
	InventoryLogReasonLingTongEatClu
	InventoryLogReasonLingTongFashionActivate
	InventoryLogReasonLingTongFashionUpstar
	InventoryLogReasonLingTongRename
	InventoryLogReasonKaiFuMuBiaoGroupRew
	InventoryLogReasonFeiShenEatDan
	InventoryLogReasonShenMoWar
	InventoryLogReasonHongBaoSend
	InventoryLogReasonHongBaoSnatch
	InventoryLogReasonTianFuAwaken
	InventoryLogReasonTianFuUpgrade
	InventoryLogReasonAdditionSysHuaLing
	InventoryLogReasonZhuanShengGiftReceive
	InventoryLogReasonSupremeTitleActive
	InventoryLogReasonEquipBaoKuAttend
	InventoryLogReasonEquipBaoKuGet
	InventoryLogReasonTakeOutMiBaoDepot
	InventoryLogReasonEquipBaoKuResolveUse
	InventoryLogReasonEquipBaoKuResolveReturn
	InventoryLogReasonGoldEquipExtendUse
	InventoryLogReasonEquipBaoKuLuckyBoxGet
	InventoryLogReasonMingGeMosaicUse
	InventoryLogReasonMingGeUnloadAdd
	InventoryLogReasonMingGeSynthesisUse
	InventoryLogReasonMingGeRefinedUse
	InventoryLogReasonMingGeMingLiBaptizeUse
	InventoryLogReasonMingGeSynthesisAdd
	InventoryLogReasonTuLongEquipRongHeUse
	InventoryLogReasonTuLongEquipRongHeGet
	InventoryLogReasonTuLongEquipStrengthenUse
	InventoryLogReasonTuLongEquipZhuanHuaUse
	InventoryLogReasonTuLongEquipZhuanHuaGet
	InventoryLogReasonHuntRew
	InventoryLogReasonHuntUse
	InventoryLogReasonZhenFaActivate
	InventoryLogReasonZhenFaShengJi
	InventoryLogReasonZhenQiAdvanced
	InventoryLogReasonZhenFaXianHuoShengJi
	InventoryLogReasonYingLingPuXq
	InventoryLogReasonYingLingPuLevelUp
	InventoryLogReasonShenQiDebrisUpCost
	InventoryLogReasonShenQiSmeltUpCost
	InventoryLogReasonResolveCost
	InventoryLogReasonMarryPreGift
	InventoryLogReasonTradeUpload
	InventoryLogReasonMarryJiNian
	InventoryLogReasonBabyDongFangUse
	InventoryLogReasonBabyChangeName
	InventoryLogReasonBabyEatTonicUse
	InventoryLogReasonBabyLearnUse
	InventoryLogReasonBabyDongFangReturn
	InventoryLogReasonBabyEquipToy
	InventoryLogReasonMarryDingQingJiHuo
	InventoryLogReasonBabyToyUplevelUse
	InventoryLogReasonMarryJiNianShiZhuang
	InventoryLogReasonQiYuRew
	InventoryLogReasonBabyRefreshUse
	InventoryLogReasonBabyZhuanShiReturn
	InventoryLogReasonTradeBuy
	InventoryLogReasonXianTaoCommit
	InventoryLogReasonHouseActivateUse
	InventoryLogReasonHouseUplevelUse
	InventoryLogReasonHouseRepairUse
	InventoryLogReasonHouseUplevelGet
	InventoryLogReasonAdditionSysShenZhuCost
	InventoryLogReasonAdditionSysAwake
	InventoryLogReasonShopBuyItem
	InventoryLogReasonItemChaiJieCost
	InventoryLogReasonItemChaiJieGet
	InventoryLogReasonAdditionSysTongLingCost
	InventoryLogReasonYuXiDayRew
	InventoryLogReasonGuideRew
	InventoryLogReasonShenYuRoundRew
	InventoryLogReasonWeedBuyRew
	InventoryLogReasonWeedDayRew
	InventoryLogReasonUnlockGem
	InventoryLogReasonFuShiActivite
	InventoryLogReasonFuShiUpLevel
	InventoryLogReasonMajorSaodangUse
	InventoryLogReasonMajorSaodangGet
	InventoryLogReasonLingTongUpstar
	InventoryLogReasonTitleUpStar
	InventoryLogReasonQiXueAdvanced
	InventoryLogReasonArenaRankReward
	InventoryLogReasonBeachShopActivite
	InventoryLogReasonBeachShopBuy
	InventoryLogReasonBaoMingChuangShiGet
	InventoryLogReasonInviteJieYiUse
	InventoryLogReasonJieYiTokenTypeChangeUse
	InventoryLogReasonJieYiTokenLevelChangeUse
	InventoryLogReasonJieYiJiuYuanUse
	InventoryLogReasonArenapvpRew
	InventoryLogReasonWushuangWeaponEat
	InventoryLogReasonWushuangWeaponBreakthrough
	InventoryLogReasonMiZangCost
	InventoryLogReasonMiZangGet
	InventoryLogReasonChuangShiChengFangJianShe
	InventoryLogReasonChuangShiGuanZhiUse
	InventoryLogReasonChuangShiGuanZhiGet
	InventoryLogReasonChuangShiFashionGet
	InventoryLogReasonChuangShiJianSheActivateSkillUse
	InventoryLogReasonCastingSpiritUplevel
	InventoryLogReasonGodCastingUpLevel
	InventoryLogReasonForgeSoulUplevel
	InventoryLogReasonGodCastingInherit
	InventoryLogReasonWushuangPutOn
	InventoryLogReasonWushuangTakeOff
	InventoryLogReasonZhenXiCost
	InventoryLogReasonArenaBossCost
	InventoryLogReasonMarryDingQingTokenGive
	InventoryLogReasonTongTianTaLingTongReward
	InventoryLogReasonTongTianTaMingGeReward
	InventoryLogReasonTongTianTaTuLongReward
	InventoryLogReasonYunYinShopBuy
	InventoryLogReasonYunYinShopReward
	InventoryLogReasonRewPoolsDrew
	InventoryLogReasonAdditionSysLingZhuUplevel
	InventoryLogReasonXianZunCardActiviteAdd
	InventoryLogReasonXianZunCardReceiveAdd
	InventoryLogReasonLingShouUplevel
	InventoryLogReasonLingWenUplevel
	InventoryLogReasonLingShouUpRank
	InventoryLogReasonLingShouLinglian
	InventoryLogReasonRingAdvance
	InventoryLogReasonRingStrengthen
	InventoryLogReasonRingJingLing
	InventoryLogReasonRingBaoKuAttend
	InventoryLogReasonRingBaoKuGet
	InventoryLogReasonRingBaoKuLuckyBoxGet
	InventoryLogReasonLingShouReceive
	InventoryLogReasonRingFuseGet
)

func (ilr InventoryLogReason) Reason() int32 {
	return int32(ilr)
}

var (
	inventoryLogReasonMap = map[InventoryLogReason]string{
		InventoryLogReasonGM:                               "gm修改",
		InventoryLogReasonPutOn:                            "穿戴",
		InventoryLogReasonTakeOff:                          "脱下",
		InventoryLogReasonSell:                             "出售",
		InventoryLogReasonUse:                              "使用",
		InventoryLogReasonEquipSlotStrengthUpgrade:         "装备槽强化升级",
		InventoryLogReasonEquipSlotStrengthStar:            "装备槽强化升星",
		InventoryLogReasonEquipUpgrade:                     "装备升阶",
		InventoryLogReasonDanEat:                           "全部食丹使用",
		InventoryLogReasonAchemyStart:                      "炼丹药id:%d",
		InventoryLogReasonAchemyRec:                        "完成炼丹,领取丹药",
		InventoryLogReasonShieldAdvanced:                   "护体盾进阶使用,当前阶数:%d",
		InventoryLogReasonAnqiAdvanced:                     "暗器进阶使用,当前阶数:%d",
		InventoryLogReasonAnqiEatDan:                       "暗器食丹",
		InventoryLogReasonShieldJJDan:                      "食用金甲丹",
		InventoryLogReasonMountAdvanced:                    "坐骑进阶使用,当前阶数:%d",
		InventoryLogReasonMountEatClu:                      "坐骑食培养丹",
		InventoryLogReasonMountEatUn:                       "坐骑食幻化丹",
		InventoryLogReasonMountUnreal:                      "坐骑幻化,幻化的阶数:%d",
		InventoryLogReasonWingAdvanced:                     "战翼进阶使用,当前阶数:%d",
		InventoryLogReasonWingUseTrialCard:                 "使用战翼试用卡",
		InventoryLogReasonWingEatUn:                        "战翼食幻化丹",
		InventoryLogReasonWingUnreal:                       "战翼幻化,幻化的阶数:%d",
		InventoryLogReasonTitleActive:                      "称号激活,称号id:%d",
		InventoryLogReasonQuestAccept:                      "任务接取消耗,任务id:%d",
		InventoryLogReasonQuestCommit:                      "任务完成消耗,任务id:%d",
		InventoryLogReasonQuestReward:                      "任务完成奖励,任务id:%d",
		InventoryLogReasonTuMoFinishAllRew:                 "屠魔任务一键完成奖励,次数:%d",
		InventoryLogReasonShopBuy:                          "商铺购买道具",
		InventoryLogReasonFashionActive:                    "时装激活,时装id:%d",
		InventoryLogReasonWeaponActive:                     "兵魂激活,兵魂id:%d",
		InventoryLogReasonWeaponEatDan:                     "兵魂食培养丹",
		InventoryLogReasonWeaponUpstar:                     "兵魂升星",
		InventoryLogReasonWeaponAwaken:                     "兵魂觉醒",
		InventoryLogReasonSoulActive:                       "帝魂激活,帝魂tag:%d",
		InventoryLogReasonSoulFeed:                         "帝魂喂养,帝魂tag:%d",
		InventoryLogReasonSoulUpgrade:                      "帝魂升级,帝魂tag:%d,当前级别:%d",
		InventoryLogReasonSoulAwaken:                       "帝魂觉醒,帝魂tag:%d",
		InventoryLogReasonPutOnGem:                         "镶嵌宝石",
		InventoryLogReasonTakeOffGem:                       "卸下宝石",
		InventoryLogReasonJueXueActive:                     "绝学激活,绝学typ:%d",
		InventoryLogReasonJueXueUpgrade:                    "绝学升级",
		InventoryLogReasonXinFaActive:                      "心法激活,心法typ:%d",
		InventoryLogReasonXinFaUpgrade:                     "心法升级",
		InventoryLogReasonGemMineActive:                    "矿工激活",
		InventoryLogReasonGemGambleDrop:                    "赌石掉落物品",
		InventoryLogReasonSynthesisStart:                   "物品合成,合成id:%d,数量:%d",
		InventoryLogReasonSynthesisReceive:                 "合成成功,加入背包",
		InventoryLogReasonGift:                             "赠送好友礼物",
		InventoryLogReasonUseToken:                         "使用屠魔令",
		InventoryLogReasonEmailAttachment:                  "邮件附件",
		InventoryLogReasonSoulRuinsRewChapter:              "帝陵遗迹章节奖励领取chapter:%d typ:%d",
		InventoryLogReasonSoulRuinsSweepDrop:               "帝陵遗迹扫荡掉落物品chapter:%d typ:%d level:%d num:%d",
		InventoryLogReasonSoulRuinsFinishAll:               "帝陵遗迹一键完成chapter:%d typ:%d level:%d num:%d",
		InventoryLogReasonSoulRuinsSweep:                   "帝陵遗迹扫荡消耗物品chapter:%d typ:%d level:%d num:%d",
		InventoryLogReasonSoulRuins:                        "帝陵遗迹挑战奖励chapter:%d typ:%d level:%d",
		InventoryLogReasonSoulRuinsForceGet:                "帝陵遗迹帝魂降临传功奖励chapter:%d typ:%d level:%d",
		InventoryLogReasonXianfuUpgrade:                    "秘境仙府升级,仙府id:%d,类型:%d",
		InventoryLogReasonXianfuSaodang:                    "秘境仙府扫荡消耗,仙府id:%d,类型:%d",
		InventoryLogReasonXianfuSaodangRewards:             "秘境仙府扫荡奖励,仙府id:%d,类型:%d",
		InventoryLogReasonXianfuChallenge:                  "秘境仙府挑战消耗,仙府id:%d,类型:%d",
		InventoryLogReasonRealmToKillRew:                   "天劫塔前往击杀奖励,层数:%d",
		InventoryLogReasonMonsterKilled:                    "杀怪奖励",
		InventoryLogReasonMoonloveGivePresents:             "月下情缘赠送消耗",
		InventoryLogReasonMoonloveTickRew:                  "月下情缘定时获取",
		InventoryLogReasonAllianceCreate:                   "创建仙盟消耗",
		InventoryLogReasonAllianceDonate:                   "仙盟捐献,仙盟id:%d",
		InventoryLogReasonAllianceDonateHuFu:               "仙盟捐献虎符,仙盟id:%d",
		InventoryLogReasonAllianceCallGuard:                "城战召唤守卫,仙盟id:%d,地图id:%d,守卫id:%d",
		InventoryLogReasonAllianceSceneDoorReward:          "城战城门奖励,仙盟id:%d,门:%d",
		InventoryLogReasonAllianceSkillUpgrade:             "仙术升级消耗",
		InventoryLogReasonAllianceConvert:                  "腰牌兑换物品",
		InventoryLogReasonSecretCardStarRew:                "天机牌领取运势箱",
		InventoryLogReasonSecretCardFinishAll:              "天机牌一键完成奖励",
		InventoryLogReasonSecretCardFinish:                 "天机牌任务完成",
		InventoryLogReasonFoundResource:                    "资源找回奖励,资源id:%d,次数:%d",
		InventoryLogReasonDragonFeed:                       "神龙现世喂养",
		InventoryLogReasonDragonAdvancedRew:                "神龙进阶奖励",
		InventoryLogReasonFourGodOpenBox:                   "四神遗迹开启宝箱奖励,钥匙数:%d",
		InventoryLogReasonFourGodBlack:                     "四神遗迹蒙面衣使用",
		InventoryLogReasonFeatherAdvanced:                  "护体仙羽进阶",
		InventoryLogReasonShieldSpikesAdvanced:             "神盾尖刺进阶",
		InventoryLogReasonLingyuEatUn:                      "领域食幻化丹",
		InventoryLogReasonLingyuAdvanced:                   "领域进阶,当前阶数:%d",
		InventoryLogReasonLingyuUnreal:                     "领域幻化,幻化的阶数:%d",
		InventoryLogReasonMarryProposal:                    "求婚扣除婚戒",
		InventoryLogReasonMarryProposalFail:                "求婚失败婚戒返还",
		InventoryLogReasonMarryRingFeedUse:                 "婚戒培养,当前等级:%d",
		InventoryLogReasonMarryTreeFeedUse:                 "爱情树培养,当前等级:%d",
		InventoryLogReasonMarryWedGift:                     "婚宴赠送贺礼",
		InventoryLogReasonMarryRingReplace:                 "婚戒替换",
		InventoryLogReasonMarryGift:                        "结婚贺礼",
		InventoryLogReasonBoxUse:                           "宝箱开启消耗",
		InventoryLogReasonBoxRew:                           "宝箱开启获得",
		InventoryLogReasonGoldEquipStrengthUpgrade:         "元神金装强化升级",
		InventoryLogReasonGoldEquipChongzhuGet:             "元神金装重铸获得",
		InventoryLogReasonGoldEquipChongzhuUse:             "元神金装重铸消耗",
		InventoryLogReasonChessGet:                         "苍龙棋局奖励,类型:%s",
		InventoryLogReasonShenfaEatUn:                      "身法食幻化丹",
		InventoryLogReasonShenfaAdvanced:                   "身法进阶,当前阶数:%d",
		InventoryLogReasonShenfaUnreal:                     "身法幻化,幻化的阶数:%d",
		InventoryLogReasonSaveInDepot:                      "保存到仓库",
		InventoryLogReasonTakeOutDepot:                     "从仓库取出物品",
		InventoryLogReasonOneArenaOutputKun:                "灵池产出鲲",
		InventoryLogReasonOneArenaSellKun:                  "出售鲲",
		InventoryLogReasonChangeScene:                      "切换世界地图,地图:%d",
		InventoryLogReasonRelive:                           "复活,次数:%d",
		InventoryLogReasonRecover:                          "恢复血量",
		InventoryLogReasonFirstCharge:                      "首冲奖励",
		InventoryLogReasonChessAttend:                      "苍龙棋局消耗,棋局类型:%d",
		InventoryLogReasonOpenActivityRew:                  "开服活动奖励,类型:%d,子类型:%d",
		InventoryLogReasonXianfuChallengeSuccessRewards:    "仙府通关奖励,仙府id:%d,类型:%d",
		InventoryLogReasonMajorRew:                         "夫妻双修奖励",
		InventoryLogReasonArena:                            "3v3竞技场,层数:%d",
		InventoryLogReasonArenaExtra:                       "3v3竞技场,层数:%d,额外奖励",
		InventoryLogReasonArenaRelive:                      "3v3竞技场复活,次数:%d",
		InventoryLogReasonFashionUpstar:                    "时装升星消耗",
		InventoryLogReasonBlessDan:                         "祝福丹消耗,类型:%d",
		InventoryLogReasonTransportRew:                     "押镖成功奖励,类型:%d,状态:%d",
		InventoryLogReasonTransportFailRew:                 "押镖失败奖励,类型:%d,状态:%d",
		InventoryLogReasonMountUpstar:                      "坐骑皮肤升星",
		InventoryLogReasonFuncopenRew:                      "功能开启奖励,功能id:%d",
		InventoryLogReasonCollect:                          "采集物奖励",
		InventoryLogReasonGoldEquipEat:                     "吞噬元神金装消耗,是否自动分解：%d",
		InventoryLogReasonVipGiftGet:                       "vip礼包奖励",
		InventoryLogReasonTuMoImmediateFinishRew:           "屠魔任务一键完成奖励",
		InventoryLogReasonVipFreeGiftGet:                   "vip免费礼包奖励",
		InventoryLogReasonWingUpstar:                       "战翼皮肤升星",
		InventoryLogReasonLingYuUpstar:                     "领域皮肤升星",
		InventoryLogReasonShenFaUpstar:                     "身法皮肤升星",
		InventoryLogReasonEmperorRobReward:                 "抢帝王奖励",
		InventoryLogReasonEmperorOpenBox:                   "帝王开宝箱奖励",
		InventoryLogReasonShootFireworks:                   "放烟花,个数:%d",
		InventoryLogReasonSaveInAllianceDepot:              "仙盟仓库存入",
		InventoryLogReasonTakeOutAllianceDepot:             "仙盟仓库取出",
		InventoryLogReasonMassacreAdvanced:                 "戮仙刃升阶消耗，当前%d阶%d星",
		InventoryLogReasonEnterTower:                       "进入打宝塔消耗",
		InventoryLogReasonTianShuUplevel:                   "天书升级消耗,type:%s",
		InventoryLogReasonMyBossUse:                        "个人boss消耗",
		InventoryLogReasonGoldEquipOpenLight:               "金装开光消耗",
		InventoryLogReasonGoldEquipTunShiReturn:            "金装吞噬返还奖励，是否自动分解：%d",
		InventoryLogReasonGoldEquipUpstar:                  "金装升星强化消耗",
		InventoryLogReasonSystemSkillActive:                "系统技能激活,模板id:%d",
		InventoryLogReasonSystemSkillUpgrade:               "系统技能升级,模板id:%d",
		InventoryLogReasonFoeTrackUse:                      "仇人追踪使用",
		InventoryLogReasonDeliverUse:                       "仙盟救援，小飞鞋消耗",
		InventoryLogReasonQuestQuickTransfer:               "任务快速传送,消耗小飞鞋,任务id:%d",
		InventoryLogReasonUnrealBossGet:                    "击杀幻境boss获得",
		InventoryLogReasonFaBaoAdvanced:                    "法宝进阶,当前阶数:%d",
		InventoryLogReasonFaBaoUnreal:                      "法宝幻化,幻化的阶数:%d",
		InventoryLogReasonFaBaoEatUn:                       "法宝食幻化丹",
		InventoryLogReasonFaBaoUpstar:                      "法宝皮肤升星",
		InventoryLogReasonFaBaoTongLing:                    "法宝通灵消耗",
		InventoryLogReasonMaterialChallenge:                "材料副本挑战消耗,type:%s",
		InventoryLogReasonMaterialSaodangUse:               "材料副本扫荡%d次消耗,type:%s",
		InventoryLogReasonMaterialSaodangGet:               "材料副本扫荡%d次奖励,type:%s",
		InventoryLogReasonXueDunEatClu:                     "血盾食培养丹",
		InventoryLogReasonXianTiAdvanced:                   "仙体进阶使用,当前阶数:%d",
		InventoryLogReasonXianTiEatUn:                      "仙体食幻化丹",
		InventoryLogReasonXianTiUnreal:                     "仙体幻化,幻化的阶数:%d",
		InventoryLogReasonXianTiUpstar:                     "仙体皮肤升星",
		InventoryLogReasonXueDunUpgrade:                    "血盾升阶消耗",
		InventoryLogReasonLivenessStarRew:                  "活跃度领取宝箱",
		InventoryLogReasonAdditionSysShengJi:               "系统%s升级使用,当前等级:%d",
		InventoryLogReasonHuiYuanRew:                       "会员每日奖励,类型:%d",
		InventoryLogReasonTianShuRew:                       "领取天书奖励,类型:%d",
		InventoryLogReasonAddFriendRew:                     "添加好友奖励",
		InventoryLogReasonDailyQuestReward:                 "日环任务奖励物品,日环类型：%s",
		InventoryLogReasonDailyQuestFinishAllReward:        "日环任务一键完成奖励",
		InventoryLogReasonBaGuaToKillRew:                   "八卦秘境前往击杀奖励,层数:%d",
		InventoryLogReasonOutlandBossGet:                   "击杀外域boss获得",
		InventoryLogReasonSongBuTingRew:                    "元宝送不停奖励",
		InventoryLogReasonTeamCopyReward:                   "组队副本奖励,副本类型:%d",
		InventoryLogReasonQuizAnswer:                       "仙尊问答奖励,玩家等级:%d ,答题结果:%d",
		InventoryLogReasonPlayerTransfer:                   "玩家场景传送",
		InventoryLogReasonDianXingAdvanced:                 "点星系统升级,当前%d星谱%d星%d进度%d次数",
		InventoryLogReasonFriendNoticeFeedback:             "赞赏好友信息奖励,推送类型:%d,条件:%d",
		InventoryLogReasonFriendGiveGet:                    "好友赠送获取",
		InventoryLogReasonFriendGiveUse:                    "好友赠送消耗",
		InventoryLogReasonDianXingJieFengAdvanced:          "点星解封升级，当前%d阶级%d进度%d次数",
		InventoryLogReasonWardrobeEatDan:                   "衣橱套装%d,配置资质丹",
		InventoryLogReasonTianMoAdvanced:                   "天魔体进阶消耗，阶数:%d",
		InventoryLogReasonTianMoEatDan:                     "天魔体食丹消耗",
		InventoryLogReasonShiHunFanAdvanced:                "噬魂幡进阶使用,当前阶数:%d",
		InventoryLogReasonShiHunFanEatDan:                  "噬魂幡食丹消耗",
		InventoryLogReasonActivityTickRew:                  "活动定时奖励",
		InventoryLogReasonLingTongAdvaced:                  "灵童%s进阶消耗,阶数:%d",
		InventoryLogReasonLingTongDevTongLing:              "灵童%s通灵消耗",
		InventoryLogReasonLingTongDevUnreal:                "灵童%s幻化,幻化的阶数:%d",
		InventoryLogReasonLingTongDevEatUn:                 "灵童%s食幻化丹",
		InventoryLogReasonLingTongDevUpstar:                "灵童%s皮肤升星",
		InventoryLogReasonLingTongDevEatClu:                "灵童%s食培养丹",
		InventoryLogReasonOpenActivityUse:                  "开服活动消耗,类型:%d,子类型:%d",
		InventoryLogReasonLingTongActivate:                 "灵童激活消耗,灵童id：%d",
		InventoryLogReasonLingTongUpgrade:                  "灵童%d升级消耗,当前级数:%d",
		InventoryLogReasonLingTongEatClu:                   "灵童id:%d食培养丹",
		InventoryLogReasonLingTongFashionActivate:          "灵童时装激活消耗,时装id:%d",
		InventoryLogReasonLingTongFashionUpstar:            "灵童时装升星消耗",
		InventoryLogReasonLingTongRename:                   "灵童改名消耗,灵童id：%d",
		InventoryLogReasonKaiFuMuBiaoGroupRew:              "开服目标组奖励,开服解锁天数:%d",
		InventoryLogReasonFeiShenEatDan:                    "飞升食丹,飞升等级:%d",
		InventoryLogReasonShenMoWar:                        "神魔战场领取周排行榜奖励,仙盟排名:%d",
		InventoryLogReasonHongBaoSend:                      "发送红包消耗",
		InventoryLogReasonHongBaoSnatch:                    "抢送红包获得",
		InventoryLogReasonTianFuAwaken:                     "技能:%d,天赋:%d,觉醒",
		InventoryLogReasonTianFuUpgrade:                    "天赋升级,技能:%d,天赋:%d,等级:%d",
		InventoryLogReasonAdditionSysHuaLing:               "系统%s化灵食用,当前等级:%d",
		InventoryLogReasonZhuanShengGiftReceive:            "转生大礼包赠品领取:groupId:%d,gift:%d",
		InventoryLogReasonSupremeTitleActive:               "至尊称号激活,称号id:%d",
		InventoryLogReasonEquipBaoKuAttend:                 "装备宝库消耗,等级：%d，转数：%d",
		InventoryLogReasonEquipBaoKuGet:                    "装备宝库奖励,等级：%d，转数：%d",
		InventoryLogReasonTakeOutMiBaoDepot:                "从秘宝仓库取出物品",
		InventoryLogReasonEquipBaoKuResolveUse:             "装备宝库装备分解消耗",
		InventoryLogReasonEquipBaoKuResolveReturn:          "装备宝库装备分解返还获取",
		InventoryLogReasonGoldEquipExtendUse:               "元神金装继承消耗, 继承等级：%d",
		InventoryLogReasonEquipBaoKuLuckyBoxGet:            "装备宝库幸运宝箱奖励",
		InventoryLogReasonMingGeMosaicUse:                  "命格镶嵌消耗",
		InventoryLogReasonMingGeUnloadAdd:                  "命格卸下获得",
		InventoryLogReasonMingGeSynthesisUse:               "命格合成消耗",
		InventoryLogReasonMingGeRefinedUse:                 "命格祭炼消耗",
		InventoryLogReasonMingGeMingLiBaptizeUse:           "命理洗练消耗",
		InventoryLogReasonMingGeSynthesisAdd:               "命格合成获得",
		InventoryLogReasonTuLongEquipRongHeUse:             "屠龙装备融合消耗",
		InventoryLogReasonTuLongEquipRongHeGet:             "屠龙装备融合结果",
		InventoryLogReasonTuLongEquipStrengthenUse:         "屠龙装备强化消耗",
		InventoryLogReasonTuLongEquipZhuanHuaUse:           "屠龙装备转化消耗",
		InventoryLogReasonTuLongEquipZhuanHuaGet:           "屠龙装备转化结果",
		InventoryLogReasonHuntRew:                          "寻宝奖励，类型：%d",
		InventoryLogReasonHuntUse:                          "寻宝消耗，类型：%d",
		InventoryLogReasonZhenFaActivate:                   "阵法激活消耗",
		InventoryLogReasonZhenFaShengJi:                    "阵法升级消耗",
		InventoryLogReasonZhenQiAdvanced:                   "阵旗进阶消耗",
		InventoryLogReasonZhenFaXianHuoShengJi:             "阵法仙火升级消耗",
		InventoryLogReasonYingLingPuXq:                     "英灵普位置镶嵌,类型[%s],id[%d],等级[%d],碎片[%d]",
		InventoryLogReasonYingLingPuLevelUp:                "英灵普升级,类型[%s],id[%d],等级[%d]",
		InventoryLogReasonShenQiDebrisUpCost:               "神器碎片升级消耗,类型：%s，部位：%s",
		InventoryLogReasonShenQiSmeltUpCost:                "神器淬炼升级消耗,类型：%s，部位：%s,",
		InventoryLogReasonResolveCost:                      "分解消耗",
		InventoryLogReasonMarryPreGift:                     "结婚游车回馈",
		InventoryLogReasonTradeUpload:                      "交易上架,价格[%d]",
		InventoryLogReasonMarryJiNian:                      "纪念奖励,类型[%d]",
		InventoryLogReasonBabyDongFangUse:                  "玩家洞房消耗",
		InventoryLogReasonBabyChangeName:                   "宝宝改名消耗,宝宝id：%d",
		InventoryLogReasonBabyEatTonicUse:                  "宝宝吃补品消耗",
		InventoryLogReasonBabyLearnUse:                     "宝宝读书消耗,babyId:%d, level:%d",
		InventoryLogReasonBabyDongFangReturn:               "洞房失败返还物品",
		InventoryLogReasonBabyEquipToy:                     "装备宝宝玩具消耗",
		InventoryLogReasonMarryDingQingJiHuo:               "定情信物激活",
		InventoryLogReasonBabyToyUplevelUse:                "宝宝玩具升级消耗,suitType:%d,pos:%s,level:%d",
		InventoryLogReasonMarryJiNianShiZhuang:             "结婚纪念时装赠送",
		InventoryLogReasonQiYuRew:                          "奇遇任务奖励，qiyuId:%d",
		InventoryLogReasonBabyRefreshUse:                   "宝宝天赋洗练消耗,babyId:%d,第%d次",
		InventoryLogReasonBabyZhuanShiReturn:               "宝宝转世返还，等级：%d,洗练消耗：%d",
		InventoryLogReasonTradeBuy:                         "交易市场获得,订单id:%d,商品id:%d,花费:%d",
		InventoryLogReasonXianTaoCommit:                    "提交仙桃奖励",
		InventoryLogReasonHouseActivateUse:                 "房子激活消耗，第%d套，类型：%d,等级：%d",
		InventoryLogReasonHouseUplevelUse:                  "房子升级消耗，第%d套，类型：%d,等级：%d",
		InventoryLogReasonHouseRepairUse:                   "房子维修消耗，第%d套，类型：%d,等级：%d",
		InventoryLogReasonHouseUplevelGet:                  "房子升级奖励，第%d套，类型：%d,等级：%d",
		InventoryLogReasonAdditionSysShenZhuCost:           "%s系统%s部位神铸消耗，当前%d等级%d进度%d次数",
		InventoryLogReasonAdditionSysAwake:                 "系统%s觉醒食用",
		InventoryLogReasonShopBuyItem:                      "商城购买shopId:%d,购买个数:%d",
		InventoryLogReasonItemChaiJieCost:                  "物品拆解消耗，物品id:%d",
		InventoryLogReasonItemChaiJieGet:                   "物品拆解获得，物品id:%d",
		InventoryLogReasonAdditionSysTongLingCost:          "%s系统通灵消耗，当前%d等级%d进度%d次数",
		InventoryLogReasonYuXiDayRew:                       "玉玺获胜每日奖励",
		InventoryLogReasonGuideRew:                         "引导副本奖励，类型：%s",
		InventoryLogReasonShenYuRoundRew:                   "神域之战奖励，参赛轮：%d",
		InventoryLogReasonWeedBuyRew:                       "购买周卡奖励，类型：%s",
		InventoryLogReasonWeedDayRew:                       "周卡每日奖励，类型：%s",
		InventoryLogReasonUnlockGem:                        "解锁宝石槽位，位置：%s，槽位：%d",
		InventoryLogReasonFuShiActivite:                    "八卦符石激活，类型：%s",
		InventoryLogReasonFuShiUpLevel:                     "八卦符石升级，类型：%s",
		InventoryLogReasonMajorSaodangUse:                  "夫妻副本扫荡%d次消耗,type:%s,fubenId:%d",
		InventoryLogReasonMajorSaodangGet:                  "夫妻副本扫荡%d次奖励,type:%s,fubenId:%d",
		InventoryLogReasonLingTongUpstar:                   "灵童%d升星消耗,当前星级:%d",
		InventoryLogReasonTitleUpStar:                      "称号升星，称号类型：%d，星级：%d",
		InventoryLogReasonQiXueAdvanced:                    "泣血枪升阶消耗，当前%d阶%d星",
		InventoryLogReasonArenaRankReward:                  "领取3v3周排名奖励,排名:%d",
		InventoryLogReasonBeachShopActivite:                "在%s活动中激活沙滩商店消耗物品",
		InventoryLogReasonBeachShopBuy:                     "购买沙滩商店获得物品，活动type:%d, 活动subType: %d, 商品type: %d",
		InventoryLogReasonBaoMingChuangShiGet:              "玩家报名创世之战获得",
		InventoryLogReasonInviteJieYiUse:                   "玩家邀请结义使用道具, 道具类型：%s, 消耗方式：%s",
		InventoryLogReasonJieYiTokenTypeChangeUse:          "信物类型改变使用，信物类型：%s, 消耗方式：%s",
		InventoryLogReasonJieYiTokenLevelChangeUse:         "信物升级使用,信物类型:%s,信物等级:%d",
		InventoryLogReasonJieYiJiuYuanUse:                  "结义救援使用，小飞鞋消耗",
		InventoryLogReasonArenapvpRew:                      "比武大会奖励，获胜：%v， 类型：%s",
		InventoryLogReasonWushuangWeaponEat:                "无双神器吞噬物品[%s],数量%d",
		InventoryLogReasonWushuangWeaponBreakthrough:       "无双神器突破消耗物品,位置[%s],等级[%d]",
		InventoryLogReasonMiZangCost:                       "密藏花费[%s]",
		InventoryLogReasonMiZangGet:                        "密藏获得[%s]",
		InventoryLogReasonChuangShiChengFangJianShe:        "创世城池建设消耗,数量[%d]",
		InventoryLogReasonChuangShiGuanZhiUse:              "创世之战升职使用,等级:%d",
		InventoryLogReasonChuangShiGuanZhiGet:              "创世之战升职获得,等级:%d",
		InventoryLogReasonChuangShiFashionGet:              "创世之战官职等级达到%d, 领取时装",
		InventoryLogReasonChuangShiJianSheActivateSkillUse: "创世之战建设技能激活消耗，技能等级：%d",
		InventoryLogReasonCastingSpiritUplevel:             "神铸装备铸灵升级消耗物品[%s]，数量%d，部位：%s，铸灵类型：%s",
		InventoryLogReasonWushuangPutOn:                    "无双神器穿戴装备[%s]，部位：%s",
		InventoryLogReasonWushuangTakeOff:                  "无双神器脱下装备[%s]，部位：%s",
		InventoryLogReasonZhenXiCost:                       "进入珍稀boss[%s]",
		InventoryLogReasonGodCastingUpLevel:                "神铸装备升级消耗物品[%s]，数量%s，部位：%s，物品【%s】升级为物品【%s】",
		InventoryLogReasonForgeSoulUplevel:                 "神铸装备锻魂升级消耗物品[%s]，数量%d，部位：%s，锻魂类型：%s",
		InventoryLogReasonGodCastingInherit:                "神铸装备继承消耗物品[%s]，数量%d，部位：%s，物品【%s】继承物品【%s】",
		InventoryLogReasonArenaBossCost:                    "进入圣兽boss",
		InventoryLogReasonMarryDingQingTokenGive:           "结婚定情信物赠送消耗，定情信物：%s",
		InventoryLogReasonTongTianTaLingTongReward:         "领取通天塔-灵童奖励物品，战力：%d",
		InventoryLogReasonTongTianTaMingGeReward:           "领取通天塔-命格奖励物品，战力：%d",
		InventoryLogReasonTongTianTaTuLongReward:           "领取通天塔-屠龙装奖励物品，战力：%d",
		InventoryLogReasonYunYinShopBuy:                    "购买沙滩商店获得物品，活动type:%d, 活动subType: %d, 商品type: %d",
		InventoryLogReasonYunYinShopReward:                 "领取通天塔-灵童奖励物品，花费元宝：%d",
		InventoryLogReasonRewPoolsDrew:                     "仙人指路-奖池抽奖，奖池位置[%d]",
		InventoryLogReasonAdditionSysLingZhuUplevel:        "附加系统灵珠升级，消耗物品，系统类型[%s]，灵珠类型[%s]",
		InventoryLogReasonXianZunCardActiviteAdd:           "激活仙尊特权卡添加物品，类型：%s",
		InventoryLogReasonXianZunCardReceiveAdd:            "领取每日仙尊特权卡奖励，类型%s",
		InventoryLogReasonLingShouUplevel:                  "上古之灵灵兽升级，灵兽类型:[%s]",
		InventoryLogReasonLingWenUplevel:                   "上古之灵灵纹升级，灵兽类型:[%s], 灵纹类型[%s]",
		InventoryLogReasonLingShouUpRank:                   "上古之灵灵兽进阶，灵兽类型:[%s]",
		InventoryLogReasonLingShouLinglian:                 "上古之灵灵兽灵炼，灵兽类型:[%s]",
		InventoryLogReasonRingAdvance:                      "特戒进阶消耗，特戒类型：%s",
		InventoryLogReasonRingStrengthen:                   "特戒强化消耗，特戒类型：%s",
		InventoryLogReasonRingJingLing:                     "特戒净灵消耗，特戒类型：%s",
		InventoryLogReasonRingBaoKuAttend:                  "特戒宝库探索消耗",
		InventoryLogReasonRingBaoKuGet:                     "特戒宝库探索添加",
		InventoryLogReasonRingBaoKuLuckyBoxGet:             "特戒宝库幸运宝箱奖励",
		InventoryLogReasonLingShouReceive:                  "上古之灵领取奖励，灵兽类型:[%s]",
		InventoryLogReasonRingFuseGet:                      "特戒融合获取，特戒类型：%s",
	}
)

func (ilr InventoryLogReason) String() string {
	return inventoryLogReasonMap[ilr]
}
