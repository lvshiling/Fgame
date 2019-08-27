package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//奇遇岛
type GroupTemplateQiYuDao struct {
	*welfaretemplate.GroupTemplateBase //
}

func (gt *GroupTemplateQiYuDao) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}

	t := gt.GetFirstOpenTemp()
	to := template.GetTemplateService().Get(int(t.Value1), (*gametemplate.WelfareSceneTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] Value1 invalid", t.Value1)
		err = welfaretypes.NewWelfareRecordError(t.Id, err)
	}

	return nil
}

func CreateGroupTemplateQiYuDao(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateQiYuDao{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeQiYu, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateQiYuDao))
}
