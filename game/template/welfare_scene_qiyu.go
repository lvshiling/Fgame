package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	template.Register((*WelfareSceneQiYuTemplate)(nil))
}

type WelfareSceneQiYuTemplate struct {
	*WelfareSceneQiYuTemplateVO
	boPosInitTemp  *BiologyPosTemplate
	boPosTempList   []*BiologyPosTemplate
	startTime       int64
	endTime         int64
	refreshTimeList []int64
}

func (t *WelfareSceneQiYuTemplate) GetEndTime(now int64) int64 {
	beginDay, _ := timeutils.BeginOfNow(now)
	return t.endTime + beginDay
}

func (t *WelfareSceneQiYuTemplate) GetStartTime(now int64) int64 {
	beginDay, _ := timeutils.BeginOfNow(now)
	return t.startTime + beginDay
}

func (t *WelfareSceneQiYuTemplate) GetRefreshTimeList(now int64) (refreshTimeList []int64) {
	beginDay, _ := timeutils.BeginOfNow(now)
	for i := int32(0); i < t.RefreshBiologyTimes; i++ {
		refreshTime := beginDay + t.startTime + t.RefreshBiologyTime*int64(i)
		refreshTimeList = append(refreshTimeList, refreshTime)
	}
	return
}

func (t *WelfareSceneQiYuTemplate) TemplateId() int {
	return t.Id
}

func (t *WelfareSceneQiYuTemplate) GetBiologyPosList() []*BiologyPosTemplate {
	return t.boPosTempList
}

func (t *WelfareSceneQiYuTemplate) FileName() string {
	return "tb_yunying_qiyudao.json"
}

//组合成需要的数据
func (t *WelfareSceneQiYuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.BiologyBeginId > 0 {
		bo := template.GetTemplateService().Get(int(t.BiologyBeginId), (*BiologyPosTemplate)(nil))
		if bo == nil {
			err = fmt.Errorf("BiologyId [%d] no exist", t.BiologyBeginId)
			return err
		}
		boPosTemp, ok := bo.(*BiologyPosTemplate)
		if !ok {
			err = fmt.Errorf("BiologyId [%d] no exist", t.BiologyBeginId)
			return
		}
		t.boPosInitTemp = boPosTemp
	}
	//

	t.startTime, err = timeutils.ParseDayOfHHMM(t.RefreshBiologyBeginTime)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RefreshBiologyBeginTime)
		return template.NewTemplateFieldError("RefreshBiologyBeginTime", err)
	}
	return nil
}

//检查有效性
func (t *WelfareSceneQiYuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 刷新间隔
	err = validator.RangeValidate(float64(t.RefreshBiologyTime), 1, true, float64(common.DAY), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RefreshBiologyTime)
		return template.NewTemplateFieldError("RefreshBiologyTime", err)
	}
	// 刷新次数
	err = validator.MinValidate(float64(t.RefreshBiologyTimes), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RefreshBiologyTimes)
		return template.NewTemplateFieldError("RefreshBiologyTimes", err)
	}

	//
	endTime := t.startTime + t.RefreshBiologyTime*int64(t.RefreshBiologyTimes)
	if endTime > int64(common.DAY) {
		err = fmt.Errorf("[%d] invalid", t.RefreshBiologyTimes)
		return template.NewTemplateFieldError("RefreshBiologyTimes", err)
	}

	return nil
}

//检验后组合
func (t *WelfareSceneQiYuTemplate) PatchAfterCheck() {
	for initTemp := t.boPosInitTemp; initTemp != nil; initTemp = initTemp.GetNextTemp() {
		t.boPosTempList = append(t.boPosTempList, initTemp)
	}
}
