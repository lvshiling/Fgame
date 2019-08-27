package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arenapvp/arenapvp"
	"fgame/fgame/game/arenapvp/pbutil"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENAPVP_GUESS_TYPE), dispatch.HandlerFunc(handleArenapvpGuess))
}

//处理竞猜信息
func handleArenapvpGuess(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理竞猜信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSArenapvpGuess)
	guessId := csMsg.GetGuessId()

	err = arenapvpGuess(tpl, guessId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理竞猜信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理竞猜信息,完成")
	return nil
}

func arenapvpGuess(pl player.Player, guessId int64) (err error) {
	guessData := arenapvp.GetArenapvpService().GetArenapvpGuessData()
	if guessData == nil {
		log.WithFields(
			log.Fields{
				"playerId":          pl.GetId(),
				"guessData.PvpType": guessData.PvpType,
			}).Warn("arenapvp:参与竞猜,竞猜不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	pvpTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpTemplate(guessData.PvpType)
	if pvpTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"guessType": guessData.PvpType,
			}).Warn("arenapvp:参与竞猜,竞猜类型错误")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if guessData.GetWinnerId() != 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"winnerId": guessData.GetWinnerId(),
			}).Warn("arenapvp:参与竞猜,对战已经结束")
		playerlogic.SendSystemMessage(pl, lang.ArenapvpGuessFail)
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if !pvpTemp.IfCanGuess(now) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arenapvp:参与竞猜,竞猜时间已过")
		playerlogic.SendSystemMessage(pl, lang.ArenapvpGuessFail)
		return
	}

	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	lastGuessObj := arenapvpManager.GetLastGuessLog()
	if lastGuessObj != nil && lastGuessObj.IfAttendGuess(guessData.RaceNumber, guessData.PvpType) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"guessId":  lastGuessObj.GetGuessId(),
			}).Warn("arenapvp:参与竞猜,已经参与过竞猜")
		playerlogic.SendSystemMessage(pl, lang.ArenapvpGuessFail)
		return
	}

	needBindGold := int64(pvpTemp.JingchaiUseBindgold)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needBindGold, true) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"needBindGold": needBindGold,
			}).Warn("arenapvp:参与竞猜,元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	useReason := commonlog.GoldLogReasonArenapvpAttendGuess
	useReasonText := fmt.Sprintf(useReason.String(), pvpTemp.GetArenapvpType().String())
	flag := propertyManager.CostGold(needBindGold, true, useReason, useReasonText)
	if !flag {
		panic("arenapvp:竞猜消耗元宝应该成功")
	}
	propertylogic.SnapChangedProperty(pl)

	flag = arenapvpManager.AddGuessLog(guessData.RaceNumber, guessData.PvpType, guessId)
	if !flag {
		panic(fmt.Errorf("竞猜失败，已参与竞猜。guessId:%d，guessType:%v", guessId, guessData.PvpType))
	}
	arenapvp.GetArenapvpService().AttendGuess(pl.GetId(), guessData.RaceNumber, guessId, guessData.PvpType)

	scMsg := pbutil.BuildSCArenapvpGuess(guessId)
	pl.SendMsg(scMsg)
	return
}
