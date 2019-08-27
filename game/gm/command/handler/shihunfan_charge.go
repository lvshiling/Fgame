package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/shihunfan/pbutil"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeShiHunFanCharge, command.CommandHandlerFunc(handleShiHunFanCharge))

}

func handleShiHunFanCharge(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	chargeStr := c.Args[0]
	chargeNum, err := strconv.ParseInt(chargeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"chargeNum": chargeNum,
				"error":     err,
			}).Warn("gm:处理设置护噬魂幡充值数,chargeNum不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	manager.GmSetShiHunFanShiHunFanChargeVal(int32(chargeNum))
	scChargeMsg := pbutil.BuildSCShiHunFanChargeVary(manager.GetShiHunFanInfo().ChargeVal)
	pl.SendMsg(scChargeMsg)
	return
}
