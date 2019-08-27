package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/marry/marry"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_DIVORCE_DEAL_TYPE), dispatch.HandlerFunc(handleMarryDivorceDeal))
}

//处理离婚对方决策信息
func handleMarryDivorceDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理离婚对方决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMarryDivorceDeal := msg.(*uipb.CSMarryDivorceDeal)
	result := csMarryDivorceDeal.GetResult()

	err = marryDivorceDeal(tpl, marrytypes.MarryResultType(result))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"error":    err,
			}).Error("marry:处理离婚对方决策消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理离婚对方决策消息完成")
	return nil
}

//处理离婚对方决策信息逻辑
func marryDivorceDeal(pl player.Player, result marrytypes.MarryResultType) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	if !result.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
		}).Warn("marry:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag := manager.IsReceiveDivorce()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
		}).Warn("marry:还没收到离婚决策")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	marryInfo := manager.GetMarryInfo()
	spouseId := marryInfo.SpouseId
	marry.GetMarryService().DivorceDeal(pl, spouseId, result)
	return
}
