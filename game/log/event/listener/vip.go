package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	vipeventtypes "fgame/fgame/game/vip/event/types"
	playervip "fgame/fgame/game/vip/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家VIP升级日志
func playerVipUplevelLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*vipeventtypes.PlayerVipAdvancedLogEventData)
	if !ok {
		return
	}

	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	curLevel, _ := vipManager.GetVipLevel()
	curChargeGold := vipManager.GetChargeNum()
	beforeLevel := eventData.GetBeforeLevel()
	uplevel := curLevel - beforeLevel

	logVipAdvanced := &logmodel.PlayerVip{}
	logVipAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logVipAdvanced.BeforeVipLevel = eventData.GetBeforeLevel()
	logVipAdvanced.CurVipLevel = curLevel
	logVipAdvanced.Uplevel = uplevel
	logVipAdvanced.BeforeGold = eventData.GetBeforeGold()
	logVipAdvanced.CurGold = curChargeGold
	logVipAdvanced.AddGold = eventData.GetAddGold()
	logVipAdvanced.Reason = int32(eventData.GetReason())
	logVipAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logVipAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(vipeventtypes.EventTypeVipLevelChangedLog, event.EventListenerFunc(playerVipUplevelLog))
}
