package template

import (
	"fgame/fgame/core/template"
	hunttypes "fgame/fgame/game/hunt/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//寻宝模板处理
type HuntTemplateService interface {
	GetHuntTemplat(huntType hunttypes.HuntType) *gametemplate.HuntTemplate
	GetHuntTimesTemplat(times int32) *gametemplate.HuntTimesTemplate
}

type huntTemplateService struct {
	//寻宝配置
	huntMap      map[hunttypes.HuntType]*gametemplate.HuntTemplate
	huntTimesMap map[int32]*gametemplate.HuntTimesTemplate
}

//初始化
func (s *huntTemplateService) init() error {
	s.huntMap = make(map[hunttypes.HuntType]*gametemplate.HuntTemplate)
	//寻宝
	templateMap := template.GetTemplateService().GetAll((*gametemplate.HuntTemplate)(nil))
	for _, templateObject := range templateMap {
		huntTemplate, _ := templateObject.(*gametemplate.HuntTemplate)
		s.huntMap[huntTemplate.GetHuntType()] = huntTemplate
	}

	s.huntTimesMap = make(map[int32]*gametemplate.HuntTimesTemplate)
	//寻宝
	timesTemplateMap := template.GetTemplateService().GetAll((*gametemplate.HuntTimesTemplate)(nil))
	for _, templateObject := range timesTemplateMap {
		huntTimesTemplate, _ := templateObject.(*gametemplate.HuntTimesTemplate)
		s.huntTimesMap[huntTimesTemplate.Times] = huntTimesTemplate
	}

	return nil
}

func (s huntTemplateService) GetHuntTemplat(huntType hunttypes.HuntType) *gametemplate.HuntTemplate {
	temp, ok := s.huntMap[huntType]
	if !ok {
		return nil
	}

	return temp
}
func (s huntTemplateService) GetHuntTimesTemplat(times int32) *gametemplate.HuntTimesTemplate {
	temp, ok := s.huntTimesMap[times]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	cs   *huntTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &huntTemplateService{}
		err = cs.init()
	})
	return err
}

func GetHuntTemplateService() HuntTemplateService {
	return cs
}
