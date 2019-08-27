package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"
)

//神器器灵配置
type ShenQiQiLingTemplate struct {
	*ShenQiQiLingTemplateVO
	qiLingType    shenqitypes.QiLingType
	qiLingSubType shenqitypes.QiLingSubType
}

func (est *ShenQiQiLingTemplate) TemplateId() int {
	return est.Id
}

func (est *ShenQiQiLingTemplate) GetQiLingType() shenqitypes.QiLingType {
	return est.qiLingType
}

func (est *ShenQiQiLingTemplate) GetQiLingSubType() shenqitypes.QiLingSubType {
	return est.qiLingSubType
}

func (est *ShenQiQiLingTemplate) IsTypeByArg(typ shenqitypes.QiLingType, subType shenqitypes.QiLingSubType) bool {
	if int32(typ) == int32(est.qiLingType) && subType.SubType() == est.qiLingSubType.SubType() {
		return true
	}
	return false
}

func (et *ShenQiQiLingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	return nil
}
func (et *ShenQiQiLingTemplate) PatchAfterCheck() {

}
func (et *ShenQiQiLingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

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

	//生命
	err = validator.MinValidate(float64(et.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}
	//攻击
	err = validator.MinValidate(float64(et.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}
	//防御
	err = validator.MinValidate(float64(et.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//分解器灵获得的灵气值
	err = validator.MinValidate(float64(et.FenJieGet), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.FenJieGet)
		return template.NewTemplateFieldError("FenJieGet", err)
	}

	return nil
}

func (edt *ShenQiQiLingTemplate) FileName() string {
	return "tb_shenqi_qiling.json"
}

func init() {
	template.Register((*ShenQiQiLingTemplate)(nil))
}
