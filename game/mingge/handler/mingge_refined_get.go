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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MINGGE_REFINED_GET_TYPE), dispatch.HandlerFunc(handleMingGeRefinedGet))
}

//处理命盘祭炼信息
func handleMingGeRefinedGet(s session.Session, msg interface{}) (err error) {
	log.Debug("mingge:处理命盘祭炼信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = mingGeRefinedGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("mingge:处理命盘祭炼信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mingge:处理命盘祭炼信息完成")
	return nil
}

//处理命盘祭炼信息逻辑
func mingGeRefinedGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	refinedMap := manager.GetMingGePanRefinedMap()
	scMingGeRefinedGet := pbutil.BuildSCMingGeRefinedGet(refinedMap)
	pl.SendMsg(scMingGeRefinedGet)
	return
}
