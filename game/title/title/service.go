package title

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	titletypes "fgame/fgame/game/title/types"
	"fmt"
	"math/rand"
	"sync"
)

//称号接口处理
type TitleService interface {
	//获取称号配置
	GetTitleTemplate(id int) *gametemplate.TitleTemplate
	//获取称号类型map
	GetTitleTypeMap(titleType titletypes.TitleType) map[titletypes.TitleSubType]int32
	//获取称号id
	GetTitleId(typ titletypes.TitleType, subType titletypes.TitleSubType) (titleId int32, flag bool)
	//随机
	RandomTitleTemplate() *gametemplate.TitleTemplate
	//获取称号升星模板
	GetTitleUpStarTemplate(titleId int, starLev int32) *gametemplate.TitleUpStarTemplate
}

type titleService struct {
	//称号配置
	titleMap map[int]*gametemplate.TitleTemplate
	//称号类型map
	titleTypeMap map[titletypes.TitleType]map[titletypes.TitleSubType]int32

	titleList []*gametemplate.TitleTemplate
}

//初始化
func (ts *titleService) init() error {
	ts.titleMap = make(map[int]*gametemplate.TitleTemplate)
	ts.titleTypeMap = make(map[titletypes.TitleType]map[titletypes.TitleSubType]int32)

	//称号
	templateMap := template.GetTemplateService().GetAll((*gametemplate.TitleTemplate)(nil))
	for _, templateObject := range templateMap {
		titleTemplate, _ := templateObject.(*gametemplate.TitleTemplate)
		ts.titleMap[titleTemplate.TemplateId()] = titleTemplate

		typ := titleTemplate.GetTitleType()
		subType := titleTemplate.GetTitleSubType()
		typeMap, exist := ts.titleTypeMap[typ]
		if !exist {
			typeMap = make(map[titletypes.TitleSubType]int32)
			ts.titleTypeMap[typ] = typeMap
		}
		typeMap[subType] = int32(titleTemplate.TemplateId())

		ts.titleList = append(ts.titleList, titleTemplate)
	}

	//校验 排行榜称号id
	titleRankMap := ts.titleTypeMap[titletypes.TitleTypeRank]
	titleRankSubMap := titletypes.GetTitleRankSubMap()
	for rankSubType, _ := range titleRankSubMap {
		_, exist := titleRankMap[rankSubType]
		if !exist {
			return fmt.Errorf("titleservice: 排行榜 SubType:%d 应该存在", rankSubType)
		}
	}

	//检验大皇帝称号
	kingTitleMap := ts.titleTypeMap[titletypes.TitleTypeKing]
	_, exist := kingTitleMap[titletypes.TitleCommonSubTypeDefault]
	if !exist {
		return fmt.Errorf("titleservice:战力排行第一称号应该存在")
	}
	return nil
}

//获取称号配置
func (ts *titleService) GetTitleTemplate(id int) *gametemplate.TitleTemplate {
	to, ok := ts.titleMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取称号类型map
func (ts *titleService) GetTitleTypeMap(titleType titletypes.TitleType) map[titletypes.TitleSubType]int32 {
	return ts.titleTypeMap[titleType]
}

//获取称号id
func (ts *titleService) GetTitleId(typ titletypes.TitleType, subType titletypes.TitleSubType) (titleId int32, flag bool) {
	typMap, exist := ts.titleTypeMap[typ]
	if !exist {
		return 0, false
	}
	titleId = typMap[subType]
	return titleId, true
}

//获取称号升星模板
func (ts *titleService) GetTitleUpStarTemplate(titleId int, starLev int32) *gametemplate.TitleUpStarTemplate {
	titleTemp, exist := ts.titleMap[titleId]
	if !exist {
		return nil
	}

	upStarTemp := titleTemp.GetTitleUpStarTemplateByStarLev(starLev)
	return upStarTemp
}

func (ts *titleService) RandomTitleTemplate() *gametemplate.TitleTemplate {
	num := len(ts.titleList)
	index := rand.Intn(num)
	return ts.titleList[index]
}

var (
	once sync.Once
	cs   *titleService
)

func Init() (err error) {
	once.Do(func() {
		cs = &titleService{}
		err = cs.init()
	})
	return err
}

func GetTitleService() TitleService {
	return cs
}
