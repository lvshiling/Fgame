package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/major/major"
	"fgame/fgame/game/major/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MAJOR_INVITE_CANCLE_TYPE), dispatch.HandlerFunc(handleMajorInviteCancle))
}

//处理取消双修邀请信息
func handleMajorInviteCancle(s session.Session, msg interface{}) (err error) {
	log.Debug("major:处理取消双修邀请消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = majorInviteCancle(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("major:处理取消双修邀请消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("major:处理取消双修邀请消息完成")
	return nil
}

//处理取消双修邀请信息逻辑
func majorInviteCancle(pl player.Player) (err error) {
	codeResult := major.GetMajorService().CancleMajorInvite(pl)
	scMajorInviteCancle := pbutil.BuildSCMajorInviteCancle(int32(codeResult))
	pl.SendMsg(scMajorInviteCancle)
	return
}
