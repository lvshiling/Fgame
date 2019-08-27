package lang

const (
	FireworksNotice LangCode = FireworksBase + iota
)

var (
	fireworksLangMap = map[LangCode]string{
		FireworksNotice: "%s燃放了烟花%s,为盛大婚礼增加了更多喜庆氛围",
	}
)

func init() {
	mergeLang(fireworksLangMap)
}
