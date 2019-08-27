package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/lingtongdev/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
	"strconv"
)

//灵身配置
type LingTongShenFaTemplate struct {
	*LingTongShenFaTemplateVO
	typ                       types.LingTongDevType               //灵身类型
	magicParamMap             map[int32]string                    //幻化条件
	magicParamXUMap           map[int32]int32                     //幻化条件1,2
	magicParamIMap            map[int32]int32                     //幻化条件3消耗物品
	itemMap                   map[int32]int32                     //进阶物品
	shenFaUpstartTemplateMap  map[int32]LingTongDevUpstarTemplate //灵身皮肤升星map
	shenFaUpstartTemplate     LingTongDevUpstarTemplate           //灵身皮肤升星
	battlePropertyMap         map[propertytypes.BattlePropertyType]int64
	lingTongBattlePropertyMap map[propertytypes.BattlePropertyType]int64
	next                      LingTongDevTemplate
}

func (t *LingTongShenFaTemplate) TemplateId() int {
	return t.Id
}

func (t *LingTongShenFaTemplate) GetNextId() int32 {
	return t.NextId
}

func (t *LingTongShenFaTemplate) GetUpdateWfb() int32 {
	return t.UpdateWfb
}

func (t *LingTongShenFaTemplate) GetAddMin() int32 {
	return t.AddMin
}

func (t *LingTongShenFaTemplate) GetAddMax() int32 {
	return t.AddMax
}

func (t *LingTongShenFaTemplate) GetTimesMin() int32 {
	return t.TimesMin
}

func (t *LingTongShenFaTemplate) GetTimesMax() int32 {
	return t.TimesMax
}

func (t *LingTongShenFaTemplate) GetZhuFuMax() int32 {
	return t.ZhufuMax
}

func (t *LingTongShenFaTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *LingTongShenFaTemplate) GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.lingTongBattlePropertyMap
}

func (t *LingTongShenFaTemplate) GetType() types.LingTongDevType {
	return t.typ
}

func (t *LingTongShenFaTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *LingTongShenFaTemplate) GetMagicParamIMap() map[int32]int32 {
	return t.magicParamIMap
}

func (t *LingTongShenFaTemplate) GetMagicParamXUMap() map[int32]int32 {
	return t.magicParamXUMap
}

func (t *LingTongShenFaTemplate) GetIsClear() bool {
	return t.IsClear != 0
}

func (t *LingTongShenFaTemplate) GetUpstarBeginId() int32 {
	return t.LingTongShenFaUpstarBeginId
}

func (t *LingTongShenFaTemplate) GetGold() int64 {
	return t.UseMoney
}

func (t *LingTongShenFaTemplate) GetBindGold() int64 {
	return 0
}

func (t *LingTongShenFaTemplate) GetSilver() int64 {
	return t.UseYinliang
}

func (t *LingTongShenFaTemplate) GetNumber() int32 {
	return t.Number
}

func (t *LingTongShenFaTemplate) GetShiDanLimit() int32 {
	return t.ShidanLimit
}

func (t *LingTongShenFaTemplate) GetCulDanLimit() int32 {
	return 0
}

func (t *LingTongShenFaTemplate) GetNext() LingTongDevTemplate {
	return t.next
}

func (t *LingTongShenFaTemplate) GetClassType() types.LingTongDevSysType {
	return types.LingTongDevSysTypeLingShen
}

func (t *LingTongShenFaTemplate) GetName() string {
	return t.Name
}

func (t *LingTongShenFaTemplate) GetLingTongDevUpstarByLevel(level int32) LingTongDevUpstarTemplate {
	if v, ok := t.shenFaUpstartTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *LingTongShenFaTemplate) GetLingTongDevPeiYangByLevel(level int32) LingTongDevPeiYangTemplate {
	return nil
}
func (t *LingTongShenFaTemplate) GetLingTongDevTongLingByLevel(level int32) LingTongDevTongLingTemplate {
	return nil
}

func (t *LingTongShenFaTemplate) PatchAfterCheck() {
	if t.shenFaUpstartTemplate != nil {
		t.shenFaUpstartTemplateMap = make(map[int32]LingTongDevUpstarTemplate)
		//赋值shenFaUpstartTemplateMap
		for tempTemplate := t.shenFaUpstartTemplate; tempTemplate != nil; tempTemplate = tempTemplate.GetNext() {
			level := tempTemplate.GetLevel()
			t.shenFaUpstartTemplateMap[level] = tempTemplate
		}
	}
}

