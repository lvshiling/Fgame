package lang

const (
	PkStateInvalid LangCode = PkBase + iota
	PkStateSame
	PkStateForbid
	PkStateProtect
	PkStateProtectCanNotAttack
)

var pkLangMap = map[LangCode]string{
	PkStateInvalid:             "pk状态无效",
	PkStateSame:                "pk状态一样",
	PkStateForbid:              "地图禁止此pk模式",
	PkStateProtect:             "该玩家等级不足%s级，无法切换pk模式",
	PkStateProtectCanNotAttack: "该玩家等级不足%s级，无法被攻击",
}

func init() {
	mergeLang(pkLangMap)
}
