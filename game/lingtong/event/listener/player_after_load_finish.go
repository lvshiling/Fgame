package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		return
	}

	lingTongId := lingTongObj.GetLingTongId()
	//灵童
	lingTongMap := manager.GetLingTongMap()
	lingTongFashion := manager.GetLingTongFashion()
	scLingTongGet := pbutil.BuildSCLingTongGet(lingTongId, lingTongMap, lingTongFashion)
	pl.SendMsg(scLingTongGet)

	//灵童时装
	fashionMap := manager.GetActivateFashionMap()
	scLingTongFashionGet := pbutil.BuildSCLingTongFashionGet(lingTongId, fashionMap)
	pl.SendMsg(scLingTongFashionGet)

	//更新灵童
	lingtonglogic.InitLingTong(pl)

	//巡游阶段隐藏灵童
	if pl.GetWeddingStatus() == int32(marrytypes.MarryWedStatusSelfTypeCruise) {
		pl.HiddenLingTong(true)
	} else {
		pl.HiddenLingTong(false)
	}

	// 灵童基础战力
	basePower := manager.GetBasePower()
	scLingTongPowerNotice := pbutil.BuildSCLingTongPowerNotice(basePower)
	pl.SendMsg(scLingTongPowerNotice)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
