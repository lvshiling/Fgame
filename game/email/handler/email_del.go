package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/email/pbutil"
	playeremail "fgame/fgame/game/email/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DEL_EMAIL_TYPE), dispatch.HandlerFunc(handlerDelEmail))
}

//删除邮件请求
func handlerDelEmail(s session.Session, msg interface{}) (err error) {
	log.Debug("email：处理删除邮件请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	emailId := msg.(*uipb.CSDelEmail).GetEmailId()
	err = delEmail(tpl, emailId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("email:处理删除邮件请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("email:处理删除邮件请求完成")
	return
}

//删除邮件逻辑
func delEmail(pl player.Player, emailId int64) (err error) {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	_, emailObj := emailManager.GetEmail(emailId)
	//验证参数
	if emailObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"emailId":  emailId,
			}).Warn("email:删除邮件请求参数错误，无效的emailId")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//尚有附件未领取，请先领取附件
	if !emailManager.HasNotOrReceiveAttachment(emailId) {
		playerlogic.SendSystemMessage(pl, lang.EmailNoReceiveAttachment)
		return
	}
	//删除【无or已领取】附件的邮件
	emailManager.DelEmail(emailId)

	scDelEmail := pbutil.BuildSCDelEmail(emailId)
	pl.SendMsg(scDelEmail)

	return
}
