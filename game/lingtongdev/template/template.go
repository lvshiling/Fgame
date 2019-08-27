package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/lingtongdev/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//灵童养成模板接口处理
type LingTongDevTemplateService interface {
	//获取灵童养成模板进阶配置
	GetLingTongDevNumber(classType types.LingTongDevSysType, number int32) gametemplate.LingTongDevTemplate
	//获取灵童养成模板配置
	GetLingTongDevTemplate(classType types.LingTongDevSysType, seqId int) gametemplate.LingTongDevTemplate
	//获取食幻化丹配置
	GetLingTongDevHuanHuaTemplate(classType types.LingTongDevSysType, level int32) gametemplate.LingTongDevHuanHuaTemplate
	//获取通灵配置
	GetLingTongDevTongLingTemplate(classType types.LingTongDevSysType, level int32) gametemplate.LingTongDevTongLingTemplate
	//获取培养配置
	GetLingTongDevPeiYangTemplate(classType types.LingTongDevSysType, level int32) gametemplate.LingTongDevPeiYangTemplate
	//吃幻化丹升级
	GetLingTongDevEatHuanHuanTemplate(classType types.LingTongDevSysType, curLevel int32, num int32) (gametemplate.LingTongDevHuanHuaTemplate, bool)
	//灵童养成通灵升级
	GetLingTongDevTongLingUpgrade(classType types.LingTongDevSysType, curLevel int32, num int32) (gametemplate.LingTongDevTongLingTemplate, bool)
	//灵童养成培养
	GetLingTongDevPeiYangUpgrade(classType types.LingTongDevSysType, curLevel int32, num int32) (gametemplate.LingTongDevPeiYangTemplate, bool)

	//RandomLingTongDevTemplate() *gametemplate.LingTongDevTemplate
}

type lingTongDevTemplateService struct {
	//灵童养成模板进阶配置
	lingTongDevNumberMap map[types.LingTongDevSysType]map[int32]gametemplate.LingTongDevTemplate
	//灵童养成模板配置
	lingTongDevMap map[types.LingTongDevSysType]map[int]gametemplate.LingTongDevTemplate
	//灵童养成模板幻化配置
	huanHuaMap map[types.LingTongDevSysType]map[int32]gametemplate.LingTongDevHuanHuaTemplate
	//灵童养成模板通灵配置
	tongLingMap map[types.LingTongDevSysType]map[int32]gametemplate.LingTongDevTongLingTemplate
	//灵童养成模板培养配置
	peiYangMap map[types.LingTongDevSysType]map[int32]gametemplate.LingTongDevPeiYangTemplate
	//lingTongDevList []*gametemplate.LingTongDevTemplate
}

func (ws *lingTongDevTemplateService) initMount() (err error) {
	//灵骑模板
	numberMap, ok := ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingQi]
	if !ok {
		numberMap = make(map[int32]gametemplate.LingTongDevTemplate)
		ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingQi] = numberMap
	}
	devMap, ok := ws.lingTongDevMap[types.LingTongDevSysTypeLingQi]
	if !ok {
		devMap = make(map[int]gametemplate.LingTongDevTemplate)
		ws.lingTongDevMap[types.LingTongDevSysTypeLingQi] = devMap
	}
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongMountTemplate)(nil))
	for _, templateObject := range templateMap {
		devTemplate, _ := templateObject.(*gametemplate.LingTongMountTemplate)
		devMap[devTemplate.TemplateId()] = devTemplate

		typ := devTemplate.GetType()
		if typ == types.LingTongDevTypeAdvanced {
			numberMap[devTemplate.Number] = devTemplate
		}
	}

	//幻化
	huanHuaMap, ok := ws.huanHuaMap[types.LingTongDevSysTypeLingQi]
	if !ok {
		huanHuaMap = make(map[int32]gametemplate.LingTongDevHuanHuaTemplate)
		ws.huanHuaMap[types.LingTongDevSysTypeLingQi] = huanHuaMap
	}
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongMountHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.LingTongMountHuanHuaTemplate)
		huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	//培养
	peiYangMap, ok := ws.peiYangMap[types.LingTongDevSysTypeLingQi]
	if !ok {
		peiYangMap = make(map[int32]gametemplate.LingTongDevPeiYangTemplate)
		ws.peiYangMap[types.LingTongDevSysTypeLingQi] = peiYangMap
	}
	peiYangtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongMountPeiYangTemplate)(nil))
	for _, templateObject := range peiYangtemplateMap {
		peiYangTemplate, _ := templateObject.(*gametemplate.LingTongMountPeiYangTemplate)
		peiYangMap[peiYangTemplate.Level] = peiYangTemplate
	}

	//通灵
	tongLingMap, ok := ws.tongLingMap[types.LingTongDevSysTypeLingQi]
	if !ok {
		tongLingMap = make(map[int32]gametemplate.LingTongDevTongLingTemplate)
		ws.tongLingMap[types.LingTongDevSysTypeLingQi] = tongLingMap
	}
	tongLingtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongMountTongLingTemplate)(nil))
	for _, templateObject := range tongLingtemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.LingTongMountTongLingTemplate)
		tongLingMap[tongLingTemplate.Level] = tongLingTemplate
	}
	return
}

