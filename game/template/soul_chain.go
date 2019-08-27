package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	soultypes "fgame/fgame/game/soul/types"
	"fmt"
)

//帝魂锁链配置
type SoulChainTemplate struct {
	*SoulChainTemplateVO
	effectTyp soultypes.SoulEffectType //效果类型
}

func (sct *SoulChainTemplate) TemplateId() int {
	return sct.Id
}

func (sct *SoulChainTemplate) PatchAfterCheck() {

}

func (sct *SoulChainTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(sct.FileName(), sct.TemplateId(), err)
			return
		}
	}()

	//验证 type
	sct.effectTyp = soultypes.SoulEffectType(sct.Type)
	if !sct.effectTyp.Valid() {
		err = fmt.Errorf("[%d] invalid", sct.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	return nil
}

func (sct *SoulChainTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(sct.FileName(), sct.TemplateId(), err)
			return
		}
	}()

	//验证 need_count
	err = validator.MinValidate(float64(sct.NeedCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sct.NeedCount)
		return template.NewTemplateFieldError("NeedCount", err)
	}

	//验证 value
	err = validator.MinValidate(float64(sct.Value), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sct.Value)
		return template.NewTemplateFieldError("Value", err)
	}

	return nil
}

func (sct *SoulChainTemplate) FileName() string {
	return "tb_soul_chain.json"
}

func init() {
	template.Register((*SoulChainTemplate)(nil))
}
