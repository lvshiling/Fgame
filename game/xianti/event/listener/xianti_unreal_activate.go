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
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	xiantitypes "fgame/fgame/game/xianti/types"
	"fgame/fgame/game/xianti/xianti"
	"fmt"
)

//玩家幻化激活
func playerXianTiUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	xianTiId := data.(int)
	xianTiTemplate := xianti.GetXianTiService().GetXianTi(xianTiId)
	if xianTiTemplate == nil {
		return
	}
	if xianTiTemplate.GetTyp() != xiantitypes.XianTiTypeSkin {
		return
	}

	power := propertylogic.CulculateForce(xianTiTemplate.GetBattleProperty())

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	xianTiName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(xianTiTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", power))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.XianTiUnrealActivateNotice), playerName, xianTiName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiUnrealActivate, event.EventListenerFunc(playerXianTiUnrealActivate))
}
