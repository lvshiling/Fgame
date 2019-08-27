package template

import (
	"fgame/fgame/core/template"
	droptemplate "fgame/fgame/game/drop/template"
	inventorytypes "fgame/fgame/game/inventory/types"
	gametemplate "fgame/fgame/game/template"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fmt"
	"sync"
)

//屠龙装备模板处理
type TuLongEquipTemplateService interface {
	//融合模板
	GetTuLongEquipRongHeTemplat(jieShu int32) *gametemplate.TuLongEquipRongHeTemplate
	//转化模板
	GetTuLongEquipZhuanHuaTemplat(posType inventorytypes.BodyPositionType, jieShu int32) *gametemplate.TuLongEquipZhuanHuaTemplate
	//屠龙套装模板
	GetTuLongEquipTemplateBySuitGroup(suitGroup int32) *gametemplate.TuLongEquipSuitGroupTemplate
	//获取强化模板
	GetTuLongEquipStrengthenTemplate(suitType tulongequiptypes.TuLongSuitType, posType inventorytypes.BodyPositionType, level int32) *gametemplate.TuLongEquipStrengthenTemplate
	//套装技能模板（手动激活）
	GetTuLongEquipTemplateSkill(suitType tulongequiptypes.TuLongSuitType, level int32) *gametemplate.TuLongEquipSkillTemplate
}

type tuLongEquipTemplateService struct {
	// 屠龙装备配置
	tuLongEquipMap map[int]*gametemplate.TuLongEquipTemplate
	// 屠龙装备模板按套装
	tuLongEquipSuitGroupMap map[int32]*gametemplate.TuLongEquipSuitGroupTemplate
	// 屠龙装备重铸配置
	tuLongEquipRongHeMap map[int32]*gametemplate.TuLongEquipRongHeTemplate
	// 屠龙装备强化配置
	tuLongEquipStrengthenMap map[tulongequiptypes.TuLongSuitType]map[inventorytypes.BodyPositionType][]*gametemplate.TuLongEquipStrengthenTemplate
	// 屠龙装备转化配置
	tuLongEquipZhuanHuaMap map[inventorytypes.BodyPositionType]map[int32]*gametemplate.TuLongEquipZhuanHuaTemplate
	zhuanHuaMaxNumber      map[inventorytypes.BodyPositionType]int32
	// 屠龙装备技能
	tuLongSkillMap map[tulongequiptypes.TuLongSuitType]map[int32]*gametemplate.TuLongEquipSkillTemplate
}

