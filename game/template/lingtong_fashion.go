package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/lingtong/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//灵童时装配置
type LingTongFashionTemplate struct {
	*LingTongFashionTemplateVO
	fashionType               types.LingTongFashionType                //类型
	battleAttrTemplate        *AttrTemplate                            //阶别属性
	needItemMap               map[int32]int32                          //激活需要物品
	fashionUpstarTemplateMap  map[int32]*LingTongFashionUpstarTemplate //灵童时装升星map
	fashionUpstarTemplate     *LingTongFashionUpstarTemplate           //灵童时装升星
	battlePropertyMap         map[propertytypes.BattlePropertyType]int64
	lingTongBattlePropertyMap map[propertytypes.BattlePropertyType]int64
}

//永久性
func (t *LingTongFashionTemplate) Permanent() bool {
	return t.Time == 0
}

func (t *LingTongFashionTemplate) TemplateId() int {
	return t.Id
}

func (t *LingTongFashionTemplate) GetFashionType() types.LingTongFashionType {
	return t.fashionType
}

func (t *LingTongFashionTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *LingTongFashionTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *LingTongFashionTemplate) GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.lingTongBattlePropertyMap
}

func (t *LingTongFashionTemplate) GetLingTongFashionUpstarByLevel(level int32) *LingTongFashionUpstarTemplate {
	if v, ok := t.fashionUpstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *LingTongFashionTemplate) PatchAfterCheck() {
	if t.fashionUpstarTemplate != nil {
		t.fashionUpstarTemplateMap = make(map[int32]*LingTongFashionUpstarTemplate)
		//赋值fashionUpstarTemplateMap
		for tempTemplate := t.fashionUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextLingTongFashionUpstarTemplate {
			level := tempTemplate.Level
			t.fashionUpstarTemplateMap[level] = tempTemplate
		}
	}
}

func (t *LingTongFashionTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 Type
	t.fashionType = types.LingTongFashionType(t.Type)
	if !t.fashionType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	t.needItemMap = make(map[int32]int32)
	if t.NeedItemId != 0 {
		needItemTemplateVO := template.GetTemplateService().Get(int(t.NeedItemId), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(t.NeedItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
		t.needItemMap[t.NeedItemId] = t.NeedItemCount
	}

	//验证 fashion_upgrade_begin_id
	if t.LingTongUpstarId != 0 {
		to := template.GetTemplateService().Get(int(t.LingTongUpstarId), (*LingTongFashionUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.LingTongUpstarId)
			return template.NewTemplateFieldError("LingTongUpstarId", err)
		}

		fashionUpstarTemplate, ok := to.(*LingTongFashionUpstarTemplate)
		if !ok {
			return fmt.Errorf("LingTongUpstarId [%d] invalid", t.LingTongUpstarId)
		}
		if fashionUpstarTemplate.Level != 1 {
			return fmt.Errorf("LingTongUpstarId Level [%d] invalid", fashionUpstarTemplate.Level)
		}
		t.fashionUpstarTemplate = fashionUpstarTemplate
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

func (t *LingTongFashionTemplate) PatchAterCheck() {
	if t.fashionUpstarTemplate != nil {
		t.fashionUpstarTemplateMap = make(map[int32]*LingTongFashionUpstarTemplate)
		//赋值fashionUpstarTemplateMap
		for tempTemplate := t.fashionUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextLingTongFashionUpstarTemplate {
			level := tempTemplate.Level
			t.fashionUpstarTemplateMap[level] = tempTemplate
		}
	}
}

func (t *LingTongFashionTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *LingTongFashionTemplate) FileName() string {
	return "tb_lingtong_fashion.json"
}

func init() {
	template.Register((*LingTongFashionTemplate)(nil))
}
