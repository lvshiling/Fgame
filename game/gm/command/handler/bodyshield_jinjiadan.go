package handler

import (
	"fgame/fgame/common/lang"
	bodyshieldservice "fgame/fgame/game/bodyshield/bodyshield"
	bodyshieldlogic "fgame/fgame/game/bodyshield/logic"
	"fgame/fgame/game/bodyshield/pbutil"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
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
	command.Register(gmcommandtypes.CommandTypeJinJiaDan, command.CommandHandlerFunc(handleBodyShieldJinJiaDan))

}

func handleBodyShieldJinJiaDan(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	jinJiaDanStr := c.Args[0]
	jinJiaDanLevel, err := strconv.ParseInt(jinJiaDanStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"jinJiaDanLevel": jinJiaDanLevel,
				"error":          err,
			}).Warn("gm:处理设置护体盾食金甲丹等级,jinJiaDanLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := bodyshieldservice.GetBodyShieldService().GetBodyShieldJinJia(int32(jinJiaDanLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"jinJiaDanLevel": jinJiaDanLevel,
				"error":          err,
			}).Warn("gm:处理设置护体盾食金甲丹等级,jinJiaDanLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	manager.GmSetBodyShieldJinJiaDanLevel(int32(jinJiaDanLevel))

	//同步属性
	bodyshieldlogic.BodyShieldPropertyChanged(pl)

	scBodyShieldJJDan := pbutil.BuildSCBodyShieldJJDan(int32(jinJiaDanLevel), 0)
	pl.SendMsg(scBodyShieldJJDan)
	return
}
