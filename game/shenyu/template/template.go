package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"

	"sync"
)

//
const (
	shenYuInitRound = 1 //神域之战初始轮
)

//神域配置服务
type ShenYuTemplaterService interface {
	// 神域配置
	GetShenYuTemplate(round int32) *gametemplate.ShenYuTemplate
	GetShenYuInitRoundTemplate() *gametemplate.ShenYuTemplate
	// 神域常量配置
	GetShenYuConstantTemplate() *gametemplate.ShenYuConstantTemplate
	// 神域排行榜配置
	GetShenYuRankTemplate(round, ranking int32) *gametemplate.ShenYuRankTemplate
}

type shenyuTemplaterService struct {
	shenYuMap          map[int32]*gametemplate.ShenYuTemplate
	shenYuConstantTemp *gametemplate.ShenYuConstantTemplate
	shenYuRankMap      map[int32][]*gametemplate.ShenYuRankTemplate
}

//初始化
func (ts *shenyuTemplaterService) init() error {
	ts.shenYuMap = make(map[int32]*gametemplate.ShenYuTemplate)
	var mapIdList []int32

	templateMap := template.GetTemplateService().GetAll((*gametemplate.ShenYuTemplate)(nil))
	for _, temp := range templateMap {
		shenyuTemplate := temp.(*gametemplate.ShenYuTemplate)
		ts.shenYuMap[shenyuTemplate.RoundType] = shenyuTemplate

		mapIdList = append(mapIdList, shenyuTemplate.MapId)
	}

	if len(mapIdList) == 0 || utils.IfRepeatElementInt32(mapIdList) {
		return fmt.Errorf("神域地图配置错误")
	}

	// 神域常量配置
	constantTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ShenYuConstantTemplate)(nil))
	if len(constantTemplateMap) != 1 {
		return fmt.Errorf("shenyu：神域常量配置只有一条")
	}
	for _, temp := range constantTemplateMap {
		constantTemp := temp.(*gametemplate.ShenYuConstantTemplate)
		ts.shenYuConstantTemp = constantTemp
	}

	//排行榜
	ts.shenYuRankMap = make(map[int32][]*gametemplate.ShenYuRankTemplate)
	rankTempMap := template.GetTemplateService().GetAll((*gametemplate.ShenYuRankTemplate)(nil))
	for _, to := range rankTempMap {
		rankTemp, _ := to.(*gametemplate.ShenYuRankTemplate)
		ts.shenYuRankMap[rankTemp.RoundType] = append(ts.shenYuRankMap[rankTemp.RoundType], rankTemp)
	}

	for roundType, rankList := range ts.shenYuRankMap {
		sort.Sort(gametemplate.ShenYuRankTemplateList(rankList))
		ts.shenYuRankMap[roundType] = rankList
	}

	return nil
}

func (ts *shenyuTemplaterService) GetShenYuTemplate(round int32) *gametemplate.ShenYuTemplate {
	return ts.shenYuMap[round]
}

func (ts *shenyuTemplaterService) GetShenYuInitRoundTemplate() *gametemplate.ShenYuTemplate {
	return ts.shenYuMap[shenYuInitRound]
}

func (ts *shenyuTemplaterService) GetShenYuConstantTemplate() *gametemplate.ShenYuConstantTemplate {
	return ts.shenYuConstantTemp
}

func (ts *shenyuTemplaterService) GetShenYuRankTemplate(round, ranking int32) *gametemplate.ShenYuRankTemplate {
	rankList, ok := ts.shenYuRankMap[round]
	if !ok {
		return nil
	}

	for _, temp := range rankList {
		if ranking >= temp.RankMin && ranking <= temp.RankMax {
			return temp
		}
	}

	return nil
}

var (
	once sync.Once
	cs   *shenyuTemplaterService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shenyuTemplaterService{}
		err = cs.init()
	})
	return err
}

func GetShenYuTemplateService() ShenYuTemplaterService {
	return cs
}
