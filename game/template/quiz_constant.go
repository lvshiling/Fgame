package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
	"strings"
)

//仙尊问答常量模板配置
type QuizConstantTemplate struct {
	*QuizConstantTemplateVO
	//活动开始时间
	beginTimeHour int64
	//活动结束时间
	endTimeHour int64
	//必定出题时间
	refreshTimeArr []int64
}

func (t *QuizConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//必定出题时间
	refreshTimeStr := strings.Split(t.RefreshTime, ",")
	if len(refreshTimeStr) > 0 {
		t.refreshTimeArr = make([]int64, 0, len(refreshTimeStr))
		for i := 0; i < len(refreshTimeStr); i++ {
			val := refreshTimeStr[i]
			refreshTime, err := timeutils.ParseDayOfHHMMSS(val)
			if err != nil {
				return template.NewTemplateFieldError("RefreshTime", fmt.Errorf("[%s] invalid", t.RefreshTime))
			}
			if len(t.refreshTimeArr) > 0 && t.refreshTimeArr[len(t.refreshTimeArr)-1] >= refreshTime {
				return template.NewTemplateFieldError("RefreshTime", fmt.Errorf("[%s] invalid", t.RefreshTime))
			}
			t.refreshTimeArr = append(t.refreshTimeArr, refreshTime)
		}
	}

	//活动开始时间
	beginTime, err := timeutils.ParseDayOfHHMM(t.BeginTime)
	if err != nil {
		return template.NewTemplateFieldError("BeginTime", fmt.Errorf("[%s] invalid", t.BeginTime))
	}
	t.beginTimeHour = beginTime

	//活动结束时间
	endInt, err := strconv.ParseInt(t.EndTime, 10, 64)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.EndTime)
		return template.NewTemplateFieldError("EndTime", err)
	}
	if endInt == 2400 {
		t.endTimeHour = int64(common.DAY)
	} else {
		endTime, err := timeutils.ParseDayOfHHMM(t.EndTime)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.EndTime)
			return template.NewTemplateFieldError("EndTime", err)
		}
		t.endTimeHour = endTime
	}

	return
}

func (t *QuizConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//每题题目的答题时间
	err = validator.MinValidate(float64(t.DaTiTime), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("DaTiTime", err)
	}

	//判断刷新间隔时间大于等于答题时间
	//题目刷新间隔时间
	err = validator.MinValidate(float64(t.IntervalTime), float64(t.DaTiTime), true)
	if err != nil {
		return template.NewTemplateFieldError("IntervalTime", err)
	}

	//假消息时间下限
	err = validator.MinValidate(float64(t.MsgTimeMin), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("MsgTimeMin", err)
	}

	//假消息时间上限
	err = validator.MinValidate(float64(t.MsgTimeMax), float64(t.MsgTimeMin), true)
	if err != nil {
		return template.NewTemplateFieldError("MsgTimeMax", err)
	}

	//验证刷题时间顺序 开始时间 结束时间
	err = validator.MinValidate(float64(t.endTimeHour), float64(t.beginTimeHour), true)
	if err != nil {
		return template.NewTemplateFieldError("BeginTime, EndTime", err)
	}
	if len(t.refreshTimeArr) > 0 {
		err = validator.MinValidate(float64(t.refreshTimeArr[0]), float64(t.beginTimeHour), true)
		if err != nil {
			return template.NewTemplateFieldError("BeginTime, EndTime", err)
		}
		err = validator.MaxValidate(float64(t.refreshTimeArr[len(t.refreshTimeArr)-1]), float64(t.endTimeHour), true)
		if err != nil {
			return template.NewTemplateFieldError("BeginTime, EndTime", err)
		}
		for i := 0; i+1 < len(t.refreshTimeArr); i++ {
			val := t.refreshTimeArr[i]
			nextVal := t.refreshTimeArr[i+1]
			if t.DaTiTime > int32(nextVal-val) {
				return template.NewTemplateFieldError("RefreshTime with DaTiTime", fmt.Errorf("[%s] invalid", t.RefreshTime))
			}
		}
	}

	return
}

func (t *QuizConstantTemplate) PatchAfterCheck() {

}
func (t *QuizConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *QuizConstantTemplate) FileName() string {
	return "tb_quiz_constant.json"
}

func (t *QuizConstantTemplate) GetNearRefreshTime(now int64) (int64, error) {
	nearRefreshTime := int64(0)
	for i := 0; i < len(t.refreshTimeArr); i++ {
		val := t.refreshTimeArr[i]
		beginDay, err := timeutils.BeginOfNow(now)
		if err != nil {
			return nearRefreshTime, err
		}
		tempTime := beginDay + val
		if now > tempTime && tempTime > nearRefreshTime {
			nearRefreshTime = tempTime
		}
	}
	return nearRefreshTime, nil
}

func (t *QuizConstantTemplate) GetBeginTime(now int64) (int64, error) {
	beginDay, err := timeutils.BeginOfNow(now)
	if err != nil {
		return 0, err
	}

	return beginDay + t.beginTimeHour, nil
}

func (t *QuizConstantTemplate) GetEndTime(now int64) (int64, error) {
	beginDay, err := timeutils.BeginOfNow(now)
	if err != nil {
		return 0, err
	}
	return beginDay + t.endTimeHour, nil
}

func init() {
	template.Register((*QuizConstantTemplate)(nil))
}
