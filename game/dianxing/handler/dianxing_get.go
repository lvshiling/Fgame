package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dianxing/pbutil"
	playerdianxing "fgame/fgame/game/dianxing/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DIANXING_GET_TYPE), dispatch.HandlerFunc(handleDianXingGet))
}

//处理点星系统信息
func handleDianXingGet(s session.Session, msg interface{}) (err error) {
	log.Debug("dianxing:处理获取点星系统消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = dianXingGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("dianxing:处理获取点星系统消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dianxing:处理获取点星系统消息完成")
	return nil
}

//获取点星系统信息
func dianXingGet(pl player.Player) (err error) {
	dianxingManager := pl.GetPlayerDataManager(playertypes.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	dianXingObject := dianxingManager.GetDianXingObject()
	scDianXingGet := pbutil.BuildSCDianXingGet(dianXingObject)
	pl.SendMsg(scDianXingGet)
	return
}
