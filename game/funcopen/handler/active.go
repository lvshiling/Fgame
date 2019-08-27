package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	funcopenlogic "fgame/fgame/game/funcopen/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FUNC_OPEN_MANUAL_ACTIVE_TYPE), dispatch.HandlerFunc(handleManualActive))
}

//处理功能开启手动激活
func handleManualActive(s session.Session, msg interface{}) (err error) {
	log.Debug("funcopen:处理功能开启手动激活")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFuncOpenManualActive := msg.(*uipb.CSFuncOpenManualActive)
	moduleId := csFuncOpenManualActive.GetModuleId()

	err = funcopenlogic.HandleManualActive(tpl, moduleId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"moduleId": moduleId,
				"error":    err,
			}).Error("funcopen:处理功能开启手动激活,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"moduleId": moduleId,
		}).Debug("funcopen:处理功能开启手动激活完成")
	return nil
}
