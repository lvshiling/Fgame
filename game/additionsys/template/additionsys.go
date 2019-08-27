package template

import (
	"fgame/fgame/core/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	itemtypes "fgame/fgame/game/item/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

// 系统升级通用接口
type SystemShengJiCommonTemplate interface {
	GetUseItemTemplate() *gametemplate.ItemTemplate
	// filed
	GetLevel() int32
	GetHp() int32
	GetAttack() int32
	GetDefence() int32
	GetUseMoney() int32
	GetUseItem() int32
	GetItemCount() int32
	GetUpdateWfb() int32
	GetZhufuMax() int32
	GetAddMin() int32
	GetAddMax() int32
	GetTimesMin() int32
	GetTimesMax() int32
}

// 系统新升级通用接口
type SystemUpgradeCommonTemplate interface {
	// filed
	GetLevel() int32
	GetHp() int32
	GetAttack() int32
	GetDefence() int32
	GetPercent() int32
	GetUseMoney() int32
	GetItemCount() int32
	GetUpdateWfb() int32
	GetZhufuMax() int32
	GetAddMin() int32
	GetAddMax() int32
	GetTimesMin() int32
	GetTimesMax() int32
}

//接口处理
type AdditionSysTemplateService interface {
	//获取装备配置
	GetEquipTemplate(id int) *gametemplate.SystemEquipTemplate
	//获取部位强化配置
	GetBodyStrengthenByArg(typ additionsystypes.AdditionSysType, styp additionsystypes.SlotPositionType, lev int32) *gametemplate.SystemStrengthenTemplate
	//获取套装配置
	GetTaoZhuangByArg(typ additionsystypes.AdditionSysType, qualityType itemtypes.ItemQualityType) *gametemplate.SystemTaozhuangTemplate
	//获取升级配置
	GetShengJiByArg(typ additionsystypes.AdditionSysType, lev int32) SystemShengJiCommonTemplate
	//获取化灵配置
	GetHuaLingByArg(typ additionsystypes.AdditionSysType, lev int32) (levelTemp *gametemplate.SystemHuaLingTemplate, typeTemp *gametemplate.SystemHuaLingUseTemplate)
	//获取觉醒配置
	GetAwakeByType(typ additionsystypes.AdditionSysType) *gametemplate.SystemAwakeUseTemplate
	//获取觉醒属性配置
	GetAwakeByArg(typ additionsystypes.AdditionSysType, advanced int32) *gametemplate.SystemAwakeTemplate
	//获取觉醒等级属性配置
	GetAwakeLevelByArg(typ additionsystypes.AdditionSysType, advanced int32, level int32) *gametemplate.SystemAwakeLevelTemplate
	//获取食用物品数量达到化灵配置
	GetHuaLingReachByArg(typ additionsystypes.AdditionSysType, curCulLevel int32, num int32) (reachTemplate *gametemplate.SystemHuaLingTemplate, flag bool)
	//系统神铸
	GetShenZhuByArg(pos additionsystypes.SlotPositionType, lev int32) (temp *gametemplate.SystemShenZhuTemplate)
	//系统神铸使用道具
	GetShenZhuUseByType(typ additionsystypes.AdditionSysType) (typeTemp *gametemplate.SystemShenZhuUseTemplate)
	// 系统通灵
	GetTongLingByLevel(lev int32) (levelTemp SystemUpgradeCommonTemplate)
	//系统通灵使用道具
	GetTongLingUseByType(typ additionsystypes.AdditionSysType) (typeTemp *gametemplate.SystemTongLingUseTemplate)
	//获取食用物品数量达到通灵配置
	GetTongLingReachByArg(typ additionsystypes.AdditionSysType, curCulLevel int32, num int32) (reachTemplate SystemUpgradeCommonTemplate, flag bool)
	//获取灵珠模板
	GetLingZhuTemplate(lingzhuType additionsystypes.LingZhuType) *gametemplate.SystemLingZhuTemplate
	//获取灵珠技能模板
	GetLingZhuSkillTemplate(lingtongId int32, level int32) (skillTemp, beforeSkillTemp *gametemplate.SystemLingZhuSkillTemplate)
}

type additionSysTemplateService struct {
	//装备配置
	equipByIdMap map[int]*gametemplate.SystemEquipTemplate
	//部位强化配置
	bodyStrengthenByArgMap map[additionsystypes.AdditionSysType]map[additionsystypes.SlotPositionType]map[int32]*gametemplate.SystemStrengthenTemplate
	//套装配置
	taoZhuangByArgMap map[additionsystypes.AdditionSysType]map[itemtypes.ItemQualityType]*gametemplate.SystemTaozhuangTemplate
	//升级配置
	shengJiByArgMap map[additionsystypes.AdditionSysType]map[int32]SystemShengJiCommonTemplate
	//化灵等级配置
	huaLingByLevelMap map[int32]*gametemplate.SystemHuaLingTemplate
	//化灵系统类型配置
	huaLingByTypeMap map[additionsystypes.AdditionSysType]*gametemplate.SystemHuaLingUseTemplate
	//觉醒配置
	awakeByTypeMap map[additionsystypes.AdditionSysType]*gametemplate.SystemAwakeUseTemplate
	//觉醒属性配置
	AwakeByTypeAndAdvancedMap map[additionsystypes.AdditionSysType]map[int32]*gametemplate.SystemAwakeTemplate
	//神铸等级配置
	shenZhuByArgMap map[additionsystypes.SlotPositionType]map[int32]*gametemplate.SystemShenZhuTemplate
	//神铸系统类型配置
	shenZhuByTypeMap map[additionsystypes.AdditionSysType]*gametemplate.SystemShenZhuUseTemplate
	//通灵等级配置
	tongLingByLevelMap map[int32]SystemUpgradeCommonTemplate
	//通灵系统类型配置
	tongLingByTypeMap map[additionsystypes.AdditionSysType]*gametemplate.SystemTongLingUseTemplate
	//灵珠配置
	lingzhuMap map[additionsystypes.LingZhuType]*gametemplate.SystemLingZhuTemplate
	//灵珠技能配置
	lingzhuSkillMap map[int32]map[int32]*gametemplate.SystemLingZhuSkillTemplate
}

//初始化
func (s *additionSysTemplateService) init() error {
	s.equipByIdMap = make(map[int]*gametemplate.SystemEquipTemplate)
	//装备
	templateMap := template.GetTemplateService().GetAll((*gametemplate.SystemEquipTemplate)(nil))
	for _, templateObject := range templateMap {
		equipTemplate, _ := templateObject.(*gametemplate.SystemEquipTemplate)
		s.equipByIdMap[equipTemplate.TemplateId()] = equipTemplate
	}

	//部位强化
	s.bodyStrengthenByArgMap = make(map[additionsystypes.AdditionSysType]map[additionsystypes.SlotPositionType]map[int32]*gametemplate.SystemStrengthenTemplate)
	bodyStrengthenTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemStrengthenTemplate)(nil))
	for _, templateObject := range bodyStrengthenTemplateMap {
		strengthenTemplate, _ := templateObject.(*gametemplate.SystemStrengthenTemplate)
		tempMM, ok := s.bodyStrengthenByArgMap[additionsystypes.AdditionSysType(strengthenTemplate.Type)]
		if !ok {
			tempMM = make(map[additionsystypes.SlotPositionType]map[int32]*gametemplate.SystemStrengthenTemplate)
			s.bodyStrengthenByArgMap[additionsystypes.AdditionSysType(strengthenTemplate.Type)] = tempMM
		}
		tempM, ok := tempMM[additionsystypes.SlotPositionType(strengthenTemplate.SubType)]
		if !ok {
			tempM = make(map[int32]*gametemplate.SystemStrengthenTemplate)
			tempMM[additionsystypes.SlotPositionType(strengthenTemplate.SubType)] = tempM
		}
		tempM[strengthenTemplate.Level] = strengthenTemplate
	}

	//套装
	s.taoZhuangByArgMap = make(map[additionsystypes.AdditionSysType]map[itemtypes.ItemQualityType]*gametemplate.SystemTaozhuangTemplate)
	taoZhuangTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemTaozhuangTemplate)(nil))
	for _, templateObject := range taoZhuangTemplateMap {
		groupTemplate, _ := templateObject.(*gametemplate.SystemTaozhuangTemplate)
		tempM, ok := s.taoZhuangByArgMap[groupTemplate.GetTaozhuangType()]
		if !ok {
			tempM = make(map[itemtypes.ItemQualityType]*gametemplate.SystemTaozhuangTemplate)
			s.taoZhuangByArgMap[groupTemplate.GetTaozhuangType()] = tempM
		}
		tempM[groupTemplate.GetTaozhuangQuality()] = groupTemplate
	}

	//系统升级
	s.shengJiByArgMap = make(map[additionsystypes.AdditionSysType]map[int32]SystemShengJiCommonTemplate)
	shengJiTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemShengJiTemplate)(nil))
	for _, templateObject := range shengJiTemplateMap {
		shengJiTemplate, _ := templateObject.(*gametemplate.SystemShengJiTemplate)
		tempM, ok := s.shengJiByArgMap[additionsystypes.AdditionSysType(shengJiTemplate.SysType)]
		if !ok {
			tempM = make(map[int32]SystemShengJiCommonTemplate)
			s.shengJiByArgMap[additionsystypes.AdditionSysType(shengJiTemplate.SysType)] = tempM
		}
		tempM[shengJiTemplate.Level] = shengJiTemplate
	}
	//系统升级-圣痕
	shengHenTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemShengJiShengHenTemplate)(nil))
	for _, templateObject := range shengHenTemplateMap {
		shengJiTemplate, _ := templateObject.(*gametemplate.SystemShengJiShengHenTemplate)
		tempM, ok := s.shengJiByArgMap[additionsystypes.AdditionSysType(shengJiTemplate.SysType)]
		if !ok {
			tempM = make(map[int32]SystemShengJiCommonTemplate)
			s.shengJiByArgMap[additionsystypes.AdditionSysType(shengJiTemplate.SysType)] = tempM
		}
		tempM[shengJiTemplate.Level] = shengJiTemplate
	}

	//系统化灵
	s.huaLingByLevelMap = make(map[int32]*gametemplate.SystemHuaLingTemplate)
	huaLingTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemHuaLingTemplate)(nil))
	for _, templateObject := range huaLingTemplateMap {
		huaLingTemplate, _ := templateObject.(*gametemplate.SystemHuaLingTemplate)
		s.huaLingByLevelMap[huaLingTemplate.Level] = huaLingTemplate
	}
	s.huaLingByTypeMap = make(map[additionsystypes.AdditionSysType]*gametemplate.SystemHuaLingUseTemplate)
	huaLingUseTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemHuaLingUseTemplate)(nil))
	for _, templateObject := range huaLingUseTemplateMap {
		huaLingUseTemplate, _ := templateObject.(*gametemplate.SystemHuaLingUseTemplate)
		s.huaLingByTypeMap[additionsystypes.AdditionSysType(huaLingUseTemplate.SysType)] = huaLingUseTemplate
	}

	//系统神铸
	s.shenZhuByArgMap = make(map[additionsystypes.SlotPositionType]map[int32]*gametemplate.SystemShenZhuTemplate)
	shenZhuTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemShenZhuTemplate)(nil))
	for _, templateObject := range shenZhuTemplateMap {
		shenZhuTemplate, _ := templateObject.(*gametemplate.SystemShenZhuTemplate)
		tempM, ok := s.shenZhuByArgMap[shenZhuTemplate.GetPos()]
		if !ok {
			tempM = make(map[int32]*gametemplate.SystemShenZhuTemplate)
			s.shenZhuByArgMap[shenZhuTemplate.GetPos()] = tempM
		}
		tempM[shenZhuTemplate.Level] = shenZhuTemplate
	}
	s.shenZhuByTypeMap = make(map[additionsystypes.AdditionSysType]*gametemplate.SystemShenZhuUseTemplate)
	shenZhuUseTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemShenZhuUseTemplate)(nil))
	for _, templateObject := range shenZhuUseTemplateMap {
		shenZhuUseTemplate, _ := templateObject.(*gametemplate.SystemShenZhuUseTemplate)
		s.shenZhuByTypeMap[additionsystypes.AdditionSysType(shenZhuUseTemplate.SysType)] = shenZhuUseTemplate
	}

	//系统觉醒
	s.awakeByTypeMap = make(map[additionsystypes.AdditionSysType]*gametemplate.SystemAwakeUseTemplate)
	awakeTempMap := template.GetTemplateService().GetAll((*gametemplate.SystemAwakeUseTemplate)(nil))
	for _, temp := range awakeTempMap {
		awakeTemp, _ := temp.(*gametemplate.SystemAwakeUseTemplate)
		s.awakeByTypeMap[awakeTemp.GetSysType()] = awakeTemp
	}

	s.AwakeByTypeAndAdvancedMap = make(map[additionsystypes.AdditionSysType]map[int32]*gametemplate.SystemAwakeTemplate)
	awakeTempMap = template.GetTemplateService().GetAll((*gametemplate.SystemAwakeTemplate)(nil))
	for _, temp := range awakeTempMap {
		awakeTemp, _ := temp.(*gametemplate.SystemAwakeTemplate)
		awakeMap, ok := s.AwakeByTypeAndAdvancedMap[awakeTemp.GetSysType()]
		if !ok {
			awakeMap = make(map[int32]*gametemplate.SystemAwakeTemplate)
			s.AwakeByTypeAndAdvancedMap[awakeTemp.GetSysType()] = awakeMap
		}
		awakeMap[awakeTemp.Number] = awakeTemp
	}

	//系统通灵
	s.tongLingByLevelMap = make(map[int32]SystemUpgradeCommonTemplate)
	tongLingTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemTongLingTemplate)(nil))
	for _, templateObject := range tongLingTemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.SystemTongLingTemplate)
		s.tongLingByLevelMap[tongLingTemplate.Level] = tongLingTemplate
	}
	s.tongLingByTypeMap = make(map[additionsystypes.AdditionSysType]*gametemplate.SystemTongLingUseTemplate)
	tongLingUseTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SystemTongLingUseTemplate)(nil))
	for _, templateObject := range tongLingUseTemplateMap {
		tongLingUseTemplate, _ := templateObject.(*gametemplate.SystemTongLingUseTemplate)
		s.tongLingByTypeMap[additionsystypes.AdditionSysType(tongLingUseTemplate.SysType)] = tongLingUseTemplate
	}

	//系统灵珠
	s.lingzhuMap = make(map[additionsystypes.LingZhuType]*gametemplate.SystemLingZhuTemplate)
	lingzhulist := template.GetTemplateService().GetAll((*gametemplate.SystemLingZhuTemplate)(nil))
	for _, tempObj := range lingzhulist {
		temp, _ := tempObj.(*gametemplate.SystemLingZhuTemplate)
		s.lingzhuMap[temp.GetLingZhuType()] = temp
	}
	s.lingzhuSkillMap = make(map[int32]map[int32]*gametemplate.SystemLingZhuSkillTemplate)
	lingzhuskilllist := template.GetTemplateService().GetAll((*gametemplate.SystemLingZhuSkillTemplate)(nil))
	for _, tempObj := range lingzhuskilllist {
		temp, _ := tempObj.(*gametemplate.SystemLingZhuSkillTemplate)
		tempMap, ok := s.lingzhuSkillMap[temp.LingtongId]
		if !ok {
			tempMap = make(map[int32]*gametemplate.SystemLingZhuSkillTemplate)
		}
		tempMap[temp.Level] = temp
		s.lingzhuSkillMap[temp.LingtongId] = tempMap
	}

	return nil
}

