package logic

import (
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongplayertypes "fgame/fgame/game/lingtong/player/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//灵童自身属性变更
func LingTongSelfAllPropertyChanged(pl player.Player) {
	//同步属性
	lingTongDataManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongDataManager.UpdateBattleProperty(lingtongplayertypes.PropertyEffectorTypeMaskAll)
	return
}

//灵童自身属性变更
func LingTongSelfFashionPropertyChanged(pl player.Player) {
	//同步属性
	lingTongDataManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongDataManager.UpdateBattleProperty(lingtongplayertypes.LingTongPropertyEffectorTypeFashion.Mask())
	return
}
