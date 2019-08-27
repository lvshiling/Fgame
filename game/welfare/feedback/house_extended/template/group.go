package template

import (
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	houseextendedtypes "fgame/fgame/game/welfare/feedback/house_extended/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"sort"
)

//房产活动
type GroupTemplateHouseExtended struct {
	*welfaretemplate.GroupTemplateBase
	houseRewTempMap map[houseextendedtypes.HouseRewType][]*gametemplate.OpenserverActivityTemplate
}

func (gt *GroupTemplateHouseExtended) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	for _, temp := range gt.GetOpenTempMap() {
		rewType := houseextendedtypes.HouseRewType(temp.Value1)
		if !rewType.Valid() {
			err = fmt.Errorf("value1 [%d] error ", temp.Value1)
			return welfaretypes.NewWelfareRecordError(temp.Id, err)
		}

		err = validator.MinValidate(float64(temp.Value2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("Value2 [%d] error", temp.Value2)
			return welfaretypes.NewWelfareRecordError(temp.Id, err)
		}

		gt.houseRewTempMap[rewType] = append(gt.houseRewTempMap[rewType], temp)
	}

	// 校验装修礼包等级差值
	upLevelRewTempList := gt.houseRewTempMap[houseextendedtypes.HouseRewTypeUplevel]
	sort.Sort(welfaretemplate.SortTempListThree(upLevelRewTempList))
	preLevel := int32(0)
	for index, temp := range upLevelRewTempList {
		if index == 0 {
			preLevel = temp.Value3
			continue
		}

		curLevel := temp.Value3
		diff := curLevel - preLevel
		if diff != 1 {
			err = fmt.Errorf("Value3 [%d] error", temp.Value3)
			return welfaretypes.NewWelfareRecordError(temp.Id, err)
		}
		preLevel = curLevel
	}

	// 校验激活礼包
	activateRewTempList := gt.houseRewTempMap[houseextendedtypes.HouseRewTypeActivate]
	if len(activateRewTempList) != 1 {
		return fmt.Errorf("房产活动：激活礼包应该有且只有一条")
	}

	return
}

func (gt *GroupTemplateHouseExtended) GetActivateRewTemp(chargeNum int32) *gametemplate.OpenserverActivityTemplate {
	temp := gt.houseRewTempMap[houseextendedtypes.HouseRewTypeActivate][0]

	return temp
}

func (gt *GroupTemplateHouseExtended) GetActivateCanRewTemp(chargeNum int32) *gametemplate.OpenserverActivityTemplate {
	temp := gt.houseRewTempMap[houseextendedtypes.HouseRewTypeActivate][0]
	needCharge := temp.Value2
	if chargeNum < needCharge {
		return nil
	}

	return temp
}

func (gt *GroupTemplateHouseExtended) GetUplevelCanRewTemp(chargeNum, curLevel int32) *gametemplate.OpenserverActivityTemplate {
	for _, temp := range gt.houseRewTempMap[houseextendedtypes.HouseRewTypeUplevel] {
		level := temp.Value3
		if curLevel != level {
			continue
		}

		needCharge := temp.Value2
		if chargeNum < needCharge {
			continue
		}

		return temp
	}

	return nil
}

func CreateGroupTemplateHouseExtended(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	g := &GroupTemplateHouseExtended{}
	g.GroupTemplateBase = base
	g.houseRewTempMap = make(map[houseextendedtypes.HouseRewType][]*gametemplate.OpenserverActivityTemplate)
	return g
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseExtended, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateHouseExtended))
}
