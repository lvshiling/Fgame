package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	"fmt"
)

type ShangguzhilingLingwenTemplate struct {
	*ShangguzhilingLingwenTemplateVO
	//灵兽类型
	lingshouType shangguzhilingtypes.LingshouType
	//灵纹类型
	lingwenType shangguzhilingtypes.LingwenType
	//升级起始Id
	firstLevelTemp *ShangguzhilingLingwenLevelTemplate
	levelMap       map[int32]*ShangguzhilingLingwenLevelTemplate
	//灵纹升级使用的物品ID列表
	lingwenUseItemIdList []int32
}

// 灵纹升级可使用的物品列表
func (t *ShangguzhilingLingwenTemplate) GetLingWenUpLevelUseItemList() []int32 {
	return t.lingwenUseItemIdList
}

//灵纹类型
func (t *ShangguzhilingLingwenTemplate) GetLingWenType() shangguzhilingtypes.LingwenType {
	return t.lingwenType
}

//灵兽类型
func (t *ShangguzhilingLingwenTemplate) GetLingShouType() shangguzhilingtypes.LingshouType {
	return t.lingshouType
}

func (t *ShangguzhilingLingwenTemplate) GetLevelTemp(level int32) *ShangguzhilingLingwenLevelTemplate {
	temp, ok := t.levelMap[level]
	if !ok {
		return nil
	}
	return temp
}

func (t *ShangguzhilingLingwenTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingLingwenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingLingwenTemplate) PatchAfterCheck() {
	t.levelMap = make(map[int32]*ShangguzhilingLingwenLevelTemplate)
	for temp := t.firstLevelTemp; temp != nil; temp = temp.GetNextLevelTemp() {
		t.levelMap[temp.Level] = temp
	}
}

func (t *ShangguzhilingLingwenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//灵兽类型
	lingshouType := shangguzhilingtypes.LingshouType(t.Type)
	if !lingshouType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	t.lingshouType = lingshouType

	//灵纹类型
	lingwenType := shangguzhilingtypes.LingwenType(t.SubType)
	if !lingwenType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SubType)
		return template.NewTemplateFieldError("SubType", err)
	}
	t.lingwenType = lingwenType

	//所需的上古之灵等级
	err = validator.MinValidate(float64(t.NeedSgzlLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedSgzlLevel)
		return template.NewTemplateFieldError("NeedSgzlLevel", err)
	}

	//灵纹升级起始ID
	lingwenLevelTempInterface := template.GetTemplateService().Get(int(t.LingwenLevelBeginId), (*ShangguzhilingLingwenLevelTemplate)(nil))
	if lingwenLevelTempInterface == nil {
		err = fmt.Errorf("ShangguzhilingLingwenLevelTemplate [%d] no exist", t.LingwenLevelBeginId)
		return template.NewTemplateFieldError("LingwenLevelBeginId", err)
	}
	lingwenLevelTemp, ok := lingwenLevelTempInterface.(*ShangguzhilingLingwenLevelTemplate)
	if !ok {
		err = fmt.Errorf("ShangguzhilingLingwenLevelTemplate assert [%d] no exist", t.LingwenLevelBeginId)
		return template.NewTemplateFieldError("LingwenLevelBeginId", err)
	}
	t.firstLevelTemp = lingwenLevelTemp

	//上古之灵灵纹升级使用的物品ID
	lingwenUseItemIdList, err := coreutils.SplitAsIntArray(t.LingwenLevelUseItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingwenLevelUseItemId)
		return template.NewTemplateFieldError("LingwenLevelUseItemId", err)
	}
	t.lingwenUseItemIdList = lingwenUseItemIdList

	return nil
}

func (t *ShangguzhilingLingwenTemplate) FileName() string {
	return "tb_sgzl_lingwen_pos.json"
}

func init() {
	template.Register((*ShangguzhilingLingwenTemplate)(nil))
}
