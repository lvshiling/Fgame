package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/weapon/types"
	"fmt"
)

//兵魂配置
type WeaponTemplate struct {
	*WeaponTemplateVO
	weaponTagType            types.WeaponTagType                      //标识
	weaponType               types.WeaponType                         //类型
	battleAttrTemplate       *AttrTemplate                            //阶别属性
	awakenAttrTemplate       *AttrTemplate                            //觉醒属性
	needItemMap              map[playertypes.RoleType]map[int32]int32 //激活需要物品
	awakenItemMap            map[int32]int32                          //觉醒需要物品
	weaponUpstarTemplateMap  map[int32]*WeaponUpstarTemplate          //兵魂升星map
	weaponUpstarTemplate     *WeaponUpstarTemplate                    //兵魂升星
	weaponPeiYangTemplateMap map[int32]*WeaponPeiYangTemplate         //兵魂培养map
	weaponPeiYangTemplate    *WeaponPeiYangTemplate                   //兵魂培养
}

func (wt *WeaponTemplate) TemplateId() int {
	return wt.Id
}

func (wt *WeaponTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return wt.battleAttrTemplate
}

func (wt *WeaponTemplate) GetNeedItemMap(roleType playertypes.RoleType) map[int32]int32 {
	return wt.needItemMap[roleType]
}

func (wt *WeaponTemplate) GetName(roleType playertypes.RoleType) string {
	switch roleType {
	case playertypes.RoleTypeKaiTian:
		return wt.Name1
	case playertypes.RoleTypeYiJian:
		return wt.Name2
	case playertypes.RoleTypePoYue:
		return wt.Name3
	}

	return ""
}

func (wt *WeaponTemplate) GetAwakenItemMap() map[int32]int32 {
	return wt.awakenItemMap
}

func (wt *WeaponTemplate) GetAwakenAttrTemplate() *AttrTemplate {
	return wt.awakenAttrTemplate
}

