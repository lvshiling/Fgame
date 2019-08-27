package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	"fgame/fgame/game/arenapvp/pbutil"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

//竞猜推送
func arenapvpGuessBroadcast(target event.EventTarget, data event.EventData) (err error) {
	guessData, ok := target.(*arenapvpdata.GuessData)
	if !ok {
		return
	}

	//竞猜推送
	alPlayers := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range alPlayers {
		ctx := scene.WithPlayer(context.Background(), pl)
		arenapvpGuessMsg := message.NewScheduleMessage(onGuessNotice, ctx, guessData, nil)
		pl.Post(arenapvpGuessMsg)
	}

	return
}

//竞猜推送
func onGuessNotice(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	guessData := result.(*arenapvpdata.GuessData)

	fmt.Println(fmt.Sprint("--------------------------------------"))
	fmt.Println(fmt.Sprint("竞猜推送成功"))

	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpObj := arenapvpManager.GetPlayerArenapvpObj()
	if !arenapvpObj.IfGuessNotice() {
		return nil
	}

	scMsg := pbutil.BuildSCArenapvpGuessBeginNotice(guessData)
	pl.SendMsg(scMsg)
	return nil
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpGuessBroadcast, event.EventListenerFunc(arenapvpGuessBroadcast))
}
