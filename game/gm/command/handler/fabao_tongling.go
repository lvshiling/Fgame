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
	command.Register(gmcommandtypes.CommandTypeFaBaoTongLing, command.CommandHandlerFunc(handleFaBaoTongLing))

}

func handleFaBaoTongLing(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	tongLingStr := c.Args[0]
	tongLingLevel, err := strconv.ParseInt(tongLingStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":            pl.GetId(),
				"tongLingLevel": tongLingLevel,
				"error":         err,
			}).Warn("gm:处理设置法宝食幻化丹等级,tongLingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := fabaotemplate.GetFaBaoTemplateService().GetFaBaoTongLingTemplate(int32(tongLingLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":            pl.GetId(),
				"tongLingLevel": tongLingLevel,
				"error":         err,
			}).Warn("gm:处理设置法宝食幻化丹等级,tongLingLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	manager.GmSetFaBaoTongLing(int32(tongLingLevel))

	faBaoInfo := manager.GetFaBaoInfo()

	//同步属性
	fabaologic.FaBaoPropertyChanged(pl)

	scFaBaoTongLing := pbutil.BuildSCFaBaoTongLing(faBaoInfo.GetTongLingLevel(), faBaoInfo.GetTongLingNum(), faBaoInfo.GetTongLingPro())
	pl.SendMsg(scFaBaoTongLing)
	return
}
