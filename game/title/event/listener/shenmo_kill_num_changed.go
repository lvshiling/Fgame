package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	crosstypes "fgame/fgame/game/cross/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shenmotemplate "fgame/fgame/game/shenmo/template"
	playertitle "fgame/fgame/game/title/player"
)

//玩家击杀人数变更
func playerShenMoGongKillNumChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	killNum := pl.GetShenMoKillNum()

	manager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	shenMoTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoTitleTemplateByKillNum(killNum)
	if shenMoTemplate == nil {
		return
	}
	if pl.GetCrossType() != crosstypes.CrossTypeShenMoWar {
		return
	}
	titleId := manager.GetTitleId()
	if titleId == shenMoTemplate.Title {
		return
	}
	manager.TempTitleRemove(titleId)
	manager.TempTitleAdd(shenMoTemplate.Title)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShenMoKillNumChanged, event.EventListenerFunc(playerShenMoGongKillNumChanged))
}
