package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	lianyulogic "fgame/fgame/game/lianyu/logic"
	"fgame/fgame/game/lianyu/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LIANYU_CANCLE_LINEUP_TYPE), dispatch.HandlerFunc(handleLianYuCancleLineUp))
}

//处理无间炼狱取消排队
func handleLianYuCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("lianyu:处理无间炼狱取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

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
		}).Debug("lianyu:处理无间炼狱取消排队")
	return nil

}

//处理无间炼狱取消排队
func lianYuCancleLineUp(pl player.Player) (err error) {
	flag := pl.IsLianYuLineUp()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("lingyu:您当前未在排队中")
		playerlogic.SendSystemMessage(pl, lang.LianYuCancleLineUpNoExist)
		return
	}
	lianyulogic.LianYuCancleLineUpSend(pl)
	scLianYuCancleLineUp := pbutil.BuildSCLianYuCancleLineUp()
	pl.SendMsg(scLianYuCancleLineUp)
	return
}
