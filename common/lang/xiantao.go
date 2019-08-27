package lang

const (
	XianTaoCollectCountNoEnough LangCode = XianTaoBase + iota
	XianTaoPeachNumNoEnough
	XianTaoCommitNoDistance
)

var (
	xiantoLangMap = map[LangCode]string{
		XianTaoCollectCountNoEnough: "采集次数不足",
		XianTaoPeachNumNoEnough:     "仙桃数量不足",
		XianTaoCommitNoDistance:     "不在提交范围",
	}
)

func init() {
	mergeLang(xiantoLangMap)
}
