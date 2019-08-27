package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	"fgame/fgame/game/mount/mount"
	mounttypes "fgame/fgame/game/mount/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家幻化激活
func playerMountUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	mountId := data.(int)
	mountTemplate := mount.GetMountService().GetMount(mountId)
	if mountTemplate == nil {
		return
	}
	if mountTemplate.GetTyp() != mounttypes.MountTypeShenLong {
		return
	}

	attrTemp := mountTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}

	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	mountName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(mountTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, coreutils.FormatNoticeStr(fmt.Sprintf("%d", power)))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MountUnrealActivateNotice), playerName, mountName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountUnrealActivate, event.EventListenerFunc(playerMountUnrealActivate))
}
