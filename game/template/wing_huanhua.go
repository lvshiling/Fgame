package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"

	"fmt"
)

//战翼幻化配置
type WingHuanHuaTemplate struct {
	*WingHuanHuaTemplateVO
	useItemMap         map[int32]int32 //进阶物品
	unrealItemTemplate *ItemTemplate
}

func (whht *WingHuanHuaTemplate) TemplateId() int {
	return whht.Id
}

func (whht *WingHuanHuaTemplate) GetUseItemTemplate() map[int32]int32 {
	return whht.useItemMap
}

func (whht *WingHuanHuaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(whht.FileName(), whht.TemplateId(), err)
			return
		}
	}()

	whht.useItemMap = make(map[int32]int32)
	//验证 UseItem
	if whht.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(whht.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", whht.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}

		whht.unrealItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(whht.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", whht.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
		whht.useItemMap[whht.UseItem] = whht.ItemCount
	}

	return nil
}

func (whht *WingHuanHuaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(whht.FileName(), whht.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(whht.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if whht.NextId != 0 {
		diff := whht.NextId - int32(whht.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", whht.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(whht.NextId), (*WingHuanHuaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", whht.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		huanHuaTemplate := to.(*WingHuanHuaTemplate)

		diffLevel := huanHuaTemplate.Level - whht.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", whht.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(whht.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(whht.TimesMin), float64(0), true, float64(whht.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(whht.TimesMax), float64(whht.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(whht.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(whht.AddMin), float64(0), true, float64(whht.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(whht.AddMax), float64(whht.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(whht.ZhufuMax), float64(whht.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(whht.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(whht.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(whht.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", whht.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	if whht.unrealItemTemplate != nil {
		if whht.unrealItemTemplate.GetItemSubType() != itemtypes.ItemWingSubTypeUnreal {
			err = fmt.Errorf("[%d] invalid", whht.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
	}

	return nil
}
func (whht *WingHuanHuaTemplate) PatchAfterCheck() {

}
func (whht *WingHuanHuaTemplate) FileName() string {
	return "tb_wing_huanhua.json"
}

func init() {
	template.Register((*WingHuanHuaTemplate)(nil))
}
