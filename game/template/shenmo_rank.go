package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//神魔排行榜配置
type ShenMoRankTemplate struct {
	*ShenMoRankTemplateVO
	rewItemMap map[int32]int32
}

func (t *ShenMoRankTemplate) TemplateId() int {
	return t.Id
}

func (t *ShenMoRankTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *ShenMoRankTemplate) PatchAfterCheck() {

}

func (t *ShenMoRankTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.rewItemMap = make(map[int32]int32)
	if t.GetItemId != "" {
		itemArr, err := utils.SplitAsIntArray(t.GetItemId)
		if err != nil {
			return err
		}

		if t.GetItemCount == "" {
			err = fmt.Errorf("[%s] invalid", t.GetItemCount)
			return template.NewTemplateFieldError("GetItemCount", err)
		}

		numArr, err := utils.SplitAsIntArray(t.GetItemCount)
		if err != nil {
			return err
		}

		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", t.GetItemCount)
			return template.NewTemplateFieldError("GetItemCount", err)
		}

		for index, itemId := range itemArr {
			t.rewItemMap[itemId] = numArr[index]
		}
	}

	return nil
}

func (t *ShenMoRankTemplate) Check() (err error) {
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
	err = validator.MinValidate(float64(t.RankMax), float64(t.RankMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RankMax)
		err = template.NewTemplateFieldError("RankMax", err)
		return
	}

	//验证 get_silver
	err = validator.MinValidate(float64(t.GetSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GetSilver)
		err = template.NewTemplateFieldError("GetSilver", err)
		return
	}

	//验证 get_bind_gold
	err = validator.MinValidate(float64(t.GetBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GetBindGold)
		err = template.NewTemplateFieldError("GetBindGold", err)
		return
	}

	for itemId, num := range t.rewItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.GetItemId)
			err = template.NewTemplateFieldError("GetItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.GetItemCount)
			err = template.NewTemplateFieldError("GetItemCount", err)
			return
		}
	}

	return nil
}

func (t *ShenMoRankTemplate) FileName() string {
	return "tb_shenmo_rank.json"
}

func init() {
	template.Register((*ShenMoRankTemplate)(nil))
}
