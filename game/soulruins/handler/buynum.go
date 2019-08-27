package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/soulruins/pbutil"
	playersoulruins "fgame/fgame/game/soulruins/player"
	"fgame/fgame/game/soulruins/soulruins"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_BUYNUM_TYPE), dispatch.HandlerFunc(handleSoulRuinsBuyNum))
}

//处理帝陵遗迹购买挑战次数信息
func handleSoulRuinsBuyNum(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理获取帝陵遗迹购买挑战次数消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulRuinsBuyNum := msg.(*uipb.CSSoulRuinsBuyNum)
	num := csSoulRuinsBuyNum.GetNum()

	err = soulRuinsBuyNum(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("soulruins:处理获取帝陵遗迹购买挑战次数消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Debug("soulruins:处理获取帝陵遗迹购买挑战次数消息完成")
	return nil

}

//获取帝陵遗迹购买挑战次数界面信息的逻辑
func soulRuinsBuyNum(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("soulruins:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	flag := manager.IfBuyNumReachLimit(num)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("soulruins:今日购买次数已达上限")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsBuyNumReachLimit)
		return
	}

	//需要元宝
	costGold := soulruins.GetSoulRuinsService().GetSoulRuinsBuyChallengeNumCostGold()
	needGold := costGold * num
	//判断元宝是否足够
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag = propertyManager.HasEnoughGold(int64(needGold), true)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("soulruins:元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	reasonGold := commonlog.GoldLogReasonSoulRuinsBuyNum
	reasonGoldText := reasonGold.String()
	flag = propertyManager.CostGold(int64(needGold), true, reasonGold, reasonGoldText)
	if !flag {
		panic(fmt.Errorf("soulruins: soulRuinsBuyNum CostGold should be ok"))
	}
	//同步元宝
	propertylogic.SnapChangedProperty(pl)

	//购买挑战次数
	manager.AddSoulRuinsBuyNum(num)
	numObj := manager.GetSoulRuinsNum()
	scSoulRuinsGet := pbutil.BuildSCSoulRuinsBuyNum(numObj)
	pl.SendMsg(scSoulRuinsGet)
	return
}
