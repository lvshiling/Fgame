package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	marrytype "fgame/fgame/game/marry/types"
	"fmt"
)

type MarryTuiSongTemplate struct {
	*MarryTuiSongTemplateVO
	PreGiftType   marrytype.MarryPreGiftType
	returnItemMap map[int32]int32
}

func (mmt *MarryTuiSongTemplate) TemplateId() int {
	return mmt.Id
}

func (mmt *MarryTuiSongTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mmt.FileName(), mmt.TemplateId(), err)
			return
		}
	}()

	preType := marrytype.MarryPreGiftType(mmt.Type)
	if !preType.Valid() {
		err = fmt.Errorf("[%d] invalid", mmt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}
	mmt.PreGiftType = preType

	returnZhuHeArray, err := utils.SplitAsIntArray(mmt.ZhuHeItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", mmt.ZhuHeItem)
		err = template.NewTemplateFieldError("zhuhe_item", err)
		return
	}

	returnZhuHeCountArray, err := utils.SplitAsIntArray(mmt.ZhuheItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", mmt.ZhuheItemCount)
		err = template.NewTemplateFieldError("zhuhe_item", err)
		return
	}

	if len(returnZhuHeArray) != len(returnZhuHeCountArray) {
		err = fmt.Errorf("[%s] [%s] invalid,count not the same", mmt.ZhuheItemCount, mmt.ZhuHeItem)
		err = template.NewTemplateFieldError("zhuhe_item", err)
		return
	}

	mmt.returnItemMap = make(map[int32]int32)
	for i := 0; i < len(returnZhuHeArray); i++ {
		mmt.returnItemMap[returnZhuHeArray[i]] = returnZhuHeCountArray[i]
	}

	return nil
}

func (mmt *MarryTuiSongTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mmt.FileName(), mmt.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(mmt.NeedGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mmt.NeedGold)
		err = template.NewTemplateFieldError("NeedGold", err)
		return
	}

	err = validator.MinValidate(float64(mmt.RewardExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mmt.RewardExp)
		err = template.NewTemplateFieldError("RewardExp", err)
		return
	}

	err = validator.MinValidate(float64(mmt.RewardExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mmt.RewardExpPoint)
		err = template.NewTemplateFieldError("RewardExpPoint", err)
		return
	}

	return nil
}
func (mmt *MarryTuiSongTemplate) PatchAfterCheck() {

}
func (mmt *MarryTuiSongTemplate) FileName() string {
	return "tb_marry_tuisong.json"
}

func (mmt *MarryTuiSongTemplate) GetReturnItemMap() map[int32]int32 {
	return mmt.returnItemMap
}

func (mmt *MarryTuiSongTemplate) GetRewardExp() int64 {
	return int64(mmt.RewardExp)
}

func (mmt *MarryTuiSongTemplate) GetRewardExpPoint() int64 {
	return int64(mmt.RewardExpPoint)
}

func init() {
	template.Register((*MarryTuiSongTemplate)(nil))
}
