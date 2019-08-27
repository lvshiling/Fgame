package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_JIEYI_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerJieYiChanged))
}

//玩家结义变化
func handlePlayerJieYiChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:玩家结义变化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siPlayerJieYiSync := msg.(*crosspb.SIPlayerJieYiSync)
	jieYiId := siPlayerJieYiSync.GetJieYiData().GetJieYiId()
	jieYiName := siPlayerJieYiSync.GetJieYiData().GetJieYiName()
	jieYiRank := siPlayerJieYiSync.GetJieYiData().GetRank()

	err = playerJieYiChanged(tpl, jieYiId, jieYiName, jieYiRank)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("jieyi:玩家结义变化,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("jieyi:玩家结义变化,完成")
	return nil

}

//结义变化
func playerJieYiChanged(pl *player.Player, jieYiId int64, jieYiName string, jieYiRank int32) (err error) {
	pl.SyncJieYi(jieYiId, jieYiName, jieYiRank)
	return
}
