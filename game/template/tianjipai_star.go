package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//天机牌星数配置
type TianJiPaiStarTemplate struct {
	*TianJiPaiStarTemplateVO
	rewData    *propertytypes.RewData //奖励属性
	rewItemMap map[int32]int32        //奖励物品
}

func (tt *TianJiPaiStarTemplate) TemplateId() int {
	return tt.Id
}

func (tt *TianJiPaiStarTemplate) GetRewData() *propertytypes.RewData {
	return tt.rewData
}

func (tt *TianJiPaiStarTemplate) GetRewItemMap() map[int32]int32 {
	return tt.rewItemMap
}

func (tt *TianJiPaiStarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//award_silver
	err = validator.MinValidate(float64(tt.AwardSilver), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("AwardSilver", err)
		return
	}

	//award_gold
	err = validator.MinValidate(float64(tt.AwardGold), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("AwardGold", err)
		return
	}

	//award_bindgold
	err = validator.MinValidate(float64(tt.AwardBindGold), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("AwardBindGold", err)
		return
	}

	if tt.AwardSilver > 0 || tt.AwardBindGold > 0 || tt.AwardGold > 0 {
		tt.rewData = propertytypes.CreateRewData(0, 0, tt.AwardSilver, tt.AwardGold, tt.AwardBindGold)
	}

	tt.rewItemMap = make(map[int32]int32)
	if tt.AwardItemId != "" {
		if tt.AwardItemIdCount == "" {
			err = fmt.Errorf("[%s] invalid", tt.AwardItemIdCount)
			return template.NewTemplateFieldError("AwardItemIdCount", err)
		}

		itemArr, err := utils.SplitAsIntArray(tt.AwardItemId)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(tt.AwardItemIdCount)
		if err != nil {
			return err
		}
		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", tt.AwardItemIdCount)
			return template.NewTemplateFieldError("AwardItemIdCount", err)
		}

		for i := 0; i < len(itemArr); i++ {
			tt.rewItemMap[itemArr[i]] = numArr[i]
		}

	}

	return nil
}

func (tt *TianJiPaiStarTemplate) PatchAfterCheck() {

}

func (tt *TianJiPaiStarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//need_star
	err = validator.MinValidate(float64(tt.NeedStar), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("StarMax", err)
		return
	}

	//next_id
	if tt.NextId != 0 {
		diff := tt.NextId - int32(tt.Id)
		to := template.GetTemplateService().Get(int(tt.NextId), (*TianJiPaiStarTemplate)(nil))
		if to == nil || diff != 1 {
			err = fmt.Errorf("[%d] invalid", tt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		tempTemplate := to.(*TianJiPaiStarTemplate)
		if tempTemplate.NeedStar < tt.NeedStar {
			err = fmt.Errorf("[%d] invalid", tt.NeedStar)
			return template.NewTemplateFieldError("NeedStar", err)
		}
	}

	//校验奖励物品
	for itemId, num := range tt.rewItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", tt.AwardItemId)
			return template.NewTemplateFieldError("AwardItemId", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", tt.AwardItemIdCount)
			return template.NewTemplateFieldError("AwardItemIdCount", err)
		}
	}

	return nil
}

func (tt *TianJiPaiStarTemplate) FileName() string {
	return "tb_tianjipai_star.json"
}

func init() {
	template.Register((*TianJiPaiStarTemplate)(nil))
}
