package unrealboss

import (
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	unrealbosstemplate "fgame/fgame/game/unrealboss/template"
	"sort"
	"sync"
)

type UnrealBossService interface {
	Start()
	//获取幻境boss列表
	GetUnrealBossList() []scene.NPC
	//获取地图幻境boss列表
	GetUnrealBossListGroupByMap(mapId int32) []scene.NPC
	//获取幻境boss
	GetUnrealBoss(biologyId int32) scene.NPC
	//筛选boss
	GetGuaiJiUnrealBossList(force int64) []scene.NPC
}

type unrealBossService struct {
	rwm sync.RWMutex
	//幻境boss
	unrealBossList     []scene.NPC
	sortUnrealBossList []scene.NPC
}

type sortUnrealBossList []scene.NPC

func (s sortUnrealBossList) Len() int {
	return len(s)
}

func (s sortUnrealBossList) Less(i, j int) bool {
	a := unrealbosstemplate.GetUnrealBossTemplateService().GetUnrealBossTemplate(int32(s[i].GetBiologyTemplate().Id))
	b := unrealbosstemplate.GetUnrealBossTemplateService().GetUnrealBossTemplate(int32(s[j].GetBiologyTemplate().Id))
	if a.RecForce == b.RecForce {
		return a.Id < b.Id
	}
	return a.RecForce < b.RecForce
}

func (s sortUnrealBossList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *unrealBossService) init() (err error) {
	return
}

func (s *unrealBossService) GetUnrealBossList() []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.unrealBossList
}

func (s *unrealBossService) GetUnrealBossListGroupByMap(mapId int32) []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	var bossArr []scene.NPC
	for _, boss := range s.unrealBossList {
		if boss.GetScene().MapId() == mapId {
			bossArr = append(bossArr, boss)
		}
	}

	return bossArr
}

func (s *unrealBossService) GetUnrealBoss(biologyId int32) scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.getBoss(biologyId)
}

func (s *unrealBossService) Start() {
	mapIdList := unrealbosstemplate.GetUnrealBossTemplateService().GetMapIdList()
	for _, mapId := range mapIdList {
		sc := scene.GetSceneService().GetBossSceneByMapId(mapId)
		if sc == nil {
			continue
		}
		//TODO:xzk:修改优化
		bossList := sc.GetNPCS(scenetypes.BiologyScriptTypeUnrealBoss)
		for _, boss := range bossList {
			s.unrealBossList = append(s.unrealBossList, boss)
			s.sortUnrealBossList = append(s.sortUnrealBossList, boss)
		}
	}
	sort.Sort(sortUnrealBossList(s.sortUnrealBossList))
	return
}

func (s *unrealBossService) getBoss(biologyId int32) (n scene.NPC) {
	for _, boss := range s.unrealBossList {
		bossBiologyId := int32(boss.GetBiologyTemplate().TemplateId())
		if bossBiologyId == biologyId {
			return boss
		}
	}
	return nil
}

func (s *unrealBossService) GetGuaiJiUnrealBossList(force int64) []scene.NPC {
	for index, boss := range s.sortUnrealBossList {
		unrealbossTemplate := unrealbosstemplate.GetUnrealBossTemplateService().GetUnrealBossTemplate(int32(boss.GetBiologyTemplate().Id))
		if int64(unrealbossTemplate.RecForce) > force {
			return s.sortUnrealBossList[:index]
		}
	}
	return s.sortUnrealBossList
}

var (
	once sync.Once
	ws   *unrealBossService
)

func Init() (err error) {
	once.Do(func() {
		ws = &unrealBossService{}
		err = ws.init()
	})
	return err
}

func GetUnrealBossService() UnrealBossService {
	return ws
}
