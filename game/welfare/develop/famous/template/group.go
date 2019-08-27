package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//名人普
type GroupTemplateFamous struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateFamous) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	famousTemp := welfaretemplate.GetWelfareTemplateService().GetFamousTemplate(gt.GetGroupId())
	if famousTemp == nil {
		return fmt.Errorf("tb_famous:名人普奖励没有配置")
	}

	return
}

func CreateGroupTemplateFamous(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateFamous{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeDevelop, welfaretypes.OpenActivityDefaultSubTypeDefault, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateFamous))
}
