package lang

const (
	GuideFinishNotSlotsMailContent = GuideBase + iota
	GuideNotInRescureMap
	GuideCommitNoDistance
	GuideNotHerbs
)

var (
	guideLangMap = map[LangCode]string{
		GuideFinishNotSlotsMailContent: "引导副本奖励背包空间不足",
		GuideNotInRescureMap:           "不是救援副本",
		GuideCommitNoDistance:          "不在提交范围内",
		GuideNotHerbs:                  "身上没有草药",
	}
)

func init() {
	mergeLang(guideLangMap)
}
