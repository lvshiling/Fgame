package template

import (
	"fgame/fgame/core/template"
	fabaotypes "fgame/fgame/game/fabao/types"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sync"
)

//法宝模板接口处理
type FaBaoTemplateService interface {
	//获取法宝模板进阶配置
	GetFaBaoNumber(number int32) *gametemplate.FaBaoTemplate
	//获取法宝模板配置
	GetFaBao(id int) *gametemplate.FaBaoTemplate
	//获取食幻化丹配置
	GetFaBaoHuanHuaTemplate(level int32) *gametemplate.FaBaoHuanHuaTemplate
	//获取通灵配置
	GetFaBaoTongLingTemplate(level int32) *gametemplate.FaBaoTongLingTemplate

	RandomFaBaoTemplate() *gametemplate.FaBaoTemplate
	//吃幻化丹升级
	GetFaBaoEatHuanHuanTemplate(curLevel int32, num int32) (*gametemplate.FaBaoHuanHuaTemplate, bool)
	//法宝通灵升级
	GetFaBaoTongLingUpgrade(curLevel int32, num int32) (*gametemplate.FaBaoTongLingTemplate, bool)
}

type faBaoTemplateService struct {
	//法宝模板进阶配置
	faBaoNumberMap map[int32]*gametemplate.FaBaoTemplate
	//法宝模板配置
	faBaoMap map[int]*gametemplate.FaBaoTemplate
	//法宝模板幻化配置
	huanHuaMap map[int32]*gametemplate.FaBaoHuanHuaTemplate
	//法宝模板通灵配置
	tongLingMap map[int32]*gametemplate.FaBaoTongLingTemplate

	faBaoList []*gametemplate.FaBaoTemplate
}

//初始化
func (ws *faBaoTemplateService) init() error {
	ws.faBaoNumberMap = make(map[int32]*gametemplate.FaBaoTemplate)
	ws.faBaoMap = make(map[int]*gametemplate.FaBaoTemplate)
	ws.huanHuaMap = make(map[int32]*gametemplate.FaBaoHuanHuaTemplate)
	ws.tongLingMap = make(map[int32]*gametemplate.FaBaoTongLingTemplate)
	//法宝模板
	templateMap := template.GetTemplateService().GetAll((*gametemplate.FaBaoTemplate)(nil))
	for _, templateObject := range templateMap {
		faBaoTemplate, _ := templateObject.(*gametemplate.FaBaoTemplate)
		ws.faBaoMap[faBaoTemplate.TemplateId()] = faBaoTemplate

		typ := faBaoTemplate.GetTyp()
		if typ == fabaotypes.FaBaoTypeAdvanced {
			ws.faBaoNumberMap[faBaoTemplate.Number] = faBaoTemplate
		}
		ws.faBaoList = append(ws.faBaoList, faBaoTemplate)
	}

	//法宝模板幻化
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.FaBaoHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.FaBaoHuanHuaTemplate)
		ws.huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	//法宝模板通灵
	tongLingTemplateMap := template.GetTemplateService().GetAll((*gametemplate.FaBaoTongLingTemplate)(nil))
	for _, templateObject := range tongLingTemplateMap {
		tongLingTemplate, _ := templateObject.(*gametemplate.FaBaoTongLingTemplate)
		ws.tongLingMap[tongLingTemplate.Level] = tongLingTemplate
	}

	return nil
}

//获取法宝模板进阶配置
func (ws *faBaoTemplateService) GetFaBaoNumber(number int32) *gametemplate.FaBaoTemplate {
	to, ok := ws.faBaoNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//获取幻化配置
func (ws *faBaoTemplateService) GetFaBaoHuanHuaTemplate(level int32) *gametemplate.FaBaoHuanHuaTemplate {
	to, ok := ws.huanHuaMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取通灵配置
func (ws *faBaoTemplateService) GetFaBaoTongLingTemplate(level int32) *gametemplate.FaBaoTongLingTemplate {
	to, ok := ws.tongLingMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取法宝模板配置
func (ws *faBaoTemplateService) GetFaBao(id int) *gametemplate.FaBaoTemplate {
	to, ok := ws.faBaoMap[id]
	if !ok {
		return nil
	}
	return to
}

func (ws *faBaoTemplateService) RandomFaBaoTemplate() *gametemplate.FaBaoTemplate {
	num := len(ws.faBaoList)
	index := rand.Intn(num)
	return ws.faBaoList[index]
}

//吃幻化丹升级
func (ws *faBaoTemplateService) GetFaBaoEatHuanHuanTemplate(curLevel int32, num int32) (faBaoHuanHuaTemplate *gametemplate.FaBaoHuanHuaTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		faBaoHuanHuaTemplate, flag = ws.huanHuaMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= faBaoHuanHuaTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

func (ws *faBaoTemplateService) GetFaBaoTongLingUpgrade(curLevel int32, num int32) (faBaoTongLingTemplate *gametemplate.FaBaoTongLingTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		faBaoTongLingTemplate, flag = ws.tongLingMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= faBaoTongLingTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *faBaoTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &faBaoTemplateService{}
		err = cs.init()
	})
	return err
}

func GetFaBaoTemplateService() FaBaoTemplateService {
	return cs
}
