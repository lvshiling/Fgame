package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	huiyuaneventtypes "fgame/fgame/game/huiyuan/event/types"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fmt"
)

//玩家购买至尊会员
func playerHuiYuanBuy(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	huiyuanType := data.(huiyuantypes.HuiYuanType)
	if huiyuanType != huiyuantypes.HuiYuanTypePlus {
		return
	}

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	rewardsStr := coreutils.FormatColor(chattypes.ColorTypeModuleName, "288元宝大礼")

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.HuiYuanBuyNotice), playerName, rewardsStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(huiyuaneventtypes.EventTypeHuiYuanBuy, event.EventListenerFunc(playerHuiYuanBuy))
}
