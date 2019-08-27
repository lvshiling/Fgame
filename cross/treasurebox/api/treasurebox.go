package api

import (
	"context"
	"fgame/fgame/cross/treasurebox/grpc"
	treasureboxpb "fgame/fgame/cross/treasurebox/pb"
	"fgame/fgame/cross/treasurebox/treasurebox"
)

//服务器服务
type TreasureBoxServer struct {
}

//开跨服宝箱
func (ss *TreasureBoxServer) OpenTreasureBox(ctx context.Context, req *treasureboxpb.TreasureBoxOpenLogRequest) (res *treasureboxpb.TreasureBoxOpenLogResponse, err error) {

	serverId := req.GetServerId()
	playerName := req.GetPlayerName()
	itemList := req.GetItemList()

	treasurebox.GetTreasureBoxService().OpenTreasureBox(serverId, playerName, itemList)
	res = &treasureboxpb.TreasureBoxOpenLogResponse{}
	return
}

//获取跨服宝箱日志
func (ss *TreasureBoxServer) GetTreasureBoxLogList(ctx context.Context, req *treasureboxpb.TreasureBoxLogListRequest) (res *treasureboxpb.TreasureBoxLogListResponse, err error) {
	logList := treasurebox.GetTreasureBoxService().GetLogList()
	logBoxList := grpc_pbutil.BuildTreasureBoxInfoList(logList)

	res = &treasureboxpb.TreasureBoxLogListResponse{}
	res.LogList = logBoxList
	return
}

func NewTreasureBoxServer() *TreasureBoxServer {
	ss := &TreasureBoxServer{}
	return ss
}
