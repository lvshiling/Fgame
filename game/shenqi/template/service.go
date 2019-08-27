package template

import (
	"fgame/fgame/core/template"
	shenqitypes "fgame/fgame/game/shenqi/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//接口处理
type ShenQiTemplateService interface {
	//神器配置
	GetShenQiByArg(typ shenqitypes.ShenQiType, lev int32) *gametemplate.ShenQiTemplate
	//神器淬炼配置
	GetShenQiSmeltByArg(typ shenqitypes.ShenQiType, lev int32) *gametemplate.ShenQiCuiLianTemplate
	//神器器灵影响配置
	GetShenQiQiLingEffectByArg(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, ssubType shenqitypes.QiLingSubType) *gametemplate.ShenQiQiLingEffectTemplate
	//神器套装配置
	GetShenQiTaoZhuangMap() map[int]*gametemplate.ShenQiTaoZhuangTemplate
	//神器套装配置
	GetShenQiTaoZhuangById(id int) *gametemplate.ShenQiTaoZhuangTemplate
	//神器碎片升级配置
	GetShenQiDebrisUpByArg(typ shenqitypes.ShenQiType, subType shenqitypes.DebrisType, lev int32) *gametemplate.ShenQiLevelTemplate
	//神器淬炼升级配置
	GetShenQiSmeltUpByArg(typ shenqitypes.ShenQiType, subType shenqitypes.SmeltType, lev int32) *gametemplate.ShenQiCuiLianLevelTemplate
	//神器注灵配置
	GetShenQiZhuLingByArg(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, ssubType shenqitypes.QiLingSubType, lev int32) *gametemplate.ShenQiZhuLingTemplate
	//获取神器技能
	GetShenQiSkillId(typ shenqitypes.ShenQiType, lev int32) int32
}

type shenQiTemplateService struct {
	//神器配置
	shenQiMap map[shenqitypes.ShenQiType]map[int32]*gametemplate.ShenQiTemplate
	//神器淬炼配置
	shenQiSmeltMap map[shenqitypes.ShenQiType]map[int32]*gametemplate.ShenQiCuiLianTemplate
	//神器器灵影响配置
	shenQiQiLingEffectMap map[shenqitypes.ShenQiType]map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*gametemplate.ShenQiQiLingEffectTemplate
	//神器套装配置
	shenQiTaoZhuangIdMap map[int]*gametemplate.ShenQiTaoZhuangTemplate
	//神器碎片升级配置
	shenQiDebrisUpMap map[shenqitypes.ShenQiType]map[shenqitypes.DebrisType]map[int32]*gametemplate.ShenQiLevelTemplate
	//神器淬炼升级配置
	shenQiSmeltUpMap map[shenqitypes.ShenQiType]map[shenqitypes.SmeltType]map[int32]*gametemplate.ShenQiCuiLianLevelTemplate
	//神器注灵配置
	shenQiZhuLingMap map[shenqitypes.ShenQiType]map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]map[int32]*gametemplate.ShenQiZhuLingTemplate
}

