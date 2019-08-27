package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	titletypes "fgame/fgame/game/title/types"
	"fmt"
)

type chanChuSilver struct {
	min    int64
	max    int64
	silver int64
}

func newChanChuSilver(min int64, max int64, silver int64) *chanChuSilver {
	chanChu := &chanChuSilver{}
	chanChu.max = max
	chanChu.min = min
	chanChu.silver = silver
	return chanChu
}

func (ccs *chanChuSilver) GetMin() int64 {
	return ccs.min
}

func (ccs *chanChuSilver) GetMax() int64 {
	return ccs.max
}

func (ccs *chanChuSilver) GetSilver() int64 {
	return ccs.silver
}

//抢龙椅配置
type DragonChairTemplate struct {
	*DragonChairTemplateVO
	firstAttrTemplate *AttrTemplate          //属性加成底数
	valueAttrTemplate *AttrTemplate          //属性加成固定
	worshipRewData    *propertytypes.RewData //奖励属性
	worshipRewItemMap map[int32]int32        //膜拜奖励物品
	chanChuSilverList []*chanChuSilver       //产出银两
	specialDropList   []int32                //高级掉落
	commonDropList    []int32                //普通掉落
}

func (dct *DragonChairTemplate) TemplateId() int {
	return dct.Id
}

func (dct *DragonChairTemplate) GetWorshipRewData() *propertytypes.RewData {
	return dct.worshipRewData
}

func (dct *DragonChairTemplate) GetWorshipRewItemMap() map[int32]int32 {
	return dct.worshipRewItemMap
}

func (dct *DragonChairTemplate) GetFirstAttrTemplate() *AttrTemplate {
	return dct.firstAttrTemplate
}

func (dct *DragonChairTemplate) GetValueAttrTemplate() *AttrTemplate {
	return dct.valueAttrTemplate
}

func (dct *DragonChairTemplate) GetSilverList() []*chanChuSilver {
	return dct.chanChuSilverList
}

func (dct *DragonChairTemplate) GetSpecialDropList() []int32 {
	return dct.specialDropList
}

func (dct *DragonChairTemplate) GetCommonDropList() []int32 {
	return dct.commonDropList
}

