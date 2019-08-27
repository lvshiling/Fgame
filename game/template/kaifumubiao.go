package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//运营活动次数配置
type KaiFuMuBiaoTemplate struct {
	*KaiFuMuBiaoTemplateVO
	rewItemMap                  map[int32]int32                     //奖励物品
	kaiFuMuBiaoQuestTemplateMap map[int32]*KaiFuMuBiaoQuestTemplate //开服目标任务map
	kaiFuMuBiaoQuestTemplate    *KaiFuMuBiaoQuestTemplate           //开服目标任务
}

func (t *KaiFuMuBiaoTemplate) TemplateId() int {
	return t.Id
}

func (t *KaiFuMuBiaoTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *KaiFuMuBiaoTemplate) GetQuestMap() map[int32]*KaiFuMuBiaoQuestTemplate {
	return t.kaiFuMuBiaoQuestTemplateMap
}

func (t *KaiFuMuBiaoTemplate) PatchAfterCheck() {
	if t.kaiFuMuBiaoQuestTemplate != nil {
		t.kaiFuMuBiaoQuestTemplateMap = make(map[int32]*KaiFuMuBiaoQuestTemplate)
		for tempTemplate := t.kaiFuMuBiaoQuestTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextKaiFuMuBiaoQuestTemplate {
			questId := tempTemplate.Quest
			t.kaiFuMuBiaoQuestTemplateMap[questId] = tempTemplate
		}
	}
}

func (t *KaiFuMuBiaoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//奖励物品id
	if t.ItemId != "" {
		t.rewItemMap = make(map[int32]int32)
		itemIdList, err := utils.SplitAsIntArray(t.ItemId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.ItemId)
			return template.NewTemplateFieldError("ItemId", err)
		}
		if len(itemIdList) == 0 {
			err = fmt.Errorf("[%s] invalid", t.ItemId)
			return template.NewTemplateFieldError("ItemId", err)
		}
		itemCountList, err := utils.SplitAsIntArray(t.ItemCount)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}

		if len(itemCountList) != len(itemIdList) {
			err = fmt.Errorf("[%s] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}

		for index, itemId := range itemIdList {
			t.rewItemMap[itemId] = itemCountList[index]
		}
	}

	//验证
	if t.GroupBeginId != 0 {
		to := template.GetTemplateService().Get(int(t.GroupBeginId), (*KaiFuMuBiaoQuestTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.GroupBeginId)
			return template.NewTemplateFieldError("GroupBeginId", err)
		}

		kaiFuMuBiaoQuestTemplate, ok := to.(*KaiFuMuBiaoQuestTemplate)
		if !ok {
			return fmt.Errorf("GroupBeginId [%d] invalid", t.GroupBeginId)
		}
		t.kaiFuMuBiaoQuestTemplate = kaiFuMuBiaoQuestTemplate
	}

	return nil
}

func (t *KaiFuMuBiaoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证
	err = validator.MinValidate(float64(t.KaiFuTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KaiFuTime)
		err = template.NewTemplateFieldError("KaiFuTime", err)
		return
	}

	//验证 FinishQuestCount
	err = validator.MinValidate(float64(t.FinishQuestCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FinishQuestCount)
		err = template.NewTemplateFieldError("FinishQuestCount", err)
		return
	}

	//验证 RewardSilver
	err = validator.MinValidate(float64(t.RewardSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewardSilver)
		err = template.NewTemplateFieldError("RewardSilver", err)
		return
	}

	//验证  物品
	for itemId, num := range t.rewItemMap {
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

	return nil
}

func (t *KaiFuMuBiaoTemplate) FileName() string {
	return "tb_kaifumubiao.json"
}

func init() {
	template.Register((*KaiFuMuBiaoTemplate)(nil))
}
