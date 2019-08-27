package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/bagua/bagua"
	"fgame/fgame/game/bagua/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BAGUA_PAIR_CANCLE_TYPE), dispatch.HandlerFunc(handleBaGuaPairCancle))
}

//处理取消夫妻助战邀请信息
func handleBaGuaPairCancle(s session.Session, msg interface{}) (err error) {
	log.Debug("bagua:处理取消夫妻助战邀请消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = baGuaPairCancle(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("bagua:处理取消夫妻助战邀请消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bagua:处理取消夫妻助战邀请消息完成")
	return nil
}

//处理取消夫妻助战邀请信息逻辑
func baGuaPairCancle(pl player.Player) (err error) {
	codeResult := bagua.GetBaGuaService().CanclePairInvite(pl)
	scBaGuaPairCancle := pbutil.BuildSCBaGuaPairCancle(int32(codeResult))
	pl.SendMsg(scBaGuaPairCancle)
	return
}