func (dct *DragonChairTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(dct.FileName(), dct.TemplateId(), err)
			return
		}
	}()

	//验证 silver_worship
	err = validator.MinValidate(float64(dct.SilverWorship), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.SilverWorship)
		return template.NewTemplateFieldError("SilverWorship", err)
	}

	//验证 bind_gold_worship
	err = validator.MinValidate(float64(dct.BindGoldWorship), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.BindGoldWorship)
		return template.NewTemplateFieldError("BindGoldWorship", err)
	}

	//验证 exp_worship
	err = validator.MinValidate(float64(dct.ExpWorship), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.ExpWorship)
		return template.NewTemplateFieldError("ExpWorship", err)
	}

	//验证 exp_point_worship
	err = validator.MinValidate(float64(dct.ExpPointWorship), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.ExpPointWorship)
		return template.NewTemplateFieldError("ExpPointWorship", err)
	}

	if dct.SilverWorship > 0 || dct.BindGoldWorship > 0 || dct.ExpWorship > 0 || dct.ExpPointWorship > 0 {
		dct.worshipRewData = propertytypes.CreateRewData(dct.ExpWorship, dct.ExpPointWorship, dct.SilverWorship, 0, dct.BindGoldWorship)
	}

	//验证 first_attr_id
	tempAttrTemplate := template.GetTemplateService().Get(int(dct.FirstAttrId), (*AttrTemplate)(nil))
	if tempAttrTemplate == nil {
		err = fmt.Errorf("[%d] invalid", dct.FirstAttrId)
		return template.NewTemplateFieldError("FirstAttrId", err)
	}
	attrTemplate, _ := tempAttrTemplate.(*AttrTemplate)
	dct.firstAttrTemplate = attrTemplate

	//验证 value_attr_id
	tempAttrTemplate = template.GetTemplateService().Get(int(dct.ValueAttrId), (*AttrTemplate)(nil))
	if tempAttrTemplate == nil {
		err = fmt.Errorf("[%d] invalid", dct.ValueAttrId)
		return template.NewTemplateFieldError("ValueAttrId", err)
	}
	attrTemplate, _ = tempAttrTemplate.(*AttrTemplate)
	dct.valueAttrTemplate = attrTemplate

	dct.worshipRewItemMap = make(map[int32]int32)
	if dct.ItemWorship != "" {
		itemArr, err := utils.SplitAsIntArray(dct.ItemWorship)
		if err != nil {
			return err
		}

		numArr, err := utils.SplitAsIntArray(dct.ItemCountWorship)
		if err != nil {
			return err
		}

		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", dct.ItemWorship)
			return template.NewTemplateFieldError("ItemWorship", err)
		}

		for i := 0; i < len(itemArr); i++ {
			itemId := itemArr[i]
			num := numArr[i]
			dct.worshipRewItemMap[itemId] = num
		}
	}

	dct.chanChuSilverList = make([]*chanChuSilver, 0, 4)
	if dct.ZhanlingTime1 != "" {
		time1Arr, err := utils.SplitAsIntArray(dct.ZhanlingTime1)
		if err != nil {
			return err
		}
		if len(time1Arr) != 2 {
			err = fmt.Errorf("[%s] invalid", dct.ZhanlingTime1)
			return template.NewTemplateFieldError("ZhanlingTime1", err)
		}
		min := int64(0)
		max := int64(0)
		for index, time1 := range time1Arr {
			if index == 0 {
				min = int64(time1) * int64(common.MINUTE)
			} else {
				max = int64(time1) * int64(common.MINUTE)
			}
		}
		if min > max {
			temp := min
			min = max
			max = temp
		}

		err = validator.MinValidate(float64(dct.ChanchuSilver1), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", dct.ChanchuSilver1)
			return template.NewTemplateFieldError("ChanchuSilver1", err)
		}
		chanChuSilver := newChanChuSilver(min, max, int64(dct.ChanchuSilver1))
		dct.chanChuSilverList = append(dct.chanChuSilverList, chanChuSilver)
	}

	if dct.ZhanlingTime2 != "" {
		time2Arr, err := utils.SplitAsIntArray(dct.ZhanlingTime2)
		if err != nil {
			return err
		}
		if len(time2Arr) != 2 {
			err = fmt.Errorf("[%s] invalid", dct.ZhanlingTime2)
			return template.NewTemplateFieldError("ZhanlingTime2", err)
		}
		min := int64(0)
		max := int64(0)
		for index, time2 := range time2Arr {
			if index == 0 {
				min = int64(time2) * int64(common.MINUTE)
			} else {
				max = int64(time2) * int64(common.MINUTE)
			}
		}
		if min > max {
			temp := min
			min = max
			max = temp
		}

		err = validator.MinValidate(float64(dct.ChanchuSilver2), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", dct.ChanchuSilver2)
			return template.NewTemplateFieldError("ChanchuSilver2", err)
		}
		chanChuSilver := newChanChuSilver(min, max, int64(dct.ChanchuSilver2))
		dct.chanChuSilverList = append(dct.chanChuSilverList, chanChuSilver)
	}

	if dct.ZhanlingTime3 != "" {
		time3Arr, err := utils.SplitAsIntArray(dct.ZhanlingTime3)
		if err != nil {
			return err
		}
		if len(time3Arr) != 2 {
			err = fmt.Errorf("[%s] invalid", dct.ZhanlingTime3)
			return template.NewTemplateFieldError("ZhanlingTime3", err)
		}
		min := int64(0)
		max := int64(0)
		for index, time3 := range time3Arr {
			if index == 0 {
				min = int64(time3) * int64(common.MINUTE)
			} else {
				max = int64(time3) * int64(common.MINUTE)
			}
		}
		if min > max {
			temp := min
			min = max
			max = temp
		}

		err = validator.MinValidate(float64(dct.ChanchuSilver3), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", dct.ChanchuSilver3)
			return template.NewTemplateFieldError("ChanchuSilver3", err)
		}
		chanChuSilver := newChanChuSilver(min, max, int64(dct.ChanchuSilver3))
		dct.chanChuSilverList = append(dct.chanChuSilverList, chanChuSilver)
	}

	if dct.ZhanlingTime4 != "" {
		time4Arr, err := utils.SplitAsIntArray(dct.ZhanlingTime4)
		if err != nil {
			return err
		}
		if len(time4Arr) != 2 {
			err = fmt.Errorf("[%s] invalid", dct.ZhanlingTime4)
			return template.NewTemplateFieldError("ZhanlingTime4", err)
		}
		min := int64(0)
		max := int64(0)
		for index, time4 := range time4Arr {
			if index == 0 {
				min = int64(time4) * int64(common.MINUTE)
			} else {
				max = int64(time4) * int64(common.MINUTE)
			}
		}
		if min > max {
			temp := min
			min = max
			max = temp
		}

		err = validator.MinValidate(float64(dct.ChanchuSilver4), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", dct.ChanchuSilver4)
			return template.NewTemplateFieldError("ChanchuSilver4", err)
		}
		chanChuSilver := newChanChuSilver(min, max, int64(dct.ChanchuSilver4))
		dct.chanChuSilverList = append(dct.chanChuSilverList, chanChuSilver)
	}

	err = validator.MinValidate(float64(dct.ChanchuSilver0), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.ChanchuSilver0)
		return template.NewTemplateFieldError("ChanchuSilver0", err)
	}

	dct.specialDropList = make([]int32, 0, 4)
	if dct.SpecialDrop != "" {
		specialDropList, err := utils.SplitAsIntArray(dct.SpecialDrop)
		if err != nil {
			return err
		}
		dct.specialDropList = append(dct.specialDropList, specialDropList...)
	}

	dct.commonDropList = make([]int32, 0, 4)
	if dct.CommonDrop != "" {
		commonDropList, err := utils.SplitAsIntArray(dct.CommonDrop)
		if err != nil {
			return err
		}
		dct.commonDropList = append(dct.commonDropList, commonDropList...)
	}

	err = validator.MinValidate(float64(dct.DropTime1), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.DropTime1)
		return template.NewTemplateFieldError("DropTime1", err)
	}

	err = validator.MinValidate(float64(dct.DropTime2), float64(dct.DropTime1), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.DropTime2)
		return template.NewTemplateFieldError("DropTime2", err)
	}

	err = validator.MinValidate(float64(dct.DropTime3), float64(dct.DropTime2), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.DropTime3)
		return template.NewTemplateFieldError("DropTime3", err)
	}

	err = validator.MinValidate(float64(dct.DropTime4), float64(dct.DropTime3), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.DropTime4)
		return template.NewTemplateFieldError("DropTime4", err)
	}

	err = validator.RangeValidate(float64(dct.GoldPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.GoldPercent)
		err = template.NewTemplateFieldError("GoldPercent", err)
		return
	}
	return nil
}

