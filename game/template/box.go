package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	playertypes "fgame/fgame/game/player/types"
	treasureboxtypes "fgame/fgame/game/treasurebox/types"
	"fmt"
)

type BoxTemplate struct {
	*BoxTemplateVO
	boxType         treasureboxtypes.BoxType
	useResMap       map[playertypes.RoleType]map[int32]int32
	nextBoxTemplate *BoxTemplate
	dropIdList      map[playertypes.RoleType]map[playertypes.SexType][]int32
}

func (t *BoxTemplate) GetUseItemMap(roleType playertypes.RoleType, num int32) map[int32]int32 {
	if num < 1 {
		panic("num必须大于0")
	}
	newMap := make(map[int32]int32)

	useResMap := t.getUseItemMap(roleType)
	for id, count := range useResMap {
		newMap[id] = count * num
	}
	return newMap
}

func (t *BoxTemplate) getUseItemMap(roleType playertypes.RoleType) map[int32]int32 {
	useResRoleMap, exist := t.useResMap[roleType]
	if !exist {
		return nil
	}
	return useResRoleMap
}

func (t *BoxTemplate) GetDropIdList(roleType playertypes.RoleType, sexType playertypes.SexType) []int32 {
	dropRoleListMap, exist := t.dropIdList[roleType]
	if !exist {
		return nil
	}
	dropIdList, exist := dropRoleListMap[sexType]
	if !exist {
		return nil
	}
	return dropIdList
}

func (t *BoxTemplate) GetBoxType() treasureboxtypes.BoxType {
	return t.boxType
}

func (t *BoxTemplate) TemplateId() int {
	return t.Id
}

