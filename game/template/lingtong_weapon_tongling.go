package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/lingtongdev/types"
	"fmt"
)

//灵兵通灵配置
type LingTongWeaponTongLingTemplate struct {
	*LingTongWeaponTongLingTemplateVO
	needItemMap     map[int32]int32 //通灵需要物品
	useItemTemplate *ItemTemplate
	next            LingTongDevTongLingTemplate
}

func (t *LingTongWeaponTongLingTemplate) TemplateId() int {
	return t.Id
}

func (t *LingTongWeaponTongLingTemplate) GetNextId() int32 {
	return t.NextId
}

func (t *LingTongWeaponTongLingTemplate) GetUpdateWfb() int32 {
	return t.UpdateWfb
}

func (t *LingTongWeaponTongLingTemplate) GetAddMin() int32 {
	return t.AddMin
}

func (t *LingTongWeaponTongLingTemplate) GetAddMax() int32 {
	return t.AddMax
}

func (t *LingTongWeaponTongLingTemplate) GetTimesMin() int32 {
	return t.TimesMin
}

func (t *LingTongWeaponTongLingTemplate) GetTimesMax() int32 {
	return t.TimesMax
}

func (t *LingTongWeaponTongLingTemplate) GetZhuFuMax() int32 {
	return t.ZhufuMax
}

func (t *LingTongWeaponTongLingTemplate) GetItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *LingTongWeaponTongLingTemplate) GetTongLingPercent() int32 {
	return t.LingTongWeaponPercent
}

func (t *LingTongWeaponTongLingTemplate) GetLevel() int32 {
	return t.Level
}

func (t *LingTongWeaponTongLingTemplate) GetNext() LingTongDevTongLingTemplate {
	return t.next
}

func (t *LingTongWeaponTongLingTemplate) GetItemCount() int32 {
	return t.ItemCount
}

func (t *LingTongWeaponTongLingTemplate) GetClassType() types.LingTongDevSysType {
	return types.LingTongDevSysTypeLingBao
}

func (t *LingTongWeaponTongLingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.needItemMap = make(map[int32]int32)
	//验证 use_item
	if t.UseItem != 0 {
		to := template.GetTemplateService().Get(int(t.UseItem), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}

		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}

		t.needItemMap[t.UseItem] = t.ItemCount
	}

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*LingTongWeaponTongLingTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*LingTongWeaponTongLingTemplate)
			diffLevel := nextTemplate.Level - t.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			t.next = nextTemplate
		}
	}

	return nil
}

func (t *LingTongWeaponTongLingTemplate) PatchAfterCheck() {

}

func (t *LingTongWeaponTongLingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 fabao_percent
	err = validator.RangeValidate(float64(t.LingTongWeaponPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongWeaponPercent)
		err = template.NewTemplateFieldError("LingTongWeaponPercent", err)
		return
	}

	//验证 level
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	for itemId, _ := range t.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
		itemTemplate := to.(*ItemTemplate)

		if itemTemplate.GetItemSubType() != itemtypes.ItemLingTongWeaponSubTypeTongLingDan {
			err = fmt.Errorf("UpstarItemId [%d]  invalid", t.UseItem)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(t.TimesMin), float64(0), true, float64(t.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(t.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(t.AddMin), float64(0), true, float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(t.AddMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return nil
}

func (t *LingTongWeaponTongLingTemplate) FileName() string {
	return "tb_lingtong_weapon_tongling.json"
}

// 使用通用 tb_system_tongling
// func init() {
// 	template.Register((*LingTongWeaponTongLingTemplate)(nil))
// }
