package lang

const (
	FuShiTemplateNotExist LangCode = FuShiBase + iota
	FuShiBaGuaMiJingLevelNotEnough
	FuShiAlreadyActivite
	FuShiNotActivite
	FuShiAlreadyTopLevel
)

var (
	fushiLangMap = map[LangCode]string{
		FuShiTemplateNotExist:          "符石模板不存在",
		FuShiBaGuaMiJingLevelNotEnough: "八卦秘境通关等级不足",
		FuShiAlreadyActivite:           "符石已经激活",
		FuShiNotActivite:               "符石未激活",
		FuShiAlreadyTopLevel:           "符石已经达到最高等级",
	}
)

func init() {
	mergeLang(fushiLangMap)
}
