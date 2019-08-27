package center

import (
	"context"
	centerclient "fgame/fgame/center/client"
	"fgame/fgame/game/global"
	"time"

	log "github.com/Sirupsen/logrus"
)

type CenterConfig struct {
	Host         string `json:"host"`
	Port         int32  `json:"port"`
	SyncInterval int64  `json:"syncInterval"`
}

type CenterService struct {
	config     *CenterConfig
	client     *centerclient.Client
	syncTicker *time.Ticker
	done       chan struct{}
	serverId   int32
}

func (s *CenterService) init() (err error) {
	//TODO 修改配置
	syncInterValNano := s.config.SyncInterval * int64(time.Millisecond)
	s.syncTicker = time.NewTicker(time.Duration(syncInterValNano))
	s.done = make(chan struct{}, 1)
	cfg := &centerclient.Config{
		Host: s.config.Host,
		Port: s.config.Port,
	}
	s.client, err = centerclient.NewClient(cfg)
	if err != nil {
		return
	}
	//同步注册
	id, err := s.register()
	if err != nil {
		return
	}
	s.serverId = id
	return
}

func (s *CenterService) GetServerId() int32 {
	return s.serverId
}

func (s *CenterService) Start() {
	log.WithFields(
		log.Fields{}).Info("center:中心服服务开启")
	go func() {
	Loop:
		for {
			select {
			case <-s.syncTicker.C:
				s.ping()
				break
			case <-s.done:
				break Loop
			}
		}
		s.unregister()
		log.WithFields(
			log.Fields{}).Info("center:中心服服务结束")
	}()
}

func (s *CenterService) Stop() {
	s.syncTicker.Stop()
	s.done <- struct{}{}
}

func (s *CenterService) register() (serverId int32, err error) {
	//TODO 添加超时机制
	ctx := context.TODO()
	serverType := global.GetGame().GetServerType()
	platform := global.GetGame().GetPlatform()
	serverIndex := global.GetGame().GetServerIndex()
	serverIp := global.GetGame().GetServerIp()
	serverPort := global.GetGame().GetServerPort()

	resp, err := s.client.Register(ctx, serverType, platform, serverIndex, serverIp, serverPort)

	if err != nil {
		return
	}
	log.WithFields(
		log.Fields{
			"serverType":  serverType.String(),
			"platform":    platform,
			"serverIndex": serverIndex,
			"serverIp":    serverIp,
			"serverPort":  serverPort,
		}).Info("center:同步服务器")
	return resp.GetId(), nil
}

func (s *CenterService) ping() {
	//TODO 添加超时机制
	_, err := s.register()
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("center:同步服务器,失败")
	}
	return
}

func (s *CenterService) unregister() {
	//TODO 添加超时机制
	ctx := context.TODO()

	serverId := s.serverId
	_, err := s.client.Unregister(ctx, serverId)

	if err != nil {
		log.WithFields(
			log.Fields{
				"serverId": serverId,
			}).Warn("center:取消注册失败")
		return
	}
	log.WithFields(
		log.Fields{
			"serverId": serverId,
		}).Info("center:取消注册服务器")
	return
}

func NewCenterService(config *CenterConfig) (s *CenterService, err error) {
	s = &CenterService{
		config: config,
	}
	err = s.init()
	if err != nil {
		return
	}
	return
}
