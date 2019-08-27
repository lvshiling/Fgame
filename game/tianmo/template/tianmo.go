package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//天魔接口处理
type TianMoTemplateService interface {
	//获取进阶天魔配置
	GetTianMoNumber(number int32) *gametemplate.TianMoTemplate
	//获取天魔配置
	GetTianMo(id int) *gametemplate.TianMoTemplate
	//天魔丹配置
	GetTianMoDan(level int32) *gametemplate.TianMoDanTemplate
	//获取天魔技能
	GetTianMoSkill(advanceId int32) int32
	//吃培养丹升级
	GetTianMoEatDanTemplate(curLevel int32, num int32) (*gametemplate.TianMoDanTemplate, bool)
}

type tianMoTemplateService struct {
	//进阶天魔配置
	tianMoNumberMap map[int32]*gametemplate.TianMoTemplate
	//天魔配置
	tianMoMap map[int]*gametemplate.TianMoTemplate
	//天魔丹
	tianMoDanMap map[int32]*gametemplate.TianMoDanTemplate
}

//初始化
func (s *tianMoTemplateService) init() error {
	s.tianMoNumberMap = make(map[int32]*gametemplate.TianMoTemplate)
	s.tianMoMap = make(map[int]*gametemplate.TianMoTemplate)
	s.tianMoDanMap = make(map[int32]*gametemplate.TianMoDanTemplate)
	//天魔
	templateMap := template.GetTemplateService().GetAll((*gametemplate.TianMoTemplate)(nil))
	for _, templateObject := range templateMap {
		tianMoTemplate, _ := templateObject.(*gametemplate.TianMoTemplate)
		s.tianMoMap[tianMoTemplate.TemplateId()] = tianMoTemplate

		s.tianMoNumberMap[tianMoTemplate.Number] = tianMoTemplate
	}

	//天魔丹
	templateMap = template.GetTemplateService().GetAll((*gametemplate.TianMoDanTemplate)(nil))
	for _, templateObject := range templateMap {
		tianMoDanTemplate, _ := templateObject.(*gametemplate.TianMoDanTemplate)
		s.tianMoDanMap[tianMoDanTemplate.Level] = tianMoDanTemplate
	}

	return nil
}

//获取进阶天魔配置
func (s *tianMoTemplateService) GetTianMoNumber(number int32) *gametemplate.TianMoTemplate {
	to, ok := s.tianMoNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//天魔丹配置
func (s *tianMoTemplateService) GetTianMoDan(level int32) *gametemplate.TianMoDanTemplate {
	to, ok := s.tianMoDanMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取天魔配置
func (s *tianMoTemplateService) GetTianMo(id int) *gametemplate.TianMoTemplate {
	to, ok := s.tianMoMap[id]
	if !ok {
		return nil
	}
	return to
}

// 获取天魔技能
func (s *tianMoTemplateService) GetTianMoSkill(advanceId int32) (skillId int32) {
	temp, ok := s.tianMoNumberMap[advanceId]
	if !ok {
		return
	}
	skillId = temp.SkillId
	return
}

//吃培养丹升级
func (s *tianMoTemplateService) GetTianMoEatDanTemplate(curLevel int32, num int32) (anQiDanTemplate *gametemplate.TianMoDanTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		anQiDanTemplate, flag = s.tianMoDanMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= anQiDanTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *tianMoTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &tianMoTemplateService{}
		err = cs.init()
	})
	return err
}

func GetTianMoTemplateService() TianMoTemplateService {
	return cs
}
