package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constantypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/unrealboss/pbutil"
	playerunrealboss "fgame/fgame/game/unrealboss/player"
	viplogic "fgame/fgame/game/vip/logic"

	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_UNREAL_BOSS_BUY_PILAO_TYPE), dispatch.HandlerFunc(handlerUnrealBossBuy))
}

//幻境boss购买疲劳
func handlerUnrealBossBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("unrealboss:处理幻境boss购买疲劳请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSUnrealBossBuyPilaoNum)
	buyNum := csMsg.GetBuyNum()

	err = unrealbossBuy(tpl, buyNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("unrealboss:处理幻境boss购买疲劳请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("unrealboss：处理幻境boss购买疲劳请求完成")

	return
}

const (
	limitLevel = 4
)

//幻境boss购买疲劳
func unrealbossBuy(pl player.Player, buyNum int32) (err error) {
	if pl.GetVip() < limitLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"vip":      pl.GetVip(),
			}).Warn("unrealboss:幻境boss购买次数，VIP等级不足")
		playerlogic.SendSystemMessage(pl, lang.VipLevelToLow)
		return
	}

	// 购买次数
	maxBuyTimes := viplogic.GetUnrealBossMaxBuyTimes(pl)
	unrealManager := pl.GetPlayerDataManager(playertypes.PlayerUnrealBossDataManagerType).(*playerunrealboss.PlayerUnrealBossDataManager)
	curBuyTimes := unrealManager.GetPilaoBuyTimes()
	if curBuyTimes+buyNum > maxBuyTimes {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"buyNum":      buyNum,
				"curBuyTimes": curBuyTimes,
			}).Warn("unrealboss:幻境boss购买次数，已达最大购买次数")
		playerlogic.SendSystemMessage(pl, lang.UnrealBossBuyNumReachLimit)
		return
	}

	//元宝足够
	price := constant.GetConstantService().GetConstant(constantypes.ConstantTypePiLaoPrice)
	needGold := int64(price * buyNum)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needGold, false) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"buyNum":      buyNum,
				"curBuyTimes": curBuyTimes,
				"needGold":    needGold,
			}).Warn("unrealboss:幻境boss购买次数，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	useReason := commonlog.GoldLogReasonUnrealBossBuyPilao
	flag := propertyManager.CostGold(needGold, false, useReason, useReason.String())
	if !flag {
		panic("unrealboss:购买疲劳，消耗元宝应该成功")
	}

	unrealManager.BuyPilaoNum(buyNum)
	propertylogic.SnapChangedProperty(pl)

	curPilao := unrealManager.GetCurPilaoNum()
	scMsg := pbutil.BuildSCUnrealBossBuyPilaoNum(curPilao)
	pl.SendMsg(scMsg)
	return
}
