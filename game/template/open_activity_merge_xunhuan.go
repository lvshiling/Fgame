package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//合服循环活动配置
type OpenActivityMergeXunHuanTemplate struct {
	*OpenActivityMergeXunHuanTemplateVO
	groupIdList []int32
}

func (t *OpenActivityMergeXunHuanTemplate) TemplateId() int {
	return t.Id
}

func (t *OpenActivityMergeXunHuanTemplate) GetGroupIdList() []int32 {
	return t.groupIdList
}

func (t *OpenActivityMergeXunHuanTemplate) Patch() (err error) {
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

func (t *OpenActivityMergeXunHuanTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 校验 循环活动天数
	err = validator.MinValidate(float64(t.CycleDay), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CycleDay)
		err = template.NewTemplateFieldError("CycleDay", err)
		return
	}

	return nil
}

func (t *OpenActivityMergeXunHuanTemplate) PatchAfterCheck() {

}

func (t *OpenActivityMergeXunHuanTemplate) FileName() string {
	return "tb_circle_hefu.json"
}
 
func init() {
	template.Register((*OpenActivityMergeXunHuanTemplate)(nil))
}
