package api

import (
	tulongpb "fgame/fgame/cross/tulong/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	tulongpb.RegisterTuLongServer(s, NewTuLongServer())
}
