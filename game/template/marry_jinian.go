package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	marrytype "fgame/fgame/game/marry/types"
	"fmt"
)

//婚车移动配置
type MarryJiNianTemplate struct {
	*MarryJiNianTemplateVO
	jinianMap    map[int32]int32 //赠送物品
	marrySubType marrytype.MarryBanquetSubTypeWed
}

//接口开始
func (mmt *MarryJiNianTemplate) TemplateId() int {
	return mmt.Id
}

func (mmt *MarryJiNianTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mmt.FileName(), mmt.TemplateId(), err)
			return
		}
	}()

	preType := marrytype.MarryBanquetSubTypeWed(mmt.Type)
	if !preType.Valid() {
		err = fmt.Errorf("[%d] invalid", mmt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}
	mmt.marrySubType = preType

	returnZhuHeArray, err := utils.SplitAsIntArray(mmt.ZhuheItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", mmt.ZhuheItem)
		err = template.NewTemplateFieldError("zhuhe_item", err)
		return
	}

	returnZhuHeCountArray, err := utils.SplitAsIntArray(mmt.ZhuheItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", mmt.ZhuheItemCount)
		err = template.NewTemplateFieldError("ZhuheItemCount", err)
		return
	}

	if len(returnZhuHeArray) != len(returnZhuHeCountArray) {
		err = fmt.Errorf("[%s] [%s] invalid,count not the same", mmt.ZhuheItemCount, mmt.ZhuheItem)
		err = template.NewTemplateFieldError("zhuhe_item", err)
		return
	}

	mmt.jinianMap = make(map[int32]int32)
	for i := 0; i < len(returnZhuHeArray); i++ {
		mmt.jinianMap[returnZhuHeArray[i]] = returnZhuHeCountArray[i]
	}

	return nil
}

func (mmt *MarryJiNianTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mmt.FileName(), mmt.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(mmt.TitleId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mmt.TitleId)
		err = template.NewTemplateFieldError("TitleId", err)
		return
	}

	err = validator.MinValidate(float64(mmt.NeedNum), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mmt.NeedNum)
		err = template.NewTemplateFieldError("NeedNum", err)
		return
	}

	return nil
}
func (mmt *MarryJiNianTemplate) PatchAfterCheck() {

}
func (mmt *MarryJiNianTemplate) FileName() string {
	return "tb_marry_jinian.json"
}

func (mmt *MarryJiNianTemplate) GetItemMap() map[int32]int32 {
	return mmt.jinianMap
}

func (mmt *MarryJiNianTemplate) GetMarrySubType() marrytype.MarryBanquetSubTypeWed {
	return mmt.marrySubType
}

func init() {
	template.Register((*MarryJiNianTemplate)(nil))
}