//系统灵珠
func (s additionSysTemplateService) GetLingZhuTemplate(lingzhuType additionsystypes.LingZhuType) *gametemplate.SystemLingZhuTemplate {
	t, ok := s.lingzhuMap[lingzhuType]
	if !ok {
		return nil
	}
	return t
}

func (s additionSysTemplateService) GetLingZhuSkillTemplate(lingtongId int32, totalLevel int32) (skillTemp, beforeSkillTemp *gametemplate.SystemLingZhuSkillTemplate) {
	_, ok := s.lingzhuSkillMap[lingtongId]
	if !ok {
		return nil, nil
	}
	for level := int32(len(s.lingzhuSkillMap[lingtongId])); level >= int32(1); level-- {
		tt, ook := s.lingzhuSkillMap[lingtongId][level]
		if !ook {
			return nil, nil
		}
		if tt.NeedLingzhuLevel <= totalLevel {
			beforeSkillTemp, ok := s.lingzhuSkillMap[lingtongId][level-1]
			if !ok {
				return tt, nil
			}
			return tt, beforeSkillTemp
		}
	}

	return nil, nil
}

//装备
func (s additionSysTemplateService) GetEquipTemplate(id int) *gametemplate.SystemEquipTemplate {
	to, ok := s.equipByIdMap[id]
	if !ok {
		return nil
	}
	return to
}

