package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	centertypes "fgame/fgame/game/center/types"
	constanttypes "fgame/fgame/game/constant/types"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fmt"
)

//会员特权配置
type HuiYuanTemplate struct {
	*HuiYuanTemplateVO
	huiyuanType          huiyuantypes.HuiYuanType
	rewItemMap           map[int32]int32
	firstRewItemMap      map[int32]int32
	emailRewItemMap      map[int32]int32
	emailFirstRewItemMap map[int32]int32
	houtaiType           centertypes.ZhiZunType
}

func (t *HuiYuanTemplate) TemplateId() int {
	return t.Id
}

func (t *HuiYuanTemplate) GetHuiYuanType() huiyuantypes.HuiYuanType {
	return t.huiyuanType
}

func (t *HuiYuanTemplate) GetHoutaiType() centertypes.ZhiZunType {
	return t.houtaiType
}

func (t *HuiYuanTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *HuiYuanTemplate) GetRewFirstItemMap() map[int32]int32 {
	return t.firstRewItemMap
}

func (t *HuiYuanTemplate) GetEmailRewItemMap() map[int32]int32 {
	return t.emailRewItemMap
}

func (t *HuiYuanTemplate) GetEmailFirstRewItemMap() map[int32]int32 {
	return t.emailFirstRewItemMap
}

func (t *HuiYuanTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewItemMap = make(map[int32]int32)
	t.emailRewItemMap = make(map[int32]int32)
	//验证 rew_item_id
	rewItemIdList, err := utils.SplitAsIntArray(t.RewItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewItem)
		return template.NewTemplateFieldError("RewItem", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.RewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewItemCount)
		return template.NewTemplateFieldError("RewItemCount", err)
	}
	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.RewItem, t.RewItemCount)
		return template.NewTemplateFieldError("RewItem or RewItemCount", err)
	}
	if len(rewItemIdList) > 0 {
		//组合数据
		for index, itemId := range rewItemIdList {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.RewItem)
				return template.NewTemplateFieldError("RewItem", err)
			}

			err = validator.MinValidate(float64(rewItemCountList[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("RewItemCount", err)
			}

			t.rewItemMap[itemId] = rewItemCountList[index]
			t.emailRewItemMap[itemId] = rewItemCountList[index]
		}
	}

	// 第一次特殊奖励
	t.firstRewItemMap = make(map[int32]int32)
	t.emailFirstRewItemMap = make(map[int32]int32)
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
	if len(getItemIdList) > 0 {
		//组合数据
		for index, itemId := range getItemIdList {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.RewItem)
				return template.NewTemplateFieldError("RewItem", err)
			}

			err = validator.MinValidate(float64(getItemCountList[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("RewItemCount", err)
			}

			t.firstRewItemMap[itemId] = getItemCountList[index]
			t.emailFirstRewItemMap[itemId] = getItemCountList[index]
		}
	}

	return nil
}

func (t *HuiYuanTemplate) PatchAfterCheck() {
	if t.RewSilver != 0 {
		t.emailRewItemMap[constanttypes.SilverItem] = t.RewSilver
		t.emailFirstRewItemMap[constanttypes.SilverItem] = t.RewSilver
	}
	if t.RewBindGold != 0 {
		t.emailRewItemMap[constanttypes.BindGoldItem] = t.RewBindGold
		t.emailFirstRewItemMap[constanttypes.BindGoldItem] = t.RewBindGold
	}
	if t.RewGold != 0 {
		t.emailRewItemMap[constanttypes.GoldItem] = t.RewGold
		t.emailFirstRewItemMap[constanttypes.GoldItem] = t.RewGold
	}
}

func (t *HuiYuanTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//等级
	t.huiyuanType = huiyuantypes.HuiYuanType(t.Level)
	if !t.huiyuanType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//购买所需元宝
	err = validator.MinValidate(float64(t.NeedGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedGold)
		return template.NewTemplateFieldError("NeedGold", err)
	}

	// 持续时间
	err = validator.MinValidate(float64(t.Duration), float64(0), true)
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

	// 屠魔特权
	err = validator.MinValidate(float64(t.TumoFour), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TumoFour)
		return template.NewTemplateFieldError("TumoFour", err)
	}

	// 帝陵特权
	err = validator.MinValidate(float64(t.DilingCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DilingCount)
		return template.NewTemplateFieldError("DilingCount", err)
	}

	// 后台版本
	t.houtaiType = centertypes.ZhiZunType(t.HoutaiType)
	if !t.houtaiType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.HoutaiType)
		return template.NewTemplateFieldError("HoutaiType", err)
	}

	return nil
}

func (edt *HuiYuanTemplate) FileName() string {
	return "tb_tequan.json"
}

func init() {
	template.Register((*HuiYuanTemplate)(nil))
}
