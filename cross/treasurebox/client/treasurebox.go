package client

import (
	"context"
	treasureboxpb "fgame/fgame/cross/treasurebox/pb"

	"google.golang.org/grpc"
)

type TreasureBoxClient interface {
	GetTreasureBoxLogList(ctx context.Context) (resp *treasureboxpb.TreasureBoxLogListResponse, err error)
	OpenTreasureBox(ctx context.Context, req *treasureboxpb.TreasureBoxOpenLogRequest) (resp *treasureboxpb.TreasureBoxOpenLogResponse, err error)
}

type treasureBoxClient struct {
	c      *grpc.ClientConn
	remote treasureboxpb.TreasureBoxClient
}

func (m *treasureBoxClient) OpenTreasureBox(ctx context.Context, req *treasureboxpb.TreasureBoxOpenLogRequest) (resp *treasureboxpb.TreasureBoxOpenLogResponse, err error) {
	resp, err = m.remote.OpenTreasureBox(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *treasureBoxClient) GetTreasureBoxLogList(ctx context.Context) (resp *treasureboxpb.TreasureBoxLogListResponse, err error) {
	req := &treasureboxpb.TreasureBoxLogListRequest{}
	resp, err = m.remote.GetTreasureBoxLogList(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewTreasureBoxClient(conn *grpc.ClientConn) TreasureBoxClient {
	m := &treasureBoxClient{}
	m.c = conn
	m.remote = treasureboxpb.NewTreasureBoxClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