// 强化
func (s *additionSysTemplateService) GetBodyStrengthenByArg(typ additionsystypes.AdditionSysType, styp additionsystypes.SlotPositionType, lev int32) *gametemplate.SystemStrengthenTemplate {
	typ = typ.ConvertToTemplateAdditionSysType()
	to, ok := s.bodyStrengthenByArgMap[typ][styp][lev]
	if !ok {
		return nil
	}

	return to
}

//套装
func (s *additionSysTemplateService) GetTaoZhuangByArg(typ additionsystypes.AdditionSysType, qualityType itemtypes.ItemQualityType) *gametemplate.SystemTaozhuangTemplate {
	typ = typ.ConvertToTemplateAdditionSysType()
	to, ok := s.taoZhuangByArgMap[typ][qualityType]
	if !ok {
		return nil
	}
	return to
}

// 系统升级
func (s *additionSysTemplateService) GetShengJiByArg(typ additionsystypes.AdditionSysType, lev int32) SystemShengJiCommonTemplate {
	typ = typ.ConvertToTemplateAdditionSysType()
	to, ok := s.shengJiByArgMap[typ][lev]
	if !ok {
		return nil
	}

	return to
}

// 系统化灵
func (s *additionSysTemplateService) GetHuaLingByArg(typ additionsystypes.AdditionSysType, lev int32) (levelTemp *gametemplate.SystemHuaLingTemplate, typeTemp *gametemplate.SystemHuaLingUseTemplate) {
	typ = typ.ConvertToTemplateAdditionSysType()
	levelTemp, ok := s.huaLingByLevelMap[lev]
	if !ok {
		return nil, nil
	}

	typeTemp, ok = s.huaLingByTypeMap[typ]
	if !ok {
		return nil, nil
	}

	return
}

