package relive_handler

import (
	crosstulong "fgame/fgame/cross/tulong/tulong"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	tulongtemplate "fgame/fgame/game/tulong/template"
	tulongtypes "fgame/fgame/game/tulong/types"
)

func init() {
	scene.RegisterReliveEntryPointHandler(scenetypes.SceneTypeCrossTuLong, scene.ReliveEntryPointHandlerFunc(ReliveEntryPoint))
}

//回进入点复活
func ReliveEntryPoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	bornBiaoShi, flag := crosstulong.GetTuLongService().GetPlayerBornBiaoShi(pl.GetAllianceId())
	if !flag {
		return false
	}

	tuLongPosTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongPosTemplate(tulongtypes.TuLongPosTypePlayer, bornBiaoShi)
	if tuLongPosTemplate == nil {
		return false
	}
	pos := tuLongPosTemplate.GetPos()
	pl.Reborn(pos)
	return true
}
