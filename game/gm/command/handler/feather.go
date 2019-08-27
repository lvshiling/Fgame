package handler

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
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
	command.Register(gmcommandtypes.CommandTypeFeatherLevel, command.CommandHandlerFunc(handleFeatherLevel))

}

func handleFeatherLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	featherId, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"featherId": featherId,
				"error":     err,
			}).Warn("gm:处理设置护体仙羽阶别,featherId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := wing.GetWingService().GetFeather(int32(featherId))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"featherId": featherId,
				"error":     err,
			}).Warn("gm:处理设置护体仙羽阶别,featherId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	manager.GmSetFeatherId(int32(featherId))

	//同步属性
	winglogic.WingPropertyChanged(pl)

	scFeatherAdvanced := pbutil.BuildSCFeatherAdvanced(int32(featherId), 0, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scFeatherAdvanced)
	return
}
