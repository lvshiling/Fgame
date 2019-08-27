package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"
)

//神器器灵影响配置
type ShenQiQiLingEffectTemplate struct {
	*ShenQiQiLingEffectTemplateVO
	shenQiType    shenqitypes.ShenQiType
	qiLingType    shenqitypes.QiLingType
	qiLingSubType shenqitypes.QiLingSubType
}

func (est *ShenQiQiLingEffectTemplate) TemplateId() int {
	return est.Id
}

func (est *ShenQiQiLingEffectTemplate) GetShenQiType() shenqitypes.ShenQiType {
	return est.shenQiType
}

func (est *ShenQiQiLingEffectTemplate) GetQiLingType() shenqitypes.QiLingType {
	return est.qiLingType
}

func (est *ShenQiQiLingEffectTemplate) GetQiLingSubType() shenqitypes.QiLingSubType {
	return est.qiLingSubType
}

func (et *ShenQiQiLingEffectTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	return nil
}
func (et *ShenQiQiLingEffectTemplate) PatchAfterCheck() {

}
func (et *ShenQiQiLingEffectTemplate) Check() (err error) {
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

	//验证 QiLingTyp
	et.qiLingType = shenqitypes.QiLingType(et.QiLingTyp)
	if !et.qiLingType.Valid() {
		err = fmt.Errorf("[%d] invalid", et.QiLingTyp)
		return template.NewTemplateFieldError("QiLingTyp", err)
	}

	//器灵部位类型
	et.qiLingSubType = shenqitypes.CreateQiLingSubType(et.qiLingType, et.QiLingSubTyp)
	if et.qiLingSubType == nil || !et.qiLingSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", et.QiLingSubTyp)
		return template.NewTemplateFieldError("QiLingSubTyp", err)
	}

	//对应神器需要达到几级开启
	err = validator.MinValidate(float64(et.NeedLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.NeedLevel)
		return template.NewTemplateFieldError("NeedLevel", err)
	}

	return nil
}

func (edt *ShenQiQiLingEffectTemplate) FileName() string {
	return "tb_shenqi_qiling_effect.json"
}

func init() {
	template.Register((*ShenQiQiLingEffectTemplate)(nil))
}
