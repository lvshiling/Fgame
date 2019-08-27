package bodyshield

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//护体盾接口处理
type BodyShieldService interface {
	//获取进阶护体盾配置
	GetBodyShieldNumber(number int32) *gametemplate.BodyShieldTemplate
	//获取护体盾配置
	GetBodyShield(id int) *gametemplate.BodyShieldTemplate
	//获取神盾尖刺配置
	GetShield(id int32) *gametemplate.ShieldTemplate
	//护体盾金甲丹配置
	GetBodyShieldJinJia(level int32) *gametemplate.BodyShieldJinJiaTemplate
	//吃金甲丹升级
	GetBodyShieldEatJinJiaTemplate(curLevel int32, num int32) (*gametemplate.BodyShieldJinJiaTemplate, bool)
}

type bodyShieldService struct {
	//进阶护体盾配置
	bShieldNumberMap map[int32]*gametemplate.BodyShieldTemplate
	//护体盾配置
	bShieldMap map[int]*gametemplate.BodyShieldTemplate
	//神盾尖刺
	shieldMap map[int32]*gametemplate.ShieldTemplate
	//护体盾金甲丹
	shieldJinJiaMap map[int32]*gametemplate.BodyShieldJinJiaTemplate
}

//初始化
func (s *bodyShieldService) init() error {
	s.bShieldNumberMap = make(map[int32]*gametemplate.BodyShieldTemplate)
	s.bShieldMap = make(map[int]*gametemplate.BodyShieldTemplate)
	s.shieldMap = make(map[int32]*gametemplate.ShieldTemplate)
	s.shieldJinJiaMap = make(map[int32]*gametemplate.BodyShieldJinJiaTemplate)
	//护体盾
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BodyShieldTemplate)(nil))
	for _, templateObject := range templateMap {
		bshieldTemplate, _ := templateObject.(*gametemplate.BodyShieldTemplate)
		s.bShieldMap[bshieldTemplate.TemplateId()] = bshieldTemplate

		s.bShieldNumberMap[bshieldTemplate.Number] = bshieldTemplate
	}
	//神盾尖刺
	templateMap = template.GetTemplateService().GetAll((*gametemplate.ShieldTemplate)(nil))
	for _, templateObject := range templateMap {
		shieldTemplate, _ := templateObject.(*gametemplate.ShieldTemplate)
		s.shieldMap[int32(shieldTemplate.TemplateId())] = shieldTemplate
	}

	//护体盾金甲丹
	templateMap = template.GetTemplateService().GetAll((*gametemplate.BodyShieldJinJiaTemplate)(nil))
	for _, templateObject := range templateMap {
		shieldJinJiaTemplate, _ := templateObject.(*gametemplate.BodyShieldJinJiaTemplate)
		s.shieldJinJiaMap[shieldJinJiaTemplate.Level] = shieldJinJiaTemplate
	}

	return nil
}

//获取进阶护体盾配置
func (s *bodyShieldService) GetBodyShieldNumber(number int32) *gametemplate.BodyShieldTemplate {
	to, ok := s.bShieldNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//护体盾金甲丹配置
func (s *bodyShieldService) GetBodyShieldJinJia(level int32) *gametemplate.BodyShieldJinJiaTemplate {
	to, ok := s.shieldJinJiaMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取护体盾配置
func (s *bodyShieldService) GetBodyShield(id int) *gametemplate.BodyShieldTemplate {
	to, ok := s.bShieldMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取神盾尖刺配置
func (s *bodyShieldService) GetShield(id int32) *gametemplate.ShieldTemplate {
	to, ok := s.shieldMap[id]
	if !ok {
		return nil
	}
	return to
}

//吃金甲丹升级
func (s *bodyShieldService) GetBodyShieldEatJinJiaTemplate(curLevel int32, num int32) (jinJiaTemplate *gametemplate.BodyShieldJinJiaTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		jinJiaTemplate, flag = s.shieldJinJiaMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= jinJiaTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *bodyShieldService
)

func Init() (err error) {
	once.Do(func() {
		cs = &bodyShieldService{}
		err = cs.init()
	})
	return err
}

func GetBodyShieldService() BodyShieldService {
	return cs
}
