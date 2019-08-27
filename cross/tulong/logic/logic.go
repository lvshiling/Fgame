package logic

import (
	"fgame/fgame/cross/tulong/pbutil"
	tulongscene "fgame/fgame/cross/tulong/scene"
	"fgame/fgame/game/scene/scene"
)

//龙蛋采集被打断
func CollectEggInterrupt(pl scene.Player) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}

	tuLongSceneData, ok := sd.(tulongscene.TuLongSceneData)
	if !ok {
		return
	}
	pos, hasCollect := tuLongSceneData.HasingCollectEgg(pl.GetId())
	if !hasCollect {
		return
	}
	tuLongSceneData.CollectEggInterrupt(pos)
	scTuLongCollectStop := pbutil.BuildSCTuLongCollectStop(pos)
	pl.SendMsg(scTuLongCollectStop)
}
