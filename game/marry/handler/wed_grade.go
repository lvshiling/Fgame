package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/marry/marry"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_WED_GRADE_TYPE), dispatch.HandlerFunc(handleMarryWedGrade))
}

//处理婚礼预定信息
func handleMarryWedGrade(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理婚礼预定处理消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMarryWedGrade := msg.(*uipb.CSMarryWedGrade)
	wedGrade := csMarryWedGrade.GetGrade()
	period := csMarryWedGrade.GetPeriod()
	grade := marrytypes.MarryBanquetSubTypeWed(wedGrade.GetGrade())
	hunCheGrade := marrytypes.MarryBanquetSubTypeHunChe(wedGrade.GetHunCheGrade())
	sugarGrade := marrytypes.MarryBanquetSubTypeSugar(wedGrade.GetSugarGrade())
	if !grade.Valid() || !hunCheGrade.Valid() || !sugarGrade.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"grade":    grade,
			"period":   period,
		}).Warn("marry:处理婚礼预定处理消息,参数无效")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	marryGrade := marrytypes.CreateMarryGrade(grade, hunCheGrade, sugarGrade)
	err = marryWedGrade(tpl, marryGrade, period)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"grade":       grade,
				"hunCheGrade": hunCheGrade,
				"sugarGrade":  sugarGrade,
				"period":      period,
				"error":       err,
			}).Error("marry:处理婚礼预定处理消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理婚礼预定处理消息完成")
	return nil
}

