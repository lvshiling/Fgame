package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/game/wardrobe/wardrobe"
	"fmt"
	"sort"
	"sync"
)

//记录排序
type WardrobeSubTypeList []int32

func (l WardrobeSubTypeList) Len() int {
	return len(l)
}

func (l WardrobeSubTypeList) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l WardrobeSubTypeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

//衣橱接口处理
type WardrobeTemplateService interface {
	//获取衣橱配置
	GetYiChuTemplate(typ int32, subType int32) *gametemplate.YiChuTemplate
	//TODO:zrc 修改到service
	//获取当前衣橱套装下系统激活数
	GetYiChuActiveNum(pl player.Player, typ int32) (permanentNum, totalNum int32)
	//获取技能id
	GetYiChuSkillId(typ int32, subType int32, peiYangLevel int32) (skillId int32, skillId2 int32)
	//获取衣橱套装配置
	GetYiChuSuitTemplate(typ int32) *gametemplate.YiChuSuitTemplate
	//获取衣橱子类型
	GetYiChuSubTypeList(typ int32) []int32
	//衣橱最后一套
	GetYiChuMaxType() int32
}

type wardrobeTemplateService struct {
	//衣橱配置
	wardrobeMap map[int32]map[int32]*gametemplate.YiChuTemplate
	//衣橱套装
	wardrobeSuitMap     map[int32]*gametemplate.YiChuSuitTemplate
	wardrobeSubTypeMap  map[int32][]int32 //衣橱套装子类型
	wardrobeMaxSuitType int32             //套装最后一套
}

//初始化
func (ws *wardrobeTemplateService) init() error {
	ws.wardrobeMap = make(map[int32]map[int32]*gametemplate.YiChuTemplate)
	ws.wardrobeSuitMap = make(map[int32]*gametemplate.YiChuSuitTemplate)
	ws.wardrobeSubTypeMap = make(map[int32][]int32)

	suitTemplateMap := template.GetTemplateService().GetAll((*gametemplate.YiChuSuitTemplate)(nil))
	for _, templateObject := range suitTemplateMap {
		suitTemplate, _ := templateObject.(*gametemplate.YiChuSuitTemplate)
		typ := suitTemplate.GetType()
		ws.wardrobeSuitMap[typ] = suitTemplate

		if ws.wardrobeMaxSuitType < suitTemplate.Type {
			ws.wardrobeMaxSuitType = suitTemplate.Type
		}
	}

	//衣橱
	templateMap := template.GetTemplateService().GetAll((*gametemplate.YiChuTemplate)(nil))
	for _, templateObject := range templateMap {
		yiChuTemplate, _ := templateObject.(*gametemplate.YiChuTemplate)
		typ := yiChuTemplate.GetType()
		subType := yiChuTemplate.GetSubType()
		subTypeMap, ok := ws.wardrobeMap[typ]
		if !ok {
			subTypeMap = make(map[int32]*gametemplate.YiChuTemplate)
			ws.wardrobeMap[typ] = subTypeMap
		}
		subTypeMap[subType] = yiChuTemplate

		suitTemplate, ok := ws.wardrobeSuitMap[typ]
		if !ok {
			return fmt.Errorf("wardrobe:套装类型%d应该是存在的", typ)
		}

		//策划可以改配置
		// if yiChuTemplate.Number > suitTemplate.GetPermanentNum() {
		// 	return fmt.Errorf("wardrobe:套装类型%d的number大于套装类型能永久激活总件数", typ)
		// }

		peiYangTemplate := suitTemplate.GetPeiYangByLevel(yiChuTemplate.ShiDanLimit)
		if peiYangTemplate == nil {
			return fmt.Errorf("wardrobe:套装类型%d的shidanlimit大于套装能升级的数量", typ)
		}

		subTypeList := ws.wardrobeSubTypeMap[typ]
		subTypeList = append(subTypeList, subType)
		ws.wardrobeSubTypeMap[typ] = subTypeList
	}

	for typ, subTypeList := range ws.wardrobeSubTypeMap {
		sort.Sort(WardrobeSubTypeList(subTypeList))
		ws.wardrobeSubTypeMap[typ] = subTypeList
	}

	return nil
}

func (ws *wardrobeTemplateService) GetYiChuSuitTemplate(typ int32) *gametemplate.YiChuSuitTemplate {
	to, ok := ws.wardrobeSuitMap[typ]
	if !ok {
		return nil
	}
	return to
}

func (ws *wardrobeTemplateService) GetYiChuSubTypeList(typ int32) []int32 {
	subTypeList, ok := ws.wardrobeSubTypeMap[typ]
	if !ok {
		return nil
	}
	return subTypeList
}

func (ws *wardrobeTemplateService) GetYiChuMaxType() int32 {
	return ws.wardrobeMaxSuitType
}

//获取衣橱配置
func (ws *wardrobeTemplateService) GetYiChuTemplate(typ int32, subType int32) *gametemplate.YiChuTemplate {
	subTypeMap, ok := ws.wardrobeMap[typ]
	if !ok {
		return nil
	}
	to, ok := subTypeMap[subType]
	if !ok {
		return nil
	}
	return to
}

func (ws *wardrobeTemplateService) GetYiChuSkillId(typ int32, subType int32, peiYangLevel int32) (skillId int32, skillId2 int32) {
	to := ws.GetYiChuTemplate(typ, subType)
	if to == nil {
		return
	}
	if to.SkillId == 0 && to.SkillId2 == 0 {
		return
	}
	if peiYangLevel == 0 {
		return to.SkillId, to.SkillId2
	}

	suitTempalte := ws.GetYiChuSuitTemplate(typ)
	if suitTempalte == nil {
		return
	}
	peiYangTemplate := suitTempalte.GetPeiYangByLevel(peiYangLevel)
	if peiYangTemplate == nil {
		return
	}
	return peiYangTemplate.SkillId, peiYangTemplate.SkillId2
}

func (ws *wardrobeTemplateService) GetYiChuActiveNum(pl player.Player, typ int32) (permanentNum, totalNum int32) {
	suitTemplate := ws.GetYiChuSuitTemplate(typ)
	if suitTemplate == nil {
		return
	}
	for sysType, sysIdInfo := range suitTemplate.GetSysIdMap() {
		if wardrobe.CheckHandle(pl, sysType, sysIdInfo.GetSeqId()) {
			if sysIdInfo.GetIsPermanent() {
				permanentNum++
			}
			totalNum++
		}
	}
	return
}

var (
	once sync.Once
	cs   *wardrobeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &wardrobeTemplateService{}
		err = cs.init()
	})
	return err
}

func GetWardrobeTemplateService() WardrobeTemplateService {
	return cs
}
