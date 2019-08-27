package soul

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	soultypes "fgame/fgame/game/soul/types"
	gametemplate "fgame/fgame/game/template"
	"sort"
	"sync"
)

//帝魂接口处理
type SoulService interface {
	//获取帝魂配置
	GetSoulTemplate(id int) *gametemplate.SoulTemplate
	//获取帝魂锁链
	GetSoulChainTemplate(num int32) *gametemplate.SoulChainTemplate

	//获取帝魂觉醒模板通过阶别
	GetSoulAwakenTemplateByOrder(typ soultypes.SoulType, order int32) *gametemplate.SoulAwakenTemplate
	//获取帝魂模板通过等级
	GetSoulTemplateByLevel(typ soultypes.SoulType, level int32) *gametemplate.SoulTemplate
	//获取帝魂强化模板通过等级
	GetSoulStrengthenTemplateByLevel(typ soultypes.SoulType, level int32) *gametemplate.SoulLevelUpTemplate
	//获取帝魂吞噬物品通过等级
	GetSoulDevourTemplateByLevel(typ soultypes.SoulType, level int32) map[int32]int32
	//当前吞噬经验对应模板
	GetSoulTemplateByExp(typ soultypes.SoulType, level int32, curExp int32) (to *gametemplate.SoulTemplate, exp int32)
	//帝魂激活模板
	GetSoulActiveTemplate(typ soultypes.SoulType) *gametemplate.SoulTemplate
	//获取帝魂种类
	GetSoulKindTemplate(typ soultypes.SoulType) soultypes.SoulKindType
	//获取帝魂强化的等级限制
	GetSoulStrengthenLevelLimit(level int32) int32
	//获取帝魂最大阶别等级
	//GetAwankenMaxOrder(typ soultypes.SoulType) *gametemplate.SoulAwakenTemplate
}

type soulService struct {
	//帝魂模板
	soulTemplateMap map[int]*gametemplate.SoulTemplate
	//帝魂等级模板
	soulLevelTemplateMap map[soultypes.SoulType]map[int32]*gametemplate.SoulTemplate
	//帝魂觉醒
	soulAwakenTemplateMap map[soultypes.SoulType]map[int32]*gametemplate.SoulAwakenTemplate
	//帝魂锁链配置
	soulChainTemplateList []*gametemplate.SoulChainTemplate
	//帝魂吞噬配置
	soulDevourTemplateMap map[int]*gametemplate.SoulDevourTemplate
	//帝魂强化配置
	soulStrengthenTemplateMap map[soultypes.SoulType]map[int32]*gametemplate.SoulLevelUpTemplate
	//最大技能等级
	//soulAwankenMaxTemplateMap map[soultypes.SoulType]*gametemplate.SoulAwakenTemplate
}

//排序使用
type soulChainTemplateList []*gametemplate.SoulChainTemplate

func (s soulChainTemplateList) Len() int {
	return len(s)
}

func (s soulChainTemplateList) Less(i, j int) bool {
	return s[i].NeedCount < s[j].NeedCount
}

