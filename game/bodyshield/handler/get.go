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
	processor.Register(codec.MessageType(uipb.MessageType_CS_BODYSHIELD_GET_TYPE), dispatch.HandlerFunc(handleBodyShieldGet))
}

//处理护体盾信息
func handleBodyShieldGet(s session.Session, msg interface{}) (err error) {
	log.Debug("bodyShield:处理获取护体盾消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = bodyShieldGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("bodyShield:处理获取护体盾消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bodyShield:处理获取护体盾消息完成")
	return nil
}

//获取护体盾信息
func bodyShieldGet(pl player.Player) (err error) {
	bshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
	bShieldInfo := bshieldManager.GetBodyShiedInfo()
	scBodyShieldGet := pbutil.BuildSCBodyShieldGet(bShieldInfo)
	pl.SendMsg(scBodyShieldGet)
	return
}
