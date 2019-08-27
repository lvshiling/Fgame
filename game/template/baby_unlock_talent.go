package template

import (
	"fgame/fgame/core/template"
	"fmt"
)

//宝宝解锁天赋配置
type BabyUnlockTalentTemplate struct {
	*BabyUnlockTalentTemplateVO
	nextTemp *BabyUnlockTalentTemplate
}

func (t *BabyUnlockTalentTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyUnlockTalentTemplate) PatchAfterCheck() {}

func (t *BabyUnlockTalentTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//下一阶强化
	if t.NextId != 0 {
		if t.NextId-t.Id != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}

		to := template.GetTemplateService().Get(t.NextId, (*BabyUnlockTalentTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}
		t.nextTemp = to.(*BabyUnlockTalentTemplate)

		if t.nextTemp.Times-t.Times != 1 {
			err = fmt.Errorf("[%d] invalid", t.Times)
			return template.NewTemplateFieldError("Times", err)
		}
	}

	return nil
}

func (t *BabyUnlockTalentTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *BabyUnlockTalentTemplate) FileName() string {
	return "tb_baobao_jiesuo.json"
}

func init() {
	template.Register((*BabyUnlockTalentTemplate)(nil))
}
