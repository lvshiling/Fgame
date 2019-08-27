package template

import (
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//每日充值
type GroupTemplateCycleCharge struct {
	*welfaretemplate.GroupTemplateBase
	cycTempMap map[int32][]*gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateCycleCharge) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	gt.cycTempMap = make(map[int32][]*gametemplate.OpenserverActivityTemplate)
	for _, t := range gt.GetOpenTempMap() {
		dayKey := t.Value1
		gt.cycTempMap[dayKey] = append(gt.cycTempMap[dayKey], t)
	}

	return
}

//每日充值奖励模板
func (gt *GroupTemplateCycleCharge) GetCurDayTempList(day int32) []*gametemplate.OpenserverActivityTemplate {
	tempList, ok := gt.cycTempMap[day]
	if !ok {
		return nil
	}

	return tempList
}

func CreateGroupTemplateCycleCharge(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateCycleCharge{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeCharge, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateCycleCharge))
}
