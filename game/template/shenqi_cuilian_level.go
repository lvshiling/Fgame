package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"
)

//神器淬炼等级配置
type ShenQiCuiLianLevelTemplate struct {
	*ShenQiCuiLianLevelTemplateVO
	nextShenQiCuiLianLevelTemplate *ShenQiCuiLianLevelTemplate //下一级
	shenQiType                     shenqitypes.ShenQiType
	cuiLianType                    shenqitypes.SmeltType
	useItemTemplate                *ItemTemplate
}

func (est *ShenQiCuiLianLevelTemplate) TemplateId() int {
	return est.Id
}

func (est *ShenQiCuiLianLevelTemplate) GetNextTemplate() *ShenQiCuiLianLevelTemplate {
	return est.nextShenQiCuiLianLevelTemplate
}

func (est *ShenQiCuiLianLevelTemplate) GetShenQiType() shenqitypes.ShenQiType {
	return est.shenQiType
}

func (est *ShenQiCuiLianLevelTemplate) GetSmeltType() shenqitypes.SmeltType {
	return est.cuiLianType
}

func (est *ShenQiCuiLianLevelTemplate) GetUseItemTemplate() *ItemTemplate {
	return est.useItemTemplate
}

func (et *ShenQiCuiLianLevelTemplate) Patch() (err error) {
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
		tempNextShenQiCuiLianLevelTemplate := template.GetTemplateService().Get(int(et.NextId), (*ShenQiCuiLianLevelTemplate)(nil))
		if tempNextShenQiCuiLianLevelTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		et.nextShenQiCuiLianLevelTemplate = tempNextShenQiCuiLianLevelTemplate.(*ShenQiCuiLianLevelTemplate)
		diffLev := et.nextShenQiCuiLianLevelTemplate.Level - et.Level
		if diffLev != 1 {
			err = fmt.Errorf("[%d] invalid", et.nextShenQiCuiLianLevelTemplate.Level)
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
func (et *ShenQiCuiLianLevelTemplate) PatchAfterCheck() {

}
func (et *ShenQiCuiLianLevelTemplate) Check() (err error) {
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

	//验证 CuiLianType
	et.cuiLianType = shenqitypes.SmeltType(et.CuiLianType)
	if !et.cuiLianType.Valid() {
		err = fmt.Errorf("[%d] invalid", et.CuiLianType)
		return template.NewTemplateFieldError("CuiLianType", err)
	}

	//等级
	err = validator.MinValidate(float64(et.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//需要神器等级
	err = validator.MinValidate(float64(et.NeedShenQiLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.NeedShenQiLevel)
		return template.NewTemplateFieldError("NeedShenQiLevel", err)
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

func (edt *ShenQiCuiLianLevelTemplate) FileName() string {
	return "tb_shenqi_cuilian_level.json"
}

func init() {
	template.Register((*ShenQiCuiLianLevelTemplate)(nil))
}
