package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENFA_HIDDEN_TYPE), dispatch.HandlerFunc(handleShenfaHidden))
}

//处理身法隐藏展示信息
func handleShenfaHidden(s session.Session, msg interface{}) (err error) {
	log.Debug("shenfa:处理身法隐藏展示信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csShenfaHidden := msg.(*uipb.CSShenfaHidden)
	hiddenFlag := csShenfaHidden.GetHidden()

	err = shenfaHidden(tpl, hiddenFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
				"error":      err,
			}).Error("shenfa:处理身法隐藏展示信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"hiddenFlag": hiddenFlag,
		}).Debug("shenfa:处理身法隐藏展示信息完成")
	return nil

}

//身法隐藏展示的逻辑
func shenfaHidden(pl player.Player, hiddenFlag bool) (err error) {
	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)

	shenfaManager.Hidden(hiddenFlag)

	scShenfaHidden := pbutil.BuildSCShenfaHidden(hiddenFlag)
	pl.SendMsg(scShenfaHidden)
	return
}
