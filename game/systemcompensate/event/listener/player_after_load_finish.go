package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/systemcompensate/systemcompensate"
	systemcompensatetemplate "fgame/fgame/game/systemcompensate/template"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	"fmt"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	if !pl.IsSystemCompensate() {
		for min := systemcompensatetypes.MinSysCompensate; min <= systemcompensatetypes.MaxSysCompensate; min++ {
			number := systemcompensate.GetSystemAdvancedNum(pl, min)
			temp := systemcompensatetemplate.GetReturnXiTongTemplateService().GetSystemCompensateTemplate(min, number)
			if temp == nil {
				continue
			}

			attachmentInfo := temp.GetReturnItemMap()
			if len(attachmentInfo) == 0 {
				continue
			}

			title := lang.GetLangService().ReadLang(lang.SystemCompensateMailTitle)
			content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.SystemCompensateMailContent), min.String())
			emaillogic.AddEmail(pl, title, content, attachmentInfo)
		}

		pl.SendSystemCompensate()
	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
