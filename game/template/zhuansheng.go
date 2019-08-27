package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//转生配置
type ZhuanShengTemplate struct {
	*ZhuanShengTemplateVO
}

func (t *ZhuanShengTemplate) TemplateId() int {
	return t.Id
}

func (t *ZhuanShengTemplate) PatchAfterCheck() {
}

func (t *ZhuanShengTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ZhuanShengTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 need_zhuanshu
	err = validator.MinValidate(float64(t.NeedZhuanshu), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedZhuanshu)
		return template.NewTemplateFieldError("NeedZhuanshu", err)
	}

	//验证 need_level
	err = validator.MinValidate(float64(t.NeedLevel), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedLevel)
		return template.NewTemplateFieldError("NeedLevel", err)
	}

	//验证 NeedFeisheng
	err = validator.MinValidate(float64(t.NeedFeisheng), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedFeisheng)
		return template.NewTemplateFieldError("NeedFeisheng", err)
	}

	//验证 need_equip_count
	err = validator.MinValidate(float64(t.NeedEquipCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedEquipCount)
		return template.NewTemplateFieldError("NeedEquipCount", err)
	}

	//need_equip_zhuanshu
	err = validator.MinValidate(float64(t.NeedEquipZhuanshu), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedEquipZhuanshu)
		return template.NewTemplateFieldError("NeedEquipZhuanshu", err)
	}

	//need_equip_level
	err = validator.MinValidate(float64(t.NeedEquipLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedEquipLevel)
		return template.NewTemplateFieldError("NeedEquipLevel", err)
	}

	//need_equip_streng
	err = validator.MinValidate(float64(t.NeedEquipStreng), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedEquipStreng)
		return template.NewTemplateFieldError("NeedEquipStreng", err)
	}

	//need_equip_streng
	err = validator.MinValidate(float64(t.NeedEquipQuality), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedEquipQuality)
		return template.NewTemplateFieldError("NeedEquipQuality", err)
	}

	//hp
	err = validator.MinValidate(float64(t.Hp), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//attack
	err = validator.MinValidate(float64(t.Attack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//defence
	err = validator.MinValidate(float64(t.Defence), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}
	return nil
}

func (t *ZhuanShengTemplate) FileName() string {
	return "tb_zhuansheng.json"
}

func init() {
	template.Register((*ZhuanShengTemplate)(nil))
}
