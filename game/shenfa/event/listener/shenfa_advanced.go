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
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	"fmt"
)

//玩家身法进阶
func playerShenfaAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId := data.(int32)
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(advanceId)
	if shenfaTemplate == nil {
		return
	}
	attrTemp := shenfaTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}

	preShenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(advanceId - 1)
	if preShenfaTemplate == nil {
		return
	}
	preAttrTemp := preShenfaTemplate.GetBattleAttrTemplate()
	if preAttrTemp == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preAttrTemp.GetAllBattleProperty())
	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	shenfaName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(shenfaTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ShenfaAdvancedNotice), playerName, shenfaName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaAdvanced, event.EventListenerFunc(playerShenfaAdavanced))
}
