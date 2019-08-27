package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	//推送所有物品
	equipBags := additionsysManager.GetAdditionSysEquipBags()
	levelMap := additionsysManager.GetAdditionSysLevelInfoAll()
	awakeMap := additionsysManager.GetAdditionSysAwakeInfoAll()
	tongLingMap := additionsysManager.GetAdditionSysTongLingInfoAll()
	equipSlotInfoList := pbutil.BuildSCAdditionSysSlotInfoList(equipBags, levelMap, tongLingMap, awakeMap)
	pl.SendMsg(equipSlotInfoList)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