//初始化
func (s *shenQiTemplateService) init() error {
	//神器配置
	s.shenQiMap = make(map[shenqitypes.ShenQiType]map[int32]*gametemplate.ShenQiTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ShenQiTemplate)(nil))
	for _, templateObject := range templateMap {
		shenQiTemplate, _ := templateObject.(*gametemplate.ShenQiTemplate)
		tempM, ok := s.shenQiMap[shenQiTemplate.GetShenQiType()]
		if !ok {
			tempM = make(map[int32]*gametemplate.ShenQiTemplate)
			s.shenQiMap[shenQiTemplate.GetShenQiType()] = tempM
		}
		tempM[shenQiTemplate.Level] = shenQiTemplate
	}

	//神器淬炼配置
	s.shenQiSmeltMap = make(map[shenqitypes.ShenQiType]map[int32]*gametemplate.ShenQiCuiLianTemplate)
	smeltTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ShenQiCuiLianTemplate)(nil))
	for _, templateObject := range smeltTemplateMap {
		smeltTemplate, _ := templateObject.(*gametemplate.ShenQiCuiLianTemplate)
		tempM, ok := s.shenQiSmeltMap[smeltTemplate.GetShenQiType()]
		if !ok {
			tempM = make(map[int32]*gametemplate.ShenQiCuiLianTemplate)
			s.shenQiSmeltMap[smeltTemplate.GetShenQiType()] = tempM
		}
		tempM[smeltTemplate.Level] = smeltTemplate
	}

	//神器器灵影响配置
	s.shenQiQiLingEffectMap = make(map[shenqitypes.ShenQiType]map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*gametemplate.ShenQiQiLingEffectTemplate)
	qiLingEffectTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ShenQiQiLingEffectTemplate)(nil))
	for _, templateObject := range qiLingEffectTemplateMap {
		qiLingEffectTemplate, _ := templateObject.(*gametemplate.ShenQiQiLingEffectTemplate)
		tempMM, ok := s.shenQiQiLingEffectMap[qiLingEffectTemplate.GetShenQiType()]
		if !ok {
			tempMM = make(map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*gametemplate.ShenQiQiLingEffectTemplate)
			s.shenQiQiLingEffectMap[qiLingEffectTemplate.GetShenQiType()] = tempMM
		}
		tempM, ok := tempMM[qiLingEffectTemplate.GetQiLingType()]
		if !ok {
			tempM = make(map[shenqitypes.QiLingSubType]*gametemplate.ShenQiQiLingEffectTemplate)
			tempMM[qiLingEffectTemplate.GetQiLingType()] = tempM
		}
		tempM[qiLingEffectTemplate.GetQiLingSubType()] = qiLingEffectTemplate
	}

	//神器套装配置
	s.shenQiTaoZhuangIdMap = make(map[int]*gametemplate.ShenQiTaoZhuangTemplate)
	taoZhuangTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ShenQiTaoZhuangTemplate)(nil))
	for _, templateObject := range taoZhuangTemplateMap {
		taoZhuangTemplate, _ := templateObject.(*gametemplate.ShenQiTaoZhuangTemplate)
		s.shenQiTaoZhuangIdMap[taoZhuangTemplate.TemplateId()] = taoZhuangTemplate
	}

	//神器碎片升级配置
	s.shenQiDebrisUpMap = make(map[shenqitypes.ShenQiType]map[shenqitypes.DebrisType]map[int32]*gametemplate.ShenQiLevelTemplate)
	debrisUpTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ShenQiLevelTemplate)(nil))
	for _, templateObject := range debrisUpTemplateMap {
		debrisUpTemplate, _ := templateObject.(*gametemplate.ShenQiLevelTemplate)
		tempMM, ok := s.shenQiDebrisUpMap[debrisUpTemplate.GetShenQiType()]
		if !ok {
			tempMM = make(map[shenqitypes.DebrisType]map[int32]*gametemplate.ShenQiLevelTemplate)
			s.shenQiDebrisUpMap[debrisUpTemplate.GetShenQiType()] = tempMM
		}
		tempM, ok := tempMM[debrisUpTemplate.GetDebrisType()]
		if !ok {
			tempM = make(map[int32]*gametemplate.ShenQiLevelTemplate)
			tempMM[debrisUpTemplate.GetDebrisType()] = tempM
		}
		tempM[debrisUpTemplate.Level] = debrisUpTemplate
	}

	//神器淬炼升级配置
	s.shenQiSmeltUpMap = make(map[shenqitypes.ShenQiType]map[shenqitypes.SmeltType]map[int32]*gametemplate.ShenQiCuiLianLevelTemplate)
	smeltUpTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ShenQiCuiLianLevelTemplate)(nil))
	for _, templateObject := range smeltUpTemplateMap {
		smeltUpTemplate, _ := templateObject.(*gametemplate.ShenQiCuiLianLevelTemplate)
		tempMM, ok := s.shenQiSmeltUpMap[smeltUpTemplate.GetShenQiType()]
		if !ok {
			tempMM = make(map[shenqitypes.SmeltType]map[int32]*gametemplate.ShenQiCuiLianLevelTemplate)
			s.shenQiSmeltUpMap[smeltUpTemplate.GetShenQiType()] = tempMM
		}
		tempM, ok := tempMM[smeltUpTemplate.GetSmeltType()]
		if !ok {
			tempM = make(map[int32]*gametemplate.ShenQiCuiLianLevelTemplate)
			tempMM[smeltUpTemplate.GetSmeltType()] = tempM
		}
		tempM[smeltUpTemplate.Level] = smeltUpTemplate
	}

	//神器注灵配置
	s.shenQiZhuLingMap = make(map[shenqitypes.ShenQiType]map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]map[int32]*gametemplate.ShenQiZhuLingTemplate)
	zhuLingTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ShenQiZhuLingTemplate)(nil))
	for _, templateObject := range zhuLingTemplateMap {
		zhuLingTemplate, _ := templateObject.(*gametemplate.ShenQiZhuLingTemplate)
		tempMMM, ok := s.shenQiZhuLingMap[zhuLingTemplate.GetShenQiType()]
		if !ok {
			tempMMM = make(map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]map[int32]*gametemplate.ShenQiZhuLingTemplate)
			s.shenQiZhuLingMap[zhuLingTemplate.GetShenQiType()] = tempMMM
		}
		tempMM, ok := tempMMM[zhuLingTemplate.GetQiLingType()]
		if !ok {
			tempMM = make(map[shenqitypes.QiLingSubType]map[int32]*gametemplate.ShenQiZhuLingTemplate)
			tempMMM[zhuLingTemplate.GetQiLingType()] = tempMM
		}
		tempM, ok := tempMM[zhuLingTemplate.GetQiLingSubType()]
		if !ok {
			tempM = make(map[int32]*gametemplate.ShenQiZhuLingTemplate)
			tempMM[zhuLingTemplate.GetQiLingSubType()] = tempM
		}
		tempM[zhuLingTemplate.Level] = zhuLingTemplate
	}

	return nil
}

