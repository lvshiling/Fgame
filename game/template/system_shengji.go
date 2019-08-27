package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//升级配置
type SystemShengJiTemplate struct {
	*SystemShengJiTemplateVO
	useItemTemplate           *ItemTemplate                              //升级物品
	nextSystemShengJiTemplate *SystemShengJiTemplate                     //下一级
	battlePropertyMap         map[propertytypes.BattlePropertyType]int64 //属性
}

func (mclt *SystemShengJiTemplate) TemplateId() int {
	return mclt.Id
}

func (t *SystemShengJiTemplate) GetLevel() int32 {
	return t.Level
}

func (t *SystemShengJiTemplate) GetHp() int32 {
	return t.Hp
}

func (t *SystemShengJiTemplate) GetAttack() int32 {
	return t.Attack
}

func (t *SystemShengJiTemplate) GetDefence() int32 {
	return t.Defence
}

func (t *SystemShengJiTemplate) GetUseMoney() int32 {
	return t.UseMoney
}
func (t *SystemShengJiTemplate) GetUseItem() int32 {
	return t.UseItem
}
func (t *SystemShengJiTemplate) GetItemCount() int32 {
	return t.ItemCount
}
func (t *SystemShengJiTemplate) GetUpdateWfb() int32 {
	return t.UpdateWfb
}
func (t *SystemShengJiTemplate) GetZhufuMax() int32 {
	return t.ZhufuMax
}
func (t *SystemShengJiTemplate) GetAddMin() int32 {
	return t.AddMin
}
func (t *SystemShengJiTemplate) GetAddMax() int32 {
	return t.AddMax
}
func (t *SystemShengJiTemplate) GetTimesMin() int32 {
	return t.TimesMin
}
func (t *SystemShengJiTemplate) GetTimesMax() int32 {
	return t.TimesMax
}

func (mclt *SystemShengJiTemplate) GetUseItemTemplate() *ItemTemplate {
	return mclt.useItemTemplate
}

func (mclt *SystemShengJiTemplate) GetNextTemplate() *SystemShengJiTemplate {
	return mclt.nextSystemShengJiTemplate
}

func (mclt *SystemShengJiTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return mclt.battlePropertyMap
}

func (mclt *SystemShengJiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//下一级
	if mclt.NextId != 0 {
		tempNextSystemShengJiTemplate := template.GetTemplateService().Get(int(mclt.NextId), (*SystemShengJiTemplate)(nil))
		if tempNextSystemShengJiTemplate == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		mclt.nextSystemShengJiTemplate = tempNextSystemShengJiTemplate.(*SystemShengJiTemplate)
	}

	//验证 UseItem
	if mclt.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(mclt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", mclt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		mclt.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(mclt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mclt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//属性
	mclt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(mclt.Hp)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(mclt.Attack)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(mclt.Defence)

	return nil
}

func (mclt *SystemShengJiTemplate) Check() (err error) {
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
		to := template.GetTemplateService().Get(int(mclt.NextId), (*SystemShengJiTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		shengJiTemplate := to.(*SystemShengJiTemplate)

		diffLevel := shengJiTemplate.Level - mclt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证SysType
	sysType := additionsystypes.AdditionSysType(mclt.SysType)
	if !sysType.Valid() {
		err = fmt.Errorf("[%d] invalid", mclt.SysType)
		return template.NewTemplateFieldError("SysType", err)
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mclt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
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
func (mclt *SystemShengJiTemplate) PatchAfterCheck() {

}
func (mclt *SystemShengJiTemplate) FileName() string {
	return "tb_system_shengji.json"
}

func init() {
	template.Register((*SystemShengJiTemplate)(nil))
}
