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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_HEGEMON_INFO_TYPE), dispatch.HandlerFunc(handleAllianceHegemon))
}

//处理霸主信息
func handleAllianceHegemon(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理霸主信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceHegemon(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理霸主信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理霸主信息,完成")
	return nil

}

//仙盟
func allianceHegemon(pl player.Player) (err error) {
	alliancelogic.SendAllianceHegemonInfo(pl)
	return
}
