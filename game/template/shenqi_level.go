package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"
)

//神器等级配置
type ShenQiLevelTemplate struct {
	*ShenQiLevelTemplateVO
	nextShenQiLevelTemplate *ShenQiLevelTemplate //下一级
	shenQiType              shenqitypes.ShenQiType
	debrisType              shenqitypes.DebrisType
	useItemTemplate         *ItemTemplate
}

func (est *ShenQiLevelTemplate) TemplateId() int {
	return est.Id
}

func (est *ShenQiLevelTemplate) GetNextTemplate() *ShenQiLevelTemplate {
	return est.nextShenQiLevelTemplate
}

func (est *ShenQiLevelTemplate) GetShenQiType() shenqitypes.ShenQiType {
	return est.shenQiType
}

func (est *ShenQiLevelTemplate) GetDebrisType() shenqitypes.DebrisType {
	return est.debrisType
}

func (est *ShenQiLevelTemplate) GetUseItemTemplate() *ItemTemplate {
	return est.useItemTemplate
}

func (et *ShenQiLevelTemplate) Patch() (err error) {
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
		tempNextShenQiLevelTemplate := template.GetTemplateService().Get(int(et.NextId), (*ShenQiLevelTemplate)(nil))
		if tempNextShenQiLevelTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		et.nextShenQiLevelTemplate = tempNextShenQiLevelTemplate.(*ShenQiLevelTemplate)
		diffLev := et.nextShenQiLevelTemplate.Level - et.Level
		if diffLev != 1 {
			err = fmt.Errorf("[%d] invalid", et.nextShenQiLevelTemplate.Level)
			err = template.NewTemplateFieldError("Next Level", err)
			return
		}
	}

	//验证 UseItem
	if et.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(et.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", et.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		et.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(et.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", et.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	return nil
}
func (et *ShenQiLevelTemplate) PatchAfterCheck() {

}
func (et *ShenQiLevelTemplate) Check() (err error) {
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

	//验证 SubType
	et.debrisType = shenqitypes.DebrisType(et.SubType)
	if !et.debrisType.Valid() {
		err = fmt.Errorf("[%d] invalid", et.SubType)
		return template.NewTemplateFieldError("SubType", err)
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

	return nil
}

func (edt *ShenQiLevelTemplate) FileName() string {
	return "tb_shenqi_level.json"
}

func init() {
	template.Register((*ShenQiLevelTemplate)(nil))
}
