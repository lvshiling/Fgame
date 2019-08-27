package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	ylptypes "fgame/fgame/game/yinglingpu/types"
	"sync"
)

type YingLingPuTemplateService interface {
	GetYingLingPuById(tujianId int32, tujianType ylptypes.YingLingPuTuJianType) *gametemplate.YinglingpuTemplate
	// GetYingLingPuLevel(tujianId int32, level int32, tujianType ylptypes.YingLingPuTuJianType) *gametemplate.YinglingpuLevelTemplate
	GetYingLingPuSuiPian(tujianId int32, suiPianId int32, tujianType ylptypes.YingLingPuTuJianType) *gametemplate.YinglingPuSuiPianTemplate
	GetYingLingPuSuitMap() map[int32]*gametemplate.YingLingPuSuitTemplate
	GetYingLingPuSuitById(id int32) *gametemplate.YingLingPuSuitTemplate
}

type yingLingPuTemplateService struct {
	yingLingPuMap map[ylptypes.YingLingPuTuJianType]map[int32]*gametemplate.YinglingpuTemplate
	ylpSuitMap    map[int32]*gametemplate.YingLingPuSuitTemplate
}

func (t *yingLingPuTemplateService) init() error {
	t.yingLingPuMap = make(map[ylptypes.YingLingPuTuJianType]map[int32]*gametemplate.YinglingpuTemplate)

	for i := ylptypes.YingLingPuTuJianTypeCommon; i <= ylptypes.GetMaxYingLingPuType(); i++ {
		t.yingLingPuMap[i] = make(map[int32]*gametemplate.YinglingpuTemplate)
	}
	allMap := template.GetTemplateService().GetAll((*gametemplate.YinglingpuTemplate)(nil))
	for _, value := range allMap {
		tempValue := value.(*gametemplate.YinglingpuTemplate)
		ylpType := ylptypes.YingLingPuTuJianType(tempValue.Type)
		if !ylpType.Valid() {
			continue
		}
		t.yingLingPuMap[ylpType][tempValue.TujianId] = tempValue
	}

	// 英灵普套装
	t.ylpSuitMap = make(map[int32]*gametemplate.YingLingPuSuitTemplate)
	suitTempMap := template.GetTemplateService().GetAll((*gametemplate.YingLingPuSuitTemplate)(nil))
	for _, to := range suitTempMap {
		suitTemp := to.(*gametemplate.YingLingPuSuitTemplate)
		t.ylpSuitMap[int32(suitTemp.Id)] = suitTemp
	}

	return nil
}

func (t *yingLingPuTemplateService) GetYingLingPuById(tujianId int32, tujianType ylptypes.YingLingPuTuJianType) *gametemplate.YinglingpuTemplate {
	return t.yingLingPuMap[tujianType][tujianId]
}

// func (t *yingLingPuTemplateService) GetYingLingPuLevel(tujianId int32, level int32, tujianType ylptypes.YingLingPuTuJianType) *gametemplate.YinglingpuLevelTemplate {
// 	item, ok := t.yingLingPuMap[tujianType][tujianId]
// 	if !ok {
// 		return nil
// 	}
// 	levelMap := item.GetLevelMap()
// 	if len(levelMap) == 0 {
// 		return nil
// 	}
// 	levelInfo, ok := levelMap[level]
// 	if !ok {
// 		return nil
// 	}
// 	return levelInfo
// }

func (t *yingLingPuTemplateService) GetYingLingPuSuiPian(tujianId int32, suiPianId int32, tujianType ylptypes.YingLingPuTuJianType) *gametemplate.YinglingPuSuiPianTemplate {
	item, ok := t.yingLingPuMap[tujianType][tujianId]
	if !ok {
		return nil
	}
	suiPianMap := item.GetSuiPianMap()
	if len(suiPianMap) == 0 {
		return nil
	}
	suiPianInfo, ok := suiPianMap[suiPianId]
	if !ok {
		return nil
	}
	return suiPianInfo
}

func (t *yingLingPuTemplateService) GetYingLingPuSuitMap() map[int32]*gametemplate.YingLingPuSuitTemplate {
	return t.ylpSuitMap
}

func (t *yingLingPuTemplateService) GetYingLingPuSuitById(id int32) *gametemplate.YingLingPuSuitTemplate {
	return t.ylpSuitMap[id]
}

var (
	once sync.Once
	cs   *yingLingPuTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &yingLingPuTemplateService{}
		err = cs.init()
	})
	return err
}

func GetYingLingPuTemplateService() YingLingPuTemplateService {
	return cs
}
