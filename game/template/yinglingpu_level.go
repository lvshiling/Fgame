package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

type YinglingpuLevelTemplate struct {
	*YinglingpuLevelTemplateVO
	useItemMap        map[int32]int32 //升级需要的物品和数量
	battleAttrMap     map[propertytypes.BattlePropertyType]int64
	nextLevelTemplate *YinglingpuLevelTemplate
}

func (t *YinglingpuLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *YinglingpuLevelTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *YinglingpuLevelTemplate) FileName() string {
	return "tb_yinglingpu_level.json"
}

//组合成需要的数据
func (t *YinglingpuLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//填充物品数量
	spilitUseItemArray, err := utils.SplitAsIntArray(t.UseItemId)
	if err != nil {
		return template.NewTemplateFieldError("UseItemId", err)
	}
	spilitUseItemCountArray, err := utils.SplitAsIntArray(t.UseItemCount)
	if err != nil {
		return template.NewTemplateFieldError("UseItemCount", err)
	}

	if len(spilitUseItemArray) != len(spilitUseItemCountArray) {
		err = fmt.Errorf("userItem[%s],useCount[%s] 数量不一样", t.UseItemId, t.UseItemCount)
		err = template.NewTemplateFieldError("UseItemId or UseItemCount", err)
		return
	}
	t.useItemMap = make(map[int32]int32)
	for i := 0; i < len(spilitUseItemArray); i++ {
		itemId := spilitUseItemArray[i]
		t.useItemMap[itemId] = spilitUseItemCountArray[i]
	}

	if t.NextId > 0 {
		nextTemplate := template.GetTemplateService().Get(int(t.NextId), (*YinglingpuLevelTemplate)(nil))
		if nextTemplate == nil {
			err = fmt.Errorf("[%d]无效", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return err
		}
		t.nextLevelTemplate = nextTemplate.(*YinglingpuLevelTemplate)
	}

	// 套装属性
	t.battleAttrMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battleAttrMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battleAttrMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battleAttrMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)

	return nil
}

//检查有效性
func (t *YinglingpuLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	if t.nextLevelTemplate != nil {
		if t.Level != t.nextLevelTemplate.Level-1 {
			err = fmt.Errorf("[%d]无效", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return err
		}
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

	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证使用数量
	for itemId, value := range t.useItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.UseItemId)
			err = template.NewTemplateFieldError("UseItemId", err)
			return err
		}

		err = validator.MinValidate(float64(value), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.UseItemCount)
			err = template.NewTemplateFieldError("UseItemCount", err)
			return
		}
	}

	return nil
}

//检验后组合
func (t *YinglingpuLevelTemplate) PatchAfterCheck() {
}

func (t *YinglingpuLevelTemplate) GetUseItemMap() map[int32]int32 {
	return t.useItemMap
}

func (t *YinglingpuLevelTemplate) GetNextLevelTemplate() *YinglingpuLevelTemplate {
	return t.nextLevelTemplate
}

func init() {
	template.Register((*YinglingpuLevelTemplate)(nil))
}
