package api

import (
	shenmopb "fgame/fgame/cross/shenmo/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	shenmopb.RegisterShenMoServer(s, NewShenMoServer())
}
