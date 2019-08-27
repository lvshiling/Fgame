package template

import (
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
)

// 运营活动-打折礼包分组模板
type BargainShopGroupTemplate struct {
	zhuanShengMap     map[playertypes.RoleType]map[playertypes.SexType][]*gametemplate.BargainShopTemplate
	sexBargainMap     map[playertypes.SexType][]*gametemplate.BargainShopTemplate
	roleBargainMap    map[playertypes.RoleType][]*gametemplate.BargainShopTemplate
	commonBargainList []*gametemplate.BargainShopTemplate
}

func CreateBargainGroupTemplate() *BargainShopGroupTemplate {
	gt := &BargainShopGroupTemplate{}
	gt.zhuanShengMap = make(map[playertypes.RoleType]map[playertypes.SexType][]*gametemplate.BargainShopTemplate)
	gt.sexBargainMap = make(map[playertypes.SexType][]*gametemplate.BargainShopTemplate)
	gt.roleBargainMap = make(map[playertypes.RoleType][]*gametemplate.BargainShopTemplate)
	return gt
}

func (gt *BargainShopGroupTemplate) AddTemplate(temp *gametemplate.BargainShopTemplate) {
	if temp.Profession != 0 && temp.Gender == 0 {
		gt.roleBargainMap[temp.GetRole()] = append(gt.roleBargainMap[temp.GetRole()], temp)
		return
	}

	if temp.Profession == 0 && temp.Gender != 0 {
		gt.sexBargainMap[temp.GetSex()] = append(gt.sexBargainMap[temp.GetSex()], temp)
		return
	}

	if temp.Profession == 0 && temp.Gender == 0 {
		gt.commonBargainList = append(gt.commonBargainList, temp)
		return
	}

	sexMap, ok := gt.zhuanShengMap[temp.GetRole()]
	if !ok {
		sexMap = make(map[playertypes.SexType][]*gametemplate.BargainShopTemplate)
		gt.zhuanShengMap[temp.GetRole()] = sexMap
	}

	sexMap[temp.GetSex()] = append(sexMap[temp.GetSex()], temp)
}

func (group *BargainShopGroupTemplate) GetDiscountBargainTemplateByType(role playertypes.RoleType, sex playertypes.SexType, typ int32) *gametemplate.BargainShopTemplate {
	tempList := group.GetDiscountBargainTemplate(role, sex)
	for _, temp := range tempList {
		if temp.Type == typ {
			return temp
		}
	}

	return nil
}

func (group *BargainShopGroupTemplate) GetDiscountBargainTemplate(role playertypes.RoleType, sex playertypes.SexType) []*gametemplate.BargainShopTemplate {
	// 全部通用
	totalTempList := group.commonBargainList

	// 性别通用
	sexCommonList := group.roleBargainMap[role]
	if len(sexCommonList) > 0 {
		totalTempList = append(totalTempList, sexCommonList...)
	}

	// 角色通用
	roleCommonList := group.sexBargainMap[sex]
	if len(roleCommonList) > 0 {
		totalTempList = append(totalTempList, roleCommonList...)
	}

	// 全部限制
	sexMap, ok := group.zhuanShengMap[role]
	if !ok {
		return totalTempList
	}
	tempList, ok := sexMap[sex]
	if !ok {
		return totalTempList
	}

	totalTempList = append(totalTempList, tempList...)
	return totalTempList
}
