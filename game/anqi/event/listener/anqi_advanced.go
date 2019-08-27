package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	anqiventtypes "fgame/fgame/game/anqi/event/types"
	anqitemplate "fgame/fgame/game/anqi/template"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家暗器进阶
func playerAnqiAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId := data.(int32)
	anqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(advanceId)
	if anqiTemplate == nil {
		return
	}

	preAnqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(advanceId - 1)
	if preAnqiTemplate == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preAnqiTemplate.GetBattleProperty())
	power := propertylogic.CulculateForce(anqiTemplate.GetBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	anqiName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(anqiTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AnqiAdvancedNotice), playerName, anqiName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(anqiventtypes.EventTypeAnqiAdvanced, event.EventListenerFunc(playerAnqiAdavanced))
}
