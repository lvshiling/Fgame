package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_RING_BAOKU_ATTEND_TYPE), dispatch.HandlerFunc(handleRingBaoKuAttend))
}

func handleRingBaoKuAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("ring: 开始处理宝库寻宝请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csRingBaoKuAttend := msg.(*uipb.CSRingBaoKuAttend)
	attendType := ringtypes.BaoKuAttendType(csRingBaoKuAttend.GetAttendType())
	if !attendType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"attendType": attendType,
			}).Warn("ring: 处理宝库寻宝请求消息,寻宝类型不合法")
		return
	}
	autoFlag := csRingBaoKuAttend.GetAutoFlag()
	typ := ringtypes.BaoKuType(csRingBaoKuAttend.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"typ":      typ,
			}).Warn("ring: 处理宝库寻宝请求消息,宝库类型不合法")
		return
	}

	err = ringBaoKuAttend(tpl, attendType, autoFlag, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("ring: 处理宝库寻宝请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("ring: 处理宝库寻宝请求消息")
	return nil

}

//批量探索特戒宝库逻辑
func ringBaoKuAttend(pl player.Player, attendType ringtypes.BaoKuAttendType, autoFlag bool, typ ringtypes.BaoKuType) (err error) {
	playerId := pl.GetId()
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeRingXunBao) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring: 宝库寻宝错误,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	attendNum := attendType.Int()
	ringBaoKuManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	ringBaoKuTemplate := ringtemplate.GetRingTemplateService().GetRingBaoKuTemplate(typ)
	if ringBaoKuTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring: 宝库寻宝错误,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	// 计算需要的钱和物品
	needBindGold := int32(ringBaoKuTemplate.BindGoldUse * attendNum)
	needGold := int32(ringBaoKuTemplate.GoldUse * attendNum)
	needSilver := int64(ringBaoKuTemplate.SilverUse * attendNum)
	needItemId := ringBaoKuTemplate.UseItemId
	needItemCount := ringBaoKuTemplate.UseItemCount * attendNum

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	//物品是否足够
	totalNum := inventoryManager.NumOfItems(int32(needItemId))
	if totalNum < needItemCount {
		if !autoFlag {
			log.WithFields(
				log.Fields{
					"playerId":      playerId,
					"typ":           typ.String(),
					"needItemId":    needItemId,
					"needItemCount": needItemCount,
				}).Warn("ring:宝库寻宝错误，道具不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动进阶
		needBuyNum := needItemCount - totalNum
		needItemCount = totalNum
		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(needItemId) {
				log.WithFields(log.Fields{
					"playerId": playerId,
					"typ":      typ.String(),
					"autoFlag": autoFlag,
				}).Warn("ring:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, needItemId, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": playerId,
					"typ":      typ.String(),
					"autoFlag": autoFlag,
				}).Warn("ring:购买物品失败,宝库探索失败")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNumInvalid)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			needGold += int32(shopNeedGold)
			needBindGold += int32(shopNeedBindGold)
			needSilver += shopNeedSilver
		}

	}

	//是否足够银两
	flag := propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring:宝库寻宝错误，银两不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(int64(needGold), false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring:宝库寻宝错误，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//是否足够绑元
	needCostBindGold := needBindGold + needGold
	flag = propertyManager.HasEnoughGold(int64(needCostBindGold), true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring:宝库寻宝错误，绑元不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//批量处理
	rewList := ringBaoKuManager.GetRingBaoKuDrop(attendNum, typ)
	if len(rewList) == 0 {
		log.WithFields(
			log.Fields{
				"playerId":     playerId,
				"typ":          typ.String(),
				"attendTimes ": attendNum,
			}).Warn("ring:宝库寻宝错误，掉落为空")
		playerlogic.SendSystemMessage(pl, lang.RingBaoKuNotGetRewards)
		return
	}

	//背包空间
	var rewItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(rewList) > 0 {
		rewItemList, resMap = droplogic.SeperateItemDatas(rewList)
	}

	if !inventoryManager.HasEnoughSlotsOfItemLevel(rewItemList) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"typ":      typ.String(),
			}).Warn("ring:宝库寻宝错误, 背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryMiBaoDepotSlotNoEnough)
		return
	}

	//消耗银两
	if needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonRingBaoKuUse
		flag := propertyManager.CostSilver(needSilver, silverUseReason, silverUseReason.String())
		if !flag {
			panic("ring:消耗银两应该成功")
		}
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonRingBaoKuUse
		flag := propertyManager.CostGold(int64(needGold), false, goldUseReason, goldUseReason.String())
		if !flag {
			panic("ring:消耗元宝应该成功")
		}
	}

	//消耗绑元
	if needBindGold > 0 {
		goldUseReason := commonlog.GoldLogReasonRingBaoKuUse
		flag := propertyManager.CostGold(int64(needBindGold), true, goldUseReason, goldUseReason.String())
		if !flag {
			panic("ring:消耗元宝应该成功")
		}
	}

	//消耗物品
	if needItemCount > 0 {
		itemUseReason := commonlog.InventoryLogReasonRingBaoKuAttend
		if flag := inventoryManager.UseItem(needItemId, needItemCount, itemUseReason, itemUseReason.String()); !flag {
			panic(fmt.Errorf("ring: attend ring use item should be ok"))
		}
	}

	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonRingBaoKuGet
		silverReason := commonlog.SilverLogReasonRingBaoKuGet
		levelReason := commonlog.LevelLogReasonRingBaoKuGet
		err = droplogic.AddRes(pl, resMap, goldReason, goldReason.String(), silverReason, silverReason.String(), levelReason, levelReason.String())
		if err != nil {
			return
		}
	}

	if len(rewItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonRingBaoKuGet
		flag = inventoryManager.BatchAddOfItemLevel(rewItemList, itemGetReason, itemGetReason.String())
		if !flag {
			panic("ring:增加物品应该成功")
		}
	}

	//宝库积分,幸运值
	addXingYunZhi := ringBaoKuTemplate.GiftXingYunZhi * attendNum
	addJiFen := ringBaoKuTemplate.GiftJiFen * attendNum

	// 判断幸运值是否已满
	obj := ringBaoKuManager.GetPlayerBaoKuObject(typ)
	baoKuTemp := ringtemplate.GetRingTemplateService().GetRingBaoKuTemplate(typ)
	if baoKuTemp == nil {
		return
	}
	var extraRewList []*droptemplate.DropItemData
	if obj.GetLuckyPoints()+addXingYunZhi > baoKuTemp.NeedXingYunZhi {
		flag, extraRewList = ringlogic.RingLuckyPointsTop(pl, typ)
		if !flag {
			return
		}
	}
	flag = ringBaoKuManager.AttendRingBaoKu(addXingYunZhi, addJiFen, attendNum, typ)
	if !flag {
		panic("ring: 宝库寻宝完成刷新数据应该成功")
	}

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scRingBaoKuAttendBatch := pbutil.BuildSCRingBaoKuAttend(autoFlag, int32(attendType), obj, rewItemList, extraRewList)
	pl.SendMsg(scRingBaoKuAttendBatch)
	return
}
