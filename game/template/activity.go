package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"time"
)

//活动模板配置
type ActivityTemplate struct {
	*ActivityTemplateVO
	activityType activitytypes.ActivityType
	funcOpenType funcopentypes.FuncOpenType
	timeList     []*ActivityTimeTemplate
	mapTemplate  *MapTemplate
}

func (t *ActivityTemplate) GetActivityType() activitytypes.ActivityType {
	return t.activityType
}

func (t *ActivityTemplate) GetTimeList() []*ActivityTimeTemplate {
	return t.timeList
}

func (t *ActivityTemplate) GetMapTemplate() *MapTemplate {
	return t.mapTemplate
}

func (t *ActivityTemplate) IsAtActivityTime(now, openTime, mergeTime int64) (bool, error) {
	tmp, err := t.GetActivityTimeTemplate(now, openTime, mergeTime)
	if err != nil {
		return false, err
	}
	if tmp == nil {
		return false, nil
	}

	return true, nil
}

func (t *ActivityTemplate) GetActivityTimeTemplate(now, openTime, mergeTime int64) (*ActivityTimeTemplate, error) {

	for _, timeTmp := range t.timeList {
		if !t.isValidTimeTemplate(timeTmp, now, openTime, mergeTime) {
			continue
		}

		//是否活动时间
		begin, err := timeTmp.GetBeginTime(now)
		if err != nil {
			return nil, err
		}

		end, err := timeTmp.GetEndTime(now)
		if err != nil {
			return nil, err
		}
		isAtTime := now >= begin && now <= end
		if isAtTime {
			return timeTmp, nil
		}
	}
	return nil, nil
}

func (t *ActivityTemplate) GetOnDateTimeTemplate(now, openTime, mergeTime int64) *ActivityTimeTemplate {
	for _, timeTemp := range t.timeList {
		if !t.isValidTimeTemplate(timeTemp, now, openTime, mergeTime) {
			continue
		}
		return timeTemp
	}
	return nil
}

func (t *ActivityTemplate) isValidTimeTemplate(timeTmp *ActivityTimeTemplate, now, openTime, mergeTime int64) bool {
	// 是否一定开活动
	if openTime != 0 && timeTmp.IsMustStartActivity(now, openTime, mergeTime) {
		return true
	}
	hasTime := true
	if openTime != 0 {
		//开服后结束活动
		hasEnd, afterOpenEnd, err := timeTmp.GetAfterOpenEndTime(openTime)
		if err != nil {
			return false
		}

		if hasEnd && now > afterOpenEnd {
			goto Merge
		}

		//开服后时间限制
		afterOpenLimit, err := timeTmp.GetAfterOpenLimitTime(openTime)
		if err != nil {
			return false
		}
		if afterOpenLimit > now {
			hasTime = false
			goto Merge
		}
		goto HasTime
	}
Merge:
	if mergeTime != 0 {
		//合服后结束活动
		hasEnd, afterMergeEnd, err := timeTmp.GetAfterMergeEndTime(mergeTime)
		if err != nil {
			return false
		}
		if hasEnd && now > afterMergeEnd {
			return false
		}

		//开服后时间限制
		afterMergeLimit, err := timeTmp.GetAfterMergeLimitTime(mergeTime)
		if err != nil {
			return false
		}
		if afterMergeLimit > now {
			return false
		}
	}
HasTime:
	if !hasTime {
		return false
	}
	//是否处于时间段
	if timeTmp.GetQuanTumStart() != 0 && timeTmp.GetQuantumEnd() != 0 {
		if now <= timeTmp.GetQuanTumStart() || now >= timeTmp.GetQuantumEnd() {
			return false
		}
	}

	//每周开启的日期
	day := timeutils.MillisecondToTime(now).Weekday()
	if day == time.Sunday {
		day = 7
	}
	dayInt := int32(day)

	for _, weekday := range timeTmp.GetWeekDayArr() {
		if dayInt == weekday {
			return true
		}
	}

	monthDay := int32(timeutils.MillisecondToTime(now).Day())

	for _, tempMonthDay := range timeTmp.GetMonthDayArr() {
		if monthDay == tempMonthDay {
			return true
		}
	}
	return false
}

