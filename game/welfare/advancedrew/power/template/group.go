package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//升阶战力奖励
type GroupTemplatePower struct {
	*welfaretemplate.GroupTemplateBase
	groupAdvancedType welfaretypes.AdvancedType
}

func (gt *GroupTemplatePower) Init() (err error) {
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
		// 校验进阶类型
		advancedType := welfaretypes.AdvancedType(t.Value1)
		if !advancedType.Valid() || advancedType != gt.groupAdvancedType {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}
	}

	return
}

func (gt *GroupTemplatePower) GetOpenTemplateListAboutReward(power int64, recordList []int64) (openTempList []*gametemplate.OpenserverActivityTemplate) {

	for _, temp := range gt.GetOpenTempMap() {
		needNum := temp.Value2

		// 战斗力判断
		if power < int64(needNum) {
			continue
		}

		// 是否领取过奖励
		if utils.ContainInt64(recordList, int64(needNum)) {
			continue
		}

		openTempList = append(openTempList, temp)
	}

	return
}

//升阶奖励-升阶奖励有效时间
func (gt *GroupTemplatePower) GetAdvancedRewExpireTime() int64 {
	for _, temp := range gt.GetOpenTempMap() {
		return int64(temp.Value3) * int64(common.DAY)
	}

	return -1
}

//升阶奖励-升阶类型
func (gt *GroupTemplatePower) GetAdvancedType() welfaretypes.AdvancedType {
	return gt.groupAdvancedType
}

func CreateGroupTemplatePower(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplatePower{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypePower, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplatePower))
}
