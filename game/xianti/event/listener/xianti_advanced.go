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
	"fgame/fgame/game/xianti/xianti"
	"fmt"
)

//玩家仙体进阶
func playerXianTiAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int)
	if !ok {
		return
	}
	xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(advancedId))
	if xianTiTemplate == nil {
		return
	}

	preXianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(advancedId) - 1)
	if preXianTiTemplate == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preXianTiTemplate.GetBattleProperty())
	power := propertylogic.CulculateForce(xianTiTemplate.GetBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	xianTiName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(xianTiTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, coreutils.FormatNoticeStr(fmt.Sprintf("%d", diffPower)))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.XianTiAdvancedNotice), playerName, xianTiName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiAdvanced, event.EventListenerFunc(playerXianTiAdavanced))
}
