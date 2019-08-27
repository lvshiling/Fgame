package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gem/gem"
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
	command.Register(gmcommandtypes.CommandTypeMineStorage, command.CommandHandlerFunc(handleMineStorage))

}

func handleMineStorage(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	storageStr := c.Args[0]
	mineStorage, err := strconv.ParseInt(storageStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"storageStr": storageStr,
				"error":      err,
			}).Warn("gm:处理设置矿山当前库存,storageStr不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	gemManager := pl.GetPlayerDataManager(types.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	level := gemManager.GetMineLevel()
	to := gem.GetGemService().GetMineTemplateByLevel(level)
	limitMax := int64(to.LimitMax)
	if mineStorage <= 0 || mineStorage > limitMax {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"storageStr": mineStorage,
				"error":      err,
			}).Warn("gm:处理设置矿山当前库存,storageStr小于等于0大于当前最大库存")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	gemManager.GmSetMineStorage(int32(mineStorage))

	mine := gemManager.GetMine()
	scGemMineActive := pbutil.BuildSCGemMineGet(mine)
	pl.SendMsg(scGemMineActive)
	return
}
