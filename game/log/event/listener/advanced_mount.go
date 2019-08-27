package listener

import (
	"fgame/fgame/core/event"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	playermount "fgame/fgame/game/mount/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家坐骑进阶日志
func playerMountAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*mounteventtypes.PlayerMountAdvancedLogEventData)
	if !ok {
		return
	}

	mountManager := pl.GetPlayerDataManager(playertypes.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()

	logMountAdvanced := &logmodel.PlayerMount{}
	logMountAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logMountAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logMountAdvanced.CurAdvancedNum = int32(mountInfo.AdvanceId)
	logMountAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logMountAdvanced.Reason = int32(eventData.GetReason())
	logMountAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logMountAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountAdvancedLog, event.EventListenerFunc(playerMountAdvancedLog))
}
