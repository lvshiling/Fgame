package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	skilltypes "fgame/fgame/game/skill/types"
	soultypes "fgame/fgame/game/soul/types"
	"fmt"
)

//帝魂觉醒配置
type SoulAwakenTemplate struct {
	*SoulAwakenTemplateVO
	kindType       soultypes.SoulKindType //帝魂种类
	soulType       soultypes.SoulType     //帝魂类型
	needItemMap    map[int32]int32        //觉醒物品
	upLevelItemMap map[int32]int32        //升级需要物品
}

func (sat *SoulAwakenTemplate) TemplateId() int {
	return sat.Id
}

func (sat *SoulAwakenTemplate) GetNeedItemMap() map[int32]int32 {
	return sat.needItemMap
}

func (sat *SoulAwakenTemplate) GetUpLevelItemMap() map[int32]int32 {
	return sat.upLevelItemMap
}

func (sat *SoulAwakenTemplate) GetSoulType() soultypes.SoulType {
	return sat.soulType
}

func (sat *SoulAwakenTemplate) PatchAfterCheck() {

}

func (sat *SoulAwakenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(sat.FileName(), sat.TemplateId(), err)
			return
		}
	}()
	//type
	sat.kindType = soultypes.SoulKindType(sat.Type)
	if !sat.kindType.Valid() {
		err = fmt.Errorf("[%d] invalid", sat.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//soul_typ
	sat.soulType = soultypes.SoulType(sat.SoulType)
	if !sat.soulType.Valid() {
		err = fmt.Errorf("[%d] invalid", sat.SoulType)
		err = template.NewTemplateFieldError("SoulType", err)
		return
	}

	//need_item_id
	if sat.NeedItemId != "" {
		itemArr, err := utils.SplitAsIntArray(sat.NeedItemId)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(sat.NeedItemCount)
		if err != nil {
			return err
		}
		if len(itemArr) <= 0 || len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", sat.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return err
		}
		sat.needItemMap = make(map[int32]int32)
		for i := 0; i < len(itemArr); i++ {
			sat.needItemMap[itemArr[i]] = numArr[i]
		}
	}

	//uplevel_needitem
	if sat.UplevelNeeditem != "" {
		itemArr, err := utils.SplitAsIntArray(sat.UplevelNeeditem)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(sat.UplevelItemCount)
		if err != nil {
			return err
		}
		if len(itemArr) <= 0 || len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", sat.UplevelNeeditem)
			err = template.NewTemplateFieldError("UplevelNeeditem", err)
			return err
		}
		sat.upLevelItemMap = make(map[int32]int32)
		for i := 0; i < len(itemArr); i++ {
			sat.upLevelItemMap[itemArr[i]] = numArr[i]
		}
	}

	return nil
}

func (sat *SoulAwakenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(sat.FileName(), sat.TemplateId(), err)
			return
		}
	}()

	for itemId, num := range sat.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", sat.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return err
		}
		//need_item_count
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", sat.NeedItemCount)
			return template.NewTemplateFieldError("NeedItemCount", err)
		}
	}

	for itemId, num := range sat.upLevelItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", sat.UplevelNeeditem)
			err = template.NewTemplateFieldError("UplevelNeeditem", err)
			return err
		}
		//uplevel_item_count
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", sat.UplevelItemCount)
			return template.NewTemplateFieldError("UplevelItemCount", err)
		}
	}

	//验证 order
	err = validator.MinValidate(float64(sat.Order), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sat.Order)
		return template.NewTemplateFieldError("Order", err)
	}

	//验证 next_id
	if sat.NextId != 0 {
		diff := sat.NextId - int32(sat.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", sat.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return err
		}
		to := template.GetTemplateService().Get(int(sat.NextId), (*SoulAwakenTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", sat.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		//验证 order
		sato := to.(*SoulAwakenTemplate)
		diffOrder := sato.Order - sat.Order
		if diffOrder != 1 {
			err = fmt.Errorf("[%d] invalid", sato.Order)
			return template.NewTemplateFieldError("Order", err)
		}

		//验证 type
		if sato.Type != sat.Type {
			err = fmt.Errorf("[%d] invalid", sato.Type)
			return template.NewTemplateFieldError("Type", err)
		}
	}

	//验证 skill_id
	if sat.SkillId != 0 {
		to := template.GetTemplateService().Get(int(sat.SkillId), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", sat.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
		typ := to.(*SkillTemplate).GetSkillFirstType()
		if typ != skilltypes.SkillFirstTypeGuHun {
			err = fmt.Errorf("[%d] invalid", sat.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
	}

	//验证 skill_id
	if sat.UplevelSkillId != 0 {
		to := template.GetTemplateService().Get(int(sat.UplevelSkillId), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", sat.UplevelSkillId)
			return template.NewTemplateFieldError("UplevelSkillId", err)
		}
		typ := to.(*SkillTemplate).GetSkillFirstType()
		if typ != skilltypes.SkillFirstTypeGuHun {
			err = fmt.Errorf("[%d] invalid", sat.UplevelSkillId)
			return template.NewTemplateFieldError("UplevelSkillId", err)
		}
	}

	return nil
}

func (sat *SoulAwakenTemplate) FileName() string {
	return "tb_soul_awaken.json"
}

func init() {
	template.Register((*SoulAwakenTemplate)(nil))
}
