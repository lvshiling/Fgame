package types

type FeedbackStatus int32

const (
	FeedbackStatusInit    FeedbackStatus = iota //初始化
	FeedbackStatusProcess                       //进行中
	FeedbackStatusFinish                        //兑换成功
	FeedbackStatusFailed                        //兑换失败

)

var (
	feedbackStatusMap = map[FeedbackStatus]string{
		FeedbackStatusInit:    "初始化",
		FeedbackStatusProcess: "进行中",
		FeedbackStatusFinish:  "兑换成功",
		FeedbackStatusFailed:  "兑换失败",
	}
)

func (s FeedbackStatus) String() string {
	return feedbackStatusMap[s]
}

type FeedbackExchangeStatus int32

const (
	FeedbackExchangeStatusInit         FeedbackExchangeStatus = iota //初始化
	FeedbackExchangeStatusGenerateCode                               //生成码
	FeedbackExchangeStatusProcess                                    //进行中
	FeedbackExchangeStatusFinish                                     //兑换成功
	FeedbackExchangeStatusFailed                                     //兑换失败
	FeedbackExchangeStatusNotify                                     //通知
)

var (
	feedbackExchangeStatusMap = map[FeedbackExchangeStatus]string{
		FeedbackExchangeStatusInit:         "初始化",
		FeedbackExchangeStatusGenerateCode: "生成码",
		FeedbackExchangeStatusProcess:      "进行中",
		FeedbackExchangeStatusFinish:       "兑换成功",
		FeedbackExchangeStatusFailed:       "兑换失败",
		FeedbackExchangeStatusNotify:       "成功失败通知",
	}
)

func (s FeedbackExchangeStatus) String() string {
	return feedbackExchangeStatusMap[s]
}

type FeedbackFeeType int32

const (
	FeedbackFeeTypeCash FeedbackFeeType = iota
	FeedbackFeeTypeGold
)

const MoneyYuan = int32(100)
