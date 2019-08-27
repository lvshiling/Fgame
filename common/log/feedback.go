package log

type FeedbackLogReason int32

const (
	FeedbackLogReasonGM FeedbackLogReason = iota + 1
	FeedbackLogReasonExchange
	FeedbackLogReasonExchangeGold
	FeedbackLogReasonExchangeRefund
	FeedbackLogReasonUseItemRew
)

func (zslr FeedbackLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	feedbackReasonMap = map[FeedbackLogReason]string{
		FeedbackLogReasonGM:             "gm修改",
		FeedbackLogReasonExchange:       "兑换现金",
		FeedbackLogReasonExchangeGold:   "兑换元宝[%d]",
		FeedbackLogReasonExchangeRefund: "兑换码过期,code[%s]",
		FeedbackLogReasonUseItemRew:     "使用现金元宝卡,数量：%d",
	}
)

func (ar FeedbackLogReason) String() string {
	return feedbackReasonMap[ar]
}
