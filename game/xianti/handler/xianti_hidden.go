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
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANTI_HIDDEN_TYPE), dispatch.HandlerFunc(handleXianTiHidden))
}

//处理仙体隐藏展示信息
func handleXianTiHidden(s session.Session, msg interface{}) (err error) {
	log.Debug("xianti:处理仙体隐藏展示信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csXianTiHidden := msg.(*uipb.CSXiantiHidden)
	hiddenFlag := csXianTiHidden.GetHidden()

	err = xianTiHidden(tpl, hiddenFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
				"error":      err,
			}).Error("xianti:处理仙体隐藏展示信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"hiddenFlag": hiddenFlag,
		}).Debug("xianti:处理仙体隐藏展示信息完成")
	return nil

}

//仙体隐藏展示的逻辑
func xianTiHidden(pl player.Player, hiddenFlag bool) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	manager.Hidden(hiddenFlag)
	scXianTiHidden := pbutil.BuildSCXianTiHidden(hiddenFlag)
	pl.SendMsg(scXianTiHidden)
	return
}
