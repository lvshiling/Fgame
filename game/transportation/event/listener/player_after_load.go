package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/transportation/pbutil"
	playertransportation "fgame/fgame/game/transportation/player"
	"fgame/fgame/game/transportation/transpotation"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	transManager := pl.GetPlayerDataManager(types.PlayerTransportationType).(*playertransportation.PlayerTransportationDataManager)
	personalTimes := transManager.GetTranspotTimes()
	allianceTimes := alliance.GetAllianceService().GetAllianceTransportTimes(pl.GetId())
	scPlayerTransportationBriefInfo := pbutil.BuildSCPlayerTransportationBriefInfo(personalTimes, allianceTimes)
	pl.SendMsg(scPlayerTransportationBriefInfo)

	//玩家上线时重新同步下数据
	biaoChe := transpotation.GetTransportService().GetTransportation(pl.GetId())
	if biaoChe == nil {
		return
	}
	obj := biaoChe.GetTransportationObject()
	scTransportBriefInfoNotice := pbutil.BuildSCTransportBriefInfoNotice(obj, biaoChe)
	pl.SendMsg(scTransportBriefInfoNotice)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
