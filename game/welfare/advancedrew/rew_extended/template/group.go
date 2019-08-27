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

//升阶额外奖励
type GroupTemplateRewExtended struct {
	*welfaretemplate.GroupTemplateBase
	groupAdvancedType welfaretypes.AdvancedType
}

func (gt *GroupTemplateRewExtended) Init() (err error) {
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

func (gt *GroupTemplateRewExtended) GetOpenTemplateListAboutReward(advancedNum int32, recordList []int32) (openTempList []*gametemplate.OpenserverActivityTemplate) {

	for _, temp := range gt.GetOpenTempMap() {
		needNum := temp.Value2

		// 进阶等级判断
		if advancedNum < needNum {
			continue
		}

		// 是否领取过奖励
		if utils.ContainInt32(recordList, needNum) {
			continue
		}

		openTempList = append(openTempList, temp)
	}

	return
}

//返还类型
func (gt *GroupTemplateRewExtended) GetAdvancedType() welfaretypes.AdvancedType {
	return gt.groupAdvancedType
}

//升阶奖励-升阶奖励有效时间
func (gt *GroupTemplateRewExtended) GetAdvancedRewExpireTime() int64 {
	for _, temp := range gt.GetOpenTempMap() {
		return int64(temp.Value3) * int64(common.DAY)
	}

	return -1
}

func CreateGroupTemplateRewExtended(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateRewExtended{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewExtended, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRewExtended))
}
