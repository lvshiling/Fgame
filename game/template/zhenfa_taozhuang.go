package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"

	"fmt"
)

//阵法套装配置
type ZhenFaTaoZhuangTemplate struct {
	*ZhenFaTaoZhuangTemplateVO
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	preTemplate       *ZhenFaTaoZhuangTemplate
}

func (mt *ZhenFaTaoZhuangTemplate) TemplateId() int {
	return mt.Id
}

func (mt *ZhenFaTaoZhuangTemplate) GetPreTemplate() *ZhenFaTaoZhuangTemplate {
	return mt.preTemplate
}

func (mt *ZhenFaTaoZhuangTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return mt.battlePropertyMap
}

func (mt *ZhenFaTaoZhuangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//属性
	mt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	if mt.Hp != 0 {
		mt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = mt.Hp
	}
	if mt.Attack != 0 {
		mt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = mt.Attack
	}
	if mt.Defence != 0 {
		mt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = mt.Defence
	}

	return nil
}

func (mt *ZhenFaTaoZhuangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证level
	err = validator.MinValidate(float64(mt.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	preId := mt.Id - 1
	to := template.GetTemplateService().Get(int(preId), (*ZhenFaTaoZhuangTemplate)(nil))
	if to != nil {
		mt.preTemplate = to.(*ZhenFaTaoZhuangTemplate)
		if mt.preTemplate.Level >= mt.Level {
			err = fmt.Errorf("[%d] invalid", mt.Level)
			err = template.NewTemplateFieldError("Level", err)
			return
		}
	}
	return nil
}
func (mt *ZhenFaTaoZhuangTemplate) PatchAfterCheck() {

}
func (mt *ZhenFaTaoZhuangTemplate) FileName() string {
	return "tb_zhenfa_taozhuang.json"
}

func init() {
	template.Register((*ZhenFaTaoZhuangTemplate)(nil))
}
