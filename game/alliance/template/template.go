package template

import (
	"fgame/fgame/core/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"
	"sync"
)

//快捷缓存
//配置的整合
type AllianceTemplateService interface {
	GetAllianceTemplate(version alliancetypes.AllianceVersionType, level int32) *gametemplate.AllianceTemplate
	GetUnionDonateTemplate(juanXianType alliancetypes.AllianceJuanXianType) *gametemplate.UnionDonateTemplate
	GetAllianceSkillTemplate(id int32) *gametemplate.AllianceSkillTemplate
	GetAllianceSkillTemplateByType(level int32, skillType alliancetypes.AllianceSkillType) *gametemplate.AllianceSkillTemplate
	GetAllianceConstantTemp() *gametemplate.UnionConstantTemplate

	GetGuardTemplate(index int32) *gametemplate.WarGuardTemplate
	GetDoorRewardTemplate(door int32) *gametemplate.WarAwardDoorTemplate
	GetDoorRewardTemplateMap() map[int32]*gametemplate.WarAwardDoorTemplate
	//获取城战奖励
	GetAllianceSceneReward(alliancetypes.AllianceSceneRewardType) *gametemplate.WarAwardTemplate
	//获取连胜奖励
	GetAllianceSceneOccupyReward(winNum int32) *gametemplate.WarAwardOccupyTemplate
	//获取城战常量
	GetWarTemplate() *gametemplate.WarTemplate
	//

	//获取仙盟boss配置
	GetAllianceBossTemplate(level int32) *gametemplate.UnionBossTemplate
	// 获取仙盟最高等级
	GetAllianceMaxLevel(version alliancetypes.AllianceVersionType) int32
}

//快捷缓存
//配置的整合
type allianceTemplateService struct {
	allianceTemplateMap    map[alliancetypes.AllianceVersionType]map[int32]*gametemplate.AllianceTemplate
	unionDonateTemplateMap map[alliancetypes.AllianceJuanXianType]*gametemplate.UnionDonateTemplate
	unionConstantTemp      *gametemplate.UnionConstantTemplate
	//仙法配置
	allianceSkillTemplateMap     map[int32]*gametemplate.AllianceSkillTemplate
	allianceSkillTypeTemplateMap map[alliancetypes.AllianceSkillType]map[int32]*gametemplate.AllianceSkillTemplate

	sceneGuardTemplateMap      map[int32]*gametemplate.WarGuardTemplate
	doorAwardTemplateMap       map[int32]*gametemplate.WarAwardDoorTemplate
	warAwardTemplateMap        map[alliancetypes.AllianceSceneRewardType]*gametemplate.WarAwardTemplate
	warAwardOccupyTemplateList []*gametemplate.WarAwardOccupyTemplate
	warTemplate                *gametemplate.WarTemplate
	allianceBossMap            map[int32]*gametemplate.UnionBossTemplate
}

type warAwardOccupyList []*gametemplate.WarAwardOccupyTemplate

func (l warAwardOccupyList) Len() int {
	return len(l)
}

func (l warAwardOccupyList) Less(i, j int) bool {
	return l[i].OccupyCityContinue < l[j].OccupyCityContinue
}

