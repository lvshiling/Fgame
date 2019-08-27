package template

import (
	"fgame/fgame/core/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//进阶次数返还14-5
type GroupTemplateAdvancedTimesReturn struct {
	*welfaretemplate.GroupTemplateBase
	groupAdvancedType welfaretypes.AdvancedType
}

func (gt *GroupTemplateAdvancedTimesReturn) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	// 进阶类型
	firstTemp := gt.GetFirstOpenTemp()
	if firstTemp != nil {
		gt.groupAdvancedType = welfaretypes.AdvancedType(firstTemp.Value1)
	}

	for _, t := range gt.GetOpenTempMap() {
		// 校验进阶类型
		advancedType := welfaretypes.AdvancedType(t.Value1)
		if !advancedType.Valid() || advancedType != gt.groupAdvancedType {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}
	}

	return
}

//返还类型
func (gt *GroupTemplateAdvancedTimesReturn) GetAdvancedType() welfaretypes.AdvancedType {
	return gt.groupAdvancedType
}

func CreateGroupTemplateAdvancedTimesReturn(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateAdvancedTimesReturn{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeTimesReturn, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateAdvancedTimesReturn))
}
