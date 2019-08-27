package relive_handler

import (
	"fgame/fgame/game/lianyu/lianyu"
	lianyuscene "fgame/fgame/game/lianyu/scene"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterRelivePointHandler(scenetypes.SceneTypeCrossLianYu, scene.RelivePointHandlerFunc(RelivePoint))
}

//复活点复活
func RelivePoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate()
	_, ok := sd.(lianyuscene.LianYuSceneData)
	if !ok {
		return
	}
	pos, flag := lianyu.GetLianYuService().GetRebornPos(pl.GetId())
	if !flag {
		return false
	}
	pl.Reborn(pos)
	return true
}
