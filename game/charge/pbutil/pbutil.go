package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playercharge "fgame/fgame/game/charge/player"
)

func BuildSCCharge(chargeId int32) *uipb.SCCharge {
	scCharge := &uipb.SCCharge{}
	scCharge.ChargeId = &chargeId
	return scCharge
}

func BuildSCChargeOrder(chargeId int32, chargeType int32, orderId string, notifyUrl string, sdkOrderId string, platformUserId string, money int32, playerId int64, playerName string, serverId int32, serverName string, extension string) *uipb.SCChargeOrder {
	scChargeOrder := &uipb.SCChargeOrder{}
	scChargeOrder.ChargeId = &chargeId
	scChargeOrder.ChargeType = &chargeType
	scChargeOrder.OrderId = &orderId
	scChargeOrder.SdkOrderId = &sdkOrderId
	scChargeOrder.NotifyUrl = &notifyUrl
	scChargeOrder.PlatformUserId = &platformUserId
	scChargeOrder.Money = &money
	scChargeOrder.PlayerId = &playerId
	scChargeOrder.PlayerName = &playerName
	scChargeOrder.ServerId = &serverId
	scChargeOrder.ServerName = &serverName
	scChargeOrder.Extension = &extension
	return scChargeOrder
}

func BuildSCFirstChargeRecordNotice(recordObj map[int32]*playercharge.PlayerFirstChargeRecordObject) *uipb.SCFirstChargeRecordNotice {
	scFirstChargeRecordNotice := &uipb.SCFirstChargeRecordNotice{}
	for chargeId, _ := range recordObj {
		scFirstChargeRecordNotice.ChargeIdList = append(scFirstChargeRecordNotice.ChargeIdList, chargeId)
	}
	return scFirstChargeRecordNotice
}

func BuildSCNewFirstChargeRecordNotice(record []int32) *uipb.SCNewFirstChargeRecordNotice {
	scNewFirstChargeRecordNotice := &uipb.SCNewFirstChargeRecordNotice{}
	scNewFirstChargeRecordNotice.ChargeIdList = record
	return scNewFirstChargeRecordNotice
}

func BuildSCChargeGoldNotice(gold int64) *uipb.SCChargeGoldNotice {
	scMsg := &uipb.SCChargeGoldNotice{}
	scMsg.ChargeGold = &gold
	return scMsg
}

func BuildSCTodayChargeGold(gold int64) *uipb.SCTodayChargeGold {
	scMsg := &uipb.SCTodayChargeGold{}
	scMsg.ChargeGold = &gold
	return scMsg
}

func BuildSCNewFirstChargeRecord(startTime int64, duration int64, record []int32) *uipb.SCNewFirstChargeRecord {
	scNewFirstChargeRecord := &uipb.SCNewFirstChargeRecord{}
	scNewFirstChargeRecord.StartTime = &startTime
	scNewFirstChargeRecord.Duration = &duration
	scNewFirstChargeRecord.ChargeIdList = record
	return scNewFirstChargeRecord
}