func (ws *lingTongDevTemplateService) initWing() (err error) {
	//灵翼模板
	numberMap, ok := ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingYi]
	if !ok {
		numberMap = make(map[int32]gametemplate.LingTongDevTemplate)
		ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingYi] = numberMap
	}
	devMap, ok := ws.lingTongDevMap[types.LingTongDevSysTypeLingYi]
	if !ok {
		devMap = make(map[int]gametemplate.LingTongDevTemplate)
		ws.lingTongDevMap[types.LingTongDevSysTypeLingYi] = devMap
	}
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongWingTemplate)(nil))
	for _, templateObject := range templateMap {
		devTemplate, _ := templateObject.(*gametemplate.LingTongWingTemplate)
		devMap[devTemplate.TemplateId()] = devTemplate

		typ := devTemplate.GetType()
		if typ == types.LingTongDevTypeAdvanced {
			numberMap[devTemplate.Number] = devTemplate
		}
	}

	//幻化
	huanHuaMap, ok := ws.huanHuaMap[types.LingTongDevSysTypeLingYi]
	if !ok {
		huanHuaMap = make(map[int32]gametemplate.LingTongDevHuanHuaTemplate)
		ws.huanHuaMap[types.LingTongDevSysTypeLingYi] = huanHuaMap
	}
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongWingHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.LingTongWingHuanHuaTemplate)
		huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	//培养
	peiYangMap, ok := ws.peiYangMap[types.LingTongDevSysTypeLingYi]
	if !ok {
		peiYangMap = make(map[int32]gametemplate.LingTongDevPeiYangTemplate)
		ws.peiYangMap[types.LingTongDevSysTypeLingYi] = peiYangMap
	}
	peiYangtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongWingPeiYangTemplate)(nil))
	for _, templateObject := range peiYangtemplateMap {
		peiYangTemplate, _ := templateObject.(*gametemplate.LingTongWingPeiYangTemplate)
		peiYangMap[peiYangTemplate.Level] = peiYangTemplate
	}

	//通灵
	tongLingMap, ok := ws.tongLingMap[types.LingTongDevSysTypeLingYi]
	if !ok {
		tongLingMap = make(map[int32]gametemplate.LingTongDevTongLingTemplate)
		ws.tongLingMap[types.LingTongDevSysTypeLingYi] = tongLingMap
	}
	tongLingtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongWingTongLingTemplate)(nil))
	for _, templateObject := range tongLingtemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.LingTongWingTongLingTemplate)
		tongLingMap[tongLingTemplate.Level] = tongLingTemplate
	}
	return
}

