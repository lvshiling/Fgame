package xianti

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	xiantitypes "fgame/fgame/game/xianti/types"
	"math/rand"
	"sync"
)

//仙体接口处理
type XianTiService interface {
	//获取仙体进阶配置
	GetXianTiNumber(number int32) *gametemplate.XianTiTemplate
	//获取仙体配置
	GetXianTi(id int) *gametemplate.XianTiTemplate
	//获取仙体幻化配置
	GetXianTiHuanHuaTemplate(level int32) *gametemplate.XianTiHuanHuaTemplate
	RandomXianTiTemplate() *gametemplate.XianTiTemplate
	//吃幻化丹升级
	GetXianTiEatHuanHuanTemplate(curLevel int32, num int32) (*gametemplate.XianTiHuanHuaTemplate, bool)
}

type xianTiService struct {
	//进阶map
	xianTiNumberMap map[int32]*gametemplate.XianTiTemplate
	//仙体配置
	xianTiMap map[int]*gametemplate.XianTiTemplate
	//仙体幻化配置
	huanHuaMap map[int32]*gametemplate.XianTiHuanHuaTemplate
	xianTiList []*gametemplate.XianTiTemplate
}

//初始化
func (ms *xianTiService) init() error {
	ms.xianTiNumberMap = make(map[int32]*gametemplate.XianTiTemplate)
	ms.xianTiMap = make(map[int]*gametemplate.XianTiTemplate)
	ms.huanHuaMap = make(map[int32]*gametemplate.XianTiHuanHuaTemplate)
	//仙体
	templateMap := template.GetTemplateService().GetAll((*gametemplate.XianTiTemplate)(nil))
	for _, templateObject := range templateMap {
		xianTiTemplate, _ := templateObject.(*gametemplate.XianTiTemplate)
		ms.xianTiMap[xianTiTemplate.TemplateId()] = xianTiTemplate

		typ := xianTiTemplate.GetTyp()
		if typ == xiantitypes.XianTiTypeAdvanced {
			ms.xianTiNumberMap[xianTiTemplate.Number] = xianTiTemplate
		}
		ms.xianTiList = append(ms.xianTiList, xianTiTemplate)
	}

	//仙体幻化
	huanHuaTemplateMap := template.GetTemplateService().GetAll((*gametemplate.XianTiHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuaTemplateMap {
		xianTiHuanHuaTemplate, _ := templateObject.(*gametemplate.XianTiHuanHuaTemplate)
		ms.huanHuaMap[xianTiHuanHuaTemplate.Level] = xianTiHuanHuaTemplate
	}

	return nil
}

//获取仙体进阶配置
func (ms *xianTiService) GetXianTiNumber(number int32) *gametemplate.XianTiTemplate {
	to, ok := ms.xianTiNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//获取仙体配置
func (ms *xianTiService) GetXianTi(id int) *gametemplate.XianTiTemplate {
	to, ok := ms.xianTiMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取仙体幻化配置
func (ms *xianTiService) GetXianTiHuanHuaTemplate(level int32) *gametemplate.XianTiHuanHuaTemplate {
	to, ok := ms.huanHuaMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取仙体配置
func (ms *xianTiService) RandomXianTiTemplate() *gametemplate.XianTiTemplate {
	num := len(ms.xianTiList)
	index := rand.Intn(num)
	return ms.xianTiList[index]
}

func (ms *xianTiService) GetXianTiEatHuanHuanTemplate(curLevel int32, num int32) (huanHuaTemplate *gametemplate.XianTiHuanHuaTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		huanHuaTemplate, flag = ms.huanHuaMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= huanHuaTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *xianTiService
)

func Init() (err error) {
	once.Do(func() {
		cs = &xianTiService{}
		err = cs.init()
	})
	return err
}

func GetXianTiService() XianTiService {
	return cs
}
