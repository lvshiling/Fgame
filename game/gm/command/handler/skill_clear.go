package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/skill/pbutil"
	playerskill "fgame/fgame/game/skill/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSkillClear, command.CommandHandlerFunc(handleSkillClear))
}

//将所有技能设置为1级
func handleSkillClear(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空技能")

	err = clearSkillAll(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空技能,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空技能完成")
	return
}

//技能设置为1级
func clearSkillAll(p scene.Player) (err error) {
	pl := p.(player.Player)
	skillManager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	skillMap := pl.GetAllSkills()
	for _, skill := range skillMap {
		skillId := skill.GetSkillId()
		obj := skillManager.GetSkill(skillId)
		if obj == nil {
			continue
		}
		skillManager.GmSetSkillClear(obj)
	}

	scSkillGet := pbutil.BuildSCSkillGet(skillMap)
	pl.SendMsg(scSkillGet)
	return nil
}
