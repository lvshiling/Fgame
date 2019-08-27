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
	souleventtypes "fgame/fgame/game/soul/event/types"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
	"fmt"
)

func playerSoulAwaken(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	soulManager := pl.GetPlayerDataManager(playertypes.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	pl.SyncSoulAwakenNum(soulManager.GetAwakenNum())

	soulTag, ok := data.(soultypes.SoulType)
	if !ok {
		return
	}

	temp := soul.GetSoulService().GetSoulActiveTemplate(soulTag)
	if temp == nil {
		return
	}

	//公告
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	soulName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(temp.Name))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.SoulAwakenNotice), playerName, soulName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulAwaken, event.EventListenerFunc(playerSoulAwaken))
}
