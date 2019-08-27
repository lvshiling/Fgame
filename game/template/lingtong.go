package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/lingtong/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//灵童配置
type LingTongTemplate struct {
	*LingTongTemplateVO
	typ                       types.LingTongActivateType         //激活类型
	useItemTemplate           *ItemTemplate                      //激活物品
	shengJiTemplateMap        map[int32]*LingTongShengJiTemplate //灵童升级map
	shengJiTemplate           *LingTongShengJiTemplate           //灵童升级
	peiYangTemplateMap        map[int32]*LingTongPeiYangTemplate //灵童培养map
	peiYangTemplate           *LingTongPeiYangTemplate           //灵童培养
	upstarTemplateMap         map[int32]*LingTongUpstarTemplate  //灵童升星map
	upstarTemplate            *LingTongUpstarTemplate            //灵童升星
	battlePropertyMap         map[propertytypes.BattlePropertyType]int64
	lingTongBattlePropertyMap map[propertytypes.BattlePropertyType]int64
	skillIdList               []int32
}

func (t *LingTongTemplate) TemplateId() int {
	return t.Id
}

func (t *LingTongTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *LingTongTemplate) GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.lingTongBattlePropertyMap
}

func (t *LingTongTemplate) GetTyp() types.LingTongActivateType {
	return t.typ
}

func (t *LingTongTemplate) GetUseItemTemplate() *ItemTemplate {
	return t.useItemTemplate
}

func (t *LingTongTemplate) GetLingTongShengJiByLevel(level int32) *LingTongShengJiTemplate {
	if v, ok := t.shengJiTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *LingTongTemplate) GetLingTongPeiYangByLevel(level int32) *LingTongPeiYangTemplate {
	if v, ok := t.peiYangTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *LingTongTemplate) GetLingTongUpstarByLevel(level int32) *LingTongUpstarTemplate {
	if v, ok := t.upstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *LingTongTemplate) GetSkillIdList() []int32 {
	return t.skillIdList
}

func (t *LingTongTemplate) PatchAfterCheck() {
	t.shengJiTemplateMap = make(map[int32]*LingTongShengJiTemplate)
	if t.shengJiTemplate != nil {

		//赋值shengJiTemplateMap
		for tempTemplate := t.shengJiTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextTemplate {
			level := tempTemplate.Level
			t.shengJiTemplateMap[level] = tempTemplate
		}
	}
	t.peiYangTemplateMap = make(map[int32]*LingTongPeiYangTemplate)
	if t.peiYangTemplate != nil {

		for tempTemplate := t.peiYangTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextTemplate {
			level := tempTemplate.Level
			t.peiYangTemplateMap[level] = tempTemplate
		}
	}
	t.upstarTemplateMap = make(map[int32]*LingTongUpstarTemplate)
	if t.upstarTemplate != nil {

		for tempTemplate := t.upstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextTemplate {
			level := tempTemplate.Level
			t.upstarTemplateMap[level] = tempTemplate
		}
	}
}

