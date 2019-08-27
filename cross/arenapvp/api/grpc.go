package api

import (
	arenapvppb "fgame/fgame/cross/arenapvp/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	arenapvppb.RegisterArenapvpServer(s, NewArenapvpServer())
}
