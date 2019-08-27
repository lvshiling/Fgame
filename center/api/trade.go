package api

import (
	"context"
	"fgame/fgame/center/center"
	centerpb "fgame/fgame/center/pb"
)

//服务器服务
type TradeManagerServer struct {
	m center.TradeManager
}

func (ss *TradeManagerServer) GetTradeServerList(ctx context.Context, req *centerpb.TradeServerListRequest) (res *centerpb.TradeServerListResponse, err error) {
	res, err = ss.m.GetTradeServerList(ctx, req)
	return
}

func NewTradeManagerServer(centerServer *center.CenterServer) *TradeManagerServer {
	ss := &TradeManagerServer{
		m: centerServer,
	}
	return ss
}
