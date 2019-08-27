package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	drewcommontypes "fgame/fgame/game/welfare/drew/common/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//幸运炸矿
type GroupTemplateDrewBomb struct {
	*welfaretemplate.GroupTemplateBase
	drewTypeMap map[drewcommontypes.LuckyDrewAttendType]*gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateDrewBomb) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	//幸运炸矿类型
	gt.drewTypeMap = make(map[drewcommontypes.LuckyDrewAttendType]*gametemplate.OpenserverActivityTemplate)
	for _, t := range gt.GetOpenTempMap() {
		drewAttendType := drewcommontypes.LuckyDrewAttendType(t.Value1)
		if !drewAttendType.Valid() {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}

		_, ok := gt.drewTypeMap[drewAttendType]
		if ok {
			return fmt.Errorf("重复配置参与方式，attendType:%d", t.Value1)
		}
		gt.drewTypeMap[drewAttendType] = t

	}
	return
}

//幸运炸矿消耗元宝
func (gt *GroupTemplateDrewBomb) GetLuckyDrewNeedGold(attendType drewcommontypes.LuckyDrewAttendType) int64 {
	temp, ok := gt.drewTypeMap[attendType]
	if !ok {
		return 0
	}

	return int64(temp.Value2)
}

func CreateGroupTemplateDrewBomb(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	g := &GroupTemplateDrewBomb{}
	g.GroupTemplateBase = base
	return g
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeBombOre, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateDrewBomb))
}
