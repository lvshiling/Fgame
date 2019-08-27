package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/marry/marry"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_WED_GRADE_SPOUSE_DEAL_TYPE), dispatch.HandlerFunc(handleMarryPreWedDeal))
}

//处理婚期预定决策信息
func handleMarryPreWedDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理婚期预定决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMarryWedGradeSpouseDeal := msg.(*uipb.CSMarryWedGradeSpouseDeal)
	result := csMarryWedGradeSpouseDeal.GetResult()
	err = marryPreWed(tpl, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"error":    err,
			}).Error("marry:处理婚期预定决策消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理婚期预定决策消息完成")
	return nil
}

//处理婚期预定决策信息逻辑
func marryPreWed(pl player.Player, result bool) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	flag := manager.GetMarryPreWedFlag()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
		}).Warn("marry:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	defer manager.MarryPreWedFlag(false)
	playerMarryInfo := manager.GetMarryInfo()
	isFirst := false
	if !playerMarryInfo.HasHunLi() {
		isFirst = true
	}

	err = marry.GetMarryService().MarryPreWedDeal(pl, result, isFirst)
	if err != nil {
		return
	}

	return
}
