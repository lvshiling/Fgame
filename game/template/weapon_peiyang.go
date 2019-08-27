package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"

	"fmt"
)

//兵魂培养配置
type WeaponPeiYangTemplate struct {
	*WeaponPeiYangTemplateVO
	useItemMap                map[int32]int32 //进阶物品
	culItemTemplate           *ItemTemplate
	nextWeaponPeiYangTemplate *WeaponPeiYangTemplate
}

func (wpyt *WeaponPeiYangTemplate) TemplateId() int {
	return wpyt.Id
}

func (wpyt *WeaponPeiYangTemplate) GetUseItemTemplate() map[int32]int32 {
	return wpyt.useItemMap
}

func (wpyt *WeaponPeiYangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wpyt.FileName(), wpyt.TemplateId(), err)
			return
		}
	}()

	wpyt.useItemMap = make(map[int32]int32)
	//验证 UseItem
	if wpyt.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(wpyt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", wpyt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		wpyt.culItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(wpyt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", wpyt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
		wpyt.useItemMap[wpyt.UseItem] = wpyt.ItemCount
	}

	return nil
}

func (wpyt *WeaponPeiYangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wpyt.FileName(), wpyt.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(wpyt.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if wpyt.NextId != 0 {
		diff := wpyt.NextId - int32(wpyt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", wpyt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(wpyt.NextId), (*WeaponPeiYangTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wpyt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		peiYangTemplate := to.(*WeaponPeiYangTemplate)

		diffLevel := peiYangTemplate.Level - wpyt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", wpyt.Level)
			return template.NewTemplateFieldError("Level", err)
		}
		wpyt.nextWeaponPeiYangTemplate = peiYangTemplate
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(wpyt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(wpyt.TimesMin), float64(0), true, float64(wpyt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(wpyt.TimesMax), float64(wpyt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(wpyt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(wpyt.AddMin), float64(0), true, float64(wpyt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(wpyt.AddMax), float64(wpyt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(wpyt.ZhufuMax), float64(wpyt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(wpyt.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(wpyt.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(wpyt.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wpyt.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	if wpyt.culItemTemplate != nil {
		if wpyt.culItemTemplate.GetItemSubType() != itemtypes.ItemSoulSubTypeCul {
			err = fmt.Errorf("[%d] invalid", wpyt.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
	}

	return nil
}
func (wpyt *WeaponPeiYangTemplate) PatchAfterCheck() {

}
func (wpyt *WeaponPeiYangTemplate) FileName() string {
	return "tb_weapon_peiyang.json"
}

func init() {
	template.Register((*WeaponPeiYangTemplate)(nil))
}
