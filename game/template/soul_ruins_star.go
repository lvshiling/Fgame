package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/soulruins/types"
	"fmt"
)

//帝魂遗迹星级奖励配置
type SoulRuinsStarTemplate struct {
	*SoulRuinsStarTemplateVO
	typ        types.SoulRuinsType
	rewItemMap map[int32]int32        //奖励物品
	rewData    *propertytypes.RewData //奖励属性
}

func (srst *SoulRuinsStarTemplate) TemplateId() int {
	return srst.Id
}

func (srst *SoulRuinsStarTemplate) GetType() types.SoulRuinsType {
	return srst.typ
}

func (srst *SoulRuinsStarTemplate) GetRewItemMap() map[int32]int32 {
	return srst.rewItemMap
}

func (srst *SoulRuinsStarTemplate) GetRewData() *propertytypes.RewData {
	return srst.rewData
}

func (srst *SoulRuinsStarTemplate) PatchAfterCheck() {

}

func (srst *SoulRuinsStarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(srst.FileName(), srst.TemplateId(), err)
			return
		}
	}()

	//type
	srst.typ = types.SoulRuinsType(srst.Type)
	if !srst.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", srst.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	srst.rewItemMap = make(map[int32]int32)
	//rew_item_id1
	if srst.RewItemId1 != 0 {
		srst.rewItemMap[srst.RewItemId1] = srst.RewItemCount1
	}
	//rew_item_id2
	if srst.RewItemId2 != 0 {
		_, ok := srst.rewItemMap[srst.RewItemId2]
		if ok {
			err = fmt.Errorf("[%d] invalid", srst.RewItemId2)
			err = template.NewTemplateFieldError("RewItemId2", err)
			return
		}
		srst.rewItemMap[srst.RewItemId2] = srst.RewItemCount2
	}
	//rew_item_id3
	if srst.RewItemId3 != 0 {
		_, ok := srst.rewItemMap[srst.RewItemId3]
		if ok {
			err = fmt.Errorf("[%d] invalid", srst.RewItemId3)
			err = template.NewTemplateFieldError("RewItemId3", err)
			return
		}
		srst.rewItemMap[srst.RewItemId3] = srst.RewItemCount3
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(srst.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srst.RewGold)
		return template.NewTemplateFieldError("RewGold", err)
	}

	//验证 rew_yinliang
	err = validator.MinValidate(float64(srst.RewYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srst.RewYinliang)
		return template.NewTemplateFieldError("RewYinliang", err)
	}

	//验证 rew_exp
	err = validator.MinValidate(float64(srst.RewExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srst.RewExp)
		return template.NewTemplateFieldError("RewExp", err)
	}

	//验证 rew_uplev
	err = validator.MinValidate(float64(srst.RewUplev), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srst.RewUplev)
		return template.NewTemplateFieldError("RewUplev", err)
	}

	if srst.RewYinliang > 0 || srst.RewGold > 0 || srst.RewUplev > 0 || srst.RewExp > 0 {
		srst.rewData = propertytypes.CreateRewData(srst.RewExp, srst.RewUplev, srst.RewYinliang, srst.RewGold, 0)
	}

	return nil
}

func (srst *SoulRuinsStarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(srst.FileName(), srst.TemplateId(), err)
			return
		}
	}()

	//验证 rew_item
	for itemId, num := range srst.rewItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("rew_item_id", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", srst.RewItemCount3)
			return template.NewTemplateFieldError("rew_item_count", err)
		}
	}

	//验证 need_star
	err = validator.MinValidate(float64(srst.NeedStar), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srst.NeedStar)
		return template.NewTemplateFieldError("NeedStar", err)
	}

	//验证 chapter
	err = validator.MinValidate(float64(srst.Chapter), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srst.Chapter)
		return template.NewTemplateFieldError("Chapter", err)
	}

	return nil
}

func (srst *SoulRuinsStarTemplate) FileName() string {
	return "tb_soul_ruins_star.json"
}

func init() {
	template.Register((*SoulRuinsStarTemplate)(nil))
}
