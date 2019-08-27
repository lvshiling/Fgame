package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//圣痕升级配置
type SystemShengJiShengHenTemplate struct {
	*SystemShengJiShengHenTemplateVO
	useItemTemplate                   *ItemTemplate                              //升级物品
	nextSystemShengJiShengHenTemplate *SystemShengJiShengHenTemplate             //下一级
	battlePropertyMap                 map[propertytypes.BattlePropertyType]int64 //属性
}

func (t *SystemShengJiShengHenTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemShengJiShengHenTemplate) GetLevel() int32 {
	return t.Level
}

func (t *SystemShengJiShengHenTemplate) GetHp() int32 {
	return t.Hp
}

func (t *SystemShengJiShengHenTemplate) GetAttack() int32 {
	return t.Attack
}

func (t *SystemShengJiShengHenTemplate) GetDefence() int32 {
	return t.Defence
}

func (t *SystemShengJiShengHenTemplate) GetUseMoney() int32 {
	return t.UseMoney
}
func (t *SystemShengJiShengHenTemplate) GetUseItem() int32 {
	return t.UseItem
}
func (t *SystemShengJiShengHenTemplate) GetItemCount() int32 {
	return t.ItemCount
}
func (t *SystemShengJiShengHenTemplate) GetUpdateWfb() int32 {
	return t.UpdateWfb
}
func (t *SystemShengJiShengHenTemplate) GetZhufuMax() int32 {
	return t.ZhufuMax
}
func (t *SystemShengJiShengHenTemplate) GetAddMin() int32 {
	return t.AddMin
}
func (t *SystemShengJiShengHenTemplate) GetAddMax() int32 {
	return t.AddMax
}
func (t *SystemShengJiShengHenTemplate) GetTimesMin() int32 {
	return t.TimesMin
}
func (t *SystemShengJiShengHenTemplate) GetTimesMax() int32 {
	return t.TimesMax
}

func (t *SystemShengJiShengHenTemplate) GetUseItemTemplate() *ItemTemplate {
	return t.useItemTemplate
}

func (t *SystemShengJiShengHenTemplate) GetNextTemplate() *SystemShengJiShengHenTemplate {
	return t.nextSystemShengJiShengHenTemplate
}

func (t *SystemShengJiShengHenTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *SystemShengJiShengHenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//下一级
	if t.NextId != 0 {
		tempNextSystemShengJiShengHenTemplate := template.GetTemplateService().Get(int(t.NextId), (*SystemShengJiShengHenTemplate)(nil))
		if tempNextSystemShengJiShengHenTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextSystemShengJiShengHenTemplate = tempNextSystemShengJiShengHenTemplate.(*SystemShengJiShengHenTemplate)
	}

	//验证 UseItem
	if t.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(t.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		t.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)

	return nil
}

func (t *SystemShengJiShengHenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*SystemShengJiShengHenTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		shengJiTemplate := to.(*SystemShengJiShengHenTemplate)

		diffLevel := shengJiTemplate.Level - t.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", t.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证SysType
	sysType := additionsystypes.AdditionSysType(t.SysType)
	if !sysType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SysType)
		return template.NewTemplateFieldError("SysType", err)
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_money
	err = validator.MinValidate(float64(t.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(t.TimesMin), float64(0), true, float64(t.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(t.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(t.AddMin), float64(0), true, float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(t.AddMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	return nil
}
func (t *SystemShengJiShengHenTemplate) PatchAfterCheck() {

}
func (t *SystemShengJiShengHenTemplate) FileName() string {
	return "tb_system_shengji_shenghen.json"
}

func init() {
	template.Register((*SystemShengJiShengHenTemplate)(nil))
}
