package listener

import (
	"fgame/fgame/core/event"
	activityeventtypes "fgame/fgame/game/activity/event/types"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家参加活动
//TODO 后面的配置是配成一个活动类型 需和策划数据一起优化
func playerJoinActivity(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ, ok := data.(activitytypes.ActivityType)
	if !ok {
		return
	}

	err = allianceActivity(pl, typ)
	if err != nil {
		return
	}

	err = moonLoveActivity(pl, typ)
	if err != nil {
		return
	}

	err = fourgodActivity(pl, typ)
	if err != nil {
		return
	}
	return
}

//参加仙盟活动(屠龙，城战，仙盟镖都算)X次
func allianceActivity(pl player.Player, typ activitytypes.ActivityType) (err error) {
	switch typ {
	case activitytypes.ActivityTypeAlliance:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAlliance, 0, 1)
	}
	return
}

//参加月下情缘活动X次
func moonLoveActivity(pl player.Player, typ activitytypes.ActivityType) (err error) {
	switch typ {
	case activitytypes.ActivityTypeMoonLove:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeMoonLove, 0, 1)
	}
	return
}

//参加四神秘境X次
func fourgodActivity(pl player.Player, typ activitytypes.ActivityType) (err error) {
	switch typ {
	case activitytypes.ActivityTypeFourGod:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFourMysterious, 0, 1)
	}
	return
}

func init() {
	gameevent.AddEventListener(activityeventtypes.EventTypeActivityJoin, event.EventListenerFunc(playerJoinActivity))
}
