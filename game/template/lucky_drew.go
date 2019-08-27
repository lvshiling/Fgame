package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sort"
)

//抽奖活动配置
type LuckyDrewTemplate struct {
	*LuckyDrewTemplateVO
	drewType       welfaretypes.OpenActivityDrewSubType
	dropByTimesMap map[int32]int32 //按次数必定掉落map
	timesDescList  []int           //循环间隔，降序
	rateList       []int64         //权重列表
	useItemMap     map[int32]int32 //消耗物品map
	giveItemMap    map[int32]int32 //额外奖励物品map
}

func (t *LuckyDrewTemplate) TemplateId() int {
	return t.Id
}

func (t *LuckyDrewTemplate) GetDrewType() welfaretypes.OpenActivityDrewSubType {
	return t.drewType
}

func (t *LuckyDrewTemplate) GetUseItemMap() map[int32]int32 {
	return t.useItemMap
}

func (t *LuckyDrewTemplate) GetGiveItemMap() map[int32]int32 {
	return t.giveItemMap
}

func (t *LuckyDrewTemplate) GetRandomResultType() welfaretypes.DrewResultType {
	index := mathutils.RandomWeights(t.rateList)
	return welfaretypes.DrewResultType(index)
}

func (t *LuckyDrewTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//消耗物品
	t.useItemMap = make(map[int32]int32)
	itemArr, err := utils.SplitAsIntArray(t.ItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}
	itemNumArr, err := utils.SplitAsIntArray(t.ItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemCount)
		return template.NewTemplateFieldError("ItemCount", err)
	}
	if len(itemArr) != len(itemNumArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.ItemId, t.ItemCount)
		return template.NewTemplateFieldError("ItemId and ItemCount", err)
	}
	if len(itemArr) > 0 {
		//组合数据
		for index, itemId := range itemArr {
			t.useItemMap[itemId] = itemNumArr[index]
		}
	}

	//额外奖励物品
	t.giveItemMap = make(map[int32]int32)
	giveItemArr, err := utils.SplitAsIntArray(t.GiveItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GiveItemId)
		return template.NewTemplateFieldError("GiveItemId", err)
	}
	giveItemNumArr, err := utils.SplitAsIntArray(t.GiveItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GiveItemCount)
		return template.NewTemplateFieldError("GiveItemCount", err)
	}
	if len(giveItemArr) != len(giveItemNumArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.GiveItemId, t.GiveItemCount)
		return template.NewTemplateFieldError("GiveItemId and GiveItemCount", err)
	}
	if len(giveItemArr) > 0 {
		//组合数据
		for index, itemId := range giveItemArr {
			t.giveItemMap[itemId] = giveItemNumArr[index]
		}
	}

	//类型
	typ := welfaretypes.OpenActivityDrewSubType(t.Type)
	if !typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	t.drewType = typ

	//按次数必定掉落map
	t.dropByTimesMap = make(map[int32]int32)
	if t.MustAmount1 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount1]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount1)
			return template.NewTemplateFieldError("MustAmount1", err)
		}
		t.dropByTimesMap[t.MustAmount1] = t.MustGet1
		t.timesDescList = append(t.timesDescList, int(t.MustAmount1))
	}
	if t.MustAmount2 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount2]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount2)
			return template.NewTemplateFieldError("MustAmount2", err)
		}
		t.dropByTimesMap[t.MustAmount2] = t.MustGet2
		t.timesDescList = append(t.timesDescList, int(t.MustAmount2))
	}
	if t.MustAmount3 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount3]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount3)
			return template.NewTemplateFieldError("MustAmount3", err)
		}
		t.dropByTimesMap[t.MustAmount3] = t.MustGet3
		t.timesDescList = append(t.timesDescList, int(t.MustAmount3))
	}
	if t.MustAmount4 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount4]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount4)
			return template.NewTemplateFieldError("MustAmount4", err)
		}
		t.dropByTimesMap[t.MustAmount4] = t.MustGet4
		t.timesDescList = append(t.timesDescList, int(t.MustAmount4))
	}
	if t.MustAmount5 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount5]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount5)
			return template.NewTemplateFieldError("MustAmount5", err)
		}
		t.dropByTimesMap[t.MustAmount5] = t.MustGet5
		t.timesDescList = append(t.timesDescList, int(t.MustAmount5))
	}
	if t.MustAmount6 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount6]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount6)
			return template.NewTemplateFieldError("MustAmount6", err)
		}
		t.dropByTimesMap[t.MustAmount6] = t.MustGet6
		t.timesDescList = append(t.timesDescList, int(t.MustAmount6))
	}
	if t.MustAmount7 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount7]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount7)
			return template.NewTemplateFieldError("MustAmount7", err)
		}
		t.dropByTimesMap[t.MustAmount7] = t.MustGet7
		t.timesDescList = append(t.timesDescList, int(t.MustAmount7))
	}
	if t.MustAmount8 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount8]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount8)
			return template.NewTemplateFieldError("MustAmount8", err)
		}
		t.dropByTimesMap[t.MustAmount8] = t.MustGet8
		t.timesDescList = append(t.timesDescList, int(t.MustAmount8))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(t.timesDescList)))
	return
}

