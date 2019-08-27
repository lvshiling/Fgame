package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	"fgame/fgame/game/lianyu/lianyu"
	"fgame/fgame/game/lianyu/pbutil"
	lianyuscene "fgame/fgame/game/lianyu/scene"
	"fgame/fgame/game/scene/scene"
)

//玩家进入无间炼狱
func lianYuPlayerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(lianyuscene.LianYuSceneData)
	if !ok {
		return
	}
	pl, ok := data.(scene.Player)
	if !ok {
		return
	}
	rankList := sd.GetRank()
	boss := sd.GetBoss()
	if boss == nil {
		return
	}
	bossStatus := boss.GetBossStatus()
	bossNpc := boss.GetNpc()
	num := sd.GetScenePlayerNum()
	shaQiNum := sd.GetShaQiNum(pl.GetId())
	lianyu.GetLianYuService().SyncSceneNum(num)

	scLianYuGet := pbutil.BuildSCLianYuGet(bossNpc, int32(bossStatus), rankList, shaQiNum)
	pl.SendMsg(scLianYuGet)
	return
}

func init() {
	gameevent.AddEventListener(lianyueventtypes.EventTypeLianYuPlayerEnter, event.EventListenerFunc(lianYuPlayerEnterScene))
}
