package listener

import (
	"fgame/fgame/core/event"
	shenmologic "fgame/fgame/cross/shenmo/logic"
	"fgame/fgame/cross/shenmo/pbutil"
	"fgame/fgame/cross/shenmo/shenmo"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	shenmotemplate "fgame/fgame/game/shenmo/template"
)

//采集 采集完成
func playerCollectFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	n, ok := data.(scene.NPC)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossShenMo {
		return
	}

	biologyId := int32(n.GetBiologyTemplate().TemplateId())
	shenMoTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoConstantTemplate()
	if !shenMoTemplate.IsShenMoCollect(biologyId) {
		return
	}

	addGongXunNum, addJiFenNum := shenMoTemplate.GetCollectPoint(biologyId)

	//增加仙盟积分
	allianceId := pl.GetAllianceId()
	if allianceId != 0 && addJiFenNum != 0 {
		serverId := pl.GetServerId()
		allianceName := pl.GetAllianceName()
		shenmo.GetShenMoService().AddJiFenNum(serverId, allianceId, allianceName, addJiFenNum)
		shenmologic.JiFenChangedAllianceBroadcast(s, allianceId)
	}

	//推送本服
	isPlayerGongXunAdd := pbutil.BuildISPlayerGongXunAdd(addGongXunNum)
	pl.SendMsg(isPlayerGongXunAdd)

	//推送大旗生物信息
	if shenMoTemplate.DaQiBiologyId == biologyId {
		bcMsg := pbutil.BuildSCShenMoBioBroadcast(n)
		s.BroadcastMsg(bcMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinish, event.EventListenerFunc(playerCollectFinish))
}
