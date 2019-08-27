package api

import (
	treasureboxpb "fgame/fgame/cross/treasurebox/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	treasureboxpb.RegisterTreasureBoxServer(s, NewTreasureBoxServer())
}
