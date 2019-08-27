package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

func init() {
	template.Register((*MoonloveTemplate)(nil))
}

//月下情缘配置
type MoonloveTemplate struct {
	*MoonloveTemplateVO
	rewardItemMap map[int32]int32
}

func (t *MoonloveTemplate) FileName() string {
	return "tb_moon_love.json"
}

func (t *MoonloveTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewardItemMap = make(map[int32]int32)

	//奖励物品
	itemIdArr, err := utils.SplitAsIntArray(t.RewItemId)
	if err != nil {
		return template.NewTemplateFieldError("RewItemId", fmt.Errorf("[%s] invalid", t.RewItemId))
	}
	//奖励数量
	itemCountArr, err := utils.SplitAsIntArray(t.RewItemCount)
	if err != nil {
		return template.NewTemplateFieldError("RewItemCount", fmt.Errorf("[%s] invalid", t.RewItemCount))
	}
	if len(itemIdArr) != len(itemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.RewItemId, t.RewItemCount)
		return template.NewTemplateFieldError("RewItemId or RewItemCount", err)
	}
	if len(itemIdArr) > 0 {
		t.rewardItemMap = make(map[int32]int32)
		for index, itemId := range itemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("RewItemId", fmt.Errorf("[%s] invalid", t.RewItemId))
			}

			err = validator.MinValidate(float64(itemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("RewItemCount", err)
			}

			t.rewardItemMap[itemId] = itemCountArr[index]
		}

	}

	return nil
}

func (t *MoonloveTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//最小等级
	if err = validator.MinValidate(float64(t.MinLev), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("MinLev", err)
		return
	}
	//最大等级
	if err = validator.MaxValidate(float64(t.MaxLev), float64(999), true); err != nil {
		err = template.NewTemplateFieldError("MaxLev", err)
		return
	}
	//进入后多少毫秒获得奖励
	if err = validator.MinValidate(float64(t.FristTiem), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("FristTiem", err)
		return
	}
	//奖励时间间隔
	if err = validator.MinValidate(float64(t.RewTiem), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewTiem", err)
		return
	}
	//单次发放的经验
	if err = validator.MinValidate(float64(t.RewExp), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}
	//单次发放的经验点
	if err = validator.MinValidate(float64(t.RewExpPoint), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewExpPoint", err)
		return
	}
	//奖励银两
	if err = validator.MinValidate(float64(t.RewSilver), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}
	//奖励元宝
	if err = validator.MinValidate(float64(t.RewGold), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}
	//奖励绑元
	if err = validator.MinValidate(float64(t.RewBindGold), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewBindGold", err)
		return
	}
	//双人状态系数
	if err = validator.MinValidate(float64(t.DoubleMan), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("DoubleMan", err)
		return
	}

	return nil
}

func (t *MoonloveTemplate) PatchAfterCheck() {

}

func (t *MoonloveTemplate) TemplateId() int {
	return t.Id
}

func (t *MoonloveTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewardItemMap
}
