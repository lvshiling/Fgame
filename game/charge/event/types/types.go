package types

type ChargeEventType string

const (
	ChargeEventTypeChargeGold                  ChargeEventType = "ChargeGold" //充值元宝
	ChargeEventTypeGetOrderFailed                              = "GetOrderFailed"
	ChargeEventTypeGetOrderFinish                              = "GetOrderFinish"
	ChargeEventTypeOrderCharge                                 = "OrderCharge"
	ChargeEventTypePrivilegeCharge                             = "PrivilegeCharge"
	ChargeEventTypeFirstCycleCharge                            = "FirstCycleCharge"      //每日首充
	ChargeEventTypeChargeSuccess                               = "ChargeSuccess"         //充值成功
	ChargeEventTypeNewFirstChargeTimeChangeLog                 = "SetNewFirstChargeTime" // 新首充活动时间改变
)

//
type PlayerChargeSuccessEventData struct {
	chargeId   int32
	chargeGold int32
}

func CreatePlayerChargeSuccessEventData(chargeId, chargeGold int32) *PlayerChargeSuccessEventData {
	d := &PlayerChargeSuccessEventData{
		chargeId:   chargeId,
		chargeGold: chargeGold,
	}
	return d
}

func (d *PlayerChargeSuccessEventData) GetChargeId() int32 {
	return d.chargeId
}

func (d *PlayerChargeSuccessEventData) GetChargeGold() int32 {
	return d.chargeGold
}
