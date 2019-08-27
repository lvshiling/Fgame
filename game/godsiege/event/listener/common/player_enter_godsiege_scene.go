package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	"fgame/fgame/game/godsiege/godsiege"
	"fgame/fgame/game/godsiege/pbutil"
	godsiegescene "fgame/fgame/game/godsiege/scene"
	"fgame/fgame/game/scene/scene"
)

//玩家进入神兽攻城
func godSiegePlayerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(godsiegescene.GodSiegeSceneData)
	if !ok {
		return
	}
	pl, ok := data.(scene.Player)
	if !ok {
		return
	}
	godType := sd.GetGodType()
	boss := sd.GetBoss()
	if boss == nil {
		return
	}
	bossStatus := boss.GetBossStatus()
	bossNpc := boss.GetNpc()
	num := sd.GetScenePlayerNum()
	itemMap := sd.GetItemMapByPlayer(pl)
	collectList := sd.GetCollectNpcList()

	godsiege.GetGodSiegeService().SyncSceneNum(godType, num)
	scGodSiegeGet := pbutil.BuildSCGodSiegeGet(pl, int32(godType), bossNpc, int32(bossStatus), itemMap, collectList)
	pl.SendMsg(scGodSiegeGet)

	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegePlayerEnter, event.EventListenerFunc(godSiegePlayerEnterScene))
}
