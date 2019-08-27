package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"sort"
)

type GroupTemplateRewMax struct {
	*welfaretemplate.GroupTemplateBase
	groupAdvancedType welfaretypes.AdvancedType
	rewDescTempList   []*gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateRewMax) Init() (err error) {
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

		gt.rewDescTempList = append(gt.rewDescTempList, t)
	}

	//排序
	sort.Sort(sort.Reverse(welfaretemplate.SortTempListTwo(gt.rewDescTempList)))

	return
}

//返还类型
func (gt *GroupTemplateRewMax) GetAdvancedType() welfaretypes.AdvancedType {
	return gt.groupAdvancedType
}

//升阶奖励模板
func (gt *GroupTemplateRewMax) GetRewTempDescList() []*gametemplate.OpenserverActivityTemplate {
	return gt.rewDescTempList
}

func CreateGroupTemplateRewMax(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateRewMax{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewMax, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRewMax))
}
