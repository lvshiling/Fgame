package fourgod

import (
	fourgodscene "fgame/fgame/game/fourgod/scene"
	"sync"
)

type FourGodService interface {
	//创建战场数据
	CreateFourGodSceneData(warId int32, endTime int64) fourgodscene.FourGodSceneData
	//获取四神遗迹数据
	GetFourGodSceneData() fourgodscene.FourGodSceneData
	//四神遗迹活动结束
	FourGodSceneFinish()
}

type fourGodService struct {
	rwm              sync.RWMutex
	fourGodSceneData fourgodscene.FourGodSceneData
}

func (s *fourGodService) init() (err error) {
	return nil
}

func (s *fourGodService) CreateFourGodSceneData(warId int32, endTime int64) (data fourgodscene.FourGodSceneData) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.fourGodSceneData != nil {
		data = s.fourGodSceneData
		return
	}
	fourGodSceneData := fourgodscene.CreateFourGodSceneData(warId, endTime)
	s.fourGodSceneData = fourGodSceneData
	return s.fourGodSceneData
}

func (s *fourGodService) GetFourGodSceneData() fourgodscene.FourGodSceneData {
	return s.fourGodSceneData
}

func (s *fourGodService) FourGodSceneFinish() {
	s.rwm.Lock()
	s.rwm.Unlock()
	s.fourGodSceneData = nil
}

var (
	once sync.Once
	cs   *fourGodService
)

func Init() (err error) {
	once.Do(func() {
		cs = &fourGodService{}
		err = cs.init()
	})
	return err
}

func GetFourGodService() FourGodService {
	return cs
}
