package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	itemtypes "fgame/fgame/game/item/types"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//称号定制配置
type TitleDingZhiTemplate struct {
	*TitleDingZhiTemplateVO
	needItemMap       map[int32]int32 //激活需要物品
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	buffTemplate      *BuffTemplate
}

func (tt *TitleDingZhiTemplate) TemplateId() int {
	return tt.Id
}

func (tt *TitleDingZhiTemplate) GetNeedItemMap() map[int32]int32 {
	return tt.needItemMap
}

func (tt *TitleDingZhiTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return tt.battlePropertyMap
}

func (tt *TitleDingZhiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//属性
	tt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	tt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(tt.Hp)
	tt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(tt.Attack)
	tt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(tt.Defence)

	//验证 NeedItemId
	tt.needItemMap = make(map[int32]int32)
	if tt.NeedItemId != 0 {
		tt.needItemMap[tt.NeedItemId] = tt.NeedItemCount
	}

	if tt.BuffId != 0 {
		buffTemplateVO := template.GetTemplateService().Get(int(tt.BuffId), (*BuffTemplate)(nil))
		if buffTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", tt.BuffId)
			err = template.NewTemplateFieldError("BuffId", err)
			return
		}

		tt.buffTemplate = buffTemplateVO.(*BuffTemplate)

	}

	return nil
}

func (tt *TitleDingZhiTemplate) PatchAfterCheck() {

}

func (tt *TitleDingZhiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	for itemId, num := range tt.needItemMap {
		needItemTemplateVO := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", tt.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return
		}

		itemTemplate := needItemTemplateVO.(*ItemTemplate)
		if itemTemplate.GetItemSubType() != itemtypes.ItemTitleSubTypeDingZhiCard {
			err = template.NewTemplateFieldError("NeedItemId", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
	}

	if tt.buffTemplate != nil {
		if tt.buffTemplate.GetBuffType() != scenetypes.BuffTypeTitleDingZhi {
			err = fmt.Errorf("[%d] invalid", tt.BuffId)
			err = template.NewTemplateFieldError("BuffId", err)
			return
		}
	}

	return nil
}

func (tt *TitleDingZhiTemplate) FileName() string {
	return "tb_title_dingzhi.json"
}

func init() {
	template.Register((*TitleDingZhiTemplate)(nil))
}
