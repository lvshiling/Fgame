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
	zhenfalogic "fgame/fgame/game/zhenfa/logic"
	"fgame/fgame/game/zhenfa/pbutil"
	playerzhenfa "fgame/fgame/game/zhenfa/player"
	zhenfatemplate "fgame/fgame/game/zhenfa/template"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ZHENFA_SHENGJI_TYPE), dispatch.HandlerFunc(handleZhenFaShengJi))
}

//处理阵法升级信息
func handleZhenFaShengJi(s session.Session, msg interface{}) (err error) {
	log.Debug("zhenfa:处理阵法升级信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csZhenFaShengJi := msg.(*uipb.CSZhenFaShengJi)
	autoFlag := csZhenFaShengJi.GetAutoFlag()
	zhenFaType := csZhenFaShengJi.GetZhenFaType()
	err = zhenFaShengJi(tpl, autoFlag, zhenfatypes.ZhenFaType(zhenFaType))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"autoFlag":   autoFlag,
				"zhenFaType": zhenFaType,
				"error":      err,
			}).Error("zhenfa:处理阵法升级信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("zhenfa:处理阵法升级信息完成")
	return nil
}

//处理阵法升级信息逻辑
func zhenFaShengJi(pl player.Player, autoFlag bool, zhenFaType zhenfatypes.ZhenFaType) (err error) {
	if !zhenFaType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"autoFlag":   autoFlag,
			"zhenFaType": zhenFaType,
		}).Warn("zhenfa:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	obj := manager.GetZhenFaByType(zhenFaType)
	if obj == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"autoFlag":   autoFlag,
			"zhenFaType": zhenFaType,
		}).Warn("zhenfa:未激活的阵法,无法升级")
		playerlogic.SendSystemMessage(pl, lang.ZhenFaShengJiNoActivate)
		return
	}

	curLevel := obj.GetLevel()
	nextLevel := curLevel + 1
	zhenFaTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaTempalte(zhenFaType, nextLevel)
	if zhenFaTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"autoFlag":   autoFlag,
			"zhenFaType": zhenFaType,
		}).Warn("zhenfa:阵法已达最高级")
		playerlogic.SendSystemMessage(pl, lang.ZhenFaShengJiFullLevel)
		return
	}

	needItems := zhenFaTemplate.GetNeedItemMap()
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(needItems) != 0 {
		if !inventoryManager.HasEnoughItems(needItems) && !autoFlag {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"autoFlag":   autoFlag,
				"zhenFaType": zhenFaType,
			}).Warn("zhenfa:物品不足")
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

		isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayerMap(pl, buyItems)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"autoFlag":   autoFlag,
				"zhenFaType": zhenFaType,
			}).Warn("zhenfa:购买物品失败")
			playerlogic.SendSystemMessage(pl, lang.ShopZhenFaShengJiAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		gold += int32(shopNeedGold)
		bindGold += int32(shopNeedBindGold)
		sliver += shopNeedSilver

		flag := propertyManager.HasEnoughCost(int64(bindGold), int64(gold), sliver)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"autoFlag":   autoFlag,
				"zhenFaType": zhenFaType,
			}).Warn("zhenfa:元宝不足，无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}

		reasonGoldText := commonlog.GoldLogReasonZhenFaShengJiCost.String()
		reasonSliverText := commonlog.SilverLogReasonZhenFaShengJiCost.String()
		flag = propertyManager.Cost(int64(bindGold), int64(gold), commonlog.GoldLogReasonZhenFaShengJiCost, reasonGoldText, sliver, commonlog.SilverLogReasonZhenFaShengJiCost, reasonSliverText)
		if !flag {
			panic(fmt.Errorf("zhenfa: zhenFaShengJi Cost should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗物品
	if len(items) != 0 {
		reasonText := commonlog.InventoryLogReasonZhenFaShengJi.String()
		flag := inventoryManager.BatchRemove(items, commonlog.InventoryLogReasonZhenFaShengJi, reasonText)
		if !flag {
			panic(fmt.Errorf("zhenfa: zhenFaShengJi use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}
	pro, _, sucess := zhenfalogic.ZhenFaShengJi(obj.GetLevelNum(), obj.GetLevelPro(), zhenFaTemplate)
	flag := manager.ZhenFaShengJi(zhenFaType, sucess, pro)
	if !flag {
		panic(fmt.Errorf("zhenfa: zhenFaShengJi should be ok"))
	}

	// 属性变化
	zhenfalogic.ZhenFaPropertyChanged(pl)

	scZhenFaShengJi := pbutil.BuildSCZhenFaShengJi(sucess, obj)
	pl.SendMsg(scZhenFaShengJi)
	return
}
