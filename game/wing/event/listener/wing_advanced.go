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
	wingeventtypes "fgame/fgame/game/wing/event/types"
	"fgame/fgame/game/wing/wing"
	"fmt"
)

//玩家战翼进阶
func playerWingAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int)
	if !ok {
		return
	}
	wingTemplate := wing.GetWingService().GetWingNumber(int32(advanceId))
	if wingTemplate == nil {
		return
	}
	attrTemp := wingTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}

	preWingTemplate := wing.GetWingService().GetWingNumber(int32(advanceId) - 1)
	if preWingTemplate == nil {
		return
	}
	preAttrTemp := preWingTemplate.GetBattleAttrTemplate()
	if preAttrTemp == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preAttrTemp.GetAllBattleProperty())
	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	wingName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(wingTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.WingAdvancedNotice), playerName, wingName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingAdvanced, event.EventListenerFunc(playerWingAdavanced))
}
