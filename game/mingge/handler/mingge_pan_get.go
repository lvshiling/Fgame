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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MINGGE_PAN_GET_TYPE), dispatch.HandlerFunc(handleMingGePanGet))
}

//处理命格盘信息
func handleMingGePanGet(s session.Session, msg interface{}) (err error) {
	log.Debug("mingge:处理命格盘信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = mingGePanGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("mingge:处理命格盘信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mingge:处理命格盘信息完成")
	return nil
}

//处理命格盘信息逻辑
func mingGePanGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	mingGePanMap := manager.GetMingGePanMap()
	scMingGePanGet := pbutil.BuildSCMingGePanGet(mingGePanMap)
	pl.SendMsg(scMingGePanGet)
	return
}
