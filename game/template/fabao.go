package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	fabaotypes "fgame/fgame/game/fabao/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
	"strconv"
)

//法宝配置
type FaBaoTemplate struct {
	*FaBaoTemplateVO
	typ                     fabaotypes.FaBaoType           //法宝类型
	magicParamMap           map[int32]string               //幻化条件
	magicParamXUMap         map[int32]int32                //幻化条件1,2
	magicParamIMap          map[int32]int32                //幻化条件3消耗物品
	useItemTemplate         *ItemTemplate                  //进阶物品
	faBaoUpstartTemplateMap map[int32]*FaBaoUpstarTemplate //法宝皮肤升星map
	faBaoUpstartTemplate    *FaBaoUpstarTemplate           //法宝皮肤升星
	battlePropertyMap       map[propertytypes.BattlePropertyType]int64
}

func (wt *FaBaoTemplate) TemplateId() int {
	return wt.Id
}

func (wt *FaBaoTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return wt.battlePropertyMap
}

func (wt *FaBaoTemplate) GetTyp() fabaotypes.FaBaoType {
	return wt.typ
}

func (wt *FaBaoTemplate) GetUseItemTemplate() *ItemTemplate {
	return wt.useItemTemplate
}

func (wt *FaBaoTemplate) GetMagicParamIMap() map[int32]int32 {
	return wt.magicParamIMap
}

func (wt *FaBaoTemplate) GetMagicParamXUMap() map[int32]int32 {
	return wt.magicParamXUMap
}

func (wt *FaBaoTemplate) GetIsClear() bool {
	return wt.IsClear != 0
}

func (wt *FaBaoTemplate) GetFaBaoUpstarByLevel(level int32) *FaBaoUpstarTemplate {
	if v, ok := wt.faBaoUpstartTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (wt *FaBaoTemplate) PatchAfterCheck() {
	if wt.faBaoUpstartTemplate != nil {
		wt.faBaoUpstartTemplateMap = make(map[int32]*FaBaoUpstarTemplate)
		//赋值faBaoUpstartTemplateMap
		for tempTemplate := wt.faBaoUpstartTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextFaBaoUpstarTemplate {
			level := tempTemplate.Level
			wt.faBaoUpstartTemplateMap[level] = tempTemplate
		}
	}
}

func (wt *FaBaoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wt.FileName(), wt.TemplateId(), err)
			return
		}
	}()

	wt.typ = fabaotypes.FaBaoType(wt.Type)
	if !wt.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", wt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//幻化条件
	wt.magicParamMap = make(map[int32]string)
	wt.magicParamXUMap = make(map[int32]int32)
	wt.magicParamIMap = make(map[int32]int32)

	wt.magicParamMap[wt.MagicConditionType1] = wt.MagicConditionParameter1
	wt.magicParamMap[wt.MagicConditionType2] = wt.MagicConditionParameter2
	wt.magicParamMap[wt.MagicConditionType3] = wt.MagicConditionParameter3

	for condType, condParam := range wt.magicParamMap {
		cType := fabaotypes.FaBaoUCondType(condType)
		if !cType.Valid() {
			err = fmt.Errorf("[%d] invalid", condType)
			return template.NewTemplateFieldError("magic_condition_type", err)
		}
		switch cType {
		case fabaotypes.FaBaoUCondTypeX,
			fabaotypes.FaBaoUCondTypeU,
			fabaotypes.FaBaoUCondTypeW:
			num, err := strconv.ParseInt(condParam, 10, 32)
			if err != nil {
				return err
			}
			wt.magicParamXUMap[condType] = int32(num)
			break
		case fabaotypes.FaBaoUCondTypeI:
			itemArr, err := utils.SplitAsIntArray(condParam)
			if err != nil {
				return err
			}
			if len(itemArr) != 2 {
				err = fmt.Errorf("[%s] invalid", condParam)
				return template.NewTemplateFieldError("magic_condition_parameter", err)
			}
			wt.magicParamIMap[itemArr[0]] = itemArr[1]
			break
		default:
			break
		}
	}
	//幻化条件1、2
	for condType, condParam := range wt.magicParamXUMap {
		if condType == int32(fabaotypes.FaBaoUCondTypeX) {
			err = validator.MinValidate(float64(condParam), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", condParam)
				err = template.NewTemplateFieldError("MagicConditionParameter", err)
				return
			}
		}
	}
	//幻化条件3
	for item, num := range wt.magicParamIMap {
		itemTemplate := template.GetTemplateService().Get(int(item), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("item[%d] invalid", item)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
			return
		}
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("num[%d] invalid", num)
			return template.NewTemplateFieldError("MagicConditionParameter", err)
		}
	}

	//验证 UseItem
	if wt.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(wt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", wt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		wt.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(wt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", wt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//验证 wing_upgrade_begin_id
	if wt.FaBaoUpstarBeginId != 0 {
		to := template.GetTemplateService().Get(int(wt.FaBaoUpstarBeginId), (*FaBaoUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.FaBaoUpstarBeginId)
			return template.NewTemplateFieldError("FaBaoUpstarBeginId", err)
		}

		faBaoUpstartTemplate, ok := to.(*FaBaoUpstarTemplate)
		if !ok {
			return fmt.Errorf("FaBaoUpgradeBeginId [%d] invalid", wt.FaBaoUpstarBeginId)
		}

		wt.faBaoUpstartTemplate = faBaoUpstartTemplate
	}

	//属性
	wt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	wt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(wt.Hp)
	wt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(wt.Attack)
	wt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(wt.Defence)

	return nil
}

func (wt *FaBaoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wt.FileName(), wt.TemplateId(), err)
			return
		}
	}()

	//验证Number
	err = validator.MinValidate(float64(wt.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证 next_id
	if wt.NextId != 0 {

		to := template.GetTemplateService().Get(int(wt.NextId), (*FaBaoTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp := to.(*FaBaoTemplate)

		diff := nextTemp.Number - int32(wt.Number)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", wt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	//验证 update_wfb
	err = validator.RangeValidate(float64(wt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 use_money
	err = validator.MinValidate(float64(wt.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(wt.TimesMin), float64(0), true, float64(wt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(wt.TimesMax), float64(wt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(wt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(wt.AddMin), float64(0), true, float64(wt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(wt.AddMax), float64(wt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(wt.ZhufuMax), float64(wt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 UseYinliang
	err = validator.MinValidate(float64(wt.UseYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.UseYinliang)
		err = template.NewTemplateFieldError("UseYinliang", err)
		return
	}

	//验证 ShidanLimit
	err = validator.MinValidate(float64(wt.ShidanLimit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.ShidanLimit)
		err = template.NewTemplateFieldError("ShidanLimit", err)
		return
	}

	err = validator.MinValidate(float64(wt.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}

func (wt *FaBaoTemplate) FileName() string {
	return "tb_fabao.json"
}

func init() {
	template.Register((*FaBaoTemplate)(nil))
}
