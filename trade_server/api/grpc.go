package api

import (
	"fgame/fgame/trade_server/pb"
	"fgame/fgame/trade_server/remote"
	"fgame/fgame/trade_server/trade"

	"google.golang.org/grpc"
)

func Server(s *trade.TradeServer, r remote.RemoteService, gopts ...grpc.ServerOption) *grpc.Server {

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
	pb.RegisterTradeManageServer(grpcServer, NewTradeManagerServer(s, r))

	return grpcServer
}
