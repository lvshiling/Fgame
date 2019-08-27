package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	groupcollectenum "fgame/fgame/game/welfare/group/collect/enum"
	"fmt"
)

//摸金奖励配置
type ChouJiangPokerTemplate struct {
	*ChouJiangPokerTemplateVO
	pokerType groupcollectenum.PokerType
	// rewItemMap map[int32]int32
}

func (t *ChouJiangPokerTemplate) TemplateId() int {
	return t.Id
}

func (t *ChouJiangPokerTemplate) GetPokerType() groupcollectenum.PokerType {
	return t.pokerType
}

func (t *ChouJiangPokerTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// t.rewItemMap = make(map[int32]int32)
	// //验证 rew_item_id
	// rewItemIdList, err := utils.SplitAsIntArray(t.RawItem)
	// if err != nil {
	// 	err = fmt.Errorf("[%s] split invalid", t.RawItem)
	// 	return template.NewTemplateFieldError("RawItem", err)
	// }
	// rewItemCountList, err := utils.SplitAsIntArray(t.RawCount)
	// if err != nil {
	// 	err = fmt.Errorf("[%s] invalid", t.RawCount)
	// 	return template.NewTemplateFieldError("RawCount", err)
	// }
	// if len(rewItemIdList) != len(rewItemCountList) {
	// 	err = fmt.Errorf("[%s] invalid", t.RawCount)
	// 	return template.NewTemplateFieldError("RawCount", err)
	// }
	// if len(rewItemIdList) > 0 {
	// 	//组合数据
	// 	for index, itemId := range rewItemIdList {
	// 		_, ok := t.rewItemMap[itemId]
	// 		if !ok {
	// 			t.rewItemMap[itemId] = rewItemCountList[index]
	// 		} else {
	// 			t.rewItemMap[itemId] += rewItemCountList[index]
	// 		}
	// 	}
	// }

	return nil
}

func (t *ChouJiangPokerTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 校验 卡牌类型
	t.pokerType = groupcollectenum.PokerType(t.Type)
	if !t.pokerType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 rew_silver
	err = validator.MinValidate(float64(t.RawYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RawYinliang)
		err = template.NewTemplateFieldError("RawYinliang", err)
		return
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(t.RawGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RawGold)
		err = template.NewTemplateFieldError("RawGold", err)
		return
	}

	//验证 rew_bind_gold
	err = validator.MinValidate(float64(t.RawBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RawBindGold)
		err = template.NewTemplateFieldError("RawBindGold", err)
		return
	}

	// // 验证物品
	// for itemId, num := range t.rewItemMap {
	// 	itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
	// 	if itemTmpObj == nil {
	// 		err = fmt.Errorf("[%s] nil invalid", itemId)
	// 		return template.NewTemplateFieldError("RewItemId", err)
	// 	}

	// 	err = validator.MinValidate(float64(num), float64(1), true)
	// 	if err != nil {
	// 		return template.NewTemplateFieldError("RewItemCount", err)
	// 	}
	// }

	return nil
}

func (t *ChouJiangPokerTemplate) PatchAfterCheck() {

}

func (t *ChouJiangPokerTemplate) FileName() string {
	return "tb_mojin.json"
}

func init() {
	template.Register((*ChouJiangPokerTemplate)(nil))
}
