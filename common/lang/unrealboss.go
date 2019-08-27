package lang

const (
	UnrealBossBuyNumReachLimit LangCode = UnrealBoss + iota
	UnrealBossPilaoNumNotEnough
)

var (
	unrealBossLangMap = map[LangCode]string{
		UnrealBossBuyNumReachLimit:  "今日购买次数已达上限",
		UnrealBossPilaoNumNotEnough: "您的疲劳值不足，击杀BOSS将无法获得任何奖励！",
	}
)

func init() {
	mergeLang(unrealBossLangMap)
}
