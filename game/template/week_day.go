package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//周卡每日奖励配置
type WeekDayTemplate struct {
	*WeekDayTemplateVO
	nextTemp   *WeekDayTemplate
	rewItemMap map[int32]int32
}

func (t *WeekDayTemplate) TemplateId() int {
	return t.Id
}

func (t *WeekDayTemplate) GetNextTemp() *WeekDayTemplate {
	return t.nextTemp
}

func (t *WeekDayTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *WeekDayTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewItemMap = make(map[int32]int32)
	//验证 rew_item_id
	rewItemIdList, err := utils.SplitAsIntArray(t.GetItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GetItem)
		return template.NewTemplateFieldError("GetItem", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.GetItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GetItemCount)
		return template.NewTemplateFieldError("GetItemCount", err)
	}
	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.GetItem, t.GetItemCount)
		return template.NewTemplateFieldError("GetItem or GetItemCount", err)
	}
	//组合数据
	for index, itemId := range rewItemIdList {
		t.rewItemMap[itemId] = rewItemCountList[index]
	}

	//
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*WeekDayTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		t.nextTemp = to.(*WeekDayTemplate)
	}
	return nil
}

func (t *WeekDayTemplate) PatchAfterCheck() {

}

func (t *WeekDayTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 奖励银两
	err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewSilver)
		return template.NewTemplateFieldError("RewSilver", err)
	}
	// 奖励元宝
	err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewGold)
		return template.NewTemplateFieldError("RewGold", err)
	}
	// 奖励绑元
	err = validator.MinValidate(float64(t.RewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewBindGold)
		return template.NewTemplateFieldError("RewBindGold", err)
	}
	// 天数
	err = validator.MinValidate(float64(t.DayInt), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DayInt)
		return template.NewTemplateFieldError("DayInt", err)
	}

	//
	for itemId, num := range t.rewItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			err = fmt.Errorf("[%s] invalid", t.GetItem)
			return template.NewTemplateFieldError("GetItem", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.GetItemCount)
			return template.NewTemplateFieldError("GetItemCount", err)
		}
	}
	return nil
}

func (t *WeekDayTemplate) FileName() string {
	return "tb_zhouka_day.json"
}

func init() {
	template.Register((*WeekDayTemplate)(nil))
}
