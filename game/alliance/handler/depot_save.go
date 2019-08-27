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
	processor.Register(codec.MessageType(uipb.MessageType_CS_SAVE_IN_ALLIANCE_DEPOT_TYPE), dispatch.HandlerFunc(handleSaveDepot))
}

//处理仙盟仓库存入
func handleSaveDepot(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟仓库存入")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSSaveInAllianceDepot)
	index := csMsg.GetIndex()
	num := csMsg.GetNum()
	err = saveAllianceDepot(tpl, index, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"error":    err,
			}).Error("alliance:处理仙盟仓库存入,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"index":    index,
		}).Debug("alliance:处理仙盟仓库存入,完成")
	return nil

}

//仙盟仓库存入
func saveAllianceDepot(pl player.Player, index, num int32) (err error) {

	return alliancelogic.HandleSaveAllianceDepot(pl, index, num)
}
