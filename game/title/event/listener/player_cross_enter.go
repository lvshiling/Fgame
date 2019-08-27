package listener

import (
	"fgame/fgame/core/event"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	crosstypes "fgame/fgame/game/cross/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shenmotemplate "fgame/fgame/game/shenmo/template"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	playertitle "fgame/fgame/game/title/player"
)

//玩家跨服进入
func crossEnter(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	if pl.GetCrossType() != crosstypes.CrossTypeShenMoWar {
		return
	}

	killNum := pl.GetShenMoKillNum()
	//卸下称号
	manager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	manager.TitleNoWear()
	supremeTitleManager := pl.GetPlayerDataManager(types.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	supremeTitleManager.TitleNoWear()
	shenMoTitleTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoTitleTemplateByKillNum(killNum)
	if shenMoTitleTemplate == nil {
		return
	}
	if shenMoTitleTemplate.Title == 0 {
		return
	}
	manager.TempTitleAdd(shenMoTitleTemplate.Title)
	return nil
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossEnter, event.EventListenerFunc(crossEnter))
}
