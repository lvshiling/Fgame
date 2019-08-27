package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

//宝库积分商店配置
type BaoKuJiFenTemplate struct {
	*BaoKuJiFenTemplateVO
	baoKuType  equipbaokutypes.BaoKuType
	itemIdMMap map[playertypes.RoleType]map[playertypes.SexType]int32
}

func (st *BaoKuJiFenTemplate) TemplateId() int {
	return st.Id
}

func (st *BaoKuJiFenTemplate) GetItemIdByRoleAndSex(role playertypes.RoleType, sex playertypes.SexType) int32 {
	itemIdMap, exist := st.itemIdMMap[role]
	if !exist {
		return 0
	}
	itemId, exist := itemIdMap[sex]
	if !exist {
		return 0
	}
	return itemId
}

func (st *BaoKuJiFenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	//区分职业性别道具id
	st.itemIdMMap = make(map[playertypes.RoleType]map[playertypes.SexType]int32)
	//开天
	kaiTianMap, exist := st.itemIdMMap[playertypes.RoleTypeKaiTian]
	if !exist {
		kaiTianMap = make(map[playertypes.SexType]int32)
		st.itemIdMMap[playertypes.RoleTypeKaiTian] = kaiTianMap
	}
	kaiTianMap[playertypes.SexTypeMan] = st.ItemIdKaiTianNan
	kaiTianMap[playertypes.SexTypeWoman] = st.ItemIdKaiTianNv
	//弈剑
	yiJianMap, exist := st.itemIdMMap[playertypes.RoleTypeYiJian]
	if !exist {
		yiJianMap = make(map[playertypes.SexType]int32)
		st.itemIdMMap[playertypes.RoleTypeYiJian] = yiJianMap
	}
	yiJianMap[playertypes.SexTypeMan] = st.ItemIdYiJianNan
	yiJianMap[playertypes.SexTypeWoman] = st.ItemIdYiJianNv
	//破月
	poYueMap, exist := st.itemIdMMap[playertypes.RoleTypePoYue]
	if !exist {
		poYueMap = make(map[playertypes.SexType]int32)
		st.itemIdMMap[playertypes.RoleTypePoYue] = poYueMap
	}
	poYueMap[playertypes.SexTypeMan] = st.ItemIdPoYueNan
	poYueMap[playertypes.SexTypeWoman] = st.ItemIdPoYueNv

	return nil
}

func (st *BaoKuJiFenTemplate) PatchAfterCheck() {

}

func (st *BaoKuJiFenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	//验证 Type
	typ := equipbaokutypes.BaoKuType(st.Type)
	if !typ.Valid() {
		err = fmt.Errorf("[%d] invalid", st.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	st.baoKuType = typ

	//验证 buyCount
	err = validator.MinValidate(float64(st.BuyCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.BuyCount)
		err = template.NewTemplateFieldError("BuyCount", err)
		return
	}

	//验证 maxCount
	err = validator.MinValidate(float64(st.MaxCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.MaxCount)
		err = template.NewTemplateFieldError("MaxCount", err)
		return
	}

	//验证 ItemId
	itemTemp := template.GetTemplateService().Get(int(st.ItemIdKaiTianNan), (*ItemTemplate)(nil))
	if itemTemp == nil {
		err = fmt.Errorf("[%d] invalid", st.ItemIdKaiTianNan)
		err = template.NewTemplateFieldError("ItemIdKaiTianNan", err)
		return
	}

	itemTemp = template.GetTemplateService().Get(int(st.ItemIdKaiTianNv), (*ItemTemplate)(nil))
	if itemTemp == nil {
		err = fmt.Errorf("[%d] invalid", st.ItemIdKaiTianNv)
		err = template.NewTemplateFieldError("ItemIdKaiTianNv", err)
		return
	}

	itemTemp = template.GetTemplateService().Get(int(st.ItemIdYiJianNan), (*ItemTemplate)(nil))
	if itemTemp == nil {
		err = fmt.Errorf("[%d] invalid", st.ItemIdYiJianNan)
		err = template.NewTemplateFieldError("ItemIdYiJianNan", err)
		return
	}

	itemTemp = template.GetTemplateService().Get(int(st.ItemIdYiJianNv), (*ItemTemplate)(nil))
	if itemTemp == nil {
		err = fmt.Errorf("[%d] invalid", st.ItemIdYiJianNv)
		err = template.NewTemplateFieldError("ItemIdYiJianNv", err)
		return
	}

	itemTemp = template.GetTemplateService().Get(int(st.ItemIdPoYueNan), (*ItemTemplate)(nil))
	if itemTemp == nil {
		err = fmt.Errorf("[%d] invalid", st.ItemIdPoYueNan)
		err = template.NewTemplateFieldError("ItemIdPoYueNan", err)
		return
	}

	itemTemp = template.GetTemplateService().Get(int(st.ItemIdPoYueNv), (*ItemTemplate)(nil))
	if itemTemp == nil {
		err = fmt.Errorf("[%d] invalid", st.ItemIdPoYueNv)
		err = template.NewTemplateFieldError("ItemIdPoYueNv", err)
		return
	}

	//验证 UseJiFen
	err = validator.MinValidate(float64(st.UseJiFen), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.UseJiFen)
		err = template.NewTemplateFieldError("UseJiFen", err)
		return
	}

	return nil
}

func (st *BaoKuJiFenTemplate) FileName() string {
	return "tb_baoku_jifen.json"
}

func init() {
	template.Register((*BaoKuJiFenTemplate)(nil))
}
