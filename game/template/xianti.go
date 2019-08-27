package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	xiantitypes "fgame/fgame/game/xianti/types"
	"fmt"
	"strconv"
)

//仙体配置
type XianTiTemplate struct {
	*XianTiTemplateVO
	typ                     xiantitypes.XianTiType //仙体类型
	magicParamMap           map[int32]string       //幻化条件
	magicParamXUMap         map[int32]int32        //幻化条件1,2
	magicParamIMap          map[int32]int32        //幻化条件3消耗物品
	useItemTemplate         *ItemTemplate          //进阶物品
	battlePropertyMap       map[propertytypes.BattlePropertyType]int64
	xianTiUpstarTemplateMap map[int32]*XianTiUpstarTemplate //仙体皮肤升星map
	xianTiUpstarTemplate    *XianTiUpstarTemplate           //仙体皮肤升星
}

func (mt *XianTiTemplate) TemplateId() int {
	return mt.Id
}

func (mt *XianTiTemplate) GetTyp() xiantitypes.XianTiType {
	return mt.typ
}

func (mt *XianTiTemplate) GetUseItemTemplate() *ItemTemplate {
	return mt.useItemTemplate
}

func (mt *XianTiTemplate) GetMagicParamIMap() map[int32]int32 {
	return mt.magicParamIMap
}

func (mt *XianTiTemplate) GetMagicParamXUMap() map[int32]int32 {
	return mt.magicParamXUMap
}

func (mt *XianTiTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return mt.battlePropertyMap
}

func (mt *XianTiTemplate) GetIsClear() bool {
	return mt.IsClear != 0
}

func (mt *XianTiTemplate) GetXianTiUpstarByLevel(level int32) *XianTiUpstarTemplate {
	if v, ok := mt.xianTiUpstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (mt *XianTiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//仙体类型
	mt.typ = xiantitypes.XianTiType(mt.Type)
	if !mt.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//幻化条件
	mt.magicParamMap = make(map[int32]string)
	mt.magicParamXUMap = make(map[int32]int32)
	mt.magicParamIMap = make(map[int32]int32)

	mt.magicParamMap[mt.MagicConditionType1] = mt.MagicConditionParameter1
	mt.magicParamMap[mt.MagicConditionType2] = mt.MagicConditionParameter2
	mt.magicParamMap[mt.MagicConditionType3] = mt.MagicConditionParameter3

	for condType, condParam := range mt.magicParamMap {
		cType := xiantitypes.XianTiUCondType(condType)
		if !cType.Valid() {
			err = fmt.Errorf("[%d] invalid", condType)
			return template.NewTemplateFieldError("magic_condition_type", err)
		}
		switch cType {
		case xiantitypes.XianTiUCondTypeX,
			xiantitypes.XianTiUCondTypeU,
			xiantitypes.XianTiUCondTypeW:
			num, err := strconv.ParseInt(condParam, 10, 32)
			if err != nil {
				return err
			}
			mt.magicParamXUMap[condType] = int32(num)
			break
		case xiantitypes.XianTiUCondTypeI:
			itemArr, err := utils.SplitAsIntArray(condParam)
			if err != nil {
				return err
			}
			if len(itemArr) != 2 {
				err = fmt.Errorf("[%s] invalid", condParam)
				return template.NewTemplateFieldError("magic_condition_parameter", err)
			}
			mt.magicParamIMap[itemArr[0]] = itemArr[1]
			break
		default:
			break
		}
	}
	//幻化条件1、2
	for condType, condParam := range mt.magicParamXUMap {
		if condType == int32(xiantitypes.XianTiUCondTypeX) {
			err = validator.MinValidate(float64(condParam), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", condParam)
				err = template.NewTemplateFieldError("MagicConditionParameter", err)
				return
			}
		}
	}
	//幻化条件3
	for item, num := range mt.magicParamIMap {
		itemTemplate := template.GetTemplateService().Get(int(item), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("item [%d] invalid", item)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
			return
		}
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("num [%d] invalid", num)
			return template.NewTemplateFieldError("MagicConditionParameter", err)
		}
	}

	//验证 UseItem
	if mt.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(mt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", mt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		mt.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(mt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//验证 xianti_upgrade_begin_id
	if mt.XianTiUpstarBeginId != 0 {
		to := template.GetTemplateService().Get(int(mt.XianTiUpstarBeginId), (*XianTiUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.XianTiUpstarBeginId)
			return template.NewTemplateFieldError("XianTiUpstarBeginId", err)
		}

		xiantiUpstarTemplate, ok := to.(*XianTiUpstarTemplate)
		if !ok {
			return fmt.Errorf("XianTiUpgradeBeginId [%d] invalid", mt.XianTiUpstarBeginId)
		}

		mt.xianTiUpstarTemplate = xiantiUpstarTemplate
	}

	//属性
	mt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(mt.Hp)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(mt.Attack)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(mt.Defence)

	return nil
}

func (mt *XianTiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if mt.NextId != 0 {
		to := template.GetTemplateService().Get(int(mt.NextId), (*XianTiTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp := to.(*XianTiTemplate)

		diff := nextTemp.Number - int32(mt.Number)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	//验证Number
	err = validator.MinValidate(float64(mt.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_money
	err = validator.MinValidate(float64(mt.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mt.TimesMin), float64(0), true, float64(mt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mt.TimesMax), float64(mt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mt.AddMin), float64(0), true, float64(mt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mt.AddMax), float64(mt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(mt.ZhufuMax), float64(mt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 UseYinliang
	err = validator.MinValidate(float64(mt.UseYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UseYinliang)
		err = template.NewTemplateFieldError("UseYinliang", err)
		return
	}

	//验证 ShidanLimit
	err = validator.MinValidate(float64(mt.ShidanLimit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ShidanLimit)
		err = template.NewTemplateFieldError("ShidanLimit", err)
		return
	}

	err = validator.MinValidate(float64(mt.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}
func (mt *XianTiTemplate) PatchAfterCheck() {
	if mt.xianTiUpstarTemplate != nil {
		mt.xianTiUpstarTemplateMap = make(map[int32]*XianTiUpstarTemplate)
		//赋值xianTiUpstarTemplateMap
		for tempTemplate := mt.xianTiUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextXianTiUpstarTemplate {
			level := tempTemplate.Level
			mt.xianTiUpstarTemplateMap[level] = tempTemplate
		}
	}

}
func (mt *XianTiTemplate) FileName() string {
	return "tb_xianti.json"
}

func init() {
	template.Register((*XianTiTemplate)(nil))
}
