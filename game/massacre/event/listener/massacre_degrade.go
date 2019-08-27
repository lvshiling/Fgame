package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	massacreeventtypes "fgame/fgame/game/massacre/event/types"
	massacreventtypes "fgame/fgame/game/massacre/event/types"
	"fgame/fgame/game/massacre/pbutil"
	massacretemplate "fgame/fgame/game/massacre/template"
	"fgame/fgame/game/player"
)

//玩家戮仙刃降阶
func playerMassacreDegrade(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	playerMassacreDegradeEventData := data.(*massacreeventtypes.PlayerMassacreDegradeEventData)
	oldAdvanceId := playerMassacreDegradeEventData.GetOldAdvanceId()
	newAdvanceId := playerMassacreDegradeEventData.GetNewAdvanceId()
	attackName := playerMassacreDegradeEventData.GetAttackName()
	hasWeaponBefore := false
	oldAdvanceTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(int(oldAdvanceId))
	if oldAdvanceTemplate.WeaponId != 0 {
		hasWeaponBefore = true
	}
	hasWeaponAfter := false
	newAdvanceTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(int(newAdvanceId))
	if newAdvanceTemplate != nil {
		if newAdvanceTemplate.WeaponId != 0 {
			hasWeaponAfter = true
		}
	}

	//冰魂变化
	if hasWeaponBefore && !hasWeaponAfter {
		eventData := massacreeventtypes.CreatePlayerMassacreWeaponEventData(oldAdvanceTemplate.WeaponId, false)
		gameevent.Emit(massacreeventtypes.EventTypeMassacreWeapon, pl, eventData)
		if attackName != "" {
			//告诉前端兵魂变化了
			scMassacreWeaponLose := pbutil.BuildSCMassacreWeaponLose(int32(oldAdvanceId), int32(newAdvanceId), attackName)
			pl.SendMsg(scMassacreWeaponLose)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(massacreventtypes.EventTypeMassacreDegrade, event.EventListenerFunc(playerMassacreDegrade))
}
