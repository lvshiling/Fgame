package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

type ShengTanTemplateService interface {
	//获取等级奖励
	GetShengTanAwardTemplate(lev int32) *gametemplate.UnionShengTanAwardTemplate
	//圣坛
	GetShengTanTemplate() *gametemplate.UnionShengTanTemplate
}

type shengTanTemplateService struct {
	awardTemplateList []*gametemplate.UnionShengTanAwardTemplate
	shengTanTemplate  *gametemplate.UnionShengTanTemplate
}

func (s *shengTanTemplateService) init() (err error) {
	s.awardTemplateList = make([]*gametemplate.UnionShengTanAwardTemplate, 0, 8)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.UnionShengTanAwardTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.UnionShengTanAwardTemplate)
		s.awardTemplateList = append(s.awardTemplateList, ftem)
	}

	//TODO: 排序
	tempShengTanTemplate := template.GetTemplateService().Get(1, (*gametemplate.UnionShengTanTemplate)(nil))
	if tempShengTanTemplate == nil {
		return fmt.Errorf("shengtan:圣坛常量不能为空")
	}
	s.shengTanTemplate = tempShengTanTemplate.(*gametemplate.UnionShengTanTemplate)

	return
}

func (s *shengTanTemplateService) GetShengTanAwardTemplate(lev int32) *gametemplate.UnionShengTanAwardTemplate {
	var awardTemplate *gametemplate.UnionShengTanAwardTemplate
	for _, tempAwardTemplate := range s.awardTemplateList {
		if lev >= tempAwardTemplate.MinLev {
			awardTemplate = tempAwardTemplate
			continue
		}
		break
	}
	return awardTemplate
}

func (s *shengTanTemplateService) GetShengTanTemplate() *gametemplate.UnionShengTanTemplate {

	return s.shengTanTemplate
}

var (
	once sync.Once
	s    *shengTanTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &shengTanTemplateService{}
		err = s.init()
	})

	return
}

func GetShengTanTemplateService() ShengTanTemplateService {
	return s
}
