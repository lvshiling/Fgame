package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

func init() {
	template.Register((*GiftCodeTemplate)(nil))
}

//礼包激活配置
type GiftCodeTemplate struct {
	*GiftCodeTemplateVO
	giftItemMap map[int32]int32
}

func (t *GiftCodeTemplate) TemplateId() int {
	return t.Id
}

func (t *GiftCodeTemplate) GetItemMap() map[int32]int32 {
	return t.giftItemMap
}

func (t *GiftCodeTemplate) FileName() string {
	return "tb_duihuan.json"
}

//检查有效性
func (t *GiftCodeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//组合成需要的数据
func (t *GiftCodeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证: 扫荡所需的物品ID,逗号隔开
	t.giftItemMap = make(map[int32]int32)

	intGiftItemIdArr, err := utils.SplitAsIntArray(t.ItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}
	intGiftItemNumArr, err := utils.SplitAsIntArray(t.ItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemCount)
		return template.NewTemplateFieldError("ItemCount", err)
	}
	if len(intGiftItemIdArr) != len(intGiftItemNumArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.ItemId, t.ItemCount)
		err = template.NewTemplateFieldError("ItemId or ItemCount", err)
		return err
	}
	if len(intGiftItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intGiftItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("ItemId", fmt.Errorf("[%s] invalid", t.ItemId))
			}

			err = validator.MinValidate(float64(intGiftItemNumArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("ItemCount", err)
			}

			t.giftItemMap[itemId] = intGiftItemNumArr[index]
		}
	}

	return nil
}

//检验后组合
func (t *GiftCodeTemplate) PatchAfterCheck() {
}
