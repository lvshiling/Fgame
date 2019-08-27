package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	relivelogic "fgame/fgame/game/relive/logic"
	"fgame/fgame/game/relive/pbutil"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_RELIVE_TYPE), dispatch.HandlerFunc(handlePlayerRelive))
}

//玩家复活
func handlePlayerRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("relive:玩家跨服复活")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isPlayerRelive := msg.(*crosspb.ISPlayerRelive)
	reliveTime := isPlayerRelive.GetReliveTime()

	err = playerRelive(tpl, reliveTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"reliveTime": reliveTime,

				"error": err,
			}).Error("relive:玩家跨服复活,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"reliveTime": reliveTime,
		}).Debug("relive:玩家跨服复活")
	return nil
}

//处理设置复活
func playerRelive(pl player.Player, reliveTime int32) (err error) {
	flag := relivelogic.PlayerReliveTimeCost(pl, reliveTime, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("relive:原地复活,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		siPlayerRelive := pbutil.BuildSIPlayerRelive(false)
		pl.SendCrossMsg(siPlayerRelive)
		return
	}
	inventorylogic.SnapInventoryChanged(pl)
	siPlayerRelive := pbutil.BuildSIPlayerRelive(true)
	pl.SendCrossMsg(siPlayerRelive)
	return
}
