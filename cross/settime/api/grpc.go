package api

import (
	settimepb "fgame/fgame/cross/settime/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	settimepb.RegisterSetTimeServer(s, NewSetTimeServer())
}
