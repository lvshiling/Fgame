package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	viplogic "fgame/fgame/game/vip/logic"
	playervip "fgame/fgame/game/vip/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeVipGift, command.CommandHandlerFunc(handleVipGiftReset))
}

func handleVipGiftReset(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理vip礼包次数重置")

	err = resetVipGift(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理vip礼包次数重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理vip礼包次数重置完成")
	return
}

func resetVipGift(p scene.Player) (err error) {
	pl := p.(player.Player)
	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	vipManager.GMResetBuyGift()

	viplogic.VipInfoNotice(pl)
	return
}
