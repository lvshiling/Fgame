package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sort"
)

//疯狂宝箱
type GroupTemplateCrazyBox struct {
	*welfaretemplate.GroupTemplateBase
	crazyBoxSortList groupSortCrazyBoxList
	weightList       []int64
}

func CreateGroupTemplateCrazyBox(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateCrazyBox{}
	gt.GroupTemplateBase = base
	return gt
}

//分组模板排序类型
type groupSortCrazyBoxList []*gametemplate.OpenserverActivityTemplate

func (s groupSortCrazyBoxList) Len() int           { return len(s) }
func (s groupSortCrazyBoxList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s groupSortCrazyBoxList) Less(i, j int) bool { return s[i].Value1 < s[j].Value1 }

//获取分组模板排序的groupTemp切片
func (gt *GroupTemplateCrazyBox) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	//验证 抽奖单价
	for _, t := range gt.GetOpenTempMap() {
		err = validator.MinValidate(float64(t.Value2), float64(0), false)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value2)
			err = template.NewTemplateFieldError("Value2", err)
			return
		}
	}

	openLen := len(gt.GroupTemplateBase.GetOpenTempMap())
	groupTempList := make(groupSortCrazyBoxList, 0, openLen)
	gt.weightList = make([]int64, 0, openLen)
	for _, temp := range gt.GroupTemplateBase.GetOpenTempMap() {
		groupTempList = append(groupTempList, temp)
		gt.weightList = append(gt.weightList, int64(1))
	}
	sort.Sort(groupTempList)
	gt.crazyBoxSortList = groupTempList
	return
}

//获取疯狂宝箱总开箱次数(已用和未用的)
func (gt *GroupTemplateCrazyBox) GetCrazyBoxTotalTimes(costNum int32) int32 {
	totalTimes := int32(0)
	for _, temp := range gt.crazyBoxSortList {
		levCost := temp.Value2 * temp.Value3
		if costNum < levCost {
			totalTimes += costNum / temp.Value2
			break
		}
		totalTimes += temp.Value3
		costNum -= levCost
	}
	return totalTimes
}

//获取疯狂宝箱等级和当前宝箱剩开箱次数
func (gt *GroupTemplateCrazyBox) GetCrazyBoxArg(useTimes int32) (int32, int32) {
	curBoxLev := int32(0)
	leftTimes := int32(0)
	for _, temp := range gt.crazyBoxSortList {
		if useTimes < temp.Value3 {
			curBoxLev = temp.Value1
			leftTimes = temp.Value3 - useTimes
			break
		}
		useTimes -= temp.Value3
		curBoxLev = temp.Value1
	}
	return curBoxLev, leftTimes
}

//获取当前等级疯狂宝箱配置
func (gt *GroupTemplateCrazyBox) GetOpenActivityCrazyBox(level int32) *gametemplate.OpenserverActivityTemplate {
	for _, temp := range gt.crazyBoxSortList {
		if temp.Value1 == level {
			return temp
		}
	}
	return nil
}

//随机疯狂宝箱配置
func (gt *GroupTemplateCrazyBox) GetRandomOpenActivityCrazyBoxLevel() int32 {
	index := mathutils.RandomWeights(gt.weightList)
	if index == -1 {
		return 0
	}
	ch := gt.crazyBoxSortList[index]
	return ch.Value1
}

//获取当前等级疯狂宝箱开箱次数上限
func (gt *GroupTemplateCrazyBox) GetOpenActivityCrazyBoxUpTimes(level int32) int32 {
	for _, temp := range gt.crazyBoxSortList {
		if temp.Value1 == level {
			return temp.Value3
		}
	}
	return 0
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCrazyBox, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateCrazyBox))
}
