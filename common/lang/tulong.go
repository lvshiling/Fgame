package lang

const (
	TuLongUnionEachPeople LangCode = TuLongBase + iota
	TuLongCollectNoMengZhu
	TuLongHasedCollectEgg
	TuLongOtherCollect
	TuLongCollectEggShouldExist
	TuLongKillBossTitle
	TuLongKillBossContent
	TuLongBossBorn
	TuLongSmallBossBorn
	TuLongAllianceAreaRankNoEnough
)

var (
	tuLongLangMap = map[LangCode]string{
		TuLongUnionEachPeople:          "您所在仙盟参与的人数已达上限",
		TuLongCollectNoMengZhu:         "你是普通成员,无法开启龙蛋",
		TuLongHasedCollectEgg:          "您今日已开启过龙蛋,无法再次开启",
		TuLongOtherCollect:             "其它玩家正在采集",
		TuLongCollectEggShouldExist:    "采集的龙蛋应该是存在的",
		TuLongKillBossTitle:            "跨服屠龙",
		TuLongKillBossContent:          "恭喜您所在仙盟%s在跨服屠龙中成功击杀%s,您同时获得了奖励",
		TuLongBossBorn:                 "%s已经复活！大家赶紧前往%s抢夺Boss吧",
		TuLongSmallBossBorn:            "%s在%s处偷偷开了一只小龙！大家赶紧去抢",
		TuLongAllianceAreaRankNoEnough: "所在仙盟区排名不满足条件",
	}
)

func init() {
	mergeLang(tuLongLangMap)
}
