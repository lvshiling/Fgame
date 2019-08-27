package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	"fgame/fgame/game/marry/marry"
	marryscene "fgame/fgame/game/marry/scene"
	marrytemplate "fgame/fgame/game/marry/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMarry, command.CommandHandlerFunc(handleMarry))
}

func handleMarry(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理进入结婚场景")

	err = enterMarry(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),

				"error": err,
			}).Warn("gm:处理跳转结婚场景,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理跳转结婚场景,完成")
	return
}

func enterMarry(p scene.Player) (err error) {
	pl := p.(player.Player)
	plScene := pl.GetScene()
	sd := marry.GetMarryService().GetMarrySceneData()
	switch sd.Status {
	case marryscene.MarrySceneStatusTypeInit,
		marryscene.MarrySceneStatusBanquet:
		{
			marryTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
			marryMapTemplate := marryTemplate.GetMarryMapTemplate()
			marryScene := marry.GetMarryService().GetScene()

			if plScene == marryScene {
				return
			}
			scenelogic.PlayerEnterScene(pl, marryScene, marryMapTemplate.GetBornPos())
			break
		}
	case marryscene.MarrySceneStatusCruise:
		{
			marryConstTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
			moveScene := scene.GetSceneService().GetWorldSceneByMapId(marryConstTemplate.CarMapId)
			marryMoveTemplate := marrytemplate.GetMarryTemplateService().GetMarryMoveTeamplate(1)
			if plScene == moveScene {
				return
			}
			scenelogic.PlayerEnterScene(pl, moveScene, marryMoveTemplate.GetPos())
			break
		}
	}
	return
}
