package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//聚宝盆
type GroupTemplateGoldBowl struct {
	*welfaretemplate.GroupTemplateBase //
}

func (gt *GroupTemplateGoldBowl) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}

	return nil
}

//聚宝盆返利比例
func (gt *GroupTemplateGoldBowl) GetGoldBowlRate() int32 {
	return gt.GetFirstValue1()
}

func CreateGroupTemplateGoldBowl(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	g := &GroupTemplateGoldBowl{}
	g.GroupTemplateBase = base
	return g
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldBowl, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateGoldBowl))
}
