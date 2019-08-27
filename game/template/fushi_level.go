package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	fushitypes "fgame/fgame/game/fushi/types"
	"fmt"
)

type FuShiLevelTemplate struct {
	*FuShiLevelTemplateVO
	fushiType fushitypes.FuShiType
	itemMap   map[int32]int32
}

// 返回符石类型
func (t *FuShiLevelTemplate) GetFuShiType() fushitypes.FuShiType {
	return t.fushiType
}

// 返回升级消耗物品
func (t *FuShiLevelTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *FuShiLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *FuShiLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 升级消耗物品
	t.itemMap = make(map[int32]int32)
	useItemIdArr, err := utils.SplitAsIntArray(t.UpLevelItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UpLevelItem)
		return template.NewTemplateFieldError("UpLevelItem", err)
	}
	useItemCountArr, err := utils.SplitAsIntArray(t.UpLevelItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UpLevelItemCount)
		return template.NewTemplateFieldError("UpLevelItemCount", err)
	}
	if len(useItemIdArr) != len(useItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.UpLevelItem, t.UpLevelItemCount)
		return template.NewTemplateFieldError("UpLevelItem or UpLevelItemCount", err)
	}
	if len(useItemIdArr) > 0 {
		for index, itemId := range useItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.UpLevelItem)
				return template.NewTemplateFieldError("UpLevelItem", err)
			}

			err = validator.MinValidate(float64(useItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("UpLevelItemCount", err)
			}

			t.itemMap[itemId] = useItemCountArr[index]
		}
	}

	return
}

func (t *FuShiLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*FuShiLevelTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		temp := to.(*FuShiLevelTemplate)

		if temp.Type != t.Type {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		diff := temp.Level - int32(t.Level)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	// 验证 符石类型
	t.fushiType = fushitypes.FuShiType(t.Type)
	if !t.fushiType.Vaild() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	// 验证 等级
	err = validator.MinValidate(float64(t.Level), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("Level", err)
	}

	// 验证 技能id
	if t.SkillId != 0 {
		to := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
	}

	//验证 成功率
	err = validator.RangeValidate(float64(t.UpLevelRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpLevelRate)
		err = template.NewTemplateFieldError("UpLevelRate", err)
		return
	}

	return
}

func (t *FuShiLevelTemplate) PatchAfterCheck() {
}

func (t *FuShiLevelTemplate) FileName() string {
	return "tb_baguaifushi_level.json"
}

func init() {
	template.Register((*FuShiLevelTemplate)(nil))
}
