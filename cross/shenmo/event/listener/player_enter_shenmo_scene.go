package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/shenmo/pbutil"
	crossshenmo "fgame/fgame/cross/shenmo/shenmo"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	shenmoscene "fgame/fgame/game/shenmo/scene"
	gameshenmo "fgame/fgame/game/shenmo/shenmo"
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
	gameshenmo.GetShenMoService().SyncSceneNum(num)

	gongXunNum := pl.GetShenMoGongXunNum()
	killNum := pl.GetShenMoKillNum()
	jiFenNum := int32(0)
	allianceId := pl.GetAllianceId()
	if allianceId != 0 {
		jiFenNum = crossshenmo.GetShenMoService().GetJiFenNum(allianceId)
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
