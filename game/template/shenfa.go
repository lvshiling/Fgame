package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	shenfatypes "fgame/fgame/game/shenfa/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
	"strconv"
)

//身法配置
type ShenfaTemplate struct {
	*ShenfaTemplateVO
	nextTemp                *ShenfaTemplate
	typ                     shenfatypes.ShenfaType          //身法类型
	magicParamMap           map[int32]string                //幻化条件
	magicParamXUMap         map[int32]int32                 //幻化条件1,2
	magicParamIMap          map[int32]int32                 //幻化条件3消耗物品
	useItemTemplate         *ItemTemplate                   //进阶物品
	battleAttrTemplate      *AttrTemplate                   //阶别属性
	skillTemplate           *SkillTemplate                  //身法技能
	shenFaUpstarTemplateMap map[int32]*ShenFaUpstarTemplate //身法皮肤升星map
	shenFaUpstarTemplate    *ShenFaUpstarTemplate           //身法皮肤升星
}

func (t *ShenfaTemplate) TemplateId() int {
	return t.Id
}

func (t *ShenfaTemplate) GetTyp() shenfatypes.ShenfaType {
	return t.typ
}

func (t *ShenfaTemplate) GetUseItemTemplate() *ItemTemplate {
	return t.useItemTemplate
}

func (t *ShenfaTemplate) GetMagicParamIMap() map[int32]int32 {
	return t.magicParamIMap
}

func (t *ShenfaTemplate) GetMagicParamXUMap() map[int32]int32 {
	return t.magicParamXUMap
}

func (t *ShenfaTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return t.battleAttrTemplate
}

func (t *ShenfaTemplate) GetSkillTemplate() *SkillTemplate {
	return t.skillTemplate
}

func (t *ShenfaTemplate) GetNextTemplate() *ShenfaTemplate {
	return t.nextTemp
}

func (t *ShenfaTemplate) GetIsClear() bool {
	return t.IsClear != 0
}

func (t *ShenfaTemplate) GetShenFaUpstarByLevel(level int32) *ShenFaUpstarTemplate {
	if v, ok := t.shenFaUpstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *ShenfaTemplate) PatchAfterCheck() {
	if t.shenFaUpstarTemplate != nil {
		t.shenFaUpstarTemplateMap = make(map[int32]*ShenFaUpstarTemplate)
		//赋值shenFaUpstarTemplateMap
		for tempTemplate := t.shenFaUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextShenFaUpstarTemplate {
			level := tempTemplate.Level
			t.shenFaUpstarTemplateMap[level] = tempTemplate
		}
	}

}
func (t *ShenfaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.typ = shenfatypes.ShenfaType(t.Type)
	if !t.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	if t.Attr != 0 {
		//阶别attr属性
		to := template.GetTemplateService().Get(int(t.Attr), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		t.battleAttrTemplate = attrTemplate
	}
	//身法技能
	if t.Skill != 0 {
		to := template.GetTemplateService().Get(int(t.Skill), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.Skill)
			return template.NewTemplateFieldError("Skill", err)
		}
		t.skillTemplate = to.(*SkillTemplate)
	}

	//幻化条件
	t.magicParamMap = make(map[int32]string)
	t.magicParamXUMap = make(map[int32]int32)
	t.magicParamIMap = make(map[int32]int32)

	t.magicParamMap[t.MagicConditionType1] = t.MagicConditionParameter1
	t.magicParamMap[t.MagicConditionType2] = t.MagicConditionParameter2
	t.magicParamMap[t.MagicConditionType3] = t.MagicConditionParameter3

	for condType, condParam := range t.magicParamMap {
		cType := shenfatypes.ShenfaUCondType(condType)
		if !cType.Valid() {
			err = fmt.Errorf("[%d] invalid", condType)
			return template.NewTemplateFieldError("magic_condition_type", err)
		}
		switch cType {
		case shenfatypes.ShenfaUCondTypeX,
			shenfatypes.ShenfaUCondTypeU:
			num, err := strconv.ParseInt(condParam, 10, 32)
			if err != nil {
				return err
			}
			t.magicParamXUMap[condType] = int32(num)
			break
		case shenfatypes.ShenfaUCondTypeI:
			itemArr, err := utils.SplitAsIntArray(condParam)
			if err != nil {
				return err
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
		if condType == int32(shenfatypes.ShenfaUCondTypeX) {
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

	//验证 shenfa_upgrade_begin_id
	if t.ShenfaUpstarBeginId != 0 {
		to := template.GetTemplateService().Get(int(t.ShenfaUpstarBeginId), (*ShenFaUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.ShenfaUpstarBeginId)
			return template.NewTemplateFieldError("ShenfaUpstarBeginId", err)
		}

		shenFaUpstarTemplate, ok := to.(*ShenFaUpstarTemplate)
		if !ok {
			return fmt.Errorf("ShenfaUpstarBeginId [%d] invalid", t.ShenfaUpstarBeginId)
		}

		t.shenFaUpstarTemplate = shenFaUpstarTemplate
	}

	return nil
}

func (t *ShenfaTemplate) Check() (err error) {
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

		to := template.GetTemplateService().Get(int(t.NextId), (*ShenfaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*ShenfaTemplate)

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
		if skilltyp != skilltypes.SkillFirstTypeShenfa {
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

func (t *ShenfaTemplate) FileName() string {
	return "tb_shenfa.json"
}

func init() {
	template.Register((*ShenfaTemplate)(nil))
}
