package lang

const (
	ArenaMailTitle LangCode = ArenaBase + iota
	ArenaMailContent
	ArenaInvite
	ArenaInviteCD
	ArenaSpecialReward
	ArenaExpTreeCollect
	ArenaOtherIsCollecting
	AreanaCollectBoxGongGao
	ArenaGetRewardNoInRank
	ArenaGetRewardHasedGet
	ArenaMailTitleFail
	ArenaMailContentFail
)

var arenaLangMap = map[LangCode]string{
	ArenaMailTitle:          "3v3竞技场",
	ArenaMailContent:        "3v3竞技场",
	ArenaInvite:             "[%s]战队邀请大家参与3V3跨服战，获胜可得大量银两以及极品道具，最终场景更可获得极品稀有装备%s",
	ArenaInviteCD:           "3v3吆喝CD中",
	ArenaSpecialReward:      "恭喜%s战队获胜轮数达到%s，系统随机给予该战队玩家%s3V3跨服宝箱奖励，可随机开出稀有装备以及极品道具",
	ArenaExpTreeCollect:     "恭喜%s成功采集经验树，其所在战队“%s的战队”所有成员经验获得大涨！已经退出战队的玩家不再享受此加成",
	ArenaOtherIsCollecting:  "当前有人正在采集,请稍后再试",
	AreanaCollectBoxGongGao: "%s圣兽已被击杀，恭喜%s成功采集稀有装备，获得极品稀有装备%s",
	ArenaGetRewardNoInRank:  "您的连胜排名未在榜单内,无法领取奖励",
	ArenaGetRewardHasedGet:  "您已经领过周排行奖励了",
	ArenaMailTitleFail:      "3V3战败",
	ArenaMailContentFail:    "由于你从3V3竞技场中提前退出，鉴于你的行为，在上一场3V3战斗中你被判定为战败！",
}

func init() {
	mergeLang(arenaLangMap)
}
