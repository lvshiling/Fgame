package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/mingge/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//命格接口处理
type MingGeTemplateService interface {
	//获取命格模板
	GetMingGeTempalte(id int32) *gametemplate.MingGeTemplate
	//获取命格合成模板
	GetMingGeSynthesisTemplate(id int32) *gametemplate.MingGeSynthesisTemplate
	//获取命格命盘模板
	GetMingGeMingPanTemplate(subType types.MingGeAllSubType, number int32, star int32) *gametemplate.MingGeMingPanTemplate
	//获取命格命宫模板
	GetMingGeMingGongTemplate(typ types.MingGongType) *gametemplate.MingGeMingGongTemplate
	//获取命格命理模板
	GetMingGeMingLiTemplate(typ types.MingGongType, subType types.MingGongSubType) *gametemplate.MingGeMingLiTemplate
	//获取命格单倍
	GetMingGeDanBeiTemplate(typ types.MingGePropertyType) *gametemplate.MingGeDanBeiTemplate
	//获取命盘模板
	GetMingGeMingPanTemplateById(id int32) *gametemplate.MingGeMingPanTemplate
	//获取属性池概率
	GetMingGeDanBeiRateList(propertyTypeList []types.MingGePropertyType) ([]int64, bool)
	//获取命格补偿
	GetMingGeBuChang(num int32) *gametemplate.MingGeBuchangTemplate
}

type mingGeTemplateService struct {
	//命格模板
	mingGeTemplateMap map[int32]*gametemplate.MingGeTemplate
	//命格合成
	mingGeSynthesisTemplateMap map[int32]*gametemplate.MingGeSynthesisTemplate
	//命格命盘
	mingGeMingPanTemplateMap map[types.MingGeAllSubType]map[int32]map[int32]*gametemplate.MingGeMingPanTemplate
	//命格命宫
	mingGeMingGongTemplateMap map[types.MingGongType]*gametemplate.MingGeMingGongTemplate
	//命格命理
	mingGeMingLiTemplateMap map[types.MingGongType]map[types.MingGongSubType]*gametemplate.MingGeMingLiTemplate
	//命格单倍
	mingGeDanBeiTemplateMap map[types.MingGePropertyType]*gametemplate.MingGeDanBeiTemplate
	//命格命盘
	mingGeMingPanIdMap map[int32]*gametemplate.MingGeMingPanTemplate
	//命格补偿
	mingGeBuChangMap map[int32]*gametemplate.MingGeBuchangTemplate
}

//初始化
func (ts *mingGeTemplateService) init() (err error) {
	err = ts.initMingGe()
	if err != nil {
		return
	}

	err = ts.initMingGeSysthesis()
	if err != nil {
		return
	}

	err = ts.initMingGeMingPan()
	if err != nil {
		return
	}

	err = ts.initMingGeMingGong()
	if err != nil {
		return
	}

	err = ts.initMingGeMingLi()
	if err != nil {
		return
	}

	err = ts.initMingGeDanBei()
	if err != nil {
		return
	}
	err = ts.initMingGeBuchang()
	if err != nil {
		return
	}
	return nil
}

//初始化命格
func (ts *mingGeTemplateService) initMingGe() (err error) {
	ts.mingGeTemplateMap = make(map[int32]*gametemplate.MingGeTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MingGeTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.MingGeTemplate)
		ts.mingGeTemplateMap[int32(tempTemplate.TemplateId())] = tempTemplate
	}
	return
}

//初始化命格合成
func (ts *mingGeTemplateService) initMingGeSysthesis() (err error) {
	ts.mingGeSynthesisTemplateMap = make(map[int32]*gametemplate.MingGeSynthesisTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MingGeSynthesisTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.MingGeSynthesisTemplate)
		ts.mingGeSynthesisTemplateMap[int32(tempTemplate.TemplateId())] = tempTemplate
	}
	return
}

//初始化命盘
func (ts *mingGeTemplateService) initMingGeMingPan() (err error) {
	ts.mingGeMingPanTemplateMap = make(map[types.MingGeAllSubType]map[int32]map[int32]*gametemplate.MingGeMingPanTemplate)
	ts.mingGeMingPanIdMap = make(map[int32]*gametemplate.MingGeMingPanTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MingGeMingPanTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.MingGeMingPanTemplate)

		mingGeMingPanTypeMap, ok := ts.mingGeMingPanTemplateMap[tempTemplate.GetMingPanType()]
		if !ok {
			mingGeMingPanTypeMap = make(map[int32]map[int32]*gametemplate.MingGeMingPanTemplate)
			ts.mingGeMingPanTemplateMap[tempTemplate.GetMingPanType()] = mingGeMingPanTypeMap
		}

		mingGeMingPanNumberMap, ok := mingGeMingPanTypeMap[tempTemplate.Number]
		if !ok {
			mingGeMingPanNumberMap = make(map[int32]*gametemplate.MingGeMingPanTemplate)
			mingGeMingPanTypeMap[tempTemplate.Number] = mingGeMingPanNumberMap
		}

		mingGeMingPanNumberMap[tempTemplate.Star] = tempTemplate

		ts.mingGeMingPanIdMap[int32(tempTemplate.TemplateId())] = tempTemplate
	}
	return
}

