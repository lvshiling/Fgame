package shenyu

import (
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	shenyuscene "fgame/fgame/game/shenyu/scene"
	shenyutemplate "fgame/fgame/game/shenyu/template"
	"sync"
)

type ShenYuService interface {
	GetShenYuScene() scene.Scene                                            //
	CreateShenYuScene(mapId int32, endTime, sceneEndTime int64) scene.Scene //
	ShenYuSceneFinish()                                                     //神域之战活动结束
}

type shenYuService struct {
	rwm             sync.RWMutex
	shenYuSceneData shenyuscene.ShenYuSceneData
}

func (s *shenYuService) init() (err error) {
	return
}

func (s *shenYuService) GetShenYuScene() scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	if s.shenYuSceneData == nil {
		return nil
	}

	return s.shenYuSceneData.GetScene()
}

func (s *shenYuService) CreateShenYuScene(mapId int32, endTime, sceneEndTime int64) scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	return s.createShenYuScene(mapId, endTime, sceneEndTime)
}

func (s *shenYuService) ShenYuSceneFinish() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.shenYuSceneData = nil
}

func (s *shenYuService) Start() {
}

func (s *shenYuService) Heartbeat() {
}

func (s *shenYuService) createShenYuScene(mapId int32, endTime, sceneEndTime int64) (sc scene.Scene) {

	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeShenYu {
		return nil
	}

	shenYuTemp := shenyutemplate.GetShenYuTemplateService().GetShenYuInitRoundTemplate()
	sd := shenyuscene.CreateShenYuSceneData(shenYuTemp, endTime)
	sc = scene.CreateActivityScene(mapId, sceneEndTime, sd)
	if sc != nil {
		s.shenYuSceneData = sd
	}
	return sc
}

var (
	once sync.Once
	s    *shenYuService
)

func Init() (err error) {
	once.Do(func() {
		s = &shenYuService{}
		err = s.init()
	})
	return err
}

func GetShenYuService() ShenYuService {
	return s
}
