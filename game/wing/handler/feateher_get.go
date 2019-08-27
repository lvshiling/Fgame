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
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEATHER_GET_TYPE), dispatch.HandlerFunc(handleFeatherGet))

}

//处理护体仙羽信息
func handleFeatherGet(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理获取护体仙羽消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = featherGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("Wing:处理获取护体仙羽消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("Wing:处理获取护体仙羽消息完成")
	return nil

}

//获取护体仙羽界面信息逻辑
func featherGet(pl player.Player) (err error) {
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	scFeatherGet := pbutil.BuildSCFeatherGet(wingInfo)
	pl.SendMsg(scFeatherGet)
	return
}
