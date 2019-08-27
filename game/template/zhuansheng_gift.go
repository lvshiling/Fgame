package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

//转生大礼包
type ZhuanShengGiftTemplate struct {
	*ZhuanShengGiftTemplateVO
	itemMap     map[int32]int32
	useItemMap  map[int32]int32
	giftItemMap map[int32]int32
	role        playertypes.RoleType
	sex         playertypes.SexType
}

func (t *ZhuanShengGiftTemplate) TemplateId() int {
	return t.Id
}

func (t *ZhuanShengGiftTemplate) GetUseItemMap() map[int32]int32 {
	return t.useItemMap
}

func (t *ZhuanShengGiftTemplate) GetGiftItemMap() map[int32]int32 {
	return t.giftItemMap
}

func (t *ZhuanShengGiftTemplate) GetRole() playertypes.RoleType {
	return t.role
}

func (t *ZhuanShengGiftTemplate) GetSex() playertypes.SexType {
	return t.sex
}

func (t *ZhuanShengGiftTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *ZhuanShengGiftTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//消耗物品
	t.useItemMap = make(map[int32]int32)
	itemArr, err := utils.SplitAsIntArray(t.UseItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}
	itemNumArr, err := utils.SplitAsIntArray(t.UseItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemCount)
		return template.NewTemplateFieldError("UseItemCount", err)
	}
	if len(itemArr) != len(itemNumArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.UseItemId, t.ItemCount)
		return template.NewTemplateFieldError("UseItemId or UseItemCount", err)
	}
	//组合数据
	if len(itemArr) > 0 {
		for index, itemId := range itemArr {
			t.useItemMap[itemId] += itemNumArr[index]
		}
	}

	//折扣物品
	t.itemMap = make(map[int32]int32)
	rewItemIdList, err := utils.SplitAsIntArray(t.ItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.ItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemCount)
		return template.NewTemplateFieldError("ItemCount", err)
	}
	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.ItemId, t.ItemCount)
		return template.NewTemplateFieldError("ItemId or ItemCount", err)
	}
	if len(rewItemIdList) > 0 {
		//组合数据
		for index, itemId := range rewItemIdList {
			_, ok := t.itemMap[itemId]
			if ok {
				t.itemMap[itemId] += rewItemCountList[index]
			} else {
				t.itemMap[itemId] = rewItemCountList[index]
			}
		}
	}

	//赠品物品
	t.giftItemMap = make(map[int32]int32)
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
		return template.NewTemplateFieldError("GiveItemId or GiveItemCount", err)
	}

	//组合数据
	if len(giveItemArr) > 0 {
		for index, itemId := range giveItemArr {
			t.giftItemMap[itemId] += giveItemNumArr[index]
		}
	}

	return nil
}

func (t *ZhuanShengGiftTemplate) PatchAfterCheck() {

}

func (t *ZhuanShengGiftTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 buyCount（废弃）
	err = validator.MinValidate(float64(t.BuyCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BuyCount)
		err = template.NewTemplateFieldError("BuyCount", err)
		return
	}

	//验证 maxCount（废弃）
	err = validator.MinValidate(float64(t.MaxCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MaxCount)
		err = template.NewTemplateFieldError("MaxCount", err)
		return
	}

	// 原价
	err = validator.MinValidate(float64(t.YuanGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("YuanGold", err)
	}

	//现价
	err = validator.MinValidate(float64(t.UseGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UseGold", err)
	}

	//现积分价
	err = validator.MinValidate(float64(t.UsePoint), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UsePoint", err)
	}

	//充值限制
	err = validator.MinValidate(float64(t.NeedChongZhi), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("NeedChongZhi", err)
	}

	//最大总购买数
	err = validator.MinValidate(float64(t.BuyMax), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BuyMax", err)
	}

	//批量购买折扣
	err = validator.MinValidate(float64(t.Bargain), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Bargain", err)
	}

	//验证 礼包物品
	for itemId, num := range t.itemMap {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("ItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//角色
	if t.Profession != 0 {
		t.role = playertypes.RoleType(t.Profession)
		if !t.role.Valid() {
			err = template.NewTemplateFieldError("Profession", err)
			return
		}
	}

	//性别
	if t.Gender != 0 {
		t.sex = playertypes.SexType(t.Gender)
		if !t.sex.Valid() {
			err = template.NewTemplateFieldError("Gender", err)
			return
		}
	}

	//校验物品
	for itemId, num := range t.useItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("UseItemId", fmt.Errorf("[%s] invalid", t.UseItemId))
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("UseItemCount", fmt.Errorf("[%s] invalid", t.UseItemCount))
		}
	}

	//校验赠送物品
	for itemId, num := range t.giftItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("GiveItemId", fmt.Errorf("[%s] invalid", t.GiveItemId))
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("GiveItemCount", fmt.Errorf("[%s] invalid", t.GiveItemCount))
		}
	}

	return nil
}

func (t *ZhuanShengGiftTemplate) FileName() string {
	return "tb_zhuansheng_gift.json"
}

func init() {
	template.Register((*ZhuanShengGiftTemplate)(nil))
}
