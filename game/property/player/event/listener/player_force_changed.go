package listener

// //玩家玩家进入场景
// func playerForceChanged(target event.EventTarget, data event.EventData) (err error) {
// 	p, ok := target.(player.Player)
// 	if !ok {
// 		return
// 	}
// 	m := p.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
// 	m.SetPower(p.GetForce())
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerForceChanged, event.EventListenerFunc(playerForceChanged))
// }
