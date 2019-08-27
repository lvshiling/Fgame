package handler

import (
	"fgame/fgame/game/charge/pbutil"
	playercharge "fgame/fgame/game/charge/player"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeChargeGift, command.CommandHandlerFunc(handleChargeGiftReset))
}

func handleChargeGiftReset(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理档次首充次数重置")

	err = resetChargeGift(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理档次首充次数重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理档次首充次数重置完成")
	return
}

func resetChargeGift(p scene.Player) (err error) {
	pl := p.(player.Player)
	chargeManager := pl.GetPlayerDataManager(types.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	now := global.GetGame().GetTimeService().Now()
	chargeManager.ResetFirstChargeRecord(now)
	recordMap := chargeManager.GetFirstChargeRecord()

	scMsg := pbutil.BuildSCFirstChargeRecordNotice(recordMap)
	pl.SendMsg(scMsg)
	return
}
