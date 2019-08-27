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
	playertypes "fgame/fgame/game/player/types"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	playerxianfu "fgame/fgame/game/xianfu/player"
	xianfutypes "fgame/fgame/game/xianfu/types"
	"fmt"
)

//玩家仙府升级
func playerXianFuUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	xianfuType := data.(xianfutypes.XianfuType)
	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*playerxianfu.PlayerXinafuDataManager)
	xianfuLevel := xianfuManager.GetXianfuId(xianfuType)

	// 公告
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	xianfuName := coreutils.FormatColor(chattypes.ColorTypeModuleName, fmt.Sprintf("%s副本", xianfuType.String()))
	xianfuTypeName := coreutils.FormatColor(chattypes.ColorTypeModuleName, xianfuType.String())
	levelStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", xianfuLevel))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.XinafuUpgradeNotice), playerName, xianfuName, levelStr, xianfuTypeName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuUpgradeSuccess, event.EventListenerFunc(playerXianFuUpgrade))
}
