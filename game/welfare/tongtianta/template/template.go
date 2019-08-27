package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"sort"
)

// 通天塔通用模板
type GroupTemplateTongTianTa struct {
	*welfaretemplate.GroupTemplateBase
	tongTianTaTemplateAscList groupSortTongTianTaList
}

// 升序
type groupSortTongTianTaList []*gametemplate.OpenserverActivityTemplate

func (s groupSortTongTianTaList) Len() int           { return len(s) }
func (s groupSortTongTianTaList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s groupSortTongTianTaList) Less(i, j int) bool { return s[i].Value1 < s[j].Value1 }

// 初始化排序
func (g *GroupTemplateTongTianTa) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(g.GetActivityName(), g.GetGroupId(), err)
			return
		}
	}()

	openLen := len(g.GetOpenTempMap())
	groupTempList := make(groupSortTongTianTaList, 0, openLen)
	for _, temp := range g.GetOpenTempMap() {
		// 验证value_1
		err = validator.MinValidate(float64(temp.Value1), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", temp.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}

		groupTempList = append(groupTempList, temp)
	}
	sort.Sort(groupTempList)
	g.tongTianTaTemplateAscList = groupTempList
	return
}

// 获取接近战力的一档奖励模板
func (g *GroupTemplateTongTianTa) GetTongTianTaTemplateByNearForce(force int32) (lastTemp *gametemplate.OpenserverActivityTemplate) {
	for _, temp := range g.tongTianTaTemplateAscList {
		if temp.Value1 > force {
			return
		}
		lastTemp = temp
	}
	return
}

func (g *GroupTemplateTongTianTa) GetTongTianTaForceTemplateListByForce(minForce int32, maxForce int32) (list []*gametemplate.OpenserverActivityTemplate) {
	for _, temp := range g.tongTianTaTemplateAscList {
		if temp.Value1 >= minForce && temp.Value1 <= maxForce {
			list = append(list, temp)
		}
	}
	return list
}

func CreateGroupTemplateTongTianTa(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateTongTianTa{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {

	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeLingTong, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateTongTianTa))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeMingGe, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateTongTianTa))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeShengHen, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateTongTianTa))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeTuLong, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateTongTianTa))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeZhenFa, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateTongTianTa))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeDianXing, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateTongTianTa))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeBaby, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateTongTianTa))
}
