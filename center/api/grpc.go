package api

import (
	"fgame/fgame/center/center"
	centerpb "fgame/fgame/center/pb"

	"google.golang.org/grpc"
)

func Server(s *center.CenterServer, gopts ...grpc.ServerOption) *grpc.Server {

	var opts []grpc.ServerOption
	// opts = append(opts, grpc.CustomCodec(&codec{}))
	// if tls != nil {
	// 	opts = append(opts, grpc.Creds(credentials.NewTLS(tls)))
	// }
	// opts = append(opts, grpc.UnaryInterceptor(newUnaryInterceptor(s)))
	// opts = append(opts, grpc.StreamInterceptor(newStreamInterceptor(s)))
	// opts = append(opts, grpc.MaxRecvMsgSize(int(s.Cfg.MaxRequestBytes+grpcOverheadBytes)))
	// opts = append(opts, grpc.MaxSendMsgSize(maxSendBytes))
	// opts = append(opts, grpc.MaxConcurrentStreams(maxStreams))
	grpcServer := grpc.NewServer(append(opts, gopts...)...)
	centerpb.RegisterServerManageServer(grpcServer, NewServerManagerServer(s))
	centerpb.RegisterPlayerManageServer(grpcServer, NewPlayerManagerServer(s))
	centerpb.RegisterLoginManageServer(grpcServer, NewLoginManagerServer(s))
	centerpb.RegisterNoticeManageServer(grpcServer, NewNoticeManagerServer(s))
	centerpb.RegisterConfigManageServer(grpcServer, NewConfigManagerServer(s))
	centerpb.RegisterTradeServerManageServer(grpcServer, NewTradeManagerServer(s))
	centerpb.RegisterForbidManageServer(grpcServer, NewForbidManagerServer(s))

	return grpcServer
}
