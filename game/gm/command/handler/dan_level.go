package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/dan/dan"
	"fgame/fgame/game/dan/pbutil"
	playerdan "fgame/fgame/game/dan/player"
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
	command.Register(gmcommandtypes.CommandTypeDanLevel, command.CommandHandlerFunc(handleDanLevel))

}

func handleDanLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	danLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"danLevel": levelStr,
				"error":    err,
			}).Warn("gm:处理设置食丹等级,danLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if danLevel <= 1 {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"danLevel": levelStr,
				"error":    err,
			}).Warn("gm:处理设置食丹等级,danLevel小于等于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := dan.GetDanService().GetEatDan(int(danLevel))
	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"danLevel": levelStr,
				"error":    err,
			}).Warn("gm:处理设置食丹等级,danLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	manager.GmSetDanLevel(int32(danLevel))

	danInfo := manager.GetDanInfo()
	scDanUpgrade := pbuitl.BuildSCDanUpgrade(danInfo.LevelId)
	pl.SendMsg(scDanUpgrade)
	return
}
