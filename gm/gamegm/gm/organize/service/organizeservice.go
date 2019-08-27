package service

import (
	"context"
	centerServermodel "fgame/fgame/gm/gamegm/gm/center/model"
	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"
	platformmodel "fgame/fgame/gm/gamegm/gm/platform/model"
	platform "fgame/fgame/gm/gamegm/gm/platform/service"
	"net/http"

	"github.com/codegangsta/negroni"
)

type IOrganizeService interface {
	//获取服务器列表，p_channelList：gm渠道id，p_platformList：gm平台id，p_serverList：中心服务器主键id
	//result：key sdktype，value：服务器序号列表
	GetSdkServer(p_channelList []int64, p_platformList []int, p_serverList []int) (map[int][]int, error)

	GetSdkServerCount(p_channelList []int64, p_platformList []int, p_serverList []int) (int, error)

	GetSdkList(p_channelList []int64) (sdkList []int, err error)
	GetSdkListByPlatform(p_platformList []int64) (sdkList []int, err error)
}

type organizeService struct {
	serverCenterService centerserver.ICenterServerService
	platformGmService   platform.IPlatformService
}

func (m *organizeService) GetSdkList(p_channelList []int64) (sdkList []int, err error) {
	var platformList []*platformmodel.PlatformInfo
	if len(p_channelList) != 0 {
		platformList, err = m.platformGmService.GetPlatformByChannelArray(p_channelList)
		if err != nil {
			return nil, err
		}
	} else {
		platformList, err = m.platformGmService.GetAllPlatformList()
		if err != nil {
			return nil, err
		}
	}
	for _, platform := range platformList {
		sdkList = append(sdkList, platform.SdkType)
	}
	return sdkList, nil
}

func (m *organizeService) GetSdkListByPlatform(p_platformList []int64) (sdkList []int, err error) {

	platformList, err := m.platformGmService.GetPlatformInfoArray(p_platformList)
	if err != nil {
		return nil, err
	}

	for _, platform := range platformList {
		sdkList = append(sdkList, platform.SdkType)
	}
	return sdkList, nil
}

func (m *organizeService) GetSdkServer(p_channelList []int64, p_platformList []int, p_serverList []int) (map[int][]int, error) {
	serverQueryMap := make(map[int][]int)
	allPlatForm, err := m.platformGmService.GetAllPlatformList()
	if err != nil {
		return nil, err
	}
	allPlatformMap := changePlatformToMap(allPlatForm)
	allCenterPlatformMap := changePlatformToSdkMap(allPlatForm)

	if len(p_platformList) > 0 { //有选平台
		for _, value := range p_platformList {
			if _, ok := allPlatformMap[int64(value)]; !ok {
				continue
			}
			platinfo := allPlatformMap[int64(value)]
			serverQueryMap[platinfo.SdkType] = make([]int, 0)
		}
		if len(p_serverList) == 0 { //如果没有服务器
			return serverQueryMap, nil
		}
	}

	if len(p_serverList) > 0 { //有选服务器
		allServer, err := m.serverCenterService.GetAllCenterServerList()
		if err != nil {
			return nil, err
		}
		serverMap := changeToMap(allServer)
		for _, value := range p_serverList {
			if _, ok := serverMap[int64(value)]; !ok {
				continue
			}
			serverInfo := serverMap[int64(value)]
			if _, cenOk := allCenterPlatformMap[serverInfo.Platform]; !cenOk { //异常的中心平台
				continue
			}
			sdkTypeArray := allCenterPlatformMap[serverInfo.Platform]
			if len(sdkTypeArray) == 0 {
				continue
			}
			for _, skdValue := range sdkTypeArray {
				if _, quOk := serverQueryMap[skdValue]; quOk { //有在查询的sdk列表中
					serverQueryMap[skdValue] = append(serverQueryMap[skdValue], serverInfo.ServerId)
				}
			}
		}
		noServerPlat := make([]int, 0)
		for key, value := range serverQueryMap {
			if len(value) == 0 {
				noServerPlat = append(noServerPlat, key)
			}
		}
		for _, noid := range noServerPlat {
			delete(serverQueryMap, noid)
		}
		return serverQueryMap, nil
	}

	if len(p_channelList) > 0 {
		platformList, err := m.platformGmService.GetPlatformByChannelArray(p_channelList)
		if err != nil {
			return serverQueryMap, nil
		}
		for _, value := range platformList {
			serverQueryMap[value.SdkType] = make([]int, 0)
		}
	}
	return serverQueryMap, nil
}

