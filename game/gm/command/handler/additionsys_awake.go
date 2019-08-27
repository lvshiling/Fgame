package handler

import (
	"fgame/fgame/common/lang"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeAdditionsysAwake, command.CommandHandlerFunc(handleAdditionSysAwake))
}

func handleAdditionSysAwake(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	typStr := c.Args[0]
	levelStr := c.Args[1]

	typ, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    pl.GetId(),
			"typ":   typ,
			"error": err,
		}).Warn("additionsysGm:处理设置附加系统觉醒，参数错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//参数验证
	sysType := additionsystypes.AdditionSysType(typ)
	if !sysType.Valid() {
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("additionsysGm:系统类型,错误")
		return
	}

	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    pl.GetId(),
			"level": level,
			"error": err,
		}).Warn("additionsysGm:处理设置附加系统觉醒，参数错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	temp := additionsystemplate.GetAdditionSysTemplateService().GetAwakeByType(sysType)
	if temp == nil {
		log.WithFields(
			log.Fields{
				"id":  pl.GetId(),
				"typ": typ,
			}).Warn("additionsysGm:处理设置附加系统觉醒,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	flag := manager.GmSetAwake(sysType)
	if !flag {
		log.WithFields(
			log.Fields{
				"id":  pl.GetId(),
				"typ": typ,
			}).Warn("additionsysGm:处理设置附加系统觉醒,错误")
		return
	}

	//同步属性
	additionsyslogic.UpdataAdditionSysPropertyByType(pl, sysType)

	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCAdditionSysAwakeEat(int32(sysType), int32(level))
	pl.SendMsg(scMsg)

	return
}
