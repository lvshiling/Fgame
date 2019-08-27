package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	feedbackchargedeveloptypes "fgame/fgame/game/welfare/feedback/charge_develop/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//金鸡培育
type GroupTemplateDevelop struct {
	*welfaretemplate.GroupTemplateBase
	dayTempMap    map[int32]*gametemplate.OpenserverActivityTemplate //每日奖励模板
	conditionTemp *gametemplate.OpenserverActivityTemplate
	totalDayTemp  *gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateDevelop) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	gt.dayTempMap = make(map[int32]*gametemplate.OpenserverActivityTemplate)

	for _, t := range gt.GetOpenTempMap() {
		//培养奖励类型
		rewType := feedbackchargedeveloptypes.FeedbackDevelopRewType(t.Value1)
		if !rewType.Valid() {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}

		switch rewType {
		case feedbackchargedeveloptypes.FeedbackDevelopRewTypeSingleDay:
			{
				_, ok := gt.dayTempMap[t.Value2]
				if ok {
					err = fmt.Errorf("[%d] invalid", t.Value2)
					err = template.NewTemplateFieldError("Value2", err)
					return
				}
				gt.dayTempMap[t.Value2] = t
			}
		case feedbackchargedeveloptypes.FeedbackDevelopRewTypeCondition:
			{
				if gt.conditionTemp != nil {
					err = fmt.Errorf("[%d] invalid", t.Value1)
					err = template.NewTemplateFieldError("Value1", err)
					return
				}
				gt.conditionTemp = t
			}
		case feedbackchargedeveloptypes.FeedbackDevelopRewTypeCountDay:
			{
				if gt.totalDayTemp != nil {
					err = fmt.Errorf("[%d] invalid", t.Value1)
					err = template.NewTemplateFieldError("Value1", err)
					return
				}
				gt.totalDayTemp = t
			}
		}
	}

	if gt.totalDayTemp == nil {
		return fmt.Errorf("没有总奖励配置")
	}
	if gt.conditionTemp == nil {
		return fmt.Errorf("没有条件配置")
	}

	return nil
}

// 金鸡培养结束奖励
func (gt *GroupTemplateDevelop) GetDevelopEndRewTemp() *gametemplate.OpenserverActivityTemplate {
	return gt.totalDayTemp
}

// 金鸡培养每日奖励模板
func (gt *GroupTemplateDevelop) GetDevelopDayRewTemp(feedTimes int32) *gametemplate.OpenserverActivityTemplate {
	for _, temp := range gt.dayTempMap {
		if temp.Value2 == feedTimes {
			return temp
		}
	}

	return nil
}

//返利金鸡每日培养条件
func (gt *GroupTemplateDevelop) GetDevelopFeedCondition(feedTimes int32) int64 {
	for _, temp := range gt.dayTempMap {
		if temp.Value2 == feedTimes {
			return int64(temp.Value3)
		}
	}

	return 0
}

//返利金鸡培养累计条件
func (gt *GroupTemplateDevelop) GetDevelopNeedTotalTimes() int32 {
	if gt.totalDayTemp == nil {
		return 0
	}

	return gt.totalDayTemp.Value2
}

//金鸡复活/激活价格
func (gt *GroupTemplateDevelop) GetReviveNeedGold() int64 {
	if gt.conditionTemp == nil {
		return 0
	}

	return int64(gt.conditionTemp.Value2)

}

func CreateGroupTemplateDevelop(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateDevelop{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateDevelop))
}