//初始化
func (s *tuLongEquipTemplateService) init() error {
	s.tuLongEquipMap = make(map[int]*gametemplate.TuLongEquipTemplate)
	//屠龙装备
	templateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongEquipTemplate)(nil))
	for _, templateObject := range templateMap {
		tulongEquipTemplate, _ := templateObject.(*gametemplate.TuLongEquipTemplate)
		s.tuLongEquipMap[tulongEquipTemplate.TemplateId()] = tulongEquipTemplate
	}

	s.tuLongEquipSuitGroupMap = make(map[int32]*gametemplate.TuLongEquipSuitGroupTemplate)
	suitGroupTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongEquipSuitGroupTemplate)(nil))
	for _, templateObject := range suitGroupTemplateMap {
		suitGroupEquipTemplate, _ := templateObject.(*gametemplate.TuLongEquipSuitGroupTemplate)
		s.tuLongEquipSuitGroupMap[int32(suitGroupEquipTemplate.TemplateId())] = suitGroupEquipTemplate
	}

	// 强化概率
	s.tuLongEquipStrengthenMap = make(map[tulongequiptypes.TuLongSuitType]map[inventorytypes.BodyPositionType][]*gametemplate.TuLongEquipStrengthenTemplate)
	strengthenTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongEquipStrengthenTemplate)(nil))
	for _, temp := range strengthenTemplateMap {
		strengthenTemp, _ := temp.(*gametemplate.TuLongEquipStrengthenTemplate)
		subMap, ok := s.tuLongEquipStrengthenMap[strengthenTemp.GetSuitType()]
		if !ok {
			subMap = make(map[inventorytypes.BodyPositionType][]*gametemplate.TuLongEquipStrengthenTemplate)
			s.tuLongEquipStrengthenMap[strengthenTemp.GetSuitType()] = subMap
		}
		subMap[strengthenTemp.GetPosType()] = append(subMap[strengthenTemp.GetPosType()], strengthenTemp)
	}

	s.tuLongEquipRongHeMap = make(map[int32]*gametemplate.TuLongEquipRongHeTemplate)
	//屠龙重铸
	chongzhuTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongEquipRongHeTemplate)(nil))
	for _, templateObject := range chongzhuTemplateMap {
		chongzhuTemplate, _ := templateObject.(*gametemplate.TuLongEquipRongHeTemplate)
		s.tuLongEquipRongHeMap[chongzhuTemplate.Level] = chongzhuTemplate

		//验证掉落
		dropId := chongzhuTemplate.DropId
		flag := droptemplate.GetDropTemplateService().CheckSureDrop(dropId)
		if !flag {
			return fmt.Errorf("tulongequip: 屠龙装备融合配置掉落应该是必定掉落的，id:%d", chongzhuTemplate.Id)
		}
	}

	s.tuLongEquipZhuanHuaMap = make(map[inventorytypes.BodyPositionType]map[int32]*gametemplate.TuLongEquipZhuanHuaTemplate)
	s.zhuanHuaMaxNumber = make(map[inventorytypes.BodyPositionType]int32)
	//屠龙转化
	zhuanhuaTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongEquipZhuanHuaTemplate)(nil))
	for _, templateObject := range zhuanhuaTemplateMap {
		zhuanhuaTemplate, _ := templateObject.(*gametemplate.TuLongEquipZhuanHuaTemplate)
		subMap, ok := s.tuLongEquipZhuanHuaMap[zhuanhuaTemplate.GetPosType()]
		if !ok {
			subMap = make(map[int32]*gametemplate.TuLongEquipZhuanHuaTemplate)
			s.tuLongEquipZhuanHuaMap[zhuanhuaTemplate.GetPosType()] = subMap
		}
		subMap[zhuanhuaTemplate.Level] = zhuanhuaTemplate

		//记录最高转化阶数
		maxLevel, ok := s.zhuanHuaMaxNumber[zhuanhuaTemplate.GetPosType()]
		if !ok {
			s.zhuanHuaMaxNumber[zhuanhuaTemplate.GetPosType()] = zhuanhuaTemplate.Level
		}
		if ok && zhuanhuaTemplate.Level > maxLevel {
			s.zhuanHuaMaxNumber[zhuanhuaTemplate.GetPosType()] = zhuanhuaTemplate.Level
		}

		//验证掉落
		dropId := zhuanhuaTemplate.DropId
		flag := droptemplate.GetDropTemplateService().CheckSureDrop(dropId)
		if !flag {
			return fmt.Errorf("tulongequip: 屠龙装备转化配置掉落应该是必定掉落的，id:%d", zhuanhuaTemplate.Id)
		}
	}

	s.tuLongSkillMap = make(map[tulongequiptypes.TuLongSuitType]map[int32]*gametemplate.TuLongEquipSkillTemplate)
	//屠龙技能
	skillTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TuLongEquipSkillTemplate)(nil))
	for _, templateObject := range skillTemplateMap {
		skillTemplate, _ := templateObject.(*gametemplate.TuLongEquipSkillTemplate)
		subMap, ok := s.tuLongSkillMap[skillTemplate.GetSuitType()]
		if !ok {
			subMap = make(map[int32]*gametemplate.TuLongEquipSkillTemplate)
			s.tuLongSkillMap[skillTemplate.GetSuitType()] = subMap
		}
		subMap[skillTemplate.Level] = skillTemplate
	}

	return nil
}

func (s *tuLongEquipTemplateService) GetTuLongEquipStrengthenTemplate(suitType tulongequiptypes.TuLongSuitType, posType inventorytypes.BodyPositionType, level int32) *gametemplate.TuLongEquipStrengthenTemplate {
	subMap, ok := s.tuLongEquipStrengthenMap[suitType]
	if !ok {
		return nil
	}

	tempList, ok := subMap[posType]
	if !ok {
		return nil
	}

	for _, temp := range tempList {
		if temp.Level != level {
			continue
		}
		return temp
	}

	return nil
}

//套装模板
func (s *tuLongEquipTemplateService) GetTuLongEquipTemplateBySuitGroup(suitGroup int32) *gametemplate.TuLongEquipSuitGroupTemplate {
	return s.tuLongEquipSuitGroupMap[suitGroup]
}

//屠龙融合模板
func (s *tuLongEquipTemplateService) GetTuLongEquipRongHeTemplat(jieShu int32) *gametemplate.TuLongEquipRongHeTemplate {
	return s.tuLongEquipRongHeMap[jieShu]
}

//屠龙转化模板
func (s *tuLongEquipTemplateService) GetTuLongEquipZhuanHuaTemplat(posType inventorytypes.BodyPositionType, jieShu int32) *gametemplate.TuLongEquipZhuanHuaTemplate {
	subMap, ok := s.tuLongEquipZhuanHuaMap[posType]
	if !ok {
		return nil
	}

	temp, ok := subMap[jieShu]
	if !ok {
		maxJieShu := s.getZhuanHuaMaxNumber(posType)
		return subMap[maxJieShu]
	}

	return temp
}

func (s *tuLongEquipTemplateService) GetTuLongEquipTemplateSkill(posType tulongequiptypes.TuLongSuitType, level int32) *gametemplate.TuLongEquipSkillTemplate {
	subMap, ok := s.tuLongSkillMap[posType]
	if !ok {
		return nil
	}

	temp, ok := subMap[level]
	if !ok {
		return nil
	}

	return temp
}

func (s *tuLongEquipTemplateService) getZhuanHuaMaxNumber(posType inventorytypes.BodyPositionType) int32 {
	return s.zhuanHuaMaxNumber[posType]
}

var (
	once sync.Once
	cs   *tuLongEquipTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &tuLongEquipTemplateService{}
		err = cs.init()
	})
	return err
}

func GetTuLongEquipTemplateService() TuLongEquipTemplateService {
	return cs
}
