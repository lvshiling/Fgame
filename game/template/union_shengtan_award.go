package template

import (
	"fgame/fgame/core/template"
)

//圣坛配置
type UnionShengTanAwardTemplate struct {
	*UnionShengTanAwardTemplateVO
	itemMap map[int32]int32
}

func (t *UnionShengTanAwardTemplate) TemplateId() int {
	return t.Id
}

func (t *UnionShengTanAwardTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *UnionShengTanAwardTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *UnionShengTanAwardTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *UnionShengTanAwardTemplate) PatchAfterCheck() {

}

func (t *UnionShengTanAwardTemplate) FileName() string {
	return "tb_union_shengta_award.json"
}

func init() {
	template.Register((*UnionShengTanAwardTemplate)(nil))
}
