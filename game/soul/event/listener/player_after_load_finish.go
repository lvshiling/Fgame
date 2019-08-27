package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	"fgame/fgame/game/soul/pbutil"
	playersoul "fgame/fgame/game/soul/player"
)

//加载完成后
func playerSoulAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	tagList := manager.GetSoulEmbedTagList()
	for _, soulTag := range tagList {
		curOrder := manager.GetSoulOrderByTag(soulTag)
		newSkillId := manager.GetSkillId(soulTag, curOrder)
		skilllogic.TempSkillChangeNoUpdate(p, 0, newSkillId)
	}
	embedList := manager.GetSoulEmbed()
	soulInfo := manager.GetSoulInfoAll()
	scSoulGet := pbutil.BuildSCSoulGet(embedList, soulInfo)
	p.SendMsg(scSoulGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerSoulAfterLoadFinish))
}
