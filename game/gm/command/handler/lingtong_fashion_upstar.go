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
	command.Register(gmcommandtypes.CommandTypeLingTongFashionUpstar, command.CommandHandlerFunc(handleLingTongFashionUpstar))

}

func handleLingTongFashionUpstar(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	fashionIdStr := c.Args[0]
	levelStr := c.Args[1]
	fashionId, err := strconv.ParseInt(fashionIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置灵童灵童时装星级,fashionId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置灵童灵童时装星级,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(int32(fashionId))
	//灵童时装
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置灵童灵童时装星级,fashionId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	if tempTemplateObject.LingTongUpstarId == 0 {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置灵童灵童时装星级,该时装无法升星")
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

	lingTongFashionUpstarTemplate := tempTemplateObject.GetLingTongFashionUpstarByLevel(int32(level))
	if lingTongFashionUpstarTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置灵童灵童时装星级,星级模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	fashionInfo := manager.GetFashionInfoById(int32(fashionId))
	if fashionInfo == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"fashionId": fashionId,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置灵童灵童时装星级,请先激活该时装")
		return
	}

	flag = manager.GmLingTongFashionUpstar(int32(fashionId), int32(level))
	if !flag {
		return
	}
	lingtonglogic.LingTongPropertyChanged(pl)

	scLingTongFashionUpstar := pbutil.BuildSCLingTongFashionUpstar(int32(fashionId), fashionInfo.GetUpgradeLevel(), fashionInfo.GetUpgradePro(), true)
	pl.SendMsg(scLingTongFashionUpstar)
	return
}
