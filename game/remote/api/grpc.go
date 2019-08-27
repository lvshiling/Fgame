package api

import (
	remotepb "fgame/fgame/game/remote/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	remotepb.RegisterRemoteServer(s, NewRemoteServer())
}
