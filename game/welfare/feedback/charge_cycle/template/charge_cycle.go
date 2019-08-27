package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	feedbackchargecycletypes "fgame/fgame/game/welfare/feedback/charge_cycle/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//连续充值
type GroupTemplateCycleCharge struct {
	*welfaretemplate.GroupTemplateBase
	dayTempMap      map[int32]*gametemplate.OpenserverActivityTemplate //每日奖励模板
	totalDayTempMap map[int32]*gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateCycleCharge) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	gt.dayTempMap = make(map[int32]*gametemplate.OpenserverActivityTemplate)
	gt.totalDayTempMap = make(map[int32]*gametemplate.OpenserverActivityTemplate)
	for _, t := range gt.GetOpenTempMap() {
		//返利连续充值奖励类型
		rewType := feedbackchargecycletypes.FeedbackCycleRewType(t.Value1)
		if !rewType.Valid() {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}

		switch rewType {
		case feedbackchargecycletypes.FeedbackCycleRewTypeSingleDay:
			{
				_, ok := gt.dayTempMap[t.Value2]
				if ok {
					err = fmt.Errorf("[%d] invalid", t.Value2)
					err = template.NewTemplateFieldError("Value2", err)
					return
				}
				gt.dayTempMap[t.Value2] = t
			}
		case feedbackchargecycletypes.FeedbackCycleRewTypeCountDay:
			{
				_, ok := gt.totalDayTempMap[t.Value2]
				if ok {
					err = fmt.Errorf("[%d] invalid", t.Value2)
					err = template.NewTemplateFieldError("Value2", err)
					return
				}
				gt.totalDayTempMap[t.Value2] = t
			}
		}
	}

	return
}

// 连续充值结束奖励
func (gt *GroupTemplateCycleCharge) GetEndRewTempMap() (rewTempMap map[int32]*gametemplate.OpenserverActivityTemplate) {
	return gt.totalDayTempMap
}

// 每日奖励模板
func (gt *GroupTemplateCycleCharge) GetCrossDayRewTemp(dayNum int32) *gametemplate.OpenserverActivityTemplate {
	for _, temp := range gt.dayTempMap {
		if temp.Value2 == dayNum {
			return temp
		}
	}

	return nil
}

//每日奖励充值条件
func (gt *GroupTemplateCycleCharge) GetDayRewCondition(dayNum int32) int32 {
	for _, temp := range gt.dayTempMap {
		if temp.Value2 == dayNum {
			return temp.Value3
		}
	}

	return 0
}

func CreateGroupTemplateCycleCharge(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateCycleCharge{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCycleCharge, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateCycleCharge))
}
