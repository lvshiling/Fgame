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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ZHENQI_GET_TYPE), dispatch.HandlerFunc(handleZhenQiGet))
}

//处理阵旗信息
func handleZhenQiGet(s session.Session, msg interface{}) (err error) {
	log.Debug("zhenfa:处理阵旗信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = zhenQiGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("zhenfa:处理阵旗信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("zhenfa:处理阵旗信息完成")
	return nil
}

//处理阵旗信息逻辑
func zhenQiGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	zhenQiMap := manager.GetZhenQiMap()
	scZhenQiGet := pbutil.BuildSCZhenQiGet(zhenQiMap)
	pl.SendMsg(scZhenQiGet)
	return
}
