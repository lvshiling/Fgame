package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_UP_LEV_TYPE), dispatch.HandlerFunc(handlePlayerTokenUpLev))
}

func handlePlayerTokenUpLev(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理玩家信物升级请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csJieYiTokenUpLev := msg.(*uipb.CSJieYiTokenUpLev)
	autoFlag := csJieYiTokenUpLev.GetAutoFlag()

	err = playerTokenUpLev(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理玩家信物升级请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理玩家信物升级请求消息,成功")

	return
}

func playerTokenUpLev(pl player.Player, autoFlag bool) (err error) {
	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if plObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)

	if !jieYiManager.IsTokenActivite() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 信物未激活")
		playerlogic.SendSystemMessage(pl, lang.JieYiTokenNotActivite)
		return
	}

	playerJieYiObj := jieYiManager.GetPlayerJieYiObj()
	token := playerJieYiObj.GetTokenType()
	level := playerJieYiObj.GetTokenLevel() + 1

	tokenTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiTokenLevelTemplate(token, level)
	if tokenTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"tokenType":  int32(token),
				"tokenLevel": level,
			}).Warn("jieyi: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.JieYiTemplateNotExist)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := int64(0)
	//进阶需要消耗的银两
	costSilver := int64(0)
	//进阶需要消耗的绑元
	costBindGold := int64(0)
	shopIdMap := make(map[int32]int32)
	curNum := inventoryManager.NumOfItems(tokenTemp.UseItemId)
	needNum := tokenTemp.UseItemCount - curNum
	useItem := tokenTemp.UseItemId
	useNum := int32(0)
	if needNum > 0 {
		if autoFlag == false {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"tokenType":  int32(token),
					"tokenLevel": level,
				}).Warn("jieyi: 所需物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		useNum = curNum
		if !shop.GetShopService().ShopIsSellItem(useItem) {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("jieyi:商铺没有该道具,无法自动购买")
			playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
			return
		}

		isEnoughBuyTimes, shopIdMap := shoplogic.MaxBuyTimesForPlayer(pl, useItem, needNum)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("jieyi:购买物品失败,自动进阶已停止")
			playerlogic.SendSystemMessage(pl, lang.ShopAdvancedAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		costGold += shopNeedGold
		costBindGold += shopNeedBindGold
		costSilver += shopNeedSilver
	} else {
		useNum = tokenTemp.UseItemCount
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(costSilver)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("jieyi:银两不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(costGold, false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("jieyi:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(needBindGold, true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("jieyi:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonJieYiTokenUpLev
	silverUseReason := commonlog.SilverLogReasonJieYiTokenUpLev
	flag := propertyManager.Cost(costBindGold, costGold, goldUseReason, goldUseReason.String(), costSilver, silverUseReason, silverUseReason.String())
	if !flag {
		panic(fmt.Errorf("jieyi: 信物升级自动购买应该成功"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	// if curNum > tokenTemp.UseItemCount {
	// 	curNum = tokenTemp.UseItemCount
	// }

	if useNum > 0 {
		reason := commonlog.InventoryLogReasonJieYiTokenLevelChangeUse
		reasonText := fmt.Sprintf(reason.String(), token.String(), level)
		flag = inventoryManager.UseItem(useItem, useNum, reason, reasonText)
		if !flag {
			panic("jieyi: 消耗物品应该成功")
		}
	}

	// 物品改变推送
	inventorylogic.SnapInventoryChanged(pl)

	// 升级判断
	pro, randBless, success := jieyilogic.JieYiTokenUpLev(playerJieYiObj.GetTokenNum(), playerJieYiObj.GetTokenPro(), tokenTemp)
	if !success {
		level--
	}

	// 同步数据
	jieYiManager.TokenUpLevel(pro, success)
	if success {
		//同步等级
		jieyi.GetJieYiService().TokenChangeLevel(pl.GetId(), level)

		jieYi := plObj.GetJieYi()
		scJieBrotherInfoOnChange := pbutil.BuildSCJieBrotherInfoOnChange(plObj)
		jieyilogic.BroadcastJieYi(jieYi, scJieBrotherInfoOnChange)

		jieyilogic.JieYiMemberChanged(jieYi)
	}

	scMsg := pbutil.BuildSCJieYiTokenUpLev(int32(token), level, randBless, pro, success)
	pl.SendMsg(scMsg)

	return
}
