package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"sort"
	"sync"
)

//套装排序
type ZhenFaTaoZhuangTemplateList []*gametemplate.ZhenFaTaoZhuangTemplate

func (zl ZhenFaTaoZhuangTemplateList) Len() int {
	return len(zl)
}

func (zl ZhenFaTaoZhuangTemplateList) Less(i, j int) bool {
	return zl[i].Level < zl[j].Level
}

func (zl ZhenFaTaoZhuangTemplateList) Swap(i, j int) {
	zl[i], zl[j] = zl[j], zl[i]
}

//阵法接口处理
type ZhenFaTemplateService interface {
	//获取阵法模板
	GetZhenFaTempalte(zhenFaType zhenfatypes.ZhenFaType, level int32) *gametemplate.ZhenFaTemplate
	//获取阵法阵旗
	GetZhenFaZhenQiTemplate(zhenFaType zhenfatypes.ZhenFaType, zhenQiPos zhenfatypes.ZhenQiType, level int32) *gametemplate.ZhenFaZhenQiTemplate
	//获取阵法仙火
	GetZhenFaXianHuoTemplate(zhenFaType zhenfatypes.ZhenFaType, level int32) *gametemplate.ZhenFaXianHuoTemplate
	//获取阵法套装
	GetZhenFaTaoZhuangTemplate(totalLevel int32) *gametemplate.ZhenFaTaoZhuangTemplate
	//阵法激活
	GetZhenFaJiHuoTemplate(zhenFaType zhenfatypes.ZhenFaType) *gametemplate.ZhenFaJiHuoTemplate
}

type zhenFaTemplateService struct {
	//阵法模板
	zhenFaTemplateMap map[zhenfatypes.ZhenFaType]map[int32]*gametemplate.ZhenFaTemplate
	//阵法阵旗模板
	zhenFaZhenQiTemplateMap map[zhenfatypes.ZhenFaType]map[zhenfatypes.ZhenQiType]map[int32]*gametemplate.ZhenFaZhenQiTemplate
	//阵旗仙火
	zhenQiXianHuoTemplateMap map[zhenfatypes.ZhenFaType]map[int32]*gametemplate.ZhenFaXianHuoTemplate
	//阵法套装
	zhenFaTaoZhuangTemplateList []*gametemplate.ZhenFaTaoZhuangTemplate
	//阵法激活
	zhenFaJiHuoTemplateMap map[zhenfatypes.ZhenFaType]*gametemplate.ZhenFaJiHuoTemplate
}

//初始化
func (ts *zhenFaTemplateService) init() (err error) {

	err = ts.initZhenFaJiHuo()
	if err != nil {
		return
	}

	err = ts.initZhenFa()
	if err != nil {
		return
	}

	err = ts.initZhenFaZhenQi()
	if err != nil {
		return
	}

	err = ts.initZhenFaXianHuo()
	if err != nil {
		return
	}

	err = ts.initZhenFaTaoZhuang()
	if err != nil {
		return
	}

	return nil
}

func (ts *zhenFaTemplateService) initZhenFaJiHuo() (err error) {
	ts.zhenFaJiHuoTemplateMap = make(map[zhenfatypes.ZhenFaType]*gametemplate.ZhenFaJiHuoTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ZhenFaJiHuoTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.ZhenFaJiHuoTemplate)
		ts.zhenFaJiHuoTemplateMap[tempTemplate.GetZhenFaType()] = tempTemplate
	}
	return
}

//初始化阵法
func (ts *zhenFaTemplateService) initZhenFa() (err error) {
	ts.zhenFaTemplateMap = make(map[zhenfatypes.ZhenFaType]map[int32]*gametemplate.ZhenFaTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ZhenFaTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.ZhenFaTemplate)

		zhenFaTemplataMap, ok := ts.zhenFaTemplateMap[tempTemplate.GetZhenFaType()]
		if !ok {
			zhenFaTemplataMap = make(map[int32]*gametemplate.ZhenFaTemplate)
			ts.zhenFaTemplateMap[tempTemplate.GetZhenFaType()] = zhenFaTemplataMap
		}
		zhenFaTemplataMap[tempTemplate.Level] = tempTemplate
	}
	return
}

//初始化阵法合成
func (ts *zhenFaTemplateService) initZhenFaZhenQi() (err error) {
	ts.zhenFaZhenQiTemplateMap = make(map[zhenfatypes.ZhenFaType]map[zhenfatypes.ZhenQiType]map[int32]*gametemplate.ZhenFaZhenQiTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ZhenFaZhenQiTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.ZhenFaZhenQiTemplate)

		zhenFaZhenQiMap, ok := ts.zhenFaZhenQiTemplateMap[tempTemplate.GetZhenFaType()]
		if !ok {
			zhenFaZhenQiMap = make(map[zhenfatypes.ZhenQiType]map[int32]*gametemplate.ZhenFaZhenQiTemplate)
			ts.zhenFaZhenQiTemplateMap[tempTemplate.GetZhenFaType()] = zhenFaZhenQiMap
		}

		zhenQiMap, ok := zhenFaZhenQiMap[tempTemplate.GetZhenQiType()]
		if !ok {
			zhenQiMap = make(map[int32]*gametemplate.ZhenFaZhenQiTemplate)
			zhenFaZhenQiMap[tempTemplate.GetZhenQiType()] = zhenQiMap
		}
		zhenQiMap[tempTemplate.Level] = tempTemplate
	}
	return
}

