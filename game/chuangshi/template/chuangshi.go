package template

import (
	"fgame/fgame/core/template"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

type ChuangShiTemplateService interface {
	// 获得创世之战预告模板
	GetChuangShiYuGaoTemplate() *gametemplate.ChuangShiYuGaoTemplate
	// 获取阵营模板
	GetChuangShiCampTempAll() map[chuangshitypes.ChuangShiCampType]*gametemplate.ChuangShiZhenyingTemplate
	//获取城池模板
	GetAllChuangShiCityTemplate() map[chuangshitypes.ChuangShiCampType]map[chuangshitypes.ChuangShiCityType]map[int32]*gametemplate.ChuangShiCityTemplate
	GetChuangShiCityTemp(campType chuangshitypes.ChuangShiCampType, cityType chuangshitypes.ChuangShiCityType, index int32) *gametemplate.ChuangShiCityTemplate
	//获取城池建设模板
	GetAllChuangShiChengFangTemp() map[chuangshitypes.ChuangShiCityJianSheType]*gametemplate.ChuangShiChengFangTemplate
	GetChuangShiChengFangTemp(jianSheType chuangshitypes.ChuangShiCityJianSheType) *gametemplate.ChuangShiChengFangTemplate
	// 创世常量
	GetChuangshiConstantTemp() *gametemplate.ChuangShiConstantTemplate
	// 获取创世官职模板
	GetChuangShiGuanZhiTemplate(level int32) *gametemplate.ChuangShiGuanZhiTemplate

	// 创世城战模板
	GetChuangShiWarTemp(campType chuangshitypes.ChuangShiCampType) *gametemplate.ChuangShiWarTemplate
}

type chuangShiTemplateService struct {
	//预告模板
	chuangShiYuGaoTemp *gametemplate.ChuangShiYuGaoTemplate
	//阵营模板
	campTempMap map[chuangshitypes.ChuangShiCampType]*gametemplate.ChuangShiZhenyingTemplate
	//城池模板
	campCityTemplateMap map[chuangshitypes.ChuangShiCampType]map[chuangshitypes.ChuangShiCityType]map[int32]*gametemplate.ChuangShiCityTemplate
	//城池建设模板
	cityChengFangTempMap map[chuangshitypes.ChuangShiCityJianSheType]*gametemplate.ChuangShiChengFangTemplate
	//创世常量模板
	chuangshiConstantTemp *gametemplate.ChuangShiConstantTemplate
	//创世官职模板
	chuangShiGuanZhiMap map[int32]*gametemplate.ChuangShiGuanZhiTemplate
	// 创世城战模板
	chuangShiWarMap map[chuangshitypes.ChuangShiCampType]*gametemplate.ChuangShiWarTemplate
}

func (t *chuangShiTemplateService) init() (err error) {
	err = t.initYuGao()
	if err != nil {
		return
	}
	err = t.initChuangShiCamps()
	if err != nil {
		return
	}
	err = t.initChuangShiCities()
	if err != nil {
		return
	}
	err = t.initChuangShiConstant()
	if err != nil {
		return
	}
	err = t.initChuangShiCityJianShe()
	if err != nil {
		return
	}
	err = t.initChuangShiGuanZhi()
	if err != nil {
		return
	}
	err = t.initChuangShiWar()
	if err != nil {
		return
	}

	return
}

func (t *chuangShiTemplateService) initYuGao() (err error) {
	yugaoTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ChuangShiYuGaoTemplate)(nil))
	if len(yugaoTemplateMap) != 1 {
		return fmt.Errorf("chuangshi: 创世之战预告模板配置应该只有一条")
	}
	for _, temp := range yugaoTemplateMap {
		yugaoTemp, _ := temp.(*gametemplate.ChuangShiYuGaoTemplate)
		t.chuangShiYuGaoTemp = yugaoTemp
	}

	return
}

func (t *chuangShiTemplateService) initChuangShiCamps() (err error) {
	t.campTempMap = make(map[chuangshitypes.ChuangShiCampType]*gametemplate.ChuangShiZhenyingTemplate)
	tempMap := template.GetTemplateService().GetAll((*gametemplate.ChuangShiZhenyingTemplate)(nil))
	for _, temp := range tempMap {
		campTemp, _ := temp.(*gametemplate.ChuangShiZhenyingTemplate)
		t.campTempMap[campTemp.GetCampType()] = campTemp
	}
	return
}

func (t *chuangShiTemplateService) initChuangShiGuanZhi() (err error) {
	t.chuangShiGuanZhiMap = make(map[int32]*gametemplate.ChuangShiGuanZhiTemplate)
	tempMap := template.GetTemplateService().GetAll((*gametemplate.ChuangShiGuanZhiTemplate)(nil))
	for _, temp := range tempMap {
		guanzhiTemp, _ := temp.(*gametemplate.ChuangShiGuanZhiTemplate)
		t.chuangShiGuanZhiMap[guanzhiTemp.Level] = guanzhiTemp
	}
	return
}

