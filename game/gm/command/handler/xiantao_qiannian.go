package handler

import (
	"fgame/fgame/common/lang"
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/xiantao/pbutil"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXianTaoQianNianCaiJiTimes, command.CommandHandlerFunc(handleSetXianTaoQianNianCaiJiTimes))
}

func handleSetXianTaoQianNianCaiJiTimes(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   num,
				"error": err,
			}).Warn("gm:处理千年仙桃采集点采集次数任务,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}

	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeXianTaoDaHui {
		return
	}

	npcMap := s.GetNPCS(scenetypes.BiologyScriptTypeXianTaoQianNianCollect)
	for _, npc := range npcMap {
		n, ok := npc.(*collectnpc.CollectPointNPC)
		if !ok {
			continue
		}
		obj := n.GetCollect()
		totalCount := obj.GetTotalCount()
		if totalCount > 0 && totalCount >= int32(num) {
			obj.GmSetUseCount(int32(num))
			plMap := s.GetAllPlayers()
			for _, ppl := range plMap {
				scMsg := pbutil.BuildSCXiantaoPeachPointChange(n)
				ppl.SendMsg(scMsg)
			}
		}
	}
	return
}
