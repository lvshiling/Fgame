package template

import (
	"fgame/fgame/core/template"
	coreutils "fgame/fgame/core/utils"
	gametemplate "fgame/fgame/game/template"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"sync"
)

type ShareBossTemplateService interface {
	//获取跨服世界boss配置
	GetShareBossTemplateByBiologyId(typ worldbosstypes.BossType, biologyId int32) gametemplate.WorldBossTemplateInterface
	GetShareBossTemplateMap(typ worldbosstypes.BossType) map[int32]gametemplate.WorldBossTemplateInterface
	//获取跨服世界boss地图list
	GetMapIdList(typ worldbosstypes.BossType) []int32
}

type shareBossTemplateService struct {
	bossMapOfMap   map[worldbosstypes.BossType]map[int32]gametemplate.WorldBossTemplateInterface
	mapIdListOfMap map[worldbosstypes.BossType][]int32
}

func (s *shareBossTemplateService) init() (err error) {
	s.bossMapOfMap = make(map[worldbosstypes.BossType]map[int32]gametemplate.WorldBossTemplateInterface)
	s.mapIdListOfMap = make(map[worldbosstypes.BossType][]int32)
	err = s.loadShareBoss()
	if err != nil {
		return
	}
	err = s.loadZhenXiBoss()
	if err != nil {
		return
	}
	err = s.loadShengShouBoss()
	if err != nil {
		return
	}
	return
}

func (s *shareBossTemplateService) loadShareBoss() (err error) {
	bossType := worldbosstypes.BossTypeShareBoss
	tempMap := template.GetTemplateService().GetAll((*gametemplate.WorldBossTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.WorldBossTemplate)
		if ftem.GetBossType() != worldbosstypes.WorldBossTypeCross {
			continue
		}
		bossMap, ok := s.bossMapOfMap[bossType]
		if !ok {
			bossMap = make(map[int32]gametemplate.WorldBossTemplateInterface)
			s.bossMapOfMap[bossType] = bossMap
		}
		bossMap[ftem.GetBiologyId()] = ftem

		mapIdList, ok := s.mapIdListOfMap[bossType]
		if !ok {
			mapIdList = make([]int32, 0, 8)
			s.mapIdListOfMap[bossType] = mapIdList
		}
		if !coreutils.ContainInt32(mapIdList, ftem.GetMapId()) {
			mapIdList = append(mapIdList, ftem.GetMapId())
			s.mapIdListOfMap[bossType] = mapIdList
		}

	}
	return nil
}

func (s *shareBossTemplateService) loadZhenXiBoss() (err error) {
	bossType := worldbosstypes.BossTypeZhenXi
	tempMap := template.GetTemplateService().GetAll((*gametemplate.ZhenXiBossTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.ZhenXiBossTemplate)

		bossMap, ok := s.bossMapOfMap[bossType]
		if !ok {
			bossMap = make(map[int32]gametemplate.WorldBossTemplateInterface)
			s.bossMapOfMap[bossType] = bossMap
		}
		bossMap[ftem.GetBiologyId()] = ftem
		mapIdList, ok := s.mapIdListOfMap[bossType]
		if !ok {
			mapIdList = make([]int32, 0, 8)
			s.mapIdListOfMap[bossType] = mapIdList
		}
		if !coreutils.ContainInt32(mapIdList, ftem.GetMapId()) {
			mapIdList = append(mapIdList, ftem.GetMapId())
			s.mapIdListOfMap[bossType] = mapIdList
		}

	}
	return nil
}

func (s *shareBossTemplateService) loadShengShouBoss() (err error) {
	bossType := worldbosstypes.BossTypeArena
	tempMap := template.GetTemplateService().GetAll((*gametemplate.ShengShouBossTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.ShengShouBossTemplate)

		bossMap, ok := s.bossMapOfMap[bossType]
		if !ok {
			bossMap = make(map[int32]gametemplate.WorldBossTemplateInterface)
			s.bossMapOfMap[bossType] = bossMap
		}
		bossMap[ftem.GetBiologyId()] = ftem
		mapIdList, ok := s.mapIdListOfMap[bossType]
		if !ok {
			mapIdList = make([]int32, 0, 8)
			s.mapIdListOfMap[bossType] = mapIdList
		}
		if !coreutils.ContainInt32(mapIdList, ftem.GetMapId()) {
			mapIdList = append(mapIdList, ftem.GetMapId())
			s.mapIdListOfMap[bossType] = mapIdList
		}

	}
	return nil
}

func (s *shareBossTemplateService) GetShareBossTemplateByBiologyId(typ worldbosstypes.BossType, biologyId int32) gametemplate.WorldBossTemplateInterface {
	bossMap, ok := s.bossMapOfMap[typ]
	if !ok {
		return nil
	}
	boss, ok := bossMap[biologyId]
	if !ok {
		return nil
	}
	return boss
}

func (s *shareBossTemplateService) GetShareBossTemplateMap(typ worldbosstypes.BossType) map[int32]gametemplate.WorldBossTemplateInterface {
	bossMap, ok := s.bossMapOfMap[typ]
	if !ok {
		return nil
	}

	return bossMap
}

func (s *shareBossTemplateService) GetMapIdList(typ worldbosstypes.BossType) []int32 {
	mapIdList, ok := s.mapIdListOfMap[typ]
	if !ok {
		return nil
	}
	return mapIdList
}

var (
	once sync.Once
	s    *shareBossTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &shareBossTemplateService{}
		err = s.init()
	})

	return
}

func GetShareBossTemplateService() ShareBossTemplateService {
	return s
}
