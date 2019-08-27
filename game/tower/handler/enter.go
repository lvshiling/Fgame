package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	towerlogic "fgame/fgame/game/tower/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TOWER_ENTER_TYPE), dispatch.HandlerFunc(handleTowerEnter))
}

//处理进入打宝塔
func handleTowerEnter(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理进入打宝塔消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTowerEnter)
	floor := csMsg.GetFloor()

	err = enterTower(tpl, floor)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("tower:处理进入打宝塔消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tower:处理进入打宝塔消息完成")
	return nil

}

func enterTower(pl player.Player, floor int32) (err error) {
	return towerlogic.HandleEnterTower(pl, floor)
}
