package center

import (
	centerclient "fgame/fgame/center/client"
	"sync"
)

type CenterConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type CenterService struct {
	rwm          sync.RWMutex
	config       *CenterConfig
	loginClient  *centerclient.Client
	playerClient *centerclient.Client
	serverClient *centerclient.Client
	noticeClient *centerclient.Client
}

func (s *CenterService) init(config *CenterConfig) (err error) {
	s.config = config

	cfg := &centerclient.Config{
		Host: s.config.Host,
		Port: s.config.Port,
	}
	c, err := centerclient.NewClient(cfg)
	if err != nil {
		return
	}
	s.loginClient = c
	s.playerClient = c
	s.serverClient = c
	s.noticeClient = c
	return
}

func (s *CenterService) GetLoginClient() *centerclient.Client {
	return s.loginClient
}

func (s *CenterService) GetPlayerClient() *centerclient.Client {
	return s.playerClient
}

func (s *CenterService) GetServerClient() *centerclient.Client {
	return s.serverClient
}

func (s *CenterService) GetNoticeClient() *centerclient.Client {
	return s.noticeClient
}

func (s *CenterService) Start() {

}

func (s *CenterService) Stop() {

}

var (
	once sync.Once
	cs   *CenterService
)

func Init(config *CenterConfig) (err error) {
	once.Do(func() {
		cs = &CenterService{}
		err = cs.init(config)
	})
	return err
}

func GetCenterService() *CenterService {
	return cs
}
