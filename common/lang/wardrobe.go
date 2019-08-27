package lang

const (
	WardrobeNotActiveNotEat LangCode = WardrobeBase + iota
	WardrobeEatDanReachedLimit
	WardrobeEatDanNumReachedLimit
)

var (
	wardrobeLangMap = map[LangCode]string{
		WardrobeNotActiveNotEat:       "至少需要激活一条【%s】属性才可食用资质丹",
		WardrobeEatDanReachedLimit:    "当前套装食用资质丹数量已达上限,请选择其他套装食用",
		WardrobeEatDanNumReachedLimit: "当前套装食用资质丹数量超过上限,无法食用",
	}
)

func init() {
	mergeLang(wardrobeLangMap)
}
