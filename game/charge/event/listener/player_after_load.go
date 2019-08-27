package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/charge/charge"
	chargelogic "fgame/fgame/game/charge/logic"
	"fgame/fgame/game/charge/pbutil"
	playercharge "fgame/fgame/game/charge/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//档次首充记录推送
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)

	//重置首冲
	chargeManager := p.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	chargeTimeObj := charge.GetChargeService().GetChargeTime()
	if chargeTimeObj.ChargeTime != 0 {
		chargeManager.ResetFirstChargeRecord(chargeTimeObj.ChargeTime)
	}

	//恢复订单
	unfinishOrderList := charge.GetChargeService().GetUnfinishOrderList(p)
	for _, orderObj := range unfinishOrderList {
		chargelogic.OnPlayerCharge(p, orderObj.GetOrderId(), orderObj.GetChargeId())
	}

	recordMap := chargeManager.GetFirstChargeRecord()
	scMsg := pbutil.BuildSCFirstChargeRecordNotice(recordMap)
	p.SendMsg(scMsg)

	// 后台充值
	chargelogic.OnPlayerPrivilegeCharge(p)

	// 推送今日重置元宝数量
	todayGoldNum := chargeManager.GetTodayChargeNum()
	scGoldMsg := pbutil.BuildSCTodayChargeGold(todayGoldNum)
	p.SendMsg(scGoldMsg)

	// 推送新首充活动信息
	startTime, duration := charge.GetChargeService().GetNewFirstChargeTime()
	info := chargeManager.GetNewFirstChargeRecordInfo(startTime)
	reocrd := info.GetRecord()
	scNewFirstChargeRecord := pbutil.BuildSCNewFirstChargeRecord(startTime, duration, reocrd)
	p.SendMsg(scNewFirstChargeRecord)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
