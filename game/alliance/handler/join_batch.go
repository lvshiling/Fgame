package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_JOIN_APPLAY_BATCH_TYPE), dispatch.HandlerFunc(handleAllianceJoinApplyBatch))
}

//处理批量加入仙盟
func handleAllianceJoinApplyBatch(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理批量加入仙盟")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceJoinBatch(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理批量加入仙盟,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理批量加入仙盟,完成")
	return nil

}

//批量加入仙盟
func allianceJoinBatch(pl player.Player) (err error) {
	return alliancelogic.HandleAllianceJoinBatch(pl)
}
