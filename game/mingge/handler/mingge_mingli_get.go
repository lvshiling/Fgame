package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MINGGE_MINGLI_GET_TYPE), dispatch.HandlerFunc(handleMingGeMingLiGet))
}

//处理命格命理信息
func handleMingGeMingLiGet(s session.Session, msg interface{}) (err error) {
	log.Debug("mingge:处理命格命理信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = mingGeMingLiGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("mingge:处理命格命理信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mingge:处理命格命理信息完成")
	return nil
}

//处理命格命理信息逻辑
func mingGeMingLiGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	mingLiMap := manager.GetMingLiMap()
	scMingGeMingLiGet := pbutil.BuildSCMingGeMingLiGet(mingLiMap)
	pl.SendMsg(scMingGeMingLiGet)
	return
}
