package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	marryeventtypes "fgame/fgame/game/marry/event/types"
)

//玩家表白等级变化
func marryDevelopLevelChanged(target event.EventTarget, data event.EventData) (err error) {
	marryObj, ok := target.(*marry.MarryObject)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryDevelopLevelChangedEventData)
	if !ok {
		return
	}

	splId := int64(0)
	if eventData.GetChangedPlayerId() == marryObj.PlayerId {
		splId = marryObj.SpouseId
	} else {
		splId = marryObj.PlayerId
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(splId)
	if spl == nil {
		return
	}

	splCtx := scene.WithPlayer(context.Background(), spl)
	msg := message.NewScheduleMessage(onCoupleMarryDevelopChanged, splCtx, eventData.GetDevelopLevel(), nil)
	spl.Post(msg)
	return
}

func onCoupleMarryDevelopChanged(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	developLevel := result.(int32)

	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.UpdateCoupleMarryDevelopLevel(developLevel)
	marrylogic.MarryPropertyChanged(pl)

	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryDevelopLevelChanged, event.EventListenerFunc(marryDevelopLevelChanged))
}
