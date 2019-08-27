package handler

import (
	"fgame/fgame/game/email/pbutil"
	playeremail "fgame/fgame/game/email/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeEmailClear, command.CommandHandlerFunc(handleEmailClear))
}

func handleEmailClear(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理发送邮件")
	pl := p.(player.Player)

	err = clearEmail(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理发送邮件,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理发送邮件,完成")
	return
}

func clearEmail(pl player.Player) (err error) {
	emailManager := pl.GetPlayerDataManager(playertypes.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	emailManager.GMClearEmail()

	emailList := emailManager.GetEmails()
	scEmailsGet := pbutil.BuildSCEmailsGet(emailList)
	pl.SendMsg(scEmailsGet)
	return
}
