package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shenmo/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_SHENMO_KILLNUM_CHANGED_TYPE), dispatch.HandlerFunc(handleShenMoKillNumChanged))
}

//处理跨服神魔战场击杀人数变更
func handleShenMoKillNumChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理跨服神魔战场击杀人数变更")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isShenMoKillNumChanged := msg.(*crosspb.ISShenMoKillNumChanged)
	killNum := isShenMoKillNumChanged.GetKillNum()
	err = shenMoKillNumChanged(tpl, killNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"killNum":  killNum,
				"err":      err,
			}).Error("shenmo:处理跨服神魔战场击杀人数变更,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"killNum":  killNum,
		}).Debug("shenmo:处理跨服神魔战场击杀人数变更,完成")
	return nil

}

//神魔战场击杀人数变更
func shenMoKillNumChanged(pl player.Player, killNum int32) (err error) {
	pl.SetShenMoKillNum(killNum)
	siShenMoKillNumChanged := pbutil.BuildSIShenMoKillNumChanged()
	pl.SendCrossMsg(siShenMoKillNumChanged)
	return
}
