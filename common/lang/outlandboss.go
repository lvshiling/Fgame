package lang

const (
	OutlandBossZhuoQiNumNotEnough LangCode = OutlandBossBase + iota
	OutlandBossZhuoQiNumIsZero
)

var (
	outlandBossLangMap = map[LangCode]string{
		OutlandBossZhuoQiNumNotEnough: "您当前浊气值已达上限，攻击BOSS将无法造成任何伤害！",
		OutlandBossZhuoQiNumIsZero:    "当前浊气值为0，无需使用",
	}
)

func init() {
	mergeLang(outlandBossLangMap)
}
