package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//开服活动配置
type OpenActivityXunHuanTemplate struct {
	*OpenActivityXunHuanTemplateVO
	groupIdList []int32
}

func (t *OpenActivityXunHuanTemplate) TemplateId() int {
	return t.Id
}

func (t *OpenActivityXunHuanTemplate) GetGroupIdList() []int32 {
	return t.groupIdList
}

func (t *OpenActivityXunHuanTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.groupIdList, err = utils.SplitAsIntArray(t.GroupId)
	if err != nil {
		return template.NewTemplateFieldError("GroupId", err)
	}

	return nil
}

func (t *OpenActivityXunHuanTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 校验 循环活动组
	err = validator.MinValidate(float64(t.ArrIndex), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ArrIndex)
		err = template.NewTemplateFieldError("ArrIndex", err)
		return
	}

	// 校验 循环活动天数
	err = validator.MinValidate(float64(t.CycleDay), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CycleDay)
		err = template.NewTemplateFieldError("CycleDay", err)
		return
	}

	return nil
}

func (t *OpenActivityXunHuanTemplate) PatchAfterCheck() {

}

func (t *OpenActivityXunHuanTemplate) FileName() string {
	return "tb_circle_openserver.json"
}

func init() {
	template.Register((*OpenActivityXunHuanTemplate)(nil))
}
