package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type CdTemplateService interface {
	GetCdGroup(id int32) *gametemplate.CdGroupTemplate
}

type cdTemplateService struct {
	cdGroupMap map[int32]*gametemplate.CdGroupTemplate
}

func (s *cdTemplateService) init() (err error) {

	s.cdGroupMap = make(map[int32]*gametemplate.CdGroupTemplate)
	templateCdGroupTemplateMap := template.GetTemplateService().GetAll((*gametemplate.CdGroupTemplate)(nil))
	for _, templateCdGroupTemplate := range templateCdGroupTemplateMap {
		cdGroupTemplate := templateCdGroupTemplate.(*gametemplate.CdGroupTemplate)
		s.cdGroupMap[int32(cdGroupTemplate.TemplateId())] = cdGroupTemplate
	}
	return
}

func (s *cdTemplateService) GetCdGroup(id int32) *gametemplate.CdGroupTemplate {
	cdGroupTemplate, ok := s.cdGroupMap[id]
	if !ok {
		return nil
	}
	return cdGroupTemplate
}

var (
	once sync.Once
	s    *cdTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &cdTemplateService{}
		err = s.init()
	})
	return err
}

func GetCdTemplateService() CdTemplateService {
	return s
}
