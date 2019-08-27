package lang

const (
	ChessNotEnougTimes LangCode = ChessBase + iota
	ChessNotGetRewards
)

var (
	chessLangMap = map[LangCode]string{
		ChessNotEnougTimes: "棋局抽奖次数不足",
		ChessNotGetRewards: "运气不好，再来一次",
	}
)

func init() {
	mergeLang(chessLangMap)
}
