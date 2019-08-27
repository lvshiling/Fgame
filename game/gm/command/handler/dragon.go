package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/dragon/dragon"
	dragonlogic "fgame/fgame/game/dragon/logic"
	"fgame/fgame/game/dragon/pbutil"
	playerdragon "fgame/fgame/game/dragon/player"
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
	command.Register(gmcommandtypes.CommandTypeDragon, command.CommandHandlerFunc(handleDragonStage))

}

func handleDragonStage(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	stageStr := c.Args[0]
	dragonStage, err := strconv.ParseInt(stageStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"dragonStage": stageStr,
				"error":       err,
			}).Warn("gm:处理设置神龙现世阶段,dragonStage不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if dragonStage <= 1 {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"dragonStage": stageStr,
				"error":       err,
			}).Warn("gm:处理设置神龙现世阶段,dragonStage小于等于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := dragon.GetDragonService().GetDragonTemplate(int32(dragonStage))
	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"dragonStage": stageStr,
				"error":       err,
			}).Warn("gm:处理设置神龙现世阶段,dragonStage模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerDragonDataManagerType).(*playerdragon.PlayerDragonDataManager)
	manager.GmSetDragonStage(int32(dragonStage))

	dragonlogic.DragonPropertyChanged(pl)

	dragonInfo := manager.GetDragon()
	scDragonGet := pbuitl.BuildSCDragonGet(dragonInfo)
	pl.SendMsg(scDragonGet)
	return
}
