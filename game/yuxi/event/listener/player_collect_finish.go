package listener

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	yuxiscene "fgame/fgame/game/yuxi/scene"
	yuxitemplate "fgame/fgame/game/yuxi/template"
)

//采集 采集完成
func playerCollectFinishWith(target event.EventTarget, data event.EventData) (err error) {
	spl, ok := target.(scene.Player)
	if !ok {
		return
	}
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeYuXi {
		return
	}

	sd := s.SceneDelegate()
	yuxiSd, ok := sd.(yuxiscene.YuXiSceneData)
	if !ok {
		return
	}

	//加buff
	buffId := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate().BuffId
	scenelogic.AddBuff(pl, buffId, pl.GetId(), common.MAX_RATE)
	//buff持续计时
	yuxiSd.CollectYuXiFinish(spl)
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinishWith, event.EventListenerFunc(playerCollectFinishWith))
}
