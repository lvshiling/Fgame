package template

import (
	"fgame/fgame/core/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//进阶祝福日14-3
type GroupTemplateAdvancedBlessCrit struct {
	*welfaretemplate.GroupTemplateBase
	groupAdvancedType welfaretypes.AdvancedType
}

func (gt *GroupTemplateAdvancedBlessCrit) Init() (err error) {
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

//系统类型
func (gt *GroupTemplateAdvancedBlessCrit) GetAdvancedType() welfaretypes.AdvancedType {
	return gt.groupAdvancedType
}

//暴击率
func (gt *GroupTemplateAdvancedBlessCrit) GetCritRate() int32 {
	return gt.GetFirstValue2()
}

//额外增加次数
func (gt *GroupTemplateAdvancedBlessCrit) GetExtralAddTimes() int32 {
	return gt.GetFirstValue3()
}

func CreateGroupTemplateAdvancedBlessCrit(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateAdvancedBlessCrit{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeBlessCrit, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateAdvancedBlessCrit))
}
