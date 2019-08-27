package template

import (
	"fgame/fgame/core/template"
	cangjinggetypes "fgame/fgame/game/cangjingge/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type CangJingGeTemplateService interface {
	//获取藏经阁boss配置
	GetCangJingGeTemplateByBiologyId(biologyId int32) *gametemplate.CangJingGeTemplate
	//获取藏经阁boss地图list
	GetMapIdList() []int32
}

type cangJingGeTemplateService struct {
	bossMap   map[int32]*gametemplate.CangJingGeTemplate
	mapIdList []int32
}

func (s *cangJingGeTemplateService) init() (err error) {
	s.bossMap = make(map[int32]*gametemplate.CangJingGeTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.CangJingGeTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.CangJingGeTemplate)
		if ftem.GetBossType() != cangjinggetypes.CangJingGeBossTypeLocal {
			continue
		}

		s.bossMap[int32(ftem.BiologyId)] = ftem

		if !s.isExsitMap(ftem.MapId) {
			s.mapIdList = append(s.mapIdList, ftem.MapId)
		}
	}
	return
}

func (s *cangJingGeTemplateService) GetCangJingGeTemplateByBiologyId(biologyId int32) *gametemplate.CangJingGeTemplate {
	return s.bossMap[biologyId]
}

func (s *cangJingGeTemplateService) GetMapIdList() []int32 {
	return s.mapIdList
}

func (s *cangJingGeTemplateService) isExsitMap(mapId int32) bool {
	for _, value := range s.mapIdList {
		if value == mapId {
			return true
		}
	}

	return false
}

var (
	once sync.Once
	s    *cangJingGeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &cangJingGeTemplateService{}
		err = s.init()
	})

	return
}

func GetCangJingGeTemplateService() CangJingGeTemplateService {
	return s
}
