package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
	"strings"
)

//活动模板配置
type ActivityTimeTemplate struct {
	*ActivityTimeTemplateVO
	//活动时间段开始
	quantumStart int64
	//活动时间段结束
	quantumEnd int64
	//每周活动天
	weekdayArr []int32
	//每月活动天
	monthdayArr []int32
	//活动开始时间
	beginTimeHour int64
	//活动结束时间
	endTimeHour int64
	//开服时间条件
	afterOpenLimitTime int64
	//开服后结束条件
	afterOpenEndTime int64
	//合服时间条件
	afterMergeLimitTime int64
	//合服后结束条件
	afterMergeEndTime int64
}

func (t *ActivityTimeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//校验活动时间段:yyyymmdd
	quantumStr := strings.Split(t.TimeQuantum, ",")
	if len(quantumStr)%2 == 0 {
		for index, value := range quantumStr {
			quantum, err := timeutils.ParseYYYYMMDD(value)
			if err != nil {
				return template.NewTemplateFieldError("TimeQuantum", fmt.Errorf("[%s] invalid", t.TimeQuantum))
			}

			if index%2 == 0 || index/2 == 0 {
				t.quantumStart = quantum
			}

			if index%2 != 0 {
				t.quantumEnd = quantum
			}
		}
	}

	//每周几活动时间
	weekday, err := utils.SplitAsIntArray(t.Weekday)
	if err != nil {
		return template.NewTemplateFieldError("Weekday", fmt.Errorf("[%s] invalid", t.Weekday))
	}
	t.weekdayArr = append(t.weekdayArr, weekday...)

	monthDay, err := utils.SplitAsIntArray(t.Monthday)
	if err != nil {
		return template.NewTemplateFieldError("MonthDay", fmt.Errorf("[%s] invalid", t.Monthday))
	}
	t.monthdayArr = append(t.monthdayArr, monthDay...)

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

func (t *ActivityTimeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//活动Id
	activityTemp := template.GetTemplateService().Get(int(t.ActivityId), (*ActivityTemplate)(nil))
	if activityTemp == nil {
		return template.NewTemplateFieldError("ActivityId", fmt.Errorf("[%s] invalid", t.ActivityId))
	}
	//开服几天
	if t.AfterOpensvrDays > 0 {
		err = validator.MinValidate(float64(t.AfterOpensvrDays), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("AfterOpensvrDays", fmt.Errorf("[%s] invalid", t.AfterOpensvrDays))
		}
		t.afterOpenLimitTime = int64(t.AfterOpensvrDays-1) * int64(common.DAY)
	}

	//开服几天结束活动
	if t.AfterKaifuDayOver > 0 {
		err = validator.MinValidate(float64(t.AfterKaifuDayOver), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("AfterKaifuDayOver", fmt.Errorf("[%s] invalid", t.AfterKaifuDayOver))
		}
		t.afterOpenEndTime = int64(t.AfterKaifuDayOver-1) * int64(common.DAY)
	}

	//开服几天
	if t.AfterHefuDays > 0 {
		err = validator.MinValidate(float64(t.AfterHefuDays), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("AfterHefuDays", fmt.Errorf("[%s] invalid", t.AfterHefuDays))
		}
		t.afterMergeLimitTime = int64(t.AfterHefuDays-1) * int64(common.DAY)
	}

	//合服几天结束活动
	if t.AfterHefuDayOver > 0 {
		err = validator.MinValidate(float64(t.AfterHefuDayOver), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("AfterHefuDayOver", fmt.Errorf("[%s] invalid", t.AfterHefuDayOver))
		}
		t.afterMergeEndTime = int64(t.AfterHefuDayOver-1) * int64(common.DAY)
	}

	//倍数
	if t.BeiShu > 0 {
		err = validator.MinValidate(float64(t.BeiShu), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("BeiShu", fmt.Errorf("[%s] invalid", t.BeiShu))
		}
	}

	return
}

func (t *ActivityTimeTemplate) PatchAfterCheck() {

}
func (t *ActivityTimeTemplate) TemplateId() int {
	return t.Id
}

func (t *ActivityTimeTemplate) FileName() string {
	return "tb_activity_time.json"
}

func (t *ActivityTimeTemplate) GetWeekDayArr() []int32 {
	return t.weekdayArr
}

func (t *ActivityTimeTemplate) GetMonthDayArr() []int32 {
	return t.monthdayArr
}
func (t *ActivityTimeTemplate) GetQuanTumStart() int64 {
	return t.quantumStart
}

func (t *ActivityTimeTemplate) GetQuantumEnd() int64 {
	return t.quantumEnd
}

func (t *ActivityTimeTemplate) GetBeginTime(now int64) (int64, error) {
	beginDay, err := timeutils.BeginOfNow(now)
	if err != nil {
		return 0, err
	}

	return beginDay + t.beginTimeHour, nil
}

func (t *ActivityTimeTemplate) GetEndTime(now int64) (int64, error) {
	beginDay, err := timeutils.BeginOfNow(now)
	if err != nil {
		return 0, err
	}
	return beginDay + t.endTimeHour, nil
}

func (t *ActivityTimeTemplate) IsOnActivityTime(now int64) bool {
	startTime, _ := t.GetBeginTime(now)
	endTime, _ := t.GetEndTime(now)
	if now >= startTime && now <= endTime {
		return true
	}
	return false
}

func (t *ActivityTimeTemplate) GetAfterOpenLimitTime(openTime int64) (int64, error) {
	openTime, err := timeutils.BeginOfNow(openTime)
	if err != nil {
		return 0, err
	}
	return openTime + t.afterOpenLimitTime, nil
}

func (t *ActivityTimeTemplate) GetAfterOpenEndTime(openTime int64) (bool, int64, error) {
	if t.afterOpenEndTime == 0 {
		return false, 0, nil
	}
	openTime, err := timeutils.BeginOfNow(openTime)
	if err != nil {
		return false, 0, err
	}
	return true, openTime + t.afterOpenEndTime, nil
}

func (t *ActivityTimeTemplate) GetAfterMergeLimitTime(mergeTime int64) (int64, error) {
	mergeTime, err := timeutils.BeginOfNow(mergeTime)
	if err != nil {
		return 0, err
	}
	return mergeTime + t.afterMergeLimitTime, nil
}

func (t *ActivityTimeTemplate) GetAfterMergeEndTime(mergeTime int64) (bool, int64, error) {
	if t.afterMergeEndTime == 0 {
		return false, 0, nil
	}
	mergeTime, err := timeutils.BeginOfNow(mergeTime)
	if err != nil {
		return false, 0, err
	}
	return true, mergeTime + t.afterMergeEndTime, nil
}

func (t *ActivityTimeTemplate) IsMustStartActivity(now, openTime, mergeTime int64) bool {
	if t.KaiQiOpensvrDay > 0 {
		diff, _ := timeutils.DiffDay(now, openTime)
		if diff+1 == t.KaiQiOpensvrDay {
			return true
		}
	}

	if t.KaiQiHeFuDay > 0 {
		diff, _ := timeutils.DiffDay(now, mergeTime)
		if diff+1 == t.KaiQiHeFuDay {
			return true
		}
	}

	return false
}

func init() {
	template.Register((*ActivityTimeTemplate)(nil))
}
