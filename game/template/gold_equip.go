package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//装备配置
type GoldEquipTemplate struct {
	*GoldEquipTemplateVO
	//套装
	tempTaozhuangTemplate *GoldEquipSuitGroupTemplate
	//一级强化模板
	minGoldEquipStrengthenTemplate *GoldEquipStrengthenTemplate
	//强化模板Map
	strengthenTemplateMap map[int32]*GoldEquipStrengthenTemplate
	//起始开光模板
	minGoldEquipOpenLightTemplate *GoldEquipOpenLightTemplate
	//开光模板map
	openlightMap map[int32]*GoldEquipOpenLightTemplate
	//起始强化升星
	minGoldEquipUpstarTemplate *GoldEquipUpstarTemplate
	//强化升星模板map
	upstarMap map[int32]*GoldEquipUpstarTemplate
	//随机属性模板
	goldequipFuJiaTemp *GoldEquipFuJiaTemplate
	//激活属性条件
	attrConditionMap map[int32]int32
	//属性
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//强化升星最高等级
	maxUpstarLevel int32
	//神铸模板
	godCastingEquipTemp *GodCastingEquipTemplate
}

func (t *GoldEquipTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipTemplate) GetGodCastingForgeSoulMaxLevel() int32 {
	return t.DuanhunLevelMax
}

func (t *GoldEquipTemplate) GetGodCastingEquipTemp() *GodCastingEquipTemplate {
	return t.godCastingEquipTemp
}

func (t *GoldEquipTemplate) IsGodCastingEquip() bool {
	if t.ShenzhuequipLevel == 0 {
		return false
	} else {
		return true
	}
}

func (t *GoldEquipTemplate) GetGodCastingEquipLevel() int32 {
	return t.ShenzhuequipLevel
}

func (t *GoldEquipTemplate) GetActivateCondition(attrIndex int32) int32 {
	return t.attrConditionMap[attrIndex]
}

func (t *GoldEquipTemplate) GetMaxUpstarLevel() int32 {
	return t.maxUpstarLevel
}

func (t *GoldEquipTemplate) RandomGoldEquipAttr() (attrList []int32) {
	if t.goldequipFuJiaTemp == nil {
		return
	}
	return t.goldequipFuJiaTemp.RandomAttr()
}

func (t *GoldEquipTemplate) GetTaozhuangTemplate() *GoldEquipSuitGroupTemplate {
	return t.tempTaozhuangTemplate
}

func (t *GoldEquipTemplate) GetStrengthenTemplate(level int32) *GoldEquipStrengthenTemplate {
	return t.strengthenTemplateMap[level]
}

func (t *GoldEquipTemplate) GetUpstarTemplate(level int32) *GoldEquipUpstarTemplate {
	return t.upstarMap[level]
}

func (t *GoldEquipTemplate) GetOpenLightTemplate(tiems int32) *GoldEquipOpenLightTemplate {
	return t.openlightMap[tiems]
}

func (t *GoldEquipTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

//获取装备技能
func (m *GoldEquipTemplate) GetGoldEquipGroupSuitSkill(equipNum int32) []int32 {
	var skillList []int32
	groupSuitTemplate := m.tempTaozhuangTemplate
	if groupSuitTemplate == nil {
		return skillList
	}
	for _, temp := range groupSuitTemplate.GetSuitEffectTemplate() {
		suitSkillEffectLimitNum := temp.Num
		if equipNum >= suitSkillEffectLimitNum {
			skillList = append(skillList, temp.Value1)
		}
	}

	return skillList
}

func (t *GoldEquipTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//金装附件属性
	if t.FujiaId > 0 {
		tempObj := template.GetTemplateService().Get(int(t.FujiaId), (*GoldEquipFuJiaTemplate)(nil))
		if tempObj == nil {
			err = fmt.Errorf("[%d] invalid", t.FujiaId)
			err = template.NewTemplateFieldError("FujiaId", err)
			return err
		}
		t.goldequipFuJiaTemp = tempObj.(*GoldEquipFuJiaTemplate)
	}

	//套装
	if t.SuitGroup != 0 {
		tempTaozhuangTemplate := template.GetTemplateService().Get(int(t.SuitGroup), (*GoldEquipSuitGroupTemplate)(nil))
		if tempTaozhuangTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.SuitGroup)
			err = template.NewTemplateFieldError("SuitGroup", err)
			return
		}
		t.tempTaozhuangTemplate = tempTaozhuangTemplate.(*GoldEquipSuitGroupTemplate)
	}

	//神铸ID不为0时是可神铸的
	if t.ShenzhuequipId != 0 {
		godCastingTemplate := template.GetTemplateService().Get(int(t.ShenzhuequipId), (*GodCastingEquipTemplate)(nil))
		if godCastingTemplate == nil {
			return fmt.Errorf("ShenzhuequipId [%d] invalid get failed", t.ShenzhuequipId)
		}
		godCastingTemp, ok := godCastingTemplate.(*GodCastingEquipTemplate)
		if !ok {
			return fmt.Errorf("ShenzhuequipId [%d] invalid convert failed", t.ShenzhuequipId)
		}
		t.godCastingEquipTemp = godCastingTemp
	}

	//激活条件
	t.attrConditionMap = make(map[int32]int32)
	t.attrConditionMap[1] = t.NeedStrengthen1
	t.attrConditionMap[2] = t.NeedStrengthen2
	t.attrConditionMap[3] = t.NeedStrengthen3
	t.attrConditionMap[4] = t.NeedStrengthen4
	t.attrConditionMap[5] = t.NeedStrengthen5
	t.attrConditionMap[6] = t.NeedStrengthen6
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	return nil
}

