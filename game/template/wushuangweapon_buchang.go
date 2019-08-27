package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
)

type WushuangWeaponBuchangTemplate struct {
	*WushuangWeaponBuchangTemplateVO
	rewItemMap map[int32]int32
	rewTime    int64
}

func (t *WushuangWeaponBuchangTemplate) TemplateId() int {
	return t.Id
}

func (t *WushuangWeaponBuchangTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *WushuangWeaponBuchangTemplate) GetRewTime() int64 {
	return t.rewTime
}

//检查有效性
func (t *WushuangWeaponBuchangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//时间
	err = validator.MinValidate(float64(t.Time), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Time)
		return template.NewTemplateFieldError("Time", err)
	}
	rewTime, err := timeutils.ParseDayOfYYYYMMDDHHMM(strconv.FormatInt(int64(t.Time), 10))
	if err != nil {
		return template.NewTemplateFieldError("Time", fmt.Errorf("[%d] invalid", t.Time))
	}
	t.rewTime = rewTime

	//辨识物品Id
	err = validator.MinValidate(float64(t.ItemId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}

	//奖励物品
	itemIdArr, err := coreutils.SplitAsIntArray(t.RawItem_id)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RawItem_id)
		return template.NewTemplateFieldError("RawItem_id", err)
	}
	itemCountArr, err := coreutils.SplitAsIntArray(t.RawItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RawItemCount)
		return template.NewTemplateFieldError("RawItemCount", err)
	}
	if len(itemIdArr) != len(itemCountArr) {
		err = fmt.Errorf("[%s] [%s] invalid", t.RawItem_id, t.RawItemCount)
		return template.NewTemplateFieldError("RawItem_id or RawItemCount", err)
	}
	t.rewItemMap = make(map[int32]int32)
	for i := 0; i < len(itemIdArr); i++ {
		itemId := itemIdArr[i]
		itemCount := itemCountArr[i]
		tempItemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.RawItem_id)
			return template.NewTemplateFieldError("RawItem_id", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(itemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.RawItemCount)
			return template.NewTemplateFieldError("RawItemCount", err)
		}
		t.rewItemMap[itemId] = itemCount
	}

	return
}

//组合成需要的数据
func (t *WushuangWeaponBuchangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return
}

//检验后组合
func (t *WushuangWeaponBuchangTemplate) PatchAfterCheck() {
}

func (t *WushuangWeaponBuchangTemplate) FileName() string {
	return "tb_shenjia_buchang.json"
}

func init() {
	template.Register((*WushuangWeaponBuchangTemplate)(nil))
}
