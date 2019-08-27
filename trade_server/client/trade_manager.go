package client

import "context"
import tradeserverpb "fgame/fgame/trade_server/pb"

type TradeManager interface {
	Upload(ctx context.Context, platform int32, serverId int32, tradeId int64, playerId int64, playerName string, itemId int32, itemNum int32, gold int32, propertyData string, level int32) (resp *tradeserverpb.TradeUploadItemResponse, err error)
	// Upload(ctx context.Context, platform int32, serverId int32, tradeId int64) (resp *tradeserverpb.TradeUploadItemResponse, err error)
	Withdraw(ctx context.Context, platform int32, serverId int32, tradeId int64) (resp *tradeserverpb.TradeWithdrawItemResponse, err error)
	Trade(ctx context.Context, platform int32, serverId int32, buyPlayerId int64, buyPlayerName string, globalTradeId int64) (resp *tradeserverpb.TradeItemResponse, err error)
	GetTradeList(ctx context.Context, platform int32, serverId int32) (resp *tradeserverpb.TradeItemListResponse, err error)
}

type tradeManager struct {
	c      *Client
	remote tradeserverpb.TradeManageClient
}

func (m *tradeManager) Upload(ctx context.Context, platform int32, serverId int32, tradeId int64, playerId int64, playerName string, itemId int32, itemNum int32, gold int32, propertyData string, level int32) (resp *tradeserverpb.TradeUploadItemResponse, err error) {
	req := &tradeserverpb.TradeUploadItemRequest{}
	req.Platform = int32(platform)
	req.ServerId = serverId
	req.TradeId = tradeId
	req.PlayerId = playerId
	req.PlayerName = playerName
	req.ItemId = itemId
	req.ItemNum = itemNum
	req.Gold = gold
	req.PropertyData = propertyData
	req.Level = level
	resp, err = m.remote.Upload(ctx, req)
	if err != nil {
		return
	}
	return
}

func (m *tradeManager) Withdraw(ctx context.Context, platform int32, serverId int32, tradeId int64) (resp *tradeserverpb.TradeWithdrawItemResponse, err error) {
	req := &tradeserverpb.TradeWithdrawItemRequest{}
	req.ServerId = serverId
	req.Platform = platform
	req.GlobalTradeId = tradeId
	resp, err = m.remote.Withdraw(ctx, req)
	if err != nil {
		return
	}
	return
}

func (m *tradeManager) Trade(ctx context.Context, platform int32, serverId int32, buyPlayerId int64, buyPlayerName string, globalTradeId int64) (resp *tradeserverpb.TradeItemResponse, err error) {
	req := &tradeserverpb.TradeItemRequest{}
	req.ServerId = serverId
	req.Platform = platform
	req.BuyPlayerId = buyPlayerId
	req.BuyPlayerName = buyPlayerName
	req.GlobalTradeId = globalTradeId
	resp, err = m.remote.Trade(ctx, req)
	if err != nil {
		return
	}
	return
}

func (m *tradeManager) GetTradeList(ctx context.Context, platform int32, serverId int32) (resp *tradeserverpb.TradeItemListResponse, err error) {
	req := &tradeserverpb.TradeItemListRequest{}
	req.ServerId = serverId
	req.Platform = platform
	resp, err = m.remote.GetTradeList(ctx, req)
	if err != nil {
		return
	}
	return
}

func NewTradeManager(c *Client) TradeManager {
	m := &tradeManager{}
	m.c = c
	m.remote = tradeserverpb.NewTradeManageClient(c.conn)
	return m
}
