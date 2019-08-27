package cangjingge

import (
	cangjinggetemplate "fgame/fgame/game/cangjingge/template"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"sort"
	"sync"
)

type CangJingGeService interface {
	Start()
	//获取藏经阁boss列表
	GetCangJingGeBossList() []scene.NPC
	//获取地图藏经阁boss列表
	GetCangJingGeBossListGroupByMap(mapId int32) []scene.NPC
	//获取藏经阁boss
	GetCangJingGeBoss(biologyId int32) scene.NPC
	//筛选boss
	GetGuaiJiCangJingGeBossList(force int64) []scene.NPC
}

type sortCangJingGeBossList []scene.NPC

func (s sortCangJingGeBossList) Len() int {
	return len(s)
}

func (s sortCangJingGeBossList) Less(i, j int) bool {
	a := cangjinggetemplate.GetCangJingGeTemplateService().GetCangJingGeTemplateByBiologyId(int32(s[i].GetBiologyTemplate().Id))
	b := cangjinggetemplate.GetCangJingGeTemplateService().GetCangJingGeTemplateByBiologyId(int32(s[j].GetBiologyTemplate().Id))
	if a.RecForce == b.RecForce {
		return a.Id < b.Id
	}
	return a.RecForce < b.RecForce
}

func (s sortCangJingGeBossList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type cangJingGeService struct {
	rwm sync.RWMutex
	//藏经阁boss
	cangJingGeBossList []scene.NPC
	//按战斗力排序
	sortCangJingGeBossList []scene.NPC
}

func (s *cangJingGeService) init() (err error) {
	return
}

func (s *cangJingGeService) GetCangJingGeBossList() []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.cangJingGeBossList
}

func (s *cangJingGeService) GetCangJingGeBossListGroupByMap(mapId int32) []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	var bossArr []scene.NPC
	for _, boss := range s.cangJingGeBossList {
		if boss.GetScene().MapId() == mapId {
			bossArr = append(bossArr, boss)
		}
	}
	return bossArr
}

func (s *cangJingGeService) GetCangJingGeBoss(biologyId int32) scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.getBoss(biologyId)
}

func (s *cangJingGeService) Start() {
	mapIdList := cangjinggetemplate.GetCangJingGeTemplateService().GetMapIdList()
	for _, mapId := range mapIdList {
		sc := scene.GetSceneService().GetBossSceneByMapId(mapId)
		if sc == nil {
			continue
		}
		//TODO:xzk:修改优化
		bossList := sc.GetNPCS(scenetypes.BiologyScriptTypeCangJingGeBoss)
		for _, boss := range bossList {
			s.cangJingGeBossList = append(s.cangJingGeBossList, boss)
			s.sortCangJingGeBossList = append(s.sortCangJingGeBossList, boss)
		}
	}
	sort.Sort(sortCangJingGeBossList(s.sortCangJingGeBossList))
	return
}

func (s *cangJingGeService) getBoss(biologyId int32) (n scene.NPC) {
	for _, boss := range s.cangJingGeBossList {
		bossBiologyId := int32(boss.GetBiologyTemplate().TemplateId())
		if bossBiologyId == biologyId {
			return boss
		}
	}
	return nil
}

func (s *cangJingGeService) GetGuaiJiCangJingGeBossList(force int64) []scene.NPC {
	for index, boss := range s.sortCangJingGeBossList {
		template := cangjinggetemplate.GetCangJingGeTemplateService().GetCangJingGeTemplateByBiologyId(int32(boss.GetBiologyTemplate().Id))
		if int64(template.RecForce) > force {
			return s.sortCangJingGeBossList[:index]
		}
	}
	return s.sortCangJingGeBossList
}

var (
	once sync.Once
	ws   *cangJingGeService
)

func Init() (err error) {
	once.Do(func() {
		ws = &cangJingGeService{}
		err = ws.init()
	})
	return err
}

func GetCangJingGeService() CangJingGeService {
	return ws
}
