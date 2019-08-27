package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
	"math"
)

type YinglingpuLevelSuanTemplate struct {
	*YinglingpuLevelSuanTemplateVO

	battleAttrMap map[propertytypes.BattlePropertyType]int64
}

func (t *YinglingpuLevelSuanTemplate) TemplateId() int {
	return t.Id
}

func (t *YinglingpuLevelSuanTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *YinglingpuLevelSuanTemplate) FileName() string {
	return "tb_yinglingpu_level_suan.json"
}

//组合成需要的数据
func (t *YinglingpuLevelSuanTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 套装属性
	t.battleAttrMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battleAttrMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battleAttrMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battleAttrMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)

	return nil
}

func (t *YinglingpuLevelSuanTemplate) GetConsumeNum(level int32) int32 {
	return int32(math.Ceil(float64(level)/(float64(t.ExpendValueA)/float64(common.MAX_RATE)))) * int32(math.Ceil((float64(t.ExpendValueB) / float64(common.MAX_RATE))))
}

const (
	levelMax = 50
)

func (t *YinglingpuLevelSuanTemplate) GetLevelBattlePropertyMap(level int32) map[propertytypes.BattlePropertyType]int64 {
	if level > levelMax {
		return t.getLevelBattlePropertyMapAbove50(level)
	} else {
		return t.getLevelBattlePropertyMapBelow50(level)
	}
}

func (t *YinglingpuLevelSuanTemplate) getLevelBattlePropertyMapBelow50(level int32) map[propertytypes.BattlePropertyType]int64 {
	battleMap := make(map[propertytypes.BattlePropertyType]int64)
	for k, v := range t.battleAttrMap {
		actualValue := (float64(level) +
			float64(t.AttrValueA)/float64(common.MAX_RATE)*math.Pow(float64(level), float64(t.AttrValueB)/float64(common.MAX_RATE)) +
			float64(t.AttrValueC)/float64(common.MAX_RATE)*math.Pow(float64(level), float64(t.AttrValueD)/float64(common.MAX_RATE))) *
			(float64(v) / float64(common.MAX_RATE))
		battleMap[k] = int64(math.Ceil(actualValue))
	}
	return battleMap
}

func (t *YinglingpuLevelSuanTemplate) getLevelBattlePropertyMapAbove50(level int32) map[propertytypes.BattlePropertyType]int64 {
	battleMap := make(map[propertytypes.BattlePropertyType]int64)
	for k, v := range t.battleAttrMap {
		actualValue := (float64(level)+
			float64(t.AttrValueA)/float64(common.MAX_RATE)*math.Pow(float64(levelMax), float64(t.AttrValueB)/float64(common.MAX_RATE))+
			float64(t.AttrValueC)/float64(common.MAX_RATE)*math.Pow(float64(levelMax), float64(t.AttrValueD)/float64(common.MAX_RATE)))*
			(float64(v)/float64(common.MAX_RATE)) +
			float64(t.AttrValueE)/float64(common.MAX_RATE)*float64(level-levelMax)*(float64(v)/float64(common.MAX_RATE))
		battleMap[k] = int64(math.Ceil(actualValue))
	}
	return battleMap
}

//检查有效性
func (t *YinglingpuLevelSuanTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	for _, v := range t.battleAttrMap {
		err = validator.MinValidate(float64(v), float64(0), false)
		if err != nil {
			return template.NewTemplateFieldError("Attr", fmt.Errorf("[%s] invalid", v))
		}
	}
	err = validator.MinValidate(float64(t.ExpendValueA), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("ExpendValueA", fmt.Errorf("[%s] invalid", t.ExpendValueA))
	}
	err = validator.MinValidate(float64(t.ExpendValueB), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("ExpendValueB", fmt.Errorf("[%s] invalid", t.ExpendValueB))
	}
	err = validator.MinValidate(float64(t.AttrValueA), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("AttrValueB", fmt.Errorf("[%s] invalid", t.AttrValueA))
	}
	err = validator.MinValidate(float64(t.AttrValueB), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("AttrValueB", fmt.Errorf("[%s] invalid", t.AttrValueB))
	}
	err = validator.MinValidate(float64(t.AttrValueC), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("AttrValueC", fmt.Errorf("[%s] invalid", t.AttrValueC))
	}
	err = validator.MinValidate(float64(t.ExpendValueB), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("AttrValueD", fmt.Errorf("[%s] invalid", t.AttrValueD))
	}
	err = validator.MinValidate(float64(t.AttrValueE), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("AttrValueE", fmt.Errorf("[%s] invalid", t.AttrValueE))
	}
	return nil
}

//检验后组合
func (t *YinglingpuLevelSuanTemplate) PatchAfterCheck() {
}

func init() {
	template.Register((*YinglingpuLevelSuanTemplate)(nil))
}
