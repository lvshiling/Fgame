package api

import (
	chuangshipb "fgame/fgame/cross/chuangshi/pb"

	"google.golang.org/grpc"
)

func Server(s *grpc.Server) {
	chuangshipb.RegisterChuangshiServer(s, NewChuangShiServer())
}