//神器配置
func (s shenQiTemplateService) GetShenQiByArg(typ shenqitypes.ShenQiType, lev int32) *gametemplate.ShenQiTemplate {
	to, ok := s.shenQiMap[typ][lev]
	if !ok {
		return nil
	}
	return to
}

//神器淬炼配置
func (s shenQiTemplateService) GetShenQiSmeltByArg(typ shenqitypes.ShenQiType, lev int32) *gametemplate.ShenQiCuiLianTemplate {
	to, ok := s.shenQiSmeltMap[typ][lev]
	if !ok {
		return nil
	}
	return to
}

//神器器灵影响配置
func (s shenQiTemplateService) GetShenQiQiLingEffectByArg(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, ssubType shenqitypes.QiLingSubType) *gametemplate.ShenQiQiLingEffectTemplate {
	to, ok := s.shenQiQiLingEffectMap[typ][subType][ssubType]
	if !ok {
		return nil
	}
	return to
}

//神器套装配置
func (s *shenQiTemplateService) GetShenQiTaoZhuangMap() map[int]*gametemplate.ShenQiTaoZhuangTemplate {
	return s.shenQiTaoZhuangIdMap
}

//神器套装配置
func (s *shenQiTemplateService) GetShenQiTaoZhuangById(id int) *gametemplate.ShenQiTaoZhuangTemplate {
	to, ok := s.shenQiTaoZhuangIdMap[id]
	if !ok {
		return nil
	}
	return to
}

//神器碎片升级配置
func (s shenQiTemplateService) GetShenQiDebrisUpByArg(typ shenqitypes.ShenQiType, subType shenqitypes.DebrisType, lev int32) *gametemplate.ShenQiLevelTemplate {
	to, ok := s.shenQiDebrisUpMap[typ][subType][lev]
	if !ok {
		return nil
	}
	return to
}

//神器淬炼升级配置
func (s shenQiTemplateService) GetShenQiSmeltUpByArg(typ shenqitypes.ShenQiType, subType shenqitypes.SmeltType, lev int32) *gametemplate.ShenQiCuiLianLevelTemplate {
	to, ok := s.shenQiSmeltUpMap[typ][subType][lev]
	if !ok {
		return nil
	}
	return to
}

//神器注灵配置
func (s shenQiTemplateService) GetShenQiZhuLingByArg(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, ssubType shenqitypes.QiLingSubType, lev int32) *gametemplate.ShenQiZhuLingTemplate {
	to, ok := s.shenQiZhuLingMap[typ][subType][ssubType][lev]
	if !ok {
		return nil
	}
	return to
}

// 获取神器技能
func (s *shenQiTemplateService) GetShenQiSkillId(typ shenqitypes.ShenQiType, lev int32) (skillId int32) {
	temp, ok := s.shenQiMap[typ][lev]
	if !ok {
		return
	}
	skillId = temp.SkillId
	return
}

var (
	once sync.Once
	cs   *shenQiTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shenQiTemplateService{}
		err = cs.init()
	})
	return err
}

func GetShenQiTemplateService() ShenQiTemplateService {
	return cs
}
