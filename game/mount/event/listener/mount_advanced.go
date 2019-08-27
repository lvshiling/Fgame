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
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家坐骑进阶
func playerMountAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int)
	if !ok {
		return
	}
	mountTemplate := mount.GetMountService().GetMountNumber(int32(advancedId))
	if mountTemplate == nil {
		return
	}
	attrTemp := mountTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}

	preMountTemplate := mount.GetMountService().GetMountNumber(int32(advancedId) - 1)
	if preMountTemplate == nil {
		return
	}
	preAttrTemp := preMountTemplate.GetBattleAttrTemplate()
	if preAttrTemp == nil {
		return
	}

	prePower := propertylogic.CulculateForce(preAttrTemp.GetAllBattleProperty())
	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	mountName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(mountTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, coreutils.FormatNoticeStr(fmt.Sprintf("%d", diffPower)))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MountAdvancedNotice), playerName, mountName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountAdvanced, event.EventListenerFunc(playerMountAdavanced))
}
