package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/supremetitle/pbutil"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"
	titleeventtypes "fgame/fgame/game/title/event/types"
	playertitle "fgame/fgame/game/title/player"
)

//称号改变
func titleChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	titleId := m.GetTitleId()
	//卸下至尊称号
	supremeTitleManager := pl.GetPlayerDataManager(playertypes.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	curTitleId := supremeTitleManager.GetTitleId()
	if curTitleId != 0 && titleId != 0 {
		supremeTitleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(curTitleId)
		if supremeTitleTemplate != nil {
			scenelogic.RemoveBuff(pl, supremeTitleTemplate.BuffId)
		}
		supremeTitleManager.TitleNoWear()
		scSupremeTitleUnload := pbutil.BuildSCSupremeTitleUnload(curTitleId)
		pl.SendMsg(scSupremeTitleUnload)
	}
	pl.SetTitleId(titleId)
	return
}

func init() {
	gameevent.AddEventListener(titleeventtypes.EventTypeTitleChanged, event.EventListenerFunc(titleChanged))
}
