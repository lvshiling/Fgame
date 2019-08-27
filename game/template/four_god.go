package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/scene/types"
	"fmt"
)

//四神遗迹常量配置
type FourGodTemplate struct {
	*FourGodTemplateVO
	//boss模板
	bossTemplate *BiologyTemplate
	//special模板
	specialTemplate *BiologyTemplate
	//boss血量下降公告
	gongGaoThresholdList []int32
}

func (t *FourGodTemplate) TemplateId() int {
	return t.Id
}

func (t *FourGodTemplate) GetBossTemplate() *BiologyTemplate {
	return t.bossTemplate
}

func (t *FourGodTemplate) GetSpecialTemplate() *BiologyTemplate {
	return t.specialTemplate
}

func (t *FourGodTemplate) GetGongGaoThresholdList() []int32 {
	return t.gongGaoThresholdList
}

func (t *FourGodTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//boss_id
	to := template.GetTemplateService().Get(int(t.BossId), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.BossId)
		return template.NewTemplateFieldError("BossId", err)
	}

	t.bossTemplate = to.(*BiologyTemplate)

	//special_id
	to = template.GetTemplateService().Get(int(t.SpecialId), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.SpecialId)
		return template.NewTemplateFieldError("SpecialId", err)
	}

	t.specialTemplate = to.(*BiologyTemplate)
	if t.GongGao != "" {
		gongGaoArr, err := utils.SplitAsIntArray(t.GongGao)
		if err != nil {
			return err
		}
		for _, gongGao := range gongGaoArr {
			t.gongGaoThresholdList = append(t.gongGaoThresholdList, gongGao)
		}
	}

	return nil
}

func (t *FourGodTemplate) PatchAfterCheck() {

}

func (t *FourGodTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//key_max
	err = validator.MinValidate(float64(t.KeyMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KeyMax)
		return template.NewTemplateFieldError("KeyMax", err)
	}

	//robot_time
	err = validator.MinValidate(float64(t.RobotTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RobotTime)
		return template.NewTemplateFieldError("RobotTime", err)
	}

	//robot_num
	err = validator.MinValidate(float64(t.RobotNum), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RobotNum)
		return template.NewTemplateFieldError("RobotNum", err)
	}

	//robot_key_min
	err = validator.MinValidate(float64(t.RobotKeyMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RobotKeyMin)
		return template.NewTemplateFieldError("RobotKeyMin", err)
	}

	//robot_key_max
	err = validator.MinValidate(float64(t.RobotKeyMax), float64(t.RobotKeyMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RobotKeyMax)
		return template.NewTemplateFieldError("RobotKeyMax", err)
	}

	//special_time
	err = validator.MinValidate(float64(t.SpecialTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SpecialTime)
		return template.NewTemplateFieldError("SpecialTime", err)
	}

	//special_probability
	err = validator.MinValidate(float64(t.SpecialProbability), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SpecialProbability)
		return template.NewTemplateFieldError("SpecialProbability", err)
	}

	//special_probability_time
	err = validator.MinValidate(float64(t.SpecialProbabilityTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SpecialProbabilityTime)
		return template.NewTemplateFieldError("SpecialProbabilityTime", err)
	}

	//special_num
	err = validator.MinValidate(float64(t.SpecialNum), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SpecialNum)
		return template.NewTemplateFieldError("SpecialNum", err)
	}

	//blacker_buff_id
	to := template.GetTemplateService().Get(int(t.BlackerBuffId), (*BuffTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.BlackerBuffId)
		return template.NewTemplateFieldError("BlackerBuffId", err)
	}

	//boss_time
	err = validator.MinValidate(float64(t.BossTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BossTime)
		return template.NewTemplateFieldError("BossTime", err)
	}

	//item_id
	to = template.GetTemplateService().Get(int(t.ItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}

	//min_stack
	err = validator.MinValidate(float64(t.MinStack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MinStack)
		return template.NewTemplateFieldError("MinStack", err)
	}

	//max_stack
	err = validator.MinValidate(float64(t.MaxStack), float64(t.MinStack), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MaxStack)
		return template.NewTemplateFieldError("MaxStack", err)
	}

	//exist_time
	err = validator.MinValidate(float64(t.ExistTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ExistTime)
		return template.NewTemplateFieldError("ExistTime", err)
	}

	//protected_time
	err = validator.MinValidate(float64(t.ProtectedTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ProtectedTime)
		return template.NewTemplateFieldError("ProtectedTime", err)
	}

	//fail_time
	err = validator.MinValidate(float64(t.FailTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FailTime)
		return template.NewTemplateFieldError("FailTime", err)
	}

	//boss类型
	if t.bossTemplate.GetBiologyScriptType() != types.BiologyScriptTypeFourGodBoss {
		err = fmt.Errorf("[%d] invalid", t.BossId)
		return template.NewTemplateFieldError("BossId", err)
	}

	//特殊怪类型
	if t.specialTemplate.GetBiologyScriptType() != types.BiologyScriptTypeFourGodSpecial {
		err = fmt.Errorf("[%d] invalid", t.SpecialId)
		return template.NewTemplateFieldError("SpecialId", err)
	}

	return nil
}

func (t *FourGodTemplate) FileName() string {
	return "tb_four_god.json"
}

func init() {
	template.Register((*FourGodTemplate)(nil))
}
