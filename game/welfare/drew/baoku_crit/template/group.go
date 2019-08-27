package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//装备宝库暴击日7-6
type GroupTemplateDrewBaoKuCrit struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateDrewBaoKuCrit) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}

	return
}

//幸运值暴击数
func (gt *GroupTemplateDrewBaoKuCrit) GetLuckyPointCritNum() int32 {
	return gt.GetFirstValue1()
}

//积分暴击数
func (gt *GroupTemplateDrewBaoKuCrit) GetAttendPointCritNum() int32 {
	return gt.GetFirstValue2()
}

//暴击率
func (gt *GroupTemplateDrewBaoKuCrit) GetCritRate() int32 {
	return gt.GetFirstValue3()
}

func CreateGroupTemplateDrewBaoKuCrit(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateDrewBaoKuCrit{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeBaoKuCrit, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateDrewBaoKuCrit))
}
