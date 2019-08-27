package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

//兵魂升星配置
type WeaponUpstarTemplate struct {
	*WeaponUpstarTemplateVO
	needItemMap              map[playertypes.RoleType]map[int32]int32 //升星需要物品
	nextWeaponUpstarTemplate *WeaponUpstarTemplate
	useItemTemplate          map[playertypes.RoleType]*ItemTemplate //check校验使用
}

func (wut *WeaponUpstarTemplate) TemplateId() int {
	return wut.Id
}

func (wut *WeaponUpstarTemplate) GetNeedItemMap(roleType playertypes.RoleType) map[int32]int32 {
	return wut.needItemMap[roleType]
}

func (wut *WeaponUpstarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wut.FileName(), wut.TemplateId(), err)
			return
		}
	}()

	wut.needItemMap = make(map[playertypes.RoleType]map[int32]int32)
	wut.useItemTemplate = make(map[playertypes.RoleType]*ItemTemplate)
	//验证 upstar_item_id(开天)
	if wut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(wut.UpstarItemId), (*ItemTemplate)(nil))
		if to == nil {

			err = fmt.Errorf("[%d] invalid", wut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}

		wut.useItemTemplate[playertypes.RoleTypeKaiTian] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(wut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", wut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}
		kaiTianNeddItemMap, exist := wut.needItemMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianNeddItemMap = make(map[int32]int32)
			wut.needItemMap[playertypes.RoleTypeKaiTian] = kaiTianNeddItemMap
		}
		kaiTianNeddItemMap[wut.UpstarItemId] = wut.UpstarItemCount
	}

	//验证 upstar_item_id2(奕剑)
	if wut.UpstarItemId2 != 0 {
		to := template.GetTemplateService().Get(int(wut.UpstarItemId2), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wut.UpstarItemId2)
			return template.NewTemplateFieldError("UpstarItemId2", err)
		}

		wut.useItemTemplate[playertypes.RoleTypeYiJian] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(wut.UpstarItemCount2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", wut.UpstarItemCount2)
			return template.NewTemplateFieldError("UpstarItemCount2", err)
		}
		yiJianNeddItemMap, exist := wut.needItemMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianNeddItemMap = make(map[int32]int32)
			wut.needItemMap[playertypes.RoleTypeYiJian] = yiJianNeddItemMap
		}
		yiJianNeddItemMap[wut.UpstarItemId2] = wut.UpstarItemCount2
	}

	//验证 upstar_item_id3(破月)
	if wut.UpstarItemId3 != 0 {
		to := template.GetTemplateService().Get(int(wut.UpstarItemId3), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wut.UpstarItemId3)
			return template.NewTemplateFieldError("UpstarItemId3", err)
		}

		wut.useItemTemplate[playertypes.RoleTypePoYue] = to.(*ItemTemplate)

		err = validator.MinValidate(float64(wut.UpstarItemCount3), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", wut.UpstarItemCount3)
			return template.NewTemplateFieldError("UpstarItemCount3", err)
		}
		poYueNeddItemMap, exist := wut.needItemMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueNeddItemMap = make(map[int32]int32)
			wut.needItemMap[playertypes.RoleTypePoYue] = poYueNeddItemMap
		}
		poYueNeddItemMap[wut.UpstarItemId3] = wut.UpstarItemCount3
	}

	//验证 next_id
	if wut.NextId != 0 {
		to := template.GetTemplateService().Get(int(wut.NextId), (*WeaponUpstarTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*WeaponUpstarTemplate)

			diffLevel := nextTemplate.Level - wut.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			wut.nextWeaponUpstarTemplate = nextTemplate
		}
	}

	return nil
}

func (wut *WeaponUpstarTemplate) PatchAfterCheck() {

}

func (wut *WeaponUpstarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wut.FileName(), wut.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(wut.UpstarRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.UpstarRate)
		err = template.NewTemplateFieldError("UpstarRate", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(wut.TimesMin), float64(0), true, float64(wut.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(wut.TimesMax), float64(wut.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(wut.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 level
	err = validator.MinValidate(float64(wut.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(wut.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(wut.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(wut.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	// for roleType, useItemTemplate := range wut.useItemTemplate {
	// 	if useItemTemplate != nil {
	// 		if useItemTemplate.GetItemSubType() != itemtypes.ItemSoulSubTypeDebris {
	// 			err = fmt.Errorf("UpstarItemId role [%d] invalid", int32(roleType))
	// 			return template.NewTemplateFieldError("", err)
	// 		}
	// 	}
	// }

	return nil
}

func (wut *WeaponUpstarTemplate) FileName() string {
	return "tb_weapon_upstar.json"
}

func init() {
	template.Register((*WeaponUpstarTemplate)(nil))
}
