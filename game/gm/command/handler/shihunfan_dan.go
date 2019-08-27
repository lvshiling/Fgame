package handler

import (
	"fgame/fgame/common/lang"
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
	command.Register(gmcommandtypes.CommandTypeShiHunFanDan, command.CommandHandlerFunc(handleShiHunFanDan))

}

func handleShiHunFanDan(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	shihunfanDanStr := c.Args[0]
	shihunfanDanLevel, err := strconv.ParseInt(shihunfanDanStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":                pl.GetId(),
				"shihunfanDanLevel": shihunfanDanLevel,
				"error":             err,
			}).Warn("gm:处理设置护噬魂幡丹等级,shihunfanDanLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanDan(int32(shihunfanDanLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":                pl.GetId(),
				"shihunfanDanLevel": shihunfanDanLevel,
				"error":             err,
			}).Warn("gm:处理设置护噬魂幡丹等级,shihunfanDanLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	manager.GmSetShiHunFanShiHunFanDanLevel(int32(shihunfanDanLevel))

	//同步属性
	shihunfanlogic.ShiHunFanPropertyChanged(pl)

	scBodyShieldJJDan := pbutil.BuildSCShiHunFanEatDan(manager.GetShiHunFanInfo())
	pl.SendMsg(scBodyShieldJJDan)
	return
}
