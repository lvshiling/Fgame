package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenmologic "fgame/fgame/game/shenmo/logic"
	"fgame/fgame/game/shenmo/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENMO_CANCLE_LINEUP_TYPE), dispatch.HandlerFunc(handleShenMoCancleLineUp))
}

//处理神魔战场取消排队
func handleShenMoCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shenMoCancleLineUp(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("shenmo:处理神魔战场取消排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenmo:处理神魔战场取消排队")
	return nil

}

//处理神魔战场取消排队
func shenMoCancleLineUp(pl player.Player) (err error) {
	flag := pl.IsShenMoLineUp()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("shenmo:您当前未在排队中")
		playerlogic.SendSystemMessage(pl, lang.ShenMoCancleLineUpNoExist)
		return
	}
	shenmologic.ShenMoCancleLineUpSend(pl)
	scShenMoCancleLineUp := pbutil.BuildSCShenMoCancleLineUp()
	pl.SendMsg(scShenMoCancleLineUp)
	return
}
