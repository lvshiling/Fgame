package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tianshutypes "fgame/fgame/game/tianshu/types"
	"math"
)

//玩家充值成功
func playerChargeSuccess(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*chargeeventtypes.PlayerChargeSuccessEventData)
	if !ok {
		return
	}

	chargeGold := eventData.GetChargeGold()

	// 特权充值
	goldBackRate := pl.GetTianShuRate(tianshutypes.TianShuTypeGold)
	if goldBackRate > 0 {
		extraGold := int32(math.Ceil(float64(chargeGold) * float64(goldBackRate) / float64(common.MAX_RATE)))
		title := lang.GetLangService().ReadLang(lang.EmailTianShuGoldFeedbackTitle)
		content := lang.GetLangService().ReadLang(lang.EmailTianShuGoldFeedbackContent)
		attachment := make(map[int32]int32)
		attachment[int32(constanttypes.GoldItem)] = extraGold
		emaillogic.AddEmail(pl, title, content, attachment)
	}
	bindGoldBackRate := pl.GetTianShuRate(tianshutypes.TianShuTypeBindGold)
	if bindGoldBackRate > 0 {
		extraBindGold := int32(math.Ceil(float64(chargeGold) * float64(bindGoldBackRate) / float64(common.MAX_RATE)))
		title := lang.GetLangService().ReadLang(lang.EmailTianShuBindGoldFeedbackTitle)
		content := lang.GetLangService().ReadLang(lang.EmailTianShuBindGoldFeedbackContent)
		attachment := make(map[int32]int32)
		attachment[int32(constanttypes.BindGoldItem)] = extraBindGold
		emaillogic.AddEmail(pl, title, content, attachment)
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeSuccess, event.EventListenerFunc(playerChargeSuccess))
}
