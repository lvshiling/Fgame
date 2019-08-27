package handler

import (
	"fgame/fgame/common/lang"
	fabaologic "fgame/fgame/game/fabao/logic"
	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotemplate "fgame/fgame/game/fabao/template"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeFaBaoUnreal, command.CommandHandlerFunc(handleFaBaoUnreal))

}

func handleFaBaoUnreal(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	faBaoIdStr := c.Args[0]
	faBaoId, err := strconv.ParseInt(faBaoIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"faBaoId": faBaoId,
				"error":   err,
			}).Warn("gm:处理设置法宝幻化,faBaoId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoId))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"faBaoId": faBaoId,
				"error":   err,
			}).Warn("gm:处理设置法宝幻化,faBaoId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	manager.GmSetFaBaoUnreal(int(faBaoId))

	//同步属性
	fabaologic.FaBaoPropertyChanged(pl)

	scFaBaoUnreal := pbutil.BuildSCFaBaoUnreal(int32(faBaoId))
	pl.SendMsg(scFaBaoUnreal)
	return
}
