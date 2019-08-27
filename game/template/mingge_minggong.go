package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/mingge/types"
	"fmt"
)

//命格命宫配置
type MingGeMingGongTemplate struct {
	*MingGeMingGongTemplateVO
	mingGongType   types.MingGongType
	parentTemplate *MingGeMingGongTemplate
}

func (mt *MingGeMingGongTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MingGeMingGongTemplate) GetMingGongType() types.MingGongType {
	return mt.mingGongType
}

func (mt *MingGeMingGongTemplate) GetParentTemplate() *MingGeMingGongTemplate {
	return mt.parentTemplate
}

func (mt *MingGeMingGongTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	mt.mingGongType = types.MingGongType(mt.Type)
	if !mt.mingGongType.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	if mt.ParentId != 0 {
		to := template.GetTemplateService().Get(int(mt.ParentId), (*MingGeMingGongTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.ParentId)
			err = template.NewTemplateFieldError("ParentId", err)
			return
		}
		mt.parentTemplate = to.(*MingGeMingGongTemplate)
	}
	return nil
}

func (mt *MingGeMingGongTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	if mt.NextId != 0 {
		to := template.GetTemplateService().Get(int(mt.NextId), (*MingGeMingGongTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		nextTemplate := to.(*MingGeMingGongTemplate)
		if nextTemplate.NeedLevel < mt.NeedLevel {
			err = fmt.Errorf("[%d] invalid", mt.NeedLevel)
			err = template.NewTemplateFieldError("NeedLevel", err)
			return
		}

		if nextTemplate.NeedZhuanShu < mt.NeedZhuanShu {
			err = fmt.Errorf("[%d] invalid", mt.NeedZhuanShu)
			err = template.NewTemplateFieldError("NeedZhuanShu", err)
			return
		}
	}

	//验证 NeedParentZhanLi
	err = validator.MinValidate(float64(mt.NeedParentZhanLi), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedParentZhanLi)
		err = template.NewTemplateFieldError("NeedParentZhanLi", err)
		return
	}

	err = validator.MinValidate(float64(mt.NeedLevel), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedLevel)
		err = template.NewTemplateFieldError("NeedLevel", err)
		return
	}

	err = validator.MinValidate(float64(mt.NeedZhuanShu), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedZhuanShu)
		err = template.NewTemplateFieldError("NeedZhuanShu", err)
		return
	}

	return nil
}

func (mt *MingGeMingGongTemplate) PatchAfterCheck() {

}

func (mt *MingGeMingGongTemplate) FileName() string {
	return "tb_mingge_minggong.json"
}

func init() {
	template.Register((*MingGeMingGongTemplate)(nil))
}
