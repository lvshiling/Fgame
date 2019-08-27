package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//表白系统配置
type MarryDevelopTemplate struct {
	*MarryDevelopTemplateVO
	nextMarryDevelopTemplate *MarryDevelopTemplate                      //下一级
	battlePropertyMap        map[propertytypes.BattlePropertyType]int64 //属性
}

func (t *MarryDevelopTemplate) TemplateId() int {
	return t.Id
}

func (t *MarryDevelopTemplate) GetLevel() int32 {
	return t.Level
}

func (t *MarryDevelopTemplate) GetHp() int32 {
	return t.AddHp
}

func (t *MarryDevelopTemplate) GetAttack() int32 {
	return t.AddAttack
}

func (t *MarryDevelopTemplate) GetDefence() int32 {
	return t.AddDefence
}

func (t *MarryDevelopTemplate) GetNextTemplate() *MarryDevelopTemplate {
	return t.nextMarryDevelopTemplate
}

func (t *MarryDevelopTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *MarryDevelopTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.AddHp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.AddAttack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.AddDefence)

	return nil
}

func (t *MarryDevelopTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}
	//验证等级
	err = validator.MinValidate(float64(t.Experience), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Experience)
		err = template.NewTemplateFieldError("Experience", err)
		return
	}

	//验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*MarryDevelopTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextMarryDevelopTemplate = to.(*MarryDevelopTemplate)

		diffLevel := t.nextMarryDevelopTemplate.Level - t.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", t.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证 hp
	err = validator.MinValidate(float64(t.AddHp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddHp)
		err = template.NewTemplateFieldError("AddHp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(t.AddAttack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddAttack)
		err = template.NewTemplateFieldError("AddAttack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(t.AddDefence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddDefence)
		err = template.NewTemplateFieldError("AddDefence", err)
		return
	}

	return nil
}
func (t *MarryDevelopTemplate) PatchAfterCheck() {

}
func (t *MarryDevelopTemplate) FileName() string {
	return "tb_biaobai.json"
}

func init() {
	template.Register((*MarryDevelopTemplate)(nil))
}
