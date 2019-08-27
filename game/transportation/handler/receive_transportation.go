package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	transpotationlogic "fgame/fgame/game/transportation/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RECEIVE_TRANSPORT_REW_TYPE), dispatch.HandlerFunc(handlerReceiveTransportationRew))
}

//领取押镖奖励
func handlerReceiveTransportationRew(s session.Session, msg interface{}) (err error) {
	log.Debug("transport:处理领取押镖奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = receiveTransportationRew(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("transport:处理领取押镖奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("transport:处理领取押镖奖励请求完成")

	return
}

func receiveTransportationRew(pl player.Player) (err error) {

	return transpotationlogic.HandleReceiveTransportationRew(pl)
}
