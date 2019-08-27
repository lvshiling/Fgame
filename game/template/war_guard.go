package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	"fmt"
)

//配置
type WarGuardTemplate struct {
	*WarGuardTemplateVO
	mapTemplate *MapTemplate
	itemMap     map[int32]int32
}

func (t *WarGuardTemplate) TemplateId() int {
	return t.Id
}

func (t *WarGuardTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *WarGuardTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	tempMapTemplate := template.GetTemplateService().Get(int(t.Map), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.Map)
		return template.NewTemplateFieldError("Map", err)
	}
	t.mapTemplate = tempMapTemplate.(*MapTemplate)

	t.itemMap = make(map[int32]int32)
	itemIdList, err := coreutils.SplitAsIntArray(t.NeedItemId)
	if err != nil {
		return template.NewTemplateFieldError("NeedItemId", err)
	}
	itemCountList, err := coreutils.SplitAsIntArray(t.NeedItemCount)
	if err != nil {
		return template.NewTemplateFieldError("NeedItemCount", err)
	}
	if len(itemIdList) != len(itemCountList) {
		return template.NewTemplateFieldError("NeedItem", err)
	}
	for i, itemId := range itemIdList {
		tempItemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			return template.NewTemplateFieldError("NeedItemId", err)
		}
		c := itemCountList[i]
		err = validator.MinValidate(float64(c), float64(0), false)
		if err != nil {
			return template.NewTemplateFieldError("NeedItemCount", fmt.Errorf("[%s] invalid", t.NeedItemCount))
		}
		t.itemMap[itemId] = c
	}
	return nil
}

func (t *WarGuardTemplate) PatchAfterCheck() {

}

func (t *WarGuardTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//TODO 验证scene_id
	//银两
	err = validator.MinValidate(float64(t.NeedSilver), float64(0), true)
	if err != nil {

		return template.NewTemplateFieldError("NeedSilver", err)
	}
	//元宝
	err = validator.MinValidate(float64(t.NeedGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("NeedGold", err)
	}
	//绑元
	err = validator.MinValidate(float64(t.NeedBindGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("NeedBindGold", err)
	}

	return nil
}

func (tt *WarGuardTemplate) FileName() string {
	return "tb_war_guard.json"
}

func init() {
	template.Register((*WarGuardTemplate)(nil))
}
