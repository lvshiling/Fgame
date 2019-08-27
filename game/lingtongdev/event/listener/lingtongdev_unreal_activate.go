package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	lingtongdevplayer "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家灵童养成类幻化激活
func playerLingTongDevUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTongDevObj := data.(*lingtongdevplayer.PlayerLingTongDevObject)
	classType := lingTongDevObj.GetClassType()
	seqId := lingTongDevObj.GetSeqId()
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, int(seqId))
	if lingTongDevTemplate == nil {
		return
	}

	battlePropertyMap := lingTongDevTemplate.GetBattlePropertyMap()
	power := propertylogic.CulculateForce(battlePropertyMap)

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	lingTongDevName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(lingTongDevTemplate.GetName()))
	classTypeName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(classType.String()))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", power))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.LingTongDevUnrealActivateNotice), playerName, classTypeName, lingTongDevName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevUnrealActivate, event.EventListenerFunc(playerLingTongDevUnrealActivate))
}
