package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type ZhenXiTemplateService interface {
	//获取藏经阁boss配置
	GetZhenXiBossTemplateByBiologyId(biologyId int32) *gametemplate.ZhenXiBossTemplate
	//获取藏经阁boss地图list
	GetMapIdList() []int32
	GetBossUseTemplate(mapId int32) *gametemplate.ZhenXiBossUseTemplate
}

type zhenXiTemplateService struct {
	bossMap    map[int32]*gametemplate.ZhenXiBossTemplate
	bossUseMap map[int32]*gametemplate.ZhenXiBossUseTemplate
	mapIdList  []int32
}

func (s *zhenXiTemplateService) init() (err error) {
	s.bossMap = make(map[int32]*gametemplate.ZhenXiBossTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.ZhenXiBossTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.ZhenXiBossTemplate)

		s.bossMap[int32(ftem.BiologyId)] = ftem

		if !s.isExsitMap(ftem.MapId) {
			s.mapIdList = append(s.mapIdList, ftem.MapId)
		}
	}
	s.bossUseMap = make(map[int32]*gametemplate.ZhenXiBossUseTemplate)
	tempBossUseMap := template.GetTemplateService().GetAll((*gametemplate.ZhenXiBossUseTemplate)(nil))
	for _, tem := range tempBossUseMap {
		ftem, _ := tem.(*gametemplate.ZhenXiBossUseTemplate)

		s.bossUseMap[int32(ftem.MapId)] = ftem

	}
	return
}

func (s *zhenXiTemplateService) GetZhenXiBossTemplateByBiologyId(biologyId int32) *gametemplate.ZhenXiBossTemplate {
	return s.bossMap[biologyId]
}

func (s *zhenXiTemplateService) GetBossUseTemplate(mapId int32) *gametemplate.ZhenXiBossUseTemplate {
	return s.bossUseMap[mapId]
}

func (s *zhenXiTemplateService) GetMapIdList() []int32 {
	return s.mapIdList
}

func (s *zhenXiTemplateService) isExsitMap(mapId int32) bool {
	for _, value := range s.mapIdList {
		if value == mapId {
			return true
		}
	}

	return false
}

var (
	once sync.Once
	s    *zhenXiTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &zhenXiTemplateService{}
		err = s.init()
	})

	return
}

func GetZhenXiTemplateService() ZhenXiTemplateService {
	return s
}
