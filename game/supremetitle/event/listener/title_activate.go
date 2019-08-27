package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	supremetitleeventtypes "fgame/fgame/game/supremetitle/event/types"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"
	"fmt"
)

//玩家至尊称号激活
func playerSupremeTitleActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	titleId := data.(int32)
	titleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(titleId)
	if titleTemplate == nil {
		return
	}

	power := propertylogic.CulculateForce(titleTemplate.GetBattleProperty())
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	titleName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(titleTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, coreutils.FormatNoticeStr(fmt.Sprintf("%d", power)))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.TitleActivateNotice), playerName, titleName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(supremetitleeventtypes.EventTypeSupremeTitleActivate, event.EventListenerFunc(playerSupremeTitleActivate))
}
