package advance

import (
	"fgame/fgame/game/funcopen/funcopen"
	funcopenlogic "fgame/fgame/game/funcopen/logic"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeFuncOpen, guaji.GuaJiCheckHandlerFunc(funcOpenCheck))
}

func funcOpenCheck(pl player.Player) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeMingRiKaiQi) {
		return
	}

	for funcOpenType, _ := range funcopen.GetFuncOpenService().GetManualFuncOpenMap() {
		activeType, err := funcopenlogic.CheckFuncOpenByType(pl, funcOpenType)
		if err != nil {
			continue
		}
		if !activeType.Valid() {
			continue
		}
		funcopenlogic.HandleManualActive(pl, int32(funcOpenType))
	}

	return
}
