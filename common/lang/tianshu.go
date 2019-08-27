package lang

const (
	TianShuHadActivate LangCode = TianShuBase + iota
	TianShuNotActivate
	TianShuNotEnoughCondition
	TianShuHadReceive
	TianShuHadFullLevel
)

var (
	tianshuLangMap = map[LangCode]string{
		TianShuHadActivate:        "天书已激活",
		TianShuNotActivate:        "天书未激活",
		TianShuNotEnoughCondition: "不满足天书激活条件",
		TianShuHadReceive:         "天书奖励已领取",
		TianShuHadFullLevel:       "天书已满级",
	}
)

func init() {
	mergeLang(tianshuLangMap)
}
