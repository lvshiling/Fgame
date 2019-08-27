package lang

const (
	MaterialNotEnoughChallengeTimes LangCode = MaterialBase + iota
	MaterialGroupNotEnough
)

var (
	materialLangMap = map[LangCode]string{
		MaterialNotEnoughChallengeTimes: "副本次数不足",
		MaterialGroupNotEnough:          "挑战波数不足，无法扫荡",
	}
)

func init() {
	mergeLang(materialLangMap)
}
