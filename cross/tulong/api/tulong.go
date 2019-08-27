package api

import (
	"context"
	"fgame/fgame/cross/tulong/grpc"
	tulongpb "fgame/fgame/cross/tulong/pb"
	"fgame/fgame/cross/tulong/tulong"
)

//服务器服务
type TuLongServer struct {
}

func (ts *TuLongServer) GetTuLongRankList(ctx context.Context, req *tulongpb.TuLongRankListRequest) (res *tulongpb.TuLongRankListResponse, err error) {
	rankList := tulong.GetTuLongService().GetRankList()
	rankInfoList := grpc_pbutil.BuildTuLongRankInfoList(rankList)
	res = &tulongpb.TuLongRankListResponse{}
	res.RankInfoList = rankInfoList
	return
}

func NewTuLongServer() *TuLongServer {
	ss := &TuLongServer{}
	return ss
}
