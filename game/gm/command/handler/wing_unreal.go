package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	winglogic "fgame/fgame/game/wing/logic"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"
	"fgame/fgame/game/wing/wing"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeWingUnreal, command.CommandHandlerFunc(handleWingUnreal))

}

func handleWingUnreal(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	wingIdStr := c.Args[0]
	wingId, err := strconv.ParseInt(wingIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"wingId": wingId,
				"error":  err,
			}).Warn("gm:处理设置战翼幻化,wingId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := wing.GetWingService().GetWing(int(wingId))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"wingId": wingId,
				"error":  err,
			}).Warn("gm:处理设置战翼幻化,wingId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	manager.GmSetWingUnreal(int(wingId))

	//同步属性
	winglogic.WingPropertyChanged(pl)

	scWingUnreal := pbutil.BuildSCWingUnreal(int32(wingId))
	pl.SendMsg(scWingUnreal)
	return
}
