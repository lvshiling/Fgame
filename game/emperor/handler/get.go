package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/emperor/emperor"
	"fgame/fgame/game/emperor/pbutil"
	playeremperor "fgame/fgame/game/emperor/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EMPEROR_GET_TYPE), dispatch.HandlerFunc(handleEmperorGet))
}

//处理龙椅信息
func handleEmperorGet(s session.Session, msg interface{}) (err error) {
	log.Debug("emperor:处理龙椅信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = emperorGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("emperor:处理龙椅信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("emperor:处理龙椅信息完成")
	return nil
}

//处理龙椅界面信息逻辑
func emperorGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerEmperorDataManagerType).(*playeremperor.PlayerEmperorDataManager)
	emperorObj := emperor.GetEmperorService().GetEmperorInfo()
	worshipNum := manager.GetWorshipNum()
	scEmperorGet := pbuitl.BuildSCEmperorGet(emperorObj, worshipNum)
	pl.SendMsg(scEmperorGet)
	return
}
