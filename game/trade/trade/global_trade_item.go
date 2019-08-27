package trade

import (
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
)

type GlobalTradeItemObject struct {
	id                int64
	tradeId           int64
	platform          int32
	serverId          int32
	playerId          int64
	playerName        string
	itemId            int32
	num               int32
	gold              int32
	propertyData      inventorytypes.ItemPropertyData
	level             int32
	buyPlayerPlatform int32
	buyPlayerServerId int32
	buyPlayerId       int64
	buyPlayerName     string
	updateTime        int64
	createTime        int64
	deleteTime        int64
}

func (o *GlobalTradeItemObject) GetId() int64 {
	return o.id
}
func (o *GlobalTradeItemObject) GetTradeId() int64 {
	return o.tradeId
}

func (o *GlobalTradeItemObject) GetPlatform() int32 {
	return o.platform
}

func (o *GlobalTradeItemObject) GetServerId() int32 {
	return o.serverId
}

func (o *GlobalTradeItemObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *GlobalTradeItemObject) GetPlayerName() string {
	return o.playerName
}

func (o *GlobalTradeItemObject) GetItemId() int32 {
	return o.itemId
}

func (o *GlobalTradeItemObject) GetItemNum() int32 {
	return o.num
}

func (o *GlobalTradeItemObject) GetGold() int32 {
	return o.gold
}

func (o *GlobalTradeItemObject) GetPropertyData() inventorytypes.ItemPropertyData {
	return o.propertyData
}

func (o *GlobalTradeItemObject) GetLevel() int32 {
	return o.level
}

func (o *GlobalTradeItemObject) GetCreateTime() int64 {
	return o.createTime
}

func NewGlobalTradeItemObject(
	id int64,
	tradeId int64,
	platform int32,
	serverId int32,
	playerId int64,
	playerName string,
	itemId int32,
	num int32,
	gold int32,
	propertyData string,
	level int32,
	buyPlayerPlatform int32,
	buyPlayerServerId int32,
	buyPlayerId int64,
	buyPlayerName string,
	updateTime int64,
	createTime int64,
	deleteTime int64) (obj *GlobalTradeItemObject, err error) {
	obj = &GlobalTradeItemObject{}
	obj.id = id
	obj.tradeId = tradeId
	obj.platform = platform
	obj.serverId = serverId
	obj.playerId = playerId
	obj.playerName = playerName
	obj.itemId = itemId
	obj.num = num
	obj.gold = gold
	itemTemplate := item.GetItemService().GetItem(int(obj.itemId))
	if itemTemplate == nil {
		return
	}

	data, err := inventory.CreatePropertyData(itemTemplate.GetItemType(), propertyData)
	if err != nil {
		return
	}

	obj.propertyData = data
	obj.level = level
	obj.buyPlayerPlatform = buyPlayerPlatform
	obj.buyPlayerServerId = buyPlayerServerId
	obj.buyPlayerId = buyPlayerId
	obj.buyPlayerName = buyPlayerName
	obj.updateTime = updateTime
	obj.createTime = createTime
	obj.deleteTime = deleteTime
	return
}
