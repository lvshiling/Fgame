package logic

import (
	"context"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coretypes "fgame/fgame/core/types"
	commomlogic "fgame/fgame/game/common/logic"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	fireworkslogic "fgame/fgame/game/fireworks/logic"
	playerfriend "fgame/fgame/game/friend/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/marry/marry"
	playermarry "fgame/fgame/game/marry/player"
	marryscene "fgame/fgame/game/marry/scene"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//变更结婚属性
func MarryPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeMarry.Mask())
	return
}

//婚车掉落
func HunCheDrop(pos coretypes.Position, owerId int64, sugarType marrytypes.MarryBanquetSubTypeSugar) {
	hunCheNpc := marry.GetMarryService().GetHunCheNpc()
	if hunCheNpc == nil {
		return
	}
	marryConstTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
	moveScene := scene.GetSceneService().GetWorldSceneByMapId(marryConstTemplate.CarMapId)
	banquetTemplate := marrytemplate.GetMarryTemplateService().GetMarryBanquetTeamplate(marrytypes.MarryBanquetTypeSugar, sugarType)
	if banquetTemplate == nil {
		return
	}
	dropList := banquetTemplate.GetDropList()
	scenelogic.CustomDrop(moveScene, hunCheNpc, pos, owerId, dropList, 1)
}

//玩家传送
func PlayerEnterMarry(pl player.Player) (err error) {
	plScene := pl.GetScene()
	sd := marry.GetMarryService().GetMarrySceneData()
	switch sd.Status {
	case marryscene.MarrySceneStatusBanquet:
		{
			marryTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
			marryMapTemplate := marryTemplate.GetMarryMapTemplate()
			marryScene := marry.GetMarryService().GetScene()

			if plScene == marryScene {
				playerlogic.SendSystemMessage(pl, lang.MarryTransferSameScene)
				return
			}
			scenelogic.PlayerEnterScene(pl, marryScene, marryMapTemplate.GetBornPos())
			break
		}
	case marryscene.MarrySceneStatusCruise:
		{
			marryConstTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
			moveScene := scene.GetSceneService().GetWorldSceneByMapId(marryConstTemplate.CarMapId)
			hunCheNpc := marry.GetMarryService().GetHunCheNpc()
			if hunCheNpc == nil {
				return
			}

			pos := hunCheNpc.GetPosition()
			if plScene == moveScene {
				playerlogic.SendSystemMessage(pl, lang.MarryTransferSameScene)
				return
			}
			scenelogic.PlayerEnterScene(pl, moveScene, pos)
			break
		}
	default:
		{
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("marry:目前无正在举办的婚礼")
			playerlogic.SendSystemMessage(pl, lang.MarryHoldWedNoExist)
			return
		}
	}
	return
}

