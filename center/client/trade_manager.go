package client

import "context"
import centerpb "fgame/fgame/center/pb"

type TradeManager interface {
	GetTradeServerList(ctx context.Context) (resp *centerpb.TradeServerListResponse, err error)
}

type tradeManager struct {
	c      *Client
	remote centerpb.TradeServerManageClient
}

func (m *tradeManager) GetTradeServerList(ctx context.Context) (resp *centerpb.TradeServerListResponse, err error) {
	req := &centerpb.TradeServerListRequest{}

	resp, err = m.remote.GetTradeServerList(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewTradeManager(c *Client) TradeManager {
	m := &tradeManager{}
	m.c = c
	m.remote = centerpb.NewTradeServerManageClient(c.conn)
	return m
}
