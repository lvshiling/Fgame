package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

//玩家属性变化
func playerPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	propertyEffectorType, ok := data.(playerpropertytypes.PropertyEffectorType)
	if !ok {
		return
	}
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingTong,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongFashion,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongFaBao,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongWeapon,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongShenFa,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongXianTi,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongWing,
		playerpropertytypes.PlayerPropertyEffectorTypeLingTongLingYu:
	default:
		return
	}

	lingTongMountPower := lingtonglogic.GetLingTongMountPower(pl)
	lingTongWingPower := lingtonglogic.GetLingTongWingPower(pl)
	lingTongFaBaoPower := lingtonglogic.GetLingTongFaBaoPower(pl)
	lingTongXianTiPower := lingtonglogic.GetLingTongXianTiPower(pl)
	lingTongShenFaPower := lingtonglogic.GetLingTongShenFaPower(pl)
	lingTongWeaponPower := lingtonglogic.GetLingTongWeaponPower(pl)
	lingTongLingYuPower := lingtonglogic.GetLingTongLingYuPower(pl)
	lingTongPower := lingtonglogic.GetLingTongPower(pl)
	lingTongFashionPower := lingtonglogic.GetLingTongFashionPower(pl)

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	power := lingTongMountPower + lingTongWingPower + lingTongFaBaoPower + lingTongXianTiPower + lingTongShenFaPower + lingTongWeaponPower + lingTongLingYuPower + lingTongPower + lingTongFashionPower
	flag := manager.BasePower(lingTongPower + lingTongFashionPower)
	if flag {
		// 灵童基础战力
		scLingTongPowerNotice := pbutil.BuildSCLingTongPowerNotice(lingTongPower + lingTongFashionPower)
		pl.SendMsg(scLingTongPowerNotice)
	}

	manager.Power(power)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(playerPropertyChange))
}