func (s *additionSysTemplateService) GetAwakeLevelByArg(typ additionsystypes.AdditionSysType, advanced int32, level int32) *gametemplate.SystemAwakeLevelTemplate {
	typ = typ.ConvertToTemplateAdditionSysType()
	awakeTemp := s.GetAwakeByArg(typ, advanced)
	if awakeTemp == nil {
		return nil
	}
	awakeLevelTemp := awakeTemp.GetAwakeLevelTemplate(level)
	return awakeLevelTemp
}

//系统觉醒
func (s *additionSysTemplateService) GetAwakeByType(typ additionsystypes.AdditionSysType) *gametemplate.SystemAwakeUseTemplate {
	typ = typ.ConvertToTemplateAdditionSysType()
	awakeTemp, ok := s.awakeByTypeMap[typ]
	if !ok {
		return nil
	}
	return awakeTemp
}

func (s *additionSysTemplateService) GetAwakeByArg(typ additionsystypes.AdditionSysType, advanced int32) *gametemplate.SystemAwakeTemplate {
	typ = typ.ConvertToTemplateAdditionSysType()
	m, ok := s.AwakeByTypeAndAdvancedMap[typ]
	if !ok {
		return nil
	}

	a, ok := m[advanced]
	if !ok {
		return nil
	}

	return a
}

