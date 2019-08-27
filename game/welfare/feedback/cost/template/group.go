package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//消费返利
type GroupTemplateCost struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateCost) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	return
}

func CreateGroupTemplateCost(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateCost{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCost, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateCost))
}
