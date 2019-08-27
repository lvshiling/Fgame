package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGYU_HIDDEN_TYPE), dispatch.HandlerFunc(handleLingyuHidden))
}

//处理领域隐藏展示信息
func handleLingyuHidden(s session.Session, msg interface{}) (err error) {
	log.Debug("lingyu:处理领域隐藏展示信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csLingyuHidden := msg.(*uipb.CSLingyuHidden)
	hiddenFlag := csLingyuHidden.GetHidden()

	err = lingyuHidden(tpl, hiddenFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
				"error":      err,
			}).Error("lingyu:处理领域隐藏展示信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"hiddenFlag": hiddenFlag,
		}).Debug("lingyu:处理领域隐藏展示信息完成")
	return nil

}

//领域隐藏展示的逻辑
func lingyuHidden(pl player.Player, hiddenFlag bool) (err error) {
	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuManager.Hidden(hiddenFlag)

	scLingyuHidden := pbutil.BuildSCLingyuHidden(hiddenFlag)
	pl.SendMsg(scLingyuHidden)
	return
}
