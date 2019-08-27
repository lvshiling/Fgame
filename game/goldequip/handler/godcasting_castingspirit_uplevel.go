package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	commonlogic "fgame/fgame/game/common/logic"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

const (
	MAXLEVEL = int32(999)
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GODCASTING_CASTINGSPIRIT_UPLEVEL_TYPE), dispatch.HandlerFunc(handleGodCastingCastingSpiritUplevel))
}

func handleGodCastingCastingSpiritUplevel(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGodCastingCastingSpiritUplevel)
	bodyPos := inventorytypes.BodyPositionType(csMsg.GetBodyPos())
	spiritType := goldequiptypes.SpiritType(csMsg.GetSpiritType())
	autoFlag := csMsg.GetAutoFlag()
	if !bodyPos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("goldequip:处理神铸铸灵升级请求失败，装备部位不合法")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	if !spiritType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"spiritType": spiritType,
			}).Warn("goldequip:处理神铸铸灵升级请求失败，铸灵类型不合法")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = godCastingCastingSpiritUplevel(tpl, bodyPos, spiritType, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理神铸铸灵升级请求,错误")
		return err
	}
	return
}

func godCastingCastingSpiritUplevel(pl player.Player, bodyPos inventorytypes.BodyPositionType, spiritType goldequiptypes.SpiritType, autoFlag bool) (err error) {
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	castingSpiritTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetCastingSpiritTemplate(bodyPos, spiritType)
	goldequipBag := goldequipManager.GetGoldEquipBag()

	//判断是否有装备
	equip := goldequipBag.GetByPosition(bodyPos)
	if equip == nil || equip.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸铸灵升级请求,装备未装上")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentSlotNoEquip)
		return
	}

	//判断是不是神铸装备
	itemTemp := item.GetItemService().GetItem(int(equip.GetItemId()))
	if itemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸铸灵升级请求,物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	goldequipTemp := itemTemp.GetGoldEquipTemplate()
	if goldequipTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸铸灵升级请求,元神金装模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	if !goldequipTemp.IsGodCastingEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸铸灵升级请求,不是神铸装备")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	curLevel := goldequipTemp.GetGodCastingEquipLevel()

	//判断是否解锁了铸灵
	if !castingSpiritTemplate.IsActive(curLevel) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"pos":        bodyPos.String(),
				"spiritType": spiritType,
				"curLevel":   curLevel,
			}).Warn("goldequip:处理神铸铸灵升级请求,神铸等级不足")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipGodCastingLevelNotEnough)
		return
	}

	//获取铸灵升级模板
	spiritInfo := equip.GetCastingSpiritInfo(spiritType)
	spiritUpLevelTemp := castingSpiritTemplate.GetLevelTemplate(spiritInfo.Level + 1) //升级条件取下一级数据
	if spiritUpLevelTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"pos":         bodyPos.String(),
				"spiritType":  spiritType,
				"spiritLevel": spiritInfo.Level,
			}).Warn("goldequip:处理神铸铸灵升级请求,铸灵升级模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//判断铸灵是否满级
	if spiritUpLevelTemp.IsMaxLevel() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"pos":         bodyPos.String(),
				"spiritType":  spiritType,
				"spiritLevel": spiritInfo.Level,
			}).Warn("goldequip:处理神铸铸灵升级请求,铸灵已经满级")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipGodCastingCastingSpiritLevelFull)
		return
	}

	//判断物品够不够
	useItemId := castingSpiritTemplate.UseItemId
	useItemCnt := spiritUpLevelTemp.UseItemCount
	curUseItemNum := inventoryManager.NumOfItems(useItemId)
	finalUseItemNum := useItemCnt
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	needGold := int64(0)
	needBindGold := int64(0)
	needSilver := int64(0)
	if curUseItemNum < useItemCnt {
		if !autoFlag {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"pos":        bodyPos.String(),
					"spiritType": spiritType,
					"useItemId":  useItemId,
					"useItemCnt": useItemCnt,
				}).Warn("goldequip:处理神铸铸灵升级请求,铸灵升级物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动进阶
		needBuyNum := useItemCnt - curUseItemNum
		finalUseItemNum = curUseItemNum
		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItemId) {
				log.WithFields(log.Fields{
					"playerId":  pl.GetId(),
					"useItemId": useItemId,
					"autoFlag":  autoFlag,
				}).Warn("goldequip:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItemId, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId":  pl.GetId(),
					"useItemId": useItemId,
					"autoFlag":  autoFlag,
				}).Warn("goldequip:购买物品失败,铸灵升级失败")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNumInvalid)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			needGold += shopNeedGold
			needBindGold += shopNeedBindGold
			needSilver += shopNeedSilver

		}
	}

	//是否足够银两
	flag := propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("goldequip:铸灵升级，银两不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(needGold, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("goldequip:铸灵升级，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//是否足够绑元
	needCostBindGold := needBindGold + needGold
	flag = propertyManager.HasEnoughGold(needCostBindGold, true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("goldequip:铸灵升级，绑元不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//部位对象升级处理
	updateRate := spiritUpLevelTemp.UpdateWfb
	blessMax := spiritUpLevelTemp.ZhufuMax
	addMin := spiritUpLevelTemp.AddMin
	addMax := spiritUpLevelTemp.AddMax + 1

	randBless := int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes := int32(1)
	curTimesNum := spiritInfo.Times
	curTimesNum += addTimes
	curBless := spiritInfo.Bless

	pro, sucess := commonlogic.AdvancedStatusAndProgress(curTimesNum, curBless, spiritUpLevelTemp.TimesMin, spiritUpLevelTemp.TimesMax, randBless, updateRate, blessMax)
	isSuccess := int32(0)
	if sucess {
		isSuccess = int32(1)
	}
	// equip.UplevelSpirit(spiritType, sucess, pro)
	goldequipBag.UplevelSpirit(bodyPos, spiritType, sucess, pro)

	//自动购买消耗金钱
	//消耗银两
	if needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonCastingSpiritUplevel
		silverUseReasonText := fmt.Sprintf(silverUseReason.String())
		flag := propertyManager.CostSilver(needSilver, silverUseReason, silverUseReasonText)
		if !flag {
			panic("goldequip:消耗银两应该成功")
		}
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonCastingSpiritUpLevel
		goldUseReasonText := fmt.Sprintf(goldUseReason.String())
		flag := propertyManager.CostGold(needGold, false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("goldequip:消耗元宝应该成功")
		}
	}

	//消耗绑元
	if needBindGold > 0 {
		goldUseReason := commonlog.GoldLogReasonCastingSpiritUpLevel
		goldUseReasonText := fmt.Sprintf(goldUseReason.String())
		flag := propertyManager.CostGold(needBindGold, true, goldUseReason, goldUseReasonText)
		if !flag {
			panic("goldequip:消耗元宝应该成功")
		}
	}

	//同步物品（删掉吃掉的物品）
	useReason := commonlog.InventoryLogReasonCastingSpiritUplevel

	useItemTemp := item.GetItemService().GetItem(int(useItemId))
	useReasonText := fmt.Sprintf(useReason.String(), useItemTemp.Name, useItemCnt, bodyPos.String(), spiritType.String())
	if curUseItemNum > 0 {
		flag = inventoryManager.UseItem(useItemId, finalUseItemNum, useReason, useReasonText)
		if !flag {
			panic("inventory:移除物品应该是可以的")
		}
	}

	goldequiplogic.GoldEquipPropertyChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGodCastingCastingSpiritUpLevel(bodyPos, spiritType, spiritInfo, isSuccess)
	pl.SendMsg(scMsg)
	return
}
