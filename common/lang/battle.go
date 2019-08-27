package lang

const (
	BattleProtectChangeScene LangCode = BattleBase + iota
	BattleProtectRelive
	BattleProtectPK
	BattleProtectNoAttacked
)

var battleLangMap = map[LangCode]string{
	BattleProtectChangeScene: "该玩家处于过场保护中",
	BattleProtectRelive:      "该玩家处于复活保护中",
	BattleProtectPK:          "该玩家处于PK保护中",
	BattleProtectNoAttacked:  "该目标处于无敌状态中，无法被攻击",
}

func init() {
	mergeLang(battleLangMap)
}
