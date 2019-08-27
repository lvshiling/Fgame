package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//运营活动次数配置
type TimesRewTemplate struct {
	*TimesRewTemplateVO
	rewItemMap      map[int32]int32 //奖励物品
	emailRewItemMap map[int32]int32 //邮件
}

func (t *TimesRewTemplate) TemplateId() int {
	return t.Id
}

func (t *TimesRewTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *TimesRewTemplate) GetEmailRewItemMap() map[int32]int32 {
	return t.emailRewItemMap
}

func (t *TimesRewTemplate) PatchAfterCheck() {
}

func (t *TimesRewTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//奖励物品id
	t.rewItemMap = make(map[int32]int32)
	t.emailRewItemMap = make(map[int32]int32)
	itemIdList, err := utils.SplitAsIntArray(t.RawId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RawId)
		return template.NewTemplateFieldError("RawId", err)
	}
	itemCountList, err := utils.SplitAsIntArray(t.RawCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RawCount)
		return template.NewTemplateFieldError("RawCount", err)
	}
	if len(itemIdList) != len(itemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.RawId, t.RawCount)
		return template.NewTemplateFieldError("RawId or RawCount", err)
	}
	if len(itemIdList) > 0 {
		//组合数据
		for index, itemId := range itemIdList {
			_, ok := t.rewItemMap[itemId]
			if ok {
				t.rewItemMap[itemId] += itemCountList[index]
			} else {
				t.rewItemMap[itemId] = itemCountList[index]
			}

			// _, ok = t.emailRewItemMap[itemId]
			// if ok {
			// 	t.emailRewItemMap[itemId] += itemCountList[index]
			// } else {
			// 	t.emailRewItemMap[itemId] = itemCountList[index]
			// }
		}
	}

	return nil
}

func (t *TimesRewTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 活动id
	err = validator.MinValidate(float64(t.Group), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Group)
		err = template.NewTemplateFieldError("Group", err)
		return
	}

	//验证 次数
	err = validator.MinValidate(float64(t.DrawTimes), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DrawTimes)
		err = template.NewTemplateFieldError("DrawTimes", err)
		return
	}

	//验证 vip
	err = validator.MinValidate(float64(t.VipLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.VipLevel)
		err = template.NewTemplateFieldError("VipLevel", err)
		return
	}

	//验证  物品
	for itemId, num := range t.rewItemMap {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("FreeGiftId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("FreeGiftCount", err)
			return
		}
	}

	return nil
}

func (t *TimesRewTemplate) FileName() string {
	return "tb_leichong.json"
}

func init() {
	template.Register((*TimesRewTemplate)(nil))
}
