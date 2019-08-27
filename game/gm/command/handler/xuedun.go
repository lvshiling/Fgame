package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	xuedunlogic "fgame/fgame/game/xuedun/logic"
	"fgame/fgame/game/xuedun/pbutil"
	playerxuedun "fgame/fgame/game/xuedun/player"
	xueduntemplate "fgame/fgame/game/xuedun/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXueDunLevel, command.CommandHandlerFunc(handleXueDunLevel))

}

func handleXueDunLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	xueDunLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"xueDunLevel": xueDunLevel,
				"error":       err,
			}).Warn("gm:处理设置血盾阶别,xueDunLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(int32(xueDunLevel), 1)

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"xueDunLevel": xueDunLevel,
				"error":       err,
			}).Warn("gm:处理设置血盾阶别,xueDunLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	manager.GmSetXueDunNumber(int32(xueDunLevel))
	xueDunInfo := manager.GetXueDunInfo()

	//同步属性
	xuedunlogic.XueDunPropertyChanged(pl)

	scXueDunUpgrade := pbutil.BuildSCXueDunUpgrade(true, xueDunInfo)
	pl.SendMsg(scXueDunUpgrade)
	return
}
