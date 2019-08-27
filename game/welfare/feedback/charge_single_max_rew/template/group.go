package template

import (
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"sort"
)

//单笔充值（最近档次）
type GroupTemplateSingleChargeMaxRew struct {
	*welfaretemplate.GroupTemplateBase
	singleTempDesc []*gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateSingleChargeMaxRew) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	for _, t := range gt.GetOpenTempMap() {
		gt.singleTempDesc = append(gt.singleTempDesc, t)
	}

	//排序
	sort.Sort(sort.Reverse(welfaretemplate.SortTempListOne(gt.singleTempDesc)))

	return
}

//单笔充值奖励模板
func (gt *GroupTemplateSingleChargeMaxRew) GetTempDescList() []*gametemplate.OpenserverActivityTemplate {
	return gt.singleTempDesc
}

func CreateGroupTemplateSingleChargeMaxRew(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateSingleChargeMaxRew{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateSingleChargeMaxRew))
}
