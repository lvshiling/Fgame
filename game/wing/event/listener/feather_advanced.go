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

//玩家护体仙羽进阶
func playerFeatherAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	featherId := data.(int32)
	featherTemplate := wing.GetWingService().GetFeather(featherId)
	if featherTemplate == nil {
		return
	}
	attrTemp := featherTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}

	preFeatherTemplate := wing.GetWingService().GetFeather(featherId - 1)
	if preFeatherTemplate == nil {
		return
	}
	preAttrTemp := preFeatherTemplate.GetBattleAttrTemplate()
	if preAttrTemp == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preAttrTemp.GetAllBattleProperty())
	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	featherName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(featherTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FeatherAdvancedNotice), playerName, featherName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeFeatherAdvanced, event.EventListenerFunc(playerFeatherAdavanced))
}
