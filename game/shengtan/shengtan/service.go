package shengtan

import (
	shengtanscene "fgame/fgame/game/shengtan/scene"
	"sync"
)

type ShengTanService interface {
	GetShengTanScene(allianceId int64) shengtanscene.ShengTanSceneData
	CreateShengTanScene(allianceId int64, mapId int32, endTime int64) shengtanscene.ShengTanSceneData
	ShengTanSceneClose(allianceId int64)
}

type shengTanService struct {
	rwm         sync.RWMutex
	shengTanMap map[int64]shengtanscene.ShengTanSceneData
}

func (s *shengTanService) init() (err error) {
	s.shengTanMap = make(map[int64]shengtanscene.ShengTanSceneData)
	return
}

func (s *shengTanService) GetShengTanScene(allianceId int64) shengtanscene.ShengTanSceneData {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	sc, ok := s.shengTanMap[allianceId]
	if !ok {
		return nil
	}
	return sc
}

func (s *shengTanService) CreateShengTanScene(allianceId int64, mapId int32, endTime int64) shengtanscene.ShengTanSceneData {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	sc, ok := s.shengTanMap[allianceId]
	if ok {
		return sc
	}
	sc = shengtanscene.CreateShengTanSceneData(allianceId, mapId, endTime)
	s.shengTanMap[allianceId] = sc
	return sc
}

func (s *shengTanService) ShengTanSceneClose(allianceId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	_, ok := s.shengTanMap[allianceId]
	if !ok {
		return
	}
	delete(s.shengTanMap, allianceId)
}

func (s *shengTanService) Start() {

}

func (s *shengTanService) Heartbeat() {

}

var (
	once sync.Once
	s    *shengTanService
)

func Init() (err error) {
	once.Do(func() {
		s = &shengTanService{}
		err = s.init()
	})
	return err
}

func GetShengTanService() ShengTanService {
	return s
}
