package template

import (
	"fgame/fgame/game/global"
	gametemplate "fgame/fgame/game/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type GroupTemplateI interface {
	AddOpenTemp(temp *gametemplate.OpenserverActivityTemplate)
	GetOpenTempMap() map[int32]*gametemplate.OpenserverActivityTemplate
	GetFirstOpenTemp() *gametemplate.OpenserverActivityTemplate
	GetTimeTemplate() *gametemplate.OpenserverTimeTemplate
	GetType() welfaretypes.OpenActivityType
	GetSubType() welfaretypes.OpenActivitySubType
	GetActivityName() string
	GetActivityTime(openTime, mergeTime int64) (start int64, end int64)
	GetMaxValue1() int32
	GetMaxValue2() int32
	GetFirstValue1() int32
	GetFirstValue2() int32
	GetFirstValue3() int32
	GetFirstValue4() int32
	GetGroupId() int32
	Init() error
}

// value1排序
type SortTempListOne []*gametemplate.OpenserverActivityTemplate

func (s SortTempListOne) Len() int           { return len(s) }
func (s SortTempListOne) Less(i, j int) bool { return s[i].Value1 < s[j].Value1 }
func (s SortTempListOne) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// value2排序
type SortTempListTwo []*gametemplate.OpenserverActivityTemplate

func (s SortTempListTwo) Len() int           { return len(s) }
func (s SortTempListTwo) Less(i, j int) bool { return s[i].Value2 < s[j].Value2 }
func (s SortTempListTwo) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// value3排序
type SortTempListThree []*gametemplate.OpenserverActivityTemplate

func (s SortTempListThree) Len() int           { return len(s) }
func (s SortTempListThree) Less(i, j int) bool { return s[i].Value3 < s[j].Value3 }
func (s SortTempListThree) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type GroupTemplateBase struct {
	timeTemp    *gametemplate.OpenserverTimeTemplate
	openTempMap map[int32]*gametemplate.OpenserverActivityTemplate
}

func CreateGroupTemplateBase(timeTemp *gametemplate.OpenserverTimeTemplate) *GroupTemplateBase {
	base := &GroupTemplateBase{}
	base.timeTemp = timeTemp
	base.openTempMap = make(map[int32]*gametemplate.OpenserverActivityTemplate)
	return base
}

func (g *GroupTemplateBase) Init() error {
	return nil
}

func (g *GroupTemplateBase) GetOpenTempMap() map[int32]*gametemplate.OpenserverActivityTemplate {
	return g.openTempMap
}

func (g *GroupTemplateBase) AddOpenTemp(temp *gametemplate.OpenserverActivityTemplate) {
	tempId := int32(temp.Id)
	_, ok := g.openTempMap[tempId]
	if !ok {
		g.openTempMap[tempId] = temp
	}
}

func (g *GroupTemplateBase) GetType() welfaretypes.OpenActivityType {
	return g.timeTemp.GetOpenType()
}

func (g *GroupTemplateBase) GetGroupId() int32 {
	return g.timeTemp.Group
}

func (g *GroupTemplateBase) GetSubType() welfaretypes.OpenActivitySubType {
	return g.timeTemp.GetOpenSubType()
}

func (g *GroupTemplateBase) GetActivityName() string {
	return g.timeTemp.Name
}

func (g *GroupTemplateBase) GetActivityTime(openServerTime, mergeTime int64) (startTime int64, endTime int64) {
	now := global.GetGame().GetTimeService().Now()
	switch g.timeTemp.GetOpenTimeType() {
	case welfaretypes.OpenTimeTypeMerge:
		{
			// mergeTime := merge.GetMergeService().GetMergeTime()
			startTime, _ = g.timeTemp.GetBeginTime(now, mergeTime)
			endTime, _ = g.timeTemp.GetEndTime(now, mergeTime)
		}
	case welfaretypes.OpenTimeTypeOpenActivity,
		welfaretypes.OpenTimeTypeOpenActivityNoMerge:
		{
			// openServerTime := global.GetGame().GetServerTime()
			startTime, _ = g.timeTemp.GetBeginTime(now, openServerTime)
			endTime, _ = g.timeTemp.GetEndTime(now, openServerTime)
		}
	case welfaretypes.OpenTimeTypeNotTimeliness,
		welfaretypes.OpenTimeTypeSchedule,
		welfaretypes.OpenTimeTypeXunHuan,
		welfaretypes.OpenTimeTypeMergeXunHuan,
		welfaretypes.OpenTimeTypeWeek,
		welfaretypes.OpenTimeTypeMonth:
		{
			startTime, _ = g.timeTemp.GetBeginTime(now, 0)
			endTime, _ = g.timeTemp.GetEndTime(now, 0)
		}
	}

	return
}

func (g *GroupTemplateBase) GetMaxValue1() int32 {
	max := int32(0)
	for _, temp := range g.openTempMap {
		if temp.Value1 < max {
			continue
		}

		max = temp.Value1
	}

	return max
}

func (g *GroupTemplateBase) GetMaxValue2() int32 {
	max := int32(0)
	for _, temp := range g.openTempMap {
		if temp.Value2 < max {
			continue
		}

		max = temp.Value2
	}

	return max
}

func (g *GroupTemplateBase) GetFirstValue1() int32 {
	for _, temp := range g.openTempMap {
		return temp.Value1
	}
	return 0

}

func (g *GroupTemplateBase) GetFirstValue2() int32 {
	for _, temp := range g.openTempMap {
		return temp.Value2
	}
	return 0
}

func (g *GroupTemplateBase) GetFirstValue3() int32 {
	for _, temp := range g.openTempMap {
		return temp.Value3
	}
	return 0
}

func (g *GroupTemplateBase) GetFirstValue4() int32 {
	for _, temp := range g.openTempMap {
		return temp.Value4
	}
	return 0
}

func (g *GroupTemplateBase) GetFirstOpenTemp() *gametemplate.OpenserverActivityTemplate {
	for _, temp := range g.openTempMap {
		return temp
	}
	return nil
}

func (g *GroupTemplateBase) GetTimeTemplate() *gametemplate.OpenserverTimeTemplate {
	return g.timeTemp
}
