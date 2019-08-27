package lang

const (
	OpenActivityHadBuyInvest = OpenActivityBase + iota
	OpenActivityNotBuyInvest
	OpenActivityNotCanReceiveRewards
	OpenActivityNotOnTime
	OpenActivityHadReceiveFirstCharge
	OpenActivityNotEnoughDrewTimes
	OpenActivityNotFirstChargeUser
	OpenActivityDiscountLimit
	OpenActivityDiscountZhuanShengLimit
	OpenActivityFirstChargeRewardsNotice
	OpenActivityCycleChargeRewardsNotice
	OpenActivityInvestLevelRewardsNotice
	OpenActivityInvestDayRewardsNotice
	OpenActivityFeedbackChargeRewardsNotice
	OpenActivityFeedbackCostRewardsNotice
	OpenActivityWelfareRealmRewardsNotice
	MergeActivityAdvancedRewardsNotice
	MergeActivitySingleChargeRewardsNotice
	OpenActivityWelfareGiftCode
	OpenActivityGlobalNotTimesReceive
	OpenActivityGlobalNotTimesReceiveDiscount
	OpenActivityLaBaHadNotTimes
	OpenActivityLaBaRewGoldNotice
	OpenActivityDiscountGiftNotice
	OpenActivityAdvancedPowerRewNotice
	OpenActivityAdvancedRewExtendedNotice
	OpenActivityMailFeedbackGoldPigContent
	OpenActivityDrewRewNotice
	OpenActivityTimesRewContent
	OpenActivityHandlerNotExist
	OpenActivityWelfareUplevelMailContent
	OpenActivityMadeResTimesNotEnough
	OpenActivityCycleTypeNotEqual
	OpenActivityFeedbackDevelopCanNotActivate
	OpenActivityFeedbackDevelopNotDead
	OpenActivityFeedbackDevelopHadDead
	OpenActivityFeedbackDevelopNotFeedTimes
	OpenActivityCostNotEnoughCondition
	OpenActivityFeedbackCycleChargeEndMailContent
	OpenActivityDiscountKanJiaNotEnoughTimes
	OpenActivityFeedbackHouseInvestEndMailContent
	OpenActivityCycleSingleChargeMailContent
	OpenActivityDevelopHadFullTimes
	OpenActivityAllianceCheerMailContent
	OpenActivityAdvancedExpendReturnEndMailContent
	OpenActivityChargeNotEnoughCondition
	OpenActivitySystemHadActivate
	OpenActivityFeedbackDevelopCycleDayMailContent
	OpenActivityAllianceCheerBuyFaild
	OpenActivityFeedbackDevelopFameEndMailContent
	OpenActivityFeedbackDevelopNotActivate
	OpenActivityDiscountGiftUseItemNotice
	OpenActivityBuyFaild
	OpenActivityAdvancedrewTimesReturnEndMailContent
	OpenActivityRankContentNotOnLevel
	OpenActivityFeedbackChargeReturnMultipleEndMailContent
	OpenActivityChargeReturnMultipleRewardLimit
	OpenActivityChargeReturnMultipleChargeNoEnough
	OpenActivityDiscountChargePointNoEnough
	OpenActivityRankContentMarryDevelop
	OpenActivityRankContentMarryDevelopSpouse
	OpenActivityArgumentInvalid
	OpenActivityFeedbackCostContent
	OpenActivityInvestLevelBackGoldLabel
	OpenActivityInvestLevelBackGoldContent
	OpenActivityRankContentMarryDevelopNotOnLevel
	OpenActivityRankContentMarryDevelopSpouseNotOnLevel
	OpenActivityRewardsHadNoneContent
	OpenActivityRewardsPlayerTimesNotEnoughContent
	OpenActivityRankContentCharm
	OpenActivityRankContentCharmSpouse
	OpenActivityRankContentCharmNotOnLevel
	OpenActivityRankContentCharmSpouseNotOnLevel
	OpenActivityFeedbackHouseExtendedEndContent
	OpenActivityEmailCommonGoldString
	OpenActivityEmailCommonSystemString
	OpenActivityEmailCommonSystemLevelString
	OpenActivityEmailCommonDiscountString
	OpenActivityEmailCommonRate
	OpenActivityEmailAdvancedPowerRewardsContent
	OpenActivityEmailAdvancedRewContent
	OpenActivityInvestLevelObjectNotExist
	OpenActivityInvestLevelNotCanUpLev
	OpenActivityEmailUpLevelInvest
	OpenActivityYuanBaoKaPlayerChargeGoldNotEnough
	OpenActivityYuanBaoKaUseNotice
	OpenActivityFeedbackCycleChargeDuanWuContent
	OpenActivityRankContentBoatRaceForce
	OpenActivityRankContentBoatRaceForceNotOnLevel
	OpenActivityGoalMailContent
	OpenActivityDiscountBeachShopActivite
	OpenActivityDiscountBeachShopNotActivite
	OpenActivityDiscountBeachShopReachPurchaseCeiling
	OpenActivityArenapvpRankElection
	OpenActivityArenapvpRankTop32
	OpenActivityArenapvpRankTop16
	OpenActivityArenapvpRankTop8
	OpenActivityArenapvpRankTop4
	OpenActivityArenapvpRankSecond
	OpenActivityArenapvpRankFirst
	OpenActivityFeedbackChargeArenapvpAssistReturnContent
	OpenActivityFeedbackChargeArenapvpAssistReturnContentExtral
	OpenActivityTongTianTaNotReceive
	OpenActivityTongTianTaAlreadyReceive
	OpenActivityObjectNotExist
	OpenActivityTongTianTaNotReachGoal
	OpenActivityNewSevenInvestTypeWorry
	OpenActivityNewSevenInvestNotEnoughMaxSigle
	OpenActivityHallZhuanShengEndContent
	OpenActivityNewSevenDayInvestEmail
	OpenActivityEmailTongTianTa
	OpenActivityEmailYunYinShop
	OpenActivityYunYinShopNotEnough
)

