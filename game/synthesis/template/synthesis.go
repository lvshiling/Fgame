package synthesis

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type SynthesisTemplateService interface {
	//获取合成配置
	GetSynthesis(id int32) *gametemplate.SynthesisTemplate
}

type synthesisTemplateService struct {
	//合成配置
	synMap map[int32]*gametemplate.SynthesisTemplate
}

//初始化合成配置
func (s *synthesisTemplateService) init() (err error) {
	s.synMap = make(map[int32]*gametemplate.SynthesisTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.SynthesisTemplate)(nil))
	for _, temp := range tempMap {
		synTemp := temp.(*gametemplate.SynthesisTemplate)
		s.synMap[int32(synTemp.Id)] = synTemp
	}

	return nil
}

func (s *synthesisTemplateService) GetSynthesis(id int32) *gametemplate.SynthesisTemplate {
	if value, ok := s.synMap[id]; ok {
		return value
	}
	return nil
}

var (
	once sync.Once
	s    *synthesisTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &synthesisTemplateService{}
		err = s.init()
	})
	return
}

func GetSynthesisTemplateService() SynthesisTemplateService {
	return s
}
