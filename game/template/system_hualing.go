package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//化灵配置
type SystemHuaLingTemplate struct {
	*SystemHuaLingTemplateVO
	nextSystemHuaLingTemplate *SystemHuaLingTemplate                     //下一级
	battlePropertyMap         map[propertytypes.BattlePropertyType]int64 //属性
}

func (mclt *SystemHuaLingTemplate) TemplateId() int {
	return mclt.Id
}

func (mclt *SystemHuaLingTemplate) GetNextTemplate() *SystemHuaLingTemplate {
	return mclt.nextSystemHuaLingTemplate
}

func (mclt *SystemHuaLingTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return mclt.battlePropertyMap
}

func (mclt *SystemHuaLingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//下一级
	if mclt.NextId != 0 {
		tempNextSystemHuaLingTemplate := template.GetTemplateService().Get(int(mclt.NextId), (*SystemHuaLingTemplate)(nil))
		if tempNextSystemHuaLingTemplate == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		mclt.nextSystemHuaLingTemplate = tempNextSystemHuaLingTemplate.(*SystemHuaLingTemplate)
	}

	//属性
	mclt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(mclt.Hp)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(mclt.Attack)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(mclt.Defence)

	return nil
}

func (mclt *SystemHuaLingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(mclt.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if mclt.NextId != 0 {
		diff := mclt.NextId - int32(mclt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(mclt.NextId), (*SystemHuaLingTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		huaLingTemplate := to.(*SystemHuaLingTemplate)

		diffLevel := huaLingTemplate.Level - mclt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mclt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 ItemCount
	err = validator.MinValidate(float64(mclt.ItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.ItemCount)
		err = template.NewTemplateFieldError("ItemCount", err)
		return
	}

	//验证use_money
	err = validator.MinValidate(float64(mclt.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mclt.TimesMin), float64(0), true, float64(mclt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mclt.TimesMax), float64(mclt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mclt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mclt.AddMin), float64(0), true, float64(mclt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mclt.AddMax), float64(mclt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(mclt.ZhufuMax), float64(mclt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(mclt.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(mclt.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(mclt.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	return nil
}
func (mclt *SystemHuaLingTemplate) PatchAfterCheck() {

}
func (mclt *SystemHuaLingTemplate) FileName() string {
	return "tb_system_hualing.json"
}

func init() {
	template.Register((*SystemHuaLingTemplate)(nil))
}
