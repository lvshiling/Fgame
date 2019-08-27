package lang

const (
	EmailNoReceiveAttachment LangCode = EmailBase + iota
	EmailNoAttachment
	EmailRepeatedRead
	EmailCharmName
	EmailGenerousName
	EmailGenerouseContent
	EmailCharmContent
	EmailMoonloveTickName
	EmailMoonloveTickConten
	EmailAllianceTransportationFinishName
	EmailAllianceTransportationFinishContent
	EmailAllianceTransportationRobName
	EmailAllianceTransportationRobContent
	EmailAllianceTransportationBeRobTitle
	EmailAllianceTransportationBeRobContent
	EmailOpenActivityRankTitle
	EmailOpenActivityRankContent
	EmailOpenActivityFeedbackCycleChargeTitle
	EmailOpenActivityFeedbackCycleChargeContent
	EmailOpenActivityCycleChargeContent
	EmailOpenActivityHuiYuanTitle
	EmailOpenActivityHuiYuanContent
	EmailOpenActivityRealmTitle
	EmailOpenActivityRealmContent
	EmailOpenActivitySingleChargeTitle
	EmailOpenActivitySingleChargeContent
	EmailOpenActivityGoldBowlContent
	EmailOpenActivityAdvancedTitle
	EmailOpenActivityAdvancedContent
	EmailOpenActivityFeedbackChargeContent
	EmailOpenActivityAdvancedBlessTitle
	EmailOpenActivityAdvancedBlessContent
	EmailFuncOpenRewTitle
	EmailFuncOpenRewContent
	EmailXianfuChallengeRewTitle
	EmailXianfuChallengeRewContent
	EmailTianShuBindGoldFeedbackTitle
	EmailTianShuBindGoldFeedbackContent
	EmailTianShuGoldFeedbackTitle
	EmailTianShuGoldFeedbackContent
	EmailOpenActivityLongFengRobTitle
	EmailOpenActivityLongFengRobContent
	EmailOpenActivityChargeRewardsContent
	EmailOpenActivityCostRewardsContent
	EmailOpenActivityFeedbackChargeLimitContent
	EmailOpenActivityFeedbackChargeRewContent
	EmailOpenActivityAdvancedRewContent
	EmailInventorySlotNoEnough
	EmailOutlandBossDropItems
	EmailSystemRewardTitle
	EmailPrivilegeChargeContent
	EmailEquipBaoKuLuckyBoxTitle
	EmailEquipBaoKuLuckyBoxContent
	EmailJieYiDaoJuReturnTitle
	EmailJieYiDaoJuReturnContent
	EmailJieYiDaoJuFashionTitle
	EmailJieYiDaoJuFashionContent
	EmailOpenActivityFeedbackChargeArenapvpAssistContent
	EmailOpenActivityXiuXianBook
	EmailTitleTimeExpireTitle
	EmailTitleTimeExpireContent
	EmailXianZunCardCrossTitle
	EmailXianZunCardCrossContent
)

