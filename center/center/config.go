package center

import (
	"context"
	centerpb "fgame/fgame/center/pb"

	log "github.com/Sirupsen/logrus"
)

type ClientVersion struct {
	iosVersion     string
	androidVersion string
}

type ServerConfig struct {
	tradeIp string
}

func (s *CenterServer) loadClientVersion() (err error) {
	log.WithFields(
		log.Fields{}).Info("server:加载客户端版本号")

	clientVersionEntity, err := s.clientVersionStore.GetClientVersion()
	if err != nil {
		return
	}

	clientVersion := &ClientVersion{}
	if clientVersionEntity != nil {
		clientVersion.iosVersion = clientVersionEntity.IosVersion
		clientVersion.androidVersion = clientVersionEntity.AndroidVersion
	}
	s.clientVersion = clientVersion
	log.WithFields(
		log.Fields{}).Info("server:加载客户端版本号成功")
	return
}

func (s *CenterServer) RefreshClientVersion(ctx context.Context, req *centerpb.ClientVersionRefreshRequest) (res *centerpb.ClientVersionRefreshResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	log.WithFields(
		log.Fields{}).Debug("server:刷新客户端版本")
	err = s.loadClientVersion()
	if err != nil {
		return
	}
	res = &centerpb.ClientVersionRefreshResponse{}
	log.WithFields(
		log.Fields{}).Debug("server:刷新客户端版本成功")
	return
}

func (s *CenterServer) loadServerConfig() (err error) {
	log.WithFields(
		log.Fields{}).Info("server:加载服务器配置")

	sreverConfigEntity, err := s.serverConfigStore.GetServerConfig()
	if err != nil {
		return
	}
	serverConfig := &ServerConfig{}
	if sreverConfigEntity != nil {
		serverConfig.tradeIp = sreverConfigEntity.TradeServerIp

	}

	s.serverConfig = serverConfig
	log.WithFields(
		log.Fields{}).Info("server:加载服务器配置成功")
	return
}

func (s *CenterServer) RefreshServerConfig(ctx context.Context, req *centerpb.ServerConfigRefreshRequest) (res *centerpb.ServerConfigRefreshResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	log.WithFields(
		log.Fields{}).Debug("server:刷新服务器端配置")

	res = &centerpb.ServerConfigRefreshResponse{}
	err = s.loadServerConfig()
	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{}).Debug("server:刷新服务器端配置成功")
	return
}

func (s *CenterServer) loadPlatformSettings() (err error) {
	log.WithFields(
		log.Fields{}).Info("server:加载平台配置")
	s.platformSettingMap = make(map[int32]*PlatformSetting)
	platformSettingEntityList, err := s.platformSettingStore.GetAllPlatformSetting()
	if err != nil {
		return
	}

	for _, platformSettingEntity := range platformSettingEntityList {
		platformSetting := newPlatformSetting()
		err = platformSetting.FromEntity(platformSettingEntity)
		if err != nil {
			return
		}
		s.platformSettingMap[platformSetting.GetPlatform()] = platformSetting
	}

	s.platformChatSetMap = make(map[int32]*PlatformChatset)
	platformChatsetEntityList, err := s.platformChatsetStore.GetAllPlatformChatset()
	if err != nil {
		return
	}

	for _, platformChatsetEntity := range platformChatsetEntityList {
		platformChatset := newPlatformChatset()
		err = platformChatset.FromEntity(platformChatsetEntity)
		if err != nil {
			return
		}
		s.platformChatSetMap[platformChatset.GetPlatform()] = platformChatset
	}

	log.WithFields(
		log.Fields{}).Info("server:加载平台配置成功")
	return
}

func (s *CenterServer) RefreshPlatformConfig(ctx context.Context, req *centerpb.PlatformConfigRefreshRequest) (res *centerpb.PlatformConfigRefreshResponse, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	platform := req.GetPlatform()
	log.WithFields(
		log.Fields{
			"platformId": platform,
		}).Info("server:刷新平台配置")
	platformSettingEntity, err := s.platformSettingStore.GetPlatformSetting(platform)
	if err != nil {
		return
	}
	if platformSettingEntity != nil {
		platformSetting := newPlatformSetting()
		err = platformSetting.FromEntity(platformSettingEntity)
		if err != nil {
			return
		}
		s.platformSettingMap[platformSetting.GetPlatform()] = platformSetting
	}
	platformChatSetEntity, err := s.platformChatsetStore.GetPlatformChatset(platform)
	if err != nil {
		return
	}
	if platformChatSetEntity != nil {
		platformChatset := newPlatformChatset()
		err = platformChatset.FromEntity(platformChatSetEntity)
		if err != nil {
			return
		}
		s.platformChatSetMap[platformChatset.GetPlatform()] = platformChatset
	} else {
		delete(s.platformChatSetMap, platform)
	}

	res = &centerpb.PlatformConfigRefreshResponse{}
	res.Platform = platform
	log.WithFields(
		log.Fields{
			"platformId": platform,
		}).Info("server:刷新平台配置")
	return
}
