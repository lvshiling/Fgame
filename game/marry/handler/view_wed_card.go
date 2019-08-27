package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_VIEW_WED_CARD_TYPE), dispatch.HandlerFunc(handleMarryViewWedCard))
}

//处理查看喜帖信息
func handleMarryViewWedCard(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理查看喜帖消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMarryViewWedCard := msg.(*uipb.CSMarryViewWedCard)
	cardId := csMarryViewWedCard.GetCardId()
	err = marryViewWedCard(tpl, cardId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"cardId":   cardId,
				"error":    err,
			}).Error("marry:处理查看喜帖消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理查看喜帖消息完成")
	return nil
}

//处理查看喜帖信息逻辑
func marryViewWedCard(pl player.Player, wedCardId int64) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	flag := marry.GetMarryService().WeddingCardIsExist(wedCardId)
	if flag {
		manager.ViewWedCard(wedCardId)
	}
	scMarryViewWedCard := pbuitl.BuildSCMarryViewWedCard(wedCardId)
	pl.SendMsg(scMarryViewWedCard)
	return
}