func (s soulChainTemplateList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//初始化
func (ss *soulService) init() error {
	ss.soulTemplateMap = make(map[int]*gametemplate.SoulTemplate)
	ss.soulAwakenTemplateMap = make(map[soultypes.SoulType]map[int32]*gametemplate.SoulAwakenTemplate)
	ss.soulChainTemplateList = make([]*gametemplate.SoulChainTemplate, 0, 8)
	ss.soulLevelTemplateMap = make(map[soultypes.SoulType]map[int32]*gametemplate.SoulTemplate)
	ss.soulDevourTemplateMap = make(map[int]*gametemplate.SoulDevourTemplate)
	ss.soulStrengthenTemplateMap = make(map[soultypes.SoulType]map[int32]*gametemplate.SoulLevelUpTemplate)
	//ss.soulAwankenMaxTemplateMap = make(map[soultypes.SoulType]*gametemplate.SoulAwakenTemplate)

	//赋值soulChainTemplateMap
	scTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SoulChainTemplate)(nil))
	for _, templateObject := range scTemplateMap {
		soulChainTemplate, _ := templateObject.(*gametemplate.SoulChainTemplate)
		ss.soulChainTemplateList = append(ss.soulChainTemplateList, soulChainTemplate)
	}

	//排序
	sort.Sort(soulChainTemplateList(ss.soulChainTemplateList))

	//赋值soulLevelTemplateMap
	sTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SoulTemplate)(nil))
	for _, templateObject := range sTemplateMap {
		soulTemplate, _ := templateObject.(*gametemplate.SoulTemplate)

		//赋值soulTemplateMap
		ss.soulTemplateMap[soulTemplate.TemplateId()] = soulTemplate
		soulLevelMap, ok := ss.soulLevelTemplateMap[soulTemplate.GetSoulType()]
		if !ok {
			soulLevelMap = make(map[int32]*gametemplate.SoulTemplate)
			ss.soulLevelTemplateMap[soulTemplate.GetSoulType()] = soulLevelMap
		}
		_, ok = soulLevelMap[soulTemplate.Level]
		if !ok {
			soulLevelMap[soulTemplate.Level] = soulTemplate
		}
	}

	//赋值soulAwakenTemplateMap
	saTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SoulAwakenTemplate)(nil))
	for _, templateObject := range saTemplateMap {
		soulAwakenTemplate, _ := templateObject.(*gametemplate.SoulAwakenTemplate)

		soulOrderMap, ok := ss.soulAwakenTemplateMap[soulAwakenTemplate.GetSoulType()]
		if !ok {
			soulOrderMap = make(map[int32]*gametemplate.SoulAwakenTemplate)
			ss.soulAwakenTemplateMap[soulAwakenTemplate.GetSoulType()] = soulOrderMap
		}
		_, ok = soulOrderMap[soulAwakenTemplate.Order]
		if !ok {
			soulOrderMap[soulAwakenTemplate.Order] = soulAwakenTemplate
		}
	}

	//赋值soulDevourTemplateMap
	sdTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SoulDevourTemplate)(nil))
	for _, templateObject := range sdTemplateMap {
		soulDevourTemplate, _ := templateObject.(*gametemplate.SoulDevourTemplate)
		ss.soulDevourTemplateMap[soulDevourTemplate.TemplateId()] = soulDevourTemplate
	}

	ssTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SoulLevelUpTemplate)(nil))
	for _, templateObject := range ssTemplateMap {
		soulStrengthenTemplate, _ := templateObject.(*gametemplate.SoulLevelUpTemplate)

		//赋值soulStrengthenTemplateMap
		soulLevelMap, ok := ss.soulStrengthenTemplateMap[soulStrengthenTemplate.GetSoulType()]
		if !ok {
			soulLevelMap = make(map[int32]*gametemplate.SoulLevelUpTemplate)
			ss.soulStrengthenTemplateMap[soulStrengthenTemplate.GetSoulType()] = soulLevelMap
		}
		_, ok = soulLevelMap[soulStrengthenTemplate.Level]
		if !ok {
			soulLevelMap[soulStrengthenTemplate.Level] = soulStrengthenTemplate
		}
	}

	return nil
}

// func (ss *soulService) initSoulMaxAwaken(temp *gametemplate.SoulAwakenTemplate) {
// 	levelTemplate, ok := ss.soulAwankenMaxTemplateMap[temp.GetSoulType()]
// 	if !ok {
// 		ss.soulAwankenMaxTemplateMap[temp.GetSoulType()] = temp
// 		return
// 	}
// 	if temp.Order > levelTemplate.Order {
// 		ss.soulAwankenMaxTemplateMap[temp.GetSoulType()] = temp
// 	}
// 	return
// }

// //获取最大等级
// func (ss *soulService) GetAwankenMaxOrder(typ soultypes.SoulType) *gametemplate.SoulAwakenTemplate {
// 	return ss.soulAwankenMaxTemplateMap[typ]
// }

