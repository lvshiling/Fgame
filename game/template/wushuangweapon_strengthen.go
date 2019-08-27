package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

type WushuangWeaponStrengthenTemplate struct {
	*WushuangWeaponStrengthenTemplateVO
	nextStrengthenTemplate *WushuangWeaponStrengthenTemplate
	totalNeedExperience    int64
	minQuality             itemtypes.ItemQualityType
}

func (t *WushuangWeaponStrengthenTemplate) TemplateId() int {
	return t.Id
}

func (t *WushuangWeaponStrengthenTemplate) SetTotalNeedExperience(exp int64) {
	t.totalNeedExperience = exp
}

func (t *WushuangWeaponStrengthenTemplate) GetAllNeedExperience() int64 {
	return t.totalNeedExperience
}

func (t *WushuangWeaponStrengthenTemplate) GetMinNeedQuality() itemtypes.ItemQualityType {
	return t.minQuality
}

//检查有效性
func (t *WushuangWeaponStrengthenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//下一等级
	err = validator.MinValidate(float64(t.NextId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NextId)
		return template.NewTemplateFieldError("NextId", err)
	}

	//Hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//攻击力
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//防御力
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//升级所需经验
	err = validator.MinValidate(float64(t.Experience), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Experience)
		return template.NewTemplateFieldError("Experience", err)
	}
	if t.Level == 0 {
		if t.Experience != 0 {
			err = fmt.Errorf("[%d] invalid", t.Experience)
			return template.NewTemplateFieldError("Experience", err)
		}
	}

	//突破所需转数
	err = validator.MinValidate(float64(t.TupoZhuanshu), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TupoZhuanshu)
		return template.NewTemplateFieldError("TupoZhuanshu", err)
	}

	//突破成功率
	err = validator.RangeValidate(float64(t.TupoRate), float64(0), true, float64(10000), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TupoRate)
		return template.NewTemplateFieldError("TupoRate", err)
	}

	//是否修改
	err = validator.RangeValidate(float64(t.IsTupo), float64(0), true, float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsTupo)
		return template.NewTemplateFieldError("IsTupo", err)
	}

	//需要的数量
	err = validator.MinValidate(float64(t.TupoNeedCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TupoNeedCount)
		return template.NewTemplateFieldError("TupoNeedCount", err)
	}

	//品质
	t.minQuality = itemtypes.ItemQualityType(t.TupoNeedQuality)
	if !t.minQuality.Valid() {
		err = fmt.Errorf("[%d] invalid", t.TupoNeedQuality)
		return template.NewTemplateFieldError("TupoNeedQuality", err)
	}

	if t.nextStrengthenTemplate != nil {
		diff := t.nextStrengthenTemplate.Level - t.Level
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	return
}

func (t *WushuangWeaponStrengthenTemplate) IsMaxLevel() bool {
	if t.NextId == 0 {
		return true
	} else {
		return false
	}
}

func (t *WushuangWeaponStrengthenTemplate) IsNeedBreakThrough() bool {
	//TODO:wayne 修改后的
	if t.IsTupo != 0 {
		return true
	}
	return false
	// if t.IsTupo == 0 {
	// 	return false
	// } else if t.IsTupo == 1 {
	// 	return true
	// }
	// return true
}

func (t *WushuangWeaponStrengthenTemplate) GetNextStrengthenTemplate() *WushuangWeaponStrengthenTemplate {
	return t.nextStrengthenTemplate
}

//组合成需要的数据
func (t *WushuangWeaponStrengthenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId == 0 {
		t.nextStrengthenTemplate = nil
	} else {
		temp := template.GetTemplateService().Get(int(t.NextId), (*WushuangWeaponStrengthenTemplate)(nil))
		t.nextStrengthenTemplate, _ = temp.(*WushuangWeaponStrengthenTemplate)
		if t.nextStrengthenTemplate == nil {
			err = fmt.Errorf("WushuangStrengthenTemplate[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
		}
	}

	return
}

//检验后组合
func (t *WushuangWeaponStrengthenTemplate) PatchAfterCheck() {
}

func (t *WushuangWeaponStrengthenTemplate) FileName() string {
	return "tb_wushuang_strengthen.json"
}

func init() {
	template.Register((*WushuangWeaponStrengthenTemplate)(nil))
}
