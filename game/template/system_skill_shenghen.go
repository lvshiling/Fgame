package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	systemskilltypes "fgame/fgame/game/systemskill/types"
	"fmt"
)

//系统圣痕技能配置
type SystemSkillShengHenTemplate struct {
	*SystemSkillShengHenTemplateVO
	sysType     systemskilltypes.SystemSkillType    //系统类型
	subType     systemskilltypes.SystemSkillSubType //技能类型
	needItemMap map[int32]int32
}

func (t *SystemSkillShengHenTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemSkillShengHenTemplate) GetSkillId() int32 {
	return t.SkillId
}

func (t *SystemSkillShengHenTemplate) GetNumber() int32 {
	return t.Number
}

func (t *SystemSkillShengHenTemplate) GetCostGold() int32 {
	return t.CostGold
}

func (t *SystemSkillShengHenTemplate) GetCostSilver() int32 {
	return t.CostSilver
}

func (t *SystemSkillShengHenTemplate) GetLevel() int32 {
	return t.Level
}

func (t *SystemSkillShengHenTemplate) GetNextId() int32 {
	return t.NextId
}

func (t *SystemSkillShengHenTemplate) GetNeedEquipQuality() int32 {
	return t.NeedEquipQuality
}

func (t *SystemSkillShengHenTemplate) GetNeedEquipCount() int32 {
	return t.NeedEquipCount
}

func (t *SystemSkillShengHenTemplate) GetType() systemskilltypes.SystemSkillType {
	return t.sysType
}

func (t *SystemSkillShengHenTemplate) GetSubType() systemskilltypes.SystemSkillSubType {
	return t.subType
}

func (t *SystemSkillShengHenTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *SystemSkillShengHenTemplate) PatchAfterCheck() {
}

func (t *SystemSkillShengHenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 type
	t.sysType = systemskilltypes.SystemSkillType(t.Type)
	if !t.sysType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 subType
	t.subType = systemskilltypes.SystemSkillSubType(t.SubType)
	if !t.subType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SubType)
		err = template.NewTemplateFieldError("SubType", err)
		return
	}

	//验证 物品
	if t.CostItemId != "" {
		t.needItemMap = make(map[int32]int32)
		needItemIdList, err := utils.SplitAsIntArray(t.CostItemId)
		if err != nil {

			return err
		}
		needItemCountList, err := utils.SplitAsIntArray(t.CostItemCount)
		if err != nil {
			return err
		}
		if len(needItemIdList) != len(needItemCountList) {
			err = fmt.Errorf("[%s] invalid", t.CostItemId)
			return template.NewTemplateFieldError("CostItemId", err)
		}
		for index, itemId := range needItemIdList {
			t.needItemMap[itemId] += needItemCountList[index]
		}
	}

	return nil
}

func (t *SystemSkillShengHenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*SystemSkillShengHenTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		nextTo := to.(*SystemSkillShengHenTemplate)

		//验证type
		if t.Type != nextTo.Type {
			err = fmt.Errorf("[%d] invalid", nextTo.Type)
			err = template.NewTemplateFieldError("Type", err)
			return
		}

		//验证type
		if t.SubType != nextTo.SubType {
			err = fmt.Errorf("[%d] invalid", nextTo.SubType)
			err = template.NewTemplateFieldError("SubType", err)
			return
		}

		//验证level
		diffLevel := nextTo.Level - t.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", nextTo.Level)
			err = template.NewTemplateFieldError("Level", err)
			return
		}

		//验证 number
		if nextTo.Number < t.Number {
			err = fmt.Errorf("[%d] invalid", nextTo.Number)
			err = template.NewTemplateFieldError("Number", err)
			return
		}
	}

	//验证 cost_yinliang
	err = validator.MinValidate(float64(t.CostSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CostSilver)
		return template.NewTemplateFieldError("CostSilver", err)
	}

	//验证 skill_id
	to := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}
	skillTemplate := to.(*SkillTemplate)
	skillFirstType := skillTemplate.GetSkillFirstType()
	if t.sysType.GetSkillFirstType() != skillFirstType {
		err = fmt.Errorf("[%d] invalid", t.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}

	// 消耗物品id
	for itemId, num := range t.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", t.CostItemId)
			err = template.NewTemplateFieldError("CostItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.CostItemCount)
			return template.NewTemplateFieldError("CostItemCount", err)
		}
	}

	// 技能激活装备数量条件
	err = validator.MinValidate(float64(t.NeedEquipQuality), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedEquipQuality)
		return template.NewTemplateFieldError("NeedEquipQuality", err)
	}
	err = validator.MinValidate(float64(t.NeedEquipCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedEquipCount)
		return template.NewTemplateFieldError("NeedEquipCount", err)
	}

	return nil
}

func (t *SystemSkillShengHenTemplate) FileName() string {
	return "tb_system_skill_shenghen.json"
}

func init() {
	template.Register((*SystemSkillShengHenTemplate)(nil))
}