func (ws *lingTongDevTemplateService) initLingYu() (err error) {
	//灵域模板
	numberMap, ok := ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingYu]
	if !ok {
		numberMap = make(map[int32]gametemplate.LingTongDevTemplate)
		ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingYu] = numberMap
	}
	devMap, ok := ws.lingTongDevMap[types.LingTongDevSysTypeLingYu]
	if !ok {
		devMap = make(map[int]gametemplate.LingTongDevTemplate)
		ws.lingTongDevMap[types.LingTongDevSysTypeLingYu] = devMap
	}
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongLingYuTemplate)(nil))
	for _, templateObject := range templateMap {
		devTemplate, _ := templateObject.(*gametemplate.LingTongLingYuTemplate)
		devMap[devTemplate.TemplateId()] = devTemplate

		typ := devTemplate.GetType()
		if typ == types.LingTongDevTypeAdvanced {
			numberMap[devTemplate.Number] = devTemplate
		}
	}

	//幻化
	huanHuaMap, ok := ws.huanHuaMap[types.LingTongDevSysTypeLingYu]
	if !ok {
		huanHuaMap = make(map[int32]gametemplate.LingTongDevHuanHuaTemplate)
		ws.huanHuaMap[types.LingTongDevSysTypeLingYu] = huanHuaMap
	}
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongLingYuHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.LingTongLingYuHuanHuaTemplate)
		huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	//培养
	peiYangMap, ok := ws.peiYangMap[types.LingTongDevSysTypeLingYu]
	if !ok {
		peiYangMap = make(map[int32]gametemplate.LingTongDevPeiYangTemplate)
		ws.peiYangMap[types.LingTongDevSysTypeLingYu] = peiYangMap
	}
	peiYangtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongLingYuPeiYangTemplate)(nil))
	for _, templateObject := range peiYangtemplateMap {
		peiYangTemplate, _ := templateObject.(*gametemplate.LingTongLingYuPeiYangTemplate)
		peiYangMap[peiYangTemplate.Level] = peiYangTemplate
	}

	//通灵
	tongLingMap, ok := ws.tongLingMap[types.LingTongDevSysTypeLingYu]
	if !ok {
		tongLingMap = make(map[int32]gametemplate.LingTongDevTongLingTemplate)
		ws.tongLingMap[types.LingTongDevSysTypeLingYu] = tongLingMap
	}
	tongLingtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongLingYuTongLingTemplate)(nil))
	for _, templateObject := range tongLingtemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.LingTongLingYuTongLingTemplate)
		tongLingMap[tongLingTemplate.Level] = tongLingTemplate
	}
	return
}

func (ws *lingTongDevTemplateService) initShenFa() (err error) {
	//灵身模板
	numberMap, ok := ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingShen]
	if !ok {
		numberMap = make(map[int32]gametemplate.LingTongDevTemplate)
		ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingShen] = numberMap
	}
	devMap, ok := ws.lingTongDevMap[types.LingTongDevSysTypeLingShen]
	if !ok {
		devMap = make(map[int]gametemplate.LingTongDevTemplate)
		ws.lingTongDevMap[types.LingTongDevSysTypeLingShen] = devMap
	}
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongShenFaTemplate)(nil))
	for _, templateObject := range templateMap {
		devTemplate, _ := templateObject.(*gametemplate.LingTongShenFaTemplate)
		devMap[devTemplate.TemplateId()] = devTemplate

		typ := devTemplate.GetType()
		if typ == types.LingTongDevTypeAdvanced {
			numberMap[devTemplate.Number] = devTemplate
		}
	}

	//幻化
	huanHuaMap, ok := ws.huanHuaMap[types.LingTongDevSysTypeLingShen]
	if !ok {
		huanHuaMap = make(map[int32]gametemplate.LingTongDevHuanHuaTemplate)
		ws.huanHuaMap[types.LingTongDevSysTypeLingShen] = huanHuaMap
	}
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongShenFaHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.LingTongShenFaHuanHuaTemplate)
		huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	//培养
	peiYangMap, ok := ws.peiYangMap[types.LingTongDevSysTypeLingShen]
	if !ok {
		peiYangMap = make(map[int32]gametemplate.LingTongDevPeiYangTemplate)
		ws.peiYangMap[types.LingTongDevSysTypeLingShen] = peiYangMap
	}
	peiYangtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongShenFaPeiYangTemplate)(nil))
	for _, templateObject := range peiYangtemplateMap {
		peiYangTemplate, _ := templateObject.(*gametemplate.LingTongShenFaPeiYangTemplate)
		peiYangMap[peiYangTemplate.Level] = peiYangTemplate
	}

	//通灵
	tongLingMap, ok := ws.tongLingMap[types.LingTongDevSysTypeLingShen]
	if !ok {
		tongLingMap = make(map[int32]gametemplate.LingTongDevTongLingTemplate)
		ws.tongLingMap[types.LingTongDevSysTypeLingShen] = tongLingMap
	}
	tongLingtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongShenFaTongLingTemplate)(nil))
	for _, templateObject := range tongLingtemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.LingTongShenFaTongLingTemplate)
		tongLingMap[tongLingTemplate.Level] = tongLingTemplate
	}
	return
}

