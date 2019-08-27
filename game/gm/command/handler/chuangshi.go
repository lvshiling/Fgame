package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/template"
	chuangshiscene "fgame/fgame/game/chuangshi/scene"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeChuangShi, command.CommandHandlerFunc(handleChuangShi))

}

func handleChuangShi(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	cityStr := c.Args[0]
	cityInt, err := strconv.ParseInt(cityStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"city":  cityStr,
				"error": err,
			}).Warn("gm:处理设置进入创世场景,city不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	now := global.GetGame().GetTimeService().Now()
	switch cityInt {
	case 0:
		mapId := int32(5000)
		tempMapTemplate := template.GetTemplateService().Get(int(mapId), (*gametemplate.MapTemplate)(nil))
		if tempMapTemplate == nil {
			return
		}
		mapTemplate := tempMapTemplate.(*gametemplate.MapTemplate)

		s := scene.GetSceneService().GetSceneByMapId(mapId)
		if s == nil {

			endTime := now + 30*int64(common.MINUTE)
			s = chuangshiscene.CreateZhongLiSceneData(int32(mapId), endTime)
		}
		if s == nil {
			return
		}
		scenelogic.PlayerEnterScene(pl, s, mapTemplate.GetBornPos())
		break
	case 1:
		mapId := int32(5001)
		tempMapTemplate := template.GetTemplateService().Get(int(mapId), (*gametemplate.MapTemplate)(nil))
		if tempMapTemplate == nil {
			return
		}
		mapTemplate := tempMapTemplate.(*gametemplate.MapTemplate)

		s := scene.GetSceneService().GetSceneByMapId(mapId)
		if s == nil {

			endTime := now + 30*int64(common.MINUTE)
			s = chuangshiscene.CreateMainSceneData(int32(mapId), endTime)
		}
		if s == nil {
			return
		}
		scenelogic.PlayerEnterScene(pl, s, mapTemplate.GetBornPos())
		break
	case 2:
		mapId := int32(5002)
		tempMapTemplate := template.GetTemplateService().Get(int(mapId), (*gametemplate.MapTemplate)(nil))
		if tempMapTemplate == nil {
			return
		}
		mapTemplate := tempMapTemplate.(*gametemplate.MapTemplate)

		s := scene.GetSceneService().GetSceneByMapId(mapId)
		if s == nil {

			endTime := now + 30*int64(common.MINUTE)
			s = chuangshiscene.CreateFuShuSceneData(int32(mapId), chuangshitypes.RandomChuangShiCamp(), endTime)
		}
		if s == nil {
			return
		}
		scenelogic.PlayerEnterScene(pl, s, mapTemplate.GetBornPos())
		break
	}

	return
}