//获取帝魂配置
func (ss *soulService) GetSoulTemplate(id int) *gametemplate.SoulTemplate {
	to, ok := ss.soulTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

func (ss *soulService) GetSoulChainTemplate(num int32) *gametemplate.SoulChainTemplate {
	var soulChainTemplate *gametemplate.SoulChainTemplate
	for _, tempSoulChainTemplate := range ss.soulChainTemplateList {
		if num < tempSoulChainTemplate.NeedCount {
			return soulChainTemplate
		}
		soulChainTemplate = tempSoulChainTemplate
	}
	return soulChainTemplate
}

//获取帝魂id通过level
func (ss *soulService) GetSoulTemplateByLevel(typ soultypes.SoulType, level int32) *gametemplate.SoulTemplate {
	levelMap, ok := ss.soulLevelTemplateMap[typ]
	if !ok {
		return nil
	}
	levelTemplate, ok := levelMap[level]
	if !ok {
		return nil
	}
	return levelTemplate
}

//获取帝魂强化模板通过等级
func (ss *soulService) GetSoulStrengthenTemplateByLevel(typ soultypes.SoulType, level int32) *gametemplate.SoulLevelUpTemplate {
	levelMap, ok := ss.soulStrengthenTemplateMap[typ]
	if !ok {
		return nil
	}
	levelTemplate, ok := levelMap[level]
	if !ok {
		return nil
	}
	return levelTemplate
	return nil
}

//获取帝魂觉醒模板通过阶别
func (ss *soulService) GetSoulAwakenTemplateByOrder(typ soultypes.SoulType, order int32) *gametemplate.SoulAwakenTemplate {
	orderMap, ok := ss.soulAwakenTemplateMap[typ]
	if !ok {
		return nil
	}
	orderTemplate, ok := orderMap[order]
	if !ok {
		return nil
	}
	return orderTemplate
}

//获取帝魂吞噬物品通过等级
func (ss *soulService) GetSoulDevourTemplateByLevel(typ soultypes.SoulType, level int32) map[int32]int32 {
	levelTemplate := ss.GetSoulTemplateByLevel(typ, level)
	if levelTemplate == nil {
		return nil
	}
	to, ok := ss.soulDevourTemplateMap[int(levelTemplate.DevourId)]
	if !ok {
		return nil
	}
	return to.GetDevourExpMap()
}

//当前吞噬经验对应模板
func (ss *soulService) GetSoulTemplateByExp(typ soultypes.SoulType, level int32, curExp int32) (to *gametemplate.SoulTemplate, exp int32) {
	levelTemplate := ss.GetSoulTemplateByLevel(typ, level)
	if levelTemplate == nil {
		return nil, 0
	}
	if levelTemplate.NextId == 0 {
		return levelTemplate, 0
	}
	nextTemplate := ss.soulTemplateMap[int(levelTemplate.NextId)]
	lastTemplate := levelTemplate
	uplevelExp := int32(0)
	passLevelExp := int32(0)
	for ; nextTemplate != nil; nextTemplate = ss.soulTemplateMap[int(nextTemplate.NextId)] {
		uplevelExp += nextTemplate.UplevelExp
		if curExp < uplevelExp {
			return lastTemplate, curExp - passLevelExp
		} else if curExp > uplevelExp {
			passLevelExp = uplevelExp
			lastTemplate = nextTemplate
		} else {
			return nextTemplate, 0
		}
	}
	return lastTemplate, 0
}

//帝魂激活模板
func (ss *soulService) GetSoulActiveTemplate(typ soultypes.SoulType) *gametemplate.SoulTemplate {
	levelMap, ok := ss.soulLevelTemplateMap[typ]
	if !ok {
		return nil
	}
	levelTemplate, ok := levelMap[1]
	if !ok {
		return nil
	}
	return levelTemplate
}

//获取帝魂种类
func (ss *soulService) GetSoulKindTemplate(typ soultypes.SoulType) soultypes.SoulKindType {
	levelMap, ok := ss.soulLevelTemplateMap[typ]
	if !ok {
		return 0
	}
	levelTemplate, ok := levelMap[1]
	if !ok {
		return 0
	}
	return levelTemplate.GetKindType()
}

//获取帝魂强化的等级限制
func (ss *soulService) GetSoulStrengthenLevelLimit(level int32) int32 {
	limit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSoulStrengthenLevelLimit)
	return int32(level / limit)
}

var (
	once sync.Once
	cs   *soulService
)

func Init() (err error) {
	once.Do(func() {
		cs = &soulService{}
		err = cs.init()
	})
	return err
}

func GetSoulService() SoulService {
	return cs
}
