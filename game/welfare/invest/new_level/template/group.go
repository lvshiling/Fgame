package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	investnewleveltypes "fgame/fgame/game/welfare/invest/new_level/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"sort"
)

//新等级投资计划
type GroupTemplateInvestNewLevel struct {
	*welfaretemplate.GroupTemplateBase
	levelInvestTempMap map[investnewleveltypes.InvestNewLevelType]map[int32]*gametemplate.OpenserverActivityTemplate
}

func CreateGroupTemplateInvestNewLevel(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateInvestNewLevel{}
	gt.GroupTemplateBase = base
	return gt
}

func (gt *GroupTemplateInvestNewLevel) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	gt.levelInvestTempMap = make(map[investnewleveltypes.InvestNewLevelType]map[int32]*gametemplate.OpenserverActivityTemplate)
	var openTempList []*gametemplate.OpenserverActivityTemplate
	for _, t := range gt.GetOpenTempMap() {
		//投资计划类型
		investType := investnewleveltypes.InvestNewLevelType(t.Value1)
		if !investType.Valid() {
			err = fmt.Errorf("[%d] Value1 invalid", t.Value1)
			err = welfaretypes.NewWelfareRecordError(t.Id, err)
			return
		}

		//等级
		err = validator.MinValidate(float64(t.Value2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value2)
			err = template.NewTemplateFieldError("Value2", err)
			return
		}

		//价格
		err = validator.MinValidate(float64(t.Value3), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value3)
			err = template.NewTemplateFieldError("Value3", err)
			return
		}

		tempM, ok := gt.levelInvestTempMap[investType]
		if !ok {
			tempM = make(map[int32]*gametemplate.OpenserverActivityTemplate)
			gt.levelInvestTempMap[investType] = tempM
		}

		//是否等级重复
		_, isRep := tempM[t.Value2]
		if isRep {
			err = fmt.Errorf("[%d] invalid", t.Value2)
			err = template.NewTemplateFieldError("Value2 repeat", err)
			return
		}
		tempM[t.Value2] = t
		openTempList = append(openTempList, t)
	}

	// 升序
	sort.Sort(welfaretemplate.SortTempListThree(openTempList))

	// 验证value3,保证高级投资投入资金比低级多
	var lastOpenTemp *gametemplate.OpenserverActivityTemplate
	index := 0
	for _, openTemp := range openTempList {
		if index != 0 {
			// 投入资金比较
			if lastOpenTemp.Value3 == openTemp.Value3 {
				// 投资等级类型比较
				if openTemp.Value1 != lastOpenTemp.Value1 {
					return fmt.Errorf("高级投资投入资金应该比低级多")
				}
				continue
			}
			// 投资等级类型比较
			if openTemp.Value1 <= lastOpenTemp.Value1 {
				return fmt.Errorf("高级投资投入资金应该比低级多")
			}
		}
		lastOpenTemp = openTemp
		index++
	}

	// 校验 购买价格
	for _, tempM := range gt.levelInvestTempMap {
		if len(tempM) < 1 {
			continue
		}
		initPrice := int32(0)
		for _, temp := range tempM {
			if initPrice == 0 {
				initPrice = temp.Value3
				continue
			} else if initPrice == temp.Value3 {
				continue
			}
			return fmt.Errorf("价格配置应该一致")
		}
	}

	return
}

//投资计划所需元宝
func (gt *GroupTemplateInvestNewLevel) GetInvestLevelNeedGold(investType investnewleveltypes.InvestNewLevelType) int32 {
	tempM := gt.levelInvestTempMap[investType]
	for _, temp := range tempM {
		return temp.Value3
	}
	return 0
}

//投资计划-可领取奖励
func (gt *GroupTemplateInvestNewLevel) GetInvestLevelRewTempByArg(investType investnewleveltypes.InvestNewLevelType, level int32) *gametemplate.OpenserverActivityTemplate {
	temp, ok := gt.levelInvestTempMap[investType][level]
	if !ok {
		return nil
	}
	return temp
}

//投资计划-可领取奖励
func (gt *GroupTemplateInvestNewLevel) GetInvestLevelTempMByType(investType investnewleveltypes.InvestNewLevelType) map[int32]*gametemplate.OpenserverActivityTemplate {
	tempM, ok := gt.levelInvestTempMap[investType]
	if !ok {
		return nil
	}
	return tempM
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewLevel, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateInvestNewLevel))
}