//婚戒返还,修改后直接返回元宝
func PlayerMarryRingGiveBack(pl player.Player, ringType marrytypes.MarryRingType, proposalName string) {
	//template 已校验存在
	// itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, ringType.ItemWedRingSubType())
	// ringItem := int32(itemTemplate.TemplateId())
	// itemMap := make(map[int32]int32)
	// itemMap[ringItem] = 1
	// inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	// flag := inventoryManager.HasEnoughSlot(ringItem, 1)
	// if !flag {
	// 	//发邮件
	// 	emailTitle := lang.GetLangService().ReadLang(lang.MarryRingGiveBackTitle)
	// 	emailContent := lang.GetLangService().ReadLang(lang.MarryRingGiveBackContent)
	// 	emaillogic.AddEmail(pl, emailTitle, emailContent, itemMap)
	// 	return
	// }
	// inventoryReason := commonlog.InventoryLogReasonMarryProposalFail
	// flag = inventoryManager.AddItem(ringItem, 1, inventoryReason, inventoryReason.String())
	// if !flag {
	// 	panic(fmt.Errorf("marry: PlayerMarryRingGiveBack AddItem should be ok"))
	// }

	// inventorylogic.SnapInventoryChanged(pl)
	// ringTypeInt := int32(ringType)
	// merrySubRing := marrytypes.CreateMarryBanquetRingSubType(ringTypeInt) //转换成婚戒子类型配置
	ringItem := marrytemplate.GetMarryTemplateService().GetMarryBanquetTemplateByHouTai(ringType.HoutaiType(), marrytypes.MarryBanquetTypeRing, ringType.BanquetSubTypeRing())
	// itemTempalte := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, ringType.ItemWedRingSubType())
	// if itemTempalte == nil {
	// 	log.WithFields(log.Fields{
	// 		"playerId": pl.GetId(),
	// 		"ringType": ringType,
	// 	}).Warn("marry回退:结婚类型错误")
	// 	return
	// }
	if ringItem == nil {
		log.WithFields(log.Fields{
			"playerId":           pl.GetId(),
			"ringType":           ringType,
			"BanquetSubTypeRing": ringType.BanquetSubTypeRing(),
		}).Warn("marry回退:结婚类型错误")
		return
	}

	// gold := marry.GetMarrySetService().GetCostGold(marrytypes.MarryBanquetTypeRing, ringType.BanquetSubTypeRing())
	buyGold := int64(ringItem.UseGold)
	bindGold := int64(ringItem.UseBinggold)
	silver := int64(ringItem.UseSilver)
	// buyGold := int64(gold) //修改了从多版本的获取
	if buyGold == 0 && bindGold == 0 && silver == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"ringType": ringType,
			"bugGold":  buyGold,
		}).Warn("marry回退:戒指购买金额为0")
		return
	}
	// propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//回退元宝
	// returnText := commonlog.GoldLogReasonMarryProposalReturn.String()
	// propertyManager.AddGold(buyGold, false, commonlog.GoldLogReasonMarryProposalReturn, returnText)
	//发送邮件
	itemMap := make(map[int32]int32)
	if buyGold > 0 {
		itemMap[constanttypes.GoldItem] = int32(buyGold)
	}
	if bindGold > 0 {
		itemMap[constanttypes.BindGoldItem] = int32(bindGold)
	}
	if silver > 0 {
		itemMap[constanttypes.SilverItem] = int32(silver)
	}
	emailTitle := lang.GetLangService().ReadLang(lang.MarryProposalGiveBackTitle)
	emailContent := lang.GetLangService().ReadLang(lang.MarryProposalGiveBackContent)
	emailContent = fmt.Sprintf(emailContent, proposalName)
	emaillogic.AddEmail(pl, emailTitle, emailContent, itemMap)
	//删除数据库求婚数据
	marry.GetMarryService().RemoveMarryRing(pl.GetId())
	return
}

//婚戒培养
func MarryRingFeed(pl player.Player, curTimesNum int32, curBless int32, ringTemplate *gametemplate.MarryRingTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeMarryRing, ringTemplate.TimesMin, ringTemplate.TimesMax)
	updateRate := ringTemplate.UpdateWfb
	blessMax := ringTemplate.NeedRate
	addMin := ringTemplate.AddMin
	addMax := ringTemplate.AddMax + 1

	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//爱情树培养
