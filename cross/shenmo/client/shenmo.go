package client

import (
	"context"
	shenmopb "fgame/fgame/cross/shenmo/pb"

	"google.golang.org/grpc"
)

type ShenMoClient interface {
	GetShenMoRankList(ctx context.Context) (resp *shenmopb.ShenMoRankListResponse, err error)
}

type shenMoClient struct {
	c      *grpc.ClientConn
	remote shenmopb.ShenMoClient
}

func (m *shenMoClient) GetShenMoRankList(ctx context.Context) (resp *shenmopb.ShenMoRankListResponse, err error) {
	req := &shenmopb.ShenMoRankListRequest{}
	resp, err = m.remote.GetShenMoRankList(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewShenMoClient(conn *grpc.ClientConn) ShenMoClient {
	m := &shenMoClient{}
	m.c = conn
	m.remote = shenmopb.NewShenMoClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
