package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/shangguzhiling/pbutil"
	playershangguzhiling "fgame/fgame/game/shangguzhiling/player"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
)

// 同步上古之灵属性变化
func LingShouPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeShangGuZhiLing.Mask())
	return
}

func SendLingShouInfo(pl player.Player) (err error) {
	lingShouManager := pl.GetPlayerDataManager(playertypes.PlayerShangguzhilingDataManagerType).(*playershangguzhiling.PlayerShangguzhilingDataManager)
	//只发送有用到过的
	lingshouList := lingShouManager.GetCurrentLingShouObjectList()
	unlockLingwenMap := make(map[shangguzhilingtypes.LingshouType][]shangguzhilingtypes.LingwenType)
	jiesuoLinglianPosMap := make(map[shangguzhilingtypes.LingshouType][]shangguzhilingtypes.LinglianPosType)
	for _, obj := range lingshouList {
		unlockLingwenMap[obj.GetLingShouType()] = lingShouManager.GetLingWenUnlockList(obj.GetLingShouType())
		jiesuoLinglianPosMap[obj.GetLingShouType()] = lingShouManager.GetLingLianPosJiesuoList(obj.GetLingShouType())
	}
	scMsg := pbutil.BuildSCShangguzhilingInfo(lingshouList, unlockLingwenMap, jiesuoLinglianPosMap)
	pl.SendMsg(scMsg)
	return nil
}
