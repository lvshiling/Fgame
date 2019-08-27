package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
	"strings"
)

//采集活动配置
type CollectActivityTemplate struct {
	*CollectActivityTemplateVO
	biology         *BiologyTemplate
	mapTemplate     *MapTemplate
	refreshTimeList []int64
}

func (t *CollectActivityTemplate) GetMapTemplate() *MapTemplate {
	return t.mapTemplate
}

func (t *CollectActivityTemplate) GetRefreshTimeList(now int64) (freshTimeList []int64) {
	beginDay, _ := timeutils.BeginOfNow(now)
	for _, tiem := range t.refreshTimeList {
		freshTimeList = append(freshTimeList, tiem+beginDay)
	}
	return
}

func (t *CollectActivityTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//校验活动时间段:yyyymmdd
	timeStr := strings.Split(t.RebornTime, ",")
	for _, value := range timeStr {
		timeInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.RebornTime)
			return template.NewTemplateFieldError("RebornTime", err)
		}
		time := int64(0)
		if timeInt == 2400 {
			time = int64(common.DAY)
		} else {
			time, err = timeutils.ParseDayOfHHMM(value)
			if err != nil {
				return template.NewTemplateFieldError("RebornTime", fmt.Errorf("[%s] invalid", t.RebornTime))
			}
		}
		t.refreshTimeList = append(t.refreshTimeList, time)
	}

	return
}
func (t *CollectActivityTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//group
	err = validator.MinValidate(float64(t.Group), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Group)
		return template.NewTemplateFieldError("Group", err)
	}

	//活动地图id
	mapTempObj := template.GetTemplateService().Get(int(t.RebornMapId), (*MapTemplate)(nil))
	if mapTempObj == nil {
		err = fmt.Errorf("[%d] invalid", t.RebornMapId)
		return template.NewTemplateFieldError("RebornMapId", err)
	}
	t.mapTemplate = mapTempObj.(*MapTemplate)

	// 采集物id
	biologyTempObj := template.GetTemplateService().Get(int(t.BiologyId), (*BiologyTemplate)(nil))
	if biologyTempObj == nil {
		err = fmt.Errorf("[%d] invalid", t.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}
	t.biology = biologyTempObj.(*BiologyTemplate)
	if t.biology.biologyScriptType != scenetypes.BiologyScriptTypeGeneralCollect {
		err = fmt.Errorf("[%d] invalid BiologyScriptType", t.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}

	// 采集物数量
	err = validator.MinValidate(float64(t.RebornCount), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("RebornCount", err)
	}

	return
}

func (t *CollectActivityTemplate) PatchAfterCheck() {

}

func (t *CollectActivityTemplate) TemplateId() int {
	return t.Id
}

func (t *CollectActivityTemplate) FileName() string {
	return "tb_caiji_activity.json"
}

func init() {
	template.Register((*CollectActivityTemplate)(nil))
}
