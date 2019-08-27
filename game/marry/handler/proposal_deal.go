package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"

	//playerfriend "fgame/fgame/game/friend/player"
	playerfriend "fgame/fgame/game/friend/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_PROPOSAL_DEAL_TYPE), dispatch.HandlerFunc(handleMarryProposalDeal))
}

//处理被求婚者决策信息
func handleMarryProposalDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理被求婚者决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMarryProposalDeal := msg.(*uipb.CSMarryProposalDeal)
	dealId := csMarryProposalDeal.GetPlayerId()
	result := csMarryProposalDeal.GetResult()
	err = marryProposalDeal(tpl, marrytypes.MarryResultType(result), dealId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"dealId":   dealId,
				"error":    err,
			}).Error("marry:处理被求婚者决策消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理被求婚者决策消息完成")
	return nil
}

//处理被求婚者决策信息逻辑
func marryProposalDeal(pl player.Player, result marrytypes.MarryResultType, dealId int64) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	if !result.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
			"dealId":   dealId,
		}).Warn("marry:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	ringObj := marry.GetMarryService().GetMarryProposalRing(dealId)
	if ringObj == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
			"dealId":   dealId,
		}).Warn("marry:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if ringObj.PeerId != pl.GetId() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
			"dealId":   dealId,
		}).Warn("marry:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	// manager.RemoveProposaled(dealId)
	ringType := ringObj.Ring //manager.GetProposalRingType(dealId)
	friendManager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	friendObj := friendManager.GetFriend(dealId)
	if friendObj == nil {
		return
	}
	point := friendObj.Point
	marryInfo := manager.GetMarryInfo()
	curRingLevel := marryInfo.RingLevel

	err = marry.GetMarryService().ProposalDeal(pl, dealId, curRingLevel, ringType, point, result)
	return
}