//初始化命宫
func (ts *mingGeTemplateService) initMingGeMingGong() (err error) {
	ts.mingGeMingGongTemplateMap = make(map[types.MingGongType]*gametemplate.MingGeMingGongTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MingGeMingGongTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.MingGeMingGongTemplate)
		ts.mingGeMingGongTemplateMap[tempTemplate.GetMingGongType()] = tempTemplate
	}
	return
}

//初始化命理
func (ts *mingGeTemplateService) initMingGeMingLi() (err error) {
	ts.mingGeMingLiTemplateMap = make(map[types.MingGongType]map[types.MingGongSubType]*gametemplate.MingGeMingLiTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MingGeMingLiTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.MingGeMingLiTemplate)
		mingGeMingLiSubTypeMap, ok := ts.mingGeMingLiTemplateMap[tempTemplate.GetMingGongType()]
		if !ok {
			mingGeMingLiSubTypeMap = make(map[types.MingGongSubType]*gametemplate.MingGeMingLiTemplate)
			ts.mingGeMingLiTemplateMap[tempTemplate.GetMingGongType()] = mingGeMingLiSubTypeMap
		}
		mingGeMingLiSubTypeMap[tempTemplate.GetMingGongSubType()] = tempTemplate
	}
	return
}

//初始化命格单倍
func (ts *mingGeTemplateService) initMingGeDanBei() (err error) {
	ts.mingGeDanBeiTemplateMap = make(map[types.MingGePropertyType]*gametemplate.MingGeDanBeiTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MingGeDanBeiTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.MingGeDanBeiTemplate)
		ts.mingGeDanBeiTemplateMap[tempTemplate.GetMingGePropertyType()] = tempTemplate
	}
	return
}

//初始化命格单倍
func (ts *mingGeTemplateService) initMingGeBuchang() (err error) {
	ts.mingGeBuChangMap = make(map[int32]*gametemplate.MingGeBuchangTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MingGeBuchangTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.MingGeBuchangTemplate)
		ts.mingGeBuChangMap[tempTemplate.MinggongCount] = tempTemplate
	}
	return
}

//获取命格模板
func (ts *mingGeTemplateService) GetMingGeTempalte(id int32) *gametemplate.MingGeTemplate {
	temTemplate, ok := ts.mingGeTemplateMap[id]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取命格合成模板
func (ts *mingGeTemplateService) GetMingGeSynthesisTemplate(id int32) *gametemplate.MingGeSynthesisTemplate {
	temTemplate, ok := ts.mingGeSynthesisTemplateMap[id]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取命格命盘模板
func (ts *mingGeTemplateService) GetMingGeMingPanTemplate(subType types.MingGeAllSubType, number int32, star int32) *gametemplate.MingGeMingPanTemplate {
	mingGeMingPanNumberMap, ok := ts.mingGeMingPanTemplateMap[subType]
	if !ok {
		return nil
	}
	mingGeMingPanStarMap, ok := mingGeMingPanNumberMap[number]
	if !ok {
		return nil
	}
	temTemplate, ok := mingGeMingPanStarMap[star]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取命格命宫模板
func (ts *mingGeTemplateService) GetMingGeMingGongTemplate(typ types.MingGongType) *gametemplate.MingGeMingGongTemplate {
	temTemplate, ok := ts.mingGeMingGongTemplateMap[typ]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取命格命理模板
func (ts *mingGeTemplateService) GetMingGeMingLiTemplate(typ types.MingGongType, subType types.MingGongSubType) *gametemplate.MingGeMingLiTemplate {
	mingGeMingLiTemplateMap, ok := ts.mingGeMingLiTemplateMap[typ]
	if !ok {
		return nil
	}
	temTemplate, ok := mingGeMingLiTemplateMap[subType]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取命格单倍
func (ts *mingGeTemplateService) GetMingGeDanBeiTemplate(typ types.MingGePropertyType) *gametemplate.MingGeDanBeiTemplate {
	temTemplate, ok := ts.mingGeDanBeiTemplateMap[typ]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取命盘模板
func (ts *mingGeTemplateService) GetMingGeMingPanTemplateById(id int32) *gametemplate.MingGeMingPanTemplate {
	temTemplate, ok := ts.mingGeMingPanIdMap[id]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取属性池概率
func (ts *mingGeTemplateService) GetMingGeDanBeiRateList(propertyTypeList []types.MingGePropertyType) (rateList []int64, flag bool) {
	for _, propertyType := range propertyTypeList {
		temTemplate, ok := ts.mingGeDanBeiTemplateMap[propertyType]
		if !ok {
			return
		}
		rateList = append(rateList, temTemplate.Rate)
	}
	flag = true
	return
}

//获取属性池概率
func (ts *mingGeTemplateService) GetMingGeBuChang(num int32) *gametemplate.MingGeBuchangTemplate {
	mingGeBuchangTemplate, ok := ts.mingGeBuChangMap[num]
	if !ok {
		return nil
	}
	return mingGeBuchangTemplate
}

var (
	once sync.Once
	cs   *mingGeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &mingGeTemplateService{}
		err = cs.init()
	})
	return err
}

func GetMingGeTemplateService() MingGeTemplateService {
	return cs
}
