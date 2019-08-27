package api

import (
	rankpb "fgame/fgame/rank/protocol/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	rankpb.RegisterRankServer(s, NewRankServer())
}
