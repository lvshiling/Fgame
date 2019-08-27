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
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	weaponlogic "fgame/fgame/game/weapon/logic"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
	"fgame/fgame/game/weapon/weapon"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WEAPON_UPSTAR_TYPE), dispatch.HandlerFunc(handleWeaponUpstar))
}

//处理兵魂升星信息
func handleWeaponUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("weapon:处理兵魂升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWeaponUpstar := msg.(*uipb.CSWeaponUpstar)
	weaponId := csWeaponUpstar.GetWeaponId()
	autoFlag := csWeaponUpstar.GetAutoFlag()

	err = weaponUpstar(tpl, weaponId, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weaponId": weaponId,
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("weapon:处理兵魂升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("weapon:处理兵魂升星完成")
	return nil
}

//兵魂升星的逻辑
func weaponUpstar(pl player.Player, weaponId int32, autoFlag bool) (err error) {
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	flag := weaponManager.IfWeaponExist(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
			"autoFlag": autoFlag,
		}).Warn("weapon:未激活的兵魂,无法升星")
		playerlogic.SendSystemMessage(pl, lang.WeaponNotActiveNotUpstar)
		return
	}

	flag = weaponManager.IfCanUpStar(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
			"autoFlag": autoFlag,
		}).Warn("weapon:兵魂已满星")
		playerlogic.SendSystemMessage(pl, lang.WeaponReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	weaponInfo := weaponManager.GetWeapon(weaponId)
	if weaponInfo == nil {
		return
	}
	curLevel := weaponInfo.Level
	nextLevel := curLevel + 1
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if to == nil {
		return
	}
	weaponUpstarTemplate := to.GetWeaponUpstarByLevel(nextLevel)
	if weaponUpstarTemplate == nil {
		return
	}

	needItems := weaponUpstarTemplate.GetNeedItemMap(pl.GetRole())
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag && autoFlag == false {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"weaponId": weaponId,
				"autoFlag": autoFlag,
			}).Warn("weapon:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	//获取背包物品和需要购买物品
	items, buyItems := inventoryManager.GetItemsAndNeedBuy(needItems)
	//计算需要元宝等
	if len(buyItems) != 0 {
		bindGold := int32(0)
		gold := int32(0)
		sliver := int64(0)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		// currencyCost, flag := shop.GetShopService().GetShopCost(buyItems)
		// if !flag {
		// 	log.WithFields(log.Fields{
		// 		"playerid": pl.GetId(),
		// 		"weaponId": weaponId,
		// 		"autoFlag": autoFlag,
		// 	}).Warn("weapon:商铺没有该道具,无法自动购买")
		// 	playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
		// 	return
		// }
		// bindGold, _ = currencyCost[shoptypes.ShopConsumeTypeBindGold]
		// gold, _ = currencyCost[shoptypes.ShopConsumeTypeGold]
		// sliver = int64(currencyCost[shoptypes.ShopConsumeTypeSliver])

		isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayerMap(pl, buyItems)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"weaponId": weaponId,
				"autoFlag": autoFlag,
			}).Warn("weapon:购买物品失败,自动升星已停止")
			playerlogic.SendSystemMessage(pl, lang.ShopUpstarAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		gold += int32(shopNeedGold)
		bindGold += int32(shopNeedBindGold)
		sliver += shopNeedSilver

		flag = propertyManager.HasEnoughCost(int64(bindGold), int64(gold), sliver)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"weaponId": weaponId,
				"autoFlag": autoFlag,
			}).Warn("weapon:元宝不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}

		reasonGoldText := commonlog.GoldLogReasonWeapUpstar.String()
		reasonSliverText := commonlog.SilverLogReasonWeapUpstar.String()
		flag = propertyManager.Cost(int64(bindGold), int64(gold), commonlog.GoldLogReasonWeapUpstar, reasonGoldText, sliver, commonlog.SilverLogReasonWeapUpstar, reasonSliverText)
		if !flag {
			panic(fmt.Errorf("weapon: weaponUpstar Cost should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗物品
	if len(items) != 0 {
		reasonText := commonlog.InventoryLogReasonWeaponUpstar.String()
		flag := inventoryManager.BatchRemove(items, commonlog.InventoryLogReasonWeaponUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("weapon: weaponUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//兵魂升星判断
	pro, _, sucess := weaponlogic.WeaponUpStar(pl, weaponInfo.UpNum, weaponInfo.UpPro, weaponUpstarTemplate)
	flag = weaponManager.Upstar(weaponId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("weapon: weaponUpstar should be ok"))
	}
	if sucess {
		//同步属性
		weaponlogic.WeaponPropertyChanged(pl)
	}
	scWeaponUpstar := pbutil.BuildSCWeaponUpstar(weaponId, weaponInfo.Level, weaponInfo.UpPro)
	pl.SendMsg(scWeaponUpstar)
	return
}
