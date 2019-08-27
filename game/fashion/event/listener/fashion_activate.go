package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	fashioneventtypes "fgame/fgame/game/fashion/event/types"
	"fgame/fgame/game/fashion/fashion"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家时装激活
func playerFashionActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fashionId := data.(int32)
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	if fashionTemplate == nil {
		return
	}
	attrTemp := fashionTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}
	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())
	fashionName := fashionTemplate.GetFashionName(pl.GetRole(), pl.GetSex())

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	fashionNameF := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(fashionName))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, coreutils.FormatNoticeStr(fmt.Sprintf("%d", power)))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FashionActivateNotice), playerName, fashionNameF, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(fashioneventtypes.EventTypeFashionActivate, event.EventListenerFunc(playerFashionActivate))
}
