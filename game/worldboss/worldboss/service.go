package worldboss

import (
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	worldbosstemplate "fgame/fgame/game/worldboss/template"
	"sort"
	"sync"
)

type WorldBossService interface {
	Start()
	//获取世界boss列表
	GetWorldBossList() []scene.NPC
	//获取地图世界boss列表
	GetWorldBossListGroupByMap(mapId int32) []scene.NPC
	//获取世界boss
	GetWorldBoss(biologyId int32) scene.NPC
	//筛选boss
	GetGuaiJiWorldBossList(force int64) []scene.NPC
}

type sortWorldBossList []scene.NPC

func (s sortWorldBossList) Len() int {
	return len(s)
}

func (s sortWorldBossList) Less(i, j int) bool {
	a := worldbosstemplate.GetWorldBossTemplateService().GetWorldBossTemplateByBiologyId(int32(s[i].GetBiologyTemplate().Id))
	b := worldbosstemplate.GetWorldBossTemplateService().GetWorldBossTemplateByBiologyId(int32(s[j].GetBiologyTemplate().Id))
	if a.RecForce == b.RecForce {
		return a.Id < b.Id
	}
	return a.RecForce < b.RecForce
}

func (s sortWorldBossList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type worldBossService struct {
	rwm sync.RWMutex
	//世界boss
	worldBossList []scene.NPC
	//按战斗力排序
	sortWorldBossList []scene.NPC
}

func (s *worldBossService) init() (err error) {
	return
}

func (s *worldBossService) GetWorldBossList() []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.worldBossList
}

func (s *worldBossService) GetWorldBossListGroupByMap(mapId int32) []scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	var bossArr []scene.NPC
	for _, boss := range s.worldBossList {
		if boss.GetScene().MapId() == mapId {
			bossArr = append(bossArr, boss)
		}
	}
	return bossArr
}

func (s *worldBossService) GetWorldBoss(biologyId int32) scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.getBoss(biologyId)
}

func (s *worldBossService) Start() {
	mapIdList := worldbosstemplate.GetWorldBossTemplateService().GetMapIdList()
	for _, mapId := range mapIdList {
		sc := scene.GetSceneService().GetBossSceneByMapId(mapId)
		if sc == nil {
			continue
		}
		//TODO:xzk:修改优化
		bossList := sc.GetNPCS(scenetypes.BiologyScriptTypeWorldBoss)
		for _, boss := range bossList {
			s.worldBossList = append(s.worldBossList, boss)
			s.sortWorldBossList = append(s.sortWorldBossList, boss)
		}
	}
	sort.Sort(sortWorldBossList(s.sortWorldBossList))
	return
}

func (s *worldBossService) getBoss(biologyId int32) (n scene.NPC) {
	for _, boss := range s.worldBossList {
		bossBiologyId := int32(boss.GetBiologyTemplate().TemplateId())
		if bossBiologyId == biologyId {
			return boss
		}
	}
	return nil
}

func (s *worldBossService) GetGuaiJiWorldBossList(force int64) []scene.NPC {
	for index, boss := range s.sortWorldBossList {
		worldbossTemplate := worldbosstemplate.GetWorldBossTemplateService().GetWorldBossTemplateByBiologyId(int32(boss.GetBiologyTemplate().Id))
		if int64(worldbossTemplate.RecForce) > force {
			return s.sortWorldBossList[:index]
		}
	}
	return s.sortWorldBossList
}

var (
	once sync.Once
	ws   *worldBossService
)

func Init() (err error) {
	once.Do(func() {
		ws = &worldBossService{}
		err = ws.init()
	})
	return err
}

func GetWorldBossService() WorldBossService {
	return ws
}
