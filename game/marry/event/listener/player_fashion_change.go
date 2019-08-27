package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fashioneventtypes "fgame/fgame/game/fashion/event/types"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家时装变化变化
func fashionChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	if marryInfo.Status == marrytypes.MarryStatusTypeUnmarried ||
		marryInfo.Status == marrytypes.MarryStatusTypeDivorce {
		return
	}

	playerId := pl.GetId()
	fashionId := pl.GetFashionId()
	scMarryFashionChange := pbuitl.BuildSCMarryFashionChange(playerId, fashionId)

	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	pl.SendMsg(scMarryFashionChange)
	if spl == nil {
		return
	}
	spl.SendMsg(scMarryFashionChange)
	return
}

func init() {
	gameevent.AddEventListener(fashioneventtypes.EventTypeFashionChanged, event.EventListenerFunc(fashionChanged))
}
