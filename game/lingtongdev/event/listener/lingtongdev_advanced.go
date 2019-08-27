package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家灵童养成类进阶
func playerLingTongDevAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTongObj, ok := data.(*playerlingtongdev.PlayerLingTongDevObject)
	if !ok {
		return
	}
	classType := lingTongObj.GetClassType()
	advanceId := lingTongObj.GetAdvancedId()
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, advanceId)
	if lingTongDevTemplate == nil {
		return
	}

	preLingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, advanceId-1)
	if preLingTongDevTemplate == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preLingTongDevTemplate.GetBattlePropertyMap())
	power := propertylogic.CulculateForce(lingTongDevTemplate.GetBattlePropertyMap())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	lingTongDevName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(lingTongDevTemplate.GetName()))
	classTypeName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(classType.String()))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.LingTongDevAdvancedNotice), playerName, classTypeName, lingTongDevName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevAdvanced, event.EventListenerFunc(playerLingTongDevAdavanced))
}
