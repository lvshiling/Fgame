package api

import (
	"context"
	"fgame/fgame/cross/arenapvp/arenapvp"
	grpc_pbutil "fgame/fgame/cross/arenapvp/grpc"
	arenapvppb "fgame/fgame/cross/arenapvp/pb"
)

//服务器服务
type ArenapvpServer struct {
}

func (ts *ArenapvpServer) GetArenapvpData(ctx context.Context, req *arenapvppb.ArenapvpRequest) (res *arenapvppb.ArenapvpResponse, err error) {
	battleMap := arenapvp.GetArenapvpService().GetPvpResultList()
	baZhuList := arenapvp.GetArenapvpService().GetBaZhuList()
	guessDataList := arenapvp.GetArenapvpService().GetGuessDataList()
	electionSceneMap := arenapvp.GetArenapvpService().GetAllElectionSceneMap()

	pvpData := grpc_pbutil.BuildArenapvpData(battleMap, baZhuList, guessDataList, electionSceneMap)
	res = &arenapvppb.ArenapvpResponse{}
	res.ArenapvpData = pvpData
	return
}

func NewArenapvpServer() *ArenapvpServer {
	ss := &ArenapvpServer{}
	return ss
}
