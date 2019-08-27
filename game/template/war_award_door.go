package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//配置
type WarAwardDoorTemplate struct {
	*WarAwardDoorTemplateVO
	itemMap      map[int32]int32
	emailItemMap map[int32]int32
	rewData      *propertytypes.RewData //奖励属性
}

func (t *WarAwardDoorTemplate) TemplateId() int {
	return t.Id
}

func (t *WarAwardDoorTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *WarAwardDoorTemplate) GetEmailItemMap() map[int32]int32 {
	return t.emailItemMap
}

func (t *WarAwardDoorTemplate) GetRewData() *propertytypes.RewData {
	return t.rewData
}

func (t *WarAwardDoorTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.itemMap = make(map[int32]int32)
	t.emailItemMap = make(map[int32]int32)
	itemIdList, err := coreutils.SplitAsIntArray(t.WarDoorItemId)
	if err != nil {
		return template.NewTemplateFieldError("WarDoorItemId", err)
	}
	itemCountList, err := coreutils.SplitAsIntArray(t.WarDoorItemIdCount)
	if err != nil {
		return template.NewTemplateFieldError("WarDoorItemIdCount", err)
	}
	if len(itemIdList) != len(itemCountList) {
		return template.NewTemplateFieldError("WarDoorItem", err)
	}
	for i, itemId := range itemIdList {
		tempItemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			return template.NewTemplateFieldError("WarDoorItemId", err)
		}
		c := itemCountList[i]
		err = validator.MinValidate(float64(c), float64(0), false)
		if err != nil {
			return template.NewTemplateFieldError("WarDoorItemIdCount", fmt.Errorf("[%s] invalid", t.WarDoorItemIdCount))
		}
		t.itemMap[itemId] = c
		t.emailItemMap[itemId] = c
	}

	t.rewData = propertytypes.CreateRewData(0, 0, t.WarDoorSilver, t.WarDoorGold, t.WarDoorBindgold)

	return nil
}

func (t *WarAwardDoorTemplate) PatchAfterCheck() {
	if t.WarDoorSilver > 0 {
		t.emailItemMap[constanttypes.SilverItem] = t.WarDoorSilver
	}
	if t.WarDoorGold > 0 {
		t.emailItemMap[constanttypes.GoldItem] = t.WarDoorGold
	}
	if t.WarDoorBindgold > 0 {
		t.emailItemMap[constanttypes.BindGoldItem] = t.WarDoorBindgold
	}
}

func (t *WarAwardDoorTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//TODO 验证scene_id
	//银两
	err = validator.MinValidate(float64(t.WarDoorSilver), float64(0), true)
	if err != nil {

		return template.NewTemplateFieldError("WarDoorSilver", err)
	}
	//元宝
	err = validator.MinValidate(float64(t.WarDoorGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("WarDoorGold", err)
	}
	//绑元
	err = validator.MinValidate(float64(t.WarDoorBindgold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("WarDoorBindgold", err)
	}

	return nil
}

func (tt *WarAwardDoorTemplate) FileName() string {
	return "tb_war_award_door.json"
}

func init() {
	template.Register((*WarAwardDoorTemplate)(nil))
}
