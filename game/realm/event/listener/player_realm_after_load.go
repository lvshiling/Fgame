package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	realmlogic "fgame/fgame/game/realm/logic"
	"fgame/fgame/game/realm/pbutil"
	playerrealm "fgame/fgame/game/realm/player"
)

//加载完成后
func playerRealmAfterLogin(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	realmManager := p.GetPlayerDataManager(playertypes.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	level := realmManager.GetTianJieTaLevel()

	scRealmLevel := pbutil.BuildSCRealmLevel(level)
	p.SendMsg(scRealmLevel)

	//天劫塔检测补偿
	realmlogic.CheckReissue(p)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerRealmAfterLogin))
}
