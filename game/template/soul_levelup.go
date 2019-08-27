package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	soultypes "fgame/fgame/game/soul/types"

	"fmt"
)

//护体仙羽配置
type SoulLevelUpTemplate struct {
	*SoulLevelUpTemplateVO
	soulType soultypes.SoulType
}

func (slt *SoulLevelUpTemplate) TemplateId() int {
	return slt.Id
}

func (slt *SoulLevelUpTemplate) GetSoulType() soultypes.SoulType {
	return slt.soulType
}

func (slt *SoulLevelUpTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(slt.FileName(), slt.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (slt *SoulLevelUpTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(slt.FileName(), slt.TemplateId(), err)
			return
		}
	}()

	slt.soulType = soultypes.SoulType(slt.Type)
	if !slt.soulType.Valid() {
		err = fmt.Errorf("[%d] invalid", slt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 next_id
	if slt.NextId != 0 {
		diff := slt.NextId - int32(slt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", slt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(slt.NextId), (*SoulLevelUpTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", slt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		soulLevelUpTemplate := to.(*SoulLevelUpTemplate)
		nextLevel := soulLevelUpTemplate.Level
		diffLevel := nextLevel - slt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", slt.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证 level
	err = validator.MinValidate(float64(slt.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(slt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_silver
	err = validator.MinValidate(float64(slt.UseSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.UseSilver)
		err = template.NewTemplateFieldError("UseSilver", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(slt.TimesMin), float64(0), true, float64(slt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(slt.TimesMax), float64(slt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(slt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(slt.AddMin), float64(0), true, float64(slt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(slt.AddMax), float64(slt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(slt.ZhufuMax), float64(slt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 HP
	err = validator.MinValidate(float64(slt.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(slt.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(slt.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", slt.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	return nil
}
func (slt *SoulLevelUpTemplate) PatchAfterCheck() {

}
func (slt *SoulLevelUpTemplate) FileName() string {
	return "tb_soul_levelup.json"
}

func init() {
	template.Register((*SoulLevelUpTemplate)(nil))
}
