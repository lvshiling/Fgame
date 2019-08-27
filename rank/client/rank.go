package client

import (
	"context"
	rankpb "fgame/fgame/rank/protocol/pb"

	"google.golang.org/grpc"
)

type RankClient interface {
	GetRankList(ctx context.Context) (resp *rankpb.RankListResponse, err error)
	RefreshRank(ctx context.Context) (resp *rankpb.RankRefreshResponse, err error)
}

type rankClient struct {
	c      *grpc.ClientConn
	remote rankpb.RankClient
}

func (m *rankClient) RefreshRank(ctx context.Context) (resp *rankpb.RankRefreshResponse, err error) {
	req := &rankpb.RankRefreshRequest{}
	resp, err = m.remote.RefreshRank(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *rankClient) GetRankList(ctx context.Context) (resp *rankpb.RankListResponse, err error) {
	req := &rankpb.RankListRequest{}
	resp, err = m.remote.GetRankList(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewRankClient(conn *grpc.ClientConn) RankClient {
	m := &rankClient{}
	m.c = conn
	m.remote = rankpb.NewRankClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
