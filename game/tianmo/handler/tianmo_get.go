package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tianmo/pbutil"
	playertianmo "fgame/fgame/game/tianmo/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TIANMOTI_GET_TYPE), dispatch.HandlerFunc(handleTianMoGet))
}

//处理天魔信息
func handleTianMoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("tianMo:处理获取天魔消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = tianMoGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("tianMo:处理获取天魔消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tianMo:处理获取天魔消息完成")
	return nil
}

//获取天魔信息
func tianMoGet(pl player.Player) (err error) {
	tianMoManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianMoInfo := tianMoManager.GetTianMoInfo()
	scTianMoGet := pbutil.BuildSCTianMoGet(tianMoInfo)
	pl.SendMsg(scTianMoGet)
	return
}
