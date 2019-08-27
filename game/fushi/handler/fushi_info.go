package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/fushi/pbutil"
	playerfushi "fgame/fgame/game/fushi/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FUSHI_INFO_TYPE), dispatch.HandlerFunc(handleGetFushiInfo))
}

func handleGetFushiInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("处理获取八卦符石请求消息,开始")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = getFushiInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("fushi:处理获取八卦符石请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fushi:处理获取八卦符石请求消息,成功")

	return
}

func getFushiInfo(pl player.Player) (err error) {
	fushiManager := pl.GetPlayerDataManager(playertypes.PlayerFuShiDataManagerType).(*playerfushi.PlayerFuShiDataManager)
	fushiInfo := fushiManager.GetFushiInfo()
	scMsg := pbutil.BuildSCFuShiInfo(fushiInfo)
	pl.SendMsg(scMsg)
	return
}
