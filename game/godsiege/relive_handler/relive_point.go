package relive_handler

import (
	"fgame/fgame/game/godsiege/godsiege"
	godsiegescene "fgame/fgame/game/godsiege/scene"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterRelivePointHandler(scenetypes.SceneTypeCrossGodSiege, scene.RelivePointHandlerFunc(RelivePoint))
	scene.RegisterRelivePointHandler(scenetypes.SceneTypeCrossDenseWat, scene.RelivePointHandlerFunc(RelivePoint))
}

//复活点复活
func RelivePoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	sd := s.SceneDelegate()
	sceneData, ok := sd.(godsiegescene.GodSiegeSceneData)
	if !ok {
		return
	}
	godType := sceneData.GetGodType()
	pos, flag := godsiege.GetGodSiegeService().GetRebornPos(godType, pl.GetId())
	if !flag {
		return false
	}
	pl.Reborn(pos)
	return true
}
