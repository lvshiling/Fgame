package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//飞升配置
type FeiShengTemplate struct {
	*FeiShengTemplateVO
	nextTemp      *FeiShengTemplate
	battleAttrMap map[propertytypes.BattlePropertyType]int64 //飞升等级属性
}

func (t *FeiShengTemplate) TemplateId() int {
	return t.Id
}

func (t *FeiShengTemplate) GetBattleAttrMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *FeiShengTemplate) GetNextTemplate() *FeiShengTemplate {
	return t.nextTemp
}

func (t *FeiShengTemplate) PatchAfterCheck() {
}

func (t *FeiShengTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//属性
	t.battleAttrMap = make(map[propertytypes.BattlePropertyType]int64)
	if t.Hp > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	}
	if t.Attack > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	}
	if t.Defence > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	}

	return nil
}

func (t *FeiShengTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - t.Id
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*FeiShengTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTemp = to.(*FeiShengTemplate)
	}

	//验证 等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 功德
	err = validator.MinValidate(float64(t.GongDe), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GongDe)
		err = template.NewTemplateFieldError("GongDe", err)
		return
	}

	//验证 概率
	err = validator.MinValidate(float64(t.Rate), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Rate)
		err = template.NewTemplateFieldError("Rate", err)
		return
	}

	//验证 物品
	itemTemp := template.GetTemplateService().Get(int(t.ItemId), (*ItemTemplate)(nil))
	if itemTemp == nil {
		err = fmt.Errorf("[%d] invalid", t.ItemId)
		err = template.NewTemplateFieldError("ItemId", err)
		return
	}

	err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ItemCount)
		err = template.NewTemplateFieldError("ItemCount", err)
		return
	}

	//验证 增加的概率
	err = validator.MinValidate(float64(t.AddRate), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddRate)
		err = template.NewTemplateFieldError("AddRate", err)
		return
	}

	//验证 潜能
	err = validator.MinValidate(float64(t.QnAdd), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.QnAdd)
		err = template.NewTemplateFieldError("QnAdd", err)
		return
	}

	//验证 洗点元宝
	err = validator.MinValidate(float64(t.XidianGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XidianGold)
		err = template.NewTemplateFieldError("XidianGold", err)
		return
	}
	//验证 功德转换率
	err = validator.MinValidate(float64(t.GongdeRatio), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GongdeRatio)
		err = template.NewTemplateFieldError("GongdeRatio", err)
		return
	}
	//验证 经验转换率
	err = validator.MinValidate(float64(t.GiveExpRatio), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GiveExpRatio)
		err = template.NewTemplateFieldError("GiveExpRatio", err)
		return
	}

	return nil
}

func (t *FeiShengTemplate) FileName() string {
	return "tb_feisheng.json"
}

func init() {
	template.Register((*FeiShengTemplate)(nil))
}
