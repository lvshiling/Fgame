package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	foelogic "fgame/fgame/game/foe/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOE_FEEDBACK_TYPE), dispatch.HandlerFunc(handleFoeFeedback))
}

//处理报复仇人
func handleFoeFeedback(s session.Session, msg interface{}) (err error) {
	log.Debug("friend:处理报复仇人")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSFoeFeedback)
	foeId := csMsg.GetFoeId()
	args := csMsg.GetArgs()

	err = foeFeedback(tpl, foeId, args)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理报复仇人,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理报复仇人,完成")
	return nil
}

//处理报复仇人
func foeFeedback(pl player.Player, foeId int64, args string) (err error) {
	foePl := player.GetOnlinePlayerManager().GetPlayerById(foeId)
	if foePl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"foeId":    foeId,
			}).Warn("foe:处理仇人推送,用户不在线")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoOnline)
		return
	}

	foelogic.FoeFeedbackNotice(foePl, pl.GetId(), pl.GetName(), args)
	return
}
