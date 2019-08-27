package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//好友添加配置
type FriendAddTemplate struct {
	*FriendAddTemplateVO
	rewItemMap map[int32]int32
}

func (t *FriendAddTemplate) TemplateId() int {
	return t.Id
}

func (t *FriendAddTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *FriendAddTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//奖励物品
	t.rewItemMap = make(map[int32]int32)
	rewItemIdList, err := utils.SplitAsIntArray(t.RewardItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewardItemId)
		return template.NewTemplateFieldError("RewardItemId", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.RewardItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewardItemCount)
		return template.NewTemplateFieldError("RewardItemCount", err)
	}
	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.RewardItemId, t.RewardItemCount)
		return template.NewTemplateFieldError("RewardItemId or RewardItemCount", err)
	}
	if len(rewItemIdList) > 0 {
		//组合数据
		for index, itemId := range rewItemIdList {
			_, ok := t.rewItemMap[itemId]
			if ok {
				t.rewItemMap[itemId] += rewItemCountList[index]
			} else {
				t.rewItemMap[itemId] = rewItemCountList[index]
			}
		}
	}

	return nil
}

func (t *FriendAddTemplate) PatchAfterCheck() {
}

func (t *FriendAddTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//好友数量
	err = validator.MinValidate(float64(t.Num), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Num)
		return template.NewTemplateFieldError("Num", err)
	}

	// 银两
	err = validator.MinValidate(float64(t.RewardSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewardSilver)
		return template.NewTemplateFieldError("RewardSilver", err)
	}

	// 经验
	err = validator.MinValidate(float64(t.RewardExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewardExp)
		return template.NewTemplateFieldError("RewardExp", err)
	}

	// 经验点
	err = validator.MinValidate(float64(t.RewardExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewardExpPoint)
		return template.NewTemplateFieldError("RewardExpPoint", err)
	}

	//奖励物品
	for itemId, num := range t.rewItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("RewardItemId", fmt.Errorf("[%d] invalid", itemId))
		}
		if err = validator.MinValidate(float64(num), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("RewardItemCount", err)
			return
		}
	}

	return nil
}

func (edt *FriendAddTemplate) FileName() string {
	return "tb_haoyou_tianjia.json"
}

func init() {
	template.Register((*FriendAddTemplate)(nil))
}
