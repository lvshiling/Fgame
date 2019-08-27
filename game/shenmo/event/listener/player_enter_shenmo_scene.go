package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	"fgame/fgame/game/shenmo/pbutil"
	shenmoscene "fgame/fgame/game/shenmo/scene"
	shenmo "fgame/fgame/game/shenmo/shenmo"
	shenmotemplate "fgame/fgame/game/shenmo/template"
)

//玩家进入神魔战场
func shenMoPlayerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(shenmoscene.ShenMoSceneData)
	if !ok {
		return
	}
	pl, ok := data.(scene.Player)
	if !ok {
		return
	}

	num := sd.GetScenePlayerNum()
	shenmo.GetShenMoService().SyncSceneNum(num)

	gongXunNum := pl.GetShenMoGongXunNum()
	killNum := pl.GetShenMoKillNum()
	jiFenNum := int32(0)
	allianceId := pl.GetAllianceId()
	if allianceId != 0 {
		jiFenNum = shenmo.GetShenMoService().GetJiFenNum(allianceId)
	}

	shenMoTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoConstantTemplate()
	s := sd.GetScene()
	daQiList := s.GetNPCListByBiology(shenMoTemplate.DaQiBiologyId)

	scLianYuGet := pbutil.BuildSCShenMoSceneInfo(gongXunNum, killNum, jiFenNum, daQiList)
	pl.SendMsg(scLianYuGet)
	return
}

func init() {
	gameevent.AddEventListener(shenmoeventtypes.EventTypeShenMoPlayerEnter, event.EventListenerFunc(shenMoPlayerEnterScene))
}
