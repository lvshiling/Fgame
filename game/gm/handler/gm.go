package handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/gm/pbutil"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	"fgame/fgame/common/codec"
	"fgame/fgame/common/dispatch"

	_ "fgame/fgame/game/gm/command/handler"
	_ "fgame/fgame/game/gm/command/handler/common"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GM_TYPE), dispatch.HandlerFunc(handleGM))
}

func handleGM(s session.Session, msg interface{}) (err error) {
	log.Debug("gm:处理消息")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csGmCommand := msg.(*uipb.CSGMCommand)
	if !global.GetGame().GMOpen() {
		log.Warn("gm:gm命令未开启")
		return
	}
	cmd := csGmCommand.GetCommand()
	err = runCommand(tpl, cmd)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    tpl.GetId(),
				"cmd":   cmd,
				"error": err,
			}).Error("gm:处理错误")
		return err
	}
	log.Debug("gm:处理消息完成")
	return nil
}

func runCommand(pl player.Player, cmd string) (err error) {
	c, err := command.ParseCommand(cmd)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:格式错误")
		err = nil
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	if pl.IsCross() && command.IsCross(c.Type) {
		csGMCommand := pbutil.BuildCSGMCommand(cmd)
		pl.SendCrossMsg(csGMCommand)
		return
	}

	err = command.RunCommand(pl, c)
	if err != nil {
		return
	}
	return
}
