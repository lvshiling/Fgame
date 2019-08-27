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
	command.Register(gmcommandtypes.CommandTypeAdditionsysTongLingLev, command.CommandHandlerFunc(handleAdditionSysTongLingLev))
}

func handleAdditionSysTongLingLev(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	typStr := c.Args[0]
	levStr := c.Args[1]
	typ, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"typ":   typ,
				"error": err,
			}).Warn("gm:处理附加系统通灵任务,类型typ不是数字")
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
			}).Warn("gm:处理附加系统通灵任务,类型lev不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
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

	tempTemplatet := additionsystemplate.GetAdditionSysTemplateService().GetTongLingByLevel(int32(lev))
	if tempTemplatet == nil {
		log.WithFields(
			log.Fields{
				"id":  pl.GetId(),
				"typ": typ,
				"lev": lev,
			}).Warn("gm:处理设置附加系统部位通灵,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//修改等级
	manager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	flag := manager.GmSetTongLingLevel(sysType, tempTemplatet.GetLevel())
	if !flag {
		log.WithFields(
			log.Fields{
				"id":  pl.GetId(),
				"typ": typ,
				"lev": lev,
			}).Warn("gm:处理设置附加系统部位通灵,错误")
		return
	}

	//同步属性
	additionsyslogic.UpdataAdditionSysPropertyByType(pl, sysType)
	//同步改变
	propertylogic.SnapChangedProperty(pl)

	tongLingInfo := manager.GetAdditionSysTongLingInfoByType(sysType)
	scMsg := pbutil.BuildSCAdditionSysTongLingUpgrade(sysType, tongLingInfo.TongLingLev, tongLingInfo.TongLingPro)
	pl.SendMsg(scMsg)
	return
}
