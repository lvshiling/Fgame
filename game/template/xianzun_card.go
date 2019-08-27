package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/scene/types"
	xianfutypes "fgame/fgame/game/xianfu/types"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
	"fmt"
)

// 仙尊特权卡配置
type XianZunCardTemplate struct {
	*XianZunCardTemplateVO
	xianZunCardType   xianzuncardtypes.XianZunCardType
	receiveItemMap    map[int32]int32
	dayReceiveItemMap map[int32]int32
	emailItemMap      map[int32]int32
	expBiologyList    []types.BiologySetType // 获得经验加成的生物类型
}

func (t *XianZunCardTemplate) GetXianFuFreeTimes(xianFuType xianfutypes.XianfuType) int32 {
	switch xianFuType {
	case xianfutypes.XianfuTypeSilver:
		return t.SilverXianFuFreeAdd
	case xianfutypes.XianfuTypeExp:
		return t.ExpXianFuFreeAdd
	default:
		return 0
	}
}

func (t *XianZunCardTemplate) TemplateId() int {
	return t.Id
}

func (t *XianZunCardTemplate) GetXianZunCardType() xianzuncardtypes.XianZunCardType {
	return t.xianZunCardType
}

func (t *XianZunCardTemplate) GetReceiveItemMap() map[int32]int32 {
	return t.receiveItemMap
}

func (t *XianZunCardTemplate) GetDayReceiveItemMap() map[int32]int32 {
	return t.dayReceiveItemMap
}

func (t *XianZunCardTemplate) GetEmailItemMap() map[int32]int32 {
	return t.emailItemMap
}

func (t *XianZunCardTemplate) GetExpBiologyList() []types.BiologySetType {
	return t.expBiologyList
}

func (t *XianZunCardTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证仙尊特权卡类型
	t.xianZunCardType = xianzuncardtypes.XianZunCardType(t.Type)
	if !t.xianZunCardType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//验证激活获得物品id
	t.receiveItemMap = make(map[int32]int32)
	receiveItemAttr, err := utils.SplitAsIntArray(t.ItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}
	receiveItemCountAttr, err := utils.SplitAsIntArray(t.ItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemCount)
		return template.NewTemplateFieldError("ItemCount", err)
	}
	if len(receiveItemAttr) != len(receiveItemCountAttr) {
		err = fmt.Errorf("[%s][%s] invalid", t.ItemId, t.ItemCount)
		return template.NewTemplateFieldError("ItemId or ItemCount", err)
	}
	if len(receiveItemAttr) > 0 {
		for index, itemId := range receiveItemAttr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.ItemId)
				return template.NewTemplateFieldError("ItemId", err)
			}

			err = validator.MinValidate(float64(receiveItemCountAttr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("ItemCount", err)
			}

			t.receiveItemMap[itemId] = receiveItemCountAttr[index]
		}
	}

	//验证每日领取物品id
	t.dayReceiveItemMap = make(map[int32]int32)
	dayReceiveItemAttr, err := utils.SplitAsIntArray(t.DayRewItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.DayRewItem)
		return template.NewTemplateFieldError("DayRewItem", err)
	}
	dayReceiveItemCountAttr, err := utils.SplitAsIntArray(t.DayRewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.DayRewItemCount)
		return template.NewTemplateFieldError("DayRewItemCount", err)
	}
	if len(dayReceiveItemAttr) != len(dayReceiveItemCountAttr) {
		err = fmt.Errorf("[%s][%s] invalid", t.DayRewItem, t.DayRewItemCount)
		return template.NewTemplateFieldError("DayRewItem or DayRewItemCount", err)
	}
	if len(dayReceiveItemAttr) > 0 {
		for index, itemId := range dayReceiveItemAttr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.DayRewItem)
				return template.NewTemplateFieldError("DayRewItem", err)
			}

			err = validator.MinValidate(float64(dayReceiveItemCountAttr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("DayRewItemCount", err)
			}

			t.dayReceiveItemMap[itemId] = dayReceiveItemCountAttr[index]
		}
	}

	//验证经验加成的生物类型
	t.expBiologyList = make([]types.BiologySetType, 0, 8)
	biologySetTypeAttr, err := utils.SplitAsIntArray(t.BiologySetType)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BiologySetType)
		return template.NewTemplateFieldError("BiologySetType", err)
	}
	for _, typ := range biologySetTypeAttr {
		biologyType := types.BiologySetType(typ)
		if !biologyType.Valid() {
			err = fmt.Errorf("[%s] invalid", t.BiologySetType)
			return template.NewTemplateFieldError("BiologySetType", err)
		}
		t.expBiologyList = append(t.expBiologyList, biologyType)
	}
	
	return
}

