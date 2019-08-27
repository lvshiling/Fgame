package handler

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	massacrelogic "fgame/fgame/game/massacre/logic"
	"fgame/fgame/game/massacre/pbutil"
	playermassacre "fgame/fgame/game/massacre/player"
	massacretemplate "fgame/fgame/game/massacre/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMassacreLevel, command.CommandHandlerFunc(handleMassacreLevel))
	command.Register(gmcommandtypes.CommandTypeSetShaQi, command.CommandHandlerFunc(handleSetShaQi))
}

func handleMassacreLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levStr := c.Args[0]
	starStr := c.Args[1]
	lev, err := strconv.ParseInt(levStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"lev":   lev,
				"error": err,
			}).Warn("gm:处理戮仙刃任务,类型lev不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	star, err := strconv.ParseInt(starStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"star":  star,
				"error": err,
			}).Warn("gm:处理戮仙刃任务,类型star不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := massacretemplate.GetMassacreTemplateService().GetMassacreNumber(int32(lev), int32(star))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"lev":   lev,
				"star":  star,
				"error": err,
			}).Warn("gm:处理设置戮仙刃阶别,戮仙刃模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	manager.GmSetMassacreAdvanced(int32(lev), int32(star))

	//同步属性
	massacrelogic.MassacrePropertyChanged(pl)
	advancedId := int32(tempTemplateObject.TemplateId())
	sqNum := manager.GetMassacreInfo().ShaQiNum
	resultType := int32(1)
	scMassacreAdvanced := pbutil.BuildSCMassacreAdavanced(advancedId, sqNum, commontypes.AdvancedTypeJinJieDan, resultType)
	pl.SendMsg(scMassacreAdvanced)
	return
}

func handleSetShaQi(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   num,
				"error": err,
			}).Warn("gm:处理杀气任务,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	manager.GmSetMassacreShaQiNum(num)

	scMassacreGet := pbutil.BuildSCMassacreGet(manager.GetMassacreInfo())
	pl.SendMsg(scMassacreGet)
	return
}
