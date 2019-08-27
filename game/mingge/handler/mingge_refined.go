package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	minggeeventtypes "fgame/fgame/game/mingge/event/types"
	minggelogic "fgame/fgame/game/mingge/logic"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	minggetemplate "fgame/fgame/game/mingge/template"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MINGGE_REFINED_TYPE), dispatch.HandlerFunc(handleMingGeRefined))
}

//处理命盘祭炼信息
func handleMingGeRefined(s session.Session, msg interface{}) (err error) {
	log.Debug("mingge:处理命盘祭炼信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMingGeRefined := msg.(*uipb.CSMingGeRefined)
	autoBuy := csMingGeRefined.GetAutoBuy()
	num := csMingGeRefined.GetNum()

	err = mingGeRefined(tpl, autoBuy, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoBuy":  autoBuy,
				"num":      num,
				"error":    err,
			}).Error("mingge:处理命盘祭炼信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mingge:处理命盘祭炼信息完成")
	return nil
}

//处理命盘祭炼信息逻辑
func mingGeRefined(pl player.Player, autoBuy bool, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"autoBuy":  autoBuy,
			"num":      num,
		}).Warn("mingge:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	maxSilver := propertyManager.GetSilver()
	maxBindGold := propertyManager.GetBindGlod()
	maxGold := propertyManager.GetGold()
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	mingGeAllSubTypeMap := make(map[minggetypes.MingGeAllSubType]bool)
	needItemMap := make(map[int32]int32)
	totalSilver := int64(0)
	totalGold := int64(0)
	totalBindGold := int64(0)
	refinedNum := int32(0)
	totalShopIdMap := make(map[int32]int32)
	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	for i := int32(1); i <= num; i++ {
		mingGePanId, allFull := manager.RefinedRandom()
		if allFull {
			if i == 1 {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoBuy":  autoBuy,
					"num":      num,
				}).Warn("mingge:所有命盘祭炼都满级了")
				playerlogic.SendSystemMessage(pl, lang.MingGeRefinedAllFull)
				return
			} else {
				//满级不用再祭炼了 截止到本次
				log.WithFields(log.Fields{
					"playerId":   pl.GetId(),
					"autoBuy":    autoBuy,
					"refinedNum": i,
				}).Warn("mingge:祭炼已全部满级")
				break
			}
		}
		curMingGePanIdTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingPanTemplateById(mingGePanId)
		if curMingGePanIdTemplate == nil {
			continue
		}
		mingGePanIdTemplate := curMingGePanIdTemplate.GetNextMingPanTemplate()
		if mingGePanIdTemplate == nil {
			continue
		}

		itemCount := mingGePanIdTemplate.UseItemCount
		totalNum := int32(0)
		useItem := mingGePanIdTemplate.UseItemId
		isEnoughBuyTimes := true
		if useItem != 0 {
			totalNum = inventoryManager.NumOfItems(useItem)
			curNeedNum, ok := needItemMap[useItem]
			if ok {
				totalNum -= curNeedNum
				if totalNum < 0 {
					totalNum = 0
				}
			}
		}
		if totalNum < itemCount {
			if autoBuy == false {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoBuy":  autoBuy,
					"num":      num,
				}).Warn("mingge:物品不足")
				playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
				return
			}
			shopIdMap := make(map[int32]int32)
			//自动购买
			needBuyNum := itemCount - totalNum

			if needBuyNum > 0 {
				if !shop.GetShopService().ShopIsSellItem(useItem) {
					if i == 1 {
						log.WithFields(log.Fields{
							"playerId": pl.GetId(),
							"autoBuy":  autoBuy,
						}).Warn("mingge:商铺没有该道具,无法自动购买")
						playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
						return
					} else {
						//本次祭炼已经无法购买物品了 截止到本次
						log.WithFields(log.Fields{
							"playerId":   pl.GetId(),
							"autoBuy":    autoBuy,
							"refinedNum": i,
						}).Warn("mingge:购买物品失败,自动祭炼失败")
						break
					}
				}

				curMaxSilver := maxSilver - totalSilver
				curMaxBindGold := maxBindGold - totalBindGold
				curMaxGold := maxGold - totalGold
				isEnoughBuyTimes, shopIdMap = shoplogic.GetPlayerShopCostForMaxMoney(pl, useItem, needBuyNum, curMaxGold, curMaxBindGold, curMaxSilver)
				if !isEnoughBuyTimes {
					if i == 1 {
						log.WithFields(log.Fields{
							"playerId": pl.GetId(),
							"autoBuy":  autoBuy,
						}).Warn("mingge:购买物品失败,自动祭炼失败")
						playerlogic.SendSystemMessage(pl, lang.ShopMingGeAutoBuyItemFail)
						return
					} else {
						//本次祭炼已经无法购买物品了 截止到本次
						log.WithFields(log.Fields{
							"playerId":   pl.GetId(),
							"autoBuy":    autoBuy,
							"refinedNum": i,
						}).Warn("mingge:购买物品失败,自动祭炼失败")
						break
					}
				}

				if totalNum > 0 {
					needItemMap[useItem] += totalNum
				}
				shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
				totalSilver += shopNeedSilver
				totalGold += shopNeedGold
				totalBindGold += shopNeedBindGold

				for shopId, num := range shopIdMap {
					totalShopIdMap[shopId] += num
				}
			}
		} else {
			needItemMap[useItem] += itemCount
		}
		//命盘祭炼
		mingPanType := mingGePanIdTemplate.GetMingPanType()
		mingPanRefinedNum := int32(0)
		mingPanRefinedPro := int32(0)
		obj := manager.GetMingGePanRefinedByType(mingPanType)
		if obj != nil {
			mingPanRefinedPro = obj.GetRefinedPro()
			mingPanRefinedNum = obj.GetRefinedNum()
		}
		refinedNum++
		mingGeAllSubTypeMap[mingGePanIdTemplate.GetMingPanType()] = true
		pro, _, sucess := minggelogic.MingGeMingPanRefined(mingPanRefinedNum, mingPanRefinedPro, mingGePanIdTemplate)
		flag := manager.Refined(mingPanType, pro, sucess)
		if !flag {
			panic("mingge: mingGeRefined Refined should be ok")
		}
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonMingGeRefinedCost.String()
	reasonSliverText := commonlog.SilverLogReasonMingGeRefinedCost.String()
	flag := propertyManager.Cost(totalBindGold, totalGold, commonlog.GoldLogReasonMingGeRefinedCost, reasonGoldText, totalSilver, commonlog.SilverLogReasonMingGeRefinedCost, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("mingge: mingGeRefined Cost should be ok"))
	}

	//消耗物品
	if len(needItemMap) != 0 {
		itemUsereason := commonlog.InventoryLogReasonMingGeRefinedUse
		if flag := inventoryManager.BatchRemove(needItemMap, itemUsereason, itemUsereason.String()); !flag {
			panic(fmt.Errorf("mingge: mingGeRefined BatchRemove should be ok"))
		}
	}

	//更新自动购买每日限购次数
	if len(totalShopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, totalShopIdMap)
	}

	//推送变化
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	eventdata := minggeeventtypes.CreatePlayerMingGeJiLianEventData(needItemMap)
	gameevent.Emit(minggeeventtypes.EventTypeMingGeJiLian, pl, eventdata)

	//更新属性
	minggelogic.MingGePropertyChanged(pl)
	mingGePanRefinedMap := manager.GetMingGePanRefinedMap()
	scMingGeRefined := pbutil.BuildSCMingGeRefined(mingGePanRefinedMap, mingGeAllSubTypeMap)
	pl.SendMsg(scMingGeRefined)
	return
}
