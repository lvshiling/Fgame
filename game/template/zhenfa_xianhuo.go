package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	zhenfatypes "fgame/fgame/game/zhenfa/types"

	"fmt"
)

//阵旗仙火配置
type ZhenFaXianHuoTemplate struct {
	*ZhenFaXianHuoTemplateVO
	typ                 zhenfatypes.ZhenFaType //阵法类型
	itemMap             map[int32]int32
	allItemMap          map[int32]int32 //含保护符物品
	battlePropertyMap   map[propertytypes.BattlePropertyType]int64
	nextTemplate        *ZhenFaXianHuoTemplate
	returnLevelTemplate *ZhenFaXianHuoTemplate
}

func (mt *ZhenFaXianHuoTemplate) TemplateId() int {
	return mt.Id
}

func (mt *ZhenFaXianHuoTemplate) GetZhenFaType() zhenfatypes.ZhenFaType {
	return mt.typ
}

func (mt *ZhenFaXianHuoTemplate) GetNeedItemMap() map[int32]int32 {
	return mt.itemMap
}

func (mt *ZhenFaXianHuoTemplate) GetAllItemMap() map[int32]int32 {
	return mt.allItemMap
}

func (mt *ZhenFaXianHuoTemplate) GetNextTemplate() *ZhenFaXianHuoTemplate {
	return mt.nextTemplate
}

func (mt *ZhenFaXianHuoTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return mt.battlePropertyMap
}

func (mt *ZhenFaXianHuoTemplate) GetReturnLevelTemplate() *ZhenFaXianHuoTemplate {
	return mt.returnLevelTemplate
}

func (mt *ZhenFaXianHuoTemplate) Patch() (err error) {
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
	mt.allItemMap = make(map[int32]int32)
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
		mt.allItemMap[mt.UseItem] += mt.ItemCount
	}

	if mt.BaoHuFuId != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(mt.BaoHuFuId), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", mt.BaoHuFuId)
			err = template.NewTemplateFieldError("BaoHuFuId", err)
			return
		}

		err = validator.MinValidate(float64(mt.BaoHuFuCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mt.BaoHuFuCount)
			err = template.NewTemplateFieldError("BaoHuFuCount", err)
			return
		}
		mt.allItemMap[mt.BaoHuFuId] += mt.BaoHuFuCount
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

func (mt *ZhenFaXianHuoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if mt.NextId != 0 {
		to := template.GetTemplateService().Get(int(mt.NextId), (*ZhenFaXianHuoTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		mt.nextTemplate = to.(*ZhenFaXianHuoTemplate)
		if mt.nextTemplate.GetZhenFaType() != mt.nextTemplate.GetZhenFaType() {
			err = fmt.Errorf("[%d] invalid", mt.Type)
			return template.NewTemplateFieldError("Type", err)
		}

		if mt.nextTemplate.Level-mt.Level != 1 {
			err = fmt.Errorf("[%d] invalid", mt.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证level
	err = validator.MinValidate(float64(mt.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

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

	if mt.ReturnLevelId != 0 {
		to := template.GetTemplateService().Get(int(mt.ReturnLevelId), (*ZhenFaXianHuoTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.ReturnLevelId)
			return template.NewTemplateFieldError("ReturnLevelId", err)
		}
		mt.returnLevelTemplate = to.(*ZhenFaXianHuoTemplate)
	}

	//验证 return_rate
	err = validator.RangeValidate(float64(mt.ReturnRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ReturnRate)
		err = template.NewTemplateFieldError("ReturnRate", err)
		return
	}

	//验证 percent
	err = validator.RangeValidate(float64(mt.Percent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Percent)
		err = template.NewTemplateFieldError("Percent", err)
		return
	}

	return nil
}
func (mt *ZhenFaXianHuoTemplate) PatchAfterCheck() {

}
func (mt *ZhenFaXianHuoTemplate) FileName() string {
	return "tb_zhenfa_xianhuo.json"
}

func init() {
	template.Register((*ZhenFaXianHuoTemplate)(nil))
}
