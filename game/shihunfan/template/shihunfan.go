package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//噬魂幡接口处理
type ShiHunFanTemplateService interface {
	//获取进阶噬魂幡配置
	GetShiHunFanNumber(number int32) *gametemplate.ShiHunFanTemplate
	//获取噬魂幡配置
	GetShiHunFan(id int) *gametemplate.ShiHunFanTemplate
	//噬魂幡丹配置
	GetShiHunFanDan(level int32) *gametemplate.ShiHunFanDanTemplate
	//获取噬魂幡技能
	GetShiHunFanSkill(advanceId int32) int32
	//吃培养丹升级
	GetShiHunFanEatDanTemplate(curLevel int32, num int32) (*gametemplate.ShiHunFanDanTemplate, bool)
}

type shiHunFanTemplateService struct {
	//进阶噬魂幡配置
	shiHunFanNumberMap map[int32]*gametemplate.ShiHunFanTemplate
	//噬魂幡配置
	shiHunFanMap map[int]*gametemplate.ShiHunFanTemplate
	//噬魂幡丹
	shiHunFanDanMap map[int32]*gametemplate.ShiHunFanDanTemplate
}

//初始化
func (s *shiHunFanTemplateService) init() error {
	s.shiHunFanNumberMap = make(map[int32]*gametemplate.ShiHunFanTemplate)
	s.shiHunFanMap = make(map[int]*gametemplate.ShiHunFanTemplate)
	s.shiHunFanDanMap = make(map[int32]*gametemplate.ShiHunFanDanTemplate)
	//噬魂幡
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ShiHunFanTemplate)(nil))
	for _, templateObject := range templateMap {
		shiHunFanTemplate, _ := templateObject.(*gametemplate.ShiHunFanTemplate)
		s.shiHunFanMap[shiHunFanTemplate.TemplateId()] = shiHunFanTemplate
		s.shiHunFanNumberMap[shiHunFanTemplate.Number] = shiHunFanTemplate
	}

	//噬魂幡丹
	templateMap = template.GetTemplateService().GetAll((*gametemplate.ShiHunFanDanTemplate)(nil))
	for _, templateObject := range templateMap {
		shiHunFanDanTemplate, _ := templateObject.(*gametemplate.ShiHunFanDanTemplate)
		s.shiHunFanDanMap[shiHunFanDanTemplate.Level] = shiHunFanDanTemplate
	}

	return nil
}

//获取进阶噬魂幡配置
func (s *shiHunFanTemplateService) GetShiHunFanNumber(number int32) *gametemplate.ShiHunFanTemplate {
	to, ok := s.shiHunFanNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//噬魂幡丹配置
func (s *shiHunFanTemplateService) GetShiHunFanDan(level int32) *gametemplate.ShiHunFanDanTemplate {
	to, ok := s.shiHunFanDanMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取噬魂幡配置
func (s *shiHunFanTemplateService) GetShiHunFan(id int) *gametemplate.ShiHunFanTemplate {
	to, ok := s.shiHunFanMap[id]
	if !ok {
		return nil
	}
	return to
}

// 获取噬魂幡技能
func (s *shiHunFanTemplateService) GetShiHunFanSkill(advanceId int32) (skillId int32) {
	temp, ok := s.shiHunFanNumberMap[advanceId]
	if !ok {
		return
	}
	skillId = temp.SkillId
	return
}

//吃培养丹升级
func (s *shiHunFanTemplateService) GetShiHunFanEatDanTemplate(curLevel int32, num int32) (shiHunFanDanTemplate *gametemplate.ShiHunFanDanTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		shiHunFanDanTemplate, flag = s.shiHunFanDanMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= shiHunFanDanTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *shiHunFanTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shiHunFanTemplateService{}
		err = cs.init()
	})
	return err
}

func GetShiHunFanTemplateService() ShiHunFanTemplateService {
	return cs
}
