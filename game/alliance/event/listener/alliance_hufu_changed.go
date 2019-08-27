package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//虎符改变
func allianceHuFuChanged(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	sd := alliance.GetAllianceService().GetAllianceSceneData()
	//判断是否在城战
	if sd != nil {
		sceneCtx := scene.WithScene(context.Background(), sd.GetScene())
		sd.GetScene().Post(message.NewScheduleMessage(onHuFuChanged, sceneCtx, al, nil))
	}

	return
}

//虎符改变回调
func onHuFuChanged(ctx context.Context, result interface{}, err error) error {
	al := result.(*alliance.Alliance)
	s := scene.SceneInContext(ctx)
	sceneData := s.SceneDelegate().(alliancescene.AllianceSceneData)
	sceneData.OnHuFuChanged(al.GetAllianceId(), al.GetAllianceObject().GetHuFu())
	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceHuFuChanged, event.EventListenerFunc(allianceHuFuChanged))
}
