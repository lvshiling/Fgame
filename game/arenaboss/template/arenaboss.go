package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type ArenaBossTemplateService interface {
	//获取跨服世界boss配置
	GetArenaBossTemplateByBiologyId(biologyId int32) *gametemplate.ShengShouBossTemplate
	//获取跨服世界boss地图list
	GetMapIdList() []int32
}

type arenaBossTemplateService struct {
	bossMap   map[int32]*gametemplate.ShengShouBossTemplate
	mapIdList []int32
}

func (s *arenaBossTemplateService) init() (err error) {
	s.bossMap = make(map[int32]*gametemplate.ShengShouBossTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.ShengShouBossTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.ShengShouBossTemplate)

		s.bossMap[int32(ftem.BiologyId)] = ftem

		if !s.isExsitMap(ftem.MapId) {
			s.mapIdList = append(s.mapIdList, ftem.MapId)
		}
	}
	return
}

func (s *arenaBossTemplateService) GetArenaBossTemplateByBiologyId(biologyId int32) *gametemplate.ShengShouBossTemplate {
	return s.bossMap[biologyId]
}

func (s *arenaBossTemplateService) GetMapIdList() []int32 {
	return s.mapIdList
}

func (s *arenaBossTemplateService) isExsitMap(mapId int32) bool {
	for _, value := range s.mapIdList {
		if value == mapId {
			return true
		}
	}

	return false
}

var (
	once sync.Once
	s    *arenaBossTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &arenaBossTemplateService{}
		err = s.init()
	})

	return
}

func GetArenaBossTemplateService() ArenaBossTemplateService {
	return s
}
