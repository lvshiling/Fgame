package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//衣橱配置
type YiChuTemplate struct {
	*YiChuTemplateVO
	wardrobeType             int32 //衣橱类型
	subType                  int32 //子类型
	battlePropertyMap        map[propertytypes.BattlePropertyType]int64
	battlePropertyPercentMap map[propertytypes.BattlePropertyType]int64
}

func (t *YiChuTemplate) TemplateId() int {
	return t.Id
}

func (t *YiChuTemplate) GetType() int32 {
	return t.wardrobeType
}

func (t *YiChuTemplate) GetSubType() int32 {
	return t.subType
}

func (t *YiChuTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *YiChuTemplate) GetBattlePropertyPercentMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyPercentMap
}

func (t *YiChuTemplate) PatchAfterCheck() {

}

func (t *YiChuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// //验证 type
	t.wardrobeType = int32(t.Type)
	// if !t.wardrobeType.Valid() {
	// 	err = fmt.Errorf("[%d] invalid", t.Type)
	// 	err = template.NewTemplateFieldError("Type", err)
	// 	return
	// }

	// //验证 subType
	t.subType = t.SubType
	// t.subType = wardrobetypes.CreateWardrobeSubType(t.wardrobeType, t.SubType)
	// if t.subType == nil || !t.subType.Valid() {
	// 	err = fmt.Errorf("[%d] invalid", t.SubType)
	// 	return template.NewTemplateFieldError("subType", err)
	// }

	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = t.Hp
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = t.Attack
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = t.Defence

	err = validator.MinValidate(float64(t.HpPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.HpPercent)
		return template.NewTemplateFieldError("HpPercent", err)
	}

	err = validator.MinValidate(float64(t.AttackPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AttackPercent)
		return template.NewTemplateFieldError("AttackPercent", err)
	}

	err = validator.MinValidate(float64(t.DefPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DefPercent)
		return template.NewTemplateFieldError("DefPercent", err)
	}
	t.battlePropertyPercentMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeMaxHP] = t.HpPercent
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeAttack] = t.AttackPercent
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeDefend] = t.DefPercent

	return nil
}

func (t *YiChuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 skill_id
	if t.SkillId != 0 {
		to := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SkillId)
			err = template.NewTemplateFieldError("SkillId", err)
			return
		}
		typ := to.(*SkillTemplate).GetSkillFirstType()
		if typ != skilltypes.SkillFirstTypeWardrobe {
			err = fmt.Errorf("[%d] invalid", t.SkillId)
			err = template.NewTemplateFieldError("SkillId", err)
			return
		}
	}

	//验证 skill_id2
	if t.SkillId2 != 0 {
		to := template.GetTemplateService().Get(int(t.SkillId2), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SkillId2)
			err = template.NewTemplateFieldError("SkillId2", err)
			return
		}
		typ := to.(*SkillTemplate).GetSkillFirstType()
		if typ != skilltypes.SkillFirstTypeWardrobe {
			err = fmt.Errorf("[%d] invalid", t.SkillId2)
			err = template.NewTemplateFieldError("SkillId2", err)
			return
		}
	}

	if t.NextId != 0 {
		nextTo := template.GetTemplateService().Get(int(t.NextId), (*YiChuTemplate)(nil))
		if nextTo == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		if t.SkillId != 0 {
			err = fmt.Errorf("[%d] invalid", t.SkillId)
			err = template.NewTemplateFieldError("SkillId", err)
			return
		}

		if t.SkillId2 != 0 {
			err = fmt.Errorf("[%d] invalid", t.SkillId2)
			err = template.NewTemplateFieldError("SkillId2", err)
			return
		}

		nextTemplate := nextTo.(*YiChuTemplate)

		if t.Type != nextTemplate.Type {
			err = fmt.Errorf("[%d] invalid", t.Type)
			err = template.NewTemplateFieldError("Type", err)
			return
		}

		if nextTemplate.Number < t.Number {
			err = fmt.Errorf("[%d] invalid", t.Number)
			err = template.NewTemplateFieldError("Number", err)
			return
		}
	}

	return nil
}

func (t *YiChuTemplate) FileName() string {
	return "tb_yichu.json"
}

func init() {
	template.Register((*YiChuTemplate)(nil))
}
