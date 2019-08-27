package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	hunttypes "fgame/fgame/game/hunt/types"
	"fmt"
	"sort"
)

//寻宝配置
type HuntTemplate struct {
	*HuntTemplateVO
	huntType       hunttypes.HuntType
	dropByTimesMap map[int32]int32 //按次数必定掉落map
	timesDescList  []int           //循环掉落
}

func (t *HuntTemplate) TemplateId() int {
	return t.Id
}

func (t *HuntTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//按次数必定掉落map
	t.dropByTimesMap = make(map[int32]int32)

	mustDropList, err := utils.SplitAsIntArray(t.MustGet)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MustGet)
		return template.NewTemplateFieldError("MustGet", err)
	}
	mustAmountList, err := utils.SplitAsIntArray(t.MustAmount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MustAmount)
		return template.NewTemplateFieldError("MustAmount", err)
	}
	if len(mustDropList) != len(mustAmountList) {
		err = fmt.Errorf("[%s] invalid", t.MustAmount)
		return template.NewTemplateFieldError("MustAmount And MustGet ", err)
	}
	for index, mustAmount := range mustAmountList {
		t.dropByTimesMap[mustAmount] = mustDropList[index]
		t.timesDescList = append(t.timesDescList, int(mustAmount))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(t.timesDescList)))
	return
}

func (t *HuntTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//类型
	t.huntType = hunttypes.HuntType(t.Type)
	if !t.huntType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	// 银两消耗
	err = validator.MinValidate(float64(t.SilverUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("SilverUse", err)
	}
	// 元宝消耗
	err = validator.MinValidate(float64(t.GoldUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("GoldUse", err)
	}
	// 绑元消耗
	err = validator.MinValidate(float64(t.BindGoldUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BindGoldUse", err)
	}

	if t.UseItemCount > 0 {
		to := template.GetTemplateService().Get(int(t.UseItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemId)
			return template.NewTemplateFieldError("UseItemId", err)
		}
	}

	return nil
}

func (t *HuntTemplate) PatchAfterCheck() {
}

func (t *HuntTemplate) FileName() string {
	return "tb_xunbao.json"
}

func (t *HuntTemplate) GetHuntType() hunttypes.HuntType {
	return t.huntType
}

func (t *HuntTemplate) GetRewDropMap() map[int32]int32 {
	return t.dropByTimesMap
}

func (t *HuntTemplate) GetDropTimesDescList() []int {
	return t.timesDescList
}

func init() {
	template.Register((*HuntTemplate)(nil))
}
