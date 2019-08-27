package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/skill/pbutil"
	playerskill "fgame/fgame/game/skill/player"
	"strconv"

	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSkillChangeLevel, command.CommandHandlerFunc(handleSkillChangeLevel))
}

func handleSkillChangeLevel(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理改变动态技能等级")
	if len(c.Args) <= 1 {
		log.Warn("gm:改变动态技能,参数少于2")

		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	skillIdStr := c.Args[0]
	skillId, err := strconv.ParseInt(skillIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"skillId": skillIdStr,
			}).Warn("gm:技能id,skillId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	levelStr := c.Args[1]
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"level": levelStr,
			}).Warn("gm:技能等级,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	err = skillChangeDynamicLevel(pl, int32(skillId), int32(level))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"skillId": skillId,
				"level":   level,
			}).Warn("gm:改变动态技能等级,错误")
	}

	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"skillId": skillId,
			"level":   level,
		}).Debug("gm:处理改变动态技能等级完成")
	return nil

}

//改变动态技能等级
func skillChangeDynamicLevel(p scene.Player, skillId int32, level int32) error {
	pl := p.(player.Player)
	skillManager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	flag := skillManager.IsValid(skillId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skillgm:无效技能")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}
	flag = skillManager.IfSkillExist(skillId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skillgm:还没有该技能,请先获取")
		playerlogic.SendSystemMessage(pl, lang.SkillNotHas)
		return nil
	}
	flag = skillManager.ChangeDynamicLevel(skillId, level)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skillgm:无效参数")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeSkill.Mask())

	scSkillUpgrade := pbutil.BuildSCSkillUpgrade(skillId)
	pl.SendMsg(scSkillUpgrade)

	return nil
}
