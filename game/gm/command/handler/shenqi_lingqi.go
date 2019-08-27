package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeShenQiLingQiNum, command.CommandHandlerFunc(handleSetShenQiLingQiNum))
}

func handleSetShenQiLingQiNum(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理灵气值任务,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	flag := manager.GmSetLingQiNum(num)
	if !flag {
		return
	}

	debrisMap := manager.GetShenQiDebrisMap()
	smeltMap := manager.GetShenQiSmeltMap()
	qiLingMap := manager.GetShenQiQiLingMap()
	shenQiOjb := manager.GetShenQiOjb()
	scMsg := pbutil.BuildSCShenQiInfoGet(qiLingMap, debrisMap, smeltMap, shenQiOjb.LingQiNum)
	pl.SendMsg(scMsg)
	return
}
