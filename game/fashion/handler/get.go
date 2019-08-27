package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/fashion/pbutil"
	playerfashion "fgame/fgame/game/fashion/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FASHION_GET_TYPE), dispatch.HandlerFunc(handleFashionGet))
}

//处理获取时装信息
func handleFashionGet(s session.Session, msg interface{}) (err error) {
	log.Debug("fashion:处理获取时装信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = fashionGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("fashion:处理获取时装信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fashion:处理获取时装信息完成")
	return nil
}

//处理时装界面信息逻辑
func fashionGet(pl player.Player) (err error) {
	fashionManager := pl.GetPlayerDataManager(types.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	fashionWear := fashionManager.GetFashionId()
	fashionMap := fashionManager.GetFashionMap()
	trialMap := fashionManager.GetTrialFashionMap()
	scfashionGet := pbutil.BuildSCFashionGet(fashionWear, fashionMap, trialMap)
	pl.SendMsg(scfashionGet)
	return
}
