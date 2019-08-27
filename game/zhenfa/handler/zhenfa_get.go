package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/zhenfa/pbutil"
	playerzhenfa "fgame/fgame/game/zhenfa/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ZHENFA_GET_TYPE), dispatch.HandlerFunc(handleZhenFaGet))
}

//处理阵法信息
func handleZhenFaGet(s session.Session, msg interface{}) (err error) {
	log.Debug("zhenfa:处理阵法信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = zhenFaGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("zhenfa:处理阵法信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("zhenfa:处理阵法信息完成")
	return nil
}

//处理阵法信息逻辑
func zhenFaGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	zhenFaMap := manager.GetZhenFaMap()
	scZhenFaGet := pbutil.BuildSCZhenFaGet(zhenFaMap)
	pl.SendMsg(scZhenFaGet)
	return
}
