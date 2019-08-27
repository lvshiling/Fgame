package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//经验炼制
type GroupTemplateMadeExp struct {
	*welfaretemplate.GroupTemplateBase //
}

func (gt *GroupTemplateMadeExp) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}

	return nil
}

//炼制等级条件
func (gt *GroupTemplateMadeExp) GetMadeLevelLimit() int32 {
	return gt.GetFirstValue1()
}

func CreateGroupTemplateMadeExp(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateMadeExp{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMade, welfaretypes.OpenActivityMadeSubTypeResource, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateMadeExp))
}
