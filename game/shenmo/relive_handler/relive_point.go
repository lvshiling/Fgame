package relive_handler

import (
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterRelivePointHandler(scenetypes.SceneTypeCrossShenMo, scene.RelivePointHandlerFunc(RelivePoint))
}

//复活点复活
func RelivePoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	pos := s.MapTemplate().GetBornPos()
	pl.Reborn(pos)
	return true
}
