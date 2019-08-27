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
	"fgame/fgame/game/supremetitle/pbutil"
	playersupremetitle "fgame/fgame/game/supremetitle/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SUPREME_TITLE_GET_TYPE), dispatch.HandlerFunc(handleSupremeTitleGet))
}

//处理获取至尊称号信息
func handleSupremeTitleGet(s session.Session, msg interface{}) (err error) {
	log.Debug("supremetitle:处理获取至尊称号信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = supremeTitleGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("supremetitle:处理获取至尊称号信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("supremetitle:处理获取至尊称号信息完成")
	return nil
}

//处理至尊称号界面信息逻辑
func supremeTitleGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	titleWear := manager.GetTitleWear().GetTitleWear()
	titleMap := manager.GetTitleMap()
	scTitleGet := pbutil.BuildSCSupremeTitleGet(titleWear, titleMap)
	pl.SendMsg(scTitleGet)
	return
}
