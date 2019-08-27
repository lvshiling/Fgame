package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/scene/scene"
	playerskill "fgame/fgame/game/skill/player"
	"fmt"

	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSkillAdd, command.CommandHandlerFunc(handleSkillAdd))
}

func handleSkillAdd(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理添加技能")
	if len(c.Args) < 1 {
		log.Warn("gm:添加技能,参数少于1")

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
				"skillId": skillId,
			}).Warn("gm:添加技能id,skillTypeId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	//TODO 添加新技能
	err = skillAdd(pl, int32(skillId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"skillId": skillIdStr,
			}).Warn("gm:添加技能,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"skillId": skillId,
		}).Debug("gm:处理添加完成")
	return
}

func skillAdd(p scene.Player, skillId int32) error {
	pl := p.(player.Player)
	skillManager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	flag := skillManager.IsValid(skillId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skill:无效技能")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}
	flag = skillManager.IfSkillExist(skillId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"skillId":  skillId,
		}).Warn("skill:该技能已存在,无需添加")
		playerlogic.SendSystemMessage(pl, lang.SkillHasExist)
		return nil
	}

	flag = skillManager.AddSkill(skillId)
	if !flag {
		panic(fmt.Errorf("skillgm: skillAdd  should be ok"))
	}
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeSkill.Mask())

	return nil
}
