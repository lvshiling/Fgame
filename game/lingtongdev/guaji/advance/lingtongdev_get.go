package advance

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

func init() {
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeLingTongWeapon, guaji.GuaJiAdvanceGetHandlerFunc(lingTongDevGet))
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeLingTongMount, guaji.GuaJiAdvanceGetHandlerFunc(lingTongDevGet))
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeLingTongWing, guaji.GuaJiAdvanceGetHandlerFunc(lingTongDevGet))
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeLingTongShenFa, guaji.GuaJiAdvanceGetHandlerFunc(lingTongDevGet))
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeLingTongLingYu, guaji.GuaJiAdvanceGetHandlerFunc(lingTongDevGet))
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeLingTongFaBao, guaji.GuaJiAdvanceGetHandlerFunc(lingTongDevGet))
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeLingTongXianTi, guaji.GuaJiAdvanceGetHandlerFunc(lingTongDevGet))
}

func lingTongDevGet(pl player.Player, typ guajitypes.GuaJiAdvanceType) int32 {
	lingTongDevType := getLingTongDevType(typ)
	if !lingTongDevType.Vaild() {
		return 0
	}
	if !pl.IsFuncOpen(lingTongDevType.GetFuncOpenType()) {
		return 0
	}
	lingTongDevManager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	advanceId := lingTongDevManager.GetLingTongDevAdvancedId(lingTongDevType)
	return advanceId
}
