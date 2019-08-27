package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	"fmt"
)

//龙宫探宝排行榜配置
type LongGongRankTemplate struct {
	*LongGongRankTemplateVO
	rewEmailItemMap map[int32]int32
	nextTemp        *LongGongRankTemplate //下一条
}

func (t *LongGongRankTemplate) TemplateId() int {
	return t.Id
}

func (t *LongGongRankTemplate) GetNextTemplate() *LongGongRankTemplate {
	return t.nextTemp
}

func (t *LongGongRankTemplate) GetRewEmailItemMap() map[int32]int32 {
	return t.rewEmailItemMap
}

func (t *LongGongRankTemplate) PatchAfterCheck() {
	if t.RewSilver != 0 {
		t.rewEmailItemMap[int32(constanttypes.SilverItem)] = t.RewSilver
	}
	if t.RewBindGold != 0 {
		t.rewEmailItemMap[int32(constanttypes.BindGoldItem)] = t.RewBindGold
	}
	if t.RewGold != 0 {
		t.rewEmailItemMap[int32(constanttypes.GoldItem)] = t.RewGold
	}
	if t.RewExp != 0 {
		t.rewEmailItemMap[int32(constanttypes.ExpItem)] = t.RewExp
	}
}

func (t *LongGongRankTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//下一条
	if t.NextId != 0 {
		if t.NextId-t.Id != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		to := template.GetTemplateService().Get(t.NextId, (*LongGongRankTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*LongGongRankTemplate)
	}

	t.rewEmailItemMap = make(map[int32]int32)
	if t.RewItemId != "" {
		itemArr, err := utils.SplitAsIntArray(t.RewItemId)
		if err != nil {
			return err
		}

		if t.RewItemCount == "" {
			err = fmt.Errorf("[%s] invalid", t.RewItemCount)
			return template.NewTemplateFieldError("RewItemCount", err)
		}

		numArr, err := utils.SplitAsIntArray(t.RewItemCount)
		if err != nil {
			return err
		}

		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", t.RewItemCount)
			return template.NewTemplateFieldError("RewItemCount", err)
		}

		for index, itemId := range itemArr {
			t.rewEmailItemMap[itemId] = numArr[index]
		}
	}

	return nil
}

func (t *LongGongRankTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.nextTemp != nil {
		//区间连续校验
		if t.nextTemp.RankMin-t.RankMax != 1 {
			err = fmt.Errorf("[%d] , [%d] invalid", t.nextTemp.RankMin, t.RankMax)
			return template.NewTemplateFieldError("next RankMin, RankMax", err)
		}
	}

	//验证 rank_min
	err = validator.MinValidate(float64(t.RankMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RankMin)
		err = template.NewTemplateFieldError("RankMin", err)
		return
	}

	//验证 rank_max
	err = validator.RangeValidate(float64(t.RankMax), float64(t.RankMin), true, common.MAX_RANK, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RankMax)
		err = template.NewTemplateFieldError("RankMax", err)
		return
	}

	//验证 get_exp
	err = validator.MinValidate(float64(t.RewExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewExp)
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}

	//验证 get_silver
	err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewSilver)
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}

	//验证 get_bind_gold
	err = validator.MinValidate(float64(t.RewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewBindGold)
		err = template.NewTemplateFieldError("RewBindGold", err)
		return
	}

	//验证 get_gold
	err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewGold)
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}

	for itemId, num := range t.rewEmailItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.RewItemId)
			err = template.NewTemplateFieldError("RewItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.RewItemCount)
			err = template.NewTemplateFieldError("RewItemCount", err)
			return
		}
	}

	return nil
}

func (t *LongGongRankTemplate) FileName() string {
	return "tb_longgong_rank.json"
}

func init() {
	template.Register((*LongGongRankTemplate)(nil))
}