func MarryTreeFeed(curTimesNum int32, curBless int32, treeTemplate *gametemplate.MarryTreeTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := treeTemplate.TimesMin
	timesMax := treeTemplate.TimesMax
	updateRate := treeTemplate.UpdateWfb
	blessMax := treeTemplate.NeedRate
	addMin := treeTemplate.AddMin
	addMax := treeTemplate.AddMax + 1

	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//玩家赠送贺礼
func PlayerGiveWeddingGift(pl player.Player, giftTemplate *gametemplate.MarryGiftTemplate, autoFlag bool) (isReturn bool) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	giftType := giftTemplate.GetGiftType()
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	switch giftType {
	case marrytypes.MarryGiftTypeItem:
		{
			costGold := int32(0)
			costBindGold := int32(0)
			costSilver := int64(0)
			itemCount := giftTemplate.UseItemAmount
			totalNum := inventoryManager.NumOfItems(giftTemplate.UseItemId)
			if totalNum < itemCount {
				if autoFlag == false {
					isReturn = true
					playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
					return
				}
				//自动
				needBuyNum := itemCount - totalNum
				itemCount = totalNum
				//获取价格
				// shopTemplate := shop.GetShopService().GetShopTemplateByItem(giftTemplate.UseItemId)
				// if shopTemplate == nil {
				// 	isReturn = true
				// 	playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				// 	return
				// }
				// costGold, costBindGold, costSilver = shopTemplate.GetConsumeData(needBuyNum)

				if needBuyNum > 0 {
					if !shop.GetShopService().ShopIsSellItem(giftTemplate.UseItemId) {
						isReturn = true
						playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
						return
					}

					isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, giftTemplate.UseItemId, needBuyNum)
					if !isEnoughBuyTimes {
						isReturn = true
						playerlogic.SendSystemMessage(pl, lang.ShopFlowerAutoBuyItemFail)
						return
					}

					shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
					costGold += int32(shopNeedGold)
					costBindGold += int32(shopNeedBindGold)
					costSilver += shopNeedSilver
				}
			}

			//是否足够元宝
			if costGold != 0 {
				flag := propertyManager.HasEnoughGold(int64(costGold), false)
				if !flag {
					isReturn = true
					playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
					return
				}
			}

			//是否足够绑元
			needCostBindGold := costGold + costBindGold
			if needCostBindGold != 0 {
				flag := propertyManager.HasEnoughGold(int64(needCostBindGold), true)
				if !flag {
					isReturn = true
					playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
					return
				}
			}

			if costSilver != 0 {
				flag := propertyManager.HasEnoughSilver(costSilver)
				if !flag {
					isReturn = true
					playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
					return
				}
			}

			//更新自动购买每日限购次数
			if len(shopIdMap) != 0 {
				shoplogic.ShopDayCountChanged(pl, shopIdMap)
			}

			if totalNum != 0 {
				reasonText := commonlog.InventoryLogReasonMarryWedGift.String()
				flag := inventoryManager.UseItem(giftTemplate.UseItemId, giftTemplate.UseItemAmount, commonlog.InventoryLogReasonMarryWedGift, reasonText)
				if !flag {
					panic(fmt.Errorf("marry: PlayerGiveWeddingGift use item should be ok"))
				}
				//同步物品
				inventorylogic.SnapInventoryChanged(pl)
			}

			//扣除元宝
			if costGold != 0 || costBindGold != 0 || costSilver != 0 {
				//消耗钱
				reasonGoldText := commonlog.GoldLogReasonBuyXianHua.String()
				reasonSliverText := commonlog.SilverLogReasonBuyXianHua.String()
				flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonBuyXianHua, reasonGoldText, 0, commonlog.SilverLogReasonBuyXianHua, reasonSliverText)
				if !flag {
					panic(fmt.Errorf("mount: mountAdvanced Cost should be ok"))
				}
			}

			break
		}
	case marrytypes.MarryGiftTypeSilver:
		{
			costSilver := int64(giftTemplate.UseSilver)
			flag := propertyManager.HasEnoughSilver(costSilver)
			if !flag {
				isReturn = true
				playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
				return
			}

			reasonSliverText := commonlog.SilverLogReasonMarryWedGift.String()
			flag = propertyManager.CostSilver(costSilver, commonlog.SilverLogReasonMarryWedGift, reasonSliverText)
			if !flag {
				panic(fmt.Errorf("marry: PlayerGiveWeddingGift CostSilver should be ok"))
			}
			//同步银两
			propertylogic.SnapChangedProperty(pl)
			break
		}
	case marrytypes.MarryGiftTypeFireworks:
		{
			itemId := giftTemplate.UseItemId
			num := giftTemplate.UseItemAmount
			flag := fireworkslogic.ShootFireworks(pl, itemId, num, false)
			if flag {
				isReturn = true
			}
		}
	}
	return
}

func OnPlayerMarryPushWedHunChe(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl, ok := tpl.(player.Player)
	if !ok {
		return nil
	}
	marryPushWedRelated, ok := result.(*marrytypes.MarryPushWedRelated)
	if !ok {
		return nil
	}
	wedId := marryPushWedRelated.WedId
	msg := marryPushWedRelated.Msg

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.PushWedRecordHunChe(wedId)
	pl.SendMsg(msg)
	return nil
}

func OnPlayerMarryPushWedBanquet(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl, ok := tpl.(player.Player)
	if !ok {
		return nil
	}
	marryPushWedRelated, ok := result.(*marrytypes.MarryPushWedRelated)
	if !ok {
		return nil
	}

	// if playerId == marryPushWedRelated.SpouseId || playerId == marryPushWedRelated.PlayerId {
	// 	return nil
	// }
	msg, successFlag := marryPushWedRelated.Msg.(*uipb.SCMarryWedPushStatus)
	if !successFlag {
		return nil
	}
	//判断是否同一仙盟或者好友
	friendManager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	isShowHunYan := false
	if pl.GetId() != marryPushWedRelated.PlayerId && pl.GetId() != marryPushWedRelated.SpouseId {
		if friendManager.IsFriend(marryPushWedRelated.PlayerId) || friendManager.IsFriend(marryPushWedRelated.SpouseId) {
			isShowHunYan = true
		} else { //统一仙盟
			marryPlayerInfo, _ := player.GetPlayerService().GetPlayerInfo(marryPushWedRelated.PlayerId)
			if marryPlayerInfo != nil && marryPlayerInfo.AllianceId == pl.GetAllianceId() && pl.GetAllianceId() != 0 {
				isShowHunYan = true
			}
			spousePlayerInfo, _ := player.GetPlayerService().GetPlayerInfo(marryPushWedRelated.SpouseId)
			if spousePlayerInfo != nil && spousePlayerInfo.AllianceId == pl.GetAllianceId() && pl.GetAllianceId() != 0 {
				isShowHunYan = true
			}
		}
	}

	msg.IsShow = &isShowHunYan

	wedId := marryPushWedRelated.WedId

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.PushWedRecordBanquet(wedId)
	pl.SendMsg(msg)
	return nil
}

