package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLingTongFashionActivate, command.CommandHandlerFunc(handleLingTongFashionActivate))

}

func handleLingTongFashionActivate(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	fashionIdStr := c.Args[0]
	fashionId, err := strconv.ParseInt(fashionIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Warn("gm:处理设置灵童时装,fashionId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(int32(fashionId))
	//时装灵童
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
			}).Warn("gm:处理设置灵童时装,fashionId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	flag := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(int32(fashionId))
	if flag {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
			}).Warn("gm:处理设置灵童时装,出生的时装")
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	fashionInfo := manager.GetFashionInfoById(int32(fashionId))
	if fashionInfo != nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:该时装已激活")
		return
	}
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:请先激活灵童时装激活系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	flag = manager.GmLingTongFashionActivate(int32(fashionId))
	if !flag {
		return
	}
	lingtonglogic.LingTongFashionPropertyChanged(pl)

	obj := manager.GetFashionInfoById(int32(fashionId))
	if obj == nil {
		return
	}
	scLingTongFashionActivate := pbutil.BuildSCLingTongFashionActivate(obj)
	pl.SendMsg(scLingTongFashionActivate)
	return
}
