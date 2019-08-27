package handler

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
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
	command.Register(gmcommandtypes.CommandTypeFaBao, command.CommandHandlerFunc(handleFaBaoLevel))

}

func handleFaBaoLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	faBaoLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"faBaoLevel": faBaoLevel,
				"error":      err,
			}).Warn("gm:处理设置法宝阶别,faBaoLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(faBaoLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"faBaoLevel": faBaoLevel,
				"error":      err,
			}).Warn("gm:处理设置法宝阶别,faBaoLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	manager.GmSetFaBaoAdvanced(int(faBaoLevel))

	//同步属性
	fabaologic.FaBaoPropertyChanged(pl)

	faBaoId := manager.GetFaBaoId()
	scFaBaoAdvanced := pbutil.BuildSCFaBaoAdavancedFinshed(int32(faBaoLevel), faBaoId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scFaBaoAdvanced)
	return
}
