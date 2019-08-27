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
	processor.Register(codec.MessageType(uipb.MessageType_CS_DEL_EMAIL_BATCH_TYPE), dispatch.HandlerFunc(handlerDelEmailBatch))
}

//删除已读请求
func handlerDelEmailBatch(s session.Session, msg interface{}) (err error) {
	log.Debug("email：处理删除已读邮件请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	err = delEmailBatch(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("email:处理邮件删除已读请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("email:处理邮件删除已读请求完成")
	return
}

//删除已读逻辑
func delEmailBatch(pl player.Player) (err error) {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)

	//删除已读【无or已领取】附件的邮件
	hadReadEmails := emailManager.GetReadEmails()
	var toDelEmailIdArr []int64
	if len(hadReadEmails) > 0 {
		for emailId, _ := range hadReadEmails {
			if emailManager.HasNotOrReceiveAttachment(emailId) {
				emailManager.DelEmail(emailId)
				toDelEmailIdArr = append(toDelEmailIdArr, emailId)
			}
		}

		//已读存在【尚有附件未领取】的邮件
		isNotReceiveAttachment := len(hadReadEmails) > len(toDelEmailIdArr)
		if isNotReceiveAttachment {
			playerlogic.SendSystemMessage(pl, lang.EmailNoReceiveAttachment)
		}
	}
	sceDelHadReadEmail := pbutil.BuildSCDelHadReadEmail(toDelEmailIdArr)
	pl.SendMsg(sceDelHadReadEmail)

	return
}
