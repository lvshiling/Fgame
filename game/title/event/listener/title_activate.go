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
	titleeventtypes "fgame/fgame/game/title/event/types"
	"fgame/fgame/game/title/title"
	"fmt"
)

//玩家称号激活
func playerTitleActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	titleId := data.(int32)
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return
	}
	attrTemp := titleTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}
	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	titleName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(titleTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, coreutils.FormatNoticeStr(fmt.Sprintf("%d", power)))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.TitleActivateNotice), playerName, titleName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(titleeventtypes.EventTypeTitleActivate, event.EventListenerFunc(playerTitleActivate))
}
