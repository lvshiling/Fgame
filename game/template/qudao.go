package template

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fmt"
)

//充值配置
type QuDaoTemplate struct {
	*QuDaoTemplateVO
	typ             logintypes.SDKType
	allianceVersion alliancetypes.AllianceVersionType
}

func (t *QuDaoTemplate) TemplateId() int {
	return t.Id
}

func (t *QuDaoTemplate) GetType() logintypes.SDKType {
	return t.typ
}

func (t *QuDaoTemplate) GetAllianceVersion() alliancetypes.AllianceVersionType {
	return t.allianceVersion
}

func (t *QuDaoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.typ = logintypes.SDKType(t.Type)
	t.allianceVersion = alliancetypes.AllianceVersionType(t.UnionVersion)

	return
}

func (t *QuDaoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	// 检查平台
	if !t.typ.Valid() {
		err = fmt.Errorf("%d invalid", t.Type)
		return template.NewTemplateFieldError("type", err)
	}

	// 检查仙盟版本
	if !t.allianceVersion.Valid() {
		err = fmt.Errorf("%d invalid", t.UnionVersion)
		return template.NewTemplateFieldError("UnionVersion", err)
	}

	// 传音排行
	err = validator.MinValidate(float64(t.ChuanYinNum), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("ChuanYinNum", err)
	}

	return nil
}

func (t *QuDaoTemplate) PatchAfterCheck() {
}

func (t *QuDaoTemplate) FileName() string {
	return "tb_qudao.json"
}

func init() {
	template.Register((*QuDaoTemplate)(nil))
}