func (m *organizeService) GetSdkServerCount(p_channelList []int64, p_platformList []int, p_serverList []int) (int, error) {
	if len(p_serverList) != 0 {
		return len(p_serverList), nil
	}
	if len(p_platformList) != 0 {
		tempPlatformList := make([]int64, 0)
		for _, value := range p_platformList {
			tempPlatformList = append(tempPlatformList, int64(value))
		}
		platformList, err := m.platformGmService.GetPlatformInfoArray(tempPlatformList)
		if err != nil {
			return 0, err
		}
		centerPlatformList := make([]int, 0)
		for _, value := range platformList {
			centerPlatformList = append(centerPlatformList, int(value.CenterPlatformID))
		}
		serverList, err := m.serverCenterService.GetCenterServerListByPlatformArray(centerPlatformList)
		if err != nil {
			return 0, nil
		}
		return len(serverList), nil
	}
	if len(p_channelList) != 0 {
		platformList, err := m.platformGmService.GetPlatformByChannelArray(p_channelList)
		if err != nil {
			return 0, err
		}
		centerPlatformList := make([]int, 0)
		for _, value := range platformList {
			centerPlatformList = append(centerPlatformList, int(value.CenterPlatformID))
		}
		serverList, err := m.serverCenterService.GetCenterServerListByPlatformArray(centerPlatformList)
		if err != nil {
			return 0, nil
		}
		return len(serverList), nil
	}

	return 0, nil
}

func changeToMap(p_list []*centerServermodel.CenterServer) map[int64]*centerServermodel.CenterServer {
	rst := make(map[int64]*centerServermodel.CenterServer)
	for _, value := range p_list {
		if value.ServerType != 0 {
			continue
		}
		rst[value.Id] = value
	}
	return rst
}

func changePlatformToMap(p_list []*platformmodel.PlatformInfo) map[int64]*platformmodel.PlatformInfo {
	rst := make(map[int64]*platformmodel.PlatformInfo)
	for _, value := range p_list {
		rst[value.PlatformID] = value
	}
	return rst
}

func changePlatformToSdkMap(p_list []*platformmodel.PlatformInfo) map[int64][]int {
	rst := make(map[int64][]int)
	for _, value := range p_list {
		if _, ok := rst[value.CenterPlatformID]; !ok {
			rst[value.CenterPlatformID] = make([]int, 0)
		}
		rst[value.CenterPlatformID] = append(rst[value.CenterPlatformID], value.SdkType)
	}
	return rst
}

func NewOrganizeService(p_centerServer centerserver.ICenterServerService, p_platform platform.IPlatformService) IOrganizeService {
	rst := &organizeService{
		serverCenterService: p_centerServer,
		platformGmService:   p_platform,
	}
	return rst
}

type contextKey string

const (
	organizeServiceKey = contextKey("OrganizeService")
)

func WithOrganizeService(ctx context.Context, ls IOrganizeService) context.Context {
	return context.WithValue(ctx, organizeServiceKey, ls)
}

func OrganizeServiceInContext(ctx context.Context) IOrganizeService {
	us, ok := ctx.Value(organizeServiceKey).(IOrganizeService)
	if !ok {
		return nil
	}
	return us
}

func SetupOrganizeServiceHandler(ls IOrganizeService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithOrganizeService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
