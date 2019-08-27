package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//玩家名字变化
func marryPlayerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	marryObj, ok := target.(*marry.MarryObject)
	if !ok {
		return
	}
	curPlayerId, ok := data.(int64)
	if !ok {
		return
	}
	peerId := int64(0)
	name := ""
	if curPlayerId == marryObj.PlayerId {
		peerId = marryObj.SpouseId
		name = marryObj.PlayerName
	}
	if curPlayerId == marryObj.SpouseId {
		peerId = marryObj.PlayerId
		name = marryObj.SpouseName
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(peerId)
	if pl != nil {
		plCtx := scene.WithPlayer(context.Background(), pl)
		playerSpouseNameMsg := message.NewScheduleMessage(onPlayerSpouseNameChanged, plCtx, name, nil)
		pl.Post(playerSpouseNameMsg)
	}
	return
}

func onPlayerSpouseNameChanged(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	name := result.(string)
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	spouseId := marryInfo.SpouseId
	if spouseId == 0 {
		return nil
	}
	manager.SpouseNameChanged(name)
	scMarryNameChange := pbuitl.BuildSCMarryNameChange(spouseId, name)
	pl.SendMsg(scMarryNameChange)
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryPlayerNameChanged, event.EventListenerFunc(marryPlayerNameChanged))
}
