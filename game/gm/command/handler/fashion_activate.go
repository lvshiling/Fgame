package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/fashion/fashion"
	fashionlogic "fgame/fgame/game/fashion/logic"
	"fgame/fgame/game/fashion/pbutil"
	playerfashion "fgame/fgame/game/fashion/player"
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
	command.Register(gmcommandtypes.CommandTypeFashionActive, command.CommandHandlerFunc(handleFashionActive))

}

func handleFashionActive(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	fashionStr := c.Args[0]
	fashionId, err := strconv.ParseInt(fashionStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Warn("gm:处理设置时装激活,fashionId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := fashion.GetFashionService().GetFashionTemplate(int(fashionId))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Warn("gm:处理设置时装激活,fashionId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	activeTime, flag := manager.GmFashionActive(int32(fashionId))
	if !flag {
		return
	}

	//同步属性
	fashionlogic.FashionPropertyChanged(pl)
	scFashionActive := pbutil.BuildSCFashionActive(int32(fashionId), activeTime)
	pl.SendMsg(scFashionActive)
	return
}
