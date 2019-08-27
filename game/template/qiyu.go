package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	questtypes "fgame/fgame/game/quest/types"
	"fmt"
)

type AcceptCondition struct {
	Level int32
	Zhuan int32
	Fei   int32
}

//奇遇配置
type QiYuTemplate struct {
	*QiYuTemplateVO
	rewItemMap      map[int32]int32          //奖励物品
	rewEmailItemMap map[int32]int32          //邮件奖励物品
	questMap        map[int32]*QuestTemplate //奇遇任务map
	conditionList   []*AcceptCondition       //接取条件
}

func (t *QiYuTemplate) TemplateId() int {
	return t.Id
}

func (t *QiYuTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *QiYuTemplate) GetRewEmailItemMap() map[int32]int32 {
	return t.rewEmailItemMap
}

func (t *QiYuTemplate) GetEndTime(now int64) int64 {
	return int64(t.GuoQiTime) + now
}

func (t *QiYuTemplate) GetConditionList() []*AcceptCondition {
	return t.conditionList
}

// 最大满足条件
func (t *QiYuTemplate) GetMatchCondition(level, zhuan, fei int32) *AcceptCondition {
	var temCondition *AcceptCondition
	for _, condition := range t.conditionList {
		if level < condition.Level {
			break
		}
		if zhuan < condition.Zhuan {
			break
		}
		if fei < condition.Fei {
			break
		}
		temCondition = condition
	}

	return temCondition
}

func (t *QiYuTemplate) GetQuestMap() map[int32]*QuestTemplate {
	return t.questMap
}

func (t *QiYuTemplate) PatchAfterCheck() {
	if t.RewExp > 0 {
		t.rewEmailItemMap[constanttypes.ExpItem] += t.RewExp
	}
	if t.RewGold > 0 {
		t.rewEmailItemMap[constanttypes.GoldItem] += t.RewGold
	}
	if t.RewBindGold > 0 {
		t.rewEmailItemMap[constanttypes.BindGoldItem] += t.RewBindGold
	}
	if t.RewSilver > 0 {
		t.rewEmailItemMap[constanttypes.SilverItem] += t.RewSilver
	}
}

func (t *QiYuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//奖励物品id
	if t.RewItemId != "" {
		t.rewItemMap = make(map[int32]int32)
		t.rewEmailItemMap = make(map[int32]int32)
		itemIdList, err := utils.SplitAsIntArray(t.RewItemId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.RewItemId)
			return template.NewTemplateFieldError("RewItemId", err)
		}
		if len(itemIdList) == 0 {
			err = fmt.Errorf("[%s] invalid", t.RewItemId)
			return template.NewTemplateFieldError("RewItemId", err)
		}
		itemCountList, err := utils.SplitAsIntArray(t.RewItemCount)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.RewItemCount)
			return template.NewTemplateFieldError("RewItemCount", err)
		}

		if len(itemCountList) != len(itemIdList) {
			err = fmt.Errorf("[%s] invalid", t.RewItemCount)
			return template.NewTemplateFieldError("RewItemCount", err)
		}

		for index, itemId := range itemIdList {
			t.rewItemMap[itemId] = itemCountList[index]
			t.rewEmailItemMap[itemId] = itemCountList[index]
		}
	}

	//验证
	t.questMap = make(map[int32]*QuestTemplate)
	if t.QuestId != "" {
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
	}

	//条件
	levelList, err := utils.SplitAsIntArray(t.Level)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}
	zhuanList, err := utils.SplitAsIntArray(t.ZhuanSheng)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ZhuanSheng)
		return template.NewTemplateFieldError("ZhuanSheng", err)
	}
	feiList, err := utils.SplitAsIntArray(t.FeiSheng)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FeiSheng)
		return template.NewTemplateFieldError("FeiSheng", err)
	}
	if len(levelList) != len(zhuanList) && len(levelList) != len(feiList) {
		err = fmt.Errorf("[%s],[%s],[%s] invalid", t.Level, t.ZhuanSheng, t.FeiSheng)
		return template.NewTemplateFieldError("Level ZhuanSheng FeiSheng ", err)
	}
	for index, level := range levelList {
		condition := &AcceptCondition{
			Level: level,
			Zhuan: zhuanList[index],
			Fei:   feiList[index],
		}
		t.conditionList = append(t.conditionList, condition)
	}

	return nil
}

func (t *QiYuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// //验证
	// err = validator.MinValidate(float64(t.Level), float64(1), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.Level)
	// 	err = template.NewTemplateFieldError("Level", err)
	// 	return
	// }

	// //验证
	// err = validator.MinValidate(float64(t.ZhuanSheng), float64(0), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.ZhuanSheng)
	// 	err = template.NewTemplateFieldError("ZhuanSheng", err)
	// 	return
	// }

	// //验证
	// err = validator.MinValidate(float64(t.FeiSheng), float64(0), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.FeiSheng)
	// 	err = template.NewTemplateFieldError("FeiSheng", err)
	// 	return
	// }
	//验证
	err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewSilver)
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}
	//验证
	err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewGold)
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}
	//验证
	err = validator.MinValidate(float64(t.RewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewBindGold)
		err = template.NewTemplateFieldError("RewBindGold", err)
		return
	}
	//验证
	err = validator.MinValidate(float64(t.RewExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewExp)
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}
	//验证
	err = validator.MinValidate(float64(t.RewExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewExpPoint)
		err = template.NewTemplateFieldError("RewExpPoint", err)
		return
	}
	//验证
	err = validator.MinValidate(float64(t.GuoQiTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GuoQiTime)
		err = template.NewTemplateFieldError("GuoQiTime", err)
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

	//校验任务
	for _, questTemplate := range t.questMap {
		if questTemplate.GetQuestType() != questtypes.QuestTypeQiYu {
			err = fmt.Errorf("[%s] type invalid", t.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}
	}

	return nil
}

func (t *QiYuTemplate) FileName() string {
	return "tb_qiyu_quest.json"
}

func init() {
	template.Register((*QiYuTemplate)(nil))
}
