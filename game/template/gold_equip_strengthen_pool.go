package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//装备掉落强化等级配置
type GoldEquipStrengthenPoolTemplate struct {
	*GoldEquipStrengthenPoolTemplateVO
	nextTemplate *GoldEquipStrengthenPoolTemplate
}

func (t *GoldEquipStrengthenPoolTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipStrengthenPoolTemplate) GetNextTemplate() *GoldEquipStrengthenPoolTemplate {
	return t.nextTemplate
}

func (t *GoldEquipStrengthenPoolTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}
func (t *GoldEquipStrengthenPoolTemplate) PatchAfterCheck() {

}
func (t *GoldEquipStrengthenPoolTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Level", err)
	}

	//等级
	err = validator.MinValidate(float64(t.Rate), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Rate", err)
	}

	//下一级
	if t.NextId != 0 {
		diff := t.NextId - t.Id
		if diff != 1 {
			return template.NewTemplateFieldError("NextId", fmt.Errorf("[%d] invalid", t.NextId))
		}
		poolTemp := template.GetTemplateService().Get(t.NextId, (*GoldEquipStrengthenPoolTemplate)(nil))
		if poolTemp == nil {
			return template.NewTemplateFieldError("NextId", fmt.Errorf("[%d] invalid", t.NextId))
		}

		t.nextTemplate = poolTemp.(*GoldEquipStrengthenPoolTemplate)
	}

	return nil
}

func (edt *GoldEquipStrengthenPoolTemplate) FileName() string {
	return "tb_gold_strengthen_pool.json"
}

func init() {
	template.Register((*GoldEquipStrengthenPoolTemplate)(nil))
}
