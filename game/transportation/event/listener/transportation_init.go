package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	transportationeventtypes "fgame/fgame/game/transportation/event/types"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	transportationtemplate "fgame/fgame/game/transportation/template"
	transportationtypes "fgame/fgame/game/transportation/types"

	log "github.com/Sirupsen/logrus"
)

//进入场景
func transportationAdd(target event.EventTarget, data event.EventData) (err error) {
	biaoChe := target.(*biaochenpc.BiaocheNPC)
	transportObj := biaoChe.GetTransportationObject()
	if transportObj.GetState() != transportationtypes.TransportStateTypeRuning {
		return
	}
	moveId := transportObj.GetTransportMoveId()
	moveTemplate := transportationtemplate.GetTransportationTemplateService().GetTransportationMoveTemplate(moveId)
	if moveTemplate == nil {
		log.WithFields(
			log.Fields{
				"moveId": moveId,
			}).Warn("transport:押镖路线模板不存在")

		return
	}
	s := scene.GetSceneService().GetWorldSceneByMapId(moveTemplate.MapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"MapId": moveTemplate.MapId,
			}).Warn("transport:押镖场景不存在")

		return
	}
	scenelogic.NPCEnterScene(biaoChe, s, moveTemplate.GetPosition())
	return
}

func init() {
	gameevent.AddEventListener(transportationeventtypes.EventTypeTransportationInit, event.EventListenerFunc(transportationAdd))
}
