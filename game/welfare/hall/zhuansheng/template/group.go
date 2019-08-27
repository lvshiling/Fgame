package template

import (
	coreutils "fgame/fgame/core/utils"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 转生冲刺
type GroupTemplateZhuanSheng struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateZhuanSheng) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	return nil
}

func (gt *GroupTemplateZhuanSheng) GetCanRewTempList(zhuanSheng int32, rewRecord []int32) (tempList []*gametemplate.OpenserverActivityTemplate) {
	for _, temp := range gt.GetOpenTempMap() {
		needZhuanSheng := temp.Value1
		if zhuanSheng < needZhuanSheng {
			continue
		}

		if coreutils.ContainInt32(rewRecord, needZhuanSheng) {
			continue
		}

		tempList = append(tempList, temp)
	}
	return
}

func CreateGroupTemplateZhuanSheng(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateZhuanSheng{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeZhaunSheng, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateZhuanSheng))
}
