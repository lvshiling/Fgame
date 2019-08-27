package api

import (
	"context"
	grpc_pbutil "fgame/fgame/cross/shenmo/grpc"
	shenmopb "fgame/fgame/cross/shenmo/pb"
	"fgame/fgame/cross/shenmo/shenmo"
)

//服务器服务
type ShenMoServer struct {
}

func (ts *ShenMoServer) GetShenMoRankList(ctx context.Context, req *shenmopb.ShenMoRankListRequest) (res *shenmopb.ShenMoRankListResponse, err error) {
	lastTime, lastRankList := shenmo.GetShenMoService().GetLastRankList()
	thisTime, thisRankList := shenmo.GetShenMoService().GetThisRankList()

	thisRankData, lastRankData := grpc_pbutil.BuildShenMoRankInfoList(thisRankList, thisTime, lastRankList, lastTime)
	res = &shenmopb.ShenMoRankListResponse{}
	res.LastRankData = lastRankData
	res.ThisRankData = thisRankData
	return
}

func NewShenMoServer() *ShenMoServer {
	ss := &ShenMoServer{}
	return ss
}
