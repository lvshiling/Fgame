package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"
)

//神器配置
type ShenQiCuiLianTemplate struct {
	*ShenQiCuiLianTemplateVO
	nextShenQiCuiLianTemplate *ShenQiCuiLianTemplate //下一级
	shenQiType                shenqitypes.ShenQiType
}

func (est *ShenQiCuiLianTemplate) TemplateId() int {
	return est.Id
}

func (est *ShenQiCuiLianTemplate) GetNextTemplate() *ShenQiCuiLianTemplate {
	return est.nextShenQiCuiLianTemplate
}

func (est *ShenQiCuiLianTemplate) GetShenQiType() shenqitypes.ShenQiType {
	return est.shenQiType
}

func (et *ShenQiCuiLianTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//下一级
	if et.NextId != 0 {
		diff := et.NextId - int32(et.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		tempNextShenQiCuiLianTemplate := template.GetTemplateService().Get(int(et.NextId), (*ShenQiCuiLianTemplate)(nil))
		if tempNextShenQiCuiLianTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		et.nextShenQiCuiLianTemplate = tempNextShenQiCuiLianTemplate.(*ShenQiCuiLianTemplate)
		diffLev := et.nextShenQiCuiLianTemplate.Level - et.Level
		if diffLev != 1 {
			err = fmt.Errorf("[%d] invalid", et.nextShenQiCuiLianTemplate.Level)
			err = template.NewTemplateFieldError("Next Level", err)
			return
		}
	}

	return nil
}
func (et *ShenQiCuiLianTemplate) PatchAfterCheck() {

}
func (et *ShenQiCuiLianTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//验证 ShenQiType
	et.shenQiType = shenqitypes.ShenQiType(et.ShenQiType)
	if !et.shenQiType.Valid() {
		err = fmt.Errorf("[%d] invalid", et.ShenQiType)
		return template.NewTemplateFieldError("ShenQiType", err)
	}

	//等级
	err = validator.MinValidate(float64(et.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//提升属性百分比
	err = validator.MinValidate(float64(et.Percent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Percent)
		return template.NewTemplateFieldError("Percent", err)
	}

	return nil
}

func (edt *ShenQiCuiLianTemplate) FileName() string {
	return "tb_shenqi_cuilian.json"
}

func init() {
	template.Register((*ShenQiCuiLianTemplate)(nil))
}