func (wt *WeaponTemplate) GetWeaponUpstarByLevel(level int32) *WeaponUpstarTemplate {
	if v, ok := wt.weaponUpstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (wt *WeaponTemplate) GetWeaponPeiYangByLevel(level int32) *WeaponPeiYangTemplate {
	if v, ok := wt.weaponPeiYangTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (wt *WeaponTemplate) GetWeaponTag() types.WeaponTagType {
	return wt.weaponTagType
}

//兵魂吃培养丹升级
func (wt *WeaponTemplate) GetWeaponEatPeiYangTemplate(curLevel int32, num int32) (weaponPeiYangTemplate *WeaponPeiYangTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		weaponPeiYangTemplate, flag = wt.weaponPeiYangTemplateMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= weaponPeiYangTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

func (wt *WeaponTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wt.FileName(), wt.TemplateId(), err)
			return
		}
	}()

	//验证 tag
	wt.weaponTagType = types.WeaponTagType(wt.Tag)
	if !wt.weaponTagType.Valid() {
		err = fmt.Errorf("[%d] invalid", wt.Tag)
		return template.NewTemplateFieldError("Tag", err)
	}

	//验证 类型
	wt.weaponType = types.WeaponType(wt.Type)
	if !wt.weaponType.Valid() {
		err = fmt.Errorf("[%d] invalid", wt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	if wt.AttrId != 0 {
		//阶别 attr属性
		to := template.GetTemplateService().Get(int(wt.AttrId), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.AttrId)
			return template.NewTemplateFieldError("AttrId", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		wt.battleAttrTemplate = attrTemplate
	}

	err = validator.MinValidate(float64(wt.EatDan), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.EatDan)
		return template.NewTemplateFieldError("EatDan", err)
	}

	itemStrList := make([]string, 0, 3)
	countStrList := make([]string, 0, 3)
	itemStrList = append(itemStrList, wt.NeedItemId, wt.NeedItemId2, wt.NeedItemId3)
	countStrList = append(countStrList, wt.NeedItemCount, wt.NeedItemCount2, wt.NeedItemCount3)

	if len(itemStrList) != len(countStrList) {
		return template.NewTemplateFieldError("NeedItemId", err)
	}
	wt.needItemMap = make(map[playertypes.RoleType]map[int32]int32)
	for i := int(0); i < len(itemStrList); i++ {
		needItemId := itemStrList[i]
		if needItemId != "" {
			itemArr, err := utils.SplitAsIntArray(needItemId)
			if err != nil {
				return err
			}
			if len(itemArr) <= 0 {
				return template.NewTemplateFieldError("NeedItemId", err)
			}
			numArr, err := utils.SplitAsIntArray(wt.NeedItemCount)
			if err != nil {
				return err
			}
			if len(itemArr) != len(numArr) {
				return template.NewTemplateFieldError("NeedItemId", err)
			}

			switch i {
			case 0:
				{
					kaiTianNeedMap, exist := wt.needItemMap[playertypes.RoleTypeKaiTian]
					if !exist {
						kaiTianNeedMap = make(map[int32]int32)
						wt.needItemMap[playertypes.RoleTypeKaiTian] = kaiTianNeedMap
					}
					for j := 0; j < len(itemArr); j++ {
						kaiTianNeedMap[itemArr[j]] = numArr[j]
					}
					break
				}
			case 1:
				{
					yiJianNeedMap, exist := wt.needItemMap[playertypes.RoleTypeYiJian]
					if !exist {
						yiJianNeedMap = make(map[int32]int32)
						wt.needItemMap[playertypes.RoleTypeYiJian] = yiJianNeedMap
					}
					for j := 0; j < len(itemArr); j++ {
						yiJianNeedMap[itemArr[j]] = numArr[j]
					}
					break
				}
			case 2:
				{
					poYueNeedMap, exist := wt.needItemMap[playertypes.RoleTypePoYue]
					if !exist {
						poYueNeedMap = make(map[int32]int32)
						wt.needItemMap[playertypes.RoleTypePoYue] = poYueNeedMap
					}
					for j := 0; j < len(itemArr); j++ {
						poYueNeedMap[itemArr[j]] = numArr[j]
					}
					break
				}
			}
		}

	}

	//验证 is_awaken
	wt.awakenItemMap = make(map[int32]int32)
	awakenType := types.WeaponAwakenType(wt.IsAwaken)
	if !awakenType.Valid() {
		err = fmt.Errorf("[%d] invalid", wt.IsAwaken)
		return template.NewTemplateFieldError("IsAwaken", err)
	}

	if wt.IsAwaken != 0 {
		to := template.GetTemplateService().Get(int(wt.AwakenItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.AwakenItemId)
			return template.NewTemplateFieldError("AwakenItemId", err)
		}
		err = validator.MinValidate(float64(wt.AwakenItemNum), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", wt.AwakenItemNum)
			return template.NewTemplateFieldError("AwakenItemNum", err)
		}
		wt.awakenItemMap[wt.AwakenItemId] = wt.AwakenItemNum

		//验证 awaken_attr
		to = template.GetTemplateService().Get(int(wt.AwakenAttr), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.AwakenAttr)
			return template.NewTemplateFieldError("AwakenAttr", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		wt.awakenAttrTemplate = attrTemplate

		//验证 awaken_success_rate
		err = validator.RangeValidate(float64(wt.AwakenSuccessRate), float64(0), true, float64(common.MAX_RATE), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", wt.AwakenSuccessRate)
			err = template.NewTemplateFieldError("AwakenSuccessRate", err)
			return
		}

	}

	//验证 weapon_upgrade_begin_id
	if wt.WeaponUpgradeBeginId != 0 {
		to := template.GetTemplateService().Get(int(wt.WeaponUpgradeBeginId), (*WeaponUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.WeaponUpgradeBeginId)
			return template.NewTemplateFieldError("WeaponUpgradeBeginId", err)
		}

		weaponUpstarTemplate, ok := to.(*WeaponUpstarTemplate)
		if !ok {
			return fmt.Errorf("WeaponUpgradeBeginId [%d] invalid", wt.WeaponUpgradeBeginId)
		}
		if weaponUpstarTemplate.Level != 1 {
			return fmt.Errorf("WeaponUpgradeBeginId Level [%d] invalid", weaponUpstarTemplate.Level)
		}
		wt.weaponUpstarTemplate = weaponUpstarTemplate
	}

	//验证 weapon_peiyang_begin_id
	if wt.WeaponPeiYangBeginId != 0 {
		to := template.GetTemplateService().Get(int(wt.WeaponPeiYangBeginId), (*WeaponPeiYangTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wt.WeaponPeiYangBeginId)
			return template.NewTemplateFieldError("WeaponPeiYangBeginId", err)
		}

		weaponPeiYangTemplate, ok := to.(*WeaponPeiYangTemplate)
		if !ok {
			return fmt.Errorf("WeaponPeiYangBeginId [%d] invalid", wt.WeaponPeiYangBeginId)
		}
		wt.weaponPeiYangTemplate = weaponPeiYangTemplate
	}

	return nil
}

func (wt *WeaponTemplate) PatchAfterCheck() {
	if wt.weaponUpstarTemplate != nil {
		wt.weaponUpstarTemplateMap = make(map[int32]*WeaponUpstarTemplate)
		//赋值weaponUpstarTemplateMap
		for tempTemplate := wt.weaponUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextWeaponUpstarTemplate {
			level := tempTemplate.Level
			wt.weaponUpstarTemplateMap[level] = tempTemplate
		}
	}

	if wt.weaponPeiYangTemplate != nil {
		wt.weaponPeiYangTemplateMap = make(map[int32]*WeaponPeiYangTemplate)
		//赋值weaponPeiYangTemplateMap
		for tempTempalte := wt.weaponPeiYangTemplate; tempTempalte != nil; tempTempalte = tempTempalte.nextWeaponPeiYangTemplate {
			level := tempTempalte.Level
			wt.weaponPeiYangTemplateMap[level] = tempTempalte
		}
	}
}

func (wt *WeaponTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wt.FileName(), wt.TemplateId(), err)
			return
		}
	}()

	//验证 awaken_need_star
	err = validator.MinValidate(float64(wt.NeedStar), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wt.NeedStar)
		return template.NewTemplateFieldError("awaken_need_star", err)
	}

	for _, roleNeddItemMap := range wt.needItemMap {
		for item, num := range roleNeddItemMap {
			to := template.GetTemplateService().Get(int(item), (*ItemTemplate)(nil))
			if to == nil {
				err = fmt.Errorf("item[%d] invalid", item)
				return template.NewTemplateFieldError("NeedItemId", err)
			}

			// weaponTemplate := to.(*ItemTemplate)
			// if weaponTemplate.GetItemSubType() != itemtypes.ItemSoulSubTypeDebris {
			// 	err = fmt.Errorf("item[%d] invalid", item)
			// 	return template.NewTemplateFieldError("NeedItemId", err)
			// }

			err = validator.MinValidate(float64(num), float64(1), true)
			if err != nil {
				err = fmt.Errorf("num [%d] invalid", num)
				return template.NewTemplateFieldError("activation_need_item_count", err)
			}
		}
	}

	return nil
}

func (wt *WeaponTemplate) FileName() string {
	return "tb_weapon.json"
}

func init() {
	template.Register((*WeaponTemplate)(nil))
}
