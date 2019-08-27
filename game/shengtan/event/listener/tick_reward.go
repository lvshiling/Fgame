package listener

import (
	"fgame/fgame/core/event"
	activitylogic "fgame/fgame/game/activity/logic"
	activitytypes "fgame/fgame/game/activity/types"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shengtaneventtypes "fgame/fgame/game/shengtan/event/types"
	shengtantemplate "fgame/fgame/game/shengtan/template"
)

//TODO 修改为通用的 活动奖励模块
//圣坛定时奖励
func shengTanSceneTickRew(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancescene.AllianceSceneData)
	s := sd.GetScene()

	shengTanTemplate := shengtantemplate.GetShengTanTemplateService().GetShengTanTemplate()
	for _, p := range s.GetAllPlayers() {
		pl, ok := p.(player.Player)
		if !ok {
			continue
		}
		if !activitylogic.IsAddTickRew(pl, activitytypes.ActivityTypeAllianceShengTan, shengTanTemplate.FirstTime, shengTanTemplate.RewTime) {
			continue
		}

		shengTanAwradTemplate := shengtantemplate.GetShengTanTemplateService().GetShengTanAwardTemplate(pl.GetLevel())
		if shengTanAwradTemplate == nil {
			continue
		}
		silver := shengTanAwradTemplate.RewSilver
		gold := shengTanAwradTemplate.RewGold
		bindGold := shengTanAwradTemplate.RewBindGold
		exp := shengTanAwradTemplate.RewExp
		expPoint := shengTanAwradTemplate.RewExpPoint
		itemMap := shengTanAwradTemplate.GetItemMap()
		activitylogic.AddActivityTickRew(pl, activitytypes.ActivityTypeAllianceShengTan, gold, bindGold, int32(exp), expPoint, int32(silver), itemMap)
	}

	return
}

func init() {
	gameevent.AddEventListener(shengtaneventtypes.EventTypeShengTanSceneTickReward, event.EventListenerFunc(shengTanSceneTickRew))
}
