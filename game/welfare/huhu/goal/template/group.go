package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//开服目标
type GroupTemplateGoal struct {
	*welfaretemplate.GroupTemplateBase       //
	goalId                             int32 //目标任务关联id
	goalTemp                           *gametemplate.YunYingGoalTemplate
}

func (gt *GroupTemplateGoal) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	for _, t := range gt.GetOpenTempMap() {
		err = validator.MinValidate(float64(t.Value1), float64(0), false)
		if err != nil {
			err = fmt.Errorf("Value1 [%d] invalid", t.Value1)
			err = welfaretypes.NewWelfareRecordError(t.Id, err)
			return
		}

		err = validator.MinValidate(float64(t.Value2), float64(0), false)
		if err != nil {
			err = fmt.Errorf("Value2 [%d] invalid", t.Value2)
			err = welfaretypes.NewWelfareRecordError(t.Id, err)
			return
		}

		if gt.goalId == 0 {
			gt.goalId = t.Value1
		}
		if gt.goalId != t.Value1 {
			err = fmt.Errorf("Value1 [%d] invalid", t.Value1)
			err = welfaretypes.NewWelfareRecordError(t.Id, err)
			return
		}
	}

	// 校验目标关联id
	to := template.GetTemplateService().Get(int(gt.goalId), (*gametemplate.YunYingGoalTemplate)(nil))
	if to == nil {
		return fmt.Errorf("关联目标任务id配置错误：[%d]", gt.goalId)
	}
	gt.goalTemp = to.(*gametemplate.YunYingGoalTemplate)

	return nil
}

func (gt *GroupTemplateGoal) GetGolaId() int32 {
	return gt.goalId
}

func (gt *GroupTemplateGoal) IsGoalQuest(questId int32) bool {
	questMap := gt.goalTemp.GetQuestMap()
	_, ok := questMap[questId]
	if !ok {
		return false
	}
	return true
}

func (gt *GroupTemplateGoal) GetCanRewTemplate(goalCount int32, receRecordMap map[int32]struct{}) (tempList []*gametemplate.OpenserverActivityTemplate) {
	for _, temp := range gt.GetOpenTempMap() {
		//是否领取
		rewGoalCount := temp.Value2
		if goalCount < rewGoalCount {
			continue
		}

		//领取记录
		_, ok := receRecordMap[rewGoalCount]
		if ok {
			continue
		}

		tempList = append(tempList, temp)
	}
	return tempList
}

func CreateGroupTemplateGoal(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateGoal{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeGoal, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateGoal))
}
