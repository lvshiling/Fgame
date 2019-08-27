package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/cross/shenmo/pbutil"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_GONGXUN_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerGongXunChanged))
}

//处理玩家功勋值改变
func handlePlayerGongXunChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理玩家功勋值改变")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siPlayerGongXunChanged := msg.(*crosspb.SIPlayerGongXunChanged)
	num := siPlayerGongXunChanged.GetNum()
	err = playerGongXunChanged(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("shenmo:处理玩家功勋值改变,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenmo:处理玩家功勋值改变,完成")
	return nil

}

//玩家功勋值改变
func playerGongXunChanged(pl *player.Player, num int32) (err error) {
	pl.SetShenMoGongXunNum(num)
	isPlayerGongXunChanged := pbutil.BuildISPlayerGongXunChanged()
	pl.SendMsg(isPlayerGongXunChanged)

	return
}
