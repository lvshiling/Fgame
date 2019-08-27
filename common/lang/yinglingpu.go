package lang

const (
	YingLingPuSpExists LangCode = YingLingPuBase + iota
	YingLingPuSpNotAllow

	YingLingPuNotExists
	YingLingPuNotExistsTemplate
	YingLingPuMaxLevel

	YingLingPuSpOther
)

var (
	yingLingPuLangMap = map[LangCode]string{
		YingLingPuSpExists:   "碎片已经拥有",
		YingLingPuSpNotAllow: "无该碎片",

		YingLingPuNotExists:         "尚未拥有英灵谱",
		YingLingPuNotExistsTemplate: "英灵谱不存在",
		YingLingPuMaxLevel:          "英灵谱等级已最高",
		YingLingPuSpOther:           "镶嵌异常",
	}
)

func init() {
	mergeLang(yingLingPuLangMap)
}
