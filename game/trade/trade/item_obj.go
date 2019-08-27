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
type TradeItemObject struct {
	id             int64
	serverId       int32
	originServerId int32
	playerId       int64
	playerName     string
	itemId         int32
	num            int32
	gold           int32
	propertyData   inventorytypes.ItemPropertyData
	level          int32
	status         tradetypes.TradeStatus
	system         int32
	globalTradeId  int64
	buyPlatform    int32
	buyPlayerId    int64
	buyPlayerName  string
	buyServerId    int32
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func createTradeItemObject() *TradeItemObject {
	o := &TradeItemObject{}
	return o
}

func convertTradeItemObjectToEntity(o *TradeItemObject) (*tradeentity.TradeItemEntity, error) {
	data, err := json.Marshal(o.propertyData)
	if err != nil {
		return nil, err
	}

	e := &tradeentity.TradeItemEntity{
		Id:             o.id,
		ServerId:       o.serverId,
		OriginServerId: o.originServerId,
		ItemId:         o.itemId,
		PlayerId:       o.playerId,
		ItemNum:        o.num,
		Gold:           o.gold,
		PropertyData:   string(data),
		Level:          o.level,
		Status:         int32(o.status),
		System:         o.system,
		GlobalTradeId:  o.globalTradeId,
		BuyPlatform:    o.buyPlatform,
		BuyPlayerId:    o.buyPlayerId,
		BuyPlayerName:  o.buyPlayerName,
		BuyServerId:    o.buyServerId,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *TradeItemObject) GetId() int64 {
	return o.id
}

func (o *TradeItemObject) GetItemId() int32 {
	return o.itemId
}

func (o *TradeItemObject) GetNum() int32 {
	return o.num
}

func (o *TradeItemObject) GetGold() int32 {
	return int32(o.gold)
}

func (o *TradeItemObject) GetPropertyData() inventorytypes.ItemPropertyData {
	return o.propertyData
}
func (o *TradeItemObject) GetLevel() int32 {
	return o.level
}

func (o *TradeItemObject) GetDBId() int64 {
	return o.id
}

func (o *TradeItemObject) GetServerId() int32 {
	return o.serverId
}

func (o *TradeItemObject) GetOriginServerId() int32 {
	return o.originServerId
}

func (o *TradeItemObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *TradeItemObject) GetPlayerName() string {
	return o.playerName
}

func (o *TradeItemObject) GetBuyServerId() int32 {
	return o.buyServerId
}

func (o *TradeItemObject) GetBuyPlayerId() int64 {
	return o.buyPlayerId
}

func (o *TradeItemObject) GetBuyPlayerName() string {
	return o.buyPlayerName
}

func (o *TradeItemObject) GetGlobalTradeId() int64 {
	return o.globalTradeId
}

func (o *TradeItemObject) GetTradeTime() int64 {
	return o.updateTime
}

func (o *TradeItemObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *TradeItemObject) IsSystem() bool {
	return o.system != 0
}

func (o *TradeItemObject) GetStatus() tradetypes.TradeStatus {
	return o.status
}

func (o *TradeItemObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertTradeItemObjectToEntity(o)
	return e, err
}

func (o *TradeItemObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*tradeentity.TradeItemEntity)

	o.id = ae.Id
	o.serverId = ae.ServerId
	o.playerId = ae.PlayerId
	o.itemId = ae.ItemId
	o.num = ae.ItemNum
	o.gold = ae.Gold
	o.status = tradetypes.TradeStatus(ae.Status)
	o.globalTradeId = ae.GlobalTradeId
	o.buyPlatform = ae.BuyPlatform
	o.buyServerId = ae.BuyServerId
	o.buyPlayerId = ae.BuyPlayerId
	o.buyPlayerName = ae.BuyPlayerName
	o.system = ae.System
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

func (o *TradeItemObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TradeItem"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *TradeItemObject) WithDrawing(system bool) bool {
	if o.status != tradetypes.TradeStatusUpload {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = tradetypes.TradeStatusWithDrawing
	if system {
		o.system = 1
	} else {
		o.system = 0
	}
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *TradeItemObject) WithDraw() bool {
	//TODO 验证
	if o.status != tradetypes.TradeStatusWithDrawing {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = tradetypes.TradeStatusWithDraw
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *TradeItemObject) Upload(globalTradeId int64) bool {
	//TODO 验证
	if o.status != tradetypes.TradeStatusInit {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = tradetypes.TradeStatusUpload
	o.globalTradeId = globalTradeId
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *TradeItemObject) Refund() bool {
	//TODO 验证
	if o.status != tradetypes.TradeStatusInit {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = tradetypes.TradeStatusRefund
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *TradeItemObject) Sell(buyPlatform int32, buyServerId int32, buyPlayerId int64, buyPlayerName string) bool {
	//TODO 验证
	if o.status != tradetypes.TradeStatusUpload && o.status != tradetypes.TradeStatusInit && o.status != tradetypes.TradeStatusWithDrawing {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = tradetypes.TradeStatusSold
	o.buyPlatform = buyPlatform
	o.buyServerId = buyServerId
	o.buyPlayerId = buyPlayerId
	o.buyPlayerName = buyPlayerName
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *TradeItemObject) SellNotice() bool {
	//TODO 验证
	if o.status != tradetypes.TradeStatusSold {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = tradetypes.TradeStatusSoldNotice
	o.updateTime = now
	o.SetModified()
	return true
}
