package api

import (
	"context"
	"fgame/fgame/center/center"
	centerpb "fgame/fgame/center/pb"
)

//登陆服务器
type LoginManagerServer struct {
	m center.LoginManager
}

func (s *LoginManagerServer) SelfLogin(ctx context.Context, req *centerpb.SelfLoginRequest) (res *centerpb.SelfLoginResponse, err error) {
	res, err = s.m.SelfLogin(ctx, req)
	return
}
func (s *LoginManagerServer) Login(ctx context.Context, req *centerpb.LoginRequest) (res *centerpb.LoginResponse, err error) {
	res, err = s.m.Login(ctx, req)
	return
}

func (s *LoginManagerServer) LoginVerify(ctx context.Context, req *centerpb.LoginVerifyRequest) (res *centerpb.LoginVerifyResponse, err error) {
	res, err = s.m.LoginVerify(ctx, req)
	return
}

func (s *LoginManagerServer) GMLogin(ctx context.Context, req *centerpb.GMLoginRequest) (res *centerpb.GMLoginResponse, err error) {
	res, err = s.m.GMLogin(ctx, req)
	return
}

func NewLoginManagerServer(centerServer *center.CenterServer) *LoginManagerServer {
	ss := &LoginManagerServer{
		m: centerServer,
	}
	return ss
}
