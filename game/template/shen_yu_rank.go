package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	"fmt"
)

type ShenYuRankTemplateList []*ShenYuRankTemplate

func (adl ShenYuRankTemplateList) Len() int {
	return len(adl)
}

func (adl ShenYuRankTemplateList) Less(i, j int) bool {
	return adl[i].RankMin < adl[j].RankMin
}

func (adl ShenYuRankTemplateList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

//神域排行榜配置
type ShenYuRankTemplate struct {
	*ShenYuRankTemplateVO
	rewItemMap      map[int32]int32
	rewEmailItemMap map[int32]int32
}

func (t *ShenYuRankTemplate) TemplateId() int {
	return t.Id
}

func (t *ShenYuRankTemplate) GetRewEmailItemMap() map[int32]int32 {
	return t.rewEmailItemMap
}

func (t *ShenYuRankTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *ShenYuRankTemplate) PatchAfterCheck() {
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

func (t *ShenYuRankTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//下一条
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(t.NextId, (*ShenYuRankTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp := to.(*ShenYuRankTemplate)

		if nextTemp.RankMin-t.RankMax != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	}

	t.rewEmailItemMap = make(map[int32]int32)
	t.rewItemMap = make(map[int32]int32)
	itemArr, err := utils.SplitAsIntArray(t.RewItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewItemId)
		err = template.NewTemplateFieldError("RewItemId", err)
		return
	}

	numArr, err := utils.SplitAsIntArray(t.RewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewItemCount)
		err = template.NewTemplateFieldError("RewItemCount", err)
		return
	}

	if len(itemArr) != len(numArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.RewItemId, t.RewItemCount)
		err = template.NewTemplateFieldError("RewItemId or RewItemCount", err)
		return
	}

	for index, itemId := range itemArr {
		t.rewEmailItemMap[itemId] = numArr[index]
		t.rewItemMap[itemId] = numArr[index]
	}

	return nil
}

func (t *ShenYuRankTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

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

func (t *ShenYuRankTemplate) FileName() string {
	return "tb_shenyu_rank.json"
}

func init() {
	template.Register((*ShenYuRankTemplate)(nil))
}
