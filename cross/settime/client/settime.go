package client

import (
	"context"
	settimepb "fgame/fgame/cross/settime/pb"

	"google.golang.org/grpc"
)

type SetTimeClient interface {
	SetTime(ctx context.Context, currentTime int64) (resp *settimepb.SetTimeResponse, err error)
}

type setTimeClient struct {
	c      *grpc.ClientConn
	remote settimepb.SetTimeClient
}

func (m *setTimeClient) SetTime(ctx context.Context, currentTime int64) (resp *settimepb.SetTimeResponse, err error) {
	req := &settimepb.SetTimeRequest{}
	req.Time = currentTime
	resp, err = m.remote.SetTime(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewSetTimeClient(conn *grpc.ClientConn) SetTimeClient {
	m := &setTimeClient{}
	m.c = conn
	m.remote = settimepb.NewSetTimeClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
