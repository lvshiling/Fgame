package client

import (
	"context"
	sharebosspb "fgame/fgame/cross/shareboss/pb"

	"google.golang.org/grpc"
)

type ShareBossClient interface {
	GetShareBossList(ctx context.Context, bossType int32) (resp *sharebosspb.ShareBossInfoListResponse, err error)
}

type shareBossClient struct {
	c      *grpc.ClientConn
	remote sharebosspb.ShareBossClient
}

func (m *shareBossClient) GetShareBossList(ctx context.Context, bossType int32) (resp *sharebosspb.ShareBossInfoListResponse, err error) {
	req := &sharebosspb.ShareBossInfoListRequest{}
	req.BossType = bossType
	resp, err = m.remote.GetAllShareBossList(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewShareBossClient(conn *grpc.ClientConn) ShareBossClient {
	m := &shareBossClient{}
	m.c = conn
	m.remote = sharebosspb.NewShareBossClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
