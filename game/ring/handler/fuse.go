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
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/inventory/inventory"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
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
	skilltemplate "fgame/fgame/game/skill/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RING_FUSE_TYPE), dispatch.HandlerFunc(handleRingFuse))
}

func handleRingFuse(s session.Session, msg interface{}) (err error) {
	log.Debug("ring: 开始处理特戒融合请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRingFuse := msg.(*uipb.CSRingFuse)
	isBag := csRingFuse.GetIsBag()
	typ := ringtypes.RingType(0)
	index := int32(0)
	if isBag {
		index = csRingFuse.GetIndex()
	} else {
		typ := ringtypes.RingType(csRingFuse.GetType())
		if !typ.Valid() {
			log.WithFields(
				log.Fields{
					"playerId": tpl.GetId(),
					"type":     int32(typ),
				}).Warn("ring: 特戒类型不合法")
			return
		}
	}

	needIndex := csRingFuse.GetNeedIndex()

	err = ringFuse(tpl, typ, isBag, index, needIndex)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("ring: 处理特戒融合请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("ring: 处理特戒融合请求消息,成功")

	return
}

func ringFuse(pl player.Player, typ ringtypes.RingType, isBag bool, index int32, needIndex int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeRingFuse) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("ring: 功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)

	itemId := int32(0)
	var itemType itemtypes.ItemType
	var propertyData inventorytypes.ItemPropertyData
	// 判断第一个槽里的特戒是否在背包
	if isBag {
		it := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
		//物品不存在
		if it == nil || it.IsEmpty() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("ring: 使用特戒,物品不存在")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
			return
		}
		itemId = it.ItemId
		propertyData = it.PropertyData

		// 判断物品是否未特戒
		itemTemplate := item.GetItemService().GetItem(int(itemId))
		if !itemTemplate.IsTeRing() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("ring: 使用特戒, 物品不是特戒")
			playerlogic.SendSystemMessage(pl, lang.RingIsNotRing)
			return
		}
		itemType = itemTemplate.GetItemType()

		fuseTemp := ringtemplate.GetRingTemplateService().GetRingFuseSynthesisTemplate(itemId)
		if fuseTemp == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("ring: 使用特戒,物品不存在")
			playerlogic.SendSystemMessage(pl, lang.RingTempalteNotExist)
			return
		}

		createItemId := fuseTemp.ItemId
		createItemNum := fuseTemp.ItemCount

		// 判断背包是否足够
		if !inventoryManager.HasEnoughSlot(createItemId, createItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":      pl.GetId(),
					"createItemId":  createItemId,
					"createItemNum": createItemNum,
				}).Warn("ring: 背包不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}

		ringTemp := ringtemplate.GetRingTemplateService().GetRingTemplate(itemId)
		if ringTemp == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
				}).Warn("ring:模板不存在")
			playerlogic.SendSystemMessage(pl, lang.RingTempalteNotExist)
			return
		}
		typ = ringTemp.GetRingType()

	} else {
		ringObj := ringManager.GetPlayerRingObject(typ)
		if ringObj == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ.String(),
				}).Warn("ring: 玩家未穿戴该特戒")
			playerlogic.SendSystemMessage(pl, lang.RingNotEquip)
			return
		}

		itemId = ringObj.GetItemId()
	}

	needIt := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, needIndex)
	// 第二槽位的物品不存在
	if needIt == nil || needIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("ring:使用特戒,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	needItemId := needIt.ItemId

	// 判断第二槽位的物品是否为特戒
	needItemTemplate := item.GetItemService().GetItem(int(needItemId))
	if !needItemTemplate.IsTeRing() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"needIndex":  needIndex,
				"needItemId": needItemId,
			}).Warn("ring: 使用特戒, 物品不是特戒")
		playerlogic.SendSystemMessage(pl, lang.RingIsNotRing)
		return
	}

	// 融合模板
	fuseTemp := ringtemplate.GetRingTemplateService().GetRingFuseSynthesisTemplate(itemId)
	if fuseTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"itemId":   itemId,
			}).Warn("ring: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.RingTempalteNotExist)
		return
	}
	if needItemId != fuseTemp.NeedItemId2 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"typ":         typ.String(),
				"needIndex":   needIndex,
				"needItemId":  needItemId,
				"NeedItemId2": fuseTemp.NeedItemId2,
			}).Warn("ring: 特戒融合需要物品与当前物品不符")
		playerlogic.SendSystemMessage(pl, lang.RingFuseItemNotSuit)
		return
	}

	// 消耗的钱
	costGold := int64(fuseTemp.NeedGold)
	costSilver := int64(fuseTemp.NeedSilver)
	costBindGold := int64(fuseTemp.NeedBindGold)

	// 是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(costSilver)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring: 银两不足,无法融合")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	// 是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(costGold, false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring:元宝不足,无法融合")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	// 是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(needBindGold, true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring:元宝不足,无法融合")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	// 判断物品是否足够
	needItemNum := fuseTemp.NeedItemCount2
	curNum := inventoryManager.NumOfItems(needItemId)
	if curNum < needItemNum {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring: 所需物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonRingFuse
	goldUseReasonStr := fmt.Sprintf(goldUseReason.String(), typ.String())
	silverUseReason := commonlog.SilverLogReasonRingFuse
	silverUseReasonStr := fmt.Sprintf(silverUseReason.String(), typ.String())
	flag := propertyManager.Cost(costBindGold, costGold, goldUseReason, goldUseReasonStr, costSilver, silverUseReason, silverUseReasonStr)
	if !flag {
		panic(fmt.Errorf("ring: 特戒融合消耗钱应该成功"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	if needItemNum > 0 {
		reason := commonlog.InventoryLogReasonRingAdvance
		reasonText := fmt.Sprintf(reason.String(), typ.String())
		flag, err = inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, needIndex, needItemNum, reason, reasonText)
		if !flag {
			panic("ring: 消耗物品应该成功")
		}
		if err != nil {
			return
		}
	}

	success := mathutils.RandomHit(common.MAX_RATE, int(fuseTemp.SuccessRate))

	// 成功消耗第一槽位物品
	if success && isBag {
		reason := commonlog.InventoryLogReasonRingAdvance
		reasonText := fmt.Sprintf(reason.String(), typ.String())
		flag, err = inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, index, 1, reason, reasonText)
		if !flag {
			panic("ring: 消耗物品应该成功")
		}
		if err != nil {
			return
		}
	}

	createItemId := fuseTemp.ItemId
	createItemNum := fuseTemp.ItemCount

	if isBag {
		if success {
			createItemTemp := item.GetItemService().GetItem(int(createItemId))
			if createItemTemp == nil {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"typ":      typ.String(),
					}).Warn("ring: 融合成功的物品模板不存在")
				playerlogic.SendSystemMessage(pl, lang.RingTempalteNotExist)
				return
			}

			reason := commonlog.InventoryLogReasonRingFuseGet
			reasonText := fmt.Sprintf(reason.String(), typ.String())
			flag = inventoryManager.AddItemLevelWithPropertyData(createItemId, createItemNum, createItemTemp.NeedLevel, createItemTemp.GetBindType(), propertyData, reason, reasonText)
			if !flag {
				panic(fmt.Errorf("ring: 添加物品应该成功"))
			}
		}
	} else {
		if success {
			ringManager.RingFuseSuccess(typ, createItemId)
		}
	}

	// 物品改变推送
	inventorylogic.SnapInventoryChanged(pl)

	// 推送属性变化
	ringlogic.RingPropertyChange(pl)
	propertylogic.SnapChangedProperty(pl)

	// 公告
	plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	fuseNum := int32(0)
	skillId := int32(0)

	rongheTemp := ringtemplate.GetRingTemplateService().GetRingTemplate(createItemId)
	if rongheTemp != nil {
		fuseNum = rongheTemp.Level
		skillId = rongheTemp.SkillId
	}
	itemTemp := item.GetItemService().GetItem(int(createItemId))
	if itemTemp == nil {
		log.Warningf("ring: 物品模板不存在,itemId:%d", createItemId)
		return
	}

	qualityType := itemtypes.ItemQualityType(itemTemp.Quality)
	itemName := coreutils.FormatColor(qualityType.GetColor(), fmt.Sprintf("[%s]", typ.String()))

	data, ok := propertyData.(*ringtypes.RingPropertyData)
	if !ok {
		base := inventorytypes.CreateDefaultItemPropertyDataBase()
		propertyData = inventory.CreatePropertyDataInterface(itemType, base)
	}

	args := []int64{int64(chattypes.ChatLinkTypeItem), int64(createItemId), int64(data.StrengthLevel), int64(data.Advance), int64(data.JingLingLevel)}
	infoLink := coreutils.FormatLink(itemName, args)
	// 计算该融合等级属性加成的战力
	power := int64(0)
	if skillId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
		power += int64(skillTemplate.AddPower)
		power += propertylogic.CulculateForce(rongheTemp.GetBattlePropertyMap())
	}
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.RingFuseNotice), plName, infoLink, fuseNum, power)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	fmt.Println("ring： 特戒应该有回复*******************")
	scRingFuse := pbutil.BuildSCRingFuse(success, isBag, int32(typ), index, needIndex, createItemId, createItemNum)
	pl.SendMsg(scRingFuse)

	return
}
