package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	foundtypes "fgame/fgame/game/found/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fmt"
)

//资源找回配置
type FoundTemplate struct {
	*FoundTemplateVO
	resType       foundtypes.FoundResourceType
	funcOpenType  funcopentypes.FuncOpenType
	freeFoundData foundtypes.FoundData //普通找回
	goldFoundData foundtypes.FoundData //完美找回
}

func (t *FoundTemplate) TemplateId() int {
	return t.Id
}

func (t *FoundTemplate) GetFreeFoundData() foundtypes.FoundData {
	return t.freeFoundData
}

func (t *FoundTemplate) GetGoldFoundData() foundtypes.FoundData {
	return t.goldFoundData
}

func (t *FoundTemplate) GetFoundData(typ foundtypes.FoundType) foundtypes.FoundData {
	if typ == foundtypes.FoundTypeFree {
		return t.freeFoundData
	} else {
		return t.goldFoundData
	}
}

func (t *FoundTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证: 找回物品ID,逗号隔开
	//验证：找回物品数量
	t.goldFoundData = foundtypes.CreateFoundData()
	foundItemIdList, err := utils.SplitAsIntArray(t.FoundItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FoundItemId)
		return template.NewTemplateFieldError("FoundItemId", err)
	}
	foundItemAmountList, err := utils.SplitAsIntArray(t.FoundItemAmount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FoundItemAmount)
		return template.NewTemplateFieldError("FoundItemAmount", err)
	}
	if len(foundItemAmountList) != len(foundItemAmountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.FoundItemId, t.FoundItemAmount)
		err = template.NewTemplateFieldError("FoundItemId or FoundItemAmount", err)
		return err
	}
	if len(foundItemIdList) > 0 {
		//组合数据
		for index, itemId := range foundItemIdList {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.FoundItemId)
				return template.NewTemplateFieldError("FoundItemId", err)
			}

			err = validator.MinValidate(float64(foundItemAmountList[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("FoundItemAmount", err)
			}

			t.goldFoundData.FoundItemMap[itemId] = foundItemAmountList[index]
		}
	}

	//验证: 找回物品ID,逗号隔开
	//验证：找回物品数量
	t.freeFoundData = foundtypes.CreateFoundData()
	foundItemId2List, err := utils.SplitAsIntArray(t.FoundItemId2)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FoundItemId2)
		return template.NewTemplateFieldError("FoundItemId2", err)
	}
	foundItemAmount2List, err := utils.SplitAsIntArray(t.FoundItemAmount2)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FoundItemAmount2)
		return template.NewTemplateFieldError("FoundItemAmount2", err)
	}
	if len(foundItemAmount2List) != len(foundItemAmount2List) {
		err = fmt.Errorf("[%s][%s] invalid", t.FoundItemId2, t.FoundItemAmount2)
		err = template.NewTemplateFieldError("FoundItemId2 or FoundItemAmount2", err)
		return err
	}
	if len(foundItemId2List) > 0 {
		//组合数据
		for index, itemId := range foundItemId2List {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.FoundItemId2)
				return template.NewTemplateFieldError("FoundItemId2", err)
			}

			err = validator.MinValidate(float64(foundItemAmount2List[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("FoundItemAmount2", err)
			}

			t.freeFoundData.FoundItemMap[itemId] = foundItemAmount2List[index]
		}
	}

	return nil
}

func (t *FoundTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//资源类型
	typ := foundtypes.FoundResourceType(t.Type)
	if !typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}
	t.resType = typ

	//最低限制
	err = validator.MinValidate(float64(t.LevelMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMin)
		err = template.NewTemplateFieldError("LevelMin", err)
		return
	}
	//最高限制
	err = validator.MinValidate(float64(t.LevelMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMin)
		err = template.NewTemplateFieldError("LevelMin", err)
		return
	}

	//验证 found_using
	err = validator.MinValidate(float64(t.FoundUsing), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundUsing)
		err = template.NewTemplateFieldError("FoundUsing", err)
		return
	}

	//验证 found_silver
	err = validator.MinValidate(float64(t.FoundSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundSilver)
		err = template.NewTemplateFieldError("FoundSilver", err)
		return
	}
	t.goldFoundData.FoundSilver = t.FoundSilver

	//验证 found_gold
	err = validator.MinValidate(float64(t.FoundGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundGold)
		err = template.NewTemplateFieldError("FoundGold", err)
		return
	}
	t.goldFoundData.FoundGold = t.FoundGold

	//验证 found_bindgold
	err = validator.MinValidate(float64(t.FoundBindgold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundBindgold)
		err = template.NewTemplateFieldError("FoundBindgold", err)
		return
	}
	t.goldFoundData.FoundBindgold = t.FoundBindgold

	//验证 found_exp
	err = validator.MinValidate(float64(t.FoundExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundExp)
		err = template.NewTemplateFieldError("FoundExp", err)
		return
	}
	t.goldFoundData.FoundExp = t.FoundExp

	//验证 found_exp_point
	err = validator.MinValidate(float64(t.FoundExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundExpPoint)
		err = template.NewTemplateFieldError("FoundExpPoint", err)
		return
	}
	t.goldFoundData.FoundExpPoint = t.FoundExpPoint

	//验证 FoundUsingSilver
	err = validator.MinValidate(float64(t.FoundUsingSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundUsingSilver)
		err = template.NewTemplateFieldError("FoundUsingSilver", err)
		return
	}

	//验证 found_silver
	err = validator.MinValidate(float64(t.FoundSilver2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundSilver2)
		err = template.NewTemplateFieldError("FoundSilver2", err)
		return
	}
	t.freeFoundData.FoundSilver = t.FoundSilver2

	//验证 found_gold
	err = validator.MinValidate(float64(t.FoundGold2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundGold2)
		err = template.NewTemplateFieldError("FoundGold2", err)
		return
	}
	t.freeFoundData.FoundGold = t.FoundGold2

	//验证 found_bindgold
	err = validator.MinValidate(float64(t.FoundBindgold2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundBindgold2)
		err = template.NewTemplateFieldError("FoundBindgold2", err)
		return
	}
	t.freeFoundData.FoundBindgold = t.FoundBindgold2

	//验证 found_exp
	err = validator.MinValidate(float64(t.FoundExp2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundExp2)
		err = template.NewTemplateFieldError("FoundExp2", err)
		return
	}
	t.freeFoundData.FoundExp = t.FoundExp2

	//验证 found_exp_point
	err = validator.MinValidate(float64(t.FoundExpPoint2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FoundExpPoint2)
		err = template.NewTemplateFieldError("FoundExpPoint2", err)
		return
	}
	t.freeFoundData.FoundExpPoint = t.FoundExpPoint2

	//验证 open_id
	temObj := template.GetTemplateService().Get(int(t.OpenId), (*ModuleOpenedTemplate)(nil))
	if temObj == nil {
		err = fmt.Errorf("[%d] invalid", t.OpenId)
		err = template.NewTemplateFieldError("OpenId", err)
		return
	}
	funcTem := temObj.(*ModuleOpenedTemplate)
	t.funcOpenType = funcTem.GetFuncOpenType()

	return nil
}

func (t *FoundTemplate) PatchAfterCheck() {

}

func (t *FoundTemplate) FileName() string {
	return "tb_found.json"
}

func (t *FoundTemplate) GetResType() foundtypes.FoundResourceType {
	return t.resType
}

func (t *FoundTemplate) GetFuncOpenType() funcopentypes.FuncOpenType {
	return t.funcOpenType
}

func init() {
	template.Register((*FoundTemplate)(nil))
}
