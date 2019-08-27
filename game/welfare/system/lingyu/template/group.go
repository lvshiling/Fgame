package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//系统激活-领域
type GroupTemplateSystemLingYu struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateSystemLingYu) Init() (err error) {
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

//返还类型
func (gt *GroupTemplateSystemLingYu) GetActivateCondition() int32 {
	return gt.GetFirstValue1()
}

func CreateGroupTemplateSystemLingYu(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateSystemLingYu{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeSystemActivate, welfaretypes.OpenActivitySystemActivateSubTypeLingYu, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateSystemLingYu))
}
