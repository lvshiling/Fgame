package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/emperor/emperor"
	emperortemplate "fgame/fgame/game/emperor/template"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
	"math"
)

//帝王属性加成
func EmperorPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeEmperor.Mask())
	return
}

//玩家膜拜奖励
func GiveWorshipReward(pl player.Player) {
	//玩家奖励
	rewData := emperortemplate.GetEmperorTemplateService().GetEmperorWorshipRew()
	if rewData != nil {
		reasonGold := commonlog.GoldLogReasonEmperorWorship
		reasonSilver := commonlog.SilverLogReasonEmperorWorship
		reasonLevel := commonlog.LevelLogReasonEmperorWorship
		reasonGoldText := reasonGold.String()
		reasonSliverText := reasonSilver.String()
		reasonlevelText := reasonLevel.String()
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		flag := propertyManager.AddRewData(rewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("emperor: emperorWorship AddRewData  should be ok"))
		}
		//同步
		propertylogic.SnapChangedProperty(pl)
	}
	return
}

//计算龙椅战力
func CountEmperorPower() map[propertytypes.BattlePropertyType]int64 {
	attrMap := make(map[propertytypes.BattlePropertyType]int64)
	_, robNum := emperor.GetEmperorService().GetEmperorIdAndRobNum()
	weight := emperortemplate.GetEmperorTemplateService().GetEmperorRobCoefficientAttr(robNum)
	dragronChairTemplate := emperortemplate.GetEmperorTemplateService().GetEmperorTemplate()
	firstAttr := dragronChairTemplate.GetFirstAttrTemplate()
	valueAttr := dragronChairTemplate.GetValueAttrTemplate()

	for typ, val := range firstAttr.GetAllBattleProperty() {
		total := int64(math.Ceil(float64(val) * weight))
		_, ok := attrMap[typ]
		if ok {
			attrMap[typ] += total
		} else {
			attrMap[typ] = total
		}
	}

	for typ, val := range valueAttr.GetAllBattleProperty() {
		_, ok := attrMap[typ]
		if ok {
			attrMap[typ] += val
		} else {
			attrMap[typ] = val
		}
	}
	return attrMap
}

func openBoxNotice(pl player.Player, dropItemList []*droptemplate.DropItemData) {
	//公告处理
	isGreaterThree := false
	itemContent := ""
	for _, itemData := range dropItemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()

		itemTemplate := item.GetItemService().GetItem(int(itemId))
		quality := itemtypes.ItemQualityType(itemTemplate.Quality)
		if quality > itemtypes.ItemQualityTypePurple {
			isGreaterThree = true
		}
		itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(num)))
		linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
		itemNameLink := coreutils.FormatLink(itemName, linkArgs)
		itemContent += itemNameLink
	}

	var content string
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	if isGreaterThree {
		content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmperorOpenBoxNoticeGood), playerName, itemContent)
	} else {
		content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmperorOpenBoxNotice), playerName, itemContent)
	}
	//跑马灯
	noticelogic.NoticeNumBroadcast([]byte(content), 0, int32(1))
	//系统频道
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
}

func OpenBoxReward(pl player.Player, dropItemList []*droptemplate.DropItemData, isOpen bool) (err error) {
	if len(dropItemList) == 0 {
		return
	}

	//公告处理
	openBoxNotice(pl, dropItemList)
	newItemList, resMap := droplogic.SeperateItemDatas(dropItemList)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//判断背包是否足够
	if len(newItemList) != 0 {
		if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemList) {
			//写邮件
			var emailTitle string
			var emailContent string
			if isOpen {
				emailTitle = fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmperorOpenBoxTitle))
				emailContent = lang.GetLangService().ReadLang(lang.EmperorOpenBoxContent)
			} else {
				emailTitle = fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmperorRobRewardTitle))
				emailContent = lang.GetLangService().ReadLang(lang.EmperorRobRewardContent)
			}
			now := global.GetGame().GetTimeService().Now()
			emaillogic.AddEmailItemLevel(pl, emailTitle, emailContent, now, newItemList)
		} else {
			var reasonText string
			flag := true
			if isOpen {
				reasonText = commonlog.InventoryLogReasonEmperorOpenBox.String()
				flag = inventoryManager.BatchAddOfItemLevel(newItemList, commonlog.InventoryLogReasonEmperorOpenBox, reasonText)
			} else {
				reasonText = commonlog.InventoryLogReasonEmperorRobReward.String()
				flag = inventoryManager.BatchAddOfItemLevel(newItemList, commonlog.InventoryLogReasonEmperorRobReward, reasonText)
			}
			if !flag {
				panic(fmt.Errorf("emperor: OpenBoxReward BatchAdd should be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//添加掉落资源
	if len(resMap) != 0 {
		reasonGoldText := commonlog.GoldLogReasonEmperorReward.String()
		reasonSliverText := commonlog.SilverLogReasonEmperorReward.String()
		reasonLevelText := commonlog.LevelLogReasonEmperorReward.String()
		err = droplogic.AddRes(pl, resMap, commonlog.GoldLogReasonEmperorReward, reasonGoldText, commonlog.SilverLogReasonEmperorReward, reasonSliverText, commonlog.LevelLogReasonEmperorReward, reasonLevelText)
		if err != nil {
			return
		}
		propertylogic.SnapChangedProperty(pl)
	}
	return
}

//合服 资源返还
func MergeServeGiveBack(playerId int64, itemMap map[int32]int32) {
	if len(itemMap) == 0 {
		return
	}
	emperorTemplate := emperortemplate.GetEmperorTemplateService().GetEmperorTemplate()
	if emperorTemplate == nil {
		return
	}
	percent := float64(emperorTemplate.GoldPercent) / float64(common.MAX_RATE)
	for itemId, num := range itemMap {
		curNum := int32(math.Ceil(float64(num) * percent))
		if curNum <= 0 {
			delete(itemMap, itemId)
		} else {
			itemMap[itemId] = curNum
		}
	}
	if len(itemMap) == 0 {
		return
	}
	percentInt := int32(math.Ceil(percent * float64(100)))
	emailTitle := lang.GetLangService().ReadLang(lang.EmperorMergeServerTitle)
	emailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmperorMergeServerContent), percentInt)
	emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, itemMap)
}
