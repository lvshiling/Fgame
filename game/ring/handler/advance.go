package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	ringlogic "fgame/fgame/game/ring/logic"
	"fgame/fgame/game/ring/pbutil"
	playerring "fgame/fgame/game/ring/player"
	ringtemplate "fgame/fgame/game/ring/template"
	ringtypes "fgame/fgame/game/ring/types"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RING_ADVANCE_TYPE), dispatch.HandlerFunc(handleRingAdvance))
}

func handleRingAdvance(s session.Session, msg interface{}) (err error) {
	log.Debug("ring: 开始处理特戒进阶请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRingAdvance := msg.(*uipb.CSRingAdvance)
	typ := ringtypes.RingType(csRingAdvance.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"type":     int32(typ),
			}).Warn("ring: 特戒类型不合法")
		return
	}
	autoFlag := csRingAdvance.GetAutoFlag()

	err = ringAdvance(tpl, typ, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("ring: 处理特戒进阶请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("ring: 处理特戒进阶请求消息,成功")

	return
}

func ringAdvance(pl player.Player, typ ringtypes.RingType, autoFlag bool) (err error) {
	playerId := pl.GetId()
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeRingAdvance) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring: 特戒进阶失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	ringObj := ringManager.GetPlayerRingObject(typ)
	if ringObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring: 玩家未穿戴该特戒")
		playerlogic.SendSystemMessage(pl, lang.RingNotEquip)
		return
	}

	// 特戒数据
	itemId := ringObj.GetItemId()
	data := ringObj.GetPropertyData()
	ringData := data.(*ringtypes.RingPropertyData)
	advance := ringData.Advance + 1

	advanceTemp := ringtemplate.GetRingTemplateService().GetRingAdvanceTemplate(itemId, advance)
	if advanceTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring: 特戒进阶已达到最高")
		playerlogic.SendSystemMessage(pl, lang.RingAdvanceAlreadyTop)
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
	curNum := inventoryManager.NumOfItems(advanceTemp.UseItem)
	needNum := advanceTemp.ItemCount - curNum
	useItem := advanceTemp.UseItem
	useNum := int32(0)
	if needNum > 0 {
		if autoFlag == false {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"typ":      typ.String(),
				}).Warn("ring: 所需物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		useNum = curNum
		if !shop.GetShopService().ShopIsSellItem(useItem) {
			log.WithFields(log.Fields{
				"playerId": playerId,
				"autoFlag": autoFlag,
			}).Warn("ring:商铺没有该道具,无法自动购买")
			playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
			return
		}

		isEnoughBuyTimes, shopIdMap := shoplogic.MaxBuyTimesForPlayer(pl, useItem, needNum)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerId": playerId,
				"autoFlag": autoFlag,
			}).Warn("ring:购买物品失败,自动进阶已停止")
			playerlogic.SendSystemMessage(pl, lang.ShopAdvancedAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		costGold += shopNeedGold
		costBindGold += shopNeedBindGold
		costSilver += shopNeedSilver
	} else {
		useNum = advanceTemp.ItemCount
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(costSilver)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": playerId,
				"autoFlag": autoFlag,
			}).Warn("ring:银两不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(costGold, false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": playerId,
				"autoFlag": autoFlag,
			}).Warn("ring:元宝不足,无法进阶")
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
				"playerId": playerId,
				"autoFlag": autoFlag,
			}).Warn("ring:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonRingAdvance
	goldUseReasonStr := fmt.Sprintf(goldUseReason.String(), typ.String())
	silverUseReason := commonlog.SilverLogReasonRingAdvance
	silverUseReasonStr := fmt.Sprintf(silverUseReason.String(), typ.String())
	flag := propertyManager.Cost(costBindGold, costGold, goldUseReason, goldUseReasonStr, costSilver, silverUseReason, silverUseReasonStr)
	if !flag {
		panic(fmt.Errorf("ring: 特戒进阶自动购买应该成功"))
	}

	// 消耗物品
	if useNum > 0 {
		reason := commonlog.InventoryLogReasonRingAdvance
		reasonText := fmt.Sprintf(reason.String(), typ.String())
		flag = inventoryManager.UseItem(useItem, useNum, reason, reasonText)
		if !flag {
			panic("ring: 特戒进阶消耗物品应该成功")
		}
	}

	// 物品改变推送
	inventorylogic.SnapInventoryChanged(pl)

	// 成功率判断
	pro, randBless, success := ringlogic.RingAdvance(ringData.AdvanceNum, ringData.AdvancePro, advanceTemp)
	if !success {
		advance--
	}

	// 刷新数据
	flag = ringManager.RingAdvanceSuccess(typ, success, pro)
	if !flag {
		panic("ring: 特戒进阶刷新数据应该成功")
	}

	// 推送属性变化
	ringlogic.RingPropertyChange(pl)
	propertylogic.SnapChangedProperty(pl)

	// 公告
	if success {
		// 计算激活的属性百分比
		hpPercent := advanceTemp.HpPercent
		attackPercent := advanceTemp.AttackPercent
		defPercent := advanceTemp.DefPercent
		lastAdvanceTemp := ringtemplate.GetRingTemplateService().GetRingAdvanceTemplate(itemId, advance-1)
		if lastAdvanceTemp != nil {
			hpPercent -= lastAdvanceTemp.HpPercent
			attackPercent -= lastAdvanceTemp.AttackPercent
			defPercent -= lastAdvanceTemp.DefPercent
		}

		// 转换
		propertyStr := ""
		if hpPercent > 0 {
			propertyStr = fmt.Sprintf("角色防御增加%d%%", hpPercent/100)
		} else if attackPercent > 0 {
			propertyStr = fmt.Sprintf("角色防御增加%d%%", attackPercent/100)
		} else if defPercent > 0 {
			propertyStr = fmt.Sprintf("角色防御增加%d%%", defPercent/100)
		}

		if propertyStr != "" {
			plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
			itemTemp := item.GetItemService().GetItem(int(itemId))
			if itemTemp == nil {
				log.Warningf("ring: 物品模板不存在,itemId:%d", itemId)
				return
			}
			qualityType := itemtypes.ItemQualityType(itemTemp.Quality)
			if qualityType < itemtypes.ItemQualityTypeOrange {
				return
			}
			propertyData := ringManager.GetPlayerRingObject(typ).GetPropertyData()
			data, ok := propertyData.(*ringtypes.RingPropertyData)
			if !ok {
				fmt.Println("ring: 数据类型转换错误")
				return
			}
			itemName := coreutils.FormatColor(qualityType.GetColor(), fmt.Sprintf("[%s]", typ.String()))
			args := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId), int64(data.StrengthLevel), int64(data.Advance), int64(data.JingLingLevel)}
			infoLink := coreutils.FormatLink(itemName, args)
			content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.RingAdvanceNotice), plName, infoLink, advance, propertyStr)
			chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
			noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
		}

	}

	scRingAdvance := pbutil.BuildSCRingAdvance(success, int32(typ), advance, ringData.AdvancePro, randBless)
	pl.SendMsg(scRingAdvance)
	return
}
