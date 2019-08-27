package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"sync"
)

type WorldBossTemplateService interface {
	//获取世界boss配置
	GetWorldBossTemplateByBiologyId(biologyId int32) *gametemplate.WorldBossTemplate
	//获取世界boss地图list
	GetMapIdList() []int32
}

type worldBossTemplateService struct {
	bossMap   map[int32]*gametemplate.WorldBossTemplate
	mapIdList []int32
}

func (s *worldBossTemplateService) init() (err error) {
	s.bossMap = make(map[int32]*gametemplate.WorldBossTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.WorldBossTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.WorldBossTemplate)
		if ftem.GetBossType() != worldbosstypes.WorldBossTypeLocal {
			continue
		}

		s.bossMap[int32(ftem.BiologyId)] = ftem

		if !s.isExsitMap(ftem.MapId) {
			s.mapIdList = append(s.mapIdList, ftem.MapId)
		}
	}
	return
}

func (s *worldBossTemplateService) GetWorldBossTemplateByBiologyId(biologyId int32) *gametemplate.WorldBossTemplate {
	return s.bossMap[biologyId]
}

func (s *worldBossTemplateService) GetMapIdList() []int32 {
	return s.mapIdList
}

func (s *worldBossTemplateService) isExsitMap(mapId int32) bool {
	for _, value := range s.mapIdList {
		if value == mapId {
			return true
		}
	}

	return false
}

var (
	once sync.Once
	s    *worldBossTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &worldBossTemplateService{}
		err = s.init()
	})

	return
}

func GetWorldBossTemplateService() WorldBossTemplateService {
	return s
}
