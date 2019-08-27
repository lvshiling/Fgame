package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	propertytypes "fgame/fgame/game/property/types"
	questtypes "fgame/fgame/game/quest/types"
	"fmt"
)

//日环配置
type UnionRiChangTemplate struct {
	*UnionRiChangTemplateVO
	dailyTimesType   questtypes.QuestDailyType
	rewData          *propertytypes.RewData //奖励属性
	rewItemMap       map[int32]int32        //奖励物品
	doubleRewData    *propertytypes.RewData //双倍奖励
	doubleRewItemMap map[int32]int32        //双倍奖励
	emailItemMap     map[int32]int32        //发邮件
}

func (t *UnionRiChangTemplate) TemplateId() int {
	return t.Id
}

func (t *UnionRiChangTemplate) GetNextId() int32 {
	return t.NextId
}

func (t *UnionRiChangTemplate) GetLevelMin() int32 {
	return t.LevelMin
}

func (t *UnionRiChangTemplate) GetLevelMax() int32 {
	return t.LevelMax
}

func (t *UnionRiChangTemplate) GetTimesMin() int32 {
	return t.Times
}

func (t *UnionRiChangTemplate) GetTimesMax() int32 {
	return t.TimesMax
}

func (t *UnionRiChangTemplate) GetQuestId() int32 {
	return t.QuestId
}

func (t *UnionRiChangTemplate) GetPercent() int32 {
	return t.Percent
}

func (t *UnionRiChangTemplate) GetRewExp() int32 {
	return t.RewExp
}

func (t *UnionRiChangTemplate) GetRewExpPoint() int32 {
	return t.RewExpPoint
}

func (t *UnionRiChangTemplate) GetRewSilver() int32 {
	return t.RewSilver
}

func (t *UnionRiChangTemplate) GetRewBindGold() int32 {
	return t.RewBindGold
}

func (t *UnionRiChangTemplate) GetRewGold() int32 {
	return t.RewGold
}

func (t *UnionRiChangTemplate) GetBossExp() int32 {
	return t.UnionBossExp
}

func (t *UnionRiChangTemplate) GetRewData() *propertytypes.RewData {
	return t.rewData
}

func (t *UnionRiChangTemplate) GetDoubleRewData() *propertytypes.RewData {
	return t.doubleRewData
}

func (t *UnionRiChangTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *UnionRiChangTemplate) GetDoubleRewItemMap() map[int32]int32 {
	return t.doubleRewItemMap
}

func (t *UnionRiChangTemplate) GetDailyTimesType() questtypes.QuestDailyType {
	return t.dailyTimesType
}

func (t *UnionRiChangTemplate) GetEmailItemMap() map[int32]int32 {
	return t.emailItemMap
}

func (t *UnionRiChangTemplate) GetUnionStorageJiFen() int32 {
	return t.UnionStorageJiFen
}

func (t *UnionRiChangTemplate) GetDropId() int32 {
	return 0
}

func (t *UnionRiChangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.emailItemMap = make(map[int32]int32)
	//rew_silver
	err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}

	//rew_bind_gold
	err = validator.MinValidate(float64(t.RewBindGold), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewBindGold", err)
		return
	}

	//rew_gold
	err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}

	//rew_exp
	err = validator.MinValidate(float64(t.RewExp), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}

	//rew_exp_point
	err = validator.MinValidate(float64(t.RewExpPoint), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewExpPoint", err)
		return
	}

	t.rewData = propertytypes.CreateRewData(t.RewExp, t.RewExpPoint, t.RewSilver, t.RewBindGold, t.RewGold)
	t.doubleRewData = propertytypes.CreateRewData(2*t.RewExp, 2*t.RewExpPoint, 2*t.RewSilver, 2*t.RewBindGold, 2*t.RewGold)

	if t.RewSilver != 0 {
		t.emailItemMap[constanttypes.SilverItem] += t.RewSilver
	}

	if t.RewBindGold != 0 {
		t.emailItemMap[constanttypes.BindGoldItem] += t.RewBindGold
	}

	if t.RewGold != 0 {
		t.emailItemMap[constanttypes.GoldItem] += t.RewGold
	}

	if t.RewExp != 0 {
		t.emailItemMap[constanttypes.ExpItem] += t.RewExp
	}

	t.rewItemMap = make(map[int32]int32)
	t.doubleRewItemMap = make(map[int32]int32)
	if t.RewItemId != "" {
		if t.RewItemCount == "" {
			err = fmt.Errorf("[%s] invalid", t.RewItemCount)
			return template.NewTemplateFieldError("RewItemCount", err)
		}

		itemArr, err := utils.SplitAsIntArray(t.RewItemId)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(t.RewItemCount)
		if err != nil {
			return err
		}
		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", t.RewItemId)
			return template.NewTemplateFieldError("RewItemId", err)
		}

		for i := range itemArr {
			t.rewItemMap[itemArr[i]] = numArr[i]
			t.doubleRewItemMap[itemArr[i]] = 2 * numArr[i]
			t.emailItemMap[itemArr[i]] += numArr[i]
		}
	}

	return nil
}

func (t *UnionRiChangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.dailyTimesType = questtypes.QuestDailyType(t.Times)
	if !t.dailyTimesType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Times)
		err = template.NewTemplateFieldError("Times", err)
		return
	}

	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*UnionRiChangTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		dailyTempalte := to.(*UnionRiChangTemplate)
		if dailyTempalte.Times != t.Times {
			err = fmt.Errorf("[%d] invalid", t.Times)
			err = template.NewTemplateFieldError("TimesMin", err)
			return
		}

		if dailyTempalte.TimesMax != t.TimesMax {
			err = fmt.Errorf("[%d] invalid", t.TimesMax)
			err = template.NewTemplateFieldError("TimesMax", err)
			return
		}

		if t.LevelMin != dailyTempalte.LevelMin {
			err = fmt.Errorf("[%d] invalid", t.LevelMin)
			err = template.NewTemplateFieldError("LevelMin", err)
			return
		}

		if t.LevelMax != dailyTempalte.LevelMax {
			err = fmt.Errorf("[%d] invalid", t.LevelMax)
			err = template.NewTemplateFieldError("LevelMax", err)
			return
		}
	}

	//quest_id
	if t.QuestId != 0 {
		to := template.GetTemplateService().Get(int(t.QuestId), (*QuestTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}

		questTemplate := to.(*QuestTemplate)
		if questTemplate.GetQuestType() != questtypes.QuestTypeDailyAlliance {
			err = fmt.Errorf("[%d] invalid", t.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}
	}

	err = validator.MinValidate(float64(t.LevelMin), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("LevelMin", err)
		return
	}

	err = validator.MinValidate(float64(t.LevelMax), float64(t.LevelMin), true)
	if err != nil {
		err = template.NewTemplateFieldError("LevelMax", err)
		return
	}

	err = validator.MinValidate(float64(t.Percent), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("Percent", err)
		return
	}

	//校验奖励物品
	for itemId, num := range t.rewItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", t.RewItemId)
			return template.NewTemplateFieldError("RewItemId", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.RewItemCount)
			return template.NewTemplateFieldError("RewItemCount", err)
		}
	}

	return nil
}

func (t *UnionRiChangTemplate) PatchAfterCheck() {

}

func (t *UnionRiChangTemplate) FileName() string {
	return "tb_union_richang.json"
}

func init() {
	template.Register((*UnionRiChangTemplate)(nil))
}
