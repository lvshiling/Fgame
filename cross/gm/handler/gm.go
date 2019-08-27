package handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/gm/command"
	"fgame/fgame/game/scene/scene"

	gamesession "fgame/fgame/game/session"

	"fgame/fgame/common/codec"
	"fgame/fgame/common/dispatch"

	_ "fgame/fgame/cross/gm/command/handler"
	_ "fgame/fgame/game/gm/command/handler/common"
	playerlogic "fgame/fgame/game/player/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GM_TYPE), dispatch.HandlerFunc(handleGM))
}

func handleGM(s session.Session, msg interface{}) (err error) {
	log.Debug("gm:处理消息")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csGmCommand := msg.(*uipb.CSGMCommand)

	cmd := csGmCommand.GetCommand()
	c, err := command.ParseCommand(cmd)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:格式错误")
		err = nil
		playerlogic.SendSystemMessage(tpl, lang.GMFormatWrong)
		return

	}
	err = command.RunCommand(tpl, c)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    tpl.GetId(),
				"cmd":   cmd,
				"error": err,
			}).Error("gm:处理错误")
		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"cmd":      cmd,
		}).Info("gm:处理消息完成")
	return nil
}
