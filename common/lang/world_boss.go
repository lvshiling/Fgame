package lang

const (
	WorldBossBeKilled LangCode = WorldBossBase + iota
	WorldBossBeKilledNotice
	WorldBossPlayerOnBattle
	WorldBossReborn
	WorldBossRebornNotice
	WorldBossHandlerNotExist
)

var (
	WorldBossLangMap = map[LangCode]string{
		WorldBossBeKilled:        "%s已被%s击杀，BOSS掉落了%s",
		WorldBossBeKilledNotice:  "%s已被%s击杀，BOSS掉落了%s",
		WorldBossPlayerOnBattle:  "当前处于战斗状态，无法传送",
		WorldBossReborn:          "%s已于%s复活，请广大英雄侠士前往击杀！%s",
		WorldBossRebornNotice:    "%s已于%s复活，请广大英雄侠士前往击杀！",
		WorldBossHandlerNotExist: "该类型的boss处理器没有注册",
	}
)

func init() {
	mergeLang(WorldBossLangMap)
}
