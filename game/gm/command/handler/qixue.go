package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	qixuelogic "fgame/fgame/game/qixue/logic"
	"fgame/fgame/game/qixue/pbutil"
	playerqixue "fgame/fgame/game/qixue/player"
	qixuetemplate "fgame/fgame/game/qixue/template"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeQiXueLevel, command.CommandHandlerFunc(handleQiXueLevel))
	command.Register(gmcommandtypes.CommandTypeQiXueShaLu, command.CommandHandlerFunc(handleSetShaLu))
}

func handleQiXueLevel(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理泣血枪任务,类型lev不是数字")
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
			}).Warn("gm:处理泣血枪任务,类型star不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(int32(lev), int32(star))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"lev":   lev,
				"star":  star,
				"error": err,
			}).Warn("gm:处理设置泣血枪阶别,泣血枪模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
	manager.GmSetQiXueAdvanced(int32(lev), int32(star))

	//同步属性
	qixuelogic.QiXuePropertyChanged(pl)

	scQiXueAdvanced := pbutil.BuildSCQiXueAdavanced(manager.GetQiXueInfo())
	pl.SendMsg(scQiXueAdvanced)
	return
}

func handleSetShaLu(p scene.Player, c *command.Command) (err error) {
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

	manager := pl.GetPlayerDataManager(types.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
	manager.GmSetQiXueShaLuNum(num)

	scQiXueGet := pbutil.BuildSCQiXueGet(manager.GetQiXueInfo())
	pl.SendMsg(scQiXueGet)
	return
}
