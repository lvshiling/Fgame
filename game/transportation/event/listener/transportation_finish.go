package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/core/utils"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	transportationeventtypes "fgame/fgame/game/transportation/event/types"
	transportationlogic "fgame/fgame/game/transportation/logic"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	"fgame/fgame/game/transportation/pbutil"
	transportationtemplate "fgame/fgame/game/transportation/template"
	"fgame/fgame/game/transportation/transpotation"
	transportationtypes "fgame/fgame/game/transportation/types"
)

//镖车完成后
func transportationFinish(target event.EventTarget, data event.EventData) (err error) {
	biaoChe := target.(*biaochenpc.BiaocheNPC)
	if biaoChe == nil {
		return
	}

	transportationObj := biaoChe.GetTransportationObject()
	transportFinish(biaoChe)

	pl := player.GetOnlinePlayerManager().GetPlayerById(transportationObj.GetPlayerId())
	if pl != nil {
		scTransportBriefInfoNotice := pbutil.BuildSCTransportBriefInfoNotice(transportationObj, biaoChe)
		pl.SendMsg(scTransportBriefInfoNotice)
	}

	return
}

func transportFinish(biaoChe *biaochenpc.BiaocheNPC) {
	transportationObj := biaoChe.GetTransportationObject()
	typ := transportationObj.GetTransportType()

	//更新镖车状态
	transpotation.GetTransportService().TransportFinish(transportationObj.GetPlayerId())

	if typ == transportationtypes.TransportationTypeAlliance {
		//仙盟所有成员都将获得镖车奖励
		title := lang.GetLangService().ReadLang(lang.EmailAllianceTransportationFinishName)
		content := lang.GetLangService().ReadLang(lang.EmailAllianceTransportationFinishContent)
		rewMap := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(typ).GetFinishItemMap()
		// 倍数奖励
		activityTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeBiaoChe)
		now := global.GetGame().GetTimeService().Now()
		openTime := global.GetGame().GetServerTime()
		mergeTime := merge.GetMergeService().GetMergeTime()
		timeTemp, _ := activityTemp.GetActivityTimeTemplate(now, openTime, mergeTime)
		if timeTemp != nil {
			ratio := timeTemp.BeiShu
			rewMap = utils.MultMap(rewMap, ratio)
		}

		transportationlogic.AllianceTransportRewBroadcast(transportationObj.GetAllianceId(), title, content, rewMap)

		//移除镖车
		transpotation.GetTransportService().RemoveTransportation(transportationObj.GetPlayerId())
	}
}

func init() {
	gameevent.AddEventListener(transportationeventtypes.EventTypeTransportationFinish, event.EventListenerFunc(transportationFinish))
}
