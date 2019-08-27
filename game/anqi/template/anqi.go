package anqi

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//暗器接口处理
type AnqiTemplateService interface {
	//获取进阶暗器配置
	GetAnqiNumber(number int32) *gametemplate.AnqiTemplate
	//获取暗器配置
	GetAnqi(id int) *gametemplate.AnqiTemplate
	//暗器丹配置
	GetAnqiDan(level int32) *gametemplate.AnqiDanTemplate
	//获取暗器技能
	GetAnqiSkill(advanceId int32) int32
	//吃培养丹升级
	GetAnQiEatDanTemplate(curLevel int32, num int32) (*gametemplate.AnqiDanTemplate, bool)
}

type anqiTemplateService struct {
	//进阶暗器配置
	anqiNumberMap map[int32]*gametemplate.AnqiTemplate
	//暗器配置
	anqiMap map[int]*gametemplate.AnqiTemplate
	//暗器丹
	anqiDanMap map[int32]*gametemplate.AnqiDanTemplate
}

//初始化
func (s *anqiTemplateService) init() error {
	s.anqiNumberMap = make(map[int32]*gametemplate.AnqiTemplate)
	s.anqiMap = make(map[int]*gametemplate.AnqiTemplate)
	s.anqiDanMap = make(map[int32]*gametemplate.AnqiDanTemplate)
	//暗器
	templateMap := template.GetTemplateService().GetAll((*gametemplate.AnqiTemplate)(nil))
	for _, templateObject := range templateMap {
		anqiTemplate, _ := templateObject.(*gametemplate.AnqiTemplate)
		s.anqiMap[anqiTemplate.TemplateId()] = anqiTemplate

		s.anqiNumberMap[anqiTemplate.Number] = anqiTemplate
	}

	//暗器丹
	templateMap = template.GetTemplateService().GetAll((*gametemplate.AnqiDanTemplate)(nil))
	for _, templateObject := range templateMap {
		anqiDanTemplate, _ := templateObject.(*gametemplate.AnqiDanTemplate)
		s.anqiDanMap[anqiDanTemplate.Level] = anqiDanTemplate
	}

	return nil
}

//获取进阶暗器配置
func (s *anqiTemplateService) GetAnqiNumber(number int32) *gametemplate.AnqiTemplate {
	to, ok := s.anqiNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//暗器丹配置
func (s *anqiTemplateService) GetAnqiDan(level int32) *gametemplate.AnqiDanTemplate {
	to, ok := s.anqiDanMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取暗器配置
func (s *anqiTemplateService) GetAnqi(id int) *gametemplate.AnqiTemplate {
	to, ok := s.anqiMap[id]
	if !ok {
		return nil
	}
	return to
}

// 获取暗器技能
func (s *anqiTemplateService) GetAnqiSkill(advanceId int32) (skillId int32) {
	temp, ok := s.anqiNumberMap[advanceId]
	if !ok {
		return
	}
	skillId = temp.Skill
	return
}

//吃培养丹升级
func (s *anqiTemplateService) GetAnQiEatDanTemplate(curLevel int32, num int32) (anQiDanTemplate *gametemplate.AnqiDanTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		anQiDanTemplate, flag = s.anqiDanMap[level]
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
	cs   *anqiTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &anqiTemplateService{}
		err = cs.init()
	})
	return err
}

func GetAnqiTemplateService() AnqiTemplateService {
	return cs
}