func (t *LingTongTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.typ = types.LingTongActivateType(t.JiHuoType)
	if !t.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.JiHuoType)
		return template.NewTemplateFieldError("JiHuoType", err)
	}

	//验证 UseItem
	if t.UseItemId != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(t.UseItemId), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemId)
			err = template.NewTemplateFieldError("UseItemId", err)
			return
		}
		t.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(t.UseItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemCount)
			err = template.NewTemplateFieldError("UseItemCount", err)
			return
		}
	}

	//验证 wing_upgrade_begin_id
	if t.LingTongShengJiId != 0 {
		to := template.GetTemplateService().Get(int(t.LingTongShengJiId), (*LingTongShengJiTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.LingTongShengJiId)
			return template.NewTemplateFieldError("LingTongShengJiId", err)
		}

		shengJiTemplate, ok := to.(*LingTongShengJiTemplate)
		if !ok {
			return fmt.Errorf("LingTongUpgradeBeginId [%d] invalid", t.LingTongShengJiId)
		}

		t.shengJiTemplate = shengJiTemplate
	}

	//验证 wing_upgrade_begin_id
	if t.LingTongPeiYangId != 0 {
		to := template.GetTemplateService().Get(int(t.LingTongPeiYangId), (*LingTongPeiYangTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.LingTongPeiYangId)
			return template.NewTemplateFieldError("LingTongPeiYangId", err)
		}

		peiYangTemplate, ok := to.(*LingTongPeiYangTemplate)
		if !ok {
			return fmt.Errorf("LingTongPeiYangId [%d] invalid", t.LingTongPeiYangId)
		}
		t.peiYangTemplate = peiYangTemplate
	}

	//验证 upstar_begin_id
	if t.LingTongUpstarBeginId != 0 {
		to := template.GetTemplateService().Get(int(t.LingTongUpstarBeginId), (*LingTongUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.LingTongUpstarBeginId)
			return template.NewTemplateFieldError("LingTongUpstarBeginId", err)
		}

		upstarTemp, ok := to.(*LingTongUpstarTemplate)
		if !ok {
			return fmt.Errorf("LingTongUpstarBeginId [%d] invalid", t.LingTongUpstarBeginId)
		}
		t.upstarTemplate = upstarTemp
	}

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

	//属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)

	err = validator.MinValidate(float64(t.LingTongAttack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongAttack)
		return template.NewTemplateFieldError("LingTongAttack", err)
	}

	err = validator.MinValidate(float64(t.LingTongCritical), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongCritical)
		return template.NewTemplateFieldError("LingTongCritical", err)
	}

	err = validator.MinValidate(float64(t.LingTongHit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongHit)
		return template.NewTemplateFieldError("LingTongHit", err)
	}

	err = validator.MinValidate(float64(t.LingTongAbnormality), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongAbnormality)
		return template.NewTemplateFieldError("LingTongAbnormality", err)
	}

	t.lingTongBattlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.lingTongBattlePropertyMap[propertytypes.BattlePropertyTypeAttack] = t.LingTongAttack
	t.lingTongBattlePropertyMap[propertytypes.BattlePropertyTypeCrit] = t.LingTongCritical
	t.lingTongBattlePropertyMap[propertytypes.BattlePropertyTypeHit] = t.LingTongHit
	t.lingTongBattlePropertyMap[propertytypes.BattlePropertyTypeAbnormality] = t.LingTongAbnormality

	return nil
}

func (t *LingTongTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 player_attack_percent
	err = validator.MinValidate(float64(t.PlayerAttackPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PlayerAttackPercent)
		err = template.NewTemplateFieldError("PlayerAttackPercent", err)
		return
	}

	if t.NameItemId != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(t.NameItemId), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.NameItemId)
			err = template.NewTemplateFieldError("NameItemId", err)
			return
		}

		//验证 ItemCount
		err = validator.MinValidate(float64(t.NameItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.NameItemCount)
			err = template.NewTemplateFieldError("NameItemCount", err)
			return
		}
	}

	to := template.GetTemplateService().Get(int(t.LingTongFashionId), (*LingTongFashionTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongFashionId)
		err = template.NewTemplateFieldError("LingTongFashionId", err)
		return
	}

	weaponTempalteVo := template.GetTemplateService().Get(int(t.LingTongWeapon), (*LingTongWeaponTemplate)(nil))
	if weaponTempalteVo == nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongWeapon)
		err = template.NewTemplateFieldError("LingTongWeapon", err)
		return
	}

	tempAttackTemplateVO := template.GetTemplateService().Get(int(t.AttackId), (*SkillTemplate)(nil))
	if tempAttackTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", t.AttackId)
		err = template.NewTemplateFieldError("AttackId", err)
		return
	}
	t.skillIdList = append(t.skillIdList, t.AttackId)

	tempSkillTemplateVO := template.GetTemplateService().Get(int(t.SkillId1), (*SkillTemplate)(nil))
	tempSkillTemplate := tempSkillTemplateVO.(*SkillTemplate)
	if tempSkillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeLingTongSkill {
		err = fmt.Errorf("[%d] invalid", t.SkillId1)
		err = template.NewTemplateFieldError("SkillId1", err)
		return
	}
	t.skillIdList = append(t.skillIdList, t.SkillId1)

	return nil
}

func (t *LingTongTemplate) FileName() string {
	return "tb_lingtong.json"
}

func init() {
	template.Register((*LingTongTemplate)(nil))
}