func (t *chuangShiTemplateService) initChuangShiWar() (err error) {
	t.chuangShiWarMap = make(map[chuangshitypes.ChuangShiCampType]*gametemplate.ChuangShiWarTemplate)
	tempMap := template.GetTemplateService().GetAll((*gametemplate.ChuangShiWarTemplate)(nil))
	for _, temp := range tempMap {
		warTemp, _ := temp.(*gametemplate.ChuangShiWarTemplate)
		t.chuangShiWarMap[warTemp.GetCampType()] = warTemp
	}
	return
}

func (t *chuangShiTemplateService) initChuangShiCities() (err error) {
	t.campCityTemplateMap = make(map[chuangshitypes.ChuangShiCampType]map[chuangshitypes.ChuangShiCityType]map[int32]*gametemplate.ChuangShiCityTemplate)
	cityTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ChuangShiCityTemplate)(nil))

	for _, temp := range cityTemplateMap {
		cityTemplate, _ := temp.(*gametemplate.ChuangShiCityTemplate)
		tempCampTemplateMap, ok := t.campCityTemplateMap[cityTemplate.GetCamp()]
		if !ok {
			tempCampTemplateMap = make(map[chuangshitypes.ChuangShiCityType]map[int32]*gametemplate.ChuangShiCityTemplate)
			t.campCityTemplateMap[cityTemplate.GetCamp()] = tempCampTemplateMap
		}
		tempCityTemplateMap, ok := tempCampTemplateMap[cityTemplate.GetCityType()]
		if !ok {
			tempCityTemplateMap = make(map[int32]*gametemplate.ChuangShiCityTemplate)
			tempCampTemplateMap[cityTemplate.GetCityType()] = tempCityTemplateMap
		}
		tempCityTemplateMap[cityTemplate.SuoyinId] = cityTemplate
	}

	return
}

func (t *chuangShiTemplateService) initChuangShiCityJianShe() (err error) {
	t.cityChengFangTempMap = make(map[chuangshitypes.ChuangShiCityJianSheType]*gametemplate.ChuangShiChengFangTemplate)
	toMap := template.GetTemplateService().GetAll((*gametemplate.ChuangShiChengFangTemplate)(nil))

	for _, to := range toMap {
		changFangTemp, _ := to.(*gametemplate.ChuangShiChengFangTemplate)
		t.cityChengFangTempMap[changFangTemp.GetJianSheType()] = changFangTemp
	}

	return
}

func (t *chuangShiTemplateService) initChuangShiConstant() (err error) {
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ChuangShiConstantTemplate)(nil))
	if len(templateMap) != 1 {
		return fmt.Errorf("创世常量模板有且只有一条")
	}
	constantTemplate, _ := templateMap[1].(*gametemplate.ChuangShiConstantTemplate)
	t.chuangshiConstantTemp = constantTemplate

	return
}

func (t *chuangShiTemplateService) GetChuangShiCampTempAll() map[chuangshitypes.ChuangShiCampType]*gametemplate.ChuangShiZhenyingTemplate {
	return t.campTempMap
}

func (t *chuangShiTemplateService) GetAllChuangShiCityTemplate() map[chuangshitypes.ChuangShiCampType]map[chuangshitypes.ChuangShiCityType]map[int32]*gametemplate.ChuangShiCityTemplate {
	return t.campCityTemplateMap
}

func (t *chuangShiTemplateService) GetChuangShiCityTemp(campType chuangshitypes.ChuangShiCampType, cityType chuangshitypes.ChuangShiCityType, index int32) *gametemplate.ChuangShiCityTemplate {
	subMap, ok := t.campCityTemplateMap[campType]
	if !ok {
		return nil
	}

	subOfSubMap, ok := subMap[cityType]
	if !ok {
		return nil
	}

	temp, ok := subOfSubMap[index]
	if !ok {
		return nil
	}

	return temp
}

func (t *chuangShiTemplateService) GetAllChuangShiChengFangTemp() map[chuangshitypes.ChuangShiCityJianSheType]*gametemplate.ChuangShiChengFangTemplate {
	return t.cityChengFangTempMap
}

func (t *chuangShiTemplateService) GetChuangShiChengFangTemp(jianSheType chuangshitypes.ChuangShiCityJianSheType) *gametemplate.ChuangShiChengFangTemplate {
	temp, ok := t.cityChengFangTempMap[jianSheType]
	if !ok {
		return nil
	}

	return temp
}

func (t *chuangShiTemplateService) GetChuangshiConstantTemp() *gametemplate.ChuangShiConstantTemplate {
	return t.chuangshiConstantTemp
}

func (t *chuangShiTemplateService) GetChuangShiYuGaoTemplate() *gametemplate.ChuangShiYuGaoTemplate {
	return t.chuangShiYuGaoTemp
}

func (t *chuangShiTemplateService) GetChuangShiGuanZhiTemplate(level int32) *gametemplate.ChuangShiGuanZhiTemplate {
	temp, ok := t.chuangShiGuanZhiMap[level]
	if !ok {
		return nil
	}
	return temp
}

func (t *chuangShiTemplateService) GetChuangShiWarTemp(campType chuangshitypes.ChuangShiCampType) *gametemplate.ChuangShiWarTemplate {
	temp, ok := t.chuangShiWarMap[campType]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	s    *chuangShiTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &chuangShiTemplateService{}
		err = s.init()
	})
	return
}

func GetChuangShiTemplateService() ChuangShiTemplateService {
	return s
}