func (l warAwardOccupyList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (s *allianceTemplateService) init() (err error) {
	s.allianceTemplateMap = make(map[alliancetypes.AllianceVersionType]map[int32]*gametemplate.AllianceTemplate)
	s.unionDonateTemplateMap = make(map[alliancetypes.AllianceJuanXianType]*gametemplate.UnionDonateTemplate)
	//加载模板
	templateObjMap := template.GetTemplateService().GetAll((*gametemplate.AllianceTemplate)(nil))
	for _, obj := range templateObjMap {
		allianceObj, _ := obj.(*gametemplate.AllianceTemplate)
		temp, ok := s.allianceTemplateMap[allianceObj.GetAllianceVersion()]
		if !ok {
			temp = make(map[int32]*gametemplate.AllianceTemplate)
			s.allianceTemplateMap[allianceObj.GetAllianceVersion()] = temp
		}
		temp[allianceObj.UnionLevel] = allianceObj
	}

	//加载模板
	templateUnionDonateMap := template.GetTemplateService().GetAll((*gametemplate.UnionDonateTemplate)(nil))
	for _, templateUnionDonateTemplate := range templateUnionDonateMap {
		unionDonateTemplate := templateUnionDonateTemplate.(*gametemplate.UnionDonateTemplate)
		s.unionDonateTemplateMap[unionDonateTemplate.GetJuanXianType()] = unionDonateTemplate
	}

	//加载仙法配置
	s.allianceSkillTemplateMap = make(map[int32]*gametemplate.AllianceSkillTemplate)
	s.allianceSkillTypeTemplateMap = make(map[alliancetypes.AllianceSkillType]map[int32]*gametemplate.AllianceSkillTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.AllianceSkillTemplate)(nil))
	for _, templateObject := range templateMap {
		alSkillTemplate, _ := templateObject.(*gametemplate.AllianceSkillTemplate)
		//按id
		s.allianceSkillTemplateMap[int32(alSkillTemplate.TemplateId())] = alSkillTemplate

		//按类型等级
		alSkillLevelTemplateMap, ok := s.allianceSkillTypeTemplateMap[alSkillTemplate.GetSkillType()]
		if !ok {
			alSkillLevelTemplateMap = make(map[int32]*gametemplate.AllianceSkillTemplate)
			s.allianceSkillTypeTemplateMap[alSkillTemplate.GetSkillType()] = alSkillLevelTemplateMap
		}
		alSkillLevelTemplateMap[alSkillTemplate.Level] = alSkillTemplate
	}

	//加载
	//s.sceneGuardTemplateMap = make(map[int32]map[int32]*gametemplate.WarGuardTemplate)
	s.sceneGuardTemplateMap = make(map[int32]*gametemplate.WarGuardTemplate)
	tempGuardTemplateMap := template.GetTemplateService().GetAll((*gametemplate.WarGuardTemplate)(nil))
	for _, tempGuardTemplate := range tempGuardTemplateMap {
		guardTemplate := tempGuardTemplate.(*gametemplate.WarGuardTemplate)
		// guardTemplateMap, ok := s.sceneGuardTemplateMap[guardTemplate.Map]
		// if !ok {
		// 	guardTemplateMap = make(map[int32]*gametemplate.WarGuardTemplate)
		// 	s.sceneGuardTemplateMap[guardTemplate.Map] = guardTemplateMap
		// }
		// guardTemplateMap[guardTemplate.SceneId] = guardTemplate
		_, exist := s.sceneGuardTemplateMap[guardTemplate.SceneId]
		if exist {
			return fmt.Errorf("alliance:scene_id应该是唯一的")
		}
		s.sceneGuardTemplateMap[guardTemplate.SceneId] = guardTemplate
	}

	//加载
	s.doorAwardTemplateMap = make(map[int32]*gametemplate.WarAwardDoorTemplate)
	tempWarAwardDoorTemplateMap := template.GetTemplateService().GetAll((*gametemplate.WarAwardDoorTemplate)(nil))
	for _, tempWarAwarddoorTemplate := range tempWarAwardDoorTemplateMap {
		warAwardDoorTemplate := tempWarAwarddoorTemplate.(*gametemplate.WarAwardDoorTemplate)
		s.doorAwardTemplateMap[int32(warAwardDoorTemplate.TemplateId())] = warAwardDoorTemplate
	}

	s.warAwardTemplateMap = make(map[alliancetypes.AllianceSceneRewardType]*gametemplate.WarAwardTemplate)
	tempWarAwardTemplateMap := template.GetTemplateService().GetAll((*gametemplate.WarAwardTemplate)(nil))
	for _, tempWarAwardTemplate := range tempWarAwardTemplateMap {
		warAwardTemplate := tempWarAwardTemplate.(*gametemplate.WarAwardTemplate)
		s.warAwardTemplateMap[warAwardTemplate.GetRewardType()] = warAwardTemplate
	}

	s.warAwardOccupyTemplateList = make([]*gametemplate.WarAwardOccupyTemplate, 0, 8)
	tempWarAwardOccupyTemplateMap := template.GetTemplateService().GetAll((*gametemplate.WarAwardOccupyTemplate)(nil))
	for _, tempWarAwardOccupyTemplate := range tempWarAwardOccupyTemplateMap {
		warAwardTemplate := tempWarAwardOccupyTemplate.(*gametemplate.WarAwardOccupyTemplate)
		s.warAwardOccupyTemplateList = append(s.warAwardOccupyTemplateList, warAwardTemplate)
	}

	sort.Sort(warAwardOccupyList(s.warAwardOccupyTemplateList))

	//加载城战常量配置
	tempWarTemplate := template.GetTemplateService().Get(1, (*gametemplate.WarTemplate)(nil))
	if tempWarTemplate == nil {
		return fmt.Errorf("alliance:城战常量不存在")
	}
	s.warTemplate = tempWarTemplate.(*gametemplate.WarTemplate)

	//加载仙盟常量配置
	unionConstantTo := template.GetTemplateService().Get(1, (*gametemplate.UnionConstantTemplate)(nil))
	if unionConstantTo == nil {
		return fmt.Errorf("alliance:仙盟常量不存在")
	}
	s.unionConstantTemp = unionConstantTo.(*gametemplate.UnionConstantTemplate)

	//加载仙盟boss配置
	s.initAllianceBoss()
	return nil
}

