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
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WING_HIDDEN_TYPE), dispatch.HandlerFunc(handleWingHidden))
}

//处理战翼隐藏展示信息
func handleWingHidden(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理战翼隐藏展示信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csWingHidden := msg.(*uipb.CSWingHidden)
	hiddenFlag := csWingHidden.GetHidden()

	err = wingHidden(tpl, hiddenFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
				"error":      err,
			}).Error("wing:处理战翼隐藏展示信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"hiddenFlag": hiddenFlag,
		}).Debug("wing:处理战翼隐藏展示信息完成")
	return nil

}

//战翼隐藏展示的逻辑
func wingHidden(pl player.Player, hiddenFlag bool) (err error) {
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingManager.Hidden(hiddenFlag)
	scWingHidden := pbutil.BuildSCWingHidden(hiddenFlag)
	pl.SendMsg(scWingHidden)
	return
}
