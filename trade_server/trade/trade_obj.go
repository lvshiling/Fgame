package trade

import (
	"fgame/fgame/trade_server/store"
	tradeservertypes "fgame/fgame/trade_server/types"
)

type TradeObject struct {
	id                int64
	platform          int32
	serverId          int32
	tradeId           int64
	playerId          int64
	playerName        string
	propertyData      string
	level             int32
	itemId            int32
	itemNum           int32
	gold              int32
	status            tradeservertypes.TradeStatus
	buyPlayerPlatform int32
	buyPlayerServerId int32
	buyPlayerId       int64
	buyPlayerName     string
	updateTime        int64
	createTime        int64
	deleteTime        int64
}

func NewTradeObject() *TradeObject {
	o := &TradeObject{}
	return o
}

func convertTradeObjectToEntity(o *TradeObject) (*store.TradeItemEntity, error) {

	e := &store.TradeItemEntity{
		Id:                o.id,
		Platform:          o.platform,
		ServerId:          o.serverId,
		TradeId:           o.tradeId,
		PlayerId:          o.playerId,
		PlayerName:        o.playerName,
		PropertyData:      o.propertyData,
		ItemId:            o.itemId,
		ItemNum:           o.itemNum,
		Level:             o.level,
		Gold:              o.gold,
		Status:            int32(o.status),
		BuyPlayerPlatform: o.buyPlayerPlatform,
		BuyPlayerServerId: o.buyPlayerServerId,
		BuyPlayerId:       o.buyPlayerId,
		BuyPlayerName:     o.buyPlayerName,
		UpdateTime:        o.updateTime,
		CreateTime:        o.createTime,
		DeleteTime:        o.deleteTime,
	}
	return e, nil
}

func (o *TradeObject) ToEntity() (e *store.TradeItemEntity, err error) {
	e, err = convertTradeObjectToEntity(o)
	return e, err
}

func (o *TradeObject) FromEntity(e *store.TradeItemEntity) (err error) {

	o.id = e.Id
	o.platform = e.Platform
	o.serverId = e.ServerId
	o.tradeId = e.TradeId
	o.playerId = e.PlayerId
	o.playerName = e.PlayerName
	o.itemId = e.ItemId
	o.itemNum = e.ItemNum
	o.level = e.Level
	o.gold = e.Gold
	o.propertyData = e.PropertyData
	o.status = tradeservertypes.TradeStatus(e.Status)
	o.buyPlayerPlatform = e.BuyPlayerPlatform
	o.buyPlayerServerId = e.BuyPlayerServerId
	o.buyPlayerId = e.BuyPlayerId
	o.buyPlayerName = e.BuyPlayerName
	o.updateTime = e.UpdateTime
	o.createTime = e.CreateTime
	o.deleteTime = e.DeleteTime
	return nil
}

func (o *TradeObject) GetId() int64 {
	return o.id
}

func (o *TradeObject) GetPlatform() int32 {
	return o.platform
}

func (o *TradeObject) GetServerId() int32 {
	return o.serverId
}

func (o *TradeObject) GetTradeId() int64 {
	return o.tradeId
}

func (o *TradeObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *TradeObject) GetPlayerName() string {
	return o.playerName
}

func (o *TradeObject) GetItemId() int32 {
	return o.itemId
}

func (o *TradeObject) GetItemNum() int32 {
	return o.itemNum
}

func (o *TradeObject) GetGold() int32 {
	return o.gold
}

func (o *TradeObject) GetPropertyData() string {
	return o.propertyData
}

func (o *TradeObject) GetLevel() int32 {
	return o.level
}

func (o *TradeObject) GetStatus() tradeservertypes.TradeStatus {
	return o.status
}

func (o *TradeObject) GetBuyPlayerPlatform() int32 {
	return o.buyPlayerPlatform
}

func (o *TradeObject) GetBuyPlayerServerId() int32 {
	return o.buyPlayerServerId
}

func (o *TradeObject) GetBuyPlayerId() int64 {
	return o.buyPlayerId
}

func (o *TradeObject) GetBuyPlayerName() string {
	return o.buyPlayerName
}

func (o *TradeObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *TradeObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *TradeObject) GetDeleteTime() int64 {
	return o.deleteTime
}

func (o *TradeObject) SellNotice(now int64) bool {
	if o.status != tradeservertypes.TradeStatusSell {
		return false
	}
	o.status = tradeservertypes.TradeStatusSellNotice
	o.updateTime = now
	return true
}
