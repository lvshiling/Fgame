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
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TITLE_GET_TYPE), dispatch.HandlerFunc(handleTitleGet))
}

//处理获取称号信息
func handleTitleGet(s session.Session, msg interface{}) (err error) {
	log.Debug("title:处理获取称号信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = titleGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("title:处理获取称号信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("title:处理获取称号信息完成")
	return nil
}

//处理称号界面信息逻辑
func titleGet(pl player.Player) (err error) {
	titleManager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	titleWear := titleManager.GetTitleWear().TitleWear
	titleMap := titleManager.GetTitleIdMap()
	scTitleGet := pbutil.BuildSCTitleGet(pl, titleWear, titleMap)
	pl.SendMsg(scTitleGet)
	return
}
