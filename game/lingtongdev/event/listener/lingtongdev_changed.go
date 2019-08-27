package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家灵童养成类变化
func playerLingTongDevChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTong := pl.GetLingTong()
	if lingTong == nil {
		return
	}
	classType, ok := data.(lingtongdevtypes.LingTongDevSysType)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	showId := manager.GetLingTongDevSeqId(classType)
	switch classType {
	case lingtongdevtypes.LingTongDevSysTypeLingBing:
		lingTong.SetLingTongWeapon(showId, 0)
		break
	case lingtongdevtypes.LingTongDevSysTypeLingQi:
		lingTong.SetLingTongMountId(showId)
		break
	case lingtongdevtypes.LingTongDevSysTypeLingYi:
		lingTong.SetLingTongWingId(showId)
		break
	case lingtongdevtypes.LingTongDevSysTypeLingShen:
		lingTong.SetLingTongShenFaId(showId)
		break
	case lingtongdevtypes.LingTongDevSysTypeLingYu:
		lingTong.SetLingTongLingYuId(showId)
		break
	case lingtongdevtypes.LingTongDevSysTypeLingBao:
		lingTong.SetLingTongFaBaoId(showId)
		break
	case lingtongdevtypes.LingTongDevSysTypeLingTi:
		lingTong.SetLingTongXianTiId(showId)
		break
	}

	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevChanged, event.EventListenerFunc(playerLingTongDevChanged))
}
