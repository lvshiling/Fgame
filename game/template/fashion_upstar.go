package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

//时装升星配置
type FashionUpstarTemplate struct {
	*FashionUpstarTemplateVO
	needItemMap               map[playertypes.RoleType]map[playertypes.SexType]map[int32]int32 //升星需要物品
	nextFashionUpstarTemplate *FashionUpstarTemplate
	useItemTemplateMap        map[playertypes.RoleType]map[playertypes.SexType]*ItemTemplate
}

func (fut *FashionUpstarTemplate) TemplateId() int {
	return fut.Id
}

func (fut *FashionUpstarTemplate) GetNeedItemMap(roleType playertypes.RoleType, sexType playertypes.SexType) map[int32]int32 {
	sexTypeMap, exist := fut.needItemMap[roleType]
	if !exist {
		return nil
	}
	return sexTypeMap[sexType]
}

func (fut *FashionUpstarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fut.FileName(), fut.TemplateId(), err)
			return
		}
	}()

	fut.needItemMap = make(map[playertypes.RoleType]map[playertypes.SexType]map[int32]int32)
	fut.useItemTemplateMap = make(map[playertypes.RoleType]map[playertypes.SexType]*ItemTemplate)
	//验证 upstar_item_id (开天男)
	if fut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(fut.UpstarItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}

		kaiTianUseItemMap, exist := fut.useItemTemplateMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianUseItemMap = make(map[playertypes.SexType]*ItemTemplate)
			fut.useItemTemplateMap[playertypes.RoleTypeKaiTian] = kaiTianUseItemMap
		}
		kaiTianUseItemMap[playertypes.SexTypeMan] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(fut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}

		kaiTianNeedMap, exist := fut.needItemMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianNeedMap = make(map[playertypes.SexType]map[int32]int32)
			fut.needItemMap[playertypes.RoleTypeKaiTian] = kaiTianNeedMap
		}
		needMap, exist := kaiTianNeedMap[playertypes.SexTypeMan]
		if !exist {
			needMap = make(map[int32]int32)
			kaiTianNeedMap[playertypes.SexTypeMan] = needMap
		}
		needMap[fut.UpstarItemId] = fut.UpstarItemCount
	}

	//验证 upstar_item_id_nv (开天女)
	if fut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(fut.UpstarItemIdNv), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemIdNv)
			return template.NewTemplateFieldError("UpstarItemIdNv", err)
		}

		kaiTianUseItemMap, exist := fut.useItemTemplateMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianUseItemMap = make(map[playertypes.SexType]*ItemTemplate)
			fut.useItemTemplateMap[playertypes.RoleTypeKaiTian] = kaiTianUseItemMap
		}
		kaiTianUseItemMap[playertypes.SexTypeWoman] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(fut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}

		kaiTianNeedMap, exist := fut.needItemMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianNeedMap = make(map[playertypes.SexType]map[int32]int32)
			fut.needItemMap[playertypes.RoleTypeKaiTian] = kaiTianNeedMap
		}
		needMap, exist := kaiTianNeedMap[playertypes.SexTypeWoman]
		if !exist {
			needMap = make(map[int32]int32)
			kaiTianNeedMap[playertypes.SexTypeWoman] = needMap
		}
		needMap[fut.UpstarItemIdNv] = fut.UpstarItemCount
	}

	//验证 upstar_item_id2 (奕剑男)
	if fut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(fut.UpstarItemId2), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemId2)
			return template.NewTemplateFieldError("UpstarItemId2", err)
		}

		yiJianUseItemMap, exist := fut.useItemTemplateMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianUseItemMap = make(map[playertypes.SexType]*ItemTemplate)
			fut.useItemTemplateMap[playertypes.RoleTypeYiJian] = yiJianUseItemMap
		}
		yiJianUseItemMap[playertypes.SexTypeMan] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(fut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}

		yiJianNeedMap, exist := fut.needItemMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianNeedMap = make(map[playertypes.SexType]map[int32]int32)
			fut.needItemMap[playertypes.RoleTypeYiJian] = yiJianNeedMap
		}
		needMap, exist := yiJianNeedMap[playertypes.SexTypeMan]
		if !exist {
			needMap = make(map[int32]int32)
			yiJianNeedMap[playertypes.SexTypeMan] = needMap
		}
		needMap[fut.UpstarItemId2] = fut.UpstarItemCount
	}

	//验证 upstar_item_id2_nv (奕剑女)
	if fut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(fut.UpstarItemId2Nv), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemId2Nv)
			return template.NewTemplateFieldError("UpstarItemId2Nv", err)
		}

		yiJianUseItemMap, exist := fut.useItemTemplateMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianUseItemMap = make(map[playertypes.SexType]*ItemTemplate)
			fut.useItemTemplateMap[playertypes.RoleTypeYiJian] = yiJianUseItemMap
		}
		yiJianUseItemMap[playertypes.SexTypeWoman] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(fut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}

		yiJianNeedMap, exist := fut.needItemMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianNeedMap = make(map[playertypes.SexType]map[int32]int32)
			fut.needItemMap[playertypes.RoleTypeYiJian] = yiJianNeedMap
		}
		needMap, exist := yiJianNeedMap[playertypes.SexTypeWoman]
		if !exist {
			needMap = make(map[int32]int32)
			yiJianNeedMap[playertypes.SexTypeWoman] = needMap
		}
		needMap[fut.UpstarItemId2Nv] = fut.UpstarItemCount
	}

	//验证 upstar_item_id3 (破月男)
	if fut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(fut.UpstarItemId3), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemId3)
			return template.NewTemplateFieldError("UpstarItemId3", err)
		}

		poYueUseItemMap, exist := fut.useItemTemplateMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueUseItemMap = make(map[playertypes.SexType]*ItemTemplate)
			fut.useItemTemplateMap[playertypes.RoleTypePoYue] = poYueUseItemMap
		}
		poYueUseItemMap[playertypes.SexTypeMan] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(fut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}

		poYueNeedMap, exist := fut.needItemMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueNeedMap = make(map[playertypes.SexType]map[int32]int32)
			fut.needItemMap[playertypes.RoleTypePoYue] = poYueNeedMap
		}
		needMap, exist := poYueNeedMap[playertypes.SexTypeMan]
		if !exist {
			needMap = make(map[int32]int32)
			poYueNeedMap[playertypes.SexTypeMan] = needMap
		}
		needMap[fut.UpstarItemId3] = fut.UpstarItemCount
	}

	//验证 upstar_item_id3_nv (破月女)
	if fut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(fut.UpstarItemId3Nv), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemId3Nv)
			return template.NewTemplateFieldError("UpstarItemId3Nv", err)
		}

		poYueUseItemMap, exist := fut.useItemTemplateMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueUseItemMap = make(map[playertypes.SexType]*ItemTemplate)
			fut.useItemTemplateMap[playertypes.RoleTypePoYue] = poYueUseItemMap
		}
		poYueUseItemMap[playertypes.SexTypeWoman] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(fut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}

		poYueNeedMap, exist := fut.needItemMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueNeedMap = make(map[playertypes.SexType]map[int32]int32)
			fut.needItemMap[playertypes.RoleTypePoYue] = poYueNeedMap
		}
		needMap, exist := poYueNeedMap[playertypes.SexTypeWoman]
		if !exist {
			needMap = make(map[int32]int32)
			poYueNeedMap[playertypes.SexTypeWoman] = needMap
		}
		needMap[fut.UpstarItemId3Nv] = fut.UpstarItemCount
	}

	//验证 next_id
	if fut.NextId != 0 {
		to := template.GetTemplateService().Get(int(fut.NextId), (*FashionUpstarTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*FashionUpstarTemplate)
			diffLevel := nextTemplate.Level - fut.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			fut.nextFashionUpstarTemplate = nextTemplate
		}
	}

	return nil
}

