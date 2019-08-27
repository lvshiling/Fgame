package template

import (
	"fgame/fgame/core/template"
	itemskilltypes "fgame/fgame/game/itemskill/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//物品技能接口处理
type ItemSkillTemplateService interface {
	//获取物品技能模板
	GetItemSkillTemplateByTypeAndLevel(typ itemskilltypes.ItemSkillType, level int32) *gametemplate.ItemSkillTemplate
	GetItemSkillTemplate(id int32) *gametemplate.ItemSkillTemplate
	GetSkillId(typ itemskilltypes.ItemSkillType, level int32) (skillId int32)
	GetTemplateBySkillId(skillId int32) *gametemplate.ItemSkillTemplate
}

type itemSkillTemplateService struct {
	//物品技能模板
	itemSkillTemplateMap map[int32]*gametemplate.ItemSkillTemplate
	//物品技能配置
	itemSkillTypeTemplateMap map[itemskilltypes.ItemSkillType]map[int32]*gametemplate.ItemSkillTemplate
	//物品技能模板
	itemSkillTemplateSkillIdMap map[int32]*gametemplate.ItemSkillTemplate
}

//初始化
func (s *itemSkillTemplateService) init() error {
	s.itemSkillTemplateMap = make(map[int32]*gametemplate.ItemSkillTemplate)
	s.itemSkillTypeTemplateMap = make(map[itemskilltypes.ItemSkillType]map[int32]*gametemplate.ItemSkillTemplate)
	s.itemSkillTemplateSkillIdMap = make(map[int32]*gametemplate.ItemSkillTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.ItemSkillTemplate)(nil))
	for _, templateObject := range templateMap {
		skTemplate, _ := templateObject.(*gametemplate.ItemSkillTemplate)

		s.itemSkillTemplateMap[int32(skTemplate.TemplateId())] = skTemplate
		s.itemSkillTemplateSkillIdMap[skTemplate.SkillId] = skTemplate

		typ := skTemplate.GetType()
		level := skTemplate.Level

		skTypeTemplateMap, ok := s.itemSkillTypeTemplateMap[typ]
		if !ok {
			skTypeTemplateMap = make(map[int32]*gametemplate.ItemSkillTemplate)
			s.itemSkillTypeTemplateMap[typ] = skTypeTemplateMap
		}
		skTypeTemplateMap[level] = skTemplate
	}

	return nil
}

func (s *itemSkillTemplateService) GetSkillId(typ itemskilltypes.ItemSkillType, level int32) (skillId int32) {
	skillId = int32(0)
	to := s.GetItemSkillTemplateByTypeAndLevel(typ, level)
	if to != nil {
		skillId = to.SkillId
	}
	return skillId
}

//获取物品技能配置
func (s *itemSkillTemplateService) GetItemSkillTemplateByTypeAndLevel(typ itemskilltypes.ItemSkillType, level int32) *gametemplate.ItemSkillTemplate {
	to, ok := s.itemSkillTypeTemplateMap[typ][level]
	if !ok {
		return nil
	}
	return to
}

func (s *itemSkillTemplateService) GetItemSkillTemplate(id int32) *gametemplate.ItemSkillTemplate {
	to, ok := s.itemSkillTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

func (s *itemSkillTemplateService) GetTemplateBySkillId(skillId int32) *gametemplate.ItemSkillTemplate {
	to, ok := s.itemSkillTemplateSkillIdMap[skillId]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *itemSkillTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &itemSkillTemplateService{}
		err = cs.init()
	})
	return err
}

func GetItemSkillTemplateService() ItemSkillTemplateService {
	return cs
}
