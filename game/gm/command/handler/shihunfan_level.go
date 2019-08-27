package handler

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	shihunfanlogic "fgame/fgame/game/shihunfan/logic"
	"fgame/fgame/game/shihunfan/pbutil"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeShiHunFanLevel, command.CommandHandlerFunc(handleShiHunFanLevel))

}

func handleShiHunFanLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	shihunfanLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"shihunfanLevel": shihunfanLevel,
				"error":          err,
			}).Warn("gm:处理设置噬魂幡阶别,shihunfanLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFan(int(shihunfanLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"shihunfanLevel": shihunfanLevel,
				"error":          err,
			}).Warn("gm:处理设置噬魂幡阶别,噬魂幡模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	manager.GmSetShiHunFanAdvanced(int(shihunfanLevel))

	//同步属性
	shihunfanlogic.ShiHunFanPropertyChanged(pl)

	scShiHunFanAdvanced := pbutil.BuildSCShiHunFanAdavancedFinshed(manager.GetShiHunFanInfo(), commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scShiHunFanAdvanced)
	return
}
