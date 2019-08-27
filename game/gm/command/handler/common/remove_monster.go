package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/template"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeRemoveMonster, command.CommandHandlerFunc(handleRemoveMonster))
}

func handleRemoveMonster(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:移除怪物")
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	monsterIdStr := c.Args[0]
	monsterId, err := strconv.ParseInt(monsterIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"monsterId": monsterIdStr,
				"error":     err,
			}).Warn("gm:移除怪物,怪物id不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	removeNumStr := c.Args[1]
	removeNum, err := strconv.ParseInt(removeNumStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"removeNum": removeNumStr,
				"error":     err,
			}).Warn("gm:移除怪物,移除数量不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = removeMonster(pl, int32(monsterId), int32(removeNum))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:移除怪物,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:移除怪物,完成")
	return
}

func removeMonster(pl scene.Player, monsterId int32, removeNum int32) (err error) {
	tempBiologyTemplate := template.GetTemplateService().Get(int(monsterId), (*gametemplate.BiologyTemplate)(nil))
	if tempBiologyTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:移除怪物,错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	biologyTemplate := tempBiologyTemplate.(*gametemplate.BiologyTemplate)
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:移除怪物,错误")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoInScene)
		return
	}

	npcList := s.GetNPCS(biologyTemplate.GetBiologyScriptType())
	if len(npcList) < int(removeNum) {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:移除怪物,没有足够的数量移除")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	for _, npc := range npcList {
		if removeNum < 1 {
			break
		}
		if npc.GetBiologyTemplate().Id == int(monsterId) {
			s.RemoveSceneObject(npc, false)
		}

		removeNum -= 1
	}
	return
}
