package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//点星系统接口处理
type DianXingTemplateService interface {
	//获取点星系统配置
	GetDianXingTemplateByArg(typ int32, lev int32) *gametemplate.DianXingTemplate
	//获取点星系统配置
	GetDianXingJieFengTemplateByLev(lev int32) *gametemplate.DianXingJieFengTemplate
	//获取点星解封套装系统配置
	GetDianXingJieFengTaoZhuangTemplateByLev(curLevel int32) *gametemplate.DianXingJieFengTaoZhuangTemplate
}

type dianXingTemplateService struct {
	//点星系统配置
	dianXingArgMap map[int32]map[int32]*gametemplate.DianXingTemplate
	//点星系统配置
	dianXingJieFengArgMap map[int32]*gametemplate.DianXingJieFengTemplate
	//点星解封套装系统配置
	dianXingJieFengTaoZhuangMap map[int32]*gametemplate.DianXingJieFengTaoZhuangTemplate
}

//初始化
func (s *dianXingTemplateService) init() error {
	s.dianXingArgMap = make(map[int32]map[int32]*gametemplate.DianXingTemplate)
	s.dianXingJieFengArgMap = make(map[int32]*gametemplate.DianXingJieFengTemplate)
	s.dianXingJieFengTaoZhuangMap = make(map[int32]*gametemplate.DianXingJieFengTaoZhuangTemplate)
	//点星系统
	templateMap := template.GetTemplateService().GetAll((*gametemplate.DianXingTemplate)(nil))
	for _, templateObject := range templateMap {
		dianXingTemplate, _ := templateObject.(*gametemplate.DianXingTemplate)
		tempMap, ok := s.dianXingArgMap[dianXingTemplate.XingPuType]
		if !ok {
			tempMap = make(map[int32]*gametemplate.DianXingTemplate)
			s.dianXingArgMap[dianXingTemplate.XingPuType] = tempMap
		}
		tempMap[dianXingTemplate.Level] = dianXingTemplate
	}
	//点星解封
	jieFengTemplateMap := template.GetTemplateService().GetAll((*gametemplate.DianXingJieFengTemplate)(nil))
	for _, templateObject := range jieFengTemplateMap {
		dianXingJieFengTemplate, _ := templateObject.(*gametemplate.DianXingJieFengTemplate)
		s.dianXingJieFengArgMap[dianXingJieFengTemplate.Level] = dianXingJieFengTemplate
	}
	//点星解封套装
	jieFengTaoZhuangTemplateMap := template.GetTemplateService().GetAll((*gametemplate.DianXingJieFengTaoZhuangTemplate)(nil))
	for _, templateObject := range jieFengTaoZhuangTemplateMap {
		jieFengTaoZhuangTemplate, _ := templateObject.(*gametemplate.DianXingJieFengTaoZhuangTemplate)
		s.dianXingJieFengTaoZhuangMap[jieFengTaoZhuangTemplate.NeedLevel] = jieFengTaoZhuangTemplate
	}

	return nil
}

//获取点星系统配置
func (s *dianXingTemplateService) GetDianXingTemplateByArg(typ int32, lev int32) *gametemplate.DianXingTemplate {
	typMap, ok := s.dianXingArgMap[typ]
	if !ok {
		return nil
	}
	levTemplate, ok := typMap[lev]
	if !ok {
		return nil
	}
	return levTemplate
}

//获取点星系统配置
func (s *dianXingTemplateService) GetDianXingJieFengTemplateByLev(lev int32) *gametemplate.DianXingJieFengTemplate {
	to, ok := s.dianXingJieFengArgMap[lev]
	if !ok {
		return nil
	}
	return to
}

//获取点星解封套装配置
func (s dianXingTemplateService) GetDianXingJieFengTaoZhuangTemplateByLev(curLevel int32) *gametemplate.DianXingJieFengTaoZhuangTemplate {
	maxLevel := int32(0)
	for needLevel, _ := range s.dianXingJieFengTaoZhuangMap {
		if curLevel < needLevel {
			continue
		}
		if maxLevel >= needLevel {
			continue
		}
		maxLevel = needLevel
	}

	temp, ok := s.dianXingJieFengTaoZhuangMap[maxLevel]
	if !ok {
		return nil
	}
	return temp
}

var (
	once sync.Once
	cs   *dianXingTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &dianXingTemplateService{}
		err = cs.init()
	})
	return err
}

func GetDianXingTemplateService() DianXingTemplateService {
	return cs
}
