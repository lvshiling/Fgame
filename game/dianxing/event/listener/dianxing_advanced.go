package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	dianxingeventtypes "fgame/fgame/game/dianxing/event/types"
	dianxingtemplate "fgame/fgame/game/dianxing/template"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fmt"
)

//玩家点星系统进阶
func playerDianXingAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(int32)

	dianxingTemplate := dianxingtemplate.GetDianXingTemplateService().GetDianXingTemplateByArg(eventData, int32(1))
	if dianxingTemplate == nil {
		return
	}

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	xingPuName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(dianxingTemplate.Name))
	args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(funcopentypes.FuncOpenTypeDianXing)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToDianXing, args)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.DianXingAdvancedNotice), playerName, xingPuName, link)

	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(dianxingeventtypes.EventTypeDianXingAdvanced, event.EventListenerFunc(playerDianXingAdavanced))
}