func (t *LingTongShenFaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.typ = types.LingTongDevType(t.Type)
	if !t.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//幻化条件
	t.magicParamMap = make(map[int32]string)
	t.magicParamXUMap = make(map[int32]int32)
	t.magicParamIMap = make(map[int32]int32)

	t.magicParamMap[t.MagicConditionType1] = t.MagicConditionParameter1
	t.magicParamMap[t.MagicConditionType2] = t.MagicConditionParameter2
	t.magicParamMap[t.MagicConditionType3] = t.MagicConditionParameter3

	for condType, condParam := range t.magicParamMap {
		cType := types.LingTongDevUCondType(condType)
		if !cType.Valid() {
			err = fmt.Errorf("[%d] invalid", condType)
			return template.NewTemplateFieldError("magic_condition_type", err)
		}
		switch cType {
		case types.LingTongDevUCondTypeX,
			types.LingTongDevUCondTypeU:
			num, err := strconv.ParseInt(condParam, 10, 32)
			if err != nil {
				return err
			}
			t.magicParamXUMap[condType] = int32(num)
			break
		case types.LingTongDevUCondTypeI:
			itemArr, err := utils.SplitAsIntArray(condParam)
			if err != nil {
				return err
			}
			if len(itemArr) != 2 {
				err = fmt.Errorf("[%s] invalid", condParam)
				return template.NewTemplateFieldError("magic_condition_parameter", err)
			}
			t.magicParamIMap[itemArr[0]] = itemArr[1]
			break
		default:
			break
		}
	}
	//幻化条件1、2
	for condType, condParam := range t.magicParamXUMap {
		if condType == int32(types.LingTongDevUCondTypeX) {
			err = validator.MinValidate(float64(condParam), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", condParam)
				err = template.NewTemplateFieldError("MagicConditionParameter", err)
				return
			}
		}
	}
	//幻化条件3
	for item, num := range t.magicParamIMap {
		itemTemplate := template.GetTemplateService().Get(int(item), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("item[%d] invalid", item)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
			return
		}
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("num[%d] invalid", num)
			return template.NewTemplateFieldError("MagicConditionParameter", err)
		}
	}

	//验证 UseItem
	t.itemMap = make(map[int32]int32)
	if t.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(t.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}

		//验证 ItemCount
		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}

		t.itemMap[t.UseItem] = t.ItemCount
	}

	//验证 wing_upgrade_begin_id
	if t.LingTongShenFaUpstarBeginId != 0 {
		to := template.GetTemplateService().Get(int(t.LingTongShenFaUpstarBeginId), (*LingTongShenFaUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.LingTongShenFaUpstarBeginId)
			return template.NewTemplateFieldError("LingTongShenFaUpstarBeginId", err)
		}

		shenFaUpstartTemplate, ok := to.(*LingTongShenFaUpstarTemplate)
		if !ok {
			return fmt.Errorf("LingTongShenFaUpgradeBeginId [%d] invalid", t.LingTongShenFaUpstarBeginId)
		}

		t.shenFaUpstartTemplate = shenFaUpstartTemplate
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

func (t *LingTongShenFaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证Number
	err = validator.MinValidate(float64(t.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*LingTongShenFaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.next = to.(*LingTongShenFaTemplate)
	}

	//验证 update_wfb
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 use_money
	err = validator.MinValidate(float64(t.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(t.TimesMin), float64(0), true, float64(t.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(t.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(t.AddMin), float64(0), true, float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(t.AddMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 UseYinliang
	err = validator.MinValidate(float64(t.UseYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseYinliang)
		err = template.NewTemplateFieldError("UseYinliang", err)
		return
	}

	//验证 ShidanLimit
	err = validator.MinValidate(float64(t.ShidanLimit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ShidanLimit)
		err = template.NewTemplateFieldError("ShidanLimit", err)
		return
	}

	err = validator.MinValidate(float64(t.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}

func (t *LingTongShenFaTemplate) FileName() string {
	return "tb_lingtong_shenfa.json"
}

func init() {
	template.Register((*LingTongShenFaTemplate)(nil))
}
