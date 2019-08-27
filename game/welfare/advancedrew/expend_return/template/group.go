package template

import (
	"fgame/fgame/core/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//进阶消耗返还14-4
type GroupTemplateAdvancedExpendReturn struct {
	*welfaretemplate.GroupTemplateBase
	groupAdvancedType welfaretypes.AdvancedType
}

func (gt *GroupTemplateAdvancedExpendReturn) Init() (err error) {
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
func (gt *GroupTemplateAdvancedExpendReturn) GetAdvancedType() welfaretypes.AdvancedType {
	return gt.groupAdvancedType
}

func CreateGroupTemplateAdvancedExpendReturn(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateAdvancedExpendReturn{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeExpendReturn, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateAdvancedExpendReturn))
}
