package handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeGoldYuanLevel, command.CommandHandlerFunc(handleGoldYuanLevel))
}

func handleGoldYuanLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置元神等级,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if level <= 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置元神等级,level小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//修改元神等级
	temp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldYuanTemplate(int32(level))
	if temp == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置元神等级,level模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curLevel := manager.GetGoldYuanLevel()
	if int32(level) > curLevel {
		//添加经验
		totalExp := int64(0)
		for l := curLevel + 1; l <= int32(level); l++ {
			tempLevelTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetGoldYuanTemplate(l)
			totalExp += int64(tempLevelTemplate.Exp)
		}
		totalExp -= manager.GetGoldYuanExp()
		if totalExp == 0 {
			return
		}
		reason := commonlog.GoldYuanLevelLogReasonGM
		reasonText := reason.String()
		manager.AddGoldYuanExp(totalExp, reason, reasonText)
	} else {
		//扣除经验
		totalExp := int64(0)
		for l := curLevel; l > int32(level); l-- {
			tempLevelTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetGoldYuanTemplate(l)
			totalExp += int64(tempLevelTemplate.Exp)
		}
		totalExp += manager.GetGoldYuanExp()
		if totalExp == 0 {
			return
		}
		reason := commonlog.GoldYuanLevelLogReasonGM
		reasonText := reason.String()
		manager.CostGoldYuanExp(totalExp, reason, reasonText)
	}

	propertylogic.SnapChangedProperty(pl)
	return
}
