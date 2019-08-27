package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	fabaotemplate "fgame/fgame/game/fabao/template"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家法宝进阶
func playerFaBaoAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}
	fabaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(advanceId))
	if fabaoTemplate == nil {
		return
	}

	preFaBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(advanceId) - 1)
	if preFaBaoTemplate == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preFaBaoTemplate.GetBattleProperty())
	power := propertylogic.CulculateForce(fabaoTemplate.GetBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	faBaoName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(fabaoTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FaBaoAdvancedNotice), playerName, faBaoName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(fabaoeventtypes.EventTypeFaBaoAdvanced, event.EventListenerFunc(playerFaBaoAdavanced))
}