//处理婚礼预定处理信息逻辑
func marryWedGrade(pl player.Player, marryGrade *marrytypes.MarryGrade, period int32) (err error) {
	wedNum := marrytemplate.GetMarryTemplateService().GetMarryConstWedNum()
	grade := marryGrade.Grade
	hunCheGrade := marryGrade.HunCheGrade
	sugarGrade := marryGrade.SugarGrade

	if int32(grade) != int32(hunCheGrade) || int32(grade) != int32(sugarGrade) {
		log.WithFields(log.Fields{
			"playerId":    pl.GetId(),
			"grade":       grade,
			"hunCheGrade": hunCheGrade,
			"sugarGrade":  sugarGrade,
			"period":      period,
		}).Warn("marry:参数无效")
		return
	}
	//参数校验
	if period < 1 || period > wedNum {
		log.WithFields(log.Fields{
			"playerId":    pl.GetId(),
			"grade":       grade,
			"hunCheGrade": hunCheGrade,
			"sugarGrade":  sugarGrade,
			"period":      period,
		}).Warn("marry:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	switch marryInfo.Status {
	case marrytypes.MarryStatusTypeUnmarried: //未婚
		{
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"grade":       grade,
				"hunCheGrade": hunCheGrade,
				"sugarGrade":  sugarGrade,
				"period":      period,
			}).Warn("marry:未结婚,无法预定举办时间")
			playerlogic.SendSystemMessage(pl, lang.MarryReserveNoMarried)
			return
		}
	case marrytypes.MarryStatusTypeEngagement: //订婚
		{
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"grade":       grade,
				"hunCheGrade": hunCheGrade,
				"sugarGrade":  sugarGrade,
				"period":      period,
			}).Warn("marry:已经预定过婚礼举办时间")
			playerlogic.SendSystemMessage(pl, lang.MarryReserveIsExist)
			return
		}
	case marrytypes.MarryStatusTypeMarried: //举办过婚礼
		{
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"grade":       grade,
				"hunCheGrade": hunCheGrade,
				"sugarGrade":  sugarGrade,
				"period":      period,
			}).Warn("marry:已经举办过婚礼了")
			// playerlogic.SendSystemMessage(pl, lang.MarryReserveIsMarried)
			// return
		}
	case marrytypes.MarryStatusTypeDivorce: //离婚
		{
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"grade":       grade,
				"hunCheGrade": hunCheGrade,
				"sugarGrade":  sugarGrade,
				"period":      period,
			}).Warn("marry:处于离婚状态")
			playerlogic.SendSystemMessage(pl, lang.MarryReserveIsDivorced)
			return
		}
	}

	if marryInfo.IsProposal != 1 {
		log.WithFields(log.Fields{
			"playerId":    pl.GetId(),
			"grade":       grade,
			"hunCheGrade": hunCheGrade,
			"sugarGrade":  sugarGrade,
			"period":      period,
		}).Warn("marry:您不是求婚者发起人,无法预定婚礼,赶快向配偶催办婚礼吧")
		playerlogic.SendSystemMessage(pl, lang.MarryPreWedNoProposal)
		return
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	if spl == nil {
		log.WithFields(log.Fields{
			"playerId":    pl.GetId(),
			"grade":       grade,
			"hunCheGrade": hunCheGrade,
			"sugarGrade":  sugarGrade,
			"period":      period,
		}).Warn("marry:您的爱人不在线,无法预定婚期")
		playerlogic.SendSystemMessage(pl, lang.MarryPreWedSpouseNoOnline)
		return
	}

	flag := marry.GetMarryService().MarryPreWedIsExist(pl.GetId())
	if flag {
		log.WithFields(log.Fields{
			"playerId":    pl.GetId(),
			"grade":       grade,
			"hunCheGrade": hunCheGrade,
			"sugarGrade":  sugarGrade,
			"period":      period,
		}).Warn("marry:您的爱人正在思考中,请过一会儿再发送")
		playerlogic.SendSystemMessage(pl, lang.MarryPreWedPeerThinking)
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	costBindGold, costGold, costSilver := marrytemplate.GetMarryTemplateService().GetMarryGradeCost(grade, hunCheGrade, sugarGrade)
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"grade":       grade,
				"hunCheGrade": hunCheGrade,
				"sugarGrade":  sugarGrade,
				"period":      period,
			}).Warn("marry:银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"grade":       grade,
				"hunCheGrade": hunCheGrade,
				"sugarGrade":  sugarGrade,
				"period":      period,
			}).Warn("marry:元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"grade":       grade,
				"hunCheGrade": hunCheGrade,
				"sugarGrade":  sugarGrade,
				"period":      period,
			}).Warn("marry:元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	now := global.GetGame().GetTimeService().Now()
	wedBeginTime, err := marrytemplate.GetMarryTemplateService().GetMarryFisrtWedTime(now)
	if err != nil {
		return err
	}

	marryDurationTime := marrytemplate.GetMarryTemplateService().GetMarryDurationTime()
	marryQingChangTime := marrytemplate.GetMarryTemplateService().GetMarryQingChangTime()
	reserveTime := wedBeginTime + int64(period-1)*(marryDurationTime+marryQingChangTime)
	if now > reserveTime {
		log.WithFields(log.Fields{
			"playerId":    pl.GetId(),
			"grade":       grade,
			"hunCheGrade": hunCheGrade,
			"sugarGrade":  sugarGrade,
			"period":      period,
		}).Warn("marry:预定的举办场次已经过期了")
		playerlogic.SendSystemMessage(pl, lang.MarryReservePeriodIsInvaild)
		return
	}

	wedCodeResult := marrytypes.MarryCodeTypeWeddingSucess
	//距离开始小于1分钟
	diffTime := reserveTime - now
	if diffTime < int64(common.MINUTE) {
		wedCodeResult = marrytypes.MarryCodeTypeWeddingTimeLimit
		scMarryWedGrade := pbuitl.BuildSCMarryWedGrade(int32(wedCodeResult), marryGrade, period)
		pl.SendMsg(scMarryWedGrade)
		return
	}
	err = marry.GetMarryService().MarryPreWedding(pl, period, marryGrade, marryInfo.SpouseId, reserveTime)
	if err != nil {
		return
	}
	goldReason := commonlog.GoldLogReasonMarryWeddingGrade
	silverReason := commonlog.SilverLogReasonMarryWeddingGrade
	goldReasonText := fmt.Sprintf(goldReason.String(), grade, hunCheGrade, sugarGrade)
	silverReasonText := fmt.Sprintf(silverReason.String(), grade, hunCheGrade, sugarGrade)
	flag = propertyManager.Cost(int64(costBindGold), int64(costGold), goldReason, goldReasonText, int64(costSilver), silverReason, silverReasonText)
	if !flag {
		panic(fmt.Errorf("marry: marryWedGrade CostGold should be ok"))
	}
	//同步元宝
	propertylogic.SnapChangedProperty(pl)
	wedCodeResult = marrytypes.MarryCodeTypeWeddingSucess
	scMarryWedGrade := pbuitl.BuildSCMarryWedGrade(int32(wedCodeResult), marryGrade, period)
	pl.SendMsg(scMarryWedGrade)
	return
}
