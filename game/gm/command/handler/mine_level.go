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
	command.Register(gmcommandtypes.CommandTypeMineLevel, command.CommandHandlerFunc(handleMineLevel))

}

func handleMineLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	mineLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"mineLevel": levelStr,
				"error":     err,
			}).Warn("gm:处理设置矿工等级,mineLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if mineLevel <= 0 {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"mineLevel": levelStr,
				"error":     err,
			}).Warn("gm:处理设置矿工等级,mineLevel小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := gem.GetGemService().GetMineTemplateByLevel(int32(mineLevel))
	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"mineLevel": levelStr,
				"error":     err,
			}).Warn("gm:处理设置矿工等级,mineLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	gemManager := pl.GetPlayerDataManager(types.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	gemManager.GmSetMineLevel(int32(mineLevel))

	mine := gemManager.GetMine()
	scGemMineActive := pbutil.BuildSCGemMineActive(mine)
	pl.SendMsg(scGemMineActive)
	return
}