func (fut *FashionUpstarTemplate) PatchAfterCheck() {

}

func (fut *FashionUpstarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fut.FileName(), fut.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(fut.UpstarRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.UpstarRate)
		err = template.NewTemplateFieldError("UpstarRate", err)
		return
	}

	//验证 equip_percent
	err = validator.RangeValidate(float64(fut.EquipPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.EquipPercent)
		err = template.NewTemplateFieldError("EquipPercent", err)
		return
	}

	//验证 level
	err = validator.MinValidate(float64(fut.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(fut.Hp), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(fut.Attack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(fut.Defence), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	for _, useItemRoleTempalteMap := range fut.useItemTemplateMap {
		for _, useItemTemplate := range useItemRoleTempalteMap {
			if useItemTemplate == nil {
				continue
			}
			if useItemTemplate.GetItemType() != itemtypes.ItemTypeFashion {
				err = fmt.Errorf("UpstarItemId [%d]  invalid", fut.UpstarItemId)
				return template.NewTemplateFieldError("UpstarItemId", err)
			}
		}
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(fut.TimesMin), float64(0), true, float64(fut.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(fut.TimesMax), float64(fut.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(fut.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(fut.AddMin), float64(0), true, float64(fut.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(fut.AddMax), float64(fut.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(fut.ZhufuMax), float64(fut.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return nil
}

func (fut *FashionUpstarTemplate) FileName() string {
	return "tb_fashion_upstar.json"
}

func init() {
	template.Register((*FashionUpstarTemplate)(nil))
}
