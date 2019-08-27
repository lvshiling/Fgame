package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/fourgod/types"
	itemtypes "fgame/fgame/game/item/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//四神遗迹宝箱配置配置
type FourGodBoxTemplate struct {
	*FourGodBoxTemplateVO
	boxType    types.FourGodBoxType   //宝箱类型
	rewData    *propertytypes.RewData //奖励属性
	dropIdList []int32
}

func (fgt *FourGodBoxTemplate) TemplateId() int {
	return fgt.Id
}

func (fgt *FourGodBoxTemplate) GetRewData() *propertytypes.RewData {
	return fgt.rewData
}

func (fgt *FourGodBoxTemplate) GetDropIdList() []int32 {
	return fgt.dropIdList
}

func (fgt *FourGodBoxTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fgt.FileName(), fgt.TemplateId(), err)
			return
		}
	}()

	fgt.boxType = types.FourGodBoxType(fgt.Type)
	if !fgt.boxType.Valid() {
		err = fmt.Errorf("[%d] invalid", fgt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//dropId
	if fgt.DropId == "" {
		err = fmt.Errorf("[%s] invalid", fgt.DropId)
		return template.NewTemplateFieldError("DropId", err)
	}
	dropIdArr, err := utils.SplitAsIntArray(fgt.DropId)
	if err != nil {
		return
	}
	fgt.dropIdList = make([]int32, 0, 4)
	for _, dropId := range dropIdArr {
		fgt.dropIdList = append(fgt.dropIdList, dropId)
	}

	//award_silver
	err = validator.MinValidate(float64(fgt.AwardSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fgt.AwardSilver)
		return template.NewTemplateFieldError("AwardSilver", err)
	}

	//award_gold
	err = validator.MinValidate(float64(fgt.AwardGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fgt.AwardGold)
		return template.NewTemplateFieldError("AwardGold", err)
	}

	//award_bindgold
	err = validator.MinValidate(float64(fgt.AwardBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fgt.AwardBindGold)
		return template.NewTemplateFieldError("AwardBindGold", err)
	}

	//award_exp
	err = validator.MinValidate(float64(fgt.AwardExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fgt.AwardExp)
		return template.NewTemplateFieldError("AwardExp", err)
	}

	//award_exp_point
	err = validator.MinValidate(float64(fgt.AwardExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fgt.AwardExpPoint)
		return template.NewTemplateFieldError("AwardExpPoint", err)
	}

	if fgt.AwardSilver > 0 || fgt.AwardGold > 0 || fgt.AwardBindGold > 0 || fgt.AwardExp > 0 || fgt.AwardExpPoint > 0 {
		fgt.rewData = propertytypes.CreateRewData(fgt.AwardExp, fgt.AwardExpPoint, fgt.AwardSilver, fgt.AwardGold, fgt.AwardBindGold)
	}

	return nil
}

func (fgt *FourGodBoxTemplate) PatchAfterCheck() {

}

func (fgt *FourGodBoxTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fgt.FileName(), fgt.TemplateId(), err)
			return
		}
	}()

	//use_item_id
	to := template.GetTemplateService().Get(int(fgt.UseItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", fgt.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}
	itemTemplate := to.(*ItemTemplate)
	if itemTemplate.GetItemType() != itemtypes.ItemTypeAutoUseRes ||
		itemTemplate.GetItemSubType() != itemtypes.ItemAutoUseResSubTypeKey {
		err = fmt.Errorf("[%d] invalid", fgt.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}

	//UseItemCount
	err = validator.MinValidate(float64(fgt.UseItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fgt.UseItemCount)
		return template.NewTemplateFieldError("UseItemCount", err)
	}

	//key_min
	err = validator.MinValidate(float64(fgt.KeyMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fgt.KeyMin)
		return template.NewTemplateFieldError("KeyMin", err)
	}

	//key_max
	err = validator.MinValidate(float64(fgt.KeyMax), float64(fgt.KeyMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fgt.KeyMax)
		return template.NewTemplateFieldError("KeyMin", err)
	}

	//next_id
	// if fgt.NextId != 0 {
	// 	diff := fgt.NextId - int32(fgt.Id)
	// 	to := template.GetTemplateService().Get(int(fgt.NextId), (*FourGodBoxTemplate)(nil))
	// 	if to == nil || diff != 1 {
	// 		err = fmt.Errorf("[%d] invalid", fgt.NextId)
	// 		return template.NewTemplateFieldError("NextId", err)
	// 	}
	// 	diffKey := to.(*FourGodBoxTemplate).KeyMin - fgt.KeyMax
	// 	if diffKey != 1 {
	// 		err = fmt.Errorf("[%d] invalid", fgt.KeyMax)
	// 		return template.NewTemplateFieldError("KeyMax", err)
	// 	}
	// }

	//biology_id
	biologyTempTemplate := template.GetTemplateService().Get(int(fgt.BiologyId), (*BiologyTemplate)(nil))
	if biologyTempTemplate == nil {
		err = fmt.Errorf("[%d] invalid", fgt.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}

	// biologyTemplate := biologyTempTemplate.(*BiologyTemplate)
	// if biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeFourGodCollect {
	// 	err = fmt.Errorf("[%d] invalid", fgt.BiologyId)
	// 	return template.NewTemplateFieldError("BiologyId", err)
	// }

	return nil
}

func (fgt *FourGodBoxTemplate) FileName() string {
	return "tb_four_box.json"
}

func init() {
	template.Register((*FourGodBoxTemplate)(nil))
}
