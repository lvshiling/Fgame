package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	longgonglogic "fgame/fgame/game/longgong/logic"
	"fgame/fgame/game/longgong/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLongGongPlayerCaiJiTimes, command.CommandHandlerFunc(handleSetLongGongPlayerCaiJiTimes))
}

func handleSetLongGongPlayerCaiJiTimes(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理龙宫玩家财宝采集次数任务,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeLongGong {
		return
	}

	sd := s.SceneDelegate()
	longgongSd, ok := sd.(longgonglogic.LongGongSceneData)
	if !ok {
		return
	}
	longgongSd.GmSetPlayerTreasureCollectCount(pl.GetId(), int32(num))

	pCollectCount := longgongSd.GetPlayerTreasureCollectCount(pl.GetId())
	scMsg := pbutil.BuildSCLonggongPlayerValChange(pCollectCount)
	pl.SendMsg(scMsg)
	return
}