func (t *ActivityTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	activityTimeTemplateMap := template.GetTemplateService().GetAll((*ActivityTimeTemplate)(nil))
	for _, tempActivityTimeTemplate := range activityTimeTemplateMap {
		activityTimeTemplate := tempActivityTimeTemplate.(*ActivityTimeTemplate)
		if activityTimeTemplate.ActivityId == t.TemplateId() {
			t.timeList = append(t.timeList, activityTimeTemplate)
		}
	}

	tempMapTemplate := template.GetTemplateService().Get(int(t.Mapid), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.Mapid)
		return template.NewTemplateFieldError("Mapid", err)
	}
	t.mapTemplate = tempMapTemplate.(*MapTemplate)
	return
}
func (t *ActivityTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//活动类型
	activityType := activitytypes.ActivityType(t.ActiveType)
	if !activityType.Valid() {
		err = template.NewTemplateFieldError("ActiveType", fmt.Errorf("[%d] invalid", t.ActiveType))
		return err
	}
	t.activityType = activityType

	//活动地图id
	mapTem := template.GetTemplateService().Get(int(t.Mapid), (*MapTemplate)(nil))
	if mapTem == nil {
		return template.NewTemplateFieldError("Mapid", fmt.Errorf("[%d] invalid", t.Mapid))
	}
	//子类地图id
	if t.SubMapid != 0 {
		subMapTem := template.GetTemplateService().Get(int(t.SubMapid), (*MapTemplate)(nil))
		if subMapTem == nil {
			return template.NewTemplateFieldError("Mapid", fmt.Errorf("[%d] invalid", t.SubMapid))
		}
	}

	//参与最小等级
	err = validator.MinValidate(float64(t.LevMin), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("LevMin", fmt.Errorf("[%d] invalid", t.LevMin))
	}
	//参与最大等级
	err = validator.MaxValidate(float64(t.LevMax), float64(common.MAX_LEVEL), true)
	if err != nil {
		return template.NewTemplateFieldError("LevMax", fmt.Errorf("[%d] invalid", t.LevMax))
	}

	//验证 KaiqiId
	temObj := template.GetTemplateService().Get(int(t.KaiqiId), (*ModuleOpenedTemplate)(nil))
	if temObj == nil {
		err = fmt.Errorf("[%d] invalid", t.KaiqiId)
		err = template.NewTemplateFieldError("KaiqiId", err)
		return
	}
	funcTem := temObj.(*ModuleOpenedTemplate)
	t.funcOpenType = funcTem.GetFuncOpenType()

	//每天限制次数
	err = validator.MinValidate(float64(t.JoinCount), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("JoinCount", fmt.Errorf("[%s] invalid", t.JoinCount))
	}

	// 验证时间模板
	for _, timeTemp := range t.timeList {
		newBegin := timeTemp.beginTimeHour
		newEnd := timeTemp.endTimeHour
		for _, temp := range t.timeList {
			if timeTemp == temp {
				continue
			}

			begin := temp.beginTimeHour
			end := temp.endTimeHour
			if newEnd < begin || newBegin > end {
				continue
			}

			err = fmt.Errorf("[%d] TimeTempId invalid", t.TemplateId())
			return template.NewTemplateFieldError("TimeTempId", err)
		}
	}

	return
}

func (t *ActivityTemplate) GetFuncOpenType() funcopentypes.FuncOpenType {
	return t.funcOpenType
}

func (t *ActivityTemplate) PatchAfterCheck() {

}

func (t *ActivityTemplate) TemplateId() int {
	return t.Id
}

//是否满足等级
func (t *ActivityTemplate) IsReacheLevel(level int32) bool {
	if level < t.LevMin {
		return false
	}
	if level > t.LevMax {
		return false
	}
	return true
}

func (t *ActivityTemplate) FileName() string {
	return "tb_activity.json"
}

func init() {
	template.Register((*ActivityTemplate)(nil))
}
