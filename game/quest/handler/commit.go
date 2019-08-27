package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	questlogic "fgame/fgame/game/quest/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_COMMIT_TYPE), dispatch.HandlerFunc(handleQuestCommit))
}

//处理任务交付
func handleQuestCommit(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理任务交付")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csQuestCommit := msg.(*uipb.CSQuestCommit)
	questId := csQuestCommit.GetQuestId()
	double := csQuestCommit.GetDouble()

	if questId <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务交付,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = questCommit(tpl, questId, double)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"double":   double,
				"error":    err,
			}).Error("quest:处理任务交付,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"questId":  questId,
			"double":   double,
		}).Debug("quest:处理任务交付,完成")
	return nil
}

//交付
func questCommit(pl player.Player, questId int32, double bool) (err error) {
	return questlogic.CommitQuest(pl, questId, double)
}
