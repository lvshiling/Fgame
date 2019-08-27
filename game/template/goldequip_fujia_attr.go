package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//金装附加属性池配置
type GoldEquipFuJiaAttrTemplate struct {
	*GoldEquipFuJiaAttrTemplateVO
	nextTemp *GoldEquipFuJiaAttrTemplate
}

func (t *GoldEquipFuJiaAttrTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipFuJiaAttrTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *GoldEquipFuJiaAttrTemplate) PatchAfterCheck() {
}

func (t *GoldEquipFuJiaAttrTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证：下一级ID
	if t.NextId != 0 {
		diff := t.NextId - t.Id
		if diff != 1 {
			return template.NewTemplateFieldError("NextId", fmt.Errorf("[%d] invalid", t.NextId))
		}
		tempObj := template.GetTemplateService().Get(t.NextId, (*GoldEquipFuJiaAttrTemplate)(nil))
		if tempObj == nil {
			return template.NewTemplateFieldError("NextId", fmt.Errorf("[%d] invalid", t.NextId))
		}
		poolTemp := tempObj.(*GoldEquipFuJiaAttrTemplate)
		t.nextTemp = poolTemp
	}

	//生命
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}
	//攻击
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}
	//防御
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}
	//装备属性万分比
	err = validator.MinValidate(float64(t.EquipPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.EquipPercent)
		return template.NewTemplateFieldError("EquipPercent", err)
	}
	//回血属性
	err = validator.MinValidate(float64(t.Huixie), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Huixie)
		return template.NewTemplateFieldError("Huixie", err)
	}
	//增伤
	err = validator.MinValidate(float64(t.BossZengshang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BossZengshang)
		return template.NewTemplateFieldError("BossZengshang", err)
	}
	//减伤
	err = validator.MinValidate(float64(t.BossJianshang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BossJianshang)
		return template.NewTemplateFieldError("BossJianshang", err)
	}
	//击杀buff
	err = validator.MinValidate(float64(t.KillMonsterBuff), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KillMonsterBuff)
		return template.NewTemplateFieldError("KillMonsterBuff", err)
	}

	return nil
}

func (edt *GoldEquipFuJiaAttrTemplate) FileName() string {
	return "tb_goldequip_fujia_attr.json"
}

func init() {
	template.Register((*GoldEquipFuJiaAttrTemplate)(nil))
}
