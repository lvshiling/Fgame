package logic

import (
	playerbagua "fgame/fgame/game/bagua/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
)

//获取前往击杀界面信息的逻辑
func CheckPlayerIfCanEnterBaGua(pl player.Player) (flag bool) {

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeBaGuaMiJing) {
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeBaGuaMiJing {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	if manager.IfFullLevel() {
		return
	}
	flag = true
	return
}
