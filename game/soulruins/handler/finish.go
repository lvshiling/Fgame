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
	soulruinslogic "fgame/fgame/game/soulruins/logic"
	"fgame/fgame/game/soulruins/pbutil"
	playersoulruins "fgame/fgame/game/soulruins/player"
	"fgame/fgame/game/soulruins/soulruins"
	soulruinstypes "fgame/fgame/game/soulruins/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_FINISH_ALL), dispatch.HandlerFunc(handleSoulRuinsFinish))
}

//处理帝陵遗迹一键完成信息
func handleSoulRuinsFinish(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理帝陵遗迹一键完成消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = soulRuinsFinish(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("soulruins:处理帝陵遗迹一键完成消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soulruins:处理帝陵遗迹一键完成消息完成")
	return nil

}

//帝陵遗迹一键完成界面信息的逻辑
func soulRuinsFinish(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	leftNum := manager.GetSoulRuinsLeftNum()
	costGold := soulruins.GetSoulRuinsService().GetSoulRuinsFinishCostGold()
	needGold := leftNum * costGold
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//判断元宝是否足够
	flag := propertyManager.HasEnoughGold(int64(needGold), true)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("soulruins:元宝不足,无法进阶")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	reasonGoldText := commonlog.GoldLogReasonSoulRuinsFinishAllCost.String()
	flag = propertyManager.CostGold(int64(needGold), true, commonlog.GoldLogReasonSoulRuinsFinishAllCost, reasonGoldText)
	if !flag {
		panic(fmt.Errorf("soulruins: soulRuinsFinishAll CostGold should be ok"))
	}

	//帝陵遗迹一键完成奖励
	chapter, typ, level := manager.GetCurMaxLevel()
	dropItemList, rewData, err := soulruinslogic.GiveSoulRuinsFinishAllReward(pl, chapter, typ, level, leftNum)
	if err != nil {
		return
	}
	//消耗挑战次数
	manager.UseChallengeNum(leftNum)
	//更新挑战信息
	manager.RefreshSoulRuins(chapter, typ, level, soulruinstypes.MaxStar, false)

	//同步属性
	propertylogic.SnapChangedProperty(pl)
	numObj := manager.GetSoulRuinsNum()
	scSoulRuinsChallenge := pbutil.BuildSCSoulRuinsSweep(numObj, rewData, dropItemList)
	pl.SendMsg(scSoulRuinsChallenge)
	return
}
