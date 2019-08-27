package handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/template"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLevel, command.CommandHandlerFunc(handleLevel))
}

func handleLevel(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理设置等级,level不是数字")
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
			}).Warn("gm:处理设置等级,level小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := template.GetTemplateService().Get(int(level), (*gametemplate.CharacterLevelTemplate)(nil))
	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置等级,level模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	if int32(level) > pl.GetLevel() {
		//添加经验
		totalExp := int64(0)
		for l := pl.GetLevel() + 1; l <= int32(level); l++ {
			tempLevelTemplate := template.GetTemplateService().Get(int(l), (*gametemplate.CharacterLevelTemplate)(nil)).(*gametemplate.CharacterLevelTemplate)
			totalExp += int64(tempLevelTemplate.Experience)
		}
		totalExp -= manager.GetExp()
		if totalExp == 0 {
			return
		}
		reason := commonlog.LevelLogReasonGM
		reasonText := reason.String()
		manager.AddExp(totalExp, reason, reasonText)
	} else {
		//扣除经验
		totalExp := int64(0)
		for l := pl.GetLevel(); l > int32(level); l-- {
			tempLevelTemplate := template.GetTemplateService().Get(int(l), (*gametemplate.CharacterLevelTemplate)(nil)).(*gametemplate.CharacterLevelTemplate)
			totalExp += int64(tempLevelTemplate.Experience)
		}
		totalExp += manager.GetExp()
		if totalExp == 0 {
			return
		}
		reason := commonlog.LevelLogReasonGM
		reasonText := reason.String()
		manager.CostExp(totalExp, reason, reasonText)
	}

	propertylogic.SnapChangedProperty(pl)
	return
}
