package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chattypes "fgame/fgame/game/chat/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	titleeventtypes "fgame/fgame/game/title/event/types"
	titletemplate "fgame/fgame/game/title/title"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家活动称号失效
func playerTitleTimeExpire(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*titleeventtypes.PlayerTitleTimeExpireEventData)

	// 称号失效发邮件通知玩家
	t := timeutils.MillisecondToTime(eventData.GetExpireTime())
	title := lang.GetLangService().ReadLang(lang.EmailTitleTimeExpireTitle)
	content := lang.GetLangService().ReadLang(lang.EmailTitleTimeExpireContent)

	titleTemp := titletemplate.GetTitleService().GetTitleTemplate(int(eventData.GetTitleId()))
	titleName := coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, titleTemp.Name)
	battlePropertyMap := titleTemp.GetBattleAttrTemplate().GetAllBattleProperty()
	power := propertylogic.CulculateForce(battlePropertyMap)
	powerStr := coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("%d", power))

	contentStr := fmt.Sprintf(content, titleName, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), powerStr)
	emaillogic.AddEmail(pl, title, contentStr, nil)

	return
}

func init() {
	gameevent.AddEventListener(titleeventtypes.EventTypeTitleTimeExpire, event.EventListenerFunc(playerTitleTimeExpire))
}
