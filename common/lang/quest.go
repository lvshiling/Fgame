package lang

const (
	QuestNoExist LangCode = QuestBase + iota
	QuestAcceptAuto
	QuestAlreadyAccepted
	QuestStillNotFinish
	QuestBuyTuMoNumReachLimit
	QuestTuMoListReachLimit
	QuestTuMoNumReachLimit
	QuestTuMoQuestIdNotGet
	QuestTuMoBuyNumReachLimit
	QuestTuMoFinishAll
	QuestFinishTitle
	QuestFinishContent
	QuestTuMoFinishAllTitle
	QuestTuMoFinishAllContent
	QuestTuMoImmediateFinishTitle
	QuestTuMoImmediateFinishContent
	QuestTuMoFinishAllVipNoEnough
	QuestDailyEmailTitle
	QuestDailyEmailContent
	QuestDailyFinishAllContent
	QuestCommitContent
	QuestDailyAlreadyFinishAll
	QuestKaiFuMuBiaoGroupReceive
	QuestDailyAllianceFinishChat
	QuestDailyAllianceFinishAllChat
	QuestQiYuReceiveFail
	QuestRewHadReceive
	QuestQiYuEndMailTitle
	QuestQiYuEndMailContent
	QuestDailyAllianceEmailTitle
	QuestDailyAllianceEmailContent
	QuestDailyAlreadyFinishOnce
)

var questLangMap = map[LangCode]string{
	QuestNoExist:                    "任务不存在",
	QuestAcceptAuto:                 "此任务为自动接取",
	QuestAlreadyAccepted:            "任务已经接取",
	QuestStillNotFinish:             "任务还没有完成",
	QuestBuyTuMoNumReachLimit:       "购买屠魔任务次数达上限",
	QuestTuMoListReachLimit:         "屠魔任务列表已满,请先完成并交付",
	QuestTuMoNumReachLimit:          "今日屠魔次数已用完",
	QuestTuMoQuestIdNotGet:          "当前服务器繁忙，建议先完成接取到的屠魔任务",
	QuestTuMoBuyNumReachLimit:       "购买屠魔任务次数已达上限",
	QuestTuMoFinishAll:              "今日屠魔任务已全部完成",
	QuestFinishTitle:                "完成%s任务奖励",
	QuestFinishContent:              "任务奖励内容",
	QuestTuMoFinishAllTitle:         "屠魔任务一键完成奖励",
	QuestTuMoFinishAllContent:       "屠魔任务一键完成奖励内容",
	QuestTuMoImmediateFinishTitle:   "屠魔任务直接完成奖励",
	QuestTuMoImmediateFinishContent: "屠魔任务直接完成奖励内容",
	QuestTuMoFinishAllVipNoEnough:   "玩家达到VIP%s以后,屠魔任务才能一键完成",
	QuestDailyEmailTitle:            "日环任务",
	QuestDailyEmailContent:          "您成功完成了日环任务,但奖励未进行领取,系统自动补发给您,敬请领取",
	QuestCommitContent:              "由于您背包空间不足,无法领取奖励,系统以邮件形式进行发送,敬请查收",
	QuestDailyFinishAllContent:      "由于您背包空间不足,日环任务一键完成,无法领取奖励,系统以邮件形式进行发送,敬请查收",
	QuestDailyAlreadyFinishAll:      "%s任务已经全部完成,直接领取奖励",
	QuestDailyAlreadyFinishOnce:     "%s任务已经完成,直接领取奖励",
	QuestKaiFuMuBiaoGroupReceive:    "开服目标领取组奖励失败",
	QuestDailyAllianceFinishChat:    "仙盟道友%s完成了%s次仙盟日常任务,为仙盟BOSS增加了%s点经验值。",
	QuestDailyAllianceFinishAllChat: "仙盟道友%s一键完成仙盟日常任务,为仙盟BOSS增加了%s点经验值。",
	QuestQiYuReceiveFail:            "条件未达成，无法领取",
	QuestRewHadReceive:              "任务奖励已领取",
	QuestQiYuEndMailTitle:           "奇遇任务奖励",
	QuestQiYuEndMailContent:         "奇遇任务奖励未领取",
	QuestDailyAllianceEmailTitle:    "仙盟日环任务",
	QuestDailyAllianceEmailContent:  "您成功完成了仙盟日环任务,但奖励未进行领取,系统自动补发给您,敬请领取",
}

func init() {
	mergeLang(questLangMap)
}
