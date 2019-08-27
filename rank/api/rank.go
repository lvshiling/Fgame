package api

import (
	"context"
	"fgame/fgame/rank/grpc"
	rankpb "fgame/fgame/rank/protocol/pb"
	"fgame/fgame/rank/rank/rank"
)

//服务器服务
type RankServer struct {
}

//刷新排行榜
func (ss *RankServer) RefreshRank(ctx context.Context, req *rankpb.RankRefreshRequest) (res *rankpb.RankRefreshResponse, err error) {
	rank.GMRankUpdate()
	res = &rankpb.RankRefreshResponse{}
	return
}

//获取排行榜
func (ss *RankServer) GetRankList(ctx context.Context, req *rankpb.RankListRequest) (res *rankpb.RankListResponse, err error) {
	rankMap := rank.GetRankService().GetRankMap()
	forceInfoList := grpc_pbutil.BuildForceInfoList(rankMap)
	gangInfoList := grpc_pbutil.BuildGangInfoList(rankMap)
	mountInfoList := grpc_pbutil.BuildMountInfoList(rankMap)
	wingInfoList := grpc_pbutil.BuildWingInfoList(rankMap)
	weaponInfoList := grpc_pbutil.BuildWeaponInfoList(rankMap)
	bodyShieldInfoList := grpc_pbutil.BuildBodyShieldInfoList(rankMap)
	shenFaInfoList := grpc_pbutil.BuildShenFaInfoList(rankMap)
	lingYuInfoList := grpc_pbutil.BuildLingYuInfoList(rankMap)
	featherInfoList := grpc_pbutil.BuildFeatherInfoList(rankMap)
	shieldInfoList := grpc_pbutil.BuildShieldInfoList(rankMap)
	anQiInfoList := grpc_pbutil.BuildAnQiInfoList(rankMap)

	res = &rankpb.RankListResponse{}
	res.Force = forceInfoList
	res.Gang = gangInfoList
	res.Mount = mountInfoList
	res.Wing = wingInfoList
	res.Weapon = weaponInfoList
	res.BodyShield = bodyShieldInfoList
	res.ShenFa = shenFaInfoList
	res.LingYu = lingYuInfoList
	res.Feather = featherInfoList
	res.Shield = shieldInfoList
	res.AnQi = anQiInfoList
	return
}

func NewRankServer() *RankServer {
	ss := &RankServer{}
	return ss
}
