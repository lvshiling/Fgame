package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerfeedbackfee "fgame/fgame/game/feedbackfee/player"
)

func BuildSCFeedbackFeeInfo(xianJinFlag bool, obj *playerfeedbackfee.PlayerFeedbackFeeObject, recordObj *playerfeedbackfee.PlayerFeedbackRecordObject) *uipb.SCFeedbackFeeInfo {
	scMsg := &uipb.SCFeedbackFeeInfo{}
	scMsg.FeeInfo = buildFeedbackFeeInfo(obj)
	scMsg.XianJinFlag = &xianJinFlag
	exchanged := false
	if recordObj == nil {
		scMsg.Exchanged = &exchanged
		return scMsg
	}
	exchanged = true
	scMsg.Exchanged = &exchanged
	scMsg.RecordInfo = buildFeedbackRecordInfo(recordObj)
	return scMsg
}

func buildFeedbackFeeInfo(obj *playerfeedbackfee.PlayerFeedbackFeeObject) *uipb.FeedbackFeeInfo {

	money := obj.GetMoney()
	info := &uipb.FeedbackFeeInfo{}
	info.Money = &money
	todayUseNum := obj.GetTodayUseNum()
	info.TodayUseNum = &todayUseNum
	return info
}

func buildFeedbackRecordInfo(obj *playerfeedbackfee.PlayerFeedbackRecordObject) *uipb.FeedbackRecordInfo {
	recordInfo := &uipb.FeedbackRecordInfo{}
	status := int32(obj.GetStatus())
	recordInfo.Status = &status
	code := obj.GetCode()
	recordInfo.Code = &code

	money := obj.GetMoney()
	recordInfo.ExchangeMoney = &money
	createTime := obj.GetCreateTime()
	recordInfo.ExchangeTime = &createTime
	expireTime := obj.GetExpiredTime()
	recordInfo.ExpireTime = &expireTime
	return recordInfo
}
