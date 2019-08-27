package listener

import (
	"fgame/fgame/core/event"
	activitylogic "fgame/fgame/game/activity/logic"
	playeractivity "fgame/fgame/game/activity/player"
	activitytypes "fgame/fgame/game/activity/types"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//九霄城战城定时奖励
func allianceSceneTickRew(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancescene.AllianceSceneData)

	propertyTickRew(sd)
	warPointTickRew(sd)
	return
}

//资源定时奖励
func propertyTickRew(sd alliancescene.AllianceSceneData) {
	s := sd.GetScene()
	warTemplate := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	for _, p := range s.GetAllPlayers() {
		pl, ok := p.(player.Player)
		if !ok {
			continue
		}

		if !activitylogic.IsAddTickRew(pl, activitytypes.ActivityTypeAlliance, warTemplate.FristTiem, warTemplate.RewTiem) {
			continue
		}

		gold := int32(0)
		bindGold := int32(0)
		silver := int64(0)
		exp := int64(0)
		expPoint := int32(0)
		itemMap := map[int32]int32{}
		if s.MapTemplate().IsSafe(pl.GetPos()) {
			silver = warTemplate.RewSilverSafeArea
			exp = warTemplate.RewExpSafeArea
			expPoint = warTemplate.RewExpPointSafeArea
		} else {
			silver = warTemplate.RewSilver
			exp = warTemplate.RewExp
			expPoint = warTemplate.RewExpPoint
		}
		activitylogic.AddActivityTickRew(pl, activitytypes.ActivityTypeAlliance, gold, bindGold, int32(exp), expPoint, int32(silver), itemMap)
	}
}

// 城战积分定时奖励
func warPointTickRew(sd alliancescene.AllianceSceneData) {
	s := sd.GetScene()
	warTemplate := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	for _, p := range s.GetAllPlayers() {
		pl, ok := p.(player.Player)
		if !ok {
			continue
		}

		if s.MapTemplate().IsSafe(pl.GetPos()) {
			continue
		}

		if !activitylogic.IsAddTickPointRew(pl, activitytypes.ActivityTypeAlliance, warTemplate.FristJiFenTime, warTemplate.RewJiFenTime) {
			continue
		}

		allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
		curPoint := allianceManager.AddWarPoint(warTemplate.RewJiFen)

		activityManager := pl.GetPlayerDataManager(playertypes.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
		activityManager.UpdateLastRewPointTime(activitytypes.ActivityTypeAlliance)

		scMsg := pbutil.BuildSCAllianceSceneWarPointChanged(curPoint)
		pl.SendMsg(scMsg)
	}
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneTickRew, event.EventListenerFunc(allianceSceneTickRew))
}
