package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//玩家预定婚期
func playerMarryPreWed(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*marryeventtypes.MarryPreWedEventData)
	if !ok {
		return
	}
	spouseId := eventData.GetSpouseId()
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl == nil {
		return
	}

	ctx := scene.WithPlayer(context.Background(), spl)
	playerPreWedMsg := message.NewScheduleMessage(onMarryPreWed, ctx, eventData, nil)
	spl.Post(playerPreWedMsg)
	return
}

func onMarryPreWed(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*marryeventtypes.MarryPreWedEventData)
	period := data.GetPeriod()
	playerName := data.GetPlayerName()
	marryGrade := data.GetMarryGrade()

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.MarryPreWedFlag(true)

	scMarryWedGradeToSpouse := pbuitl.BuildSCMarryWedGradeToSpouse(period, playerName, marryGrade)
	pl.SendMsg(scMarryWedGradeToSpouse)
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryPreWed, event.EventListenerFunc(playerMarryPreWed))
}
