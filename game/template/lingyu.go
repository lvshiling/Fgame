package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	commontypes "fgame/fgame/game/common/types"
	lingyutypes "fgame/fgame/game/lingyu/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
	"strconv"
)

//领域配置
type LingyuTemplate struct {
	*LingyuTemplateVO
	nextTemp                *LingyuTemplate
	typ                     lingyutypes.LingyuType          //领域类型
	magicParamMap           map[int32]string                //幻化条件
	magicParamXUMap         map[int32]int32                 //幻化条件1,2
	magicParamIMap          map[int32]int32                 //幻化条件3消耗物品
	useItemTemplate         *ItemTemplate                   //进阶物品
	skillTemplate           *SkillTemplate                  //领域技能
	lingYuUpstarTemplateMap map[int32]*FieldUpstarTemplate  //领域皮肤升星map
	lingYuUpstarTemplate    *FieldUpstarTemplate            //领域皮肤升星
	activateType            commontypes.SpecialAdvancedType //激活类型
	battlePropertyMap       map[propertytypes.BattlePropertyType]int64
}

func (t *LingyuTemplate) TemplateId() int {
	return t.Id
}

func (t *LingyuTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *LingyuTemplate) GetActivateType() commontypes.SpecialAdvancedType {
	return t.activateType
}

func (t *LingyuTemplate) GetTyp() lingyutypes.LingyuType {
	return t.typ
}

func (t *LingyuTemplate) GetUseItemTemplate() *ItemTemplate {
	return t.useItemTemplate
}

func (t *LingyuTemplate) GetMagicParamIMap() map[int32]int32 {
	return t.magicParamIMap
}

func (t *LingyuTemplate) GetMagicParamXUMap() map[int32]int32 {
	return t.magicParamXUMap
}

func (t *LingyuTemplate) GetSkillTemplate() *SkillTemplate {
	return t.skillTemplate
}

func (t *LingyuTemplate) GetNextTemplate() *LingyuTemplate {
	return t.nextTemp
}

func (t *LingyuTemplate) GetLingYuUpstarByLevel(level int32) *FieldUpstarTemplate {
	if v, ok := t.lingYuUpstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *LingyuTemplate) GetIsClear() bool {
	return t.IsClear != 0
}

