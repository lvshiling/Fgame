package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//赛龙舟战力
type GroupTemplateBoatRaceForce struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateBoatRaceForce) Init() (err error) {
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

func CreateGroupTemplateBoatRaceForce(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateBoatRaceForce{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeBoatRace, welfaretypes.OpenActivityDefaultSubTypeDefault, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateBoatRaceForce))
}
