package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"

	"fmt"
)

//坐骑幻化配置
type MountHuanHuaTemplate struct {
	*MountHuanHuaTemplateVO
	useItemMap         map[int32]int32 //进阶物品
	unrealItemTemplate *ItemTemplate
}

func (mhht *MountHuanHuaTemplate) TemplateId() int {
	return mhht.Id
}

func (mhht *MountHuanHuaTemplate) GetUseItemTemplate() map[int32]int32 {
	return mhht.useItemMap
}

func (mhht *MountHuanHuaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mhht.FileName(), mhht.TemplateId(), err)
			return
		}
	}()

	mhht.useItemMap = make(map[int32]int32)
	//验证 UseItem
	if mhht.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(mhht.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", mhht.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		mhht.unrealItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(mhht.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mhht.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
		mhht.useItemMap[mhht.UseItem] = mhht.ItemCount
	}

	return nil
}

func (mhht *MountHuanHuaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mhht.FileName(), mhht.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(mhht.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if mhht.NextId != 0 {
		diff := mhht.NextId - int32(mhht.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mhht.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(mhht.NextId), (*MountHuanHuaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mhht.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		huanHuaTemplate := to.(*MountHuanHuaTemplate)

		diffLevel := huanHuaTemplate.Level - mhht.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", mhht.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mhht.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mhht.TimesMin), float64(0), true, float64(mhht.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mhht.TimesMax), float64(mhht.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mhht.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mhht.AddMin), float64(0), true, float64(mhht.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mhht.AddMax), float64(mhht.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(mhht.ZhufuMax), float64(mhht.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(mhht.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(mhht.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(mhht.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mhht.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	if mhht.unrealItemTemplate != nil {
		if mhht.unrealItemTemplate.GetItemSubType() != itemtypes.ItemMountSubTypeUnreal {
			err = fmt.Errorf("[%d] invalid", mhht.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
	}

	return nil
}
func (mhht *MountHuanHuaTemplate) PatchAfterCheck() {

}
func (mhht *MountHuanHuaTemplate) FileName() string {
	return "tb_mount_huanhua.json"
}

func init() {
	template.Register((*MountHuanHuaTemplate)(nil))
}