func (t *BoxTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//开启宝箱消耗物品数量
	t.useResMap = make(map[playertypes.RoleType]map[int32]int32)
	//开天
	intUseItemIdArr, err := utils.SplitAsIntArray(t.UseItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}
	intUseItemCountArr, err := utils.SplitAsIntArray(t.UseItemAmount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemAmount)
		return template.NewTemplateFieldError("UseItemAmount", err)
	}
	if len(intUseItemIdArr) != len(intUseItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.UseItemId, t.UseItemAmount)
		err = template.NewTemplateFieldError("UseItemId or UseItemAmount", err)
		return err
	}
	if len(intUseItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intUseItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.UseItemId)
				return template.NewTemplateFieldError("UseItemId", err)
			}

			err = validator.MinValidate(float64(intUseItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("UseItemAmount", err)
			}

			kaiTianUseResMap, exist := t.useResMap[playertypes.RoleTypeKaiTian]
			if !exist {
				kaiTianUseResMap = make(map[int32]int32)
				t.useResMap[playertypes.RoleTypeKaiTian] = kaiTianUseResMap
			}

			kaiTianUseResMap[itemId] = intUseItemCountArr[index]
		}
	}

	//奕剑
	intUseItemId2Arr, err := utils.SplitAsIntArray(t.UseItemId2)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemId2)
		return template.NewTemplateFieldError("UseItemId2", err)
	}
	intUseItemCount2Arr, err := utils.SplitAsIntArray(t.UseItemAmount2)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemAmount2)
		return template.NewTemplateFieldError("UseItemAmount2", err)
	}
	if len(intUseItemId2Arr) != len(intUseItemCount2Arr) {
		err = fmt.Errorf("[%s][%s] invalid", t.UseItemId2, t.UseItemAmount2)
		err = template.NewTemplateFieldError("UseItemId2 or UseItemAmount2", err)
		return err
	}
	if len(intUseItemId2Arr) > 0 {
		//组合数据
		for index, itemId := range intUseItemId2Arr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.UseItemId2)
				return template.NewTemplateFieldError("UseItemId2", err)
			}

			err = validator.MinValidate(float64(intUseItemCount2Arr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("UseItemAmount2", err)
			}

			yiJianUseResMap, exist := t.useResMap[playertypes.RoleTypeYiJian]
			if !exist {
				yiJianUseResMap = make(map[int32]int32)
				t.useResMap[playertypes.RoleTypeYiJian] = yiJianUseResMap
			}

			yiJianUseResMap[itemId] = intUseItemCount2Arr[index]
		}
	}

	//破月
	intUseItemId3Arr, err := utils.SplitAsIntArray(t.UseItemId3)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemId3)
		return template.NewTemplateFieldError("UseItemId3", err)
	}
	intUseItemCount3Arr, err := utils.SplitAsIntArray(t.UseItemAmount3)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemAmount3)
		return template.NewTemplateFieldError("UseItemAmount3", err)
	}
	if len(intUseItemId3Arr) != len(intUseItemCount3Arr) {
		err = fmt.Errorf("[%s][%s] invalid", t.UseItemId3, t.UseItemAmount3)
		err = template.NewTemplateFieldError("UseItemId3 or UseItemAmount3", err)
		return err
	}
	if len(intUseItemId3Arr) > 0 {
		//组合数据
		for index, itemId := range intUseItemId3Arr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.UseItemId3)
				return template.NewTemplateFieldError("UseItemId3", err)
			}

			err = validator.MinValidate(float64(intUseItemCount3Arr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("UseItemAmount3", err)
			}

			poYueUseResMap, exist := t.useResMap[playertypes.RoleTypePoYue]
			if !exist {
				poYueUseResMap = make(map[int32]int32)
				t.useResMap[playertypes.RoleTypePoYue] = poYueUseResMap
			}

			poYueUseResMap[itemId] = intUseItemCount3Arr[index]
		}
	}

	//掉落id
	t.dropIdList = make(map[playertypes.RoleType]map[playertypes.SexType][]int32)
	//开天男
	if t.DropId != "" {
		intDropIdArr, err := utils.SplitAsIntArray(t.DropId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.DropId)
			return template.NewTemplateFieldError("DropId", err)
		}
		kaiTianDropMap, exist := t.dropIdList[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianDropMap = make(map[playertypes.SexType][]int32)
			t.dropIdList[playertypes.RoleTypeKaiTian] = kaiTianDropMap
		}
		kaiTianDropMap[playertypes.SexTypeMan] = intDropIdArr
	}

	//开天女
	if t.DropIdNv != "" {
		intDropIdArr, err := utils.SplitAsIntArray(t.DropIdNv)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.DropIdNv)
			return template.NewTemplateFieldError("DropIdNv", err)
		}
		kaiTianDropMap, exist := t.dropIdList[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianDropMap = make(map[playertypes.SexType][]int32)
			t.dropIdList[playertypes.RoleTypeKaiTian] = kaiTianDropMap
		}
		kaiTianDropMap[playertypes.SexTypeWoman] = intDropIdArr
	}

	//奕剑男
	if t.DropId2 != "" {
		intDropIdArr, err := utils.SplitAsIntArray(t.DropId2)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.DropId2)
			return template.NewTemplateFieldError("DropId2", err)
		}
		yiJianDropMap, exist := t.dropIdList[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianDropMap = make(map[playertypes.SexType][]int32)
			t.dropIdList[playertypes.RoleTypeYiJian] = yiJianDropMap
		}
		yiJianDropMap[playertypes.SexTypeMan] = intDropIdArr
	}

	//奕剑女
	if t.DropId2Nv != "" {
		intDropIdArr, err := utils.SplitAsIntArray(t.DropId2Nv)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.DropId2Nv)
			return template.NewTemplateFieldError("DropId2Nv", err)
		}
		yiJianDropMap, exist := t.dropIdList[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianDropMap = make(map[playertypes.SexType][]int32)
			t.dropIdList[playertypes.RoleTypeYiJian] = yiJianDropMap
		}
		yiJianDropMap[playertypes.SexTypeWoman] = intDropIdArr
	}

	//破月男
	if t.DropId3 != "" {
		intDropIdArr, err := utils.SplitAsIntArray(t.DropId3)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.DropId3)
			return template.NewTemplateFieldError("DropId3", err)
		}
		poYueDropMap, exist := t.dropIdList[playertypes.RoleTypePoYue]
		if !exist {
			poYueDropMap = make(map[playertypes.SexType][]int32)
			t.dropIdList[playertypes.RoleTypePoYue] = poYueDropMap
		}
		poYueDropMap[playertypes.SexTypeMan] = intDropIdArr
	}

	//破月女
	if t.DropId3Nv != "" {
		intDropIdArr, err := utils.SplitAsIntArray(t.DropId3Nv)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.DropId3Nv)
			return template.NewTemplateFieldError("DropId3Nv", err)
		}
		poYueDropMap, exist := t.dropIdList[playertypes.RoleTypePoYue]
		if !exist {
			poYueDropMap = make(map[playertypes.SexType][]int32)
			t.dropIdList[playertypes.RoleTypePoYue] = poYueDropMap
		}
		poYueDropMap[playertypes.SexTypeWoman] = intDropIdArr
	}

	return nil
}

func (t *BoxTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//宝箱类型
	boxType := treasureboxtypes.BoxType(t.Type)
	if !boxType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	t.boxType = boxType

	//开启消耗银两
	err = validator.MinValidate(float64(t.UseSilver), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UseSilver", err)
	}
	//开启消耗元宝
	err = validator.MinValidate(float64(t.UseGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UseGold", err)
	}
	//开启消耗绑定元宝
	err = validator.MinValidate(float64(t.UseBindgold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UseBindgold", err)
	}
	//自由选取的种类数量
	err = validator.MinValidate(float64(t.FixationItemNum), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("FixationItemNum", err)
	}

	//最小转数
	err = validator.MinValidate(float64(t.ZhuanshuMin), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("ZhuanshuMin", err)
	}

	//最小等级
	err = validator.MinValidate(float64(t.LevelMin), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("LevelMin", err)
	}

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*BoxTemplate)(nil))
		if to != nil {
			t.nextBoxTemplate = to.(*BoxTemplate)
		}
	}

	return nil
}

func (t *BoxTemplate) PatchAfterCheck() {

}

func (t *BoxTemplate) FileName() string {
	return "tb_box.json"
}

func init() {
	template.Register((*BoxTemplate)(nil))
}
