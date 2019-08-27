package template

import (
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"sort"
)

//每日充值
type GroupTemplateCycleSingleMaxRew struct {
	*welfaretemplate.GroupTemplateBase
	cycSingleTempMap map[int32][]*gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateCycleSingleMaxRew) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	gt.cycSingleTempMap = make(map[int32][]*gametemplate.OpenserverActivityTemplate)
	for _, t := range gt.GetOpenTempMap() {
		dayKey := t.Value1
		gt.cycSingleTempMap[dayKey] = append(gt.cycSingleTempMap[dayKey], t)
	}

	//排序
	for _, tempList := range gt.cycSingleTempMap {
		sort.Sort(sort.Reverse(welfaretemplate.SortTempListTwo(tempList)))
	}

	return
}

//每日单笔充值奖励模板
func (gt *GroupTemplateCycleSingleMaxRew) GetCurDayTempDescList(day int32) []*gametemplate.OpenserverActivityTemplate {
	tempList, ok := gt.cycSingleTempMap[day]
	if !ok {
		return nil
	}

	return tempList
}

func CreateGroupTemplateCycleSingleMaxRew(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateCycleSingleMaxRew{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRew, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateCycleSingleMaxRew))
}
