package template

import (
	"fgame/fgame/core/template"
	coreutils "fgame/fgame/core/utils"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

type ChristmasTemplateService interface {
	//获取采集配置
	GetChristmasTemplateList() map[int32]*gametemplate.CollectActivityTemplate
	GetChristmasTemplate(group int32) *gametemplate.CollectActivityTemplate
}

type christmasTemplateService struct {
	collectTempMap map[int32]*gametemplate.CollectActivityTemplate
}

func (s *christmasTemplateService) init() (err error) {
	s.collectTempMap = make(map[int32]*gametemplate.CollectActivityTemplate)
	var biologyIdList []int32
	tempMap := template.GetTemplateService().GetAll((*gametemplate.CollectActivityTemplate)(nil))
	for _, temp := range tempMap {
		collectTemp, _ := temp.(*gametemplate.CollectActivityTemplate)
		s.collectTempMap[collectTemp.Group] = collectTemp

		biologyIdList = append(biologyIdList, collectTemp.BiologyId)
	}
	if coreutils.IfRepeatElementInt32(biologyIdList) {
		return fmt.Errorf("christmas:采集表生物id不能相同")
	}
	return
}

func (s *christmasTemplateService) GetChristmasTemplateList() map[int32]*gametemplate.CollectActivityTemplate {
	return s.collectTempMap
}

func (s *christmasTemplateService) GetChristmasTemplate(group int32) *gametemplate.CollectActivityTemplate {
	temp, ok := s.collectTempMap[group]
	if !ok {
		return nil
	}
	return temp
}

var (
	once sync.Once
	s    *christmasTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &christmasTemplateService{}
		err = s.init()
	})

	return
}

func GetChristmasTemplateService() ChristmasTemplateService {
	return s
}
