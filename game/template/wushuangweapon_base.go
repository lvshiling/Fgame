package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
	"fmt"
)

type WushuangWeaponBaseTemplate struct {
	*WushuangWeaponBaseTemplateVO
	firstStrengthenTemplate *WushuangWeaponStrengthenTemplate
	strengthenLevelTemplate map[int32]*WushuangWeaponStrengthenTemplate
}

func (t *WushuangWeaponBaseTemplate) TemplateId() int {
	return t.Id
}

func (t *WushuangWeaponBaseTemplate) CalculateLevel(cntExperience int64) int32 {
	temp := t.GetBeginStrengthTemp()
	totalLevel := int32(0)
	for true {
		temp = temp.GetNextStrengthenTemplate()
		if temp == nil {
			break
		}
		curNeedEx := temp.Experience
		if cntExperience > curNeedEx {
			cntExperience -= curNeedEx
			totalLevel = temp.Level
		} else {
			break
		}
	}

	return totalLevel
}

func (t *WushuangWeaponBaseTemplate) IsCanActiveShow(level int32) bool {
	if level >= t.WaiguanJihuoLevel && t.WaiguanJihuoLevel != 0 {
		return true
	} else {
		return false
	}
}

func (t *WushuangWeaponBaseTemplate) GetWearType() wushuangweapontypes.BodyPosWearType {
	weartype := wushuangweapontypes.BodyPosWearType(t.WaiguanType)
	return weartype
}

func (t *WushuangWeaponBaseTemplate) GetBeginStrengthTemp() *WushuangWeaponStrengthenTemplate {
	return t.firstStrengthenTemplate
}

func (t *WushuangWeaponBaseTemplate) GetLevel(cntExperience int64) (level int32, isBorder bool) {
	totalLevel := int32(0)
	temp := t.GetBeginStrengthTemp()
	for true {
		temp = temp.GetNextStrengthenTemplate()
		if temp == nil {
			break
		}
		curNeedEx := temp.Experience
		if cntExperience > curNeedEx {
			cntExperience -= curNeedEx
			totalLevel = temp.Level
		} else if cntExperience == curNeedEx {
			isBorder = true
			break
		} else {
			isBorder = false
			break
		}
	}
	return totalLevel, isBorder
}

func (t *WushuangWeaponBaseTemplate) GetStrengthTemplateByLevel(level int32) *WushuangWeaponStrengthenTemplate {
	temp, ok := t.strengthenLevelTemplate[level]
	if !ok {
		return nil
	}
	return temp
}

//检查有效性
func (t *WushuangWeaponBaseTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//Hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//攻击力
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//防御力
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//强化起始ID
	err = validator.MinValidate(float64(t.StrengthenBeginId), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.StrengthenBeginId)
		return template.NewTemplateFieldError("StrengthenBeginId", err)
	}

	//激活外观需要的等级
	err = validator.MinValidate(float64(t.WaiguanJihuoLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WaiguanJihuoLevel)
		return template.NewTemplateFieldError("WaiguanJihuoLevel", err)
	}

	//外观类型
	err = validator.MinValidate(float64(t.WaiguanType), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WaiguanType)
		return template.NewTemplateFieldError("WaiguanType", err)
	}

	//外观关联ID
	err = validator.MinValidate(float64(t.WaiguanId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WaiguanId)
		return template.NewTemplateFieldError("WaiguanId", err)
	}

	temp := template.GetTemplateService().Get(int(t.StrengthenBeginId), (*WushuangWeaponStrengthenTemplate)(nil))
	strengthTemp, _ := temp.(*WushuangWeaponStrengthenTemplate)
	if strengthTemp == nil {
		err = fmt.Errorf("WushuangStrengthenTemplate[%d] invalid", t.StrengthenBeginId)
		err = template.NewTemplateFieldError("MagicConditionParameter", err)
	}
	t.firstStrengthenTemplate = strengthTemp

	return
}

//组合成需要的数据
func (t *WushuangWeaponBaseTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return
}

//检验后组合
func (t *WushuangWeaponBaseTemplate) PatchAfterCheck() {
	t.strengthenLevelTemplate = make(map[int32]*WushuangWeaponStrengthenTemplate)
	t.firstStrengthenTemplate.SetTotalNeedExperience(int64(0))
	for temp := t.firstStrengthenTemplate; temp != nil; temp = temp.GetNextStrengthenTemplate() {
		t.strengthenLevelTemplate[temp.Level] = temp
		if temp.GetNextStrengthenTemplate() != nil {
			temp.GetNextStrengthenTemplate().SetTotalNeedExperience(temp.GetAllNeedExperience() + temp.GetNextStrengthenTemplate().Experience)
		}
	}
}

func (t *WushuangWeaponBaseTemplate) FileName() string {
	return "tb_wushuang.json"
}

func init() {
	template.Register((*WushuangWeaponBaseTemplate)(nil))
}
