package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	wingtypes "fgame/fgame/game/wing/types"
	"fmt"
	"strconv"
)

//战翼配置
type WingTemplate struct {
	*WingTemplateVO
	typ                   wingtypes.WingType            //坐骑类型
	magicParamMap         map[int32]string              //幻化条件
	magicParamXUMap       map[int32]int32               //幻化条件1,2
	magicParamIMap        map[int32]int32               //幻化条件3消耗物品
	useItemTemplate       *ItemTemplate                 //进阶物品
	battleAttrTemplate    *AttrTemplate                 //阶别属性
	wingUpstarTemplateMap map[int32]*WingUpstarTemplate //战翼皮肤升星map
	wingUpstarTemplate    *WingUpstarTemplate           //战翼皮肤升星
}

func (wt *WingTemplate) TemplateId() int {
	return wt.Id
}

func (wt *WingTemplate) GetTyp() wingtypes.WingType {
	return wt.typ
}

func (wt *WingTemplate) GetUseItemTemplate() *ItemTemplate {
	return wt.useItemTemplate
}

func (wt *WingTemplate) GetMagicParamIMap() map[int32]int32 {
	return wt.magicParamIMap
}

func (wt *WingTemplate) GetMagicParamXUMap() map[int32]int32 {
	return wt.magicParamXUMap
}

func (wt *WingTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return wt.battleAttrTemplate
}

func (wt *WingTemplate) GetIsClear() bool {
	return wt.IsClear != 0
}

func (wt *WingTemplate) GetWingUpstarByLevel(level int32) *WingUpstarTemplate {
	if v, ok := wt.wingUpstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (wt *WingTemplate) PatchAfterCheck() {
	if wt.wingUpstarTemplate != nil {
		wt.wingUpstarTemplateMap = make(map[int32]*WingUpstarTemplate)
		//赋值wingUpstarTemplateMap
		for tempTemplate := wt.wingUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextWingUpstarTemplate {
			level := tempTemplate.Level
			wt.wingUpstarTemplateMap[level] = tempTemplate
		}
	}
}

func (wt *WingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wt.FileName(), wt.TemplateId(), err)
			return
		}
	}()

	wt.typ = wingtypes.WingType(wt.Type)
	if !wt.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", wt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//阶别attr属性
	if wt.Attr != 0 {
		to := template.GetTemplateService().Get(int(wt.Attr), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		wt.battleAttrTemplate = attrTemplate
	}

	//幻化条件
	wt.magicParamMap = make(map[int32]string)
	wt.magicParamXUMap = make(map[int32]int32)
	wt.magicParamIMap = make(map[int32]int32)

	wt.magicParamMap[wt.MagicConditionType1] = wt.MagicConditionParameter1
	wt.magicParamMap[wt.MagicConditionType2] = wt.MagicConditionParameter2
	wt.magicParamMap[wt.MagicConditionType3] = wt.MagicConditionParameter3

	for condType, condParam := range wt.magicParamMap {
		cType := wingtypes.WingUCondType(condType)
		if !cType.Valid() {
			err = fmt.Errorf("[%d] invalid", condType)
			return template.NewTemplateFieldError("magic_condition_type", err)
		}
		switch cType {
		case wingtypes.WingUCondTypeX,
			wingtypes.WingUCondTypeU:
			num, err := strconv.ParseInt(condParam, 10, 32)
			if err != nil {
				return err
			}
			wt.magicParamXUMap[condType] = int32(num)
			break
		case wingtypes.WingUCondTypeI:
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
		if condType == int32(wingtypes.WingUCondTypeX) {
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
	if wt.WingUpstarBeginId != 0 {
		to := template.GetTemplateService().Get(int(wt.WingUpstarBeginId), (*WingUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.WingUpstarBeginId)
			return template.NewTemplateFieldError("WingUpstarBeginId", err)
		}

		wingUpstarTemplate, ok := to.(*WingUpstarTemplate)
		if !ok {
			return fmt.Errorf("WingUpgradeBeginId [%d] invalid", wt.WingUpstarBeginId)
		}

		wt.wingUpstarTemplate = wingUpstarTemplate
	}

	return nil
}

func (wt *WingTemplate) Check() (err error) {
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
		to := template.GetTemplateService().Get(int(wt.NextId), (*WingTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp := to.(*WingTemplate)

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

func (wt *WingTemplate) FileName() string {
	return "tb_wing.json"
}

func init() {
	template.Register((*WingTemplate)(nil))
}