func (t *XianZunCardTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 need_gold
	err = validator.MinValidate(float64(t.NeedGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedGold)
		return template.NewTemplateFieldError("NeedGold", err)
	}
	// 验证 duration
	err = validator.MinValidate(float64(t.Duration), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Duration)
		return template.NewTemplateFieldError("Duration", err)
	}
	// 验证 jihuo_rew_silver
	err = validator.MinValidate(float64(t.JiHuoRewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JiHuoRewSilver)
		return template.NewTemplateFieldError("JiHuoRewSilver", err)
	}
	// 验证 jihuo_rew_gold
	err = validator.MinValidate(float64(t.JiHuoRewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JiHuoRewGold)
		return template.NewTemplateFieldError("JiHuoRewGold", err)
	}
	// 验证 jihuo_rew_bind_gold
	err = validator.MinValidate(float64(t.JiHuoRewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JiHuoRewBindGold)
		return template.NewTemplateFieldError("JiHuoRewBindGold", err)
	}
	// 验证 day_rew_silver
	err = validator.MinValidate(float64(t.DayRewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DayRewSilver)
		return template.NewTemplateFieldError("DayRewSilver", err)
	}
	// 验证 day_rew_gold
	err = validator.MinValidate(float64(t.DayRewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DayRewGold)
		return template.NewTemplateFieldError("DayRewGold", err)
	}
	// 验证 day_rew_bind_gold
	err = validator.MinValidate(float64(t.DayRewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DayRewBindGold)
		return template.NewTemplateFieldError("DayRewBindGold", err)
	}
	// 验证 silver_xianfu_free_add
	err = validator.MinValidate(float64(t.SilverXianFuFreeAdd), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SilverXianFuFreeAdd)
		return template.NewTemplateFieldError("SilverXianFuFreeAdd", err)
	}
	// 验证 exp_xianfu_free_add
	err = validator.MinValidate(float64(t.ExpXianFuFreeAdd), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ExpXianFuFreeAdd)
		return template.NewTemplateFieldError("ExpXianFuFreeAdd", err)
	}
	// 验证 exp_biology_add_percent
	err = validator.RangeValidate(float64(t.ExpBiologyAddPercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ExpBiologyAddPercent)
		return template.NewTemplateFieldError("Rate", err)
	}
	// 验证 mount_attr_add_percent
	err = validator.RangeValidate(float64(t.MountAttrAddPercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MountAttrAddPercent)
		return template.NewTemplateFieldError("Rate", err)
	}
	// 验证 tianjie_boss_free_add
	err = validator.MinValidate(float64(t.TianJieBossFreeAdd), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TianJieBossFreeAdd)
		return template.NewTemplateFieldError("TianJieBossFreeAdd", err)
	}
	// 验证 3v3_jifen_add_percent
	err = validator.RangeValidate(float64(t.JiFenAddPercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JiFenAddPercent)
		return template.NewTemplateFieldError("JiFenAddPercent", err)
	}
	// 验证 3v3_jifen_max_add_percent
	err = validator.RangeValidate(float64(t.JiFenMaxAddPercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JiFenMaxAddPercent)
		return template.NewTemplateFieldError("JiFenMaxAddPercent", err)
	}
	// 验证 wing_attr_add_percent
	err = validator.RangeValidate(float64(t.WingAttrAddPercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WingAttrAddPercent)
		return template.NewTemplateFieldError("WingAttrAddPercent", err)
	}

	return nil
}

func (t *XianZunCardTemplate) PatchAfterCheck() {
	t.emailItemMap = make(map[int32]int32)
	utils.MergeMap(t.emailItemMap, t.dayReceiveItemMap)
	if t.DayRewSilver != 0 {
		t.emailItemMap[constanttypes.SilverItem] = t.DayRewSilver
	}
	if t.DayRewBindGold != 0 {
		t.emailItemMap[constanttypes.BindGoldItem] = t.DayRewBindGold
	}
	if t.DayRewGold != 0 {
		t.emailItemMap[constanttypes.GoldItem] = t.DayRewGold
	}

}

func (t *XianZunCardTemplate) FileName() string {
	return "tb_tequan_xianzun.json"
}

func init() {
	template.Register((*XianZunCardTemplate)(nil))
}
