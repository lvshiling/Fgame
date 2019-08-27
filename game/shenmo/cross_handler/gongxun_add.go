package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenmologic "fgame/fgame/game/shenmo/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_GONGXUN_ADD_TYPE), dispatch.HandlerFunc(handlePlayerGongXunAdd))
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_GONGXUN_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerGongXunChanged))
}

func handlePlayerGongXunChanged(s session.Session, msg interface{}) (err error) {
	return
}

//处理玩家功勋值改变
func handlePlayerGongXunAdd(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理玩家功勋值改变")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isPlayerGongXunAdd := msg.(*crosspb.ISPlayerGongXunAdd)
	addNum := isPlayerGongXunAdd.GetNum()
	err = playerGongXunAdd(tpl, addNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"addNum":   addNum,
				"err":      err,
			}).Error("shenmo:处理玩家功勋值改变,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"addNum":   addNum,
		}).Debug("shenmo:处理玩家功勋值改变,完成")
	return nil

}

//玩家功勋值改变
func playerGongXunAdd(pl player.Player, addGongXun int32) (err error) {
	shenmologic.AddGongXun(pl, addGongXun)

	// siPlayerGongXunNumChanged := pbutil.BuildSIPlayerGongXunNumChanged(totalNum)
	// pl.SendCrossMsg(siPlayerGongXunNumChanged)
	return
}
