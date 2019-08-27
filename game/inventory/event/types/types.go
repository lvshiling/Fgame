package types

import (
	commonlog "fgame/fgame/common/log"
)

type InventoryEventType string

const (
	//背包变更
	EventTypeInventoryChanged InventoryEventType = "InventoryChanged"
	//背包变更日志
	EventTypeInventoryChangedLog InventoryEventType = "InventoryChangedLog"
	//物品获得
	EventTypeInventoryItemGet = "InventoryItemGet"
	//物品使用信息改变
	EventTypeItemUseChanged = "ItemUseChanged"
	//物品过期
	EventTypeItemExpire = "ItemExpire"
	//使用物品
	EventTypeUseItem = "UseItem"
	//装备强化等级
	EventTypeEquipmentStrengthenLevel = "EquipmentStrengthenLevel"
	//装备升星
	EventTypeEquipmentUpgradeStar = "EquipmentStrengthUpgradeStar"
	//装备升阶
	EventTypeEquipmentUpgrade = "EquipmentUpgrade"
	//镶嵌宝石
	EventTypeEquipmentEmbedGem = "EquipmentEmbedGem"
	//装备穿上
	EventTypeEquipmentPutOn = "EquipmentPutOn"
	//脱下宝石
	EventTypeEquipmentTakeOffGem = "EquipmentTakeOffGem"
)

//
type PlayerInventoryChangedLogEventData struct {
	itemId        int32
	beforeItemNum int32
	changedNum    int32
	reason        commonlog.InventoryLogReason
	reasonText    string
}

func CreatePlayerInventoryChangedLogEventData(itemId, beforeItemNum, changedNum int32, reason commonlog.InventoryLogReason, reasonText string) *PlayerInventoryChangedLogEventData {
	d := &PlayerInventoryChangedLogEventData{
		itemId:        itemId,
		beforeItemNum: beforeItemNum,
		changedNum:    changedNum,
		reason:        reason,
		reasonText:    reasonText,
	}
	return d
}

func (d *PlayerInventoryChangedLogEventData) GetItemId() int32 {
	return d.itemId
}

func (d *PlayerInventoryChangedLogEventData) GetBeforeItemNum() int32 {
	return d.beforeItemNum
}

func (d *PlayerInventoryChangedLogEventData) GetChangedNum() int32 {
	return d.changedNum
}

func (d *PlayerInventoryChangedLogEventData) GetReason() commonlog.InventoryLogReason {
	return d.reason
}

func (d *PlayerInventoryChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

//
type PlayerInventoryItemUseEventData struct {
	itemId int32
	useNum int32
}

func CreatePlayerInventoryItemUseEventData(itemId, useNum int32) *PlayerInventoryItemUseEventData {
	d := &PlayerInventoryItemUseEventData{
		itemId: itemId,
		useNum: useNum,
	}
	return d
}

func (d *PlayerInventoryItemUseEventData) GetItemId() int32 {
	return d.itemId
}

func (d *PlayerInventoryItemUseEventData) GetUseNum() int32 {
	return d.useNum
}
