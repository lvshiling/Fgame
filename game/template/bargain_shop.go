package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

//打折礼包
type BargainShopTemplate struct {
	*BargainShopTemplateVO
	itemMap          map[int32]int32
	role             playertypes.RoleType
	sex              playertypes.SexType
	startBargainTemp *BargainTemplate
	bargainTempMap   map[int32]*BargainTemplate
}

func (t *BargainShopTemplate) TemplateId() int {
	return t.Id
}

func (t *BargainShopTemplate) GetDiscount(kanjiaTimes int32) int32 {
	bargainTemp, ok := t.bargainTempMap[kanjiaTimes]
	if !ok {
		return -1
	}

	return bargainTemp.RandomDaZhe()
}

func (t *BargainShopTemplate) GetRole() playertypes.RoleType {
	return t.role
}

func (t *BargainShopTemplate) GetSex() playertypes.SexType {
	return t.sex
}

func (t *BargainShopTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *BargainShopTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

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

	return nil
}

func (t *BargainShopTemplate) PatchAfterCheck() {
	t.bargainTempMap = make(map[int32]*BargainTemplate)
	for startTemp := t.startBargainTemp; startTemp != nil; startTemp = startTemp.nextTemp {
		t.bargainTempMap[startTemp.BargainTimes] = startTemp
	}
}

func (t *BargainShopTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 原价
	err = validator.MinValidate(float64(t.YuanGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("YuanGold", err)
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

	//打折配置
	tempObj := template.GetTemplateService().Get(int(t.BargainBegain), (*BargainTemplate)(nil))
	if tempObj == nil {
		err = fmt.Errorf("[%s] invalid", t.BargainBegain)
		return template.NewTemplateFieldError("BargainBegain", err)
	}
	t.startBargainTemp = tempObj.(*BargainTemplate)

	return nil
}

func (t *BargainShopTemplate) FileName() string {
	return "tb_bargain_shop.json"
}

func init() {
	template.Register((*BargainShopTemplate)(nil))
}
