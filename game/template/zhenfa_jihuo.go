package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	zhenfatypes "fgame/fgame/game/zhenfa/types"

	"fmt"
)

//阵法激活配置
type ZhenFaJiHuoTemplate struct {
	*ZhenFaJiHuoTemplateVO
	typ               zhenfatypes.ZhenFaType //阵法类型
	itemMap           map[int32]int32
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	nextTemplate      *ZhenFaJiHuoTemplate
}

func (mt *ZhenFaJiHuoTemplate) TemplateId() int {
	return mt.Id
}

func (mt *ZhenFaJiHuoTemplate) GetZhenFaType() zhenfatypes.ZhenFaType {
	return mt.typ
}

func (mt *ZhenFaJiHuoTemplate) GetNeedItemMap() map[int32]int32 {
	return mt.itemMap
}

func (mt *ZhenFaJiHuoTemplate) GetNextTemplate() *ZhenFaJiHuoTemplate {
	return mt.nextTemplate
}

func (mt *ZhenFaJiHuoTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return mt.battlePropertyMap
}

func (mt *ZhenFaJiHuoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//阵法类型
	mt.typ = zhenfatypes.ZhenFaType(mt.Type)
	if !mt.typ.Vaild() {
		err = fmt.Errorf("[%d] invalid", mt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//验证 UseItem
	if mt.UseItem != 0 {
		mt.itemMap = make(map[int32]int32)
		useItemTemplateVO := template.GetTemplateService().Get(int(mt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", mt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}

		//验证 ItemCount
		err = validator.MinValidate(float64(mt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
		mt.itemMap[mt.UseItem] = mt.ItemCount
	}

	//属性
	mt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	if mt.Hp != 0 {
		mt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = mt.Hp
	}
	if mt.Attack != 0 {
		mt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = mt.Attack
	}
	if mt.Defence != 0 {
		mt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = mt.Defence
	}

	return nil
}

func (mt *ZhenFaJiHuoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证update_wfb
	err = validator.RangeValidate(float64(mt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
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
	err = validator.MinValidate(float64(mt.ZhuFuMax), float64(mt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ZhuFuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 need_zhenfa_type
	if mt.NeedZhenFaType != 0 {
		if !zhenfatypes.ZhenFaType(mt.NeedZhenFaType).Vaild() {
			err = fmt.Errorf("[%d] invalid", mt.NeedZhenFaType)
			err = template.NewTemplateFieldError("NeedZhenFaType", err)
			return
		}
	}

	err = validator.MinValidate(float64(mt.NeedZhenFaLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedZhenFaLevel)
		return template.NewTemplateFieldError("NeedZhenFaLevel", err)
	}

	return nil
}
func (mt *ZhenFaJiHuoTemplate) PatchAfterCheck() {

}
func (mt *ZhenFaJiHuoTemplate) FileName() string {
	return "tb_zhenfa_jihuo.json"
}

func init() {
	template.Register((*ZhenFaJiHuoTemplate)(nil))
}
