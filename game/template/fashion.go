package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/fashion/types"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

//时装配置
type FashionTemplate struct {
	*FashionTemplateVO
	fashionType              types.FashionType                                                //类型
	battleAttrTemplate       *AttrTemplate                                                    //阶别属性
	needItemMap              map[playertypes.RoleType]map[playertypes.SexType]map[int32]int32 //激活需要物品
	fashionUpstarTemplateMap map[int32]*FashionUpstarTemplate                                 //时装升星map
	fashionUpstarTemplate    *FashionUpstarTemplate                                           //时装升星
	fashionNameMap           map[playertypes.RoleType]map[playertypes.SexType]string          //时装名称
}

//永久性
func (ft *FashionTemplate) Permanent() bool {
	return ft.Time == 0
}

func (ft *FashionTemplate) TemplateId() int {
	return ft.Id
}

func (ft *FashionTemplate) GetFashionName(role playertypes.RoleType, sex playertypes.SexType) string {
	nameMap, ok := ft.fashionNameMap[role]
	if !ok {
		return ""
	}
	name, ok := nameMap[sex]
	if !ok {
		return ""
	}

	return name
}

func (ft *FashionTemplate) GetFashionType() types.FashionType {
	return ft.fashionType
}

func (ft *FashionTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return ft.battleAttrTemplate
}

func (ft *FashionTemplate) GetNeedItemMap(roleType playertypes.RoleType, sexType playertypes.SexType) map[int32]int32 {
	needItemSexMap, exist := ft.needItemMap[roleType]
	if !exist {
		return nil
	}
	return needItemSexMap[sexType]
}

