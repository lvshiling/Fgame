package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type BuffTemplateService interface {
	GetBuff(id int32) *gametemplate.BuffTemplate
	GetAllBuffs() map[int32]*gametemplate.BuffTemplate
	GetBuffDongTai(id int32) *gametemplate.BuffDongTaiTemplate
	GetBuffTimeDuration(id int32, tianFuList []int32) int64
}

type buffTemplateService struct {
	buffTemplateMap        map[int32]*gametemplate.BuffTemplate
	buffDongTaiTemplateMap map[int32]*gametemplate.BuffDongTaiTemplate
}

func (s *buffTemplateService) init() (err error) {
	s.buffTemplateMap = make(map[int32]*gametemplate.BuffTemplate)

	templateBuffTemplateMap := template.GetTemplateService().GetAll((*gametemplate.BuffTemplate)(nil))
	for _, tempBuffTemplate := range templateBuffTemplateMap {
		buffTemplate := tempBuffTemplate.(*gametemplate.BuffTemplate)
		s.buffTemplateMap[int32(buffTemplate.TemplateId())] = buffTemplate
	}
	s.buffDongTaiTemplateMap = make(map[int32]*gametemplate.BuffDongTaiTemplate)

	templateBuffDongTaiTemplateMap := template.GetTemplateService().GetAll((*gametemplate.BuffDongTaiTemplate)(nil))
	for _, tempBuffDongTaiTemplate := range templateBuffDongTaiTemplateMap {
		buffDongTaiTemplate := tempBuffDongTaiTemplate.(*gametemplate.BuffDongTaiTemplate)
		s.buffDongTaiTemplateMap[int32(buffDongTaiTemplate.TemplateId())] = buffDongTaiTemplate
	}
	return
}

func (s *buffTemplateService) GetBuff(id int32) *gametemplate.BuffTemplate {
	buffTemplate, ok := s.buffTemplateMap[id]
	if !ok {
		return nil
	}
	return buffTemplate
}

func (s *buffTemplateService) GetAllBuffs() map[int32]*gametemplate.BuffTemplate {

	return s.buffTemplateMap
}
func (s *buffTemplateService) GetBuffDongTai(id int32) *gametemplate.BuffDongTaiTemplate {
	buffDongTaiTemplate, ok := s.buffDongTaiTemplateMap[id]
	if !ok {
		return nil
	}
	return buffDongTaiTemplate
}

func (s *buffTemplateService) GetBuffTimeDuration(id int32, tianFuList []int32) int64 {
	buffTemplate, ok := s.buffTemplateMap[id]
	if !ok {
		return 0
	}
	timeDuration := buffTemplate.TimeDuration
	for _, tianFuId := range tianFuList {
		buffDongTai := s.GetBuffDongTai(tianFuId)
		if buffDongTai == nil {
			continue
		}
		if timeDuration < buffDongTai.TimeDuration {
			timeDuration = buffDongTai.TimeDuration
		}
	}

	return timeDuration
}

var (
	once sync.Once
	s    *buffTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &buffTemplateService{}
		err = s.init()
	})
	return err
}

func GetBuffTemplateService() BuffTemplateService {
	return s
}
