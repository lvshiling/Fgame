package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	investleveltypes "fgame/fgame/game/welfare/invest/level/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//投资计划
type GroupTemplateInvestLevel struct {
	*welfaretemplate.GroupTemplateBase
	levelInvestTempMap map[investleveltypes.InvestLevelType][]*gametemplate.OpenserverActivityTemplate
}

func CreateGroupTemplateInvestLevel(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateInvestLevel{}
	gt.GroupTemplateBase = base
	return gt
}

func (gt *GroupTemplateInvestLevel) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	gt.levelInvestTempMap = make(map[investleveltypes.InvestLevelType][]*gametemplate.OpenserverActivityTemplate)
	for _, t := range gt.GetOpenTempMap() {
		//投资计划类型
		investType := investleveltypes.InvestLevelType(t.Value1)
		if !investType.Valid() {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}

		//等级
		err = validator.MinValidate(float64(t.Value2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value2)
			err = template.NewTemplateFieldError("Value2", err)
			return
		}

		gt.levelInvestTempMap[investType] = append(gt.levelInvestTempMap[investType], t)

		// 校验 购买价格
		for _, tempList := range gt.levelInvestTempMap {
			if len(tempList) < 1 {
				continue
			}
			initPrice := tempList[0].Value3
			for _, temp := range tempList {
				if initPrice == temp.Value3 {
					continue
				}
				return fmt.Errorf("价格配置应该一致")
			}
		}
	}

	return
}

// 投资计划奖励最高等级
func (gt *GroupTemplateInvestLevel) GetInvestLevelMaxRewardsLevel(investType investleveltypes.InvestLevelType) int32 {
	maxLevel := int32(0)
	tempList := gt.levelInvestTempMap[investType]
	for _, temp := range tempList {
		if temp.Value2 < maxLevel {
			continue
		}

		maxLevel = temp.Value2
	}

	return maxLevel
}

//投资计划所需元宝
func (gt *GroupTemplateInvestLevel) GetInvestLevelNeedGold(investType investleveltypes.InvestLevelType) int32 {
	tempList := gt.levelInvestTempMap[investType]
	if len(tempList) > 1 {
		return tempList[0].Value3
	}

	return 0
}

//投资计划-可领取奖励
func (gt *GroupTemplateInvestLevel) GetInvestLevelTempList(investType investleveltypes.InvestLevelType, maxExclude, maxInclude int32) (newTempList []*gametemplate.OpenserverActivityTemplate) {
	tempList := gt.levelInvestTempMap[investType]
	for _, temp := range tempList {
		//领取条件
		if temp.Value2 <= maxExclude || temp.Value2 > maxInclude {
			continue
		}

		newTempList = append(newTempList, temp)
	}

	return
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeLevel, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateInvestLevel))
}