func (ws *lingTongDevTemplateService) initWeapon() (err error) {
	//灵兵模板
	numberMap, ok := ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingBing]
	if !ok {
		numberMap = make(map[int32]gametemplate.LingTongDevTemplate)
		ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingBing] = numberMap
	}
	devMap, ok := ws.lingTongDevMap[types.LingTongDevSysTypeLingBing]
	if !ok {
		devMap = make(map[int]gametemplate.LingTongDevTemplate)
		ws.lingTongDevMap[types.LingTongDevSysTypeLingBing] = devMap
	}
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongWeaponTemplate)(nil))
	for _, templateObject := range templateMap {
		devTemplate, _ := templateObject.(*gametemplate.LingTongWeaponTemplate)
		devMap[devTemplate.TemplateId()] = devTemplate

		typ := devTemplate.GetType()
		if typ == types.LingTongDevTypeAdvanced {
			numberMap[devTemplate.Number] = devTemplate
		}
	}

	//幻化
	huanHuaMap, ok := ws.huanHuaMap[types.LingTongDevSysTypeLingBing]
	if !ok {
		huanHuaMap = make(map[int32]gametemplate.LingTongDevHuanHuaTemplate)
		ws.huanHuaMap[types.LingTongDevSysTypeLingBing] = huanHuaMap
	}
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongWeaponHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.LingTongWeaponHuanHuaTemplate)
		huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	//培养
	peiYangMap, ok := ws.peiYangMap[types.LingTongDevSysTypeLingBing]
	if !ok {
		peiYangMap = make(map[int32]gametemplate.LingTongDevPeiYangTemplate)
		ws.peiYangMap[types.LingTongDevSysTypeLingBing] = peiYangMap
	}
	peiYangtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongWeaponPeiYangTemplate)(nil))
	for _, templateObject := range peiYangtemplateMap {
		peiYangTemplate, _ := templateObject.(*gametemplate.LingTongWeaponPeiYangTemplate)
		peiYangMap[peiYangTemplate.Level] = peiYangTemplate
	}

	//通灵
	tongLingMap, ok := ws.tongLingMap[types.LingTongDevSysTypeLingBing]
	if !ok {
		tongLingMap = make(map[int32]gametemplate.LingTongDevTongLingTemplate)
		ws.tongLingMap[types.LingTongDevSysTypeLingBing] = tongLingMap
	}
	tongLingtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongWeaponTongLingTemplate)(nil))
	for _, templateObject := range tongLingtemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.LingTongWeaponTongLingTemplate)
		tongLingMap[tongLingTemplate.Level] = tongLingTemplate
	}
	return
}

func (ws *lingTongDevTemplateService) initXianTi() (err error) {
	//灵体模板
	numberMap, ok := ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingTi]
	if !ok {
		numberMap = make(map[int32]gametemplate.LingTongDevTemplate)
		ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingTi] = numberMap
	}
	devMap, ok := ws.lingTongDevMap[types.LingTongDevSysTypeLingTi]
	if !ok {
		devMap = make(map[int]gametemplate.LingTongDevTemplate)
		ws.lingTongDevMap[types.LingTongDevSysTypeLingTi] = devMap
	}
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongXianTiTemplate)(nil))
	for _, templateObject := range templateMap {
		devTemplate, _ := templateObject.(*gametemplate.LingTongXianTiTemplate)
		devMap[devTemplate.TemplateId()] = devTemplate

		typ := devTemplate.GetType()
		if typ == types.LingTongDevTypeAdvanced {
			numberMap[devTemplate.Number] = devTemplate
		}
	}

	//幻化
	huanHuaMap, ok := ws.huanHuaMap[types.LingTongDevSysTypeLingTi]
	if !ok {
		huanHuaMap = make(map[int32]gametemplate.LingTongDevHuanHuaTemplate)
		ws.huanHuaMap[types.LingTongDevSysTypeLingTi] = huanHuaMap
	}
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongXianTiHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.LingTongXianTiHuanHuaTemplate)
		huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	//培养
	peiYangMap, ok := ws.peiYangMap[types.LingTongDevSysTypeLingTi]
	if !ok {
		peiYangMap = make(map[int32]gametemplate.LingTongDevPeiYangTemplate)
		ws.peiYangMap[types.LingTongDevSysTypeLingTi] = peiYangMap
	}
	peiYangtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongXianTiPeiYangTemplate)(nil))
	for _, templateObject := range peiYangtemplateMap {
		peiYangTemplate, _ := templateObject.(*gametemplate.LingTongXianTiPeiYangTemplate)
		peiYangMap[peiYangTemplate.Level] = peiYangTemplate
	}

	//通灵
	tongLingMap, ok := ws.tongLingMap[types.LingTongDevSysTypeLingTi]
	if !ok {
		tongLingMap = make(map[int32]gametemplate.LingTongDevTongLingTemplate)
		ws.tongLingMap[types.LingTongDevSysTypeLingTi] = tongLingMap
	}
	tongLingtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongXianTiTongLingTemplate)(nil))
	for _, templateObject := range tongLingtemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.LingTongXianTiTongLingTemplate)
		tongLingMap[tongLingTemplate.Level] = tongLingTemplate
	}
	return
}

