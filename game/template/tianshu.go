package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	tianshutypes "fgame/fgame/game/tianshu/types"
	"fmt"
)

//天书配置
type TianShuTemplate struct {
	*TianShuTemplateVO
	tianshuType   tianshutypes.TianShuType
	nextTemp      *TianShuTemplate
	battleAttrMap map[propertytypes.BattlePropertyType]int64
	needItemMap   map[int32]int32
	giftItemMap   map[int32]int32
}

func (t *TianShuTemplate) TemplateId() int {
	return t.Id
}

func (t *TianShuTemplate) GetTianShuType() tianshutypes.TianShuType {
	return t.tianshuType
}

func (t *TianShuTemplate) GetBattleAttrMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *TianShuTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *TianShuTemplate) GetGiftItemMap() map[int32]int32 {
	return t.giftItemMap
}

func (t *TianShuTemplate) GetNextTemplate() *TianShuTemplate {
	return t.nextTemp
}

func (t *TianShuTemplate) PatchAfterCheck() {
}

func (t *TianShuTemplate) Patch() (err error) {
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

	//免费礼包物品
	t.giftItemMap = make(map[int32]int32)
	freeItemIdList, err := utils.SplitAsIntArray(t.FreeGiftId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FreeGiftId)
		return template.NewTemplateFieldError("FreeGiftId", err)
	}
	freeItemCountList, err := utils.SplitAsIntArray(t.FreeGiftCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FreeGiftCount)
		return template.NewTemplateFieldError("FreeGiftCount", err)
	}
	if len(freeItemIdList) != len(freeItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.FreeGiftId, t.FreeGiftCount)
		return template.NewTemplateFieldError("FreeGiftId or FreeGiftCount", err)
	}
	if len(freeItemIdList) > 0 {
		//组合数据
		for index, itemId := range freeItemIdList {
			_, ok := t.giftItemMap[itemId]
			if ok {
				t.giftItemMap[itemId] += freeItemCountList[index]
			} else {
				t.giftItemMap[itemId] = freeItemCountList[index]
			}
		}
	}

	//升级物品
	t.needItemMap = make(map[int32]int32)
	needItemIdList, err := utils.SplitAsIntArray(t.LevelItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LevelItem)
		return template.NewTemplateFieldError("LevelItem", err)
	}
	needItemCountList, err := utils.SplitAsIntArray(t.LevelItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LevelItemCount)
		return template.NewTemplateFieldError("LevelItemCount", err)
	}
	if len(needItemIdList) != len(needItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.LevelItem, t.LevelItemCount)
		return template.NewTemplateFieldError("LevelItem or LevelItemCount", err)
	}
	if len(needItemIdList) > 0 {
		//组合数据
		for index, itemId := range needItemIdList {
			_, ok := t.needItemMap[itemId]
			if ok {
				t.needItemMap[itemId] += needItemCountList[index]
			} else {
				t.needItemMap[itemId] = needItemCountList[index]
			}
		}
	}

	return nil
}

func (t *TianShuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*TianShuTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTemp = to.(*TianShuTemplate)
	}

	//验证 天书类型
	t.tianshuType = tianshutypes.TianShuType(t.Type)
	if !t.tianshuType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 等级
	err = validator.MinValidate(float64(t.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}
	//验证 元宝激活条件
	err = validator.MinValidate(float64(t.NeedGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedGold)
		err = template.NewTemplateFieldError("NeedGold", err)
		return
	}

	//验证 银两
	err = validator.MinValidate(float64(t.FreeGiftSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FreeGiftSilver)
		err = template.NewTemplateFieldError("FreeGiftSilver", err)
		return
	}

	//验证 绑元
	err = validator.MinValidate(float64(t.FreeGiftBindgold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FreeGiftBindgold)
		err = template.NewTemplateFieldError("FreeGiftBindgold", err)
		return
	}

	//验证 元宝
	err = validator.MinValidate(float64(t.FreeGiftGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FreeGiftGold)
		err = template.NewTemplateFieldError("FreeGiftGold", err)
		return
	}

	// 特权率
	err = validator.MinValidate(float64(t.Tequan), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Tequan)
		err = template.NewTemplateFieldError("Tequan", err)
		return
	}

	//验证 免费礼包物品
	for itemId, num := range t.giftItemMap {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("FreeGiftId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("FreeGiftCount", err)
			return
		}
	}
	//验证 免费礼包物品
	for itemId, num := range t.needItemMap {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("LevelItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("LevelItemCount", err)
			return
		}
	}

	return nil
}

func (t *TianShuTemplate) FileName() string {
	return "tb_tianshu.json"
}

func init() {
	template.Register((*TianShuTemplate)(nil))
}
