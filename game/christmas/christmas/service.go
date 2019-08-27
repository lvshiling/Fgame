package christmas

import (
	"fgame/fgame/core/heartbeat"
	christmaseventtypes "fgame/fgame/game/christmas/event/types"
	christmastemplate "fgame/fgame/game/christmas/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	welfarelogic "fgame/fgame/game/welfare/logic"

	"sync"
)

//圣诞采集物
type CharistmasService interface {
	Heartbeat()
	Star()
}

type christmasService struct {
	rwm          sync.RWMutex
	freshTimeMap map[int32]int64
	hbRunner     heartbeat.HeartbeatTaskRunner
	endGroupMap  map[int32]int32
}

//初始化
func (s *christmasService) init() (err error) {
	s.endGroupMap = make(map[int32]int32)
	s.freshTimeMap = make(map[int32]int64)
	s.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	s.hbRunner.AddTask(CreateRefreshCollectTask(s))
	return
}

func (s *christmasService) Star() {
	now := global.GetGame().GetTimeService().Now()
	collectTempMap := christmastemplate.GetChristmasTemplateService().GetChristmasTemplateList()
	for group, _ := range collectTempMap {
		if !welfarelogic.IsOnActivityTime(group) {
			continue
		}

		s.freshTimeMap[group] = now
		gameevent.Emit(christmaseventtypes.EventTypeChristmasRefreshCollect, nil, group)
	}

}

func (s *christmasService) Heartbeat() {
	s.hbRunner.Heartbeat()
	return
}

func (s *christmasService) checkRefreshCollect() (err error) {
	now := global.GetGame().GetTimeService().Now()
	collectTempMap := christmastemplate.GetChristmasTemplateService().GetChristmasTemplateList()
	for group, collectTemp := range collectTempMap {
		if !welfarelogic.IsOnActivityTime(group) {
			_, ok := s.endGroupMap[group]
			if !ok {
				gameevent.Emit(christmaseventtypes.EventTypeChristmasStopCollect, nil, group)
				s.endGroupMap[group] = group
			}
			continue
		} else {
			delete(s.endGroupMap, group)
		}

		//刷新点
		lastRefreshTime := s.freshTimeMap[group]
		for _, freshTime := range collectTemp.GetRefreshTimeList(now) {
			if now < freshTime || lastRefreshTime >= freshTime {
				continue
			}
			collectTemp := christmastemplate.GetChristmasTemplateService().GetChristmasTemplate(group)
			if collectTemp == nil {
				continue
			}

			sc := scene.GetSceneService().GetSceneByMapId(collectTemp.RebornMapId)
			if sc == nil {
				continue
			}
			s.freshTimeMap[group] = now
			gameevent.Emit(christmaseventtypes.EventTypeChristmasRefreshCollect, nil, group)
		}

	}
	return
}

var (
	once sync.Once
	cs   *christmasService
)

func Init() (err error) {
	once.Do(func() {
		cs = &christmasService{}
		err = cs.init()
	})
	return err
}

func GetCharistmasService() CharistmasService {
	return cs
}