var (
	welfareLangMap = map[LangCode]string{
		OpenActivityHadBuyInvest:                  "已购买投资计划",
		OpenActivityNotBuyInvest:                  "未购买投资计划",
		OpenActivityNotCanReceiveRewards:          "没有可领取的奖励",
		OpenActivityNotOnTime:                     "活动未开启",
		OpenActivityHadReceiveFirstCharge:         "首充已领取",
		OpenActivityNotFirstChargeUser:            "不是首充用户",
		OpenActivityNotEnoughDrewTimes:            "抽奖次数不足",
		OpenActivityDiscountLimit:                 "超过折扣礼包每日购买上限",
		OpenActivityDiscountZhuanShengLimit:       "超过转生大礼包购买上限",
		OpenActivityWelfareGiftCode:               "激活码不存在",
		OpenActivityGlobalNotTimesReceive:         "全服剩余份数被其他玩家抢光了，无法领取该奖励",
		OpenActivityGlobalNotTimesReceiveDiscount: "您的手太慢了，本商品已售罄",
		OpenActivityLaBaHadNotTimes:               "拉霸次数已达上限",
		OpenActivityHandlerNotExist:               "处理器没有注册",
		OpenActivityMadeResTimesNotEnough:         "您本日经验炼制次数已达上限，请明日再来",
		OpenActivityCycleTypeNotEqual:             "类型不一致",
		OpenActivityFeedbackDevelopCanNotActivate: "金鸡未满足激活条件",
		OpenActivityFeedbackDevelopNotActivate:    "金鸡未激活",
		OpenActivityFeedbackDevelopNotDead:        "金鸡未死亡",
		OpenActivityFeedbackDevelopHadDead:        "金鸡已经死亡",
		OpenActivityFeedbackDevelopNotFeedTimes:   "没有喂养次数",
		OpenActivityCostNotEnoughCondition:        "不满足消费条件",
		OpenActivityDevelopHadFullTimes:           "本物品贡献次数已达上限，无法继续贡献",
		OpenActivityChargeNotEnoughCondition:      "不满足充值条件",
		OpenActivitySystemHadActivate:             "系统已激活",
		OpenActivityAllianceCheerBuyFaild:         "请在城战开启前购买",
		OpenActivityBuyFaild:                      "购买失败",
		OpenActivityArgumentInvalid:               "活动参数无效：groupId:%s",
		OpenActivityInvestLevelObjectNotExist:     "活动对象不存在",
		OpenActivityInvestLevelNotCanUpLev:        "投资计划升级未满足条件",

		OpenActivityMailFeedbackGoldPigContent:           "恭喜您在%s活动中累计充值%s，激活%s的元宝返还比例，您在活动期间内总共消费%s，您当前一共获得%s的奖励，敬请领取",
		OpenActivityTimesRewContent:                      "恭喜您在%s活动中累计抽奖%d次，本次活动中，您的VIP等级%d，获得如下奖励，敬请领取",
		OpenActivityWelfareUplevelMailContent:            "您在%s活动中等级达到%d，获得如下奖励，敬请查收",
		OpenActivityFeedbackCycleChargeEndMailContent:    "您在%s活动中存在以下未领取奖励，敬请查收",
		OpenActivityDiscountKanJiaNotEnoughTimes:         "没有砍价次数",
		OpenActivityFeedbackHouseInvestEndMailContent:    "由于%s活动已结束，系统自动将您的房子卖出，您获得了如下收益，敬请查收",
		OpenActivityCycleSingleChargeMailContent:         "恭喜您在%s活动期间内单笔充值金额达到%s，获得如下奖励，敬请领取",
		OpenActivityAllianceCheerMailContent:             "恭喜您在%s活动中您所助威仙盟获胜，您在活动中购买礼包金额累计消费%s，系统返还%s给您，以下为您的返还元宝，敬请领取",
		OpenActivityAdvancedExpendReturnEndMailContent:   "您在%s活动中未领取的奖励",
		OpenActivityFeedbackDevelopCycleDayMailContent:   "您在%s活动中未领取的奖励",
		OpenActivityFeedbackDevelopFameEndMailContent:    "恭喜您在%s活动中与%s好感度达到%s，您获得如下奖励，敬请查收",
		OpenActivityAdvancedrewTimesReturnEndMailContent: "恭喜您在%s活动期间内，成功将%s系统进阶%s次，获得如下奖励，敬请查收",
		OpenActivityRankContentNotOnLevel:                "恭喜您在%s排行榜活动中排名达到第%s名，由于您当前不满足获得本名次奖励的条件，仅能获得满足条件的奖励，敬请查收",
		OpenActivityFeedbackCostContent:                  "恭喜您在%s活动期间内累计消费金额达到%s，可领取%s消费奖励，以下为您的消费奖励，敬请领取",
		OpenActivityRewardsHadNoneContent:                "%s%d元宝的奖励全服剩余份数已为0，无法获得奖励",
		OpenActivityRewardsPlayerTimesNotEnoughContent:   "%s%d元宝的个人可奖励次数已为0，无法获得奖励",
		OpenActivityFeedbackHouseExtendedEndContent:      "恭喜您在%s活动期间内充值金额达到%s，获得如下奖励，敬请领取",
		OpenActivityHallZhuanShengEndContent:             "您在%s活动中转数达到%s，获得如下奖励，敬请查收",

		OpenActivityFirstChargeRewardsNotice:    "一掷千金，%s在首充奖励中领取%s,战力暴涨8288",
		OpenActivityCycleChargeRewardsNotice:    "恭喜%s在%s中累计充值达%s，成功领取%s %s",
		OpenActivityInvestLevelRewardsNotice:    "精打细算，%s加入等级投资，元宝轻松翻一番！%s",
		OpenActivityInvestDayRewardsNotice:      "精打细算，%s加入7日投资，7天元宝翻五倍！%s",
		OpenActivityFeedbackChargeRewardsNotice: "恭喜%s在%s中累计充值%s，领取%s %s",
		OpenActivityFeedbackCostRewardsNotice:   "%s累计消费%s，领取%s %s",
		OpenActivityWelfareRealmRewardsNotice:   "过关斩将，%s突破%s，领取%s %s",
		MergeActivityAdvancedRewardsNotice:      "%s在升阶返利活动中领取了%s %s",
		MergeActivitySingleChargeRewardsNotice:  "%s单笔充值达到%s，在活动中领取%s %s",
		OpenActivityLaBaRewGoldNotice:           "鸿运当头，恭喜%s在%s中投入%s，获得%s奖励 %s",
		OpenActivityDiscountGiftNotice:          "天降大礼，%s在%s中花费%s，成功购买%s，价值%s，限时%s，走过路过，不要错过 %s",
		OpenActivityAdvancedPowerRewNotice:      "恭喜%s在%s中%s战力达%s，免费领取酷炫称号%s %s",
		OpenActivityAdvancedRewExtendedNotice:   "恭喜%s在%s中%s阶别达到%s，免费领取%s %s",
		OpenActivityDrewRewNotice:               "天降洪福，%s在%s中受幸运女神眷顾，获得%s %s",
		OpenActivityDiscountGiftUseItemNotice:   "天降大礼，%s在%s中花费%s，成功购买%s，限时销售，走过路过，不要错过 %s",

		OpenActivityFeedbackChargeReturnMultipleEndMailContent: "恭喜您在%s活动中累计充值%s元宝,活动期间内每充值100元宝赠送奖品，敬请领取",
		OpenActivityChargeReturnMultipleRewardLimit:            "已达领取上限",
		OpenActivityChargeReturnMultipleChargeNoEnough:         "充值不足",
		OpenActivityDiscountChargePointNoEnough:                "当前充值积分不足",

		OpenActivityRankContentMarryDevelop:                         "恭喜您在%s中累计获得表白经验%d，您在本次活动中排名第%s名，获得丰厚奖励，您的伴侣也将获得同样奖励，敬请查收",
		OpenActivityRankContentMarryDevelopSpouse:                   "由于您的伴侣在%s中累计获得表白经验%d，在活动中排名第%s名，您同时获得了奖励，敬请查收",
		OpenActivityInvestLevelBackGoldLabel:                        "投资计划更新",
		OpenActivityInvestLevelBackGoldContent:                      "亲爱的玩家，由于投资计划进行更新，系统自动将您购买所花费的元宝全部进行返还，请重新前往进行投资，给您带来的不便，敬请原谅",
		OpenActivityRankContentMarryDevelopNotOnLevel:               "恭喜您在%s中累计获得表白经验%d，您在本次活动中排名第%s名，由于您当前不满足本名次奖励的条件，仅能获得满足条件的奖励，您的伴侣也将获得同样奖励，敬请查收",
		OpenActivityRankContentMarryDevelopSpouseNotOnLevel:         "由于您的伴侣在%s中累计获得表白经验%d，在活动中排名第%s名，由于不满足本名次奖励的条件，仅能获得满足条件的奖励，您同时获得了奖励，敬请查收",
		OpenActivityRankContentCharm:                                "恭喜您在%s中累计获得魅力值%d，您在本次活动中排名第%s名，获得丰厚奖励，您的伴侣也将获得同样奖励，敬请查收",
		OpenActivityRankContentCharmSpouse:                          "由于您的伴侣在%s中累计获得魅力值%d，在活动中排名第%s名，您同时获得了奖励，敬请查收",
		OpenActivityRankContentCharmNotOnLevel:                      "恭喜您在%s中累计获得魅力值%d，您在本次活动中排名第%s名，由于您当前不满足本名次奖励的条件，仅能获得满足条件的奖励，您的伴侣也将获得同样奖励，敬请查收",
		OpenActivityRankContentCharmSpouseNotOnLevel:                "由于您的伴侣在%s中累计获得魅力值验%d，在活动中排名第%s名，由于不满足本名次奖励的条件，仅能获得满足条件的奖励，您同时获得了奖励，敬请查收",
		OpenActivityEmailCommonGoldString:                           "%d元宝",
		OpenActivityEmailCommonSystemString:                         "%s系统",
		OpenActivityEmailCommonSystemLevelString:                    "%d阶",
		OpenActivityEmailCommonDiscountString:                       "%d折",
		OpenActivityEmailCommonRate:                                 "%d.00%%",
		OpenActivityEmailAdvancedPowerRewardsContent:                "恭喜您在%s活动中战斗力成功达到%d，获得如下奖励，敬请领取",
		OpenActivityEmailAdvancedRewContent:                         "恭喜您在%s活动中成功将%s进阶到第%d阶，获得如下奖励，敬请领取",
		OpenActivityEmailUpLevelInvest:                              "恭喜您成功把%s升级到%s，现将原本该给的奖励补发给您，敬请领取",
		OpenActivityYuanBaoKaPlayerChargeGoldNotEnough:              "玩家本日充值元宝未达到使用元宝卡条件",
		OpenActivityYuanBaoKaUseNotice:                              "恭喜%s成功使用%s，成功获得%s元宝 %s",
		OpenActivityFeedbackCycleChargeDuanWuContent:                "恭喜您在%s活动中获得如下奖励，敬请领取",
		OpenActivityRankContentBoatRaceForce:                        "恭喜您在%s活动中角色战力值增长%s，您本次在活动中排名%s，以下为您的排名奖励，敬请领取",
		OpenActivityRankContentBoatRaceForceNotOnLevel:              "恭喜您在%s活动中角色战力值增长%s，您本次在活动中排名%s，由于您增长值不满足本名次奖励，因此仅能获得满足条件最近档次的一个奖励，以下为您的排名奖励，敬请查收",
		OpenActivityGoalMailContent:                                 "恭喜您在%s活动中完成对应目标，获得如下奖励，敬请查收",
		OpenActivityDiscountBeachShopActivite:                       "该沙滩商店已经激活",
		OpenActivityDiscountBeachShopNotActivite:                    "该沙滩商店未激活",
		OpenActivityDiscountBeachShopReachPurchaseCeiling:           "该商品共已售罄",
		OpenActivityArenapvpRankElection:                            "海选",
		OpenActivityArenapvpRankTop32:                               "32强",
		OpenActivityArenapvpRankTop16:                               "16强",
		OpenActivityArenapvpRankTop8:                                "8强",
		OpenActivityArenapvpRankTop4:                                "4强",
		OpenActivityArenapvpRankSecond:                              "第二名",
		OpenActivityArenapvpRankFirst:                               "第一名",
		OpenActivityFeedbackChargeArenapvpAssistReturnContent:       "恭喜您在%s活动中参与比武大会名次达到%s，可享受返还绑元比例%s，活动期间内您累计消费元宝数量%s，以下为您的奖励，敬请查收%s",
		OpenActivityFeedbackChargeArenapvpAssistReturnContentExtral: "（返还绑元数量上限为%d）",
		OpenActivityTongTianTaNotReceive:                            "不能领取该奖励",
		OpenActivityTongTianTaAlreadyReceive:                        "已经领取过该奖励了",
		OpenActivityObjectNotExist:                                  "活动对象不存在",
		OpenActivityTongTianTaNotReachGoal:                          "活动期间累计充值金额达到%s元宝即可免费领取",
		OpenActivityNewSevenInvestTypeWorry:                         "投资类型错误",
		OpenActivityNewSevenInvestNotEnoughMaxSigle:                 "您在活动期间的最大单笔充值未达到购买条件",
		OpenActivityNewSevenDayInvestEmail:                          "您在%s活动中未领取的奖励现已通过邮件发给您，敬请领取",
		OpenActivityEmailTongTianTa:                                 "恭喜您在%s活动中获得如下奖励，敬请领取",
		OpenActivityEmailYunYinShop:                                 "您在%s活动中未领取的奖励现已通过邮件补发给您，敬请领取",
		OpenActivityYunYinShopNotEnough:                             "云隐商店商品不足",
	}
)

func init() {
	mergeLang(welfareLangMap)
}