func (ws *lingTongDevTemplateService) initFaBao() (err error) {
	//灵体模板
	numberMap, ok := ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingBao]
	if !ok {
		numberMap = make(map[int32]gametemplate.LingTongDevTemplate)
		ws.lingTongDevNumberMap[types.LingTongDevSysTypeLingBao] = numberMap
	}
	devMap, ok := ws.lingTongDevMap[types.LingTongDevSysTypeLingBao]
	if !ok {
		devMap = make(map[int]gametemplate.LingTongDevTemplate)
		ws.lingTongDevMap[types.LingTongDevSysTypeLingBao] = devMap
	}
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongFaBaoTemplate)(nil))
	for _, templateObject := range templateMap {
		devTemplate, _ := templateObject.(*gametemplate.LingTongFaBaoTemplate)
		devMap[devTemplate.TemplateId()] = devTemplate

		typ := devTemplate.GetType()
		if typ == types.LingTongDevTypeAdvanced {
			numberMap[devTemplate.Number] = devTemplate
		}
	}

	//幻化
	huanHuaMap, ok := ws.huanHuaMap[types.LingTongDevSysTypeLingBao]
	if !ok {
		huanHuaMap = make(map[int32]gametemplate.LingTongDevHuanHuaTemplate)
		ws.huanHuaMap[types.LingTongDevSysTypeLingBao] = huanHuaMap
	}
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongFaBaoHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.LingTongFaBaoHuanHuaTemplate)
		huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	//培养
	peiYangMap, ok := ws.peiYangMap[types.LingTongDevSysTypeLingBao]
	if !ok {
		peiYangMap = make(map[int32]gametemplate.LingTongDevPeiYangTemplate)
		ws.peiYangMap[types.LingTongDevSysTypeLingBao] = peiYangMap
	}
	peiYangtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongFaBaoPeiYangTemplate)(nil))
	for _, templateObject := range peiYangtemplateMap {
		peiYangTemplate, _ := templateObject.(*gametemplate.LingTongFaBaoPeiYangTemplate)
		peiYangMap[peiYangTemplate.Level] = peiYangTemplate
	}

	//通灵
	tongLingMap, ok := ws.tongLingMap[types.LingTongDevSysTypeLingBao]
	if !ok {
		tongLingMap = make(map[int32]gametemplate.LingTongDevTongLingTemplate)
		ws.tongLingMap[types.LingTongDevSysTypeLingBao] = tongLingMap
	}
	tongLingtemplateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongFaBaoTongLingTemplate)(nil))
	for _, templateObject := range tongLingtemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.LingTongFaBaoTongLingTemplate)
		tongLingMap[tongLingTemplate.Level] = tongLingTemplate
	}
	return
}

//初始化
func (ws *lingTongDevTemplateService) init() (err error) {
	ws.lingTongDevNumberMap = make(map[types.LingTongDevSysType]map[int32]gametemplate.LingTongDevTemplate)
	ws.lingTongDevMap = make(map[types.LingTongDevSysType]map[int]gametemplate.LingTongDevTemplate)
	ws.huanHuaMap = make(map[types.LingTongDevSysType]map[int32]gametemplate.LingTongDevHuanHuaTemplate)
	ws.tongLingMap = make(map[types.LingTongDevSysType]map[int32]gametemplate.LingTongDevTongLingTemplate)
	ws.peiYangMap = make(map[types.LingTongDevSysType]map[int32]gametemplate.LingTongDevPeiYangTemplate)

	err = ws.initMount()
	if err != nil {
		return
	}

	err = ws.initWing()
	if err != nil {
		return
	}

	err = ws.initFaBao()
	if err != nil {
		return
	}

	err = ws.initWeapon()
	if err != nil {
		return
	}

	err = ws.initXianTi()
	if err != nil {
		return
	}

	err = ws.initShenFa()
	if err != nil {
		return
	}

	err = ws.initLingYu()
	if err != nil {
		return
	}

	return nil
}

