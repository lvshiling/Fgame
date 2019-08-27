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
	processor.Register(codec.MessageType(uipb.MessageType_CS_WING_GET_TYPE), dispatch.HandlerFunc(handleWingGet))

}

//处理战翼信息
func handleWingGet(s session.Session, msg interface{}) (err error) {
	log.Debug("Wing:处理获取战翼消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = wingGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("Wing:处理获取战翼消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("Wing:处理获取战翼消息完成")
	return nil

}

//获取战翼界面信息逻辑
func wingGet(pl player.Player) (err error) {
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	wingTrialInfo := wingManager.GetWingTrialInfo()
	wingOtherMap := wingManager.GetWingOtherMap()
	scWingGet := pbutil.BuildSCWingGet(wingInfo, wingTrialInfo, wingOtherMap)
	pl.SendMsg(scWingGet)
	return
}
