package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

type FameFeedInfo struct {
	ItemId          int32
	ItemNum         int32
	AddFavorableNum int32
	FeedLimit       int32
	DropId          int32
}

//名人配置
type FamousTemplate struct {
	*FamousTemplateVo
	fameInfoList []*FameFeedInfo
}

func (t *FamousTemplate) TemplateId() int {
	return t.Id
}

func (t *FamousTemplate) GetFameFeedInfoList() []*FameFeedInfo {
	return t.fameInfoList
}

func (t *FamousTemplate) GetFameFeedInfo(feedIndex int) *FameFeedInfo {
	if len(t.fameInfoList) < feedIndex+1 {
		return nil
	}
	return t.fameInfoList[feedIndex]
}

func (t *FamousTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//消耗物品
	itemArr, err := utils.SplitAsIntArray(t.UseItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}
	itemNumArr, err := utils.SplitAsIntArray(t.UseItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemCount)
		return template.NewTemplateFieldError("UseItemCount", err)
	}
	favorableNumArr, err := utils.SplitAsIntArray(t.ItemIncreaseFavorable)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemIncreaseFavorable)
		return template.NewTemplateFieldError("ItemIncreaseFavorable", err)
	}
	limitArr, err := utils.SplitAsIntArray(t.UseLimit)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseLimit)
		return template.NewTemplateFieldError("UseLimit", err)
	}
	dropArr, err := utils.SplitAsIntArray(t.DropId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.DropId)
		return template.NewTemplateFieldError("DropId", err)
	}
	if len(itemArr) != len(itemNumArr) {
		err = fmt.Errorf("[%s] invalid", t.UseItemCount)
		return template.NewTemplateFieldError("UseItemId or UseItemCount", err)
	}
	if len(itemArr) != len(favorableNumArr) {
		err = fmt.Errorf("[%s] invalid", t.ItemIncreaseFavorable)
		return template.NewTemplateFieldError("UseItemId or ItemIncreaseFavorable", err)
	}
	if len(itemArr) != len(limitArr) {
		err = fmt.Errorf("[%s] invalid", t.UseLimit)
		return template.NewTemplateFieldError("UseItemId or UseLimit", err)
	}
	if len(itemArr) != len(dropArr) {
		err = fmt.Errorf("[%s] invalid", t.DropId)
		return template.NewTemplateFieldError("UseItemId or dropArr", err)
	}
	if len(itemArr) > 0 {
		//组合数据
		for index, itemId := range itemArr {
			info := &FameFeedInfo{
				ItemId:          itemId,
				ItemNum:         itemNumArr[index],
				AddFavorableNum: favorableNumArr[index],
				FeedLimit:       limitArr[index],
				DropId:          dropArr[index],
			}
			t.fameInfoList = append(t.fameInfoList, info)
		}
	}

	return nil
}

func (t *FamousTemplate) PatchAfterCheck() {
}

func (t *FamousTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 消耗物品
	for _, info := range t.fameInfoList {
		itemId := info.ItemId
		num := info.ItemNum
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("UseItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("UseItemCount", err)
			return
		}
	}
	return nil
}

func (edt *FamousTemplate) FileName() string {
	return "tb_famous.json"
}

func init() {
	template.Register((*FamousTemplate)(nil))
}
