package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/marry/types"
	"fmt"
)

//婚车移动配置
type MarryGiftTemplate struct {
	*MarryGiftTemplateVO
	giftType types.MarryGiftType //贺礼类型
}

func (mgt *MarryGiftTemplate) TemplateId() int {
	return mgt.Id
}

func (mgt *MarryGiftTemplate) GetGiftType() types.MarryGiftType {
	return mgt.giftType
}

func (mgt *MarryGiftTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mgt.FileName(), mgt.TemplateId(), err)
			return
		}
	}()

	mgt.giftType = types.MarryGiftType(mgt.Type)
	if !mgt.giftType.Valid() {
		err = fmt.Errorf("[%d] invalid", mgt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	switch mgt.giftType {
	case types.MarryGiftTypeItem:
		{
			to := template.GetTemplateService().Get(int(mgt.UseItemId), (*ItemTemplate)(nil))
			if to == nil {
				err = fmt.Errorf("[%d] invalid", mgt.UseItemId)
				return template.NewTemplateFieldError("UseItemId", err)
			}

			err = validator.MinValidate(float64(mgt.UseItemAmount), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", mgt.UseItemAmount)
				err = template.NewTemplateFieldError("UseItemAmount", err)
				return
			}
			break
		}
	case types.MarryGiftTypeSilver:
		{
			err = validator.MinValidate(float64(mgt.UseSilver), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", mgt.UseSilver)
				err = template.NewTemplateFieldError("UseSilver", err)
				return
			}
			break
		}
	}

	return nil
}

func (mgt *MarryGiftTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mgt.FileName(), mgt.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(mgt.BuffAmount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mgt.BuffAmount)
		err = template.NewTemplateFieldError("BuffAmount", err)
		return
	}

	return nil
}
func (mgt *MarryGiftTemplate) PatchAfterCheck() {

}
func (mgt *MarryGiftTemplate) FileName() string {
	return "tb_marry_gift.json"
}

func init() {
	template.Register((*MarryGiftTemplate)(nil))
}
