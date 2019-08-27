package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

//玩家属性变化
func playerPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	propertyEffectorType, ok := data.(playerpropertytypes.PropertyEffectorType)
	if !ok {
		return
	}
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTongFaBao,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongWeapon,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongShenFa,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongXianTi,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongWing,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongLingYu:
	default:
		return
	}
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	powerLingTongMount(pl, propertyEffectorType)
	powerLingTongWing(pl, propertyEffectorType)
	powerLingTongWeapon(pl, propertyEffectorType)
	powerLingTongShenFa(pl, propertyEffectorType)
	powerLingTongLingYu(pl, propertyEffectorType)
	powerLingTongXianTi(pl, propertyEffectorType)
	powerLingTongFaBao(pl, propertyEffectorType)
	return
}

func powerLingTongMount(pl player.Player, propertyEffectorType playerpropertytypes.PropertyEffectorType) {
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount:
		manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount)
		manager.LingTongDevPower(types.LingTongDevSysTypeLingQi, power)
		return
	}
}

func powerLingTongWing(pl player.Player, propertyEffectorType playerpropertytypes.PropertyEffectorType) {
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTongWing:
		manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongWing)
		manager.LingTongDevPower(types.LingTongDevSysTypeLingYi, power)
		return
	}
}

func powerLingTongWeapon(pl player.Player, propertyEffectorType playerpropertytypes.PropertyEffectorType) {
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTongWeapon:
		manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongWeapon)
		manager.LingTongDevPower(types.LingTongDevSysTypeLingBing, power)
		return
	}
}

func powerLingTongShenFa(pl player.Player, propertyEffectorType playerpropertytypes.PropertyEffectorType) {
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTongShenFa:
		manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongShenFa)
		manager.LingTongDevPower(types.LingTongDevSysTypeLingShen, power)
		return
	}
}

func powerLingTongLingYu(pl player.Player, propertyEffectorType playerpropertytypes.PropertyEffectorType) {
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTongLingYu:
		manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongLingYu)
		manager.LingTongDevPower(types.LingTongDevSysTypeLingYu, power)
		return
	}
}

func powerLingTongXianTi(pl player.Player, propertyEffectorType playerpropertytypes.PropertyEffectorType) {
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTongXianTi:
		manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongXianTi)
		manager.LingTongDevPower(types.LingTongDevSysTypeLingTi, power)
		return
	}
}

func powerLingTongFaBao(pl player.Player, propertyEffectorType playerpropertytypes.PropertyEffectorType) {
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTongFaBao:
		manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongFaBao)
		manager.LingTongDevPower(types.LingTongDevSysTypeLingBao, power)
		return
	}
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(playerPropertyChange))
}
