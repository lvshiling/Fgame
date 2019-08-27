package template

import (
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"sort"
)

type GroupTemplateCycleSingleAllMultiple struct {
	*welfaretemplate.GroupTemplateBase
	//时间索引
	cycSingleTempMap map[int32][]*gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateCycleSingleAllMultiple) Init() (err error) {
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

	//降序排序
	for _, tempList := range gt.cycSingleTempMap {
		sort.Sort(sort.Reverse(welfaretemplate.SortTempListTwo(tempList)))
	}

	return
}

//每日单笔充值奖励模板
func (gt *GroupTemplateCycleSingleAllMultiple) GetCurDayTempDescList(day int32) []*gametemplate.OpenserverActivityTemplate {
	tempList, ok := gt.cycSingleTempMap[day]
	if !ok {
		return nil
	}

	return tempList
}

//每日单笔充值奖励模板
func (gt *GroupTemplateCycleSingleAllMultiple) GetCurDayChargeNumTemp(day, needNum int32) *gametemplate.OpenserverActivityTemplate {
	tempList, ok := gt.cycSingleTempMap[day]
	if !ok {
		return nil
	}
	for _, temp := range tempList {
		if temp.Value2 == needNum {
			return temp
		}
	}

	return nil
}

//获取可以领取的奖励
func (gt *GroupTemplateCycleSingleAllMultiple) GetCanRewRecordMap(day int32, remainTimesMap map[int32]int32) map[int32]int32 {
	canRewRecordMap := make(map[int32]int32)
	descTempList := gt.GetCurDayTempDescList(day)

	for goldNum, times := range remainTimesMap {
		for _, temp := range descTempList {
			needGold := temp.Value2
			if goldNum < needGold {
				continue
			}
			num := goldNum / needGold
			canRewRecordMap[needGold] += num * times
		}
	}
	return canRewRecordMap
}

//获取使用的次数
func (gt *GroupTemplateCycleSingleAllMultiple) GetSingleGoldCanRewRecordMap(day int32, goldNum int32, remainTimesMap map[int32]int32) (useTimes map[int32]int32, rewTimes int32, flag bool) {
	useTimes = make(map[int32]int32)

	openserverActivityTemplate := gt.GetCurDayChargeNumTemp(day, goldNum)
	if openserverActivityTemplate == nil {
		return
	}

	for gold, timesNum := range remainTimesMap {
		if gold < goldNum {
			continue
		}

		rewTimes += gold / goldNum * timesNum
		useTimes[gold] += timesNum
	}
	if len(useTimes) != 0 {
		flag = true
	}
	return
}

func CreateGroupTemplateCycleAllMultiple(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateCycleSingleAllMultiple{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateCycleAllMultiple))
}
