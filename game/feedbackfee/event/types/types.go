package types

import (
	commonlog "fgame/fgame/common/log"
)

type FeedbackfeeEvenType string

const (
	EventTypeCodeGenerate           FeedbackfeeEvenType = "CodeGenerate" //码生成
	EventTypeCodeExpire                                 = "CodeExpire"   //码过期了
	EventTypeCodeExchange                               = "CodeExchange" //码兑换了
	EventTypeFeedbackfeeExchangeLog                     = "FeedbackfeeExchangeLog"
)

type FeedbackfeeExchangeLogEventData struct {
	curMoney    int32
	beforeMoney int32
	changed     int32
	reason      commonlog.FeedbackLogReason
	reasonText  string
}

func CreateFeedbackfeeExchangeLogEventData(curMoney int32, beforeMoney int32, changed int32, reason commonlog.FeedbackLogReason, reasonText string) *FeedbackfeeExchangeLogEventData {
	data := &FeedbackfeeExchangeLogEventData{}
	data.curMoney = curMoney
	data.beforeMoney = beforeMoney
	data.changed = changed
	data.reason = reason
	data.reasonText = reasonText
	return data
}

func (d *FeedbackfeeExchangeLogEventData) GetCurMoney() int32 {
	return d.curMoney
}

func (d *FeedbackfeeExchangeLogEventData) GetBeforeMoney() int32 {
	return d.beforeMoney
}

func (d *FeedbackfeeExchangeLogEventData) GetChanged() int32 {
	return d.changed
}

func (d *FeedbackfeeExchangeLogEventData) GetReason() commonlog.FeedbackLogReason {
	return d.reason
}

func (d *FeedbackfeeExchangeLogEventData) GetReasonText() string {
	return d.reasonText
}
