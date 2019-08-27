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
	command.Register(gmcommandtypes.CommandTypeWingUnrealDan, command.CommandHandlerFunc(handleWingUnrealDan))

}

func handleWingUnrealDan(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	unrealDanStr := c.Args[0]
	unrealDanLevel, err := strconv.ParseInt(unrealDanStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"unrealDanLevel": unrealDanLevel,
				"error":          err,
			}).Warn("gm:处理设置战翼食幻化丹等级,unrealDanLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := wing.GetWingService().GetWingHuanHuaTemplate(int32(unrealDanLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"unrealDanLevel": unrealDanLevel,
				"error":          err,
			}).Warn("gm:处理设置战翼食幻化丹等级,unrealDanLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	manager.GmSetWingUnrealDanLevel(int32(unrealDanLevel))

	//同步属性
	winglogic.WingPropertyChanged(pl)

	scWingShiDan := pbutil.BuildSCWingUnrealDan(int32(unrealDanLevel), 0)
	pl.SendMsg(scWingShiDan)
	return
}
