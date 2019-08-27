package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	lineuplogic "fgame/fgame/game/lineup/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_CHUANGSHI_ENTER_CITY_TYPE), dispatch.HandlerFunc(handleEnterCity))
}

//处理跨服进入城池
func handleEnterCity(s session.Session, msg interface{}) (err error) {
	log.Debug("chuangshi:处理跨服进入城池")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isMsg := msg.(*crosspb.ISChuangShiEnterCity)
	isLineUp := isMsg.GetIsLineUp()
	err = chuangShiEnterCity(tpl, isLineUp)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isLineUp": isLineUp,
				"err":      err,
			}).Error("chuangshi:处理跨服进入城池,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("chuangshi:处理跨服进入城池,完成")
	return nil
}

//进入城池
func chuangShiEnterCity(pl player.Player, isLineUp bool) (err error) {
	if !isLineUp {
		//进入跨服城池
		crosslogic.CrossPlayerDataLogin(pl)
	} else {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isLineUp": isLineUp,
			}).Infoln("chuangshi:处理跨服进入创世城池失败，排队")
		lineuplogic.SendCrossLineup(pl, crosstypes.CrossTypeChuangShi)
	}
	return
}
