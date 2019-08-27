package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	npceventtypes "fgame/fgame/game/npc/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	xianfulogic "fgame/fgame/game/xianfu/logic"
	"fgame/fgame/game/xianfu/pbutil"
)

//经验副本Boss血量变化
func expBossHPChanged(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	s := n.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeExperience {
		return
	}

	sd := s.SceneDelegate().(xianfulogic.XianFuSceneData)
	fubenTemp := sd.GetCurTemplate()
	ownerId := sd.GetOwnerId()
	if fubenTemp.GetBossId() != int32(n.GetBiologyTemplate().TemplateId()) {
		return
	}
	pl := s.GetPlayer(ownerId)
	if pl == nil {
		return
	}

	curHp := n.GetHP()
	scMsg := pbutil.BuildSCXianfuBossHpChangedNotice(curHp)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCHPChanged, event.EventListenerFunc(expBossHPChanged))
}
