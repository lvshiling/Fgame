package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"
	"sync"
)

type RobotTemplateService interface {
	GetRobotTemplate(force int64) *gametemplate.RobotTemplate
}

type robotTemplateService struct {
	robotList []*gametemplate.RobotTemplate
}

func (s *robotTemplateService) init() (err error) {
	s.robotList = make([]*gametemplate.RobotTemplate, 0, 8)

	templateRobotTemplateMap := template.GetTemplateService().GetAll((*gametemplate.RobotTemplate)(nil))

	for _, tempRobotTemplate := range templateRobotTemplateMap {
		robotTemplate := tempRobotTemplate.(*gametemplate.RobotTemplate)
		s.robotList = append(s.robotList, robotTemplate)
	}
	sort.Sort(robotTemplateList(s.robotList))
	if len(s.robotList) == 0 {
		err = fmt.Errorf("robot:机器人模板不能是0")
	}
	return
}

type robotTemplateList []*gametemplate.RobotTemplate

func (l robotTemplateList) Len() int {
	return len(l)
}

func (l robotTemplateList) Less(i, j int) bool {
	return l[i].ForceMin < l[j].ForceMin
}

func (l robotTemplateList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (s *robotTemplateService) GetRobotTemplate(force int64) *gametemplate.RobotTemplate {
	for _, tempRobotTemplate := range s.robotList {
		if force <= tempRobotTemplate.ForceMin {
			return tempRobotTemplate
		}
	}

	return s.robotList[len(s.robotList)-1]

}

var (
	once sync.Once
	s    *robotTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &robotTemplateService{}
		err = s.init()
	})
	return err
}

func GetRobotTemplateService() RobotTemplateService {
	return s
}
