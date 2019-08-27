package login

import (
	"context"
	"fgame/fgame/account/center/center"
	"fgame/fgame/account/login/types"
	centerclient "fgame/fgame/center/client"
	centerpb "fgame/fgame/center/pb"
)

type LoginService interface {
	SelfLogin(ctx context.Context, name string, password string) (platform int32, platformUserId string, err error)
	Login(ctx context.Context, devicePlatform types.DevicePlatformType, sdkType types.SDKType, platformUserId string, ip string) (userId int64, token string, expiredTime int64, gm int32, err error)
}

type loginService struct {
	client *centerclient.Client
}

func (s *loginService) init() (err error) {
	s.client = center.GetCenterService().GetLoginClient()
	return nil
}

func (s *loginService) Login(ctx context.Context, devicePlatform types.DevicePlatformType, sdkType types.SDKType, platformUserId string, ip string) (userId int64, token string, expiredTime int64, gm int32, err error) {
	req := &centerpb.LoginRequest{}
	req.UserId = platformUserId
	req.Platform = int32(sdkType)
	req.DevicePlatform = int32(devicePlatform)
	req.Ip = ip
	resp, err := s.client.Login(ctx, req)
	if err != nil {
		return
	}
	userId = resp.UserId
	expiredTime = resp.ExpiredTime
	token = resp.Token
	gm = resp.Gm
	return
}

func (s *loginService) SelfLogin(ctx context.Context, name string, password string) (platform int32, platformUserId string, err error) {
	req := &centerpb.SelfLoginRequest{}
	req.Name = name
	req.Password = password
	resp, err := s.client.SelfLogin(ctx, req)
	if err != nil {
		return
	}
	platform = resp.Platform
	platformUserId = resp.PlatformUserId
	return
}

func newLoginService() *loginService {
	s := &loginService{}
	return s
}

var (
	s *loginService
)

func GetLoginService() LoginService {
	return s
}

func Init() error {
	s = newLoginService()
	err := s.init()
	if err != nil {
		return err
	}
	return nil
}
