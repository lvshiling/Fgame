package handler

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	marrynpc "fgame/fgame/game/marry/npc/hunche"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeHunChe, command.CommandHandlerFunc(handleHuncheMove))
}

func handleHuncheMove(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理婚车")

	err = huncheMove(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理婚车,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理婚车,完成")
	return
}

func huncheMove(p scene.Player) (err error) {
	pl := p.(player.Player)
	marryConstTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
	moveScene := scene.GetSceneService().GetWorldSceneByMapId(marryConstTemplate.CarMapId)
	moveTemp := marrytemplate.GetMarryTemplateService().GetMarryMoveFirstTeamplate()
	scenelogic.PlayerEnterScene(pl, moveScene, moveTemp.GetPos())
	//刷婚车
	ctx := scene.WithScene(context.Background(), moveScene)
	moveScene.Post(message.NewScheduleMessage(onRefreshMarryCar, ctx, nil, nil))
	return
}

func onRefreshMarryCar(ctx context.Context, result interface{}, err error) error {
	moveScene := scene.SceneInContext(ctx)
	//婚车移动
	n := marrynpc.CreateHunCheNPC(1, scenetypes.OwnerTypeNone, 0, 0, 0, 0, 1, 1)
	//设置场景
	moveScene.AddSceneObject(n)
	return nil
}
