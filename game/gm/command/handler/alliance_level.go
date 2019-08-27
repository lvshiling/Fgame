package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/template"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeAllianceLevel, command.CommandHandlerFunc(handleAllianceJianShe))

}

func handleAllianceJianShe(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理设置仙盟等级,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	if level <= 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置仙盟等级,level小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	tempTemplateObject := template.GetTemplateService().Get(int(level), (*gametemplate.AllianceTemplate)(nil))
	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置仙盟等级,level模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	member := alliance.GetAllianceService().GetAllianceMember(pl.GetId())
	alObj := member.GetAlliance().GetAllianceObject()
	if int32(level) > alObj.GetLevel() {
		//增加
		totalBuild := int64(0)
		for l := alObj.GetLevel() + 1; l <= int32(level); l++ {
			tempLevelTemplate := template.GetTemplateService().Get(int(l), (*gametemplate.AllianceTemplate)(nil)).(*gametemplate.AllianceTemplate)
			totalBuild += tempLevelTemplate.UnionBuild
		}
		totalBuild -= member.GetAlliance().GetAllianceObject().GetJianShe()
		if totalBuild == 0 {
			return
		}

		member.GetAlliance().AddJianShe(totalBuild)
	} else {
		//减少
		totalBuild := int64(0)
		for l := alObj.GetLevel(); l > int32(level); l-- {
			tempLevelTemplate := template.GetTemplateService().Get(int(l), (*gametemplate.AllianceTemplate)(nil)).(*gametemplate.AllianceTemplate)
			totalBuild += int64(tempLevelTemplate.UnionBuild)
		}
		totalBuild += member.GetAlliance().GetAllianceObject().GetJianShe()
		if totalBuild == 0 {
			return
		}

		member.GetAlliance().CostJianShe(totalBuild)
	}

	propertylogic.SnapChangedProperty(pl)
	return
}
