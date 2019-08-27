package massacre

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//戮仙刃接口处理
type MassacreTemplateService interface {
	//获取戮仙刃配置
	GetMassacreeWeaponLev() int32
	//获取戮仙刃配置
	GetMassacre(id int) *gametemplate.MassacreTemplate
	//获取戮仙刃配置
	GetMassacreNumber(lev int32, star int32) *gametemplate.MassacreTemplate
}

type massacreTemplateService struct {
	//戮仙刃配置
	massacreWeaponLev int32
	//戮仙刃配置
	massacreMap map[int]*gametemplate.MassacreTemplate
	//戮仙刃配置
	massacreNumberMap map[int32]map[int32]*gametemplate.MassacreTemplate
}

//初始化
func (s *massacreTemplateService) init() error {
	s.massacreMap = make(map[int]*gametemplate.MassacreTemplate)
	s.massacreNumberMap = make(map[int32]map[int32]*gametemplate.MassacreTemplate)
	//戮仙刃
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MassacreTemplate)(nil))
	noWeaponMaxLev := int32(0)
	for _, templateObject := range templateMap {
		massacreTemplate, _ := templateObject.(*gametemplate.MassacreTemplate)
		s.massacreMap[massacreTemplate.TemplateId()] = massacreTemplate

		tempMap, ok := s.massacreNumberMap[massacreTemplate.Type]
		if !ok {
			tempMap = make(map[int32]*gametemplate.MassacreTemplate)
			s.massacreNumberMap[massacreTemplate.Type] = tempMap
		}
		tempMap[massacreTemplate.Star] = massacreTemplate

		if massacreTemplate.Type > noWeaponMaxLev && massacreTemplate.WeaponId == 0 {
			noWeaponMaxLev = massacreTemplate.Type
		}
	}
	s.massacreWeaponLev = noWeaponMaxLev + 1

	return nil
}

//获取戮仙刃配置
func (s *massacreTemplateService) GetMassacreeWeaponLev() int32 {
	return s.massacreWeaponLev
}

//获取戮仙刃配置
func (s *massacreTemplateService) GetMassacre(id int) *gametemplate.MassacreTemplate {
	to, ok := s.massacreMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取戮仙刃配置
func (s *massacreTemplateService) GetMassacreNumber(lev int32, star int32) *gametemplate.MassacreTemplate {
	levMap, ok := s.massacreNumberMap[lev]
	if !ok {
		return nil
	}
	starTemplate, ok := levMap[star]
	if !ok {
		return nil
	}
	return starTemplate
}

var (
	once sync.Once
	cs   *massacreTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &massacreTemplateService{}
		err = cs.init()
	})
	return err
}

func GetMassacreTemplateService() MassacreTemplateService {
	return cs
}