func (dct *DragonChairTemplate) PatchAfterCheck() {

}

func (dct *DragonChairTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(dct.FileName(), dct.TemplateId(), err)
			return
		}
	}()

	//验证 item_worship
	for itemId, num := range dct.worshipRewItemMap {
		tempAttrTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if tempAttrTemplate == nil {
			err = fmt.Errorf("[%s] invalid", dct.ItemWorship)
			return template.NewTemplateFieldError("ItemWorship", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", dct.ItemCountWorship)
			return template.NewTemplateFieldError("ItemCountWorship", err)
		}
	}

	//验证 first_gold
	err = validator.MinValidate(float64(dct.FirstGold), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.FirstGold)
		return template.NewTemplateFieldError("FirstGold", err)
	}

	//验证 coefficient_gold
	err = validator.MinValidate(dct.CoefficientGold, float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%f] invalid", dct.CoefficientGold)
		return template.NewTemplateFieldError("CoefficientGold", err)
	}

	//验证 value_gold
	err = validator.MinValidate(float64(dct.ValueGold), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.ValueGold)
		return template.NewTemplateFieldError("ValueGold", err)
	}

	//验证 coefficient_attr
	err = validator.MinValidate(dct.CoefficientAttr, float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%f] invalid", dct.CoefficientAttr)
		return template.NewTemplateFieldError("CoefficientAttr", err)
	}

	//验证 title_id
	tempTemplate := template.GetTemplateService().Get(int(dct.TitleId), (*TitleTemplate)(nil))
	if tempTemplate == nil {
		err = fmt.Errorf("[%d] invalid", dct.TitleId)
		return template.NewTemplateFieldError("TitleId", err)
	}
	titleTemplate, _ := tempTemplate.(*TitleTemplate)
	titleType := titleTemplate.GetTitleType()
	if titleType != titletypes.TitleTypeKing {
		err = fmt.Errorf("[%d] invalid", dct.TitleId)
		return template.NewTemplateFieldError("TitleId", err)
	}

	//验证 worship_count
	err = validator.MinValidate(float64(dct.WorshipCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.WorshipCount)
		return template.NewTemplateFieldError("WorshipCount", err)
	}

	//验证 worship_chest_silver
	err = validator.MinValidate(float64(dct.WorshipChestSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.WorshipChestSilver)
		return template.NewTemplateFieldError("WorshipChestSilver", err)
	}

	//验证  auto_silver
	err = validator.MinValidate(float64(dct.AutoSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.AutoSilver)
		return template.NewTemplateFieldError("AutoSilver", err)
	}

	//验证 auto_silver_time
	err = validator.MinValidate(float64(dct.AutoSilverTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", dct.AutoSilverTime)
		return template.NewTemplateFieldError("AutoSilverTime", err)
	}

	//验证 worship_chest_silver
	err = validator.MinValidate(float64(dct.ChestMax), float64(1), true)
	if err != nil || dct.ChestMax < dct.WorshipChestSilver {
		err = fmt.Errorf("[%d] invalid", dct.ChestMax)
		return template.NewTemplateFieldError("ChestMax", err)
	}

	// if dct.WorshipChestSilver > 0 {
	// 	if dct.ChestMax%dct.WorshipChestSilver != 0 {
	// 		err = fmt.Errorf("[%d] invalid", dct.ChestMax)
	// 		return template.NewTemplateFieldError("ChestMax", err)
	// 	}
	// }

	return nil
}

func (dct *DragonChairTemplate) FileName() string {
	return "tb_dragon_chair.json"
}

func init() {
	template.Register((*DragonChairTemplate)(nil))
}
