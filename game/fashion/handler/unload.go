package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/fashion/pbutil"
	playerfashion "fgame/fgame/game/fashion/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_FASHION_UNLOAD_TYPE), dispatch.HandlerFunc(handleFashionUnload))

}

//处理时装卸下信息
func handleFashionUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("fashion:处理时装卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = fashionUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("fashion:处理时装卸下信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fashion:处理时装卸下信息完成")
	return nil

}

//时装卸下逻辑
func fashionUnload(pl player.Player) (err error) {
	fashionManager := pl.GetPlayerDataManager(types.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	fashionWear := fashionManager.GetFashionId()
	bornFashion := fashionManager.GetBornFashion()
	if fashionWear == bornFashion {
		playerlogic.SendSystemMessage(pl, lang.FashionBornNoUnload)
		return
	}

	fashionManager.Unload()
	scFashionUnload := pbutil.BuildSCFashionUnload(fashionWear)
	pl.SendMsg(scFashionUnload)
	return
}