func (t *LuckyDrewTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//概率
	err = validator.MinValidate(float64(t.Rate), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Rate", err)
	}
	//翻倍系数
	err = validator.MinValidate(float64(t.RewTimes1), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RewTimes1", err)
	}
	//翻倍概率
	err = validator.MinValidate(float64(t.Percent1), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Percent1", err)
	}
	//翻倍系数
	err = validator.MinValidate(float64(t.RewTimes2), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RewTimes2", err)
	}
	//翻倍概率
	err = validator.MinValidate(float64(t.Percent2), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Percent2", err)
	}
	if t.Rate+t.Percent1+t.Percent2 != common.MAX_RATE {
		err := fmt.Errorf("概率总和不是一万")
		return template.NewTemplateFieldError("Rate", err)
	}

	if t.drewType == welfaretypes.OpenActivityDrewSubTypeChargeDrew {
		//后台规则条件
		err = validator.MinValidate(float64(t.RewCount), float64(0), false)
		if err != nil {
			return template.NewTemplateFieldError("RewCount", err)
		}
	}

	//抽奖等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Level", err)
	}

	//校验物品
	for itemId, num := range t.useItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("ItemId", fmt.Errorf("[%s] invalid", t.ItemId))
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("ItemCount", fmt.Errorf("[%s] invalid", t.ItemCount))
		}
	}

	//校验物品
	for itemId, num := range t.giveItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("GiveItemId", fmt.Errorf("[%s] invalid", t.GiveItemId))
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("GiveItemCount", fmt.Errorf("[%s] invalid", t.GiveItemCount))
		}
	}
	// if t.NeedTimes < 1 && len(t.useItemMap) == 0 {
	// 	err = fmt.Errorf("[%d] invalid", t.NeedTimes)
	// 	return template.NewTemplateFieldError("NeedTimes", err)
	// }

	// 校验消耗次数
	err = validator.MinValidate(float64(t.CostTimes), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CostTimes)
		return template.NewTemplateFieldError("CostTimes", err)
	}

	return nil
}

func (t *LuckyDrewTemplate) PatchAfterCheck() {
	if t.Rate >= 0 {
		t.rateList = append(t.rateList, int64(t.Rate))
	}

	if t.Percent1 >= 0 {
		t.rateList = append(t.rateList, int64(t.Percent1))
	}

	if t.Percent2 >= 0 {
		t.rateList = append(t.rateList, int64(t.Percent2))
	}
}

func (t *LuckyDrewTemplate) FileName() string {
	return "tb_yunying_choujiang.json"
}

func (t *LuckyDrewTemplate) GetRewDropByTimesMap() map[int32]int32 {
	return t.dropByTimesMap
}

func (t *LuckyDrewTemplate) GetTimesDesc() []int {
	return t.timesDescList
}

func init() {
	template.Register((*LuckyDrewTemplate)(nil))
}
