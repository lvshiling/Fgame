package listener

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guidereplica/pbutil"
	guidereplicascene "fgame/fgame/game/guidereplica/scene"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
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

	//救援采集
	rescureCollect(pl)

	return
}

func rescureCollect(pl player.Player) {
	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeGuideReplicaRescue {
		return
	}

	sd := s.SceneDelegate()
	rescureSd, ok := sd.(guidereplicascene.GuideRescureSceneData)
	if !ok {
		return
	}

	//加buff
	guideTemp := rescureSd.GetGuideTemp()
	buffTemp := guideTemp.GetRescureGuideTemp().GetHerbsBuffTemplate()
	if pl.GetBuff(buffTemp.Group) == nil {
		scenelogic.AddBuff(pl, int32(buffTemp.TemplateId()), pl.GetId(), common.MAX_RATE)
		scMsg := pbutil.BuildSCGuideReplicaSceneDataChangedNoticeWithRescure(int32(rescureSd.GetGuideTemp().GetGuideType()), true)
		pl.SendMsg(scMsg)
	}
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinishWith, event.EventListenerFunc(playerCollectFinishWith))
}
