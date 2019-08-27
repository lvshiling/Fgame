package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sort"
	"sync"
)

//排序
type LongGongRankTemplateList []*gametemplate.LongGongRankTemplate

func (s LongGongRankTemplateList) Len() int {
	return len(s)
}

func (s LongGongRankTemplateList) Less(i, j int) bool {
	return s[i].RankMin < s[j].RankMin
}

func (s LongGongRankTemplateList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type LongGongTemplateService interface {
	//获取龙宫探宝常量模板
	GetLongGongConstTemplate() *gametemplate.LongGongConstantTemplate
	//获取龙宫排行榜根据名次
	GetLongGongRankTemplateByRankNum(rankNum int32) *gametemplate.LongGongRankTemplate
}

type longGongTemplateService struct {
	//龙宫探宝常量模板
	longGongConstTemplate *gametemplate.LongGongConstantTemplate
	//龙宫排行
	longGongRankList []*gametemplate.LongGongRankTemplate
}

//初始化
func (xtts *longGongTemplateService) init() error {
	xtts.longGongRankList = make([]*gametemplate.LongGongRankTemplate, 0, 8)

	//常量配置
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LongGongConstantTemplate)(nil))
	for _, templateObject := range templateMap {
		xtts.longGongConstTemplate, _ = templateObject.(*gametemplate.LongGongConstantTemplate)
		break
	}

	//龙宫排行
	templateRankMap := template.GetTemplateService().GetAll((*gametemplate.LongGongRankTemplate)(nil))
	for _, templateObject := range templateRankMap {
		rankTemplate, _ := templateObject.(*gametemplate.LongGongRankTemplate)
		xtts.longGongRankList = append(xtts.longGongRankList, rankTemplate)
	}
	sort.Sort(sort.Reverse(LongGongRankTemplateList(xtts.longGongRankList)))

	return nil
}

//获取龙宫探宝常量模板
func (xtts *longGongTemplateService) GetLongGongConstTemplate() *gametemplate.LongGongConstantTemplate {
	return xtts.longGongConstTemplate
}

//获取神魔排行榜根据名次
func (xtts *longGongTemplateService) GetLongGongRankTemplateByRankNum(rankNum int32) *gametemplate.LongGongRankTemplate {
	for _, longGongRank := range xtts.longGongRankList {
		if rankNum >= longGongRank.RankMin && rankNum <= longGongRank.RankMax {
			return longGongRank
		}
	}
	return nil
}

var (
	once sync.Once
	cs   *longGongTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &longGongTemplateService{}
		err = cs.init()
	})
	return err
}

func GetLongGongTemplateService() LongGongTemplateService {
	return cs
}
