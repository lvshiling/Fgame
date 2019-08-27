package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_DENSEWAT_SYNC_TYPE), dispatch.HandlerFunc(handleDenseWatSync))
}

//金银密窟同步
func handleDenseWatSync(s session.Session, msg interface{}) (err error) {
	log.Debug("denseWat:金银密窟同步")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)

	isDenseWatSync := msg.(*crosspb.ISDenseWatSync)
	num := isDenseWatSync.GetDensWatData().GetNum()
	endTime := isDenseWatSync.GetDensWatData().GetEndTime()

	err = denseWatSync(tpl, num, endTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"endTime":  endTime,
				"error":    err,
			}).Error("denseWat:金银密窟同步,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
			"endTime":  endTime,
		}).Debug("denseWat:金银密窟同步")
	return nil
}

//处理设置金银密窟线逻辑
func denseWatSync(pl scene.Player, num int32, endTime int64) (err error) {
	pl.SyncDenseWat(num, endTime)
	return
}
