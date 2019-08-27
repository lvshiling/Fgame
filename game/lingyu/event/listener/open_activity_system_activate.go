package listener

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/event"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//运营活动-系统激活
func playerSystemActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	activitySubType, ok := data.(welfaretypes.OpenActivitySystemActivateSubType)
	if !ok {
		return
	}

	if activitySubType != welfaretypes.OpenActivitySystemActivateSubTypeLingYu {
		return
	}

	lingyuManager := pl.GetPlayerDataManager(playertypes.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()
	lingyuManager.LingyuAdvanced(0, 0, true)
	//同步属性
	lingyulogic.LingyuPropertyChanged(pl)
	scTianMoAdvanced := pbutil.BuildSCLingyuAdavancedFinshed(int32(lingyuInfo.AdvanceId), lingyuInfo.LingyuId, commontypes.AdvancedTypeOpenActivity)
	pl.SendMsg(scTianMoAdvanced)

	lingyuReason := commonlog.LingyuLogReasonAdvanced
	reasonText := fmt.Sprintf(lingyuReason.String(), commontypes.AdvancedTypeOpenActivity.String())
	logData := lingyueventtypes.CreatePlayerLingyuAdvancedLogEventData(0, 1, lingyuReason, reasonText)
	gameevent.Emit(lingyueventtypes.EventTypeLingyuAdvancedLog, pl, logData)

	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeSystemActivate, event.EventListenerFunc(playerSystemActivate))
}
