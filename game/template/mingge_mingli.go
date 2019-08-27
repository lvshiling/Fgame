package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/mingge/types"
	"fmt"
)

//命格命理配置
type MingGeMingLiTemplate struct {
	*MingGeMingLiTemplateVO
	mingGongType        types.MingGongType
	mingGongSubType     types.MingGongAllSubType
	zhiDingPropertyType types.MingGePropertyType
	propertyPoolList    []types.MingGePropertyType
}

func (mt *MingGeMingLiTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MingGeMingLiTemplate) GetMingGongType() types.MingGongType {
	return mt.mingGongType
}

func (mt *MingGeMingLiTemplate) GetMingGongSubType() types.MingGongAllSubType {
	return mt.mingGongSubType
}

func (mt *MingGeMingLiTemplate) GetZhiDingPropertyType() types.MingGePropertyType {
	return mt.zhiDingPropertyType
}

func (mt *MingGeMingLiTemplate) GetPropertyPoolList() []types.MingGePropertyType {
	return mt.propertyPoolList
}

func (mt *MingGeMingLiTemplate) Patch() (err error) {
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

	mt.mingGongSubType = types.MingGongAllSubType(mt.SubType)
	if !mt.mingGongSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	mt.zhiDingPropertyType = types.MingGePropertyType(mt.ZhiDingAttr)
	if !mt.zhiDingPropertyType.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.ZhiDingAttr)
		err = template.NewTemplateFieldError("ZhiDingAttr", err)
		return
	}

	if mt.AttrPool&mt.zhiDingPropertyType.Mask() == 0 {
		err = fmt.Errorf("[%d]  invalid", mt.AttrPool)
		return template.NewTemplateFieldError("AttrPool", err)
	}

	for i := types.MingGePropertyTypeLife; i <= types.MingGePropertyTypeBlock; i++ {
		if i.Mask()&mt.AttrPool != 0 {
			mt.propertyPoolList = append(mt.propertyPoolList, i)
		}
	}

	return nil
}

func (mt *MingGeMingLiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(mt.AttrPool), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.AttrPool)
		err = template.NewTemplateFieldError("AttrPool", err)
		return
	}

	err = validator.MinValidate(float64(mt.AttrOne), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.AttrOne)
		err = template.NewTemplateFieldError("AttrOne", err)
		return
	}

	err = validator.RangeValidate(float64(mt.ZhiDingRate1), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ZhiDingRate1)
		err = template.NewTemplateFieldError("ZhiDingRate1", err)
		return
	}

	err = validator.RangeValidate(float64(mt.ZhiDingRate2), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ZhiDingRate2)
		err = template.NewTemplateFieldError("ZhiDingRate2", err)
		return
	}

	err = validator.RangeValidate(float64(mt.ZhiDingRate3), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ZhiDingRate3)
		err = template.NewTemplateFieldError("ZhiDingRate3", err)
		return
	}

	if mt.UseItemId != 0 {
		to := template.GetTemplateService().Get(int(mt.UseItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.UseItemId)
			err = template.NewTemplateFieldError("UseItemId", err)
			return
		}
	}

	err = validator.MinValidate(float64(mt.UseItemOne), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UseItemOne)
		err = template.NewTemplateFieldError("UseItemOne", err)
		return
	}

	err = validator.MinValidate(float64(mt.CoefficientAttr1), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.CoefficientAttr1)
		err = template.NewTemplateFieldError("CoefficientAttr1", err)
		return
	}

	err = validator.MinValidate(float64(mt.CoefficientUse1), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.CoefficientUse1)
		err = template.NewTemplateFieldError("CoefficientUse1", err)
		return
	}

	err = validator.MinValidate(float64(mt.CoefficientUse2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.CoefficientUse2)
		err = template.NewTemplateFieldError("CoefficientUse2", err)
		return
	}

	err = validator.MinValidate(float64(mt.XilianLimitCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.XilianLimitCount)
		err = template.NewTemplateFieldError("XilianLimitCount", err)
		return
	}

	err = validator.MinValidate(float64(mt.ShouYiPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ShouYiPercent)
		err = template.NewTemplateFieldError("ShouYiPercent", err)
		return
	}

	err = validator.MinValidate(float64(mt.CoefficientZhiDing), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.CoefficientZhiDing)
		err = template.NewTemplateFieldError("CoefficientZhiDing", err)
		return
	}

	return nil
}

func (mt *MingGeMingLiTemplate) GetCoefficientUse1() float64 {
	return mt.CoefficientUse1 / common.MAX_RATE
}

func (mt *MingGeMingLiTemplate) GetCoefficientAttr1() float64 {
	return mt.CoefficientAttr1 / common.MAX_RATE
}

func (mt *MingGeMingLiTemplate) PatchAfterCheck() {

}

func (mt *MingGeMingLiTemplate) FileName() string {
	return "tb_mingge_mingli.json"
}

func init() {
	template.Register((*MingGeMingLiTemplate)(nil))
}
