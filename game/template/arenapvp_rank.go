package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	"fmt"
)

//3v3排行榜配置
type ArenapvpRankTemplate struct {
	*ArenapvpRankTemplateVO
	rewItemMap     map[int32]int32
	mailRewItemMap map[int32]int32
}

func (t *ArenapvpRankTemplate) TemplateId() int {
	return t.Id
}

func (t *ArenapvpRankTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *ArenapvpRankTemplate) GetMailRewItemMap() map[int32]int32 {
	return t.mailRewItemMap
}

func (t *ArenapvpRankTemplate) PatchAfterCheck() {
	if t.GetSilver > 0 {
		t.mailRewItemMap[constanttypes.SilverItem] = t.GetSilver
	}
	if t.GetGold > 0 {
		t.mailRewItemMap[constanttypes.GoldItem] = t.GetGold
	}
	if t.GetBindGold > 0 {
		t.mailRewItemMap[constanttypes.BindGoldItem] = t.GetBindGold
	}
}

func (t *ArenapvpRankTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewItemMap = make(map[int32]int32)
	t.mailRewItemMap = make(map[int32]int32)
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
			t.mailRewItemMap[itemId] = numArr[index]
		}
	}

	return nil
}

func (t *ArenapvpRankTemplate) Check() (err error) {
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

	//验证 get_gold
	err = validator.MinValidate(float64(t.GetGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GetGold)
		err = template.NewTemplateFieldError("GetGold", err)
		return
	}

	//验证 RewJifen
	err = validator.MinValidate(float64(t.RewJifen), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewJifen)
		err = template.NewTemplateFieldError("RewJifen", err)
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

func (t *ArenapvpRankTemplate) FileName() string {
	return "tb_biwudahui_rank.json"
}

func init() {
	template.Register((*ArenapvpRankTemplate)(nil))
}
