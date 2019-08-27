package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	massacreventtypes "fgame/fgame/game/massacre/event/types"
	massacretemplate "fgame/fgame/game/massacre/template"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/weapon/weapon"
	"fmt"
)

//玩家戮仙刃进阶
func playerMassacreAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*massacreventtypes.PlayerMassacreAdvancedEventData)
	oldAdvanceId := eventData.GetOldAdvanceId()
	advanceId := eventData.GetNewAdvanceId()
	massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(int(advanceId))
	if massacreTemplate == nil {
		return
	}
	hasWeaponAfter := false
	if massacreTemplate.WeaponId != 0 {
		hasWeaponAfter = true
	}

	hasWeaponBefore := false
	prePower := int64(0)
	preMassacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(int(oldAdvanceId))
	if preMassacreTemplate != nil {
		prePower = propertylogic.CulculateForce(preMassacreTemplate.GetBattleProperty())
		if preMassacreTemplate.WeaponId != 0 {
			hasWeaponBefore = true
		}
	}

	//激活冰魂
	if !hasWeaponBefore && hasWeaponAfter {
		data := massacreventtypes.CreatePlayerMassacreWeaponEventData(massacreTemplate.WeaponId, true)
		gameevent.Emit(massacreventtypes.EventTypeMassacreWeapon, pl, data)
	}

	weaponLev := massacretemplate.GetMassacreTemplateService().GetMassacreeWeaponLev()
	power := propertylogic.CulculateForce(massacreTemplate.GetBattleProperty())
	diffPower := power - prePower

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	massacreName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(fmt.Sprintf("%d阶%d星", massacreTemplate.Type, massacreTemplate.Star)))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", diffPower))
	weaponLevStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", weaponLev))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MassacreAdvancedOneNotice), playerName, massacreName, powerStr, weaponLevStr)
	if weaponLev == massacreTemplate.Type {
		weaponTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(massacreTemplate.WeaponId))
		if weaponTemplate == nil {
			return
		}
		weaponName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(weaponTemplate.GetName(pl.GetRole())))
		content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.MassacreAdvancedWeaponNotice), playerName, massacreName, powerStr, weaponName)
	} else if weaponLev < massacreTemplate.Type {
		content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.MassacreAdvancedTwoNotice), playerName, massacreName, powerStr)
	}
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(massacreventtypes.EventTypeMassacreAdvanced, event.EventListenerFunc(playerMassacreAdavanced))
}
