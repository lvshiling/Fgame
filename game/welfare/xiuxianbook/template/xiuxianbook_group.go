package template

import (
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"sort"
)

type GroupTemplateXiuxianBook struct {
	*welfaretemplate.GroupTemplateBase
	groupTempList groupXiuxianBookList
}

//分组模板排序类型
type groupXiuxianBookList []*gametemplate.OpenserverActivityTemplate

func (s groupXiuxianBookList) Len() int           { return len(s) }
func (s groupXiuxianBookList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s groupXiuxianBookList) Less(i, j int) bool { return s[i].Value1 < s[j].Value1 }

//获取分组模板排序的groupTemp切片
func (gt *GroupTemplateXiuxianBook) Init() error {
	openLen := len(gt.GetOpenTempMap())
	groupTempList := make(groupXiuxianBookList, 0, openLen)
	for _, temp := range gt.GetOpenTempMap() {
		groupTempList = append(groupTempList, temp)
	}
	sort.Sort(groupTempList)
	gt.groupTempList = groupTempList
	return nil
}

func (gt *GroupTemplateXiuxianBook) GetCanReceiveList(level, chargeNum int32) []int32 {
	list := []int32{}
	for i := 0; i < len(gt.groupTempList); i++ {
		temp := gt.groupTempList[i]
		needLevel := temp.Value1
		needChargeNum := temp.Value2
		if level >= needLevel && chargeNum >= needChargeNum {
			list = append(list, needLevel)
		}
	}
	return list
}

func (gt *GroupTemplateXiuxianBook) GetGroupXiuxianBookList() []*gametemplate.OpenserverActivityTemplate {
	return gt.groupTempList
}

func CreateGroupTemplateXiuxianBook(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateXiuxianBook{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipStrength, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateXiuxianBook))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateXiuxianBook))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipUpStar, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateXiuxianBook))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeLingTong, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateXiuxianBook))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeDianXing, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateXiuxianBook))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeShenQi, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateXiuxianBook))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillXinFa, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateXiuxianBook))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillDiHun, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateXiuxianBook))
}
