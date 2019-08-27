package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	weektypes "fgame/fgame/game/week/types"
	"fmt"
)

//周卡配置
type WeekTemplate struct {
	*WeekTemplateVO
	weekType         weektypes.WeekType
	extralRewItemMap map[int32]int32
	firstRewItemMap  map[int32]int32
	startWeekDayTemp *WeekDayTemplate
	weekDayTempMap   map[int32]*WeekDayTemplate
}

func (t *WeekTemplate) TemplateId() int {
	return t.Id
}

func (t *WeekTemplate) IsExtralRew(cycDay int32) bool {
	if cycDay%t.EwaiNeedDay == 0 {
		return true
	}

	return false
}

func (t *WeekTemplate) GetWeekType() weektypes.WeekType {
	return t.weekType
}

func (t *WeekTemplate) GetCycDayRew(day int32) *WeekDayTemplate {
	cycle := int32(len(t.weekDayTempMap))
	if cycle == 0 {
		return nil
	}

	day = day % cycle
	if day == 0 {
		day = cycle
	}
	return t.weekDayTempMap[day]
}

func (t *WeekTemplate) GetExtralRewItemMap() map[int32]int32 {
	return t.extralRewItemMap
}

func (t *WeekTemplate) GetRewFirstItemMap() map[int32]int32 {
	return t.firstRewItemMap
}

func (t *WeekTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 第一次奖励
	t.firstRewItemMap = make(map[int32]int32)
	getItemIdList, err := utils.SplitAsIntArray(t.GetItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GetItem)
		return template.NewTemplateFieldError("GetItem", err)
	}
	getItemCountList, err := utils.SplitAsIntArray(t.GetItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GetItemCount)
		return template.NewTemplateFieldError("GetItemCount", err)
	}
	if len(getItemIdList) != len(getItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.GetItem, t.GetItemCount)
		return template.NewTemplateFieldError("GetItem or GetItemCount", err)
	}
	//组合数据
	for index, itemId := range getItemIdList {
		t.firstRewItemMap[itemId] = getItemCountList[index]
	}

	//验证 rew_item_id
	t.extralRewItemMap = make(map[int32]int32)
	extralItemIdList, err := utils.SplitAsIntArray(t.EwaiGetItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.EwaiGetItem)
		return template.NewTemplateFieldError("EwaiGetItem", err)
	}
	extralItemCountList, err := utils.SplitAsIntArray(t.EwaiGetItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.EwaiGetItemCount)
		return template.NewTemplateFieldError("EwaiGetItemCount", err)
	}
	if len(extralItemIdList) != len(extralItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.EwaiGetItem, t.EwaiGetItemCount)
		return template.NewTemplateFieldError("EwaiGetItem or EwaiGetItemCount", err)
	}
	//组合数据
	for index, itemId := range extralItemIdList {
		t.extralRewItemMap[itemId] = extralItemCountList[index]
	}

	//
	if t.EveryDayGetBegin > 0 {
		to := template.GetTemplateService().Get(int(t.EveryDayGetBegin), (*WeekDayTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.EveryDayGetBegin)
			err = template.NewTemplateFieldError("EveryDayGetBegin", err)
			return
		}
		t.startWeekDayTemp = to.(*WeekDayTemplate)
	}

	return nil
}

func (t *WeekTemplate) PatchAfterCheck() {
	//加载所有
	t.weekDayTempMap = make(map[int32]*WeekDayTemplate)
	for startTemp := t.startWeekDayTemp; startTemp != nil; startTemp = startTemp.GetNextTemp() {
		t.weekDayTempMap[startTemp.DayInt] = startTemp
	}

}

func (t *WeekTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//类型
	t.weekType = weektypes.WeekType(t.Type)
	if !t.weekType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//购买所需元宝
	err = validator.MinValidate(float64(t.NeedGold), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedGold)
		return template.NewTemplateFieldError("NeedGold", err)
	}

	// 持续时间
	err = validator.MinValidate(float64(t.Duration), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Duration)
		return template.NewTemplateFieldError("Duration", err)
	}

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
	// 额外奖励天数
	err = validator.MinValidate(float64(t.EwaiNeedDay), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.EwaiNeedDay)
		return template.NewTemplateFieldError("EwaiNeedDay", err)
	}

	//
	for itemId, num := range t.firstRewItemMap {
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
	//
	for itemId, num := range t.extralRewItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			err = fmt.Errorf("[%s] invalid", t.EwaiGetItem)
			return template.NewTemplateFieldError("EwaiGetItem", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.EwaiGetItemCount)
			return template.NewTemplateFieldError("EwaiGetItemCount", err)
		}
	}

	return nil
}

func (t *WeekTemplate) FileName() string {
	return "tb_zhouka.json"
}

func init() {
	template.Register((*WeekTemplate)(nil))
}
