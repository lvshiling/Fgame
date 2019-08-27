package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	playerxianfu "fgame/fgame/game/xianfu/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家仙府升级日志
func playerXianFuLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*xianfueventtypes.PlayerXianFuLogEventData)
	if !ok {
		return
	}

	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*playerxianfu.PlayerXinafuDataManager)
	xianfuInfo := xianfuManager.GetPlayerXianfuInfo(eventData.GetXianFuType())
	logXianFuAdvanced := &logmodel.PlayerXianFu{}
	logXianFuAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logXianFuAdvanced.BeforeLevel = eventData.GetBeforeLevel()
	logXianFuAdvanced.CurLevel = int32(xianfuInfo.GetXianfuId())
	logXianFuAdvanced.Uplevel = eventData.GetUplevel()
	logXianFuAdvanced.Reason = int32(eventData.GetReason())
	logXianFuAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logXianFuAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuLog, event.EventListenerFunc(playerXianFuLog))
}
