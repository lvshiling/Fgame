package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	arenatypes "fgame/fgame/game/arena/types"
	"fmt"
)

type ArenaTemplate struct {
	*ArenaTemplateVO
	arenaType    arenatypes.ArenaType
	itemMap      map[int32]int32
	extraItemMap map[int32]int32
}

func (t *ArenaTemplate) GetArenaType() arenatypes.ArenaType {
	return t.arenaType
}

func (t *ArenaTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *ArenaTemplate) GetExtraItemMap() map[int32]int32 {
	return t.extraItemMap
}

func (t *ArenaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.arenaType = arenatypes.ArenaType(t.Type)

	itemArr, err := coreutils.SplitAsIntArray(t.ArenaItemId)
	if err != nil {
		return template.NewTemplateFieldError("ArenaItemId", err)
	}
	numArr, err := coreutils.SplitAsIntArray(t.ArenaItemAmount)
	if err != nil {
		return template.NewTemplateFieldError("ArenaItemAmount", err)
	}
	if len(itemArr) != len(numArr) {
		err = fmt.Errorf("ArenaItemId[%s]和ArenaItemAmount[%s]长度不相等", t.ArenaItemId, t.ArenaItemAmount)
		return template.NewTemplateFieldError("ArenaItemAmount", err)
	}

	t.itemMap = make(map[int32]int32)
	for i := 0; i < len(itemArr); i++ {
		t.itemMap[itemArr[i]] = numArr[i]
	}

	//
	item2Arr, err := coreutils.SplitAsIntArray(t.LianxuGetItemId)
	if err != nil {
		return template.NewTemplateFieldError("LianxuGetItemId", err)
	}
	num2Arr, err := coreutils.SplitAsIntArray(t.LianxuGetItemCount)
	if err != nil {
		return template.NewTemplateFieldError("LianxuGetItemCount", err)
	}
	if len(item2Arr) != len(num2Arr) {
		err = fmt.Errorf("LianxuGetItemId[%s]LianxuGetItemCount[%s]长度不相等", t.LianxuGetItemId, t.LianxuGetItemCount)
		return template.NewTemplateFieldError("LianxuGetItemCount", err)
	}

	t.extraItemMap = make(map[int32]int32)
	for i := 0; i < len(item2Arr); i++ {
		t.extraItemMap[item2Arr[i]] = num2Arr[i]
	}
	return
}
func (t *ArenaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	if !t.arenaType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	if err = validator.MinValidate(float64(t.ArenaSilver), 0, true); err != nil {
		return template.NewTemplateFieldError("ArenaSilver", err)
	}

	if err = validator.MinValidate(float64(t.LianXuCount), 1, true); err != nil {
		return template.NewTemplateFieldError("LianXuCount", err)
	}

	if err = validator.MinValidate(float64(t.LianXuGetJifen), 1, true); err != nil {
		return template.NewTemplateFieldError("LianXuGetJifen", err)
	}

	for itemId, itemNum := range t.itemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.ArenaItemId)
			return template.NewTemplateFieldError("ArenaItemId", err)
		}
		if itemNum <= 0 {
			err = fmt.Errorf("[%s] invalid", t.ArenaItemAmount)
			return template.NewTemplateFieldError("ArenaItemMount", err)
		}
	}

	for itemId, itemNum := range t.extraItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.ArenaItemId)
			return template.NewTemplateFieldError("ArenaItemId", err)
		}
		if itemNum <= 0 {
			err = fmt.Errorf("[%s] invalid", t.ArenaItemAmount)
			return template.NewTemplateFieldError("ArenaItemMount", err)
		}
	}
	return
}

func (t *ArenaTemplate) PatchAfterCheck() {

}
func (t *ArenaTemplate) TemplateId() int {
	return t.Id
}

func (at *ArenaTemplate) FileName() string {
	return "tb_arena.json"
}

func init() {
	template.Register((*ArenaTemplate)(nil))
}
