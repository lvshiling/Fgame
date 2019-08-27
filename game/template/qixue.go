package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//泣血枪配置
type QiXueTemplate struct {
	*QiXueTemplateVO
	nextTemp          *QiXueTemplate
	weaponTemplate    *WeaponTemplate //激活兵魂
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
}

//泣血枪掉星几率
func (t *QiXueTemplate) IfHitReduceStar() bool {
	return mathutils.RandomHit(common.MAX_RATE, int(t.GasPercent))
}

//泣血枪掉星几率
func (t *QiXueTemplate) GetRandomReduceStar() int32 {
	return int32(mathutils.RandomRange(int(t.GasMin), int(t.GasMax)))
}

func (t *QiXueTemplate) TemplateId() int {
	return t.Id
}

func (t *QiXueTemplate) GetNextTemp() *QiXueTemplate {
	return t.nextTemp
}

func (t *QiXueTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *QiXueTemplate) GetWeaponTemplate() *WeaponTemplate {
	return t.weaponTemplate
}

func (t *QiXueTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//阶别兵魂
	if t.WeaponId != 0 {
		to := template.GetTemplateService().Get(int(t.WeaponId), (*WeaponTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.WeaponId)
			return template.NewTemplateFieldError("WeaponId", err)
		}
		weaponTemplate, _ := to.(*WeaponTemplate)
		t.weaponTemplate = weaponTemplate
	}

	//属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*QiXueTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*QiXueTemplate)

		diff := t.nextTemp.Level - t.Level
		if diff != 0 {
			if diff != 1 {
				err = fmt.Errorf("NextId [%d] invalid", t.NextId)
				err = template.NewTemplateFieldError("Level", err)
				return
			}
		} else {
			diff := t.nextTemp.Star - t.Star
			if diff != 1 {
				err = fmt.Errorf("NextId [%d] invalid", t.NextId)
				err = template.NewTemplateFieldError("Star", err)
				return
			}
		}
	}

	return nil
}

func (t *QiXueTemplate) PatchAfterCheck() {

}
func (t *QiXueTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证Level 泣血枪等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
	}

	//验证star泣血枪星数
	err = validator.MinValidate(float64(t.Star), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Star)
		err = template.NewTemplateFieldError("Star", err)
		return
	}

	//验证 UseResources
	err = validator.MinValidate(float64(t.UseResources), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseResources)
		err = template.NewTemplateFieldError("UseResources", err)
		return
	}

	//验证 GasMin
	err = validator.RangeValidate(float64(t.GasMin), float64(0), true, float64(t.GasMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GasMin)
		err = template.NewTemplateFieldError("GasMin", err)
		return
	}

	//验证 GasMax
	err = validator.MinValidate(float64(t.GasMax), float64(t.GasMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GasMax)
		err = template.NewTemplateFieldError("GasMax", err)
		return
	}

	//验证 GasPercent
	err = validator.RangeValidate(float64(t.GasPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GasPercent)
		err = template.NewTemplateFieldError("GasPercent", err)
		return
	}

	//验证StarCount
	err = validator.MinValidate(float64(t.StarCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.StarCount)
		err = template.NewTemplateFieldError("StarCount", err)
		return
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

	return nil
}

func (t *QiXueTemplate) FileName() string {
	return "tb_qixue.json"
}

func init() {
	template.Register((*QiXueTemplate)(nil))
}