func (t *GoldEquipTemplate) PatchAfterCheck() {
	//动态强化模板
	t.strengthenTemplateMap = make(map[int32]*GoldEquipStrengthenTemplate)
	//赋值 strengthenTemplateMap
	for strengthenTemplate := t.minGoldEquipStrengthenTemplate; strengthenTemplate != nil; strengthenTemplate = strengthenTemplate.nextGoldEquipStrengthenTemplate {
		level := strengthenTemplate.Level
		t.strengthenTemplateMap[level] = strengthenTemplate
	}

	//动态开光模板
	t.openlightMap = make(map[int32]*GoldEquipOpenLightTemplate)
	//赋值  openlightMap
	for openlightTemplate := t.minGoldEquipOpenLightTemplate; openlightTemplate != nil; openlightTemplate = openlightTemplate.nextTemp {
		times := openlightTemplate.Times
		t.openlightMap[times] = openlightTemplate
	}

	//动态升星强化模板
	t.upstarMap = make(map[int32]*GoldEquipUpstarTemplate)
	//赋值 upstarMap
	for upstarTemplate := t.minGoldEquipUpstarTemplate; upstarTemplate != nil; upstarTemplate = upstarTemplate.nextTemp {
		level := upstarTemplate.Level
		t.upstarMap[level] = upstarTemplate
		if t.maxUpstarLevel < level {
			t.maxUpstarLevel = level
		}
	}
}

func (t *GoldEquipTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//动态强化模板
	if t.GoldequipStrenId != 0 {
		tempGoldEquipStrengthenTemplate := template.GetTemplateService().Get(int(t.GoldequipStrenId), (*GoldEquipStrengthenTemplate)(nil))
		if tempGoldEquipStrengthenTemplate == nil {
			return fmt.Errorf("GoldequipStrenId [%d] invalid", t.GoldequipStrenId)
		}
		goldEquipStrengthenTemplate, ok := tempGoldEquipStrengthenTemplate.(*GoldEquipStrengthenTemplate)
		if !ok {
			return fmt.Errorf("GoldequipStrenId [%d] invalid", t.GoldequipStrenId)
		}
		if goldEquipStrengthenTemplate.Level != 0 {
			return fmt.Errorf("GoldequipStrenId [%d] invalid", t.GoldequipStrenId)
		}
		t.minGoldEquipStrengthenTemplate = goldEquipStrengthenTemplate
	}

	//开光起始模板
	if t.GoldeuipOpenlightId != 0 {
		goldEquipOpenLightTemplate := template.GetTemplateService().Get(int(t.GoldeuipOpenlightId), (*GoldEquipOpenLightTemplate)(nil))
		if goldEquipOpenLightTemplate == nil {
			return fmt.Errorf("GoldeuipOpenlightId [%d] invalid", t.GoldeuipOpenlightId)
		}
		openlightTemplate, ok := goldEquipOpenLightTemplate.(*GoldEquipOpenLightTemplate)
		if !ok {
			return fmt.Errorf("GoldeuipOpenlightId [%d] invalid", t.GoldeuipOpenlightId)
		}
		if openlightTemplate.Times != 0 {
			return fmt.Errorf("GoldeuipOpenlightId [%d] invalid", t.GoldeuipOpenlightId)
		}
		t.minGoldEquipOpenLightTemplate = openlightTemplate
	}

	//强化升星起始模板
	if t.GoldeuipUpstarId != 0 {
		goldEquipUpstarTemplate := template.GetTemplateService().Get(int(t.GoldeuipUpstarId), (*GoldEquipUpstarTemplate)(nil))
		if goldEquipUpstarTemplate == nil {
			return fmt.Errorf("GoldeuipUpstarId [%d] invalid", t.GoldeuipUpstarId)
		}
		upstarTemplate, ok := goldEquipUpstarTemplate.(*GoldEquipUpstarTemplate)
		if !ok {
			return fmt.Errorf("GoldeuipUpstarId [%d] invalid", t.GoldeuipUpstarId)
		}
		if upstarTemplate.Level != 0 {
			return fmt.Errorf("GoldeuipUpstarId [%d] invalid", t.GoldeuipUpstarId)
		}
		t.minGoldEquipUpstarTemplate = upstarTemplate
	}

	//生命
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}
	//攻击
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}
	//防御
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}
	//生命万分比
	err = validator.MinValidate(float64(t.HpPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.HpPercent)
		return template.NewTemplateFieldError("HpPercent", err)
	}
	//攻击万分比
	err = validator.MinValidate(float64(t.AttPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AttPercent)
		return template.NewTemplateFieldError("AttPercent", err)
	}
	//防御万分比
	err = validator.MinValidate(float64(t.DefPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DefPercent)
		return template.NewTemplateFieldError("DefPercent", err)
	}

	// 无双神器分解包
	err = validator.MinValidate(float64(t.TunshiDrop), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TunshiDrop)
		err = template.NewTemplateFieldError("TunshiDrop", err)
		return
	}

	//神铸装备关联ID
	err = validator.MinValidate(float64(t.ShenzhuequipId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ShenzhuequipId)
		err = template.NewTemplateFieldError("ShenzhuequipId", err)
		return
	}

	//神铸装备等级
	err = validator.MinValidate(float64(t.ShenzhuequipLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ShenzhuequipLevel)
		err = template.NewTemplateFieldError("ShenzhuequipLevel", err)
		return
	}

	//神铸锻魂最大等级
	err = validator.MinValidate(float64(t.DuanhunLevelMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DuanhunLevelMax)
		err = template.NewTemplateFieldError("DuanhunLevelMax", err)
		return
	}

	return nil
}

func (edt *GoldEquipTemplate) FileName() string {
	return "tb_goldequip.json"
}

func init() {
	template.Register((*GoldEquipTemplate)(nil))
}