func (s *allianceTemplateService) initAllianceBoss() (err error) {
	s.allianceBossMap = make(map[int32]*gametemplate.UnionBossTemplate)

	templateUnionBossMap := template.GetTemplateService().GetAll((*gametemplate.UnionBossTemplate)(nil))
	for _, templateUnionBossTemplate := range templateUnionBossMap {
		unionBossTemplate := templateUnionBossTemplate.(*gametemplate.UnionBossTemplate)
		s.allianceBossMap[unionBossTemplate.Level] = unionBossTemplate
	}
	return
}

func (s *allianceTemplateService) GetAllianceTemplate(version alliancetypes.AllianceVersionType, level int32) *gametemplate.AllianceTemplate {
	allianceTempMap, ok := s.allianceTemplateMap[version]
	if !ok {
		return nil
	}
	allianceTemp, ok := allianceTempMap[level]
	if !ok {
		return nil
	}
	return allianceTemp
}

func (s *allianceTemplateService) GetUnionDonateTemplate(juanXianTypes alliancetypes.AllianceJuanXianType) *gametemplate.UnionDonateTemplate {
	unionDonateTemplate, ok := s.unionDonateTemplateMap[juanXianTypes]
	if !ok {
		return nil
	}
	return unionDonateTemplate
}

func (s *allianceTemplateService) GetAllianceSkillTemplateByType(level int32, skillType alliancetypes.AllianceSkillType) *gametemplate.AllianceSkillTemplate {
	skillLevelTemplateMap, ok := s.allianceSkillTypeTemplateMap[skillType]
	if !ok {
		return nil
	}
	to, ok := skillLevelTemplateMap[level]
	if !ok {
		return nil
	}
	return to
}
func (s *allianceTemplateService) GetAllianceSkillTemplate(id int32) *gametemplate.AllianceSkillTemplate {
	to, ok := s.allianceSkillTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

func (s *allianceTemplateService) GetAllianceMaxLevel(version alliancetypes.AllianceVersionType) int32 {
	allianceMap, ok := s.allianceTemplateMap[version]
	if !ok {
		return 0
	}
	for _, alliance := range allianceMap {
		if alliance.GetNextLevelAllianceTemplate() == nil {
			continue
		}
		if alliance.GetNextLevelAllianceTemplate().NextLevelId == 0 {
			return alliance.UnionLevel + 1
		}
	}
	return 0
}

// func (s *allianceTemplateService) GetGuardTemplate(mapId int32, index int32) *gametemplate.WarGuardTemplate {
// 	guardTemplateMap, ok := s.sceneGuardTemplateMap[mapId]
// 	if !ok {
// 		return nil
// 	}
// 	guardTemplate, ok := guardTemplateMap[index]
// 	if !ok {
// 		return nil
// 	}
// 	return guardTemplate
// }

func (s *allianceTemplateService) GetGuardTemplate(index int32) *gametemplate.WarGuardTemplate {
	guardTemplate, ok := s.sceneGuardTemplateMap[index]
	if !ok {
		return nil
	}
	return guardTemplate
}

func (s *allianceTemplateService) GetDoorRewardTemplate(door int32) *gametemplate.WarAwardDoorTemplate {
	id := door + 1

	doorAwardTemplate, ok := s.doorAwardTemplateMap[id]
	if !ok {
		return nil
	}
	return doorAwardTemplate
}

func (s *allianceTemplateService) GetDoorRewardTemplateMap() map[int32]*gametemplate.WarAwardDoorTemplate {
	return s.doorAwardTemplateMap
}

func (s *allianceTemplateService) GetAllianceSceneReward(typ alliancetypes.AllianceSceneRewardType) *gametemplate.WarAwardTemplate {
	t, ok := s.warAwardTemplateMap[typ]
	if !ok {
		return nil
	}
	return t
}

func (s *allianceTemplateService) GetAllianceSceneOccupyReward(winNum int32) *gametemplate.WarAwardOccupyTemplate {
	var warAwardOccupyTemplate *gametemplate.WarAwardOccupyTemplate
	for _, tempWarAwardOccupyTemplate := range s.warAwardOccupyTemplateList {
		if winNum == tempWarAwardOccupyTemplate.OccupyCityContinue {
			warAwardOccupyTemplate = tempWarAwardOccupyTemplate
			break
		}
	}
	return warAwardOccupyTemplate
}

func (s *allianceTemplateService) GetWarTemplate() *gametemplate.WarTemplate {
	return s.warTemplate
}

func (s *allianceTemplateService) GetAllianceConstantTemp() *gametemplate.UnionConstantTemplate {
	return s.unionConstantTemp
}

func (s *allianceTemplateService) GetAllianceBossTemplate(level int32) *gametemplate.UnionBossTemplate {
	allianceBossTemplate, ok := s.allianceBossMap[level]
	if !ok {
		return nil
	}
	return allianceBossTemplate
}

var (
	once sync.Once
	cs   *allianceTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &allianceTemplateService{}
		err = cs.init()
	})
	return err
}

func GetAllianceTemplateService() AllianceTemplateService {
	return cs
}