//获取灵童养成模板进阶配置
func (ws *lingTongDevTemplateService) GetLingTongDevNumber(classType types.LingTongDevSysType, number int32) gametemplate.LingTongDevTemplate {
	devNumberMap, ok := ws.lingTongDevNumberMap[classType]
	if !ok {
		return nil
	}
	to, ok := devNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//获取灵童养成模板配置
func (ws *lingTongDevTemplateService) GetLingTongDevTemplate(classType types.LingTongDevSysType, seqId int) gametemplate.LingTongDevTemplate {
	devMap, ok := ws.lingTongDevMap[classType]
	if !ok {
		return nil
	}
	to, ok := devMap[seqId]
	if !ok {
		return nil
	}
	return to
}

//获取幻化配置
func (ws *lingTongDevTemplateService) GetLingTongDevHuanHuaTemplate(classType types.LingTongDevSysType, level int32) gametemplate.LingTongDevHuanHuaTemplate {
	huanHuaMap, ok := ws.huanHuaMap[classType]
	if !ok {
		return nil
	}
	to, ok := huanHuaMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取通灵配置
func (ws *lingTongDevTemplateService) GetLingTongDevTongLingTemplate(classType types.LingTongDevSysType, level int32) gametemplate.LingTongDevTongLingTemplate {
	tongLingMap, ok := ws.tongLingMap[classType]
	if !ok {
		return nil
	}
	to, ok := tongLingMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取培养配置
func (ws *lingTongDevTemplateService) GetLingTongDevPeiYangTemplate(classType types.LingTongDevSysType, level int32) gametemplate.LingTongDevPeiYangTemplate {
	peiYangMap, ok := ws.peiYangMap[classType]
	if !ok {
		return nil
	}
	to, ok := peiYangMap[level]
	if !ok {
		return nil
	}
	return to
}

// func (ws *lingTongDevTemplateService) RandomLingTongDevTemplate() *gametemplate.LingTongDevTemplate {
// 	num := len(ws.lingTongDevList)
// 	index := rand.Intn(num)
// 	return ws.lingTongDevList[index]
// }

//吃幻化丹升级
func (ws *lingTongDevTemplateService) GetLingTongDevEatHuanHuanTemplate(classType types.LingTongDevSysType, curLevel int32, num int32) (lingTongDevHuanHuaTemplate gametemplate.LingTongDevHuanHuaTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	huanHuanMap, ok := ws.huanHuaMap[classType]
	if !ok {
		return nil, false
	}

	for level := curLevel + 1; leftNum > 0; level++ {
		lingTongDevHuanHuaTemplate, flag = huanHuanMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= lingTongDevHuanHuaTemplate.GetItemCount()
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

func (ws *lingTongDevTemplateService) GetLingTongDevTongLingUpgrade(classType types.LingTongDevSysType, curLevel int32, num int32) (lingTongDevTongLingTemplate gametemplate.LingTongDevTongLingTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	tongLingMap, ok := ws.tongLingMap[classType]
	if !ok {
		return nil, false
	}
	for level := curLevel + 1; leftNum > 0; level++ {
		lingTongDevTongLingTemplate, flag = tongLingMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= lingTongDevTongLingTemplate.GetItemCount()
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

//灵童养成培养
func (ws *lingTongDevTemplateService) GetLingTongDevPeiYangUpgrade(classType types.LingTongDevSysType, curLevel int32, num int32) (lingTongDevPeiYangTemplate gametemplate.LingTongDevPeiYangTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	peiYangMap, ok := ws.peiYangMap[classType]
	if !ok {
		return nil, false
	}
	for level := curLevel + 1; leftNum > 0; level++ {
		lingTongDevPeiYangTemplate, flag = peiYangMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= lingTongDevPeiYangTemplate.GetItemCount()
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *lingTongDevTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &lingTongDevTemplateService{}
		err = cs.init()
	})
	return err
}

func GetLingTongDevTemplateService() LingTongDevTemplateService {
	return cs
}
