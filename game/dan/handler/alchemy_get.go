package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dan/pbutil"
	playerdan "fgame/fgame/game/dan/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DAN_ALCHEMYGET_TYPE), dispatch.HandlerFunc(handleAlchemyGet))
}

//处理炼丹信息
func handleAlchemyGet(s session.Session, msg interface{}) (err error) {
	log.Debug("dan:处理获取炼丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = alchemyGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("dan:处理获取炼丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dan:处理获取炼丹信息完成")
	return nil
}

//处理炼丹界面信息逻辑
func alchemyGet(pl player.Player) (err error) {
	danManager := pl.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	achemyList := danManager.GetAlchemyInfo()
	scAchemyGet := pbuitl.BuildSCAlchemyGet(achemyList)
	pl.SendMsg(scAchemyGet)
	return
}
