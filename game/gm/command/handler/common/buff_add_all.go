package common

import (
	bufftemplate "fgame/fgame/game/buff/template"
	common "fgame/fgame/game/common/common"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeBuffAddAll, command.CommandHandlerFunc(handleBuffAddAll))
}

func handleBuffAddAll(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理设置buff")

	//TODO 修改物品数量
	err = addBuffAll(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理设置buff,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理设置buff,完成")
	return
}

func addBuffAll(pl scene.Player) (err error) {
	for _, buffTemplate := range bufftemplate.GetBuffTemplateService().GetAllBuffs() {
		scenelogic.AddBuff(pl, int32(buffTemplate.TemplateId()), 0, common.MAX_RATE)
	}
	return
}
