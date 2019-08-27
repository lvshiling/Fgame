package handler

import (
	"fgame/fgame/common/lang"
	dianxinglogic "fgame/fgame/game/dianxing/logic"
	"fgame/fgame/game/dianxing/pbutil"
	playerdianxing "fgame/fgame/game/dianxing/player"
	dianxingtemplate "fgame/fgame/game/dianxing/template"
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
	command.Register(gmcommandtypes.CommandTypeDianXing, command.CommandHandlerFunc(handleDianXing))
	command.Register(gmcommandtypes.CommandTypeDianXingJieFeng, command.CommandHandlerFunc(handleDianXingJieFeng))
	command.Register(gmcommandtypes.CommandTypeSetXingChen, command.CommandHandlerFunc(handleSetXingChen))
}

func handleDianXing(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	xingPuStr := c.Args[0]
	levStr := c.Args[1]
	xingPu, err := strconv.ParseInt(xingPuStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"xingPu": xingPu,
				"error":  err,
			}).Warn("gm:处理点星任务,类型xingPu不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	lev, err := strconv.ParseInt(levStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"lev":   lev,
				"error": err,
			}).Warn("gm:处理点星任务,类型lev不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := dianxingtemplate.GetDianXingTemplateService().GetDianXingTemplateByArg(int32(xingPu), int32(lev))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"xingPu": xingPu,
				"lev":    lev,
				"error":  err,
			}).Warn("gm:处理设置点星阶别,点星模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	manager.GmSetDianXingAdvanced(int32(xingPu), int32(lev))

	//同步属性
	dianxinglogic.DianXingPropertyChanged(pl)
	scDianXingGet := pbutil.BuildSCDianXingGet(manager.GetDianXingObject())
	pl.SendMsg(scDianXingGet)
	return
}

func handleDianXingJieFeng(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levStr := c.Args[0]
	lev, err := strconv.ParseInt(levStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"lev":   lev,
				"error": err,
			}).Warn("gm:处理点星解封任务,类型lev不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := dianxingtemplate.GetDianXingTemplateService().GetDianXingJieFengTemplateByLev(int32(lev))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"lev":   lev,
				"error": err,
			}).Warn("gm:处理设置点星解封阶别,点星模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	manager.GmSetDianXingJieFengAdvanced(int32(lev))

	//同步属性
	dianxinglogic.DianXingPropertyChanged(pl)
	scDianXingGet := pbutil.BuildSCDianXingGet(manager.GetDianXingObject())
	pl.SendMsg(scDianXingGet)
	return
}

func handleSetXingChen(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理星尘任务,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	manager.GmSetXingChenNum(num)

	scDianXingGet := pbutil.BuildSCDianXingGet(manager.GetDianXingObject())
	pl.SendMsg(scDianXingGet)
	return
}
