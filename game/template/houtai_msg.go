package template

import (
	"fgame/fgame/core/template"
)

type HoutaiMsgTemplate struct {
	*HoutaiMsgTemplateVO
}

func (t *HoutaiMsgTemplate) TemplateId() int {
	return t.Id
}

func (t *HoutaiMsgTemplate) FileName() string {
	return "tb_houtai_msg.json"
}

//组合成需要的数据
func (t *HoutaiMsgTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *HoutaiMsgTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检验后组合
func (t *HoutaiMsgTemplate) PatchAfterCheck() {

}

func init() {
	template.Register((*HoutaiMsgTemplate)(nil))
}
