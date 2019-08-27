package client

import (
	"context"
	tulongpb "fgame/fgame/cross/tulong/pb"

	"google.golang.org/grpc"
)

type TuLongClient interface {
	GetTuLongRankList(ctx context.Context) (resp *tulongpb.TuLongRankListResponse, err error)
}

type tuLongClient struct {
	c      *grpc.ClientConn
	remote tulongpb.TuLongClient
}

func (m *tuLongClient) GetTuLongRankList(ctx context.Context) (resp *tulongpb.TuLongRankListResponse, err error) {
	req := &tulongpb.TuLongRankListRequest{}
	resp, err = m.remote.GetTuLongRankList(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewTuLongClient(conn *grpc.ClientConn) TuLongClient {
	m := &tuLongClient{}
	m.c = conn
	m.remote = tulongpb.NewTuLongClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
