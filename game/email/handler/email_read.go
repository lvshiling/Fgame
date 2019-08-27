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
	processor.Register(codec.MessageType(uipb.MessageType_CS_READ_EMAIL_TYPE), dispatch.HandlerFunc(handlerReadEmail))
}

//读邮件请求
func handlerReadEmail(s session.Session, msg interface{}) error {
	log.Debug("email:处理读邮件请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	csReadEmail := msg.(*uipb.CSReadEmail)
	emailId := csReadEmail.GetEmailId()
	err := readEmail(tpl, emailId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"emailId":  emailId,
				"error":    err,
			}).Error("email:处理读取邮件请求,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"emailId":  emailId,
		}).Debug("email:处理读取邮件请求完成")

	return nil
}

//读邮件请求逻辑
func readEmail(pl player.Player, emailId int64) (err error) {
	//验证参数
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	_, emailObj := emailManager.GetEmail(emailId)
	if emailObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"emailId":  emailId,
			}).Warn("email:无效的emailId,无法读取邮件")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断是否已读
	if emailManager.IsRead(emailId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"emailId":  emailId,
			}).Warn("email:邮件已读，请勿重复读取")
		playerlogic.SendSystemMessage(pl, lang.EmailRepeatedRead)
		return
	}

	//设置邮件已读
	emailManager.ReadEmail(emailId)

	scReadEmail := pbutil.BuildSCReadEmail(emailId)
	pl.SendMsg(scReadEmail)

	return
}
