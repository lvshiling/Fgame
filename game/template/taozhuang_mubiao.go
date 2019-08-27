package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//套装目标配置
type TaozhuangMuBiaoTemplate struct {
	*TaozhuangMuBiaoTemplateVO
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	mubiaoType        goldequiptypes.TaoZhuangMuBiaoType
}

func (t *TaozhuangMuBiaoTemplate) TemplateId() int {
	return t.Id
}

func (t *TaozhuangMuBiaoTemplate) GetMuBiaoType() goldequiptypes.TaoZhuangMuBiaoType {
	return t.mubiaoType
}

func (t *TaozhuangMuBiaoTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *TaozhuangMuBiaoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = t.Hp
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = t.Attack
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = t.Defence

	return nil
}

func (t *TaozhuangMuBiaoTemplate) PatchAfterCheck() {

}

func (t *TaozhuangMuBiaoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 目标类型
	t.mubiaoType = goldequiptypes.TaoZhuangMuBiaoType(t.Type)
	if !t.mubiaoType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	for typ, val := range t.battlePropertyMap {
		//验证数量至少0
		err = validator.MinValidate(float64(val), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", val)
			err = template.NewTemplateFieldError(typ.String(), err)
			return
		}
	}

	//套装总等级
	if err = validator.MinValidate(float64(t.NeedLevel), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("NeedLevel", err)
		return
	}
	return nil
}

func (t *TaozhuangMuBiaoTemplate) FileName() string {
	return "tb_taozhuangmubiao.json"
}

func init() {
	template.Register((*TaozhuangMuBiaoTemplate)(nil))
}
