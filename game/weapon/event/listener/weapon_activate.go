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
	weaponeventtypes "fgame/fgame/game/weapon/event/types"
	"fgame/fgame/game/weapon/weapon"
	"fmt"
)

//玩家兵魂激活
func playerWeaponActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	weaponId := data.(int32)
	weaponTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if weaponTemplate == nil {
		return
	}
	attrTemp := weaponTemplate.GetBattleAttrTemplate()
	if attrTemp == nil {
		return
	}

	power := propertylogic.CulculateForce(attrTemp.GetAllBattleProperty())

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	weaponName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(weaponTemplate.GetName(pl.GetRole())))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", power))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.WeaponActivateNotice), playerName, weaponName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(weaponeventtypes.EventTypeWeaponActivate, event.EventListenerFunc(playerWeaponActivate))
}
