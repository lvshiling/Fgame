package lang

const (
	ShenQiSlotLevelMax LangCode = ShenQiBase + iota
	ShenQiResolveNotQiLing
	ShenQiBagItemNotQiLing
	ShenQiLevelNotEnough
	ShenQiQiLingQualityNoEnough
	ShenQiQiLingSlotNotMatch
	ShenQiQiLingSlotReplaceFail
	ShenQiNotInlayQiLing
	ShenQiLingQiNotEnough
)

var (
	shenQiLangMap = map[LangCode]string{
		ShenQiSlotLevelMax:          "该槽位已经满级了",
		ShenQiResolveNotQiLing:      "请选择器灵分解",
		ShenQiBagItemNotQiLing:      "物品不是器灵",
		ShenQiLevelNotEnough:        "当前神器等级不足",
		ShenQiQiLingQualityNoEnough: "器灵品质不足",
		ShenQiQiLingSlotReplaceFail: "器灵替换失败",
		ShenQiNotInlayQiLing:        "该槽位没有镶嵌器灵",
		ShenQiLingQiNotEnough:       "灵气值不足",
	}
)

func init() {
	mergeLang(shenQiLangMap)
}
