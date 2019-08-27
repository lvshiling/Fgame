package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	marryscene "fgame/fgame/game/marry/scene"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//婚礼结束
func marryWedEnd(target event.EventTarget, data event.EventData) (err error) {
	sd := marry.GetMarryService().GetMarrySceneData()

	marryScene := marry.GetMarryService().GetScene()
	ctx := scene.WithScene(context.Background(), marryScene)
	marryScene.Post(message.NewScheduleMessage(onMarryBanquetFinish, ctx, nil, nil))

	//全服发送婚礼结束
	scMarryWedPushStatus := pbuitl.BuildSCMarryWedPushStatus(sd, true)
	player.GetOnlinePlayerManager().BroadcastMsg(scMarryWedPushStatus)
	return
}

func onMarryBanquetFinish(ctx context.Context, result interface{}, err error) error {
	marryScene := marry.GetMarryService().GetScene()
	marryDelegate := marryScene.SceneDelegate()
	//更新场景数据
	marryDelegate.(marryscene.MarrySceneData).OnBanquetEnd()
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryWedEnd, event.EventListenerFunc(marryWedEnd))
}
