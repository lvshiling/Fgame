package wing

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	gametemplate "fgame/fgame/game/template"
	wingtypes "fgame/fgame/game/wing/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math/rand"
	"sync"
)

//战翼接口处理
type WingService interface {
	//获取战翼进阶配置
	GetWingNumber(number int32) *gametemplate.WingTemplate
	//获取战翼配置
	GetWing(id int) *gametemplate.WingTemplate
	//获取护体仙羽配置
	GetFeather(id int32) *gametemplate.FeatherTemplate
	//获取食幻化丹配置
	GetWingHuanHuaTemplate(level int32) *gametemplate.WingHuanHuaTemplate

	//获取试用卡试用阶数
	GetWingTrialOrderId() int32

	RandomWingTemplate() *gametemplate.WingTemplate
	//吃幻化丹升级
	GetWingEatHuanHuanTemplate(curLevel int32, num int32) (*gametemplate.WingHuanHuaTemplate, bool)
}

type wingService struct {
	//战翼进阶配置
	wingNumberMap map[int32]*gametemplate.WingTemplate
	//战翼配置
	wingMap map[int]*gametemplate.WingTemplate
	//护体仙羽配置
	featherMap map[int32]*gametemplate.FeatherTemplate
	//战翼幻化配置
	huanHuaMap map[int32]*gametemplate.WingHuanHuaTemplate
	//战翼试用id
	wingTrialList []int32

	wingList []*gametemplate.WingTemplate
}

//初始化
func (ws *wingService) init() error {
	ws.wingNumberMap = make(map[int32]*gametemplate.WingTemplate)
	ws.wingMap = make(map[int]*gametemplate.WingTemplate)
	ws.featherMap = make(map[int32]*gametemplate.FeatherTemplate)
	ws.huanHuaMap = make(map[int32]*gametemplate.WingHuanHuaTemplate)
	//战翼
	templateMap := template.GetTemplateService().GetAll((*gametemplate.WingTemplate)(nil))
	for _, templateObject := range templateMap {
		wingTemplate, _ := templateObject.(*gametemplate.WingTemplate)
		ws.wingMap[wingTemplate.TemplateId()] = wingTemplate

		typ := wingTemplate.GetTyp()
		if typ == wingtypes.WingTypeAdvanced {
			ws.wingNumberMap[wingTemplate.Number] = wingTemplate
		}
		ws.wingList = append(ws.wingList, wingTemplate)
	}
	//护体仙羽
	templateMap = template.GetTemplateService().GetAll((*gametemplate.FeatherTemplate)(nil))
	for _, templateObject := range templateMap {
		featherTemplate, _ := templateObject.(*gametemplate.FeatherTemplate)
		ws.featherMap[int32(featherTemplate.TemplateId())] = featherTemplate

	}

	//战翼幻化
	huanHuatemplateMap := template.GetTemplateService().GetAll((*gametemplate.WingHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuatemplateMap {
		huanHuaTemplate, _ := templateObject.(*gametemplate.WingHuanHuaTemplate)
		ws.huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}
	err := ws.checkData()
	if err != nil {
		return err
	}
	return nil
}

//校验数据
func (ws *wingService) checkData() (err error) {
	itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWing, itemtypes.ItemWingSubTypeTrialCard)
	if itemTemplate == nil {
		return fmt.Errorf("wingservice:战翼试用卡配置应该是存在的")
	}

	//战翼试用卡
	flag1 := itemTemplate.TypeFlag1
	flag2 := itemTemplate.TypeFlag2
	flag3 := itemTemplate.TypeFlag3
	flag4 := itemTemplate.TypeFlag4
	ws.wingTrialList = append(ws.wingTrialList, flag1, flag2, flag3, flag4)

	for _, number := range ws.wingTrialList {
		wingNumberTemplate := ws.GetWingNumber(number)
		if wingNumberTemplate == nil || wingNumberTemplate.TemplateId() == 0 {
			return fmt.Errorf("wingservice:战翼试用卡试用阶数的配值应该是有效的")
		}
	}
	return nil
}

//获取试用卡试用阶数
func (ws *wingService) GetWingTrialOrderId() int32 {
	indexLen := len(ws.wingTrialList)
	index := mathutils.RandomRange(0, int(indexLen))
	return ws.wingTrialList[index]
}

//获取战翼进阶配置
func (ws *wingService) GetWingNumber(number int32) *gametemplate.WingTemplate {
	to, ok := ws.wingNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//获取幻化配置
func (ws *wingService) GetWingHuanHuaTemplate(level int32) *gametemplate.WingHuanHuaTemplate {
	to, ok := ws.huanHuaMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取战翼配置
func (ws *wingService) GetWing(id int) *gametemplate.WingTemplate {
	to, ok := ws.wingMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取护体仙羽配置
func (ws *wingService) GetFeather(id int32) *gametemplate.FeatherTemplate {
	to, ok := ws.featherMap[id]
	if !ok {
		return nil
	}
	return to
}

func (ws *wingService) RandomWingTemplate() *gametemplate.WingTemplate {
	num := len(ws.wingList)
	index := rand.Intn(num)
	return ws.wingList[index]
}

//吃幻化丹升级
func (ws *wingService) GetWingEatHuanHuanTemplate(curLevel int32, num int32) (wingHuanHuaTemplate *gametemplate.WingHuanHuaTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		wingHuanHuaTemplate, flag = ws.huanHuaMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= wingHuanHuaTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *wingService
)

func Init() (err error) {
	once.Do(func() {
		cs = &wingService{}
		err = cs.init()
	})
	return err
}

func GetWingService() WingService {
	return cs
}
