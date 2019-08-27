package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gem/pbutil"
	playergem "fgame/fgame/game/gem/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMineStone, command.CommandHandlerFunc(handleMineStone))

}

func handleMineStone(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	stoneStr := c.Args[0]
	mineStone, err := strconv.ParseInt(stoneStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"stoneStr": stoneStr,
				"error":    err,
			}).Warn("gm:处理设置矿山原石,stoneStr不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if mineStone <= 0 {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"mineStone": mineStone,
				"error":     err,
			}).Warn("gm:处理设置矿山原石,mineStone小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	gemManager := pl.GetPlayerDataManager(types.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	gemManager.GmSetMineStone(int32(mineStone))

	mine := gemManager.GetMine()
	scGemMineActive := pbutil.BuildSCGemMineGet(mine)
	pl.SendMsg(scGemMineActive)
	return nil
}
