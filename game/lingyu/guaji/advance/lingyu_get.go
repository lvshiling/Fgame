package advance

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

func init() {
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeLingyu, guaji.GuaJiAdvanceGetHandlerFunc(lingyuGet))
}

func lingyuGet(pl player.Player, typ guajitypes.GuaJiAdvanceType) int32 {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeLingYuAdvanced) {
		return 0
	}
	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()
	return int32(lingyuInfo.AdvanceId)
}