func (ft *FashionTemplate) GetFashionUpstarByLevel(level int32) *FashionUpstarTemplate {
	if v, ok := ft.fashionUpstarTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (ft *FashionTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(ft.FileName(), ft.TemplateId(), err)
			return
		}
	}()

	//验证 Type
	ft.fashionType = types.FashionType(ft.Type)
	if !ft.fashionType.Valid() {
		err = fmt.Errorf("[%d] invalid", ft.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//验证 Attr
	if ft.Attr != 0 {
		tempAttrTemplate := template.GetTemplateService().Get(int(ft.Attr), (*AttrTemplate)(nil))
		if tempAttrTemplate == nil {
			err = fmt.Errorf("[%d] invalid", ft.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := tempAttrTemplate.(*AttrTemplate)
		ft.battleAttrTemplate = attrTemplate
	}

	ft.needItemMap = make(map[playertypes.RoleType]map[playertypes.SexType]map[int32]int32)
	//验证 NeedItemId(开天男)
	if ft.NeedItemId != 0 {
		needItemTemplateVO := template.GetTemplateService().Get(int(ft.NeedItemId), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", ft.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(ft.NeedItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
		kaiTianNeedItemMap, exist := ft.needItemMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianNeedItemMap = make(map[playertypes.SexType]map[int32]int32)
			ft.needItemMap[playertypes.RoleTypeKaiTian] = kaiTianNeedItemMap
		}
		needItemMap, exist := kaiTianNeedItemMap[playertypes.SexTypeMan]
		if !exist {
			needItemMap = make(map[int32]int32)
			kaiTianNeedItemMap[playertypes.SexTypeMan] = needItemMap
		}
		needItemMap[ft.NeedItemId] = ft.NeedItemCount
	}

	//验证 NeedItemIdNv(开天女)
	if ft.NeedItemIdNv != 0 {
		needItemTemplateVO := template.GetTemplateService().Get(int(ft.NeedItemIdNv), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", ft.NeedItemIdNv)
			err = template.NewTemplateFieldError("NeedItemIdNv", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(ft.NeedItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
		kaiTianNeedItemMap, exist := ft.needItemMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianNeedItemMap = make(map[playertypes.SexType]map[int32]int32)
			ft.needItemMap[playertypes.RoleTypeKaiTian] = kaiTianNeedItemMap
		}
		needItemMap, exist := kaiTianNeedItemMap[playertypes.SexTypeWoman]
		if !exist {
			needItemMap = make(map[int32]int32)
			kaiTianNeedItemMap[playertypes.SexTypeWoman] = needItemMap
		}
		needItemMap[ft.NeedItemIdNv] = ft.NeedItemCount
	}

	//验证 NeedItemId2(奕剑男)
	if ft.NeedItemId != 0 {
		needItemTemplateVO := template.GetTemplateService().Get(int(ft.NeedItemId2), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", ft.NeedItemId2)
			err = template.NewTemplateFieldError("NeedItemId2", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(ft.NeedItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
		yiJianNeedItemMap, exist := ft.needItemMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianNeedItemMap = make(map[playertypes.SexType]map[int32]int32)
			ft.needItemMap[playertypes.RoleTypeYiJian] = yiJianNeedItemMap
		}
		needItemMap, exist := yiJianNeedItemMap[playertypes.SexTypeMan]
		if !exist {
			needItemMap = make(map[int32]int32)
			yiJianNeedItemMap[playertypes.SexTypeMan] = needItemMap
		}
		needItemMap[ft.NeedItemId2] = ft.NeedItemCount
	}

	//验证 NeedItemId2Nv(奕剑女)
	if ft.NeedItemIdNv != 0 {
		needItemTemplateVO := template.GetTemplateService().Get(int(ft.NeedItemId2Nv), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", ft.NeedItemId2Nv)
			err = template.NewTemplateFieldError("NeedItemId2Nv", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(ft.NeedItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
		yiJianNeedItemMap, exist := ft.needItemMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianNeedItemMap = make(map[playertypes.SexType]map[int32]int32)
			ft.needItemMap[playertypes.RoleTypeYiJian] = yiJianNeedItemMap
		}
		needItemMap, exist := yiJianNeedItemMap[playertypes.SexTypeWoman]
		if !exist {
			needItemMap = make(map[int32]int32)
			yiJianNeedItemMap[playertypes.SexTypeWoman] = needItemMap
		}
		needItemMap[ft.NeedItemId2Nv] = ft.NeedItemCount
	}

	//验证 NeedItemId2(破月男)
	if ft.NeedItemId != 0 {
		needItemTemplateVO := template.GetTemplateService().Get(int(ft.NeedItemId3), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", ft.NeedItemId3)
			err = template.NewTemplateFieldError("NeedItemId3", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(ft.NeedItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
		poYueNeedItemMap, exist := ft.needItemMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueNeedItemMap = make(map[playertypes.SexType]map[int32]int32)
			ft.needItemMap[playertypes.RoleTypePoYue] = poYueNeedItemMap
		}
		needItemMap, exist := poYueNeedItemMap[playertypes.SexTypeMan]
		if !exist {
			needItemMap = make(map[int32]int32)
			poYueNeedItemMap[playertypes.SexTypeMan] = needItemMap
		}
		needItemMap[ft.NeedItemId3] = ft.NeedItemCount
	}

	//验证 NeedItemId2Nv(破月女)
	if ft.NeedItemIdNv != 0 {
		needItemTemplateVO := template.GetTemplateService().Get(int(ft.NeedItemId3Nv), (*ItemTemplate)(nil))
		if needItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", ft.NeedItemId3Nv)
			err = template.NewTemplateFieldError("NeedItemId3Nv", err)
			return
		}

		//验证 NeedItemCount
		err = validator.MinValidate(float64(ft.NeedItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
		poYueNeedItemMap, exist := ft.needItemMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueNeedItemMap = make(map[playertypes.SexType]map[int32]int32)
			ft.needItemMap[playertypes.RoleTypePoYue] = poYueNeedItemMap
		}
		needItemMap, exist := poYueNeedItemMap[playertypes.SexTypeWoman]
		if !exist {
			needItemMap = make(map[int32]int32)
			poYueNeedItemMap[playertypes.SexTypeWoman] = needItemMap
		}
		needItemMap[ft.NeedItemId3Nv] = ft.NeedItemCount
	}

	//验证 fashion_upgrade_begin_id
	if ft.FashionUpgradeBeginId != 0 {
		to := template.GetTemplateService().Get(int(ft.FashionUpgradeBeginId), (*FashionUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", ft.FashionUpgradeBeginId)
			return template.NewTemplateFieldError("FashionUpgradeBeginId", err)
		}

		fashionUpstarTemplate, ok := to.(*FashionUpstarTemplate)
		if !ok {
			return fmt.Errorf("FashionUpgradeBeginId [%d] invalid", ft.FashionUpgradeBeginId)
		}
		if fashionUpstarTemplate.Level != 1 {
			return fmt.Errorf("FashionUpgradeBeginId Level [%d] invalid", fashionUpstarTemplate.Level)
		}
		ft.fashionUpstarTemplate = fashionUpstarTemplate
	}

	//时装名称
	ft.fashionNameMap = make(map[playertypes.RoleType]map[playertypes.SexType]string)
	kaitianMap, ok := ft.fashionNameMap[playertypes.RoleTypeKaiTian]
	if !ok {
		kaitianMap = make(map[playertypes.SexType]string)
		ft.fashionNameMap[playertypes.RoleTypeKaiTian] = kaitianMap
	}
	kaitianMap[playertypes.SexTypeMan] = ft.Name11
	kaitianMap[playertypes.SexTypeWoman] = ft.Name12

	yijianMap, ok := ft.fashionNameMap[playertypes.RoleTypeYiJian]
	if !ok {
		yijianMap = make(map[playertypes.SexType]string)
		ft.fashionNameMap[playertypes.RoleTypeYiJian] = yijianMap
	}
	yijianMap[playertypes.SexTypeMan] = ft.Name21
	yijianMap[playertypes.SexTypeWoman] = ft.Name22

	poyueMap, ok := ft.fashionNameMap[playertypes.RoleTypePoYue]
	if !ok {
		poyueMap = make(map[playertypes.SexType]string)
		ft.fashionNameMap[playertypes.RoleTypePoYue] = poyueMap
	}
	poyueMap[playertypes.SexTypeMan] = ft.Name21
	poyueMap[playertypes.SexTypeWoman] = ft.Name22
	return nil
}

func (ft *FashionTemplate) PatchAfterCheck() {
	if ft.fashionUpstarTemplate != nil {
		ft.fashionUpstarTemplateMap = make(map[int32]*FashionUpstarTemplate)
		//赋值fashionUpstarTemplateMap
		for tempTemplate := ft.fashionUpstarTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextFashionUpstarTemplate {
			level := tempTemplate.Level
			ft.fashionUpstarTemplateMap[level] = tempTemplate
		}
	}
}

func (ft *FashionTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(ft.FileName(), ft.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (ft *FashionTemplate) FileName() string {
	return "tb_fashion.json"
}

func init() {
	template.Register((*FashionTemplate)(nil))
}
