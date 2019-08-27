package lang

const (
	ShenMoCheckFuncOpen LangCode = ShenMoBase + iota
	ShenMoInLineUpEnterFuBen
	ShenMoCancleLineUpNoExist
	ShenMoGetMyRankNoAlliance
	ShenMoGetRewardNoInRank
	ShenMoGetRewardHasedGet
	ShenMoGetRewardTitle
	ShenMoGetRewardContent
	ShenMoGetRewardNoEnoughTime
)

var (
	shenMoLangMap = map[LangCode]string{
		ShenMoCheckFuncOpen:         "您当前未开启神魔战场功能,无法进入活动",
		ShenMoInLineUpEnterFuBen:    "您当前正在神魔战场排队中,无法进入副本",
		ShenMoCancleLineUpNoExist:   "您当前未在排队中",
		ShenMoGetMyRankNoAlliance:   "您当前还未加入仙盟,无法领取奖励",
		ShenMoGetRewardNoInRank:     "您的仙盟排名未在榜单内,无法领取奖励",
		ShenMoGetRewardHasedGet:     "您已经领过周排行奖励了",
		ShenMoGetRewardTitle:        "神魔战场周排行",
		ShenMoGetRewardContent:      "由于您的背包空间不足,领取的周排行奖励通过邮件下发,敬请查收!",
		ShenMoGetRewardNoEnoughTime: "入盟未满三天,无法领取奖励",
	}
)

func init() {
	mergeLang(shenMoLangMap)
}
