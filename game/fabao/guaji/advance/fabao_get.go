package advance

import (
	playerfabao "fgame/fgame/game/fabao/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

func init() {
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeFabao, guaji.GuaJiAdvanceGetHandlerFunc(fabaoGet))
}

func fabaoGet(pl player.Player, typ guajitypes.GuaJiAdvanceType) int32 {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeFaBao) {
		return 0
	}
	fabaoManager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	fabaoInfo := fabaoManager.GetFaBaoInfo()
	return int32(fabaoInfo.GetAdvancedId())
}
