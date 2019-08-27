package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
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
	command.Register(gmcommandtypes.CommandTypeAllianceSkillLevel, command.CommandHandlerFunc(handleAllianceSkillLevel))

}

func handleAllianceSkillLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	typStr := c.Args[0]
	typ, err := strconv.ParseInt(typStr, 10, 64)
	skillType := alliancetypes.AllianceSkillType(typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"typ":   typStr,
				"error": err,
			}).Warn("gm:处理设置仙术等级,typ不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	if !skillType.Valid() {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"typ":   typStr,
				"error": err,
			}).Warn("gm:处理设置仙术等级,typ不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	levelStr := c.Args[1]
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置仙术等级,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	if level < 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置仙术等级,level小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	tem := alliancetemplate.GetAllianceTemplateService().GetAllianceSkillTemplateByType(int32(level), skillType)
	//修改等级
	if tem == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置仙术等级,level模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	allianceManager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.GMSetAllianceSkill(skillType, int32(level))

	id := int32(tem.TemplateId())
	scMsg := pbutil.BuildSCAllianceSkillUpgrade(id)
	pl.SendMsg(scMsg)

	propertylogic.SnapChangedProperty(pl)
	return

}
