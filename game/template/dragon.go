package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//龙神配置
type DragonTemplate struct {
	*DragonTemplateVO
	needItemMap       map[int32]int32                    //物品id
	battlePropertyMap map[types.BattlePropertyType]int64 //整阶属性
}

func (dt *DragonTemplate) TemplateId() int {
	return dt.Id
}

func (dt *DragonTemplate) GetNeedItemMap() map[int32]int32 {
	return dt.needItemMap
}

func (dt *DragonTemplate) GetBattlePropertyMap() map[types.BattlePropertyType]int64 {
	return dt.battlePropertyMap
}

func (dt *DragonTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(dt.FileName(), dt.TemplateId(), err)
			return
		}
	}()

	dt.needItemMap = make(map[int32]int32)
	if dt.DragonItemId != "" {
		if dt.DragonItemId == "" {
			err = fmt.Errorf("[%s] invalid", dt.DragonItemId)
			return template.NewTemplateFieldError("DragonItemId", err)
		}

		itemArr, err := utils.SplitAsIntArray(dt.DragonItemId)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(dt.DragonItemAmount)
		if err != nil {
			return err
		}
		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", dt.DragonItemAmount)
			return template.NewTemplateFieldError("DragonItemAmount", err)
		}

		for i := 0; i < len(itemArr); i++ {
			dt.needItemMap[itemArr[i]] = numArr[i]
		}

	}

	return nil
}

func (dt *DragonTemplate) PatchAfterCheck() {
	dt.battlePropertyMap = make(map[types.BattlePropertyType]int64)
	for itemId, num := range dt.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		itemTemplate, _ := to.(*ItemTemplate)
		attrTemplate := itemTemplate.GetDragonAttrTemplate()
		for typ, val := range attrTemplate.GetAllBattleProperty() {
			total, ok := dt.battlePropertyMap[typ]
			if !ok {
				total = 0
			}
			total += val * int64(num)
			dt.battlePropertyMap[typ] = total
		}
	}

}

func (dt *DragonTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(dt.FileName(), dt.TemplateId(), err)
			return
		}
	}()

	//校验道具物品
	for itemId, num := range dt.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", dt.DragonItemId)
			return template.NewTemplateFieldError("DragonItemId", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", dt.DragonItemAmount)
			return template.NewTemplateFieldError("DragonItemAmount", err)
		}
	}

	//item_id
	if dt.ItemId != 0 {
		to := template.GetTemplateService().Get(int(dt.ItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", dt.ItemId)
			return template.NewTemplateFieldError("ItemId", err)
		}
	}

	//next_id
	if dt.NextId != 0 {
		diff := dt.NextId - int32(dt.Id)
		to := template.GetTemplateService().Get(int(dt.NextId), (*DragonTemplate)(nil))
		if to == nil || diff != 1 {
			err = fmt.Errorf("[%d] invalid", dt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	} else {
		if dt.DragonMount == 0 {
			err = fmt.Errorf("[%d] invalid", dt.DragonMount)
			return template.NewTemplateFieldError("DragonMount", err)
		}
	}

	//dragon_skill
	if dt.DragonSkill != 0 {
		to := template.GetTemplateService().Get(int(dt.DragonSkill), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", dt.DragonSkill)
			return template.NewTemplateFieldError("DragonSkill", err)
		}

		skillTemplate := to.(*SkillTemplate)
		skillType := skillTemplate.GetSkillFirstType()
		if skillType != skilltypes.SkillFirstTypeLongShen {
			err = fmt.Errorf("[%d] invalid", dt.DragonSkill)
			return template.NewTemplateFieldError("DragonSkill", err)
		}
	}

	//dragon_mount
	if dt.DragonMount != 0 {
		to := template.GetTemplateService().Get(int(dt.DragonMount), (*MountTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", dt.DragonMount)
			return template.NewTemplateFieldError("DragonMount", err)
		}
	}

	return nil
}

func (dt *DragonTemplate) FileName() string {
	return "tb_dragon.json"
}

func init() {
	template.Register((*DragonTemplate)(nil))
}
