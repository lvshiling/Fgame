package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/material/pbutil"
	playermaterial "fgame/fgame/game/material/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	materialManager := p.GetPlayerDataManager(types.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)

	//刷新数据
	materialManager.RefreshData()

	//获取信息List
	materialArr := materialManager.GetPlayerMaterialInfoList()
	scMsg := pbutil.BuildSCMaterialInfoGet(materialArr)
	p.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
