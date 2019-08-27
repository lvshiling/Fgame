package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	ringlogic "fgame/fgame/game/ring/logic"
	ringtypes "fgame/fgame/game/ring/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RING_UNLOAD_TYPE), dispatch.HandlerFunc(handleRingUnload))
}

func handleRingUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("ring: 开始处理特戒信息请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRingUnload := msg.(*uipb.CSRingUnload)
	typ := ringtypes.RingType(csRingUnload.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"type":     int32(typ),
			}).Warn("ring: 特戒类型不合法")
		return
	}

	err = ringUnload(tpl, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("ring: 处理特戒信息请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("ring: 处理特戒信息请求消息,成功")

	return
}

func ringUnload(pl player.Player, typ ringtypes.RingType) (err error) {
	ringlogic.RingUnload(pl, typ)
	return
}