func (t *LingyuTemplate) PatchAfterCheck() {
	if t.lingYuUpstarTemplate != nil {
		t.lingYuUpstarTemplateMap = make(map[int32]*FieldUpstarTemplate)
		//赋值shenFaUpstarTemplateMap
		for tempTemplate := t.lingYuUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextFieldUpstarTemplate {
			level := tempTemplate.Level
			t.lingYuUpstarTemplateMap[level] = tempTemplate
		}
	}

}
func (t *LingyuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//天魔类型
	t.activateType = commontypes.SpecialAdvancedType(t.ShengjieType)
	if !t.activateType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.ShengjieType)
		return template.NewTemplateFieldError("ShengjieType", err)
	}

	t.typ = lingyutypes.LingyuType(t.Type)
	if !t.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//领域技能
	if t.typ == lingyutypes.LingyuTypeAdvanced {
		if t.Skill == 0 {
			err = fmt.Errorf("[%d] invalid equal 0", t.Skill)
			return template.NewTemplateFieldError("Skill", err)
		}
	}
	if t.Skill != 0 {
		to := template.GetTemplateService().Get(int(t.Skill), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.Skill)
			return template.NewTemplateFieldError("Skill", err)
		}
		skillTemplate := to.(*SkillTemplate)
		if skillTemplate.GetSkillThirdType() == skilltypes.SkillThirdTypeStatic {
			t.skillTemplate = skillTemplate
		}
	}

	//幻化条件
	t.magicParamMap = make(map[int32]string)
	t.magicParamXUMap = make(map[int32]int32)
	t.magicParamIMap = make(map[int32]int32)

	t.magicParamMap[t.MagicConditionType1] = t.MagicConditionParameter1
	t.magicParamMap[t.MagicConditionType2] = t.MagicConditionParameter2
	t.magicParamMap[t.MagicConditionType3] = t.MagicConditionParameter3

	for condType, condParam := range t.magicParamMap {
		cType := lingyutypes.LingyuUCondType(condType)
		if !cType.Valid() {
			err = fmt.Errorf("[%d] invalid", condType)
			return template.NewTemplateFieldError("magic_condition_type", err)
		}
		switch cType {
		case lingyutypes.LingyuUCondTypeX,
			lingyutypes.LingyuUCondTypeU:
			num, err := strconv.ParseInt(condParam, 10, 32)
			if err != nil {
				err = fmt.Errorf("[%s] invalid, err[%s]", condParam, err.Error())
				return template.NewTemplateFieldError("magic_condition_parameter", err)
			}
			t.magicParamXUMap[condType] = int32(num)
			break
		case lingyutypes.LingyuUCondTypeI:
			itemArr, err := utils.SplitAsIntArray(condParam)
			if err != nil {
				return template.NewTemplateFieldError("magic_condition_parameter", err)
			}
			if len(itemArr) != 2 {
				err = fmt.Errorf("[%s] invalid", condParam)
				return template.NewTemplateFieldError("magic_condition_parameter", err)
			}
			t.magicParamIMap[itemArr[0]] = itemArr[1]
			break
		default:
			break
		}
	}
	//幻化条件1、2
	for condType, condParam := range t.magicParamXUMap {
		if condType == int32(lingyutypes.LingyuUCondTypeX) {
			err = validator.MinValidate(float64(condParam), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", condParam)
				err = template.NewTemplateFieldError("MagicConditionParameter", err)
				return
			}
		}
	}
	//幻化条件3
	for item, num := range t.magicParamIMap {
		itemTemplate := template.GetTemplateService().Get(int(item), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("item[%d] invalid", item)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
			return
		}
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("num[%d] invalid", num)
			return template.NewTemplateFieldError("MagicConditionParameter", err)
		}
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

	//验证 lingYu_upgrade_begin_id
	if t.FieldUpstarBeginId != 0 {
		to := template.GetTemplateService().Get(int(t.FieldUpstarBeginId), (*FieldUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.FieldUpstarBeginId)
			return template.NewTemplateFieldError("FieldUpstarBeginId", err)
		}

		lingYuUpstarTemplate, ok := to.(*FieldUpstarTemplate)
		if !ok {
			return fmt.Errorf("MountUpgradeBeginId [%d] invalid", t.FieldUpstarBeginId)
		}

		t.lingYuUpstarTemplate = lingYuUpstarTemplate
	}

	//
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)

	return nil
}

func (t *LingyuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//验证Number
	err = validator.MinValidate(float64(t.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*LingyuTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTemp = to.(*LingyuTemplate)
		diff := t.nextTemp.Number - int32(t.Number)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	//验证 update_wfb
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 use_money
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

	//验证 Hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 Attack
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 Defence
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		err = template.NewTemplateFieldError("Defence", err)
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

	//验证 ZhufuMax
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 UseYinliang
	err = validator.MinValidate(float64(t.UseYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseYinliang)
		err = template.NewTemplateFieldError("UseYinliang", err)
		return
	}

	//验证 skill
	if t.Skill != 0 {
		tempSkillTemplate := template.GetTemplateService().Get(int(t.Skill), (*SkillTemplate)(nil))
		if tempSkillTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.Skill)
			return template.NewTemplateFieldError("Skill", err)
		}
		skilltyp := tempSkillTemplate.(*SkillTemplate).GetSkillFirstType()
		if skilltyp != skilltypes.SkillFirstTypeLingyu {
			err = fmt.Errorf("[%d] invalid", t.Skill)
			return template.NewTemplateFieldError("Skill", err)
		}
	}

	err = validator.MinValidate(float64(t.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}

func (t *LingyuTemplate) FileName() string {
	return "tb_field.json"
}

func init() {
	template.Register((*LingyuTemplate)(nil))
}
