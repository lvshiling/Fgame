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
	command.Register(gmcommandtypes.CommandTypeXueDunPeiYang, command.CommandHandlerFunc(handleXueDunPeiYang))

}

func handleXueDunPeiYang(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	peiYangStr := c.Args[0]
	peiYangLevel, err := strconv.ParseInt(peiYangStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"peiYangLevel": peiYangLevel,
				"error":        err,
			}).Warn("gm:处理设置血盾吞噬等级,peiYangLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := xueduntemplate.GetXueDunTemplateService().GetXueDunPeiYangTemplate(int32(peiYangLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"peiYangLevel": peiYangLevel,
				"error":        err,
			}).Warn("gm:处理设置血盾吞噬等级,peiYangLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	manager.GmSetXueDunPeiYangLevel(int32(peiYangLevel))
	xueDunInfo := manager.GetXueDunInfo()

	//同步属性
	xuedunlogic.XueDunPropertyChanged(pl)

	scXueDunPeiYang := pbutil.BuildSCXueDunPeiYang(xueDunInfo)
	pl.SendMsg(scXueDunPeiYang)
	return
}
