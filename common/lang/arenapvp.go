package lang

const (
	ArenapvpElectionEnd LangCode = ArenapvpBase + iota
	ArenapvpGuessFail
	ArenapvpGuessMailTitle
	ArenapvpGuessSuccessMailContent
	ArenapvpGuessFailedMailContent
	ArenapvpElectionWinMailContent
	ArenapvpElectionFailedMailTitle
	ArenapvpElectionFailedMailContent
	ArenapvpGuessHadAttend
	ArenapvpBattleFinishMailTitle
	ArenapvpGuessReturnMailTitle
	ArenapvpGuessReturnMailContent
	ArenapvpLuckyRewMailTitle
	ArenapvpLuckyRewMailContent
	ArenapvpLuckySystemContent
	ArenapvpWinnerTOP32Content
	ArenapvpWinnerTOP16Content
	ArenapvpWinnerTOP8Content
	ArenapvpWinnerTOP4Content
	ArenapvpWinnerFinalsContent
	ArenapvpWinnerChampionContent
	ArenapvpLuckyIndexText
	ArenapvpLuckyRewNameText
	ArenapvpAlreadyBuyTicket
)

var arenapvpLangMap = map[LangCode]string{
	ArenapvpElectionEnd:               "比武大会海选已经结束",
	ArenapvpGuessFail:                 "竞猜时间已过",
	ArenapvpGuessMailTitle:            "竞猜奖励",
	ArenapvpGuessSuccessMailContent:   "恭喜您成功在%s中，猜对胜利者，这是您的奖励请查收！",
	ArenapvpGuessFailedMailContent:    "很遗憾，您在%s中，猜错胜利者，没能获得奖励，请再接再厉！",
	ArenapvpElectionWinMailContent:    "海选晋级奖励",
	ArenapvpElectionFailedMailTitle:   "比武大会",
	ArenapvpElectionFailedMailContent: "很遗憾，您的积分没有排名前4，无法进入32强。已送您参与奖：%d积分，敬请到活动界面积分商城进行兑换。",
	ArenapvpGuessHadAttend:            "已经参与竞猜",
	ArenapvpBattleFinishMailTitle:     "比武大会-淘汰赛",
	ArenapvpGuessReturnMailTitle:      "竞猜退还",
	ArenapvpGuessReturnMailContent:    "竞猜退还邮件",
	ArenapvpLuckyRewMailTitle:         "比武大会幸运奖",
	ArenapvpLuckyRewMailContent:       "恭喜您在比武大会海选赛中成功中奖，获得%s，奖励如下，请查收！",
	ArenapvpLuckySystemContent:        "恭喜%s的%s在比武大会海选积分赛中成功获得%s",
	ArenapvpLuckyIndexText:            "第%d会场",
	ArenapvpLuckyRewNameText:          "幸运奖",
	ArenapvpWinnerTOP32Content:        "在一番激烈的积分争夺下，恭喜以下玩家晋级32强：%s",
	ArenapvpWinnerTOP16Content:        "在1v1淘汰赛的惨烈战斗中，恭喜以下玩家晋级16强：%s",
	ArenapvpWinnerTOP8Content:         "在1v1淘汰赛的惨烈战斗中，恭喜以下玩家晋级8强：%s",
	ArenapvpWinnerTOP4Content:         "在4轮的角逐中，%s以黑马的姿态杀入了半决赛",
	ArenapvpWinnerFinalsContent:       "在经历了5轮的激烈角逐后，恭喜%s获得争夺顶峰的资格！",
	ArenapvpWinnerChampionContent:     "在双方的激烈拼杀下，%s轻松一招夺得了冠军，拿下“千古第一人”头衔，这届比武大会就此落下帷幕，欢迎广大玩家下次参与！",
	ArenapvpAlreadyBuyTicket:          "已经购买过比武大会门票",
}

func init() {
	mergeLang(arenapvpLangMap)
}
