package client

import "context"
import centerpb "fgame/fgame/center/pb"

type LoginManager interface {
	SelfLogin(ctx context.Context, req *centerpb.SelfLoginRequest) (resp *centerpb.SelfLoginResponse, err error)
	Login(ctx context.Context, req *centerpb.LoginRequest) (resp *centerpb.LoginResponse, err error)
	LoginVerify(ctx context.Context, req *centerpb.LoginVerifyRequest) (resp *centerpb.LoginVerifyResponse, err error)
	GMLogin(ctx context.Context, req *centerpb.GMLoginRequest) (resp *centerpb.GMLoginResponse, err error)
}

type loginManager struct {
	c      *Client
	remote centerpb.LoginManageClient
}

func (m *loginManager) SelfLogin(ctx context.Context, req *centerpb.SelfLoginRequest) (resp *centerpb.SelfLoginResponse, err error) {
	resp, err = m.remote.SelfLogin(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *loginManager) Login(ctx context.Context, req *centerpb.LoginRequest) (resp *centerpb.LoginResponse, err error) {
	resp, err = m.remote.Login(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *loginManager) LoginVerify(ctx context.Context, req *centerpb.LoginVerifyRequest) (resp *centerpb.LoginVerifyResponse, err error) {
	resp, err = m.remote.LoginVerify(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *loginManager) GMLogin(ctx context.Context, req *centerpb.GMLoginRequest) (resp *centerpb.GMLoginResponse, err error) {
	resp, err = m.remote.GMLogin(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewLoginManager(c *Client) LoginManager {
	m := &loginManager{}
	m.c = c
	m.remote = centerpb.NewLoginManageClient(c.conn)
	return m
}
