package player

import (
	playertypes "fgame/fgame/game/player/types"
	playerrealm "fgame/fgame/game/realm/player"
	hallrealmtypes "fgame/fgame/game/welfare/hall/realm/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 等级冲刺
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeRealm, playerwelfare.ActivityObjInfoInitFunc(welfareRealmInitInfo))
}

func welfareRealmInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*hallrealmtypes.WelfareRealmChallengeInfo)
	info.RewRecord = []int32{}
	info.IsEmail = false

	realmManager := obj.GetPlayer().GetPlayerDataManager(playertypes.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	initLevel := realmManager.GetTianJieTaLevel()
	info.Level = initLevel
}
