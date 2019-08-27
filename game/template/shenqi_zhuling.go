package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"
)

//神器注灵配置
type ShenQiZhuLingTemplate struct {
	*ShenQiZhuLingTemplateVO
	nextShenQiZhuLingTemplate *ShenQiZhuLingTemplate //下一级
	shenQiType                shenqitypes.ShenQiType
	qiLingType                shenqitypes.QiLingType
	qiLingSubType             shenqitypes.QiLingSubType
}

func (est *ShenQiZhuLingTemplate) TemplateId() int {
	return est.Id
}

func (est *ShenQiZhuLingTemplate) GetNextTemplate() *ShenQiZhuLingTemplate {
	return est.nextShenQiZhuLingTemplate
}

func (est *ShenQiZhuLingTemplate) GetShenQiType() shenqitypes.ShenQiType {
	return est.shenQiType
}

func (est *ShenQiZhuLingTemplate) GetQiLingType() shenqitypes.QiLingType {
	return est.qiLingType
}

func (est *ShenQiZhuLingTemplate) GetQiLingSubType() shenqitypes.QiLingSubType {
	return est.qiLingSubType
}

func (et *ShenQiZhuLingTemplate) Patch() (err error) {
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
		tempNextShenQiZhuLingTemplate := template.GetTemplateService().Get(int(et.NextId), (*ShenQiZhuLingTemplate)(nil))
		if tempNextShenQiZhuLingTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		et.nextShenQiZhuLingTemplate = tempNextShenQiZhuLingTemplate.(*ShenQiZhuLingTemplate)
		diffLev := et.nextShenQiZhuLingTemplate.Level - et.Level
		if diffLev != 1 {
			err = fmt.Errorf("[%d] invalid", et.nextShenQiZhuLingTemplate.Level)
			err = template.NewTemplateFieldError("Next Level", err)
			return
		}
	}

	return nil
}
func (et *ShenQiZhuLingTemplate) PatchAfterCheck() {

}
func (et *ShenQiZhuLingTemplate) Check() (err error) {
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

	//等级
	err = validator.MinValidate(float64(et.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(et.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(et.TimesMin), float64(0), true, float64(et.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(et.TimesMax), float64(et.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(et.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(et.AddMin), float64(0), true, float64(et.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(et.AddMax), float64(et.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(et.ZhufuMax), float64(et.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 NeedZhuLing
	err = validator.MinValidate(float64(et.NeedZhuLing), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.NeedZhuLing)
		err = template.NewTemplateFieldError("NeedZhuLing", err)
		return
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

	//提升属性百分比
	err = validator.MinValidate(float64(et.Percent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Percent)
		return template.NewTemplateFieldError("Percent", err)
	}

	return nil
}

func (edt *ShenQiZhuLingTemplate) FileName() string {
	return "tb_shenqi_zhuling.json"
}

func init() {
	template.Register((*ShenQiZhuLingTemplate)(nil))
}
