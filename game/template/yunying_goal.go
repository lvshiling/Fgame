package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	questtypes "fgame/fgame/game/quest/types"
	"fmt"
)

//目标配置
type YunYingGoalTemplate struct {
	*YunYingGoalTemplateVO
	rewItemMap      map[int32]int32          //奖励物品
	rewEmailItemMap map[int32]int32          //邮件奖励物品
	questMap        map[int32]*QuestTemplate //目标任务map
}

func (t *YunYingGoalTemplate) TemplateId() int {
	return t.Id
}

func (t *YunYingGoalTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *YunYingGoalTemplate) GetRewEmailItemMap() map[int32]int32 {
	return t.rewEmailItemMap
}

func (t *YunYingGoalTemplate) GetQuestMap() map[int32]*QuestTemplate {
	return t.questMap
}

func (t *YunYingGoalTemplate) PatchAfterCheck() {
	// if t.RewExp > 0 {
	// 	t.rewEmailItemMap[constanttypes.ExpItem] += t.RewExp
	// }
	// if t.RewGold > 0 {
	// 	t.rewEmailItemMap[constanttypes.GoldItem] += t.RewGold
	// }
	// if t.RewBindGold > 0 {
	// 	t.rewEmailItemMap[constanttypes.BindGoldItem] += t.RewBindGold
	// }
	// if t.RewSilver > 0 {
	// 	t.rewEmailItemMap[constanttypes.SilverItem] += t.RewSilver
	// }
}

func (t *YunYingGoalTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//奖励物品id
	t.rewItemMap = make(map[int32]int32)
	t.rewEmailItemMap = make(map[int32]int32)
	// itemIdList, err := utils.SplitAsIntArray(t.RewItemId)
	// if err != nil {
	// 	err = fmt.Errorf("[%s] invalid", t.RewItemId)
	// 	return template.NewTemplateFieldError("RewItemId", err)
	// }
	// itemCountList, err := utils.SplitAsIntArray(t.RewItemCount)
	// if err != nil {
	// 	err = fmt.Errorf("[%s] invalid", t.RewItemCount)
	// 	return template.NewTemplateFieldError("RewItemCount", err)
	// }
	// if len(itemCountList) != len(itemIdList) {
	// 	err = fmt.Errorf("[%s] invalid", t.RewItemCount)
	// 	return template.NewTemplateFieldError("RewItemCount", err)
	// }

	// for index, itemId := range itemIdList {
	// 	t.rewItemMap[itemId] = itemCountList[index]
	// 	t.rewEmailItemMap[itemId] = itemCountList[index]
	// }

	//验证
	t.questMap = make(map[int32]*QuestTemplate)
	questIdList, err := utils.SplitAsIntArray(t.QuestId)
	for _, questId := range questIdList {
		to := template.GetTemplateService().Get(int(questId), (*QuestTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] nil invalid", t.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}

		questTemplate, ok := to.(*QuestTemplate)
		if !ok {
			return fmt.Errorf("QuestId [%s] invalid", t.QuestId)
		}
		t.questMap[questId] = questTemplate
	}

	return nil
}

func (t *YunYingGoalTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// //验证
	// err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.RewSilver)
	// 	err = template.NewTemplateFieldError("RewSilver", err)
	// 	return
	// }
	// //验证
	// err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.RewGold)
	// 	err = template.NewTemplateFieldError("RewGold", err)
	// 	return
	// }
	// //验证
	// err = validator.MinValidate(float64(t.RewBindGold), float64(0), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.RewBindGold)
	// 	err = template.NewTemplateFieldError("RewBindGold", err)
	// 	return
	// }
	// //验证
	// err = validator.MinValidate(float64(t.RewExp), float64(0), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.RewExp)
	// 	err = template.NewTemplateFieldError("RewExp", err)
	// 	return
	// }
	// //验证
	// err = validator.MinValidate(float64(t.RewExpPoint), float64(0), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.RewExpPoint)
	// 	err = template.NewTemplateFieldError("RewExpPoint", err)
	// 	return
	// }

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

	//校验任务
	for _, questTemplate := range t.questMap {
		if questTemplate.GetQuestType() != questtypes.QuestTypeYunYingGoal {
			err = fmt.Errorf("[%s] type invalid", t.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}
	}

	return nil
}

func (t *YunYingGoalTemplate) FileName() string {
	return "tb_yunying_mubiao_quest.json"
}

func init() {
	template.Register((*YunYingGoalTemplate)(nil))
}