var (
	emailLangMap = map[LangCode]string{
		//------------------邮件-------------------
		EmailNoReceiveAttachment:                             "尚有附件未领取，请先领取附件",
		EmailNoAttachment:                                    "该邮件没有附件",
		EmailRepeatedRead:                                    "邮件已读",
		EmailCharmName:                                       "月下情缘魅力榜",
		EmailGenerousName:                                    "月下情缘土豪榜",
		EmailGenerouseContent:                                "恭喜您获得%s,送您一个超级大礼包",
		EmailCharmContent:                                    "恭喜您获得%s,送您一个超级大礼包",
		EmailMoonloveTickName:                                "月下情缘奖励",
		EmailMoonloveTickConten:                              "月下情缘定时奖励",
		EmailAllianceTransportationFinishName:                "仙盟镖车奖励",
		EmailAllianceTransportationFinishContent:             "成功押送仙盟镖车,仙盟全体成员获得丰厚奖励!",
		EmailAllianceTransportationRobName:                   "仙盟劫镖奖励",
		EmailAllianceTransportationRobContent:                "成功劫获%s仙盟镖车,仙盟全体成员获得丰厚奖励!",
		EmailAllianceTransportationBeRobTitle:                "仙盟押镖失败",
		EmailAllianceTransportationBeRobContent:              "仙盟镖车被%s劫走，未能完成押镖任务，损失惨重！",
		EmailOpenActivityRankTitle:                           "排行榜活动",
		EmailOpenActivityRankContent:                         "恭喜您在%s活动中排名第%s名，获得以下奖励，敬请领取",
		EmailOpenActivityFeedbackCycleChargeTitle:            "连续充值",
		EmailOpenActivityFeedbackCycleChargeContent:          "连续充值奖励邮件",
		EmailOpenActivityCycleChargeContent:                  "恭喜您在%s活动期间内充值金额达到%s，获得如下奖励，敬请领取",
		EmailOpenActivityHuiYuanTitle:                        "会员特权",
		EmailOpenActivityHuiYuanContent:                      "会员每日奖励邮件",
		EmailOpenActivityRealmTitle:                          "天劫塔冲刺",
		EmailOpenActivityRealmContent:                        "天劫塔冲刺未领取奖励邮件",
		EmailOpenActivitySingleChargeTitle:                   "单笔充值活动",
		EmailOpenActivitySingleChargeContent:                 "单笔充值活动未领取奖励邮件",
		EmailOpenActivityGoldBowlContent:                     "恭喜您在%s活动期间中累计充值消费%s，活动期间内返还率%s，您本次获得%s，敬请领取",
		EmailOpenActivityAdvancedTitle:                       "升阶返利活动",
		EmailOpenActivityAdvancedContent:                     "升阶返利未领取奖励邮件",
		EmailOpenActivityFeedbackChargeContent:               "恭喜您在%s活动期间内累计充值金额达到%s，可领取%s累充奖励，以下为您的累充奖励，敬请领取",
		EmailOpenActivityAdvancedBlessTitle:                  "祝福丹大派送",
		EmailOpenActivityAdvancedBlessContent:                "祝福丹大派送未领取奖励邮件",
		EmailFuncOpenRewTitle:                                "功能开启",
		EmailFuncOpenRewContent:                              "功能开启奖励邮件",
		EmailXianfuChallengeRewTitle:                         "秘境仙府挑战奖励",
		EmailXianfuChallengeRewContent:                       "秘境仙府挑战奖励邮件",
		EmailTianShuBindGoldFeedbackTitle:                    "绑元天书特权",
		EmailTianShuBindGoldFeedbackContent:                  "绑元天书充值返利",
		EmailTianShuGoldFeedbackTitle:                        "元宝天书特权",
		EmailTianShuGoldFeedbackContent:                      "元宝天书充值返利",
		EmailOpenActivityLongFengRobTitle:                    "龙凤呈祥奖励",
		EmailOpenActivityLongFengRobContent:                  "恭喜您成功抢夺龙椅，龙凤呈祥活动期间每抢夺一次龙椅将额外获取【祥龙瑞凤】时装奖励，敬请领取",
		EmailOpenActivityChargeRewardsContent:                "恭喜您在%s活动中累计充值%s，活动期间内每充值%d元宝赠送奖品，敬请领取",
		EmailOpenActivityCostRewardsContent:                  "恭喜您在%s活动中累计消费%s，活动期间内每消费%d元宝赠送奖品，敬请领取",
		EmailOpenActivityFeedbackChargeLimitContent:          "恭喜您在%s活动期间内累计充值金额达到%s，获得如下奖励，敬请领取",
		EmailOpenActivityFeedbackChargeRewContent:            "恭喜您在%s活动中成功将%s进阶到第%d阶，获得如下奖励，敬请领取",
		EmailOpenActivityAdvancedRewContent:                  "恭喜您在%s活动中成功将%s进阶到第%d阶，获得如下奖励，敬请领取",
		EmailInventorySlotNoEnough:                           "背包空间不足",
		EmailOutlandBossDropItems:                            "外域BOSS掉落",
		EmailSystemRewardTitle:                               "系统奖励邮件",
		EmailPrivilegeChargeContent:                          "这是系统奖励您的元宝，请接收。",
		EmailEquipBaoKuLuckyBoxTitle:                         "装备宝库幸运宝箱",
		EmailEquipBaoKuLuckyBoxContent:                       "您在装备宝库内的幸运宝箱，未及时抽取。现已帮您兑换成奖励发送给您，敬请查收",
		EmailJieYiDaoJuReturnTitle:                           "结义失败",
		EmailJieYiDaoJuReturnContent:                         "由于对方拒绝了您的结义邀请，系统自动将您使用的结义道具以邮件形式返还给您，敬请查收",
		EmailJieYiDaoJuFashionTitle:                          "结义时装",
		EmailJieYiDaoJuFashionContent:                        "恭喜您的结义道具达到高级，获得酷炫结义时装，敬请查收",
		EmailOpenActivityFeedbackChargeArenapvpAssistContent: "恭喜您在【擂台助力】活动期间内累计充值金额达到%s，以下为您的奖励，敬请查收",
		EmailOpenActivityXiuXianBook:                         "恭喜您在%s活动期间达成指定条件，获得如下奖励，敬请领取",
		EmailTitleTimeExpireTitle:                            "限时称号过期提醒",
		EmailTitleTimeExpireContent:                          "您的限时称号【%s】在%d年%d月%d日%d点%d分过期，您的战斗力下降【%s】点",
		EmailXianZunCardCrossTitle:                           "%s特权",
		EmailXianZunCardCrossContent:                         "以下是您%s特权卡的每日奖励，由于您当天未手动领取，过天后给你补发，请查收！",
	}
)

func init() {
	mergeLang(emailLangMap)
}
