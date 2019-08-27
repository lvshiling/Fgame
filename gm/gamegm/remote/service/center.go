package service

import (
	"context"
	cnclient "fgame/fgame/center/client"
	"net/http"

	remotemodel "fgame/fgame/gm/gamegm/remote/model"

	centerpb "fgame/fgame/center/pb"

	"github.com/codegangsta/negroni"
)

type ICenterService interface {
	Refresh(platform int32) (err error)
	RefreshSDK() (err error)
	RefreshMarryPrice(platform int32) (err error)
	RefreshClientVersion() (err error)
	RefreshPlatformServerConfig() (err error)
	RefreshPlatformConfig(platformId int32) (err error)

	GMLogin(req *remotemodel.CenterGmLoginRequest) (resp *remotemodel.CenterGmLoginRespon, err error)

	ForbidIp(ip string, flag bool) (err error)
	ForbidIpSearch(ip string) (resp *remotemodel.ForbidIpSearch, err error)
}

type centerService struct {
	manager       cnclient.ServerManager
	configManager cnclient.ConfigManager
	loginManager  cnclient.LoginManager
	forbidManager cnclient.ForbidManager
}

func (m *centerService) Refresh(platform int32) (err error) {
	ctx := context.Background()
	_, err = m.manager.Refresh(ctx, platform)
	return
}

func (m *centerService) RefreshSDK() (err error) {
	ctx := context.Background()
	_, err = m.manager.RefreshSDK(ctx)
	return
}

func (m *centerService) RefreshMarryPrice(platform int32) (err error) {
	ctx := context.Background()
	_, err = m.manager.RefreshMarryPrice(ctx, platform)
	return
}

func (m *centerService) RefreshClientVersion() (err error) {
	ctx := context.Background()
	_, err = m.configManager.RefreshClientVersion(ctx)
	return
}

func (m *centerService) RefreshPlatformServerConfig() (err error) {
	ctx := context.Background()
	_, err = m.configManager.RefreshServerConfig(ctx)
	return
}

func (m *centerService) RefreshPlatformConfig(platformId int32) (err error) {
	ctx := context.Background()
	_, err = m.configManager.RefreshPlatformConfig(ctx, platformId)
	return
}

func (m *centerService) GMLogin(req *remotemodel.CenterGmLoginRequest) (resp *remotemodel.CenterGmLoginRespon, err error) {
	ctx := context.Background()
	request := &centerpb.GMLoginRequest{
		SdkType:  req.SdkType,
		UserId:   req.UserId,
		Password: req.Password,
		Name:     req.Name,
	}
	respon, err := m.loginManager.GMLogin(ctx, request)
	if err != nil {
		return
	}
	resp = &remotemodel.CenterGmLoginRespon{
		UserId:      respon.GetUserId(),
		Token:       respon.GetToken(),
		ExpiredTime: respon.GetExpiredTime(),
	}
	return
}

func (m *centerService) ForbidIp(ip string, flag bool) (err error) {
	ctx := context.Background()
	_, err = m.forbidManager.ForbidIp(ctx, ip, flag)
	return
}

func (m *centerService) ForbidIpSearch(ip string) (resp *remotemodel.ForbidIpSearch, err error) {
	ctx := context.Background()
	resp = &remotemodel.ForbidIpSearch{}
	rst, err := m.forbidManager.ForbidIpSearch(ctx, ip)
	if err != nil {
		return
	}
	resp.Result = rst.Forbid
	return
}

func NewCenterService(p_host string, p_port int32) (ICenterService, error) {
	rst := &centerService{}
	config := &cnclient.Config{
		Host: p_host,
		Port: p_port,
	}
	client, err := cnclient.NewClient(config)
	if err != nil {
		return nil, err
	}
	manager := cnclient.NewServerManager(client)
	rst.manager = manager
	rst.configManager = cnclient.NewConfigManager(client)
	rst.loginManager = cnclient.NewLoginManager(client)
	rst.forbidManager = cnclient.NewForbidManager(client)
	return rst, nil
}

const (
	centerServiceKey = contextKey("CenterService")
)

func WithCenterService(ctx context.Context, ls ICenterService) context.Context {
	return context.WithValue(ctx, centerServiceKey, ls)
}

func CenterServiceInContext(ctx context.Context) ICenterService {
	us, ok := ctx.Value(centerServiceKey).(ICenterService)
	if !ok {
		return nil
	}
	return us
}

func SetupCenterServiceHandler(ls ICenterService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithCenterService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
