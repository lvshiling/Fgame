package types

type ArenapvpEventType string

const (
	//复活次数变化
	EventTypeArenapvpReliveChanged = "ArenapvpReliveChanged"
	//竞猜结算
	EventTypeArenapvpGuessResult = "ArenapvpGuessResult"
	//竞猜推送
	EventTypeArenapvpGuessBroadcast = "ArenapvpGuessBroadcast"
	//竞猜退还
	EventTypeArenapvpGuessReturn = "ArenapvpGuessReturn"
	//pvp积分变化
	EventTypeArenapvpJiFenChanged = "ArenapvpJiFenChanged"
	//海选幸运奖推送
	EventTypeArenapvpElectionLuckyRew = "ArenapvpElectionLuckyRew"
	//晋级公告
	EventTypeArenapvpWinnerNotice = "ArenapvpWinnerNotice"
	//pvp结果
	EventTypeArenapvpResult = "ArenapvpResult"
)
