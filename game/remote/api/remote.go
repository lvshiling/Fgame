package api

import (
	"context"
	"fgame/fgame/game/remote/cmd"
	remotepb "fgame/fgame/game/remote/pb"

	"github.com/golang/protobuf/proto"
)

//远程服务器
type RemoteServer struct {
}

func (ss *RemoteServer) DoCmd(ctx context.Context, req *remotepb.RemoteCommandRequest) (res *remotepb.RemoteCommandResponse, err error) {
	defer func() {
		if err == nil {
			res = &remotepb.RemoteCommandResponse{}
			return
		}
		switch cerr := err.(type) {
		case cmd.CmdError:
			res = &remotepb.RemoteCommandResponse{}
			res.ErrorCode = int32(cerr.Code())
			res.ErrorMsg = cerr.Code().String()
			err = nil
			return
		}
	}()
	typ := cmd.CmdType(req.GetTyp())
	cmdProto := cmd.GetCmdForType(typ)
	if cmdProto == nil {
		err = cmd.ErrorCodeCommonCmdNoFound
		return
	}
	err = proto.Unmarshal(req.GetCmd(), cmdProto)
	if err != nil {
		return
	}
	err = cmd.HandlerCmd(typ, cmdProto)
	if err != nil {
		return
	}
	return
}

func NewRemoteServer() *RemoteServer {
	ss := &RemoteServer{}
	return ss
}
