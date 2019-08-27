package template

import (
	"fgame/fgame/core/template"
	itemskilltypes "fgame/fgame/game/itemskill/types"
	"fmt"
)

//系统技能配置
type ItemSkillTemplate struct {
	*ItemSkillTemplateVO
	typ itemskilltypes.ItemSkillType //类型
}

func (t *ItemSkillTemplate) TemplateId() int {
	return t.Id
}

func (t *ItemSkillTemplate) GetType() itemskilltypes.ItemSkillType {
	return t.typ
}

func (t *ItemSkillTemplate) PatchAfterCheck() {
}

func (t *ItemSkillTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 type
	t.typ = itemskilltypes.ItemSkillType(t.Type)
	if !t.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	return nil
}

func (t *ItemSkillTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*ItemSkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		nextTo := to.(*ItemSkillTemplate)

		//验证type
		if t.Type != nextTo.Type {
			err = fmt.Errorf("[%d] invalid", nextTo.Type)
			err = template.NewTemplateFieldError("Type", err)
			return
		}

		//验证level
		diffLevel := nextTo.Level - t.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", nextTo.Level)
			err = template.NewTemplateFieldError("Level", err)
			return
		}
	}

	//验证 skill_id
	to := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}

	return nil
}

func (t *ItemSkillTemplate) FileName() string {
	return "tb_jinganghuti.json"
}

func init() {
	template.Register((*ItemSkillTemplate)(nil))
}
