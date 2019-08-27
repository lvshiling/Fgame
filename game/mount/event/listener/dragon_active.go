package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	dragoneventtypes "fgame/fgame/game/dragon/event/types"
	playerdragon "fgame/fgame/game/dragon/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/mount/mount"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

//玩家神龙进阶
func dragonActive(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	mountId := data.(int32)

	to := mount.GetMountService().GetMount(int(mountId))
	if to == nil {
		return
	}

	//公告
	dragonManager := pl.GetPlayerDataManager(playertypes.PlayerDragonDataManagerType).(*playerdragon.PlayerDragonDataManager)
	dragonObj := dragonManager.GetDragon()
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	eatLevel := coreutils.FormatColor(chattypes.ColorTypeModuleName, fmt.Sprintf("%d", dragonObj.StageId))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MountDragonActivateNotice), playerName, eatLevel)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(dragoneventtypes.EventTypeDragonAdvanced, event.EventListenerFunc(dragonActive))
}
