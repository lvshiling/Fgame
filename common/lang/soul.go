package lang

const (
	SoulRepeatActive LangCode = SoulBase + iota
	SoulActiveNotPreCond
	SoulNotActiveNotAwaken
	SoulNotActiveNotUpgrade
	SoulUpgradeReackLimit
	SoulNotActiveNotEmbed
	SoulRepeatEmbed
	SoulNotActiveNotFeed
	SoulReacheFullLevel
	SoulFeedNotItem
	SoulNotActiveNotStrengthen
	SoulStrengthenReackLimit
	SoulAwakenIsExist
	SoulActivateNotice
	SoulAwakenNotice
)

var (
	soulLangMap = map[LangCode]string{
		SoulRepeatActive:           "帝魂已激活,无需激活",
		SoulActiveNotPreCond:       "激活该帝魂的前置条件不足",
		SoulNotActiveNotAwaken:     "未激活的帝魂,无法觉醒",
		SoulNotActiveNotUpgrade:    "未激活的帝魂,无法升级",
		SoulUpgradeReackLimit:      "帝魂升级已达最高阶",
		SoulNotActiveNotEmbed:      "未激活的帝魂,无法镶嵌",
		SoulRepeatEmbed:            "该帝魂已镶嵌,重复镶嵌",
		SoulNotActiveNotFeed:       "未激活的帝魂,无法喂养",
		SoulReacheFullLevel:        "帝魂已满级",
		SoulFeedNotItem:            "当前没有可以喂养的物品",
		SoulNotActiveNotStrengthen: "未激活的帝魂,无法强化",
		SoulStrengthenReackLimit:   "帝魂强化已达最高阶",
		SoulAwakenIsExist:          "帝魂已经觉醒过了",
		SoulActivateNotice:         "恭喜%s成功激活%s，帝魂可给敌人造成巨额伤害附加各种控制效果！",
		SoulAwakenNotice:           "恭喜%s成功觉醒%s，帝魂技能可对玩家生效！",
	}
)

func init() {
	mergeLang(soulLangMap)
}
