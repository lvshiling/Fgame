package template

import (
	"fgame/fgame/core/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//升阶奖励
type GroupTemplateRew struct {
	*welfaretemplate.GroupTemplateBase
	groupAdvancedType welfaretypes.AdvancedType
}

func (gt *GroupTemplateRew) Init() (err error) {
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
func (gt *GroupTemplateRew) GetAdvancedType() welfaretypes.AdvancedType {
	return gt.groupAdvancedType
}

func CreateGroupTemplateRew(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateRew{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRew, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRew))
}
