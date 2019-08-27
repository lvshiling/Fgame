package api

import (
	"context"
	"fgame/fgame/cross/shareboss/grpc_pbutil"
	sharebosspb "fgame/fgame/cross/shareboss/pb"
	"fgame/fgame/cross/shareboss/shareboss"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

//服务器服务
type ShareBossServer struct {
}

func (ss *ShareBossServer) GetAllShareBossList(ctx context.Context, req *sharebosspb.ShareBossInfoListRequest) (res *sharebosspb.ShareBossInfoListResponse, err error) {
	bossTypeInt := req.BossType
	bossType := worldbosstypes.BossType(bossTypeInt)
	bossList := shareboss.GetShareBossService().GetShareBossList(bossType)
	bossInfoList := grpc_pbutil.BuildBossInfoList(bossList)
	res = &sharebosspb.ShareBossInfoListResponse{}
	res.BossInfoList = bossInfoList
	req.BossType = bossTypeInt
	return
}

func NewShareBossServer() *ShareBossServer {
	ss := &ShareBossServer{}
	return ss
}
