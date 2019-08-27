package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/marry/marry"
	marryscene "fgame/fgame/game/marry/scene"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/scene/scene"
)

//玩家名字变化
func playerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	marry.GetMarryService().PlayerNameChanged(pl)

	//刷结婚场景玩家名字
	marryScene := marry.GetMarryService().GetScene()
	marryctx := scene.WithScene(context.Background(), marryScene)
	marryScene.Post(message.NewScheduleMessage(onRefreshMarryScenePlayerName, marryctx, pl, nil))

	return
}

func onRefreshMarryScenePlayerName(ctx context.Context, result interface{}, err error) error {
	marryScene := marry.GetMarryService().GetScene()
	pl, ok := result.(player.Player)
	if !ok {
		return nil
	}
	//更新场景数据
	marryDelegate := marryScene.SceneDelegate()
	marryDelegate.(marryscene.MarrySceneData).OnWeddingPlayerNameChanged(pl)
	return nil
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerNameChanged, event.EventListenerFunc(playerNameChanged))
}
