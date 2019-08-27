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
	wingtypes "fgame/fgame/game/wing/types"
	"fgame/fgame/game/wing/wing"
	"fmt"
)

//玩家幻化激活
func playerWingUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	wingId := data.(int)
	wingTemplate := wing.GetWingService().GetWing(wingId)
	if wingTemplate == nil {
		return
	}
	if wingTemplate.GetTyp() != wingtypes.WingTypeSkin {
		return
	}

	attrTemp := wingTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}

	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	wingName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(wingTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", power))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.WingUnrealActivateNotice), playerName, wingName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingUnrealActivate, event.EventListenerFunc(playerWingUnrealActivate))
}
