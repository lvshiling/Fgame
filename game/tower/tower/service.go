package tower

import (
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	towertemplate "fgame/fgame/game/tower/template"
	"fmt"

	"sync"
)

//打宝塔
type TowerService interface {
	Star()
	//获取打宝塔boss列表
	GetTowerBossList() map[int32]scene.NPC
	// 获取boss信息
	GetTowerBoss(floor int32) scene.NPC
}

type towerService struct {
	rwm sync.RWMutex
	//打宝塔boss列表
	bossMap map[int32]scene.NPC
}

//初始化
func (s *towerService) init() (err error) {
	s.bossMap = make(map[int32]scene.NPC)
	return
}

func (s *towerService) Star() {
	// 加载场景
	s.initTowerScene()

	// 加载打宝塔boss
	s.loadTowerBoss()
}

func (s *towerService) loadTowerBoss() {
	sList := scene.GetSceneService().GetAllTowerScene()
	for _, sc := range sList {
		//TODO:xzk:修改优化
		sd := sc.SceneDelegate().(TowerSceneData)
		towerTemplate := sd.GetTowerTemplate()
		if towerTemplate.BossId == 0 {
			continue
		}

		floor := int32(towerTemplate.TemplateId())
		npcList := sc.GetNPCS(scenetypes.BiologyScriptTypeTowerBoss)
		if len(npcList) != 1 {
			panic(fmt.Errorf("tower:打宝塔BOSS应该有一只"))
		}
		for _, npc := range npcList {
			s.bossMap[floor] = npc
		}
	}
}

func (s *towerService) initTowerScene() {
	mapIdMap := towertemplate.GetTowerTemplateService().GetTowerMapIdMap()
	for mapId, towerTemp := range mapIdMap {
		sd := CreateTowerSceneData(towerTemp)
		sc := scene.CreateTowerScene(mapId, sd)
		if sc == nil {
			panic(fmt.Errorf("tower:初始化场景应该成功"))
		}
	}
}

func (s *towerService) GetTowerBossList() map[int32]scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.bossMap
}

func (s *towerService) GetTowerBoss(floor int32) scene.NPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	for curFloor, npc := range s.bossMap {
		if curFloor != floor {
			continue
		}
		return npc
	}

	return nil
}

var (
	once sync.Once
	cs   *towerService
)

func Init() (err error) {
	once.Do(func() {
		cs = &towerService{}
		err = cs.init()
	})
	return err
}

func GetTowerService() TowerService {
	return cs
}
