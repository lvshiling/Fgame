package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"sort"
)

type GroupTemplateRewPools struct {
	*welfaretemplate.GroupTemplateBase
	luckDrewTemplateList groupDrewTempList
}

//分组模板排序类型
type groupDrewTempList []*gametemplate.LuckyDrewTemplate

func (s groupDrewTempList) Len() int           { return len(s) }
func (s groupDrewTempList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s groupDrewTempList) Less(i, j int) bool { return s[i].Level < s[j].Level }

func CreateGroupTemplateRewPools(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateRewPools{}
	gt.GroupTemplateBase = base
	return gt
}

//校验luck模板
func (gt *GroupTemplateRewPools) Init() (err error) {
	welfaretemplateService := welfaretemplate.GetWelfareTemplateService()
	groupId := gt.GetGroupId()
	gt.luckDrewTemplateList = []*gametemplate.LuckyDrewTemplate{}

	forwardRate, backRate := int32(-1), int32(-1)
	for i := int32(0); forwardRate != int32(0); i++ {
		luckDrewTemp := welfaretemplateService.GetLuckDrewTemplateByArg(groupId, i)
		if luckDrewTemp == nil {
			err = fmt.Errorf("[%d] invalid", i)
			err = template.NewTemplateFieldError("tb_yunying_choujiang Level", err)
			return
		}
		forwardRate = luckDrewTemp.Percent1
		backRate = luckDrewTemp.Percent2
		if i == int32(0) && backRate != 0 {
			err = fmt.Errorf("[%d] invalid", backRate)
			err = template.NewTemplateFieldError("tb_yunying_choujiang Percent2", err)
			return
		}
		gt.luckDrewTemplateList = append(gt.luckDrewTemplateList, luckDrewTemp)
	}
	sort.Sort(gt.luckDrewTemplateList)
	return
}

func (gt *GroupTemplateRewPools) GetLuckDrewTemp() groupDrewTempList {
	return gt.luckDrewTemplateList
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeRewPools, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRewPools))
}
