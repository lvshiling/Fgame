package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/player"
	advancedrewrewlogic "fgame/fgame/game/welfare/advancedrew/rew_max/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家灵童养成类进阶
func playerLingTongDevAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTongObj, ok := data.(*playerlingtongdev.PlayerLingTongDevObject)
	if !ok {
		return
	}
	classType := lingTongObj.GetClassType()
	advanceId := lingTongObj.GetAdvancedId()
	advancedType, ok := welfaretypes.LingTongDevTypeToAdvancedType(classType)
	if !ok {
		return
	}

	advancedrewrewlogic.UpdateAdvancedRewData(pl, advanceId, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevAdvanced, event.EventListenerFunc(playerLingTongDevAdavanced))
}
