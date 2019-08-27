package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/supremetitle/pbutil"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"
)

//玩家登录成功后下发至尊称号
func playerTitleAfterLogin(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	titleWear := manager.GetTitleId()
	titleMap := manager.GetTitleMap()

	scTitleGet := pbutil.BuildSCSupremeTitleGet(titleWear, titleMap)
	p.SendMsg(scTitleGet)

	curTitleId := manager.GetTitleId()
	supremeTitleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(curTitleId)
	if supremeTitleTemplate != nil {
		//增加buff
		buffId := supremeTitleTemplate.BuffId
		if buffId != 0 {
			scenelogic.AddBuffs(p, buffId, p.GetId(), 1, common.MAX_RATE)
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerTitleAfterLogin))
}
