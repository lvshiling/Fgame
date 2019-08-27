package welfarescene

import (
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	welfarescenescene "fgame/fgame/game/welfarescene/scene"
	welfarescenetemplate "fgame/fgame/game/welfarescene/template"
	"sync"
)

type WelfareSceneService interface {
	GetWelfareScene(groupId int32) scene.Scene                                 //
	CreateWelfareScene(groupId int32, tempId int32, endTime int64) scene.Scene //
	WelfareSceneFinish(groupId int32)                                          //奇遇岛活动结束
}

type welfareSceneService struct {
	rwm                 sync.RWMutex
	welfareSceneDataMap map[int32]scene.SceneDelegate
}

func (s *welfareSceneService) init() (err error) {
	s.welfareSceneDataMap = make(map[int32]scene.SceneDelegate)
	return
}

func (s *welfareSceneService) GetWelfareScene(groupId int32) scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	sd := s.getSceneData(groupId)
	if sd == nil {
		return nil
	}

	return sd.GetScene()
}

func (s *welfareSceneService) CreateWelfareScene(groupId int32, tempId int32, endTime int64) scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	sd := s.getSceneData(groupId)
	if sd != nil {
		return sd.GetScene()
	}

	return s.createWelfareScene(groupId, tempId, endTime)
}

func (s *welfareSceneService) WelfareSceneFinish(groupId int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.welfareSceneDataMap[groupId] = nil
}

func (s *welfareSceneService) Start() {
}

func (s *welfareSceneService) Heartbeat() {
}

func (s *welfareSceneService) createWelfareScene(groupId, tempId int32, endTime int64) (sc scene.Scene) {
	wsTemp := welfarescenetemplate.GetWelfareSceneTemplateService().GetWelfareSceneTemplate(tempId)
	if wsTemp == nil {
		return
	}
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(wsTemp.MapId)
	if mapTemplate == nil {
		return
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeWelfareQiYu {
		return
	}

	sd := welfarescenescene.GetWelfareSceneSd(wsTemp.GetWelfareSceneType(), groupId, wsTemp)
	if sd == nil {
		return
	}
	sc = scene.CreateScene(mapTemplate, endTime, sd)
	if sc != nil {
		s.welfareSceneDataMap[groupId] = sd
	}
	return sc
}

func (s *welfareSceneService) getSceneData(groupId int32) scene.SceneDelegate {
	return s.welfareSceneDataMap[groupId]
}

var (
	once sync.Once
	s    *welfareSceneService
)

func Init() (err error) {
	once.Do(func() {
		s = &welfareSceneService{}
		err = s.init()
	})
	return err
}

func GetWelfareSceneService() WelfareSceneService {
	return s
}
