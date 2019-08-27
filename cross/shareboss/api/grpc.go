package api

import (
	sharebosspb "fgame/fgame/cross/shareboss/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	sharebosspb.RegisterShareBossServer(s, NewShareBossServer())
}