func OnPlayerMarryCeremonyMsg(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl, ok := tpl.(player.Player)
	if !ok {
		return nil
	}
	msgRelated, ok := result.(*marrytypes.MarryPairMsgRelated)
	if !ok {
		return nil
	}
	playerId := msgRelated.PlayerId
	spouseId := msgRelated.SpouseId
	allianceId := msgRelated.AllianceId
	spouseAllianceId := msgRelated.SpouseAllianceId
	showMsg := msgRelated.MsgShow
	msg := msgRelated.Msg
	typ := msgRelated.Type
	switch typ {
	case marrytypes.MarryMsgRelateTypeWedCard,
		marrytypes.MarryMsgRelateTypeMarryCancle:
		{
			if pl.GetId() == playerId || pl.GetId() == spouseId {
				return nil
			}
			break
		}
	}

	manager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	isFriend := manager.IsFriend(playerId)
	isSpouseFriend := manager.IsFriend(spouseId)
	if (pl.GetAllianceId() == allianceId && allianceId != 0) ||
		(pl.GetAllianceId() == spouseAllianceId && spouseAllianceId != 0) ||
		isFriend || isSpouseFriend {
		pl.SendMsg(showMsg)
		return nil
	}

	if typ != marrytypes.MarryMsgRelateTypeWedCard {
		pl.SendMsg(msg)
	}
	return nil
}

//合服 资源返还
func MergeServeGiveBack(marryWed *marry.MarryWedObject) {
	if marryWed == nil {
		return
	}
	playerId := marryWed.PlayerId
	spouseId := marryWed.SpouseId
	grade := marryWed.Grade
	hunCheGrade := marryWed.HunCheGrade
	sugarGrade := marryWed.SugarGrade
	itemMap := MarryPreWedItemMap(grade, hunCheGrade, sugarGrade)
	emailTitle := lang.GetLangService().ReadLang(lang.MarryMergeServerTitle)
	emailContent := lang.GetLangService().ReadLang(lang.MarryMergeServerContent)
	emailContenNotice := lang.GetLangService().ReadLang(lang.MarryMergeServerContentNotice)
	emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, itemMap)
	emaillogic.AddOfflineEmail(spouseId, emailTitle, emailContenNotice, nil)
}

//关服 返还
func CloseServeGiveBack(marryWed *marry.MarryWedObject) {
	if marryWed == nil {
		return
	}
	playerId := marryWed.PlayerId
	spouseId := marryWed.SpouseId
	grade := marryWed.Grade
	hunCheGrade := marryWed.HunCheGrade
	sugarGrade := marryWed.SugarGrade

	itemMap := MarryPreWedItemMap(grade, hunCheGrade, sugarGrade)
	emailTitle := lang.GetLangService().ReadLang(lang.MarryCloseServerTitle)
	emailContent := lang.GetLangService().ReadLang(lang.MarryCloseServerContent)
	emailContenNotice := lang.GetLangService().ReadLang(lang.MarryCloseServerContentNotice)
	emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, itemMap)
	emaillogic.AddOfflineEmail(spouseId, emailTitle, emailContenNotice, nil)
}

func MarryPreWedItemMap(grade marrytypes.MarryBanquetSubTypeWed, hunCheGrade marrytypes.MarryBanquetSubTypeHunChe, sugarGrade marrytypes.MarryBanquetSubTypeSugar) (itemMap map[int32]int32) {
	itemMap = make(map[int32]int32)
	costBindGold, costGold, costSilver := marrytemplate.GetMarryTemplateService().GetMarryGradeCost(grade, hunCheGrade, sugarGrade)
	if costBindGold != 0 {
		itemMap[constanttypes.BindGoldItem] = costBindGold
	}
	if costGold != 0 {
		itemMap[constanttypes.GoldItem] = costGold
	}
	if costSilver != 0 {
		itemMap[constanttypes.SilverItem] = int32(costSilver)
	}
	return
}

//婚礼开始时 队长在3v3匹配 自动离队邮件
func MarryStartAutoLeave(pl player.Player) {
	emailTitle := lang.GetLangService().ReadLang(lang.MarryStartInMatchToLeaveTitle)
	emailContent := lang.GetLangService().ReadLang(lang.MarryStartInMatchToLeaveContent)
	emaillogic.AddEmail(pl, emailTitle, emailContent, nil)
}
