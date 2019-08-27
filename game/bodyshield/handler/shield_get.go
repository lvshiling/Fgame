package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/bodyshield/pbutil"
	playerbshield "fgame/fgame/game/bodyshield/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHIELD_GET_TYPE), dispatch.HandlerFunc(handleShieldGet))
}

//处理神盾尖刺信息
func handleShieldGet(s session.Session, msg interface{}) (err error) {
	log.Debug("bodyShield:处理获取神盾尖刺消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = shieldGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("bodyShield:处理获取神盾尖刺消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bodyShield:处理获取神盾尖刺消息完成")
	return nil
}

//获取神盾尖刺信息
func shieldGet(pl player.Player) (err error) {
	bshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
	bShieldInfo := bshieldManager.GetBodyShiedInfo()
	scShieldGet := pbutil.BuildSCShieldGet(bShieldInfo)
	pl.SendMsg(scShieldGet)
	return
}
