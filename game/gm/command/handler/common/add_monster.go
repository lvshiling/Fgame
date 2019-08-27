package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/template"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeAddMonster, command.CommandHandlerFunc(handleAddMonster))
}

func handleAddMonster(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:添加怪物")
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
			}).Warn("gm:添加怪物,怪物id不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = addMonster(pl, int32(monsterId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:添加怪物,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:添加怪物,完成")
	return
}

func addMonster(pl scene.Player, monsterId int32) (err error) {
	tempBiologyTemplate := template.GetTemplateService().Get(int(monsterId), (*gametemplate.BiologyTemplate)(nil))
	if tempBiologyTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:添加怪物,错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	biologyTemplate := tempBiologyTemplate.(*gametemplate.BiologyTemplate)
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:添加怪物,错误")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoInScene)
		return
	}
	n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), 0, biologyTemplate, pl.GetPosition(), 90, 0)
	if n == nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"biologyType": biologyTemplate.GetBiologyType().String(),
			}).Warn("gm:添加怪物,npc不存在")
		return
	}
	s.AddSceneObject(n)
	return
}
