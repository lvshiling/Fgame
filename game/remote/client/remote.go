package client

import (
	"context"
	"fgame/fgame/game/remote/cmd"
	remotepb "fgame/fgame/game/remote/pb"
	"fmt"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type RemoteClient interface {
	DoCmd(ctx context.Context, msg proto.Message) (resp *remotepb.RemoteCommandResponse, err error)
	Close() error
}

type remoteClient struct {
	c      *grpc.ClientConn
	remote remotepb.RemoteClient
}

func (m *remoteClient) DoCmd(ctx context.Context, msg proto.Message) (resp *remotepb.RemoteCommandResponse, err error) {
	req := &remotepb.RemoteCommandRequest{}
	typ, exist := cmd.GetTypeForCmd(msg)
	if !exist {
		err = fmt.Errorf("cmd type no exist")
		return
	}
	req.Typ = int32(typ)
	cmdBytes, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	req.Cmd = cmdBytes
	resp, err = m.remote.DoCmd(ctx, req)
	err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *remoteClient) Close() error {
	return m.c.Close()
}

func NewRemoteClient(conn *grpc.ClientConn) RemoteClient {
	m := &remoteClient{}
	m.c = conn
	m.remote = remotepb.NewRemoteClient(conn)
	return m
}

//TODO 修改
func toErr(ctx context.Context, err error) error {
	return err
}
