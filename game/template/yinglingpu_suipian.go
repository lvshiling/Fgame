package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	suipiantypes "fgame/fgame/game/yinglingpu/types"
	"fmt"
)

type YinglingPuSuiPianTemplate struct {
	*YinglingpuSuiPianTemplateVO
	battleAttrMap       map[propertytypes.BattlePropertyType]int64
	nextSuiPianTemplate *YinglingPuSuiPianTemplate
}

func (t *YinglingPuSuiPianTemplate) TemplateId() int {
	return t.Id
}

func (t *YinglingPuSuiPianTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *YinglingPuSuiPianTemplate) FileName() string {
	return "tb_yinglingpu_suipian.json"
}

func (t *YinglingPuSuiPianTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.nextSuiPianTemplate != nil {
		if t.nextSuiPianTemplate.SuipianId-t.SuipianId != 1 {
			err = fmt.Errorf("[%d]无效", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return err
		}

	}

	item := template.GetTemplateService().Get(int(t.UseItemId), (*ItemTemplate)(nil))
	if item == nil {
		err = fmt.Errorf("[%d]无效", t.UseItemId)
		err = template.NewTemplateFieldError("UseItemId", err)

		return err
	}

	suiPianPosType := suipiantypes.YingLingPuSuiPianPositionType(t.SuipianId)
	if !suiPianPosType.Valid() {
		err = fmt.Errorf("[%d]无效", t.SuipianId)
		err = template.NewTemplateFieldError("SuipianId", err)
		return
	}

	//使用数量
	err = validator.MinValidate(float64(t.UseItemCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemCount)
		err = template.NewTemplateFieldError("UseItemCount", err)
		return
	}

	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	return nil
}

func (t *YinglingPuSuiPianTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId != 0 {
		nextSuiPian := template.GetTemplateService().Get(int(t.NextId), (*YinglingPuSuiPianTemplate)(nil))
		if nextSuiPian == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		t.nextSuiPianTemplate = nextSuiPian.(*YinglingPuSuiPianTemplate)
	}

	// 套装属性
	t.battleAttrMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battleAttrMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battleAttrMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battleAttrMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	return nil
}

func (t *YinglingPuSuiPianTemplate) PatchAfterCheck() {
	return
}

func (t *YinglingPuSuiPianTemplate) GetNextSuiPianTemplate() *YinglingPuSuiPianTemplate {
	return t.nextSuiPianTemplate
}

func init() {
	template.Register((*YinglingPuSuiPianTemplate)(nil))
}
