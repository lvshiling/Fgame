package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/bodyshield/bodyshield"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家神盾尖刺进阶
func playerShieldAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	shieldId := data.(int32)
	shieldTemplate := bodyshield.GetBodyShieldService().GetShield(shieldId)
	if shieldTemplate == nil {
		return
	}
	attrTemp := shieldTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}

	preShieldTemplate := bodyshield.GetBodyShieldService().GetShield(shieldId - 1)
	if preShieldTemplate == nil {
		return
	}
	preAttrTemp := preShieldTemplate.GetBattleAttrTemplate()
	if preAttrTemp == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preAttrTemp.GetAllBattleProperty())
	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	featherName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(shieldTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ShieldAdvanceNotice), playerName, featherName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeShieldAdvanced, event.EventListenerFunc(playerShieldAdavanced))
}
