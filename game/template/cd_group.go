package template

import (
	"fgame/fgame/core/template"
)

type CdGroupTemplate struct {
	*CdGroupTemplateVO
}

func (cdt *CdGroupTemplate) TemplateId() int {
	return cdt.Id
}

func (cdt *CdGroupTemplate) Patch() error {
	return nil
}

func (cdt *CdGroupTemplate) Check() error {
	return nil
}
func (cdt *CdGroupTemplate) PatchAfterCheck() {

}
func (cdt *CdGroupTemplate) FileName() string {
	return "tb_cd_group.json"
}
func init() {
	template.Register((*CdGroupTemplate)(nil))
}
