package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type DingShiTemplateService interface {
	//获取定时boss配置
	GetDingShiBossTemplateByBiologyId(biologyId int32) *gametemplate.DingShiBossTemplate
	//获取藏经阁boss地图list
	GetMapIdList() []int32
	GetDingShiMap() map[int32]*gametemplate.DingShiBossTemplate
}

type dingShiTemplateService struct {
	bossMap   map[int32]*gametemplate.DingShiBossTemplate
	mapIdList []int32
}

func (s *dingShiTemplateService) init() (err error) {
	s.bossMap = make(map[int32]*gametemplate.DingShiBossTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.DingShiBossTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.DingShiBossTemplate)

		s.bossMap[int32(ftem.BiologyId)] = ftem

		if !s.isExsitMap(ftem.MapId) {
			s.mapIdList = append(s.mapIdList, ftem.MapId)
		}
	}
	return
}

func (s *dingShiTemplateService) GetDingShiMap() map[int32]*gametemplate.DingShiBossTemplate {
	return s.bossMap
}

func (s *dingShiTemplateService) GetDingShiBossTemplateByBiologyId(biologyId int32) *gametemplate.DingShiBossTemplate {
	return s.bossMap[biologyId]
}

func (s *dingShiTemplateService) GetMapIdList() []int32 {
	return s.mapIdList
}

func (s *dingShiTemplateService) isExsitMap(mapId int32) bool {
	for _, value := range s.mapIdList {
		if value == mapId {
			return true
		}
	}

	return false
}

var (
	once sync.Once
	s    *dingShiTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &dingShiTemplateService{}
		err = s.init()
	})

	return
}

func GetDingShiTemplateService() DingShiTemplateService {
	return s
}
