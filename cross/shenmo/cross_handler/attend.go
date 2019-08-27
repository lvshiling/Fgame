package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/cross/shenmo/pbutil"
	playerlogic "fgame/fgame/game/player/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shenmo/shenmo"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_SHENMO_ATTEND_TYPE), dispatch.HandlerFunc(handleShenMoAttend))
}

//处理参加神魔战场
func handleShenMoAttend(s session.Session, msg interface{}) (err error) {
	log.Info("shenmo:处理参加神魔战场")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = shenMoAttend(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("shenmo:处理参加神魔战场,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("shenmo:处理参加神魔战场,完成")
	return nil
}

//参加神魔战场
func shenMoAttend(pl *player.Player) (err error) {
	flag := shenmo.GetShenMoService().IsShenMoActivityTime()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Info("shenmo:处理参加神魔战场,活动未开始")
		//活动未开始
		playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
		pl.Close(nil)
		return
	}
	pos, flag := shenmo.GetShenMoService().GetHasLineUp(pl.GetId())
	if flag {
		isArenaMatchResult := pbutil.BuildISShenMoAttend(flag, pos)
		pl.SendMsg(isArenaMatchResult)
		return
	}

	pos, flag = shenmo.GetShenMoService().Attend(pl.GetId())
	isArenaMatchResult := pbutil.BuildISShenMoAttend(flag, pos)
	pl.SendMsg(isArenaMatchResult)
	return
}
