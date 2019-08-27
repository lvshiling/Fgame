package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/qixue/pbutil"
	playerqixue "fgame/fgame/game/qixue/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QIXUE_GET_TYPE), dispatch.HandlerFunc(handleQiXueGet))
}

//处理泣血枪信息
func handleQiXueGet(s session.Session, msg interface{}) (err error) {
	log.Debug("qixue:处理获取泣血枪消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = qixueGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("qixue:处理获取泣血枪消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("qixue:处理获取泣血枪消息完成")
	return nil
}

//获取泣血枪信息
func qixueGet(pl player.Player) (err error) {
	qixueManager := pl.GetPlayerDataManager(playertypes.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
	qixueInfo := qixueManager.GetQiXueInfo()
	scMsg := pbutil.BuildSCQiXueGet(qixueInfo)
	pl.SendMsg(scMsg)
	return
}