//初始化阵旗仙火
func (ts *zhenFaTemplateService) initZhenFaXianHuo() (err error) {
	ts.zhenQiXianHuoTemplateMap = make(map[zhenfatypes.ZhenFaType]map[int32]*gametemplate.ZhenFaXianHuoTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ZhenFaXianHuoTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.ZhenFaXianHuoTemplate)

		zhenQiXianHuoTemplateMap, ok := ts.zhenQiXianHuoTemplateMap[tempTemplate.GetZhenFaType()]
		if !ok {
			zhenQiXianHuoTemplateMap = make(map[int32]*gametemplate.ZhenFaXianHuoTemplate)
			ts.zhenQiXianHuoTemplateMap[tempTemplate.GetZhenFaType()] = zhenQiXianHuoTemplateMap
		}
		zhenQiXianHuoTemplateMap[tempTemplate.Level] = tempTemplate
	}
	return
}

//初始化阵法套装
func (ts *zhenFaTemplateService) initZhenFaTaoZhuang() (err error) {
	ts.zhenFaTaoZhuangTemplateList = make([]*gametemplate.ZhenFaTaoZhuangTemplate, 0, 10)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ZhenFaTaoZhuangTemplate)(nil))
	for _, templateObject := range templateMap {
		tempTemplate, _ := templateObject.(*gametemplate.ZhenFaTaoZhuangTemplate)
		ts.zhenFaTaoZhuangTemplateList = append(ts.zhenFaTaoZhuangTemplateList, tempTemplate)
	}
	sort.Sort(sort.Reverse(ZhenFaTaoZhuangTemplateList(ts.zhenFaTaoZhuangTemplateList)))
	return
}

func (ts *zhenFaTemplateService) GetZhenFaTempalte(zhenFaType zhenfatypes.ZhenFaType, level int32) *gametemplate.ZhenFaTemplate {
	zhenFaTemplateMap, ok := ts.zhenFaTemplateMap[zhenFaType]
	if !ok {
		return nil
	}
	temTemplate, ok := zhenFaTemplateMap[level]
	if !ok {
		return nil
	}
	return temTemplate
}

func (ts *zhenFaTemplateService) GetZhenFaZhenQiTemplate(zhenFaType zhenfatypes.ZhenFaType, zhenQiPos zhenfatypes.ZhenQiType, level int32) *gametemplate.ZhenFaZhenQiTemplate {
	zhenFaZhenQiTemplateMap, ok := ts.zhenFaZhenQiTemplateMap[zhenFaType]
	if !ok {
		return nil
	}
	zhenFaZhenQiLevelMap, ok := zhenFaZhenQiTemplateMap[zhenQiPos]
	if !ok {
		return nil
	}
	temTemplate, ok := zhenFaZhenQiLevelMap[level]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取阵法仙火
func (ts *zhenFaTemplateService) GetZhenFaXianHuoTemplate(zhenFaType zhenfatypes.ZhenFaType, level int32) *gametemplate.ZhenFaXianHuoTemplate {
	zhenQiXianHuoTemplateMap, ok := ts.zhenQiXianHuoTemplateMap[zhenFaType]
	if !ok {
		return nil
	}
	temTemplate, ok := zhenQiXianHuoTemplateMap[level]
	if !ok {
		return nil
	}
	return temTemplate
}

//获取阵法套装
func (ts *zhenFaTemplateService) GetZhenFaTaoZhuangTemplate(totalLevel int32) *gametemplate.ZhenFaTaoZhuangTemplate {
	for _, taoZhuangTemplate := range ts.zhenFaTaoZhuangTemplateList {
		if totalLevel >= taoZhuangTemplate.Level {
			return taoZhuangTemplate
		}
	}
	return nil
}

func (ts *zhenFaTemplateService) GetZhenFaJiHuoTemplate(zhenFaType zhenfatypes.ZhenFaType) *gametemplate.ZhenFaJiHuoTemplate {
	temTemplate, ok := ts.zhenFaJiHuoTemplateMap[zhenFaType]
	if !ok {
		return nil
	}
	return temTemplate
}

var (
	once sync.Once
	cs   *zhenFaTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &zhenFaTemplateService{}
		err = cs.init()
	})
	return err
}

func GetZhenFaTemplateService() ZhenFaTemplateService {
	return cs
}
