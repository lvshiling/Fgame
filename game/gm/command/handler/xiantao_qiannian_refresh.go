package handler

import (
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/xiantao/pbutil"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeRefreshXianTaoQianNian, command.CommandHandlerFunc(handleRefreshXianTaoQianNian))
}

func handleRefreshXianTaoQianNian(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	s := pl.GetScene()
	if s == nil {
		return
	}

	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeXianTaoDaHui {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	npcMap := s.GetNPCS(scenetypes.BiologyScriptTypeXianTaoQianNianCollect)
	for _, npc := range npcMap {
		n, ok := npc.(*collectnpc.CollectPointNPC)
		if !ok {
			continue
		}
		obj := n.GetCollect()
		totalCount := obj.GetTotalCount()
		if totalCount > 0 {
			obj.GmSetUseCount(0)
			obj.GmSetLastRecoverTime(now)
			plMap := s.GetAllPlayers()
			for _, ppl := range plMap {
				scMsg := pbutil.BuildSCXiantaoPeachPointChange(n)
				ppl.SendMsg(scMsg)
			}
		}
	}
	return
}
