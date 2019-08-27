package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type MarryXinWuTemplate struct {
	*MarryXinWuTemplateVO
	itemId    int32
	xinWuName string
}

func (t *MarryXinWuTemplate) TemplateId() int {
	return t.Id
}

func (t *MarryXinWuTemplate) FileName() string {
	return "tb_marry_xinwu.json"
}

func (t *MarryXinWuTemplate) Check() (err error) {
	groupTemplate := template.GetTemplateService().Get(int(t.SuitGroup), (*MarryXinWuSuitGroupTemplate)(nil))
	if groupTemplate == nil {
		err = fmt.Errorf("[%d]无效", t.SuitGroup)
		err = template.NewTemplateFieldError("SuitGroup", err)
		return err
	}
	groupTp := groupTemplate.(*MarryXinWuSuitGroupTemplate)

	t.itemId = groupTp.GetXinWuItemId(t.Type)
	item1 := template.GetTemplateService().Get(int(t.itemId), (*ItemTemplate)(nil))
	if item1 == nil {
		err = fmt.Errorf("[%d]无效,[%d]总数", t.itemId, groupTp.GetXinWuItemCount())
		err = template.NewTemplateFieldError("XinWu Template Get ItemId", err)
		return err
	}
	t.xinWuName = groupTp.GetXinWuItemName(t.Type)

	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	return
}

func (t *MarryXinWuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *MarryXinWuTemplate) PatchAfterCheck() {
	return
}

func (t *MarryXinWuTemplate) GetItemId() int32 {
	return t.itemId
}

func (t *MarryXinWuTemplate) GetXinWuName() string {
	return t.xinWuName
}

func init() {
	template.Register((*MarryXinWuTemplate)(nil))
}
