package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/lianyu/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/lianyu/lianyu"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LIANYU_CANCLE_LINEUP_TYPE), dispatch.HandlerFunc(handleLianYuCancleLineUp))
}

//处理无间炼狱取消排队
func handleLianYuCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("lianyu:处理无间炼狱取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = lianYuCancleLineUp(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("lianyu:处理无间炼狱取消排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lianyu:处理无间炼狱取消排队,完成")
	return nil

}

//无间炼狱取消排队
func lianYuCancleLineUp(pl *player.Player) (err error) {
	flag := lianyu.GetLianYuService().CancleLineUp(pl.GetId())
	if !flag {
		return
	}
	isLianYuCancleUp := pbutil.BuildISLianYuCancleUp()
	pl.SendMsg(isLianYuCancleUp)
	return
}
