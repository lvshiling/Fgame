package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	tradeentity "fgame/fgame/game/trade/entity"
	tradetypes "fgame/fgame/game/trade/types"

	"github.com/pkg/errors"
)

type PlayerTradeLogObject struct {
	player         player.Player
	id             int64
	logType        tradetypes.TradeLogType
	tradeId        int64
	sellServerId   int32
	sellPlayerId   int64
	sellPlayerName string
	buyServerId    int32
	buyPlayerId    int64
	buyPlayerName  string
	gold           int32
	getGold        int32
	fee            int32
	feeRate        int32
	itemId         int32
	itemNum        int32
	propertyData   inventorytypes.ItemPropertyData
	level          int32
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func createPlayerTradeLogObject(p player.Player) *PlayerTradeLogObject {
	o := &PlayerTradeLogObject{}
	o.player = p
	return o
}

func convertPlayerTradeLogToEntity(o *PlayerTradeLogObject) (*tradeentity.PlayerTradeLogEntity, error) {
	propertyData, err := json.Marshal(o.propertyData)
	if err != nil {
		return nil, err
	}

	e := &tradeentity.PlayerTradeLogEntity{
		Id:             o.id,
		PlayerId:       o.player.GetId(),
		LogType:        int32(o.logType),
		TradeId:        o.tradeId,
		SellServerId:   o.sellServerId,
		SellPlayerId:   o.sellPlayerId,
		SellPlayerName: o.sellPlayerName,
		BuyServerId:    o.buyServerId,
		BuyPlayerId:    o.buyPlayerId,
		BuyPlayerName:  o.buyPlayerName,
		Gold:           o.gold,
		GetGold:        o.getGold,
		Fee:            o.fee,
		FeeRate:        o.feeRate,
		ItemId:         o.itemId,
		ItemNum:        o.itemNum,
		PropertyData:   string(propertyData),
		Level:          o.level,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *PlayerTradeLogObject) GetId() int64 {
	return o.id
}

func (o *PlayerTradeLogObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerTradeLogObject) GetLogType() tradetypes.TradeLogType {
	return o.logType
}

func (o *PlayerTradeLogObject) GetTradeId() int64 {
	return o.tradeId
}
func (o *PlayerTradeLogObject) GetSellServerId() int32 {
	return o.sellServerId
}
func (o *PlayerTradeLogObject) GetSellPlayerId() int64 {
	return o.sellPlayerId
}
func (o *PlayerTradeLogObject) GetSellPlayerName() string {
	return o.sellPlayerName
}

func (o *PlayerTradeLogObject) GetBuyServerId() int32 {
	return o.buyServerId
}

func (o *PlayerTradeLogObject) GetBuyPlayerId() int64 {
	return o.buyPlayerId
}
func (o *PlayerTradeLogObject) GetBuyPlayerName() string {
	return o.buyPlayerName
}

func (o *PlayerTradeLogObject) GetGold() int32 {
	return o.gold
}

func (o *PlayerTradeLogObject) GetGetGold() int32 {
	return o.getGold
}
func (o *PlayerTradeLogObject) GetFee() int32 {
	return o.fee
}

func (o *PlayerTradeLogObject) GetItemId() int32 {
	return o.itemId
}

func (o *PlayerTradeLogObject) GetItemNum() int32 {
	return o.itemNum
}

func (o *PlayerTradeLogObject) GetPropertyData() inventorytypes.ItemPropertyData {
	return o.propertyData
}

func (o *PlayerTradeLogObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerTradeLogObject) GetFeeRate() int32 {
	return o.feeRate
}

func (o *PlayerTradeLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *PlayerTradeLogObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerTradeLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerTradeLogToEntity(o)
	return e, err
}

func (o *PlayerTradeLogObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*tradeentity.PlayerTradeLogEntity)
	o.id = ae.Id
	o.logType = tradetypes.TradeLogType(ae.LogType)

	o.tradeId = ae.TradeId
	o.sellServerId = ae.SellServerId
	o.sellPlayerId = ae.SellPlayerId
	o.sellPlayerName = ae.SellPlayerName
	o.buyServerId = ae.BuyServerId
	o.buyPlayerId = ae.BuyPlayerId
	o.buyPlayerName = ae.BuyPlayerName
	o.gold = ae.Gold
	o.getGold = ae.GetGold
	o.fee = ae.Fee
	o.itemId = ae.ItemId
	o.itemNum = ae.ItemNum
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

func (o *PlayerTradeLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerTradeLog"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}
	o.player.AddChangedObject(obj)
	return
}
