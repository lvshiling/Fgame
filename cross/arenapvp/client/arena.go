package client

import (
	"context"
	arenapvppb "fgame/fgame/cross/arenapvp/pb"

	"google.golang.org/grpc"
)

type ArenapvpClient interface {
	GetArenapvpData(ctx context.Context) (resp *arenapvppb.ArenapvpResponse, err error)
}

type arenapvpClient struct {
	c      *grpc.ClientConn
	remote arenapvppb.ArenapvpClient
}

func (m *arenapvpClient) GetArenapvpData(ctx context.Context) (resp *arenapvppb.ArenapvpResponse, err error) {
	req := &arenapvppb.ArenapvpRequest{}
	resp, err = m.remote.GetArenapvpData(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewArenapvpClient(conn *grpc.ClientConn) ArenapvpClient {
	m := &arenapvpClient{}
	m.c = conn
	m.remote = arenapvppb.NewArenapvpClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
