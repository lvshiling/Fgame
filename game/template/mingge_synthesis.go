package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fmt"
)

//命格合成配置
type MingGeSynthesisTemplate struct {
	*MingGeSynthesisTemplateVO
	needItemMap map[int32]int32
}

func (mt *MingGeSynthesisTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MingGeSynthesisTemplate) GetNeedItemMap() map[int32]int32 {
	return mt.needItemMap
}

func (mt *MingGeSynthesisTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	if mt.NeedItemId != "" {
		mt.needItemMap = make(map[int32]int32)

		itemArr, err := utils.SplitAsIntArray(mt.NeedItemId)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(mt.NeedItemCount)
		if err != nil {
			return err
		}

		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", mt.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return err
		}

		for index, itemId := range itemArr {
			mt.needItemMap[itemId] += numArr[index]
		}
	}

	return nil
}

func (mt *MingGeSynthesisTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//lastQualityType := itemtypes.ItemQualityType(-1)
	for itemId, num := range mt.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", mt.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return
		}

		//itemTemplate := to.(*ItemTemplate)

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", mt.NeedItemCount)
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}

		// //红色品质不能再合成了
		// if itemTemplate.GetQualityType() == itemtypes.ItemQualityTypeRed {
		// 	err = fmt.Errorf("[%s] invalid", mt.NeedItemId)
		// 	err = template.NewTemplateFieldError("NeedItemId", err)
		// 	return
		// }

		// //验证品质是否一样
		// if lastQualityType == -1 {
		// 	lastQualityType = itemTemplate.GetQualityType()
		// } else {
		// 	if lastQualityType != itemTemplate.GetQualityType() {
		// 		err = fmt.Errorf("[%s] invalid", mt.NeedItemId)
		// 		err = template.NewTemplateFieldError("NeedItemId", err)
		// 		return
		// 	}
		// }
	}

	if mt.ItemId != 0 {
		to := template.GetTemplateService().Get(int(mt.ItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.ItemId)
			err = template.NewTemplateFieldError("ItemId", err)
			return
		}

		err = validator.MinValidate(float64(mt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	err = validator.MinValidate(float64(mt.NeedSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedSilver)
		err = template.NewTemplateFieldError("NeedSilver", err)
		return
	}

	err = validator.MinValidate(float64(mt.NeedGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedGold)
		err = template.NewTemplateFieldError("NeedGold", err)
		return
	}

	err = validator.MinValidate(float64(mt.NeedBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedBindGold)
		err = template.NewTemplateFieldError("NeedBindGold", err)
		return
	}

	err = validator.RangeValidate(float64(mt.SuccessRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.SuccessRate)
		err = template.NewTemplateFieldError("SuccessRate", err)
		return
	}

	return nil
}

func (mt *MingGeSynthesisTemplate) PatchAfterCheck() {

}

func (mt *MingGeSynthesisTemplate) FileName() string {
	return "tb_mingge_synthesis.json"
}

func init() {
	template.Register((*MingGeSynthesisTemplate)(nil))
}
