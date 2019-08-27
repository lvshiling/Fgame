package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	shenfalogic "fgame/fgame/game/shenfa/logic"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeShenFaUnreal, command.CommandHandlerFunc(handleShenFaUnreal))

}

func handleShenFaUnreal(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	shenFaIdStr := c.Args[0]
	shenFaId, err := strconv.ParseInt(shenFaIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"shenFaId": shenFaId,
				"error":    err,
			}).Warn("gm:处理设置身法幻化,shenFaId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenFaId))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"shenFaId": shenFaId,
				"error":    err,
			}).Warn("gm:处理设置身法幻化,shenFaId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	manager.GmSetShenFaUnreal(int(shenFaId))

	//同步属性
	shenfalogic.ShenfaPropertyChanged(pl)

	scShenfaUnreal := pbutil.BuildSCShenfaUnreal(int32(shenFaId))
	pl.SendMsg(scShenfaUnreal)
	return
}
