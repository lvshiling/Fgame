package listener

// import (
// 	"fgame/fgame/core/event"
// 	battleeventtypes "fgame/fgame/game/battle/event/types"
// 	gameevent "fgame/fgame/game/event"
// 	"fgame/fgame/game/player"
// 	playertypes "fgame/fgame/game/player/types"
// 	propertylogic "fgame/fgame/game/property/logic"
// 	playerproperty "fgame/fgame/game/property/player"
// 	playerpropertytypes "fgame/fgame/game/property/player/types"
// )

// //pk值改变
// func playerPkValueChanged(target event.EventTarget, data event.EventData) (err error) {
// 	pl := target.(player.Player)

// 	propertyDataManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
// 	propertyDataManager.SetEvil(pl.GetPkValue())
// 	propertyDataManager.SetEvilOnlineTime(pl.GetPkOnlineTime())

// 	//TODO 修改红名值属性变化
// 	propertyDataManager.UpdateBattleProperty(playerpropertytypes.PropertyEffectorTypeMaskAll)
// 	propertylogic.SnapChangedProperty(pl)
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerPkValueChanged, event.EventListenerFunc(playerPkValueChanged))
// }