// 获取食用物品数量达到化灵配置
func (s *additionSysTemplateService) GetHuaLingReachByArg(typ additionsystypes.AdditionSysType, curCulLevel int32, num int32) (reachTemplate *gametemplate.SystemHuaLingTemplate, flag bool) {
	typ = typ.ConvertToTemplateAdditionSysType()
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curCulLevel + 1; leftNum > 0; level++ {
		reachTemplate, flag = s.huaLingByLevelMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= reachTemplate.ItemCount
	}
	//次数要满足刚好升级升级
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

// 系统神铸
func (s *additionSysTemplateService) GetShenZhuByArg(pos additionsystypes.SlotPositionType, lev int32) (temp *gametemplate.SystemShenZhuTemplate) {
	temp, ok := s.shenZhuByArgMap[pos][lev]
	if !ok {
		return nil
	}
	return
}

//系统神铸使用道具
func (s *additionSysTemplateService) GetShenZhuUseByType(typ additionsystypes.AdditionSysType) (typeTemp *gametemplate.SystemShenZhuUseTemplate) {
	typ = typ.ConvertToTemplateAdditionSysType()
	typeTemp, ok := s.shenZhuByTypeMap[typ]
	if !ok {
		return nil
	}
	return
}

// 系统通灵
func (s *additionSysTemplateService) GetTongLingByLevel(lev int32) (levelTemp SystemUpgradeCommonTemplate) {
	levelTemp, ok := s.tongLingByLevelMap[lev]
	if !ok {
		return nil
	}
	return
}

//系统通灵使用道具
func (s *additionSysTemplateService) GetTongLingUseByType(typ additionsystypes.AdditionSysType) (typeTemp *gametemplate.SystemTongLingUseTemplate) {
	typ = typ.ConvertToTemplateAdditionSysType()
	typeTemp, ok := s.tongLingByTypeMap[typ]
	if !ok {
		return nil
	}
	return
}

// 获取食用物品数量达到通灵配置
func (s *additionSysTemplateService) GetTongLingReachByArg(typ additionsystypes.AdditionSysType, curCulLevel int32, num int32) (reachTemplate SystemUpgradeCommonTemplate, flag bool) {
	typ = typ.ConvertToTemplateAdditionSysType()
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curCulLevel + 1; leftNum > 0; level++ {
		reachTemplate, flag = s.tongLingByLevelMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= reachTemplate.GetItemCount()
	}
	//次数要满足刚好升级升级
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *additionSysTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &additionSysTemplateService{}
		err = cs.init()
	})
	return err
}

func GetAdditionSysTemplateService() AdditionSysTemplateService {
	return cs
}
