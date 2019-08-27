package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fmt"
)

//仙盟捐献配置
type UnionDonateTemplate struct {
	*UnionDonateTemplateVO
	allianceJuanXianType alliancetypes.AllianceJuanXianType //类型

}

func (t *UnionDonateTemplate) TemplateId() int {
	return t.Id
}

func (t *UnionDonateTemplate) GetJuanXianType() alliancetypes.AllianceJuanXianType {
	return t.allianceJuanXianType
}

func (t *UnionDonateTemplate) GetJuanXianNum() int32 {
	switch t.allianceJuanXianType {
	case alliancetypes.AllianceJuanXianTypeGold:
		return t.DonateGold
	case alliancetypes.AllianceJuanXianTypeSilver:
		return t.DonateSilver
	case alliancetypes.AllianceJuanXianTypeLingPai:
		return t.DonateItemCount
	default:
		return 0
	}
}

func (t *UnionDonateTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证类型
	t.allianceJuanXianType = alliancetypes.AllianceJuanXianType(t.Type)

	return nil
}

func (t *UnionDonateTemplate) PatchAfterCheck() {

}

func (t *UnionDonateTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 type
	if !t.allianceJuanXianType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	switch t.allianceJuanXianType {
	case alliancetypes.AllianceJuanXianTypeGold:
		//最小捐献
		err = validator.MinValidate(float64(t.DonateGold), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("DonateGold", err)
		}
		break
	case alliancetypes.AllianceJuanXianTypeLingPai:
		tempItemTemplate := template.GetTemplateService().Get(int(t.DonateItemId), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.DonateItemId)
			return template.NewTemplateFieldError("DonateItemId", err)
		}
		err = validator.MinValidate(float64(t.DonateItemCount), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("DonateItemCount", err)
		}
		break
	case alliancetypes.AllianceJuanXianTypeSilver:
		//最小捐献
		err = validator.MinValidate(float64(t.DonateSilver), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("DonateSilver", err)
		}
	}

	//最小捐献
	err = validator.MinValidate(float64(t.DonateBuild), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("DonateBuild", fmt.Errorf("[%s] invalid", t.DonateBuild))
	}
	//最小捐献
	err = validator.MinValidate(float64(t.DonateContribution), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("DonateContribution", fmt.Errorf("[%s] invalid", t.DonateContribution))
	}
	//最小捐献
	err = validator.MinValidate(float64(t.DonateLimit), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("DonateLimit", fmt.Errorf("[%s] invalid", t.DonateLimit))
	}

	return nil
}

func (tt *UnionDonateTemplate) FileName() string {
	return "tb_union_donate.json"
}

func init() {
	template.Register((*UnionDonateTemplate)(nil))
}
