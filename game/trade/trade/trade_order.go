package trade

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	tradeentity "fgame/fgame/game/trade/entity"
	tradetypes "fgame/fgame/game/trade/types"

	"github.com/pkg/errors"
)

//交易对象
type TradeOrderObject struct {
	id             int64
	serverId       int32
	buyServerId    int32
	playerId       int64
	playerName     string
	tradeId        int64
	itemId         int32
	itemNum        int32
	gold           int32
	propertyData   inventorytypes.ItemPropertyData
	level          int32
	status         tradetypes.TradeOrderStatus
	sellPlatform   int32
	sellServerId   int32
	sellPlayerId   int64
	sellPlayerName string
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func createTradeOrderObject() *TradeOrderObject {
	o := &TradeOrderObject{}
	return o
}

func convertTradeOrderObjectToEntity(o *TradeOrderObject) (*tradeentity.TradeOrderEntity, error) {
	data, err := json.Marshal(o.propertyData)
	if err != nil {
		return nil, err
	}

	e := &tradeentity.TradeOrderEntity{
		Id:             o.id,
		ServerId:       o.serverId,
		BuyServerId:    o.buyServerId,
		ItemId:         o.itemId,
		PlayerId:       o.playerId,
		PlayerName:     o.playerName,
		ItemNum:        o.itemNum,
		Gold:           o.gold,
		TradeId:        o.tradeId,
		PropertyData:   string(data),
		Level:          o.level,
		Status:         int32(o.status),
		SellPlatform:   o.sellPlatform,
		SellServerId:   o.sellServerId,
		SellPlayerId:   o.sellPlayerId,
		SellPlayerName: o.sellPlayerName,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *TradeOrderObject) GetId() int64 {
	return o.id
}

func (o *TradeOrderObject) GetItemId() int32 {
	return o.itemId
}

func (o *TradeOrderObject) GetNum() int32 {
	return o.itemNum
}

func (o *TradeOrderObject) GetGold() int32 {
	return int32(o.gold)
}

func (o *TradeOrderObject) GetTradeId() int64 {
	return o.tradeId
}

func (o *TradeOrderObject) GetPropertyData() inventorytypes.ItemPropertyData {
	return o.propertyData
}

func (o *TradeOrderObject) GetLevel() int32 {
	return o.level
}

func (o *TradeOrderObject) GetDBId() int64 {
	return o.id
}

func (o *TradeOrderObject) GetServerId() int32 {
	return o.serverId
}

func (o *TradeOrderObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *TradeOrderObject) GetPlayerName() string {
	return o.playerName
}

func (o *TradeOrderObject) GetSellServerId() int32 {
	return o.sellServerId
}

func (o *TradeOrderObject) GetSellPlayerId() int64 {
	return o.sellPlayerId
}

func (o *TradeOrderObject) GetSellPlayerName() string {
	return o.sellPlayerName
}

func (o *TradeOrderObject) GetStatus() tradetypes.TradeOrderStatus {
	return o.status
}
func (o *TradeOrderObject) GetTradeTime() int64 {
	return o.updateTime
}

func (o *TradeOrderObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *TradeOrderObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertTradeOrderObjectToEntity(o)
	return e, err
}

func (o *TradeOrderObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*tradeentity.TradeOrderEntity)

	o.id = ae.Id
	o.serverId = ae.ServerId
	o.buyServerId = ae.BuyServerId
	o.playerId = ae.PlayerId
	o.playerName = ae.PlayerName
	o.itemId = ae.ItemId
	o.itemNum = ae.ItemNum
	o.gold = ae.Gold
	o.status = tradetypes.TradeOrderStatus(ae.Status)
	o.sellPlatform = ae.SellPlatform
	o.sellServerId = ae.SellServerId
	o.sellPlayerId = ae.SellPlayerId
	o.sellPlayerName = ae.SellPlayerName
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime

	itemTemp := item.GetItemService().GetItem(int(ae.ItemId))
	data, err := inventory.CreatePropertyData(itemTemp.GetItemType(), ae.PropertyData)
	if err != nil {
		return err
	}
	o.propertyData = data
	o.level = ae.Level
	return nil
}

func (o *TradeOrderObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TradeOrder"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *TradeOrderObject) Refund() bool {
	if o.status != tradetypes.TradeOrderStatusInit {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	o.status = tradetypes.TradeOrderStatusRefund
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *TradeOrderObject) Finish() bool {
	if o.status != tradetypes.TradeOrderStatusInit {
		return false
	}
	o.status = tradetypes.TradeOrderStatusFinish
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *TradeOrderObject) End() bool {
	if o.status != tradetypes.TradeOrderStatusFinish {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	o.status = tradetypes.TradeOrderStatusEnd
	o.updateTime = now
	o.SetModified()
	return true
}
