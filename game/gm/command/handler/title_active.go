package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"
	"fmt"

	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTitleActive, command.CommandHandlerFunc(handleTitleActive))
}

func handleTitleActive(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理称号激活")
	if len(c.Args) < 1 {
		log.Warn("gm:称号激活,参数少于1")

		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	titleIdStr := c.Args[0]
	titleId, err := strconv.ParseInt(titleIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"titleId": titleIdStr,
			}).Warn("gm:称号激活id,titleId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//TODO 添加称号
	err = titleActive(pl, int32(titleId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"titleId": titleIdStr,
			}).Warn("gm:称号激活,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"titleId": titleIdStr,
		}).Debug("gm:处理称号激活完成")
	return
}

func titleActive(p scene.Player, titleId int32) error {
	pl := p.(player.Player)
	titleManager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	flag := titleManager.IsValid(titleId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:无效参数")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}
	flag = titleManager.IfTitleExist(titleId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:称号已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.TitleRepeatActive)
		return nil
	}

	titleObj, flag := titleManager.TitleActive(titleId)
	if !flag {
		panic(fmt.Errorf("titlegm: titleActive  should be ok"))
	}

	scTitleActive := pbutil.BuildSCTitleActive(titleId, titleObj.ActiveTime, titleObj.ValidTime)
	pl.SendMsg(scTitleActive)
	return nil
}
