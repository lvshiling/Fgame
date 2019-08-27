package template

import (
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"sort"
)

//排行榜模板
type GroupTemplateRank struct {
	*welfaretemplate.GroupTemplateBase
	rankSortList groupSortRankList
}

//分组模板排序类型
type groupSortRankList []*gametemplate.OpenserverActivityTemplate

func (s groupSortRankList) Len() int           { return len(s) }
func (s groupSortRankList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s groupSortRankList) Less(i, j int) bool { return s[i].Value1 < s[j].Value1 }

//获取分组模板排序的groupTemp切片
func (gt *GroupTemplateRank) Init() error {
	openLen := len(gt.GetOpenTempMap())
	groupTempList := make(groupSortRankList, 0, openLen)
	for _, temp := range gt.GetOpenTempMap() {
		groupTempList = append(groupTempList, temp)
	}
	sort.Sort(groupTempList)
	gt.rankSortList = groupTempList
	return nil
}

//排行榜名次配置
func (gt *GroupTemplateRank) GetRankRewardsOpenTemp(ranking, condition int32) (isOnLevel bool, temp *gametemplate.OpenserverActivityTemplate) {
	isOnLevel = true
	for _, temp = range gt.rankSortList {
		//是否满足排名
		if ranking > temp.Value1 && ranking > temp.Value2 {
			continue
		}

		// 是否满足排名领取条件
		rewCondition := temp.Value3
		if condition < rewCondition {
			isOnLevel = false
			continue
		}
		return
	}
	return
}

//排行榜最低入榜条件
func (gt *GroupTemplateRank) GetRankMinLimitCondition() int32 {
	rankLen := len(gt.rankSortList)
	openTemp := gt.rankSortList[rankLen-1]
	return openTemp.Value3
}

func CreateGroupTemplateRank(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateRank{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeCharge, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeCost, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeMount, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeWing, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeBodyshield, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingyu, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeShenfa, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeFeather, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeShield, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeCharm, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeAnqi, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeNumber, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeFaBao, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeXianTi, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeShiHunFan, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeTianMoTi, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingBing, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingQi, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingYi, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingBao, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingTi, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingYu, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingShen, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLevel, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeMarryDevelop, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeNumberDay, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeGoldEquipForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingTongForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeDianXingForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeShenQiForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeMingGeForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeShengHenForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeZhenFaForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeTuLongEquipForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeBabyForce, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeZhuanSheng, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRank))
}
