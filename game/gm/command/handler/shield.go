package handler

import (
	"fgame/fgame/common/lang"
	bodyshieldservice "fgame/fgame/game/bodyshield/bodyshield"
	bodyshieldlogic "fgame/fgame/game/bodyshield/logic"
	"fgame/fgame/game/bodyshield/pbutil"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
	commontypes "fgame/fgame/game/common/types"
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
	command.Register(gmcommandtypes.CommandTypeShieldLevel, command.CommandHandlerFunc(handleShieldLevel))

}

func handleShieldLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	shield, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"shield": shield,
				"error":  err,
			}).Warn("gm:处理设置神盾尖刺阶别,shield不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := bodyshieldservice.GetBodyShieldService().GetShield(int32(shield))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"shield": shield,
				"error":  err,
			}).Warn("gm:处理设置神盾尖刺阶别,shield模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	manager.GmSetShieldAdvanced(int32(shield))

	//同步属性
	bodyshieldlogic.ShieldPropertyChanged(pl)
	scBodyShieldAdvanced := pbutil.BuildSCShieldAdvanced(int32(shield), 0, commontypes.AdvancedTypeJinJieDan, false)
	pl.SendMsg(scBodyShieldAdvanced)
	return
}
