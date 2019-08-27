package template

import (
	"fgame/fgame/core/template"
	feishengtypes "fgame/fgame/game/feisheng/types"
	propertytypes "fgame/fgame/game/property/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//飞升配置服务
type FeiShengTemplaterService interface {
	GetFeiShengTemplate(feiLevel int32) *gametemplate.FeiShengTemplate
	GetQianNengAttrMap(ti, li, gu int32) map[propertytypes.BattlePropertyType]int64
}

type feishengTemplaterService struct {
	feishengMap   map[int32]*gametemplate.FeiShengTemplate
	feishengQnMap map[feishengtypes.QianNengType]*gametemplate.FeiShengQianNengTemplate
}

//初始化
func (ts *feishengTemplaterService) init() error {
	ts.feishengMap = make(map[int32]*gametemplate.FeiShengTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.FeiShengTemplate)(nil))
	for _, temp := range templateMap {
		feishengTemplate, _ := temp.(*gametemplate.FeiShengTemplate)
		ts.feishengMap[feishengTemplate.Level] = feishengTemplate
	}

	//潜能模板
	ts.feishengQnMap = make(map[feishengtypes.QianNengType]*gametemplate.FeiShengQianNengTemplate)
	qnTemplateMap := template.GetTemplateService().GetAll((*gametemplate.FeiShengQianNengTemplate)(nil))
	for _, temp := range qnTemplateMap {
		qnTemplate, _ := temp.(*gametemplate.FeiShengQianNengTemplate)
		ts.feishengQnMap[qnTemplate.GetQianNengType()] = qnTemplate
	}

	return nil
}

func (ts *feishengTemplaterService) GetFeiShengTemplate(level int32) *gametemplate.FeiShengTemplate {
	temp, ok := ts.feishengMap[level]
	if !ok {
		return nil
	}

	return temp
}

func (ts *feishengTemplaterService) GetQianNengAttrMap(ti, li, gu int32) map[propertytypes.BattlePropertyType]int64 {
	newAttrMap := make(map[propertytypes.BattlePropertyType]int64)

	for qnType, qnTemp := range ts.feishengQnMap {
		baseNum := int32(0)
		switch qnType {
		case feishengtypes.QianNengTypeTiZhi:
			baseNum = ti
		case feishengtypes.QianNengTypeLiDao:
			baseNum = li
		case feishengtypes.QianNengTypeJinGu:
			baseNum = gu
		default:
			continue
		}

		for typ, val := range qnTemp.GetBattleAttrMap() {
			_, ok := newAttrMap[typ]
			if !ok {
				newAttrMap[typ] = val * int64(baseNum)
			} else {
				newAttrMap[typ] += val * int64(baseNum)
			}
		}
	}

	return newAttrMap
}

var (
	once sync.Once
	cs   *feishengTemplaterService
)

func Init() (err error) {
	once.Do(func() {
		cs = &feishengTemplaterService{}
		err = cs.init()
	})
	return err
}

func GetFeiShengTemplateService() FeiShengTemplaterService {
	return cs
}
