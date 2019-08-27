package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	weaponeventtypes "fgame/fgame/game/weapon/event/types"
	playerweapon "fgame/fgame/game/weapon/player"
)

//兵魂改变
func weaponChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	if marryInfo.Status == marrytypes.MarryStatusTypeUnmarried ||
		marryInfo.Status == marrytypes.MarryStatusTypeDivorce {
		return
	}

	weaponManager := pl.GetPlayerDataManager(playertypes.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	weaponId := weaponManager.GetWeaponWear()
	scMsg := pbuitl.BuildSCMarryWeaponChange(pl.GetId(), weaponId)
	pl.SendMsg(scMsg)

	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	if spl == nil {
		return
	}

	spl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(weaponeventtypes.EventTypeWeaponChanged, event.EventListenerFunc(weaponChanged))
}
