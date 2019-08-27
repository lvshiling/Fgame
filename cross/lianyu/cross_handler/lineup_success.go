package handler

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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LIANYU_LINEUP_SUCCESS_TYPE), dispatch.HandlerFunc(handleLianYuLineUpSuccess))
}

//处理无间炼狱排队成功
func handleLianYuLineUpSuccess(s session.Session, msg interface{}) (err error) {
	log.Debug("lianyu:处理无间炼狱排队成功")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = lianYuLineUpSuccess(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("lianyu:处理无间炼狱排队成功,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lianyu:处理无间炼狱排队成功,完成")
	return nil

}

//无间炼狱排队成功
func lianYuLineUpSuccess(pl *player.Player) (err error) {
	return
}
