package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	lingyuventtypes "fgame/fgame/game/lingyu/event/types"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家领域进阶
func playerLingyuAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId := data.(int32)
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(advanceId)
	if lingyuTemplate == nil {
		return
	}

	preLingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(advanceId - 1)
	if preLingyuTemplate == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preLingyuTemplate.GetBattlePropertyMap())
	power := propertylogic.CulculateForce(lingyuTemplate.GetBattlePropertyMap())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	lingyuName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(lingyuTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.LingyuAdvancedNotice), playerName, lingyuName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(lingyuventtypes.EventTypeLingyuAdvanced, event.EventListenerFunc(playerLingyuAdavanced))
}
