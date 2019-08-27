package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	alliancetypes "fgame/fgame/game/alliance/types"
	constanttypes "fgame/fgame/game/constant/types"
	"fmt"
)

//配置
type WarAwardTemplate struct {
	*WarAwardTemplateVO
	itemMap    map[int32]int32
	rewardType alliancetypes.AllianceSceneRewardType
}

func (t *WarAwardTemplate) TemplateId() int {
	return t.Id
}

func (t *WarAwardTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *WarAwardTemplate) GetRewardType() alliancetypes.AllianceSceneRewardType {
	return t.rewardType
}

func (t *WarAwardTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.rewardType = alliancetypes.AllianceSceneRewardType(t.Type)

	t.itemMap = make(map[int32]int32)
	itemIdList, err := coreutils.SplitAsIntArray(t.WarAwardItemId)
	if err != nil {
		return template.NewTemplateFieldError("WarAwardItemId", err)
	}
	itemCountList, err := coreutils.SplitAsIntArray(t.WarAwardItemIdCount)
	if err != nil {
		return template.NewTemplateFieldError("WarAwardItemIdCount", err)
	}
	if len(itemIdList) != len(itemCountList) {
		return template.NewTemplateFieldError("WarAwardItem", err)
	}
	for i, itemId := range itemIdList {
		tempItemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			return template.NewTemplateFieldError("WarAwardItemId", err)
		}
		c := itemCountList[i]
		err = validator.MinValidate(float64(c), float64(0), false)
		if err != nil {
			return template.NewTemplateFieldError("WarDoorItemIdWarAwardItemIdCountCount", fmt.Errorf("[%s] invalid", t.WarAwardItemIdCount))
		}
		t.itemMap[itemId] = c
	}

	return nil
}

func (t *WarAwardTemplate) PatchAfterCheck() {
	if t.WarAwardSilver > 0 {
		t.addItem(constanttypes.SilverItem, t.WarAwardSilver)
	}
	if t.WarAwardGold > 0 {
		t.addItem(constanttypes.GoldItem, t.WarAwardGold)

	}
	if t.WarAwardBindgold > 0 {
		t.addItem(constanttypes.BindGoldItem, t.WarAwardBindgold)
	}
}

func (t *WarAwardTemplate) addItem(itemId int32, num int32) {
	currentNum := t.itemMap[itemId]
	t.itemMap[itemId] = currentNum + num
}

func (t *WarAwardTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if !t.rewardType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	//银两
	err = validator.MinValidate(float64(t.WarAwardSilver), float64(0), true)
	if err != nil {

		return template.NewTemplateFieldError("WarAwardSilver", err)
	}

	//元宝
	err = validator.MinValidate(float64(t.WarAwardGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("WarAwardGold", err)
	}
	//绑元
	err = validator.MinValidate(float64(t.WarAwardBindgold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("WarAwardBindgold", err)
	}

	return nil
}

func (tt *WarAwardTemplate) FileName() string {
	return "tb_war_award.json"
}

func init() {
	template.Register((*WarAwardTemplate)(nil))
}
